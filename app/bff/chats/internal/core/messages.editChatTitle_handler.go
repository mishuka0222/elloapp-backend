package core

import (
	"math/rand"

	msgpb "gitlab.com/merehead/elloapp/backend/elloapp_backend/app/messenger/msg/msg/msg"
	chatpb "gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/chat/chat"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
)

// MessagesEditChatTitle
// messages.editChatTitle#73783ffd chat_id:long title:string = Updates;
func (c *ChatsCore) MessagesEditChatTitle(in *mtproto.TLMessagesEditChatTitle) (*mtproto.Updates, error) {
	if in.Title == "" {
		err := mtproto.ErrChatTitleEmpty
		c.Logger.Errorf("messages.editChatTitle - error: ", err)
		return nil, err
	}

	chat, err := c.svcCtx.Dao.ChatClient.Client().ChatEditChatTitle(c.ctx, &chatpb.TLChatEditChatTitle{
		ChatId:     in.ChatId,
		EditUserId: c.MD.UserId,
		Title:      in.Title,
	})
	if err != nil {
		c.Logger.Errorf("messages.editChatTitle - error: ", err)
		return nil, err
	}

	replyUpdates, err := c.svcCtx.Dao.MsgClient.MsgSendMessage(c.ctx, &msgpb.TLMsgSendMessage{
		UserId:    c.MD.UserId,
		AuthKeyId: c.MD.AuthId,
		PeerType:  mtproto.PEER_CHAT,
		PeerId:    in.ChatId,
		Message: msgpb.MakeTLOutboxMessage(&msgpb.OutboxMessage{
			NoWebpage:    true,
			Background:   false,
			RandomId:     rand.Int63(),
			Message:      chat.MakeMessageService(c.MD.UserId, mtproto.MakeMessageActionChatEditTitle(in.Title)),
			ScheduleDate: nil,
		}).To_OutboxMessage(),
	})
	if err != nil {
		c.Logger.Errorf("messages.editChatTitle - error: %v", err)
		return nil, err
	}

	return replyUpdates, nil
}
