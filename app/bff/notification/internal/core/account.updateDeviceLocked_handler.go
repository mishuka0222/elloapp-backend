package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

// AccountUpdateDeviceLocked
// account.updateDeviceLocked#38df3532 period:int = Bool;
func (c *NotificationCore) AccountUpdateDeviceLocked(in *mtproto.TLAccountUpdateDeviceLocked) (*mtproto.Bool, error) {
	// TODO: not impl
	c.Logger.Errorf("account.updateDeviceLocked blocked, License key from https://elloapp.com required to unlock enterprise features.")

	return mtproto.BoolTrue, nil
}
