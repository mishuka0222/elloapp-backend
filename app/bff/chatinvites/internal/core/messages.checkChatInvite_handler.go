package core

import (
	chatpb "gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/chat/chat"
	userpb "gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/user/user"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
)

// MessagesCheckChatInvite
// messages.checkChatInvite#3eadb1bb hash:string = ChatInvite;
func (c *ChatInvitesCore) MessagesCheckChatInvite(in *mtproto.TLMessagesCheckChatInvite) (*mtproto.ChatInvite, error) {
	// Code	Type	Description
	// 400	INVITE_HASH_EMPTY	The invite hash is empty.
	// 400	INVITE_HASH_EXPIRED	The invite link has expired.
	// 400	INVITE_HASH_INVALID	The invite hash is invalid.

	if len(in.Hash) == 0 {
		err := mtproto.ErrInviteHashEmpty
		c.Logger.Errorf("messages.checkChatInvite - error: %v", err)
		return nil, err
	}
	if len(in.Hash) != 20 {
		err := mtproto.ErrInviteHashInvalid
		c.Logger.Errorf("messages.checkChatInvite - error: %v", err)
		return nil, err
	}

	getUserListF := func(idList []int64) []*mtproto.User {
		users, _ := c.svcCtx.Dao.UserClient.UserGetMutableUsers(c.ctx, &userpb.TLUserGetMutableUsers{
			Id: append(idList, c.MD.UserId),
		})
		return users.GetUserListByIdList(c.MD.UserId, idList...)
	}

	peerType := chatpb.GetChatTypeByInviteHash(in.Hash)
	switch peerType {
	case mtproto.PEER_CHAT:
		chatInviteExt, err := c.svcCtx.Dao.ChatClient.ChatCheckChatInvite(c.ctx, &chatpb.TLChatCheckChatInvite{
			SelfId: c.MD.UserId,
			Hash:   in.Hash,
		})
		if err != nil {
			c.Logger.Errorf("messages.checkChatInvite - error: %v", err)
			return nil, err
		}

		rValue := chatInviteExt.ToChatInvite(c.MD.UserId, func(idList []int64) []*mtproto.User {
			return getUserListF(idList)
		})
		if rValue == nil {
			err = mtproto.ErrInternelServerError
			c.Logger.Errorf("messages.checkChatInvite - error: ", err)
			return nil, err
		}

		return rValue, nil
	case mtproto.PEER_CHANNEL:
		c.Logger.Errorf("messages.checkChatInvite blocked, License key from https://elloapp.com required to unlock enterprise features.")

		return nil, mtproto.ErrEnterpriseIsBlocked
	default:
		err := mtproto.ErrInviteHashInvalid
		c.Logger.Errorf("messages.checkChatInvite - error: %v", err)
		return nil, err
	}
}
