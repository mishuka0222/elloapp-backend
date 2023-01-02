package core

import (
	chatpb "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/chat/chat"
	messagepb "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/message/message"
	userpb "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/user/user"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

// MessagesGetUnreadMentions
// messages.getUnreadMentions#46578472 peer:InputPeer offset_id:int add_offset:int limit:int max_id:int min_id:int = messages.Messages;
func (c *MessagesCore) MessagesGetUnreadMentions(in *mtproto.TLMessagesGetUnreadMentions) (*mtproto.Messages_Messages, error) {
	// TODO(@benqi): 重复FromInputPeer2
	var (
		err  error
		peer = mtproto.FromInputPeer2(c.MD.UserId, in.GetPeer())
		chat *chatpb.MutableChat
		// channel   *channelpb.MutableChannel
		// isChannel bool
		minId   = in.MinId
		limit   = in.Limit
		rValues *mtproto.Messages_Messages
	)

	if limit > 50 {
		limit = 50
	}

	if c.MD.IsBot {
		err = mtproto.ErrBotMethodInvalid
		c.Logger.Errorf("messages.getHistory - error: %v", err)
		return nil, err
	}

	switch peer.PeerType {
	case mtproto.PEER_CHAT:
		// 400	CHAT_ID_INVALID	The provided chat id is invalid
		chat, err = c.svcCtx.Dao.ChatClient.Client().ChatGetMutableChat(
			c.ctx,
			&chatpb.TLChatGetMutableChat{
				ChatId: peer.PeerId,
			})
		if err != nil {
			err = mtproto.ErrPeerIdInvalid
			c.Logger.Errorf("messages.getHistory - error: %v", err)
			return nil, err
		} else {
			// TODO(@benqi): check migratedToId
			_ = chat
		}

		rValues = mtproto.MakeTLMessagesMessages(&mtproto.Messages_Messages{
			Messages: []*mtproto.Message{},
			Chats:    []*mtproto.Chat{},
			Users:    []*mtproto.User{},
		}).To_Messages_Messages()
	case mtproto.PEER_CHANNEL:
		//// 400	CHANNEL_INVALID	The provided channel is invalid
		//// 400	CHANNEL_PRIVATE	You haven't joined this channel/supergroup
		//channel, err = c.svcCtx.Dao.ChannelClient.ChannelGetMutableChannel(
		//	c.ctx,
		//	&channelpb.TLChannelGetMutableChannel{
		//		ChannelId: peer.PeerId,
		//		Id:        []int64{c.MD.UserId},
		//	})
		//if err != nil {
		//	err = mtproto.ErrPeerIdInvalid
		//	c.Logger.Errorf("messages.getHistory - error: %v", err)
		//	return nil, err
		//}
		//
		//me := channel.GetImmutableChannelParticipant(c.MD.UserId)
		//if me != nil {
		//	minId = int32(mathx.MaxInt(int(me.AvailableMinId), int(in.MinId)))
		//}
		//
		//isChannel = true
		//// TODO(@benqi): check kicked
		//_ = channel
		//
		//rValues = mtproto.MakeTLMessagesChannelMessages(&mtproto.Messages_Messages{
		//	Messages: []*mtproto.Message{},
		//	Chats:    []*mtproto.Chat{},
		//	Users:    []*mtproto.User{},
		//	Inexact:  false,
		//	Count:    0,
		//	NextRate: nil,
		//	Pts:      channel.Pts(),
		//}).To_Messages_Messages()
		c.Logger.Errorf("messages.readHistory blocked, License key from https://elloapp.com required to unlock enterprise features.")

		return nil, mtproto.ErrEnterpriseIsBlocked
	default:
		err = mtproto.ErrPeerIdInvalid
		c.Logger.Errorf("messages.getHistory - error: %v", err)
		return nil, err
	}

	// loadType := calcLoadHistoryType(isChannel, request.OffsetId, request.OffsetDate, request.AddOffset, limit, request.MaxId, request.MinId)
	boxList, _ := c.svcCtx.Dao.MessageClient.MessageGetUnreadMentions(
		c.ctx,
		&messagepb.TLMessageGetUnreadMentions{
			UserId:    c.MD.UserId,
			PeerType:  peer.PeerType,
			PeerId:    peer.PeerId,
			OffsetId:  in.OffsetId,
			AddOffset: in.AddOffset,
			Limit:     limit,
			MinId:     minId,
			MaxInt:    in.MaxId,
		})

	//if isChannel {
	//	rValues.Count = boxList.Length()
	//}

	boxList.Visit(c.MD.UserId,
		func(messageList []*mtproto.Message) {
			rValues.Messages = messageList
		},
		func(userIdList []int64) {
			mUsers, _ := c.svcCtx.Dao.UserClient.UserGetMutableUsers(c.ctx,
				&userpb.TLUserGetMutableUsers{
					Id: userIdList,
				})
			rValues.Users = append(rValues.Users, mUsers.GetUserListByIdList(c.MD.UserId, userIdList...)...)
		},
		func(chatIdList []int64) {
			mChats, _ := c.svcCtx.Dao.ChatClient.Client().ChatGetChatListByIdList(c.ctx,
				&chatpb.TLChatGetChatListByIdList{
					IdList: chatIdList,
				})
			rValues.Chats = append(rValues.Chats, mChats.GetChatListByIdList(c.MD.UserId, chatIdList...)...)
		},
		func(channelIdList []int64) {
			//mChannels, _ := c.svcCtx.Dao.ChannelClient.ChannelGetChannelListByIdList(c.ctx,
			//	&channelpb.TLChannelGetChannelListByIdList{
			//		SelfUserId: c.MD.UserId,
			//		Id:         channelIdList,
			//	})
			//if len(mChannels.GetDatas()) > 0 {
			//	rValues.Chats = append(rValues.Chats, mChannels.GetDatas()...)
			//}
		})

	return rValues, nil
}
