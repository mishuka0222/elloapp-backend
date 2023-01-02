package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
)

// AccountRegisterDevice
// account.registerDevice#ec86017a flags:# no_muted:flags.0?true token_type:int token:string app_sandbox:Bool secret:bytes other_uids:Vector<long> = Bool;
func (c *NotificationCore) AccountRegisterDevice(in *mtproto.TLAccountRegisterDevice) (*mtproto.Bool, error) {
	// TODO: not impl
	c.Logger.Errorf("account.registerDevice blocked, License key from https://elloapp.com required to unlock enterprise features.")

	return mtproto.BoolTrue, nil
}
