package core

import (
	msgpb "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/messenger/msg/msg/msg"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

// MessagesDeleteMessages
// messages.deleteMessages#e58e95d2 flags:# revoke:flags.0?true id:Vector<int> = messages.AffectedMessages;
func (c *MessagesCore) MessagesDeleteMessages(in *mtproto.TLMessagesDeleteMessages) (*mtproto.Messages_AffectedMessages, error) {
	// TODO(@benqi): Check message service.
	affectedMessages, err := c.svcCtx.Dao.MsgClient.MsgDeleteMessages(c.ctx, &msgpb.TLMsgDeleteMessages{
		UserId:    c.MD.UserId,
		AuthKeyId: c.MD.AuthId,
		PeerType:  mtproto.PEER_EMPTY,
		PeerId:    0,
		Revoke:    in.Revoke,
		Id:        in.Id,
	})

	if err != nil {
		c.Logger.Errorf("messages.deleteMessages - error: %v", err)
		return nil, err
	}

	return affectedMessages, nil
}
