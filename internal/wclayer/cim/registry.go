package cim

import (
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"unsafe"

	"github.com/Microsoft/hcsshim/internal/cimfs"
	"github.com/Microsoft/hcsshim/internal/wclayer"
	"github.com/Microsoft/hcsshim/internal/winapi"
	"github.com/Microsoft/hcsshim/osversion"
	"github.com/pkg/errors"
	"golang.org/x/sys/windows"
)

// enableCimBoot Opens the SYSTEM registry hive at path `hivePath` in
// `layerPath` and updates it to include a CIMFS Start registry key. This prepares the uvm
// to boot from a cim file if requested. The registry changes required to actually make
// the uvm boot from a cim will be added in the uvm config since we don't want every
// single uvm started with this layer to attempt to boot from a cim. (Look at
// addBootFromCimRegistryChanges for details).
// This registry key needs to be available in the early boot phase and so including it in the
// uvm config doesn't work.
func enableCimBoot(layerPath, hivePath string) (err error) {
	dataZero := make([]byte, 4)
	dataOne := make([]byte, 4)
	binary.LittleEndian.PutUint32(dataOne, 1)
	dataFour := make([]byte, 4)
	binary.LittleEndian.PutUint32(dataFour, 4)

	// bootGUID := []byte("{b890454c-80de-4e98-a7ab-56b74b4fbd0c}")
	bootGUID, err := windows.UTF16FromString("{b890454c-80de-4e98-a7ab-56b74b4fbd0c}")
	if err != nil {
		return fmt.Errorf("failed to encode boot guid to utf16: %w", err)
	}

	overrideBootPath, err := windows.UTF16FromString("\\Windows\\")
	if err != nil {
		return fmt.Errorf("failed to encode override boot path to utf16: %w", err)
	}

	regChanges := []struct {
		keyPath   string
		valueName string
		valueType winapi.RegType
		data      *byte
		dataLen   uint32
	}{
		{"ControlSet001\\Control", "BootContainerGuid", winapi.REG_TYPE_SZ, (*byte)(unsafe.Pointer(&bootGUID[0])), 2 * uint32(len(bootGUID))},
		{"ControlSet001\\Services\\UnionFS", "Start", winapi.REG_TYPE_DWORD, &dataZero[0], uint32(len(dataZero))},
		{"ControlSet001\\Services\\wcifs", "Start", winapi.REG_TYPE_DWORD, &dataFour[0], uint32(len(dataZero))},
		// The bootmgr loads the uvm files from the cim and so uses the relative path `UtilityVM\\Files` inside the cim to access the uvm files. However, once the cim is mounted UnionFS will merge the correct directory (UtilityVM\\Files) of the cim with the scratch and then that point onwards we don't need to use the relative path. Below registry key tells the kernel that the boot path that was provided in BCD should now be overriden with this new path.
		{"Setup", "BootPathOverride", winapi.REG_TYPE_SZ, (*byte)(unsafe.Pointer(&overrideBootPath[0])), 2 * uint32(len(overrideBootPath))},
	}

	var storeHandle winapi.OrHKey
	if err = winapi.OrOpenHive(hivePath, &storeHandle); err != nil {
		return fmt.Errorf("failed to open registry store at %s: %s", hivePath, err)
	}

	for _, change := range regChanges {
		var changeKey winapi.OrHKey
		if err = winapi.OrCreateKey(storeHandle, change.keyPath, 0, 0, 0, &changeKey, nil); err != nil {
			return fmt.Errorf("failed to open reg key %s: %s", change.keyPath, err)
		}

		if err = winapi.OrSetValue(changeKey, change.valueName, uint32(change.valueType), change.data, change.dataLen); err != nil {
			return fmt.Errorf("failed to set value for regkey %s\\%s : %s", change.keyPath, change.valueName, err)
		}
	}

	// remove the existing file first
	if err := os.Remove(hivePath); err != nil {
		return fmt.Errorf("failed to remove existing registry %s: %s", hivePath, err)
	}

	if err = winapi.OrSaveHive(winapi.OrHKey(storeHandle), hivePath, uint32(osversion.Get().MajorVersion), uint32(osversion.Get().MinorVersion)); err != nil {
		return fmt.Errorf("error saving the registry store: %s", err)
	}

	// close hive irrespective of the errors
	if err := winapi.OrCloseHive(winapi.OrHKey(storeHandle)); err != nil {
		return fmt.Errorf("error closing registry store; %s", err)
	}
	return nil

}

// mergeWithParentLayerHives merges the delta hives of current layer with the base registry
// hives of its parent layer. This function reads the parent layer cim to fetch registry
// hives of the parent layer and reads the `layerPath\\Hives` directory to read the hives
// form the current layer. The merged hives are stored in the directory provided by
// `outputDir`
func mergeWithParentLayerHives(layerPath, parentLayerPath, outputDir string) error {
	// create a temp directory to store parent layer hive files
	tmpParentLayer, err := ioutil.TempDir("", "")
	if err != nil {
		return fmt.Errorf("failed to create temp dir: %s", tmpParentLayer)
	}
	defer os.RemoveAll(tmpParentLayer)

	parentCimPath := GetCimPathFromLayer(parentLayerPath)

	for _, hv := range hives {
		err := cimfs.FetchFileFromCim(parentCimPath, filepath.Join(wclayer.HivesPath, hv.base), filepath.Join(tmpParentLayer, hv.base))
		if err != nil {
			return err
		}
	}

	// merge hives
	for _, hv := range hives {
		err := mergeHive(filepath.Join(tmpParentLayer, hv.base), filepath.Join(layerPath, wclayer.HivesPath, hv.delta), filepath.Join(outputDir, hv.base))
		if err != nil {
			return err
		}
	}
	return nil
}

// mergeHive merges the hive located at parentHivePath with the hive located at deltaHivePath and stores
// the result into the file at mergedHivePath. If a file already exists at path `mergedHivePath` then it
// throws an error.
func mergeHive(parentHivePath, deltaHivePath, mergedHivePath string) (err error) {
	var baseHive, deltaHive, mergedHive winapi.OrHKey
	if err := winapi.OrOpenHive(parentHivePath, &baseHive); err != nil {
		return fmt.Errorf("failed to open base hive %s: %s", parentHivePath, err)
	}
	defer func() {
		err2 := winapi.OrCloseHive(baseHive)
		if err == nil {
			err = errors.Wrap(err2, "failed to close base hive")
		}
	}()
	if err := winapi.OrOpenHive(deltaHivePath, &deltaHive); err != nil {
		return fmt.Errorf("failed to open delta hive %s: %s", deltaHivePath, err)
	}
	defer func() {
		err2 := winapi.OrCloseHive(deltaHive)
		if err == nil {
			err = errors.Wrap(err2, "failed to close delta hive")
		}
	}()
	if err := winapi.OrMergeHives([]winapi.OrHKey{baseHive, deltaHive}, &mergedHive); err != nil {
		return fmt.Errorf("failed to merge hives: %s", err)
	}
	defer func() {
		err2 := winapi.OrCloseHive(mergedHive)
		if err == nil {
			err = errors.Wrap(err2, "failed to close merged hive")
		}
	}()
	if err := winapi.OrSaveHive(mergedHive, mergedHivePath, uint32(osversion.Get().MajorVersion), uint32(osversion.Get().MinorVersion)); err != nil {
		return fmt.Errorf("failed to save hive: %s", err)
	}
	return
}
