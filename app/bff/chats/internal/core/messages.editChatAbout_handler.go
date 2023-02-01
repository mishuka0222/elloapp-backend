package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/channels/channels"
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
		_, err := c.svcCtx.Dao.ChannelsClient.EditChannelAbout(c.ctx, &channels.EditChannelAboutReq{
			ChannelId:  in.Peer.GetChannelId(),
			EditUserId: c.MD.UserId,
			About:      in.About,
		})
		if err != nil {
			c.Logger.Errorf("messages.editChannelAbout - error: %v", err)
			return nil, err
		}

	default:
		err := mtproto.ErrPeerIdInvalid
		c.Logger.Errorf("invalid peer type: {%v}")
		return nil, err
	}

	return mtproto.BoolTrue, nil
}
