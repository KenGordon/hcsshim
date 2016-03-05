// MACHINE GENERATED BY 'go generate' COMMAND; DO NOT EDIT

package hcsshim

import "unsafe"
import "syscall"

var _ unsafe.Pointer

var (
	modole32     = syscall.NewLazyDLL("ole32.dll")
	modvmcompute = syscall.NewLazyDLL("vmcompute.dll")

	procCoTaskMemFree                              = modole32.NewProc("CoTaskMemFree")
	procActivateLayer                              = modvmcompute.NewProc("ActivateLayer")
	procCopyLayer                                  = modvmcompute.NewProc("CopyLayer")
	procCreateLayer                                = modvmcompute.NewProc("CreateLayer")
	procCreateSandboxLayer                         = modvmcompute.NewProc("CreateSandboxLayer")
	procDeactivateLayer                            = modvmcompute.NewProc("DeactivateLayer")
	procDestroyLayer                               = modvmcompute.NewProc("DestroyLayer")
	procExportLayer                                = modvmcompute.NewProc("ExportLayer")
	procGetLayerMountPath                          = modvmcompute.NewProc("GetLayerMountPath")
	procGetBaseImages                              = modvmcompute.NewProc("GetBaseImages")
	procImportLayer                                = modvmcompute.NewProc("ImportLayer")
	procLayerExists                                = modvmcompute.NewProc("LayerExists")
	procNameToGuid                                 = modvmcompute.NewProc("NameToGuid")
	procPrepareLayer                               = modvmcompute.NewProc("PrepareLayer")
	procUnprepareLayer                             = modvmcompute.NewProc("UnprepareLayer")
	procProcessBaseImage                           = modvmcompute.NewProc("ProcessBaseImage")
	procProcessUtilityImage                        = modvmcompute.NewProc("ProcessUtilityImage")
	procImportLayerBegin                           = modvmcompute.NewProc("ImportLayerBegin")
	procImportLayerNext                            = modvmcompute.NewProc("ImportLayerNext")
	procImportLayerWrite                           = modvmcompute.NewProc("ImportLayerWrite")
	procImportLayerEnd                             = modvmcompute.NewProc("ImportLayerEnd")
	procExportLayerBegin                           = modvmcompute.NewProc("ExportLayerBegin")
	procExportLayerNext                            = modvmcompute.NewProc("ExportLayerNext")
	procExportLayerRead                            = modvmcompute.NewProc("ExportLayerRead")
	procExportLayerEnd                             = modvmcompute.NewProc("ExportLayerEnd")
	procCreateComputeSystem                        = modvmcompute.NewProc("CreateComputeSystem")
	procCreateProcessWithStdHandlesInComputeSystem = modvmcompute.NewProc("CreateProcessWithStdHandlesInComputeSystem")
	procResizeConsoleInComputeSystem               = modvmcompute.NewProc("ResizeConsoleInComputeSystem")
	procShutdownComputeSystem                      = modvmcompute.NewProc("ShutdownComputeSystem")
	procStartComputeSystem                         = modvmcompute.NewProc("StartComputeSystem")
	procTerminateComputeSystem                     = modvmcompute.NewProc("TerminateComputeSystem")
	procTerminateProcessInComputeSystem            = modvmcompute.NewProc("TerminateProcessInComputeSystem")
	procWaitForProcessInComputeSystem              = modvmcompute.NewProc("WaitForProcessInComputeSystem")
	procHNSCall                                    = modvmcompute.NewProc("HNSCall")
)

func coTaskMemFree(buffer unsafe.Pointer) {
	syscall.Syscall(procCoTaskMemFree.Addr(), 1, uintptr(buffer), 0, 0)
	return
}

func activateLayer(info *driverInfo, id string) (hr error) {
	var _p0 *uint16
	_p0, hr = syscall.UTF16PtrFromString(id)
	if hr != nil {
		return
	}
	return _activateLayer(info, _p0)
}

