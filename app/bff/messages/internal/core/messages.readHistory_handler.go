package core

import (
	msgpb "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/messenger/msg/msg/msg"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

// MessagesReadHistory
// messages.readHistory#e306d3a peer:InputPeer max_id:int = messages.AffectedMessages;
func (c *MessagesCore) MessagesReadHistory(in *mtproto.TLMessagesReadHistory) (*mtproto.Messages_AffectedMessages, error) {
	var (
		peer = mtproto.FromInputPeer2(c.MD.UserId, in.Peer)
	)

	switch peer.PeerType {
	case mtproto.PEER_SELF:
	case mtproto.PEER_USER:
	case mtproto.PEER_CHAT:
	default:
		c.Logger.Errorf("invalid peer: %v", in.Peer)
		err := mtproto.ErrPeerIdInvalid
		return nil, err
	}

	return c.svcCtx.Dao.MsgClient.MsgReadHistory(c.ctx, &msgpb.TLMsgReadHistory{
		UserId:    c.MD.UserId,
		AuthKeyId: c.MD.AuthId,
		PeerType:  peer.PeerType,
		PeerId:    peer.PeerId,
		MaxId:     in.MaxId,
	})
}
