package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/dialog/dialog"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
)

// DialogGetDialogUnreadMarkList
// dialog.getDialogUnreadMarkList user_id:long = Vector<DialogPeer>;
func (c *DialogCore) DialogGetDialogUnreadMarkList(in *dialog.TLDialogGetDialogUnreadMarkList) (*dialog.Vector_DialogPeer, error) {
	// TODO: not impl
	c.Logger.Errorf("dialog.getDialogUnreadMarkList - error: method DialogGetDialogUnreadMarkList not impl")

	return nil, mtproto.ErrMethodNotImpl
}
