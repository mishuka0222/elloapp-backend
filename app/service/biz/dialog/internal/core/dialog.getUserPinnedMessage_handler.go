package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/dialog/dialog"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

// DialogGetUserPinnedMessage
// dialog.getUserPinnedMessage user_id:long peer_type:int peer_id:long = Int32;
func (c *DialogCore) DialogGetUserPinnedMessage(in *dialog.TLDialogGetUserPinnedMessage) (*mtproto.Int32, error) {
	// TODO: not impl
	c.Logger.Errorf("dialog.getUserPinnedMessage - error: method DialogGetUserPinnedMessage not impl")

	return nil, mtproto.ErrMethodNotImpl
}
