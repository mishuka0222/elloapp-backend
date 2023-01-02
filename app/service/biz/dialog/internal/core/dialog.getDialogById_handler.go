package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/dialog/dialog"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

// DialogGetDialogById
// dialog.getDialogById user_id:long peer_type:int peer_id:long = DialogExt;
func (c *DialogCore) DialogGetDialogById(in *dialog.TLDialogGetDialogById) (*dialog.DialogExt, error) {
	dialogDO, err := c.svcCtx.Dao.SelectDialog(c.ctx, in.UserId, in.PeerType, in.PeerId)
	if err != nil {
		c.Logger.Errorf("dialog.getDialogById - error: %v", err)
		return nil, err
	} else if dialogDO == nil {
		c.Logger.Errorf("dialog.getDialogById - error: not found dialog (%s)", in.DebugString())
		return nil, mtproto.ErrPeerIdInvalid
	}

	return makeDialog(dialogDO), nil
}
