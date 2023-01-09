package core

import (
	chatpb "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/chat/chat"
	userpb "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/user/user"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

// MessagesGetExportedChatInvite
// messages.getExportedChatInvite#73746f5c peer:InputPeer link:string = messages.ExportedChatInvite;
func (c *ChatInvitesCore) MessagesGetExportedChatInvite(in *mtproto.TLMessagesGetExportedChatInvite) (*mtproto.Messages_ExportedChatInvite, error) {
	var (
		peer               = mtproto.FromInputPeer2(c.MD.UserId, in.Peer)
		exportedChatInvite *mtproto.ExportedChatInvite
		err                error
	)

	if !peer.IsChat() {
		err := mtproto.ErrPeerIdInvalid
		c.Logger.Errorf("messages.getExportedChatInvite - error: ", err)
		return nil, err
	}

	exportedChatInvite, err = c.svcCtx.Dao.ChatClient.ChatGetExportedChatInvite(c.ctx, &chatpb.TLChatGetExportedChatInvite{
		ChatId: peer.PeerId,
		Link:   in.GetLink(),
	})
	if err != nil {
		c.Logger.Errorf("messages.getExportedChatInvite - error: ", err)
		return nil, err
	}

	users, err2 := c.svcCtx.Dao.UserClient.UserGetMutableUsers(c.ctx, &userpb.TLUserGetMutableUsers{
		Id: []int64{c.MD.UserId, exportedChatInvite.AdminId},
	})
	if err2 != nil {
		err = mtproto.ErrPeerIdInvalid
		c.Logger.Errorf("messages.getExportedChatInvite - error: ", err)
		return nil, err
	}

	return mtproto.MakeTLMessagesExportedChatInvite(&mtproto.Messages_ExportedChatInvite{
		Invite: exportedChatInvite,
		Users:  users.GetUserListByIdList(c.MD.UserId, exportedChatInvite.AdminId),
	}).To_Messages_ExportedChatInvite(), nil
}