func _activateLayer(info *driverInfo, id *uint16) (hr error) {
	if hr = procActivateLayer.Find(); hr != nil {
		return
	}
	r0, _, _ := syscall.Syscall(procActivateLayer.Addr(), 2, uintptr(unsafe.Pointer(info)), uintptr(unsafe.Pointer(id)), 0)
	if int32(r0) < 0 {
		hr = syscall.Errno(win32FromHresult(r0))
	}
	return
}

func copyLayer(info *driverInfo, srcId string, dstId string, descriptors []WC_LAYER_DESCRIPTOR) (hr error) {
	var _p0 *uint16
	_p0, hr = syscall.UTF16PtrFromString(srcId)
	if hr != nil {
		return
	}
	var _p1 *uint16
	_p1, hr = syscall.UTF16PtrFromString(dstId)
	if hr != nil {
		return
	}
	return _copyLayer(info, _p0, _p1, descriptors)
}

func _copyLayer(info *driverInfo, srcId *uint16, dstId *uint16, descriptors []WC_LAYER_DESCRIPTOR) (hr error) {
	var _p2 *WC_LAYER_DESCRIPTOR
	if len(descriptors) > 0 {
		_p2 = &descriptors[0]
	}
	if hr = procCopyLayer.Find(); hr != nil {
		return
	}
	r0, _, _ := syscall.Syscall6(procCopyLayer.Addr(), 5, uintptr(unsafe.Pointer(info)), uintptr(unsafe.Pointer(srcId)), uintptr(unsafe.Pointer(dstId)), uintptr(unsafe.Pointer(_p2)), uintptr(len(descriptors)), 0)
	if int32(r0) < 0 {
		hr = syscall.Errno(win32FromHresult(r0))
	}
	return
}

func createLayer(info *driverInfo, id string, parent string) (hr error) {
	var _p0 *uint16
	_p0, hr = syscall.UTF16PtrFromString(id)
	if hr != nil {
		return
	}
	var _p1 *uint16
	_p1, hr = syscall.UTF16PtrFromString(parent)
	if hr != nil {
		return
	}
	return _createLayer(info, _p0, _p1)
}

func _createLayer(info *driverInfo, id *uint16, parent *uint16) (hr error) {
	if hr = procCreateLayer.Find(); hr != nil {
		return
	}
	r0, _, _ := syscall.Syscall(procCreateLayer.Addr(), 3, uintptr(unsafe.Pointer(info)), uintptr(unsafe.Pointer(id)), uintptr(unsafe.Pointer(parent)))
	if int32(r0) < 0 {
		hr = syscall.Errno(win32FromHresult(r0))
	}
	return
}

func createSandboxLayer(info *driverInfo, id string, parent string, descriptors []WC_LAYER_DESCRIPTOR) (hr error) {
	var _p0 *uint16
	_p0, hr = syscall.UTF16PtrFromString(id)
	if hr != nil {
		return
	}
	var _p1 *uint16
	_p1, hr = syscall.UTF16PtrFromString(parent)
	if hr != nil {
		return
	}
	return _createSandboxLayer(info, _p0, _p1, descriptors)
}

func _createSandboxLayer(info *driverInfo, id *uint16, parent *uint16, descriptors []WC_LAYER_DESCRIPTOR) (hr error) {
	var _p2 *WC_LAYER_DESCRIPTOR
	if len(descriptors) > 0 {
		_p2 = &descriptors[0]
	}
	if hr = procCreateSandboxLayer.Find(); hr != nil {
		return
	}
	r0, _, _ := syscall.Syscall6(procCreateSandboxLayer.Addr(), 5, uintptr(unsafe.Pointer(info)), uintptr(unsafe.Pointer(id)), uintptr(unsafe.Pointer(parent)), uintptr(unsafe.Pointer(_p2)), uintptr(len(descriptors)), 0)
	if int32(r0) < 0 {
		hr = syscall.Errno(win32FromHresult(r0))
	}
	return
}

