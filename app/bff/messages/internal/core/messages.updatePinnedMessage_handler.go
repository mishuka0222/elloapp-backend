package core

import (
	msgpb "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/messenger/msg/msg/msg"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

// MessagesUpdatePinnedMessage
// messages.updatePinnedMessage#d2aaf7ec flags:# silent:flags.0?true unpin:flags.1?true pm_oneside:flags.2?true peer:InputPeer id:int = Updates;
func (c *MessagesCore) MessagesUpdatePinnedMessage(in *mtproto.TLMessagesUpdatePinnedMessage) (*mtproto.Updates, error) {
	var (
		peer     = mtproto.FromInputPeer2(c.MD.UserId, in.Peer)
		rUpdates *mtproto.Updates
	)
	switch peer.PeerType {
	case mtproto.PEER_SELF:
	case mtproto.PEER_USER:
	case mtproto.PEER_CHAT:
	case mtproto.PEER_CHANNEL:
	default:
		c.Logger.Errorf("invalid peer: %v", in.Peer)
		err := mtproto.ErrPeerIdInvalid
		return nil, err
	}

	rUpdates, err := c.svcCtx.Dao.MsgClient.MsgUpdatePinnedMessage(c.ctx, &msgpb.TLMsgUpdatePinnedMessage{
		UserId:    c.MD.UserId,
		AuthKeyId: c.MD.AuthId,
		Silent:    in.Silent,
		Unpin:     in.Unpin,
		PmOneside: in.PmOneside,
		PeerType:  peer.PeerType,
		PeerId:    peer.PeerId,
		Id:        in.Id,
	})
	if err != nil {
		c.Logger.Errorf("messages.updatePinnedMessage - error: %v", in.Peer)
		return nil, err
	}

	return rUpdates, nil
}
