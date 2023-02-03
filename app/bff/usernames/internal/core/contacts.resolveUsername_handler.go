package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/channels/channels"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/chat/chat"
	userpb "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/user/user"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/username/username"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

// ContactsResolveUsername
// contacts.resolveUsername#f93ccba3 username:string = contacts.ResolvedPeer;
func (c *UsernamesCore) ContactsResolveUsername(in *mtproto.TLContactsResolveUsername) (*mtproto.Contacts_ResolvedPeer, error) {
	// TODO(@benqi):
	// 401	AUTH_KEY_PERM_EMPTY	The temporary auth key must be binded to the permanent auth key to use these methods.
	// 401	SESSION_PASSWORD_NEEDED	2FA is enabled, use a password to login
	// 400	USERNAME_INVALID	The provided username is not valid
	// 400	USERNAME_NOT_OCCUPIED	The provided username is not occupied
	//
	var (
		peer    *mtproto.PeerUtil
		usrname = in.GetUsername()
	)

	//if strings.Contains(usrname, "joinchannel/") {
	//	res, err := c.svcCtx.Dao.ChannelsClient.FindByJoinLink(c.ctx, &channels.FindByJoinLinkReq{
	//		Link: usrname,
	//	})
	//	if err != nil {
	//		return nil, err
	//	}
	//	peer = mtproto.MakeChannelPeerUtil(res.ChannelId)
	//} else {
	id := userpb.GetBotIdByName(usrname)
	if id > 0 {
		peer = mtproto.MakeUserPeerUtil(id)
	} else {
		rName, err := c.svcCtx.Dao.UsernameClient.UsernameResolveUsername(c.ctx, &username.TLUsernameResolveUsername{
			Username: in.GetUsername(),
		})
		if err != nil {
			c.Logger.Errorf("contacts.resolveUsername - reply: {%v}", err)
			return nil, err
		}

		peer = mtproto.FromPeer(rName)
	}
	//}

	resolvedPeer := mtproto.MakeTLContactsResolvedPeer(&mtproto.Contacts_ResolvedPeer{
		Peer:  peer.ToPeer(),
		Chats: []*mtproto.Chat{},
		Users: []*mtproto.User{},
	}).To_Contacts_ResolvedPeer()

	switch peer.PeerType {
	case mtproto.PEER_USER:
		mUsers, _ := c.svcCtx.Dao.UserClient.UserGetMutableUsers(c.ctx, &userpb.TLUserGetMutableUsers{
			Id: []int64{c.MD.UserId, peer.PeerId},
		})
		// .UserFacade.GetUserById(ctx, md.UserId, peer.PeerId)
		if mUsers != nil {
			resolvedPeer.Users = mUsers.GetUserListByIdList(c.MD.UserId, peer.PeerId)
		}
	case mtproto.PEER_CHAT:
		chat, _ := c.svcCtx.Dao.ChatClient.ChatGetChatBySelfId(c.ctx, &chat.TLChatGetChatBySelfId{
			SelfId: c.MD.UserId,
			ChatId: peer.PeerId,
		})
		if chat != nil {
			resolvedPeer.Chats = []*mtproto.Chat{chat.ToUnsafeChat(c.MD.UserId)}
		}
	case mtproto.PEER_CHANNEL:
		channel, err := c.svcCtx.Dao.ChannelsClient.GetAllChannelDataById(c.ctx, &channels.ChannelDataByIdReq{
			ChannelId: peer.PeerId,
		})
		if err != nil {
			return nil, err
		}
		rChat, err := c.svcCtx.Dao.ChannelsClient.ToChat(c.ctx, &channels.ToChatReq{
			Channel:    channel,
			SelfUserId: c.MD.UserId,
		})
		if err != nil {
			return nil, err
		}
		resolvedPeer.Chats = []*mtproto.Chat{rChat.Chat}
	}

	return resolvedPeer, nil
}