func deactivateLayer(info *driverInfo, id string) (hr error) {
	var _p0 *uint16
	_p0, hr = syscall.UTF16PtrFromString(id)
	if hr != nil {
		return
	}
	return _deactivateLayer(info, _p0)
}

func _deactivateLayer(info *driverInfo, id *uint16) (hr error) {
	if hr = procDeactivateLayer.Find(); hr != nil {
		return
	}
	r0, _, _ := syscall.Syscall(procDeactivateLayer.Addr(), 2, uintptr(unsafe.Pointer(info)), uintptr(unsafe.Pointer(id)), 0)
	if int32(r0) < 0 {
		hr = syscall.Errno(win32FromHresult(r0))
	}
	return
}

func destroyLayer(info *driverInfo, id string) (hr error) {
	var _p0 *uint16
	_p0, hr = syscall.UTF16PtrFromString(id)
	if hr != nil {
		return
	}
	return _destroyLayer(info, _p0)
}

func _destroyLayer(info *driverInfo, id *uint16) (hr error) {
	if hr = procDestroyLayer.Find(); hr != nil {
		return
	}
	r0, _, _ := syscall.Syscall(procDestroyLayer.Addr(), 2, uintptr(unsafe.Pointer(info)), uintptr(unsafe.Pointer(id)), 0)
	if int32(r0) < 0 {
		hr = syscall.Errno(win32FromHresult(r0))
	}
	return
}

func exportLayer(info *driverInfo, id string, path string, descriptors []WC_LAYER_DESCRIPTOR) (hr error) {
	var _p0 *uint16
	_p0, hr = syscall.UTF16PtrFromString(id)
	if hr != nil {
		return
	}
	var _p1 *uint16
	_p1, hr = syscall.UTF16PtrFromString(path)
	if hr != nil {
		return
	}
	return _exportLayer(info, _p0, _p1, descriptors)
}

func _exportLayer(info *driverInfo, id *uint16, path *uint16, descriptors []WC_LAYER_DESCRIPTOR) (hr error) {
	var _p2 *WC_LAYER_DESCRIPTOR
	if len(descriptors) > 0 {
		_p2 = &descriptors[0]
	}
	if hr = procExportLayer.Find(); hr != nil {
		return
	}
	r0, _, _ := syscall.Syscall6(procExportLayer.Addr(), 5, uintptr(unsafe.Pointer(info)), uintptr(unsafe.Pointer(id)), uintptr(unsafe.Pointer(path)), uintptr(unsafe.Pointer(_p2)), uintptr(len(descriptors)), 0)
	if int32(r0) < 0 {
		hr = syscall.Errno(win32FromHresult(r0))
	}
	return
}

func getLayerMountPath(info *driverInfo, id string, length *uintptr, buffer *uint16) (hr error) {
	var _p0 *uint16
	_p0, hr = syscall.UTF16PtrFromString(id)
	if hr != nil {
		return
	}
	return _getLayerMountPath(info, _p0, length, buffer)
}

func _getLayerMountPath(info *driverInfo, id *uint16, length *uintptr, buffer *uint16) (hr error) {
	if hr = procGetLayerMountPath.Find(); hr != nil {
		return
	}
	r0, _, _ := syscall.Syscall6(procGetLayerMountPath.Addr(), 4, uintptr(unsafe.Pointer(info)), uintptr(unsafe.Pointer(id)), uintptr(unsafe.Pointer(length)), uintptr(unsafe.Pointer(buffer)), 0, 0)
	if int32(r0) < 0 {
		hr = syscall.Errno(win32FromHresult(r0))
	}
	return
}

func getBaseImages(buffer **uint16) (hr error) {
	if hr = procGetBaseImages.Find(); hr != nil {
		return
	}
	r0, _, _ := syscall.Syscall(procGetBaseImages.Addr(), 1, uintptr(unsafe.Pointer(buffer)), 0, 0)
	if int32(r0) < 0 {
		hr = syscall.Errno(win32FromHresult(r0))
	}
	return
}

