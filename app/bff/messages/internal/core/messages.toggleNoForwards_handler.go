package core

import (
	chatpb "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/chat/chat"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

// MessagesToggleNoForwards
// messages.toggleNoForwards#b11eafa2 peer:InputPeer enabled:Bool = Updates;
func (c *MessagesCore) MessagesToggleNoForwards(in *mtproto.TLMessagesToggleNoForwards) (*mtproto.Updates, error) {
	var (
		peer     = mtproto.FromInputPeer2(c.MD.UserId, in.Peer)
		rUpdates *mtproto.Updates
	)

	if !peer.IsChat() {
		err := mtproto.ErrPeerIdInvalid
		c.Logger.Errorf("messages.toggleNoForwards - error: %v", err)
		return nil, err
	}

	chat, err := c.svcCtx.Dao.ChatClient.Client().ChatToggleNoForwards(c.ctx, &chatpb.TLChatToggleNoForwards{
		ChatId:     peer.PeerId,
		OperatorId: c.MD.UserId,
		Enabled:    in.Enabled,
	})
	if err != nil {
		c.Logger.Errorf("messages.toggleNoForwards - error: %v", err)
		return nil, err
	}

	rUpdates = mtproto.MakeUpdatesByUpdatesChats(
		[]*mtproto.Chat{chat.ToUnsafeChat(c.MD.UserId)},
		mtproto.MakeTLUpdateChat(&mtproto.Update{
			ChatId_INT64: peer.PeerId,
		}).To_Update())

	return rUpdates, nil
}
