package core

import (
	chatpb "gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/chat/chat"
	messagepb "gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/message/message"
	userpb "gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/user/user"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
	"math"
)

// MessagesSearchGlobal
// messages.searchGlobal#4bc6589a flags:# folder_id:flags.0?int q:string filter:MessagesFilter min_date:int max_date:int offset_rate:int offset_peer:InputPeer offset_id:int limit:int = messages.Messages;
func (c *MessagesCore) MessagesSearchGlobal(in *mtproto.TLMessagesSearchGlobal) (*mtproto.Messages_Messages, error) {
	// 400	BOT_METHOD_INVALID	This method can't be used by a bot
	// 400	SEARCH_QUERY_EMPTY	The search query is empty
	if c.MD.IsBot {
		err := mtproto.ErrBotMethodInvalid
		c.Logger.Errorf("messages.searchGlobal - error: %v", err)
		return nil, err
	}

	if in.Q == "" {
		err := mtproto.ErrSearchQueryEmpty
		c.Logger.Errorf("messages.searchGlobal - error: %v", err)
		return nil, err
	}

	var (
		offsetId = in.OffsetId
		limit    = in.Limit
	)

	if offsetId == 0 {
		offsetId = math.MaxInt32
	}

	if limit > 50 {
		limit = 50
	}

	// TODO(@benqi): Impl MessagesSearchGlobal logic
	rValues := mtproto.MakeTLMessagesMessages(&mtproto.Messages_Messages{
		Messages: []*mtproto.Message{},
		Chats:    []*mtproto.Chat{},
		Users:    []*mtproto.User{},
	}).To_Messages_Messages()

	boxList, err := c.svcCtx.Dao.MessageClient.MessageSearchGlobal(
		c.ctx,
		&messagepb.TLMessageSearchGlobal{
			UserId: c.MD.UserId,
			Q:      in.Q,
			Offset: offsetId,
			Limit:  limit,
		})
	if err != nil {
		c.Logger.Errorf("messages.searchGlobal - error: %v", err)
		return rValues, nil
	}

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