func importLayer(info *driverInfo, id string, path string, descriptors []WC_LAYER_DESCRIPTOR) (hr error) {
	var _p0 *uint16
	_p0, hr = syscall.UTF16PtrFromString(id)
	if hr != nil {
		return
	}
	var _p1 *uint16
	_p1, hr = syscall.UTF16PtrFromString(path)
	if hr != nil {
		return
	}
	return _importLayer(info, _p0, _p1, descriptors)
}

func _importLayer(info *driverInfo, id *uint16, path *uint16, descriptors []WC_LAYER_DESCRIPTOR) (hr error) {
	var _p2 *WC_LAYER_DESCRIPTOR
	if len(descriptors) > 0 {
		_p2 = &descriptors[0]
	}
	if hr = procImportLayer.Find(); hr != nil {
		return
	}
	r0, _, _ := syscall.Syscall6(procImportLayer.Addr(), 5, uintptr(unsafe.Pointer(info)), uintptr(unsafe.Pointer(id)), uintptr(unsafe.Pointer(path)), uintptr(unsafe.Pointer(_p2)), uintptr(len(descriptors)), 0)
	if int32(r0) < 0 {
		hr = syscall.Errno(win32FromHresult(r0))
	}
	return
}

func layerExists(info *driverInfo, id string, exists *uint32) (hr error) {
	var _p0 *uint16
	_p0, hr = syscall.UTF16PtrFromString(id)
	if hr != nil {
		return
	}
	return _layerExists(info, _p0, exists)
}

func _layerExists(info *driverInfo, id *uint16, exists *uint32) (hr error) {
	if hr = procLayerExists.Find(); hr != nil {
		return
	}
	r0, _, _ := syscall.Syscall(procLayerExists.Addr(), 3, uintptr(unsafe.Pointer(info)), uintptr(unsafe.Pointer(id)), uintptr(unsafe.Pointer(exists)))
	if int32(r0) < 0 {
		hr = syscall.Errno(win32FromHresult(r0))
	}
	return
}

func nameToGuid(name string, guid *GUID) (hr error) {
	var _p0 *uint16
	_p0, hr = syscall.UTF16PtrFromString(name)
	if hr != nil {
		return
	}
	return _nameToGuid(_p0, guid)
}

func _nameToGuid(name *uint16, guid *GUID) (hr error) {
	if hr = procNameToGuid.Find(); hr != nil {
		return
	}
	r0, _, _ := syscall.Syscall(procNameToGuid.Addr(), 2, uintptr(unsafe.Pointer(name)), uintptr(unsafe.Pointer(guid)), 0)
	if int32(r0) < 0 {
		hr = syscall.Errno(win32FromHresult(r0))
	}
	return
}

func prepareLayer(info *driverInfo, id string, descriptors []WC_LAYER_DESCRIPTOR) (hr error) {
	var _p0 *uint16
	_p0, hr = syscall.UTF16PtrFromString(id)
	if hr != nil {
		return
	}
	return _prepareLayer(info, _p0, descriptors)
}

func _prepareLayer(info *driverInfo, id *uint16, descriptors []WC_LAYER_DESCRIPTOR) (hr error) {
	var _p1 *WC_LAYER_DESCRIPTOR
	if len(descriptors) > 0 {
		_p1 = &descriptors[0]
	}
	if hr = procPrepareLayer.Find(); hr != nil {
		return
	}
	r0, _, _ := syscall.Syscall6(procPrepareLayer.Addr(), 4, uintptr(unsafe.Pointer(info)), uintptr(unsafe.Pointer(id)), uintptr(unsafe.Pointer(_p1)), uintptr(len(descriptors)), 0, 0)
	if int32(r0) < 0 {
		hr = syscall.Errno(win32FromHresult(r0))
	}
	return
}

