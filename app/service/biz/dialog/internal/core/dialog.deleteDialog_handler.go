package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/dialog/dialog"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

// DialogDeleteDialog
// dialog.deleteDialog user_id:long peer_type:int peer_id:long = Bool;
func (c *DialogCore) DialogDeleteDialog(in *dialog.TLDialogDeleteDialog) (*mtproto.Bool, error) {
	c.svcCtx.Dao.DialogsDAO.Delete(c.ctx, in.UserId, in.PeerType, in.PeerId)

	return mtproto.BoolTrue, nil
}
