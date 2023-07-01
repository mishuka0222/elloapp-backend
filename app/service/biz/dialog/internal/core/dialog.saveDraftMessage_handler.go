package core

import (
	"github.com/zeromicro/go-zero/core/jsonx"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/dialog/dialog"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/pkg2/hack"
)

// DialogSaveDraftMessage
// dialog.saveDraftMessage user_id:long peer_type:int peer_id:long message:DraftMessage = Bool;
func (c *DialogCore) DialogSaveDraftMessage(in *dialog.TLDialogSaveDraftMessage) (*mtproto.Bool, error) {
	draft, _ := jsonx.Marshal(in.Message)

	_, err := c.svcCtx.Dao.DialogsDAO.SaveDraft(c.ctx,
		2,
		hack.String(draft),
		in.UserId,
		in.PeerType,
		in.PeerId)
	if err != nil {
		c.Logger.Errorf("dialog.saveDraftMessage - error: %v", err)
		return nil, err
	}

	return mtproto.BoolTrue, nil
}
