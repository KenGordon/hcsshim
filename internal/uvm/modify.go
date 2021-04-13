package uvm

import (
	"context"
	"fmt"

	"github.com/Microsoft/hcsshim/internal/log"
	"github.com/Microsoft/hcsshim/internal/requesttype"
	hcsschema "github.com/Microsoft/hcsshim/internal/schema2"
	"github.com/pkg/errors"
)

// Modify modifies the compute system by sending a request to HCS.
func (uvm *UtilityVM) modify(ctx context.Context, doc *hcsschema.ModifySettingRequest) (err error) {
	if doc.GuestRequest == nil || uvm.gc == nil {
		return uvm.hcsSystem.Modify(ctx, doc)
	}

	hostdoc := *doc
	hostdoc.GuestRequest = nil
	if doc.ResourcePath != "" && doc.RequestType == requesttype.Add {
		err = uvm.hcsSystem.Modify(ctx, &hostdoc)
		if err != nil {
			return fmt.Errorf("adding VM resources: %s", err)
		}
		defer func() {
			if err != nil {
				hostdoc.RequestType = requesttype.Remove
				rerr := uvm.hcsSystem.Modify(ctx, &hostdoc)
				if rerr != nil {
					log.G(ctx).WithError(rerr).Error("failed to roll back resource add")
				}
			}
		}()
	}
	err = uvm.gc.Modify(ctx, doc.GuestRequest)
	if err != nil {
		return fmt.Errorf("guest modify: %s", err)
	}
	if doc.ResourcePath != "" && doc.RequestType == requesttype.Remove {
		err = uvm.hcsSystem.Modify(ctx, &hostdoc)
		if err != nil {
			err = fmt.Errorf("removing VM resources: %s", err)
			log.G(ctx).WithError(err).Error("failed to remove host resources after successful guest request")
			return err
		}
	}
	return nil
}

func (uvm *UtilityVM) guestRequest(ctx context.Context, request interface{}) error {
	if err := uvm.gc.Modify(ctx, request); err != nil {
		return errors.Wrap(err, "guest modify")
	}
	return nil
}
