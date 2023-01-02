package core

import (
	"math/rand"

	msgpb "gitlab.com/merehead/elloapp/backend/elloapp_backend/app/messenger/msg/msg/msg"
	chatpb "gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/chat/chat"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
)

// MessagesImportChatInvite
// messages.importChatInvite#6c50051c hash:string = Updates;
func (c *ChatInvitesCore) MessagesImportChatInvite(in *mtproto.TLMessagesImportChatInvite) (*mtproto.Updates, error) {
	// Code	Type	Description
	// 400	INVITE_HASH_EMPTY	The invite hash is empty.
	// 400	INVITE_HASH_EXPIRED	The invite link has expired.
	// 400	INVITE_HASH_INVALID	The invite hash is invalid.

	if len(in.Hash) == 0 {
		err := mtproto.ErrInviteHashEmpty
		c.Logger.Errorf("messages.importChatInvite - error: %v", err)
		return nil, err
	}
	if len(in.Hash) != 20 {
		err := mtproto.ErrInviteHashInvalid
		c.Logger.Errorf("messages.importChatInvite - error: %v", err)
		return nil, err
	}

	peerType := chatpb.GetChatTypeByInviteHash(in.Hash)
	switch peerType {
	case mtproto.PEER_CHAT:
		mChat, err := c.svcCtx.Dao.ChatClient.ChatImportChatInvite(c.ctx, &chatpb.TLChatImportChatInvite{
			SelfId: c.MD.UserId,
			Hash:   in.Hash,
		})
		if err != nil {
			c.Logger.Errorf("messages.importChatInvite - error: %v", err)
			return nil, err
		}

		// TODO: found link
		rUpdates, err := c.svcCtx.Dao.MsgClient.MsgSendMessage(
			c.ctx,
			&msgpb.TLMsgSendMessage{
				UserId:    c.MD.UserId,
				AuthKeyId: c.MD.AuthId,
				PeerType:  mtproto.PEER_CHAT,
				PeerId:    mChat.Id(),
				Message: msgpb.MakeTLOutboxMessage(&msgpb.OutboxMessage{
					NoWebpage:    true,
					Background:   false,
					RandomId:     rand.Int63(),
					Message:      mChat.MakeMessageService(c.MD.UserId, mtproto.MakeMessageActionChatJoinByLink(mChat.Creator())),
					ScheduleDate: nil,
				}).To_OutboxMessage(),
			})
		if err != nil {
			c.Logger.Errorf("messages.importChatInvite - error: %v", err)
			return nil, err
		}

		return rUpdates, nil
	case mtproto.PEER_CHANNEL:
		c.Logger.Errorf("messages.importChatInvite blocked, License key from https://elloapp.com required to unlock enterprise features.")

		return nil, mtproto.ErrEnterpriseIsBlocked
	default:
		err := mtproto.ErrInviteHashInvalid
		c.Logger.Errorf("messages.importChatInvite - error: %v", err)
		return nil, err
	}
}
