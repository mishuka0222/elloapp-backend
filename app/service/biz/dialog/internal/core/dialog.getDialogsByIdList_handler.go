package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/dialog/dialog"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/dialog/internal/dal/dataobject"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
)

// DialogGetDialogsByIdList
// dialog.getDialogsByIdList user_id:long id_list:Vector<long> = Vector<DialogExt>;
func (c *DialogCore) DialogGetDialogsByIdList(in *dialog.TLDialogGetDialogsByIdList) (*dialog.Vector_DialogExt, error) {
	var (
		dList dialog.DialogExtList
		meId  = in.GetUserId()
		// peerId := mtproto.MakePeerDialogId()
	)

	doList, _ := c.svcCtx.Dao.DialogsDAO.SelectPeerDialogList(
		c.ctx,
		meId,
		in.IdList)

	for _, id := range in.IdList {
		found := false
		for i := 0; i < len(doList); i++ {
			if doList[i].PeerDialogId == id {
				found = true
				dList = append(dList, makeDialog(&doList[i]))
				break
			}
		}
		if !found {
			peerType, peerId := mtproto.GetPeerUtilByPeerDialogId(id)
			dList = append(dList, makeDialog(&dataobject.DialogsDO{
				UserId:           in.UserId,
				PeerType:         peerType,
				PeerId:           peerId,
				PeerDialogId:     id,
				Pinned:           0,
				TopMessage:       0,
				PinnedMsgId:      0,
				ReadInboxMaxId:   0,
				ReadOutboxMaxId:  0,
				UnreadCount:      0,
				UnreadMark:       false,
				DraftType:        0,
				DraftMessageData: "null",
				FolderId:         0,
				FolderPinned:     0,
				HasScheduled:     false,
				Date2:            0,
			}))
		}
	}

	return &dialog.Vector_DialogExt{
		Datas: dList,
	}, nil
}
