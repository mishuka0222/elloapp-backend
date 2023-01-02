package core

import (
	chatpb "gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/chat/chat"
	userpb "gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/user/user"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
)

// MessagesEditExportedChatInvite
// messages.editExportedChatInvite#bdca2f75 flags:# revoked:flags.2?true peer:InputPeer link:string expire_date:flags.0?int usage_limit:flags.1?int request_needed:flags.3?Bool title:flags.4?string = messages.ExportedChatInvite;
func (c *ChatInvitesCore) MessagesEditExportedChatInvite(in *mtproto.TLMessagesEditExportedChatInvite) (*mtproto.Messages_ExportedChatInvite, error) {
	var (
		peer    = mtproto.FromInputPeer2(c.MD.UserId, in.Peer)
		invites []*mtproto.ExportedChatInvite
		rValue  *mtproto.Messages_ExportedChatInvite
	)

	if !peer.IsChat() {
		err := mtproto.ErrPeerIdInvalid
		c.Logger.Errorf("messages.editExportedChatInvite - error: ", err)
		return nil, err
	}

	chatInvites, err := c.svcCtx.Dao.ChatClient.ChatEditExportedChatInvite(c.ctx, &chatpb.TLChatEditExportedChatInvite{
		SelfId:        c.MD.UserId,
		ChatId:        peer.PeerId,
		Revoked:       in.Revoked,
		Link:          in.Link,
		ExpireDate:    in.ExpireDate,
		UsageLimit:    in.UsageLimit,
		RequestNeeded: in.RequestNeeded,
		Title:         in.Title,
	})
	if err != nil {
		c.Logger.Errorf("messages.editExportedChatInvite - error: %v", err)
		return nil, err
	}
	if len(chatInvites.Datas) == 0 || len(chatInvites.Datas) > 2 {
		err = mtproto.ErrInternelServerError
		c.Logger.Errorf("messages.editExportedChatInvite - error: %v", err)
		return nil, err
	}

	invites = chatInvites.Datas

	users, err2 := c.svcCtx.Dao.UserClient.UserGetMutableUsers(c.ctx, &userpb.TLUserGetMutableUsers{
		Id: []int64{c.MD.UserId, invites[0].AdminId},
	})
	if err2 != nil {
		err := mtproto.ErrPeerIdInvalid
		c.Logger.Errorf("messages.editExportedChatInvite - error: ", err)
		return nil, err
	}
	inviteUsers := users.GetUserListByIdList(c.MD.UserId, invites[0].AdminId)

	if len(invites) == 1 {
		rValue = mtproto.MakeTLMessagesExportedChatInvite(&mtproto.Messages_ExportedChatInvite{
			Invite: invites[0],
			Users:  inviteUsers,
		}).To_Messages_ExportedChatInvite()
	} else {
		rValue = mtproto.MakeTLMessagesExportedChatInviteReplaced(&mtproto.Messages_ExportedChatInvite{
			Invite:    invites[0],
			NewInvite: invites[1],
			Users:     inviteUsers,
		}).To_Messages_ExportedChatInvite()
	}

	return rValue, nil
}
