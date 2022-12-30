package core

import (
	chatpb "gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/chat/chat"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
)

// MessagesExportChatInvite
// messages.exportChatInvite#a02ce5d5 flags:# legacy_revoke_permanent:flags.2?true request_needed:flags.3?true peer:InputPeer expire_date:flags.0?int usage_limit:flags.1?int title:flags.4?string = ExportedChatInvite;
func (c *ChatInvitesCore) MessagesExportChatInvite(in *mtproto.TLMessagesExportChatInvite) (*mtproto.ExportedChatInvite, error) {
	var (
		peer               = mtproto.FromInputPeer2(c.MD.UserId, in.Peer)
		err                error
		exportedChatInvite *mtproto.ExportedChatInvite
	)

	if !peer.IsChat() {
		err = mtproto.ErrPeerIdInvalid
		c.Logger.Errorf("messages.exportChatInvite - error: ", err)
		return nil, err
	}

	exportedChatInvite, err = c.svcCtx.Dao.ChatClient.ChatExportChatInvite(c.ctx, &chatpb.TLChatExportChatInvite{
		ChatId:                peer.PeerId,
		AdminId:               c.MD.UserId,
		LegacyRevokePermanent: in.LegacyRevokePermanent,
		RequestNeeded:         in.RequestNeeded,
		ExpireDate:            in.ExpireDate,
		UsageLimit:            in.UsageLimit,
		Title:                 in.Title,
	})
	if err != nil {
		c.Logger.Errorf("messages.exportChatInvite - error: ", err)
		return nil, err
	}

	// TODO(FIXME): Check logic
	return exportedChatInvite, nil
}
