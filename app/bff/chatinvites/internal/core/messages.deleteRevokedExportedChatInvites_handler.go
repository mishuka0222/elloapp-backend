package core

import (
	chatpb "gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/chat/chat"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
)

// MessagesDeleteRevokedExportedChatInvites
// messages.deleteRevokedExportedChatInvites#56987bd5 peer:InputPeer admin_id:InputUser = Bool;
func (c *ChatInvitesCore) MessagesDeleteRevokedExportedChatInvites(in *mtproto.TLMessagesDeleteRevokedExportedChatInvites) (*mtproto.Bool, error) {
	var (
		err     error
		peer    = mtproto.FromInputPeer2(c.MD.UserId, in.Peer)
		adminId = mtproto.FromInputUser(c.MD.UserId, in.AdminId)
	)

	if !peer.IsChat() {
		err = mtproto.ErrPeerIdInvalid
		c.Logger.Errorf("messages.deleteRevokedExportedChatInvites - error: ", err)
		return nil, err
	}

	_, err = c.svcCtx.Dao.ChatClient.ChatDeleteRevokedExportedChatInvites(c.ctx, &chatpb.TLChatDeleteRevokedExportedChatInvites{
		SelfId:  c.MD.UserId,
		ChatId:  peer.PeerId,
		AdminId: adminId.PeerId,
	})
	if err != nil {
		c.Logger.Errorf("messages.deleteRevokedExportedChatInvites - error: %v", err)
		return nil, err
	}

	return mtproto.BoolTrue, nil
}