func unprepareLayer(info *driverInfo, id string) (hr error) {
	var _p0 *uint16
	_p0, hr = syscall.UTF16PtrFromString(id)
	if hr != nil {
		return
	}
	return _unprepareLayer(info, _p0)
}

func _unprepareLayer(info *driverInfo, id *uint16) (hr error) {
	if hr = procUnprepareLayer.Find(); hr != nil {
		return
	}
	r0, _, _ := syscall.Syscall(procUnprepareLayer.Addr(), 2, uintptr(unsafe.Pointer(info)), uintptr(unsafe.Pointer(id)), 0)
	if int32(r0) < 0 {
		hr = syscall.Errno(win32FromHresult(r0))
	}
	return
}

func processBaseImage(path string) (hr error) {
	var _p0 *uint16
	_p0, hr = syscall.UTF16PtrFromString(path)
	if hr != nil {
		return
	}
	return _processBaseImage(_p0)
}

func _processBaseImage(path *uint16) (hr error) {
	if hr = procProcessBaseImage.Find(); hr != nil {
		return
	}
	r0, _, _ := syscall.Syscall(procProcessBaseImage.Addr(), 1, uintptr(unsafe.Pointer(path)), 0, 0)
	if int32(r0) < 0 {
		hr = syscall.Errno(win32FromHresult(r0))
	}
	return
}

func processUtilityImage(path string) (hr error) {
	var _p0 *uint16
	_p0, hr = syscall.UTF16PtrFromString(path)
	if hr != nil {
		return
	}
	return _processUtilityImage(_p0)
}

func _processUtilityImage(path *uint16) (hr error) {
	if hr = procProcessUtilityImage.Find(); hr != nil {
		return
	}
	r0, _, _ := syscall.Syscall(procProcessUtilityImage.Addr(), 1, uintptr(unsafe.Pointer(path)), 0, 0)
	if int32(r0) < 0 {
		hr = syscall.Errno(win32FromHresult(r0))
	}
	return
}

func importLayerBegin(info *driverInfo, id string, descriptors []WC_LAYER_DESCRIPTOR, context *uintptr) (hr error) {
	var _p0 *uint16
	_p0, hr = syscall.UTF16PtrFromString(id)
	if hr != nil {
		return
	}
	return _importLayerBegin(info, _p0, descriptors, context)
}

func _importLayerBegin(info *driverInfo, id *uint16, descriptors []WC_LAYER_DESCRIPTOR, context *uintptr) (hr error) {
	var _p1 *WC_LAYER_DESCRIPTOR
	if len(descriptors) > 0 {
		_p1 = &descriptors[0]
	}
	if hr = procImportLayerBegin.Find(); hr != nil {
		return
	}
	r0, _, _ := syscall.Syscall6(procImportLayerBegin.Addr(), 5, uintptr(unsafe.Pointer(info)), uintptr(unsafe.Pointer(id)), uintptr(unsafe.Pointer(_p1)), uintptr(len(descriptors)), uintptr(unsafe.Pointer(context)), 0)
	if int32(r0) < 0 {
		hr = syscall.Errno(win32FromHresult(r0))
	}
	return
}

func importLayerNext(context uintptr, fileName string, fileInfo *winio.FileBasicInfo) (hr error) {
	var _p0 *uint16
	_p0, hr = syscall.UTF16PtrFromString(fileName)
	if hr != nil {
		return
	}
	return _importLayerNext(context, _p0, fileInfo)
}

func _importLayerNext(context uintptr, fileName *uint16, fileInfo *winio.FileBasicInfo) (hr error) {
	if hr = procImportLayerNext.Find(); hr != nil {
		return
	}
	r0, _, _ := syscall.Syscall(procImportLayerNext.Addr(), 3, uintptr(context), uintptr(unsafe.Pointer(fileName)), uintptr(unsafe.Pointer(fileInfo)))
	if int32(r0) < 0 {
		hr = syscall.Errno(win32FromHresult(r0))
	}
	return
}

