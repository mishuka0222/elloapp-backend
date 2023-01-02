package core

import (
	"github.com/zeromicro/go-zero/core/jsonx"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/dialog/dialog"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/dialog/internal/dal/dataobject"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

// DialogGetAllDrafts
// dialog.getAllDrafts user_id:long = Vector<PeerWithDraftMessage>;
func (c *DialogCore) DialogGetAllDrafts(in *dialog.TLDialogGetAllDrafts) (*dialog.Vector_PeerWithDraftMessage, error) {
	// var doList []dataobject.DialogsDO
	rValues := &dialog.Vector_PeerWithDraftMessage{
		Datas: []*dialog.PeerWithDraftMessage{},
	}

	if _, err := c.svcCtx.Dao.DialogsDAO.SelectAllDraftsWithCB(c.ctx, in.UserId, func(i int, v *dataobject.DialogsDO) {
		if v.DraftMessageData == "" {
			return
		}
		draft := new(mtproto.DraftMessage)
		if err2 := jsonx.UnmarshalFromString(v.DraftMessageData, draft); err2 != nil {
			c.Logger.Errorf("invalid draft: %v", v)
			return
		}

		rValues.Datas = append(rValues.Datas,
			dialog.MakeTLUpdateDraftMessage(&dialog.PeerWithDraftMessage{
				Peer:  mtproto.MakePeer(v.PeerType, v.PeerId),
				Draft: draft,
			}).To_PeerWithDraftMessage())
	}); err != nil {
		c.Logger.Errorf("dialog.getAllDrafts - error: %v", err)
		return nil, err
	}

	return rValues, nil
}
