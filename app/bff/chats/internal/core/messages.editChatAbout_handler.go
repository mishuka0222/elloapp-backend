package core

import (
	chatpb "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/chat/chat"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

// MessagesEditChatAbout
// messages.editChatAbout#def60797 peer:InputPeer about:string = Bool;
func (c *ChatsCore) MessagesEditChatAbout(in *mtproto.TLMessagesEditChatAbout) (*mtproto.Bool, error) {
	peer := mtproto.FromInputPeer2(c.MD.UserId, in.Peer)

	switch peer.PeerType {
	case mtproto.PEER_CHAT:
		_, err := c.svcCtx.Dao.ChatClient.Client().ChatEditChatAbout(c.ctx, &chatpb.TLChatEditChatAbout{
			ChatId:     peer.PeerId,
			EditUserId: c.MD.UserId,
			About:      in.About,
		})
		if err != nil {
			c.Logger.Errorf("messages.editChatAbout - error: %v", err)
			return nil, err
		}
	case mtproto.PEER_CHANNEL:
		c.Logger.Errorf("messages.editChatAbout blocked, License key from https://elloapp.com required to unlock enterprise features.")

		return nil, mtproto.ErrEnterpriseIsBlocked
	default:
		err := mtproto.ErrPeerIdInvalid
		c.Logger.Errorf("invalid peer type: {%v}")
		return nil, err
	}

	return mtproto.BoolTrue, nil
}