func importLayerWrite(context uintptr, buffer []byte) (hr error) {
	var _p0 *byte
	if len(buffer) > 0 {
		_p0 = &buffer[0]
	}
	if hr = procImportLayerWrite.Find(); hr != nil {
		return
	}
	r0, _, _ := syscall.Syscall(procImportLayerWrite.Addr(), 3, uintptr(context), uintptr(unsafe.Pointer(_p0)), uintptr(len(buffer)))
	if int32(r0) < 0 {
		hr = syscall.Errno(win32FromHresult(r0))
	}
	return
}

func importLayerEnd(context uintptr) (hr error) {
	if hr = procImportLayerEnd.Find(); hr != nil {
		return
	}
	r0, _, _ := syscall.Syscall(procImportLayerEnd.Addr(), 1, uintptr(context), 0, 0)
	if int32(r0) < 0 {
		hr = syscall.Errno(win32FromHresult(r0))
	}
	return
}

func exportLayerBegin(info *driverInfo, id string, descriptors []WC_LAYER_DESCRIPTOR, context *uintptr) (hr error) {
	var _p0 *uint16
	_p0, hr = syscall.UTF16PtrFromString(id)
	if hr != nil {
		return
	}
	return _exportLayerBegin(info, _p0, descriptors, context)
}

func _exportLayerBegin(info *driverInfo, id *uint16, descriptors []WC_LAYER_DESCRIPTOR, context *uintptr) (hr error) {
	var _p1 *WC_LAYER_DESCRIPTOR
	if len(descriptors) > 0 {
		_p1 = &descriptors[0]
	}
	if hr = procExportLayerBegin.Find(); hr != nil {
		return
	}
	r0, _, _ := syscall.Syscall6(procExportLayerBegin.Addr(), 5, uintptr(unsafe.Pointer(info)), uintptr(unsafe.Pointer(id)), uintptr(unsafe.Pointer(_p1)), uintptr(len(descriptors)), uintptr(unsafe.Pointer(context)), 0)
	if int32(r0) < 0 {
		hr = syscall.Errno(win32FromHresult(r0))
	}
	return
}

func exportLayerNext(context uintptr, fileName **uint16, fileInfo *winio.FileBasicInfo, fileSize *int64, deleted *uint32) (hr error) {
	if hr = procExportLayerNext.Find(); hr != nil {
		return
	}
	r0, _, _ := syscall.Syscall6(procExportLayerNext.Addr(), 5, uintptr(context), uintptr(unsafe.Pointer(fileName)), uintptr(unsafe.Pointer(fileInfo)), uintptr(unsafe.Pointer(fileSize)), uintptr(unsafe.Pointer(deleted)), 0)
	if int32(r0) < 0 {
		hr = syscall.Errno(win32FromHresult(r0))
	}
	return
}

func exportLayerRead(context uintptr, buffer []byte, bytesRead *uint32) (hr error) {
	var _p0 *byte
	if len(buffer) > 0 {
		_p0 = &buffer[0]
	}
	if hr = procExportLayerRead.Find(); hr != nil {
		return
	}
	r0, _, _ := syscall.Syscall6(procExportLayerRead.Addr(), 4, uintptr(context), uintptr(unsafe.Pointer(_p0)), uintptr(len(buffer)), uintptr(unsafe.Pointer(bytesRead)), 0, 0)
	if int32(r0) < 0 {
		hr = syscall.Errno(win32FromHresult(r0))
	}
	return
}

func exportLayerEnd(context uintptr) (hr error) {
	if hr = procExportLayerEnd.Find(); hr != nil {
		return
	}
	r0, _, _ := syscall.Syscall(procExportLayerEnd.Addr(), 1, uintptr(context), 0, 0)
	if int32(r0) < 0 {
		hr = syscall.Errno(win32FromHresult(r0))
	}
	return
}

func createComputeSystem(id string, configuration string) (hr error) {
	var _p0 *uint16
	_p0, hr = syscall.UTF16PtrFromString(id)
	if hr != nil {
		return
	}
	var _p1 *uint16
	_p1, hr = syscall.UTF16PtrFromString(configuration)
	if hr != nil {
		return
	}
	return _createComputeSystem(_p0, _p1)
}

