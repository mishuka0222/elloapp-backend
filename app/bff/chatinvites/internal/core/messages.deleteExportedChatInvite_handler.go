package core

import (
	chatpb "gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/chat/chat"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
)

// MessagesDeleteExportedChatInvite
// messages.deleteExportedChatInvite#d464a42b peer:InputPeer link:string = Bool;
func (c *ChatInvitesCore) MessagesDeleteExportedChatInvite(in *mtproto.TLMessagesDeleteExportedChatInvite) (*mtproto.Bool, error) {
	var (
		peer = mtproto.FromInputPeer2(c.MD.UserId, in.Peer)
		err  error
	)

	if !peer.IsChat() {
		err = mtproto.ErrPeerIdInvalid
		c.Logger.Errorf("messages.deleteExportedChatInvite - error: ", err)
		return nil, err
	}

	_, err = c.svcCtx.Dao.ChatClient.ChatDeleteExportedChatInvite(c.ctx, &chatpb.TLChatDeleteExportedChatInvite{
		SelfId: c.MD.UserId,
		ChatId: peer.PeerId,
		Link:   in.GetLink(),
	})
	if err != nil {
		c.Logger.Errorf("messages.deleteExportedChatInvite - error: %v", err)
		return nil, err
	}

	return mtproto.BoolTrue, nil
}
