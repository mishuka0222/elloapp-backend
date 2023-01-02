package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/dialog/dialog"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
)

// MessagesGetDialogUnreadMarks
// messages.getDialogUnreadMarks#22e24e22 = Vector<DialogPeer>;
func (c *DialogsCore) MessagesGetDialogUnreadMarks(in *mtproto.TLMessagesGetDialogUnreadMarks) (*mtproto.Vector_DialogPeer, error) {
	_ = in

	dialogPeerList, err := c.svcCtx.DialogClient.DialogGetDialogUnreadMarkList(
		c.ctx,
		&dialog.TLDialogGetDialogUnreadMarkList{
			UserId: c.MD.UserId,
		})
	if err != nil {
		c.Logger.Errorf("messages.getDialogUnreadMarks - error: %v", err)
		return nil, err
	}

	return &mtproto.Vector_DialogPeer{
		Datas: dialogPeerList.Datas,
	}, nil
}