func _createComputeSystem(id *uint16, configuration *uint16) (hr error) {
	if hr = procCreateComputeSystem.Find(); hr != nil {
		return
	}
	r0, _, _ := syscall.Syscall(procCreateComputeSystem.Addr(), 2, uintptr(unsafe.Pointer(id)), uintptr(unsafe.Pointer(configuration)), 0)
	if int32(r0) < 0 {
		hr = syscall.Errno(win32FromHresult(r0))
	}
	return
}

func createProcessWithStdHandlesInComputeSystem(id string, paramsJson string, pid *uint32, stdin *syscall.Handle, stdout *syscall.Handle, stderr *syscall.Handle) (hr error) {
	var _p0 *uint16
	_p0, hr = syscall.UTF16PtrFromString(id)
	if hr != nil {
		return
	}
	var _p1 *uint16
	_p1, hr = syscall.UTF16PtrFromString(paramsJson)
	if hr != nil {
		return
	}
	return _createProcessWithStdHandlesInComputeSystem(_p0, _p1, pid, stdin, stdout, stderr)
}

func _createProcessWithStdHandlesInComputeSystem(id *uint16, paramsJson *uint16, pid *uint32, stdin *syscall.Handle, stdout *syscall.Handle, stderr *syscall.Handle) (hr error) {
	if hr = procCreateProcessWithStdHandlesInComputeSystem.Find(); hr != nil {
		return
	}
	r0, _, _ := syscall.Syscall6(procCreateProcessWithStdHandlesInComputeSystem.Addr(), 6, uintptr(unsafe.Pointer(id)), uintptr(unsafe.Pointer(paramsJson)), uintptr(unsafe.Pointer(pid)), uintptr(unsafe.Pointer(stdin)), uintptr(unsafe.Pointer(stdout)), uintptr(unsafe.Pointer(stderr)))
	if int32(r0) < 0 {
		hr = syscall.Errno(win32FromHresult(r0))
	}
	return
}

func resizeConsoleInComputeSystem(id string, pid uint32, height uint16, width uint16, flags uint32) (hr error) {
	var _p0 *uint16
	_p0, hr = syscall.UTF16PtrFromString(id)
	if hr != nil {
		return
	}
	return _resizeConsoleInComputeSystem(_p0, pid, height, width, flags)
}

func _resizeConsoleInComputeSystem(id *uint16, pid uint32, height uint16, width uint16, flags uint32) (hr error) {
	if hr = procResizeConsoleInComputeSystem.Find(); hr != nil {
		return
	}
	r0, _, _ := syscall.Syscall6(procResizeConsoleInComputeSystem.Addr(), 5, uintptr(unsafe.Pointer(id)), uintptr(pid), uintptr(height), uintptr(width), uintptr(flags), 0)
	if int32(r0) < 0 {
		hr = syscall.Errno(win32FromHresult(r0))
	}
	return
}

func shutdownComputeSystem(id string, timeout uint32) (hr error) {
	var _p0 *uint16
	_p0, hr = syscall.UTF16PtrFromString(id)
	if hr != nil {
		return
	}
	return _shutdownComputeSystem(_p0, timeout)
}

func _shutdownComputeSystem(id *uint16, timeout uint32) (hr error) {
	if hr = procShutdownComputeSystem.Find(); hr != nil {
		return
	}
	r0, _, _ := syscall.Syscall(procShutdownComputeSystem.Addr(), 2, uintptr(unsafe.Pointer(id)), uintptr(timeout), 0)
	if int32(r0) < 0 {
		hr = syscall.Errno(win32FromHresult(r0))
	}
	return
}

func startComputeSystem(id string) (hr error) {
	var _p0 *uint16
	_p0, hr = syscall.UTF16PtrFromString(id)
	if hr != nil {
		return
	}
	return _startComputeSystem(_p0)
}

func _startComputeSystem(id *uint16) (hr error) {
	if hr = procStartComputeSystem.Find(); hr != nil {
		return
	}
	r0, _, _ := syscall.Syscall(procStartComputeSystem.Addr(), 1, uintptr(unsafe.Pointer(id)), 0, 0)
	if int32(r0) < 0 {
		hr = syscall.Errno(win32FromHresult(r0))
	}
	return
}

func terminateComputeSystem(id string) (hr error) {
	var _p0 *uint16
	_p0, hr = syscall.UTF16PtrFromString(id)
	if hr != nil {
		return
	}
	return _terminateComputeSystem(_p0)
}

func _terminateComputeSystem(id *uint16) (hr error) {
	if hr = procTerminateComputeSystem.Find(); hr != nil {
		return
	}
	r0, _, _ := syscall.Syscall(procTerminateComputeSystem.Addr(), 1, uintptr(unsafe.Pointer(id)), 0, 0)
	if int32(r0) < 0 {
		hr = syscall.Errno(win32FromHresult(r0))
	}
	return
}

func terminateProcessInComputeSystem(id string, pid uint32) (hr error) {
	var _p0 *uint16
	_p0, hr = syscall.UTF16PtrFromString(id)
	if hr != nil {
		return
	}
	return _terminateProcessInComputeSystem(_p0, pid)
}

func _terminateProcessInComputeSystem(id *uint16, pid uint32) (hr error) {
	if hr = procTerminateProcessInComputeSystem.Find(); hr != nil {
		return
	}
	r0, _, _ := syscall.Syscall(procTerminateProcessInComputeSystem.Addr(), 2, uintptr(unsafe.Pointer(id)), uintptr(pid), 0)
	if int32(r0) < 0 {
		hr = syscall.Errno(win32FromHresult(r0))
	}
	return
}

func waitForProcessInComputeSystem(id string, pid uint32, timeout uint32, exitCode *uint32) (hr error) {
	var _p0 *uint16
	_p0, hr = syscall.UTF16PtrFromString(id)
	if hr != nil {
		return
	}
	return _waitForProcessInComputeSystem(_p0, pid, timeout, exitCode)
}

func _waitForProcessInComputeSystem(id *uint16, pid uint32, timeout uint32, exitCode *uint32) (hr error) {
	if hr = procWaitForProcessInComputeSystem.Find(); hr != nil {
		return
	}
	r0, _, _ := syscall.Syscall6(procWaitForProcessInComputeSystem.Addr(), 4, uintptr(unsafe.Pointer(id)), uintptr(pid), uintptr(timeout), uintptr(unsafe.Pointer(exitCode)), 0, 0)
	if int32(r0) < 0 {
		hr = syscall.Errno(win32FromHresult(r0))
	}
	return
}

func _hnsCall(method string, path string, object string, response **uint16) (hr error) {
	var _p0 *uint16
	_p0, hr = syscall.UTF16PtrFromString(method)
	if hr != nil {
		return
	}
	var _p1 *uint16
	_p1, hr = syscall.UTF16PtrFromString(path)
	if hr != nil {
		return
	}
	var _p2 *uint16
	_p2, hr = syscall.UTF16PtrFromString(object)
	if hr != nil {
		return
	}
	return __hnsCall(_p0, _p1, _p2, response)
}

func __hnsCall(method *uint16, path *uint16, object *uint16, response **uint16) (hr error) {
	if hr = procHNSCall.Find(); hr != nil {
		return
	}
	r0, _, _ := syscall.Syscall6(procHNSCall.Addr(), 4, uintptr(unsafe.Pointer(method)), uintptr(unsafe.Pointer(path)), uintptr(unsafe.Pointer(object)), uintptr(unsafe.Pointer(response)), 0, 0)
	if int32(r0) < 0 {
		hr = syscall.Errno(win32FromHresult(r0))
	}
	return
}
