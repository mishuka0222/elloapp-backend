package core

import (
	chatpb "gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/chat/chat"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/pkg2/utils"
)

// MessagesGetCommonChats
// messages.getCommonChats#e40ca104 user_id:InputUser max_id:long limit:int = messages.Chats;
func (c *ChatsCore) MessagesGetCommonChats(in *mtproto.TLMessagesGetCommonChats) (*mtproto.Messages_Chats, error) {
	var (
		userId        = in.GetUserId().GetUserId()
		count         int32
		chats         []*mtproto.Chat
		messagesChats *mtproto.Messages_Chats
	)

	// TODO(@benqi): check
	// 400	MSG_ID_INVALID	Invalid message ID provided
	// 400	USER_ID_INVALID	The provided user ID is invalid

	if userId == c.MD.UserId {
		err := mtproto.ErrUserIdInvalid
		c.Logger.Errorf("messages.getCommonChats - error: %v", err)
		return nil, err
	}

	// chat
	usersChatIdList, err := c.svcCtx.Dao.ChatClient.Client().ChatGetUsersChatIdList(c.ctx, &chatpb.TLChatGetUsersChatIdList{
		Id: []int64{c.MD.UserId, userId},
	})
	if err != nil {
		c.Logger.Errorf("messages.getCommonChats - error: %v", err)
		return nil, err
	}

	if len(usersChatIdList.Datas) == 2 {
		commonChats := utils.Int64Intersect(usersChatIdList.Datas[0].ChatIdList, usersChatIdList.Datas[1].ChatIdList)
		count = int32(len(commonChats))
		c.Logger.Errorf("messages.getCommonChats - commonChats: %v", commonChats)
		found := false
		for i, id := range commonChats {
			if id > in.MaxId {
				commonChats = commonChats[i:]
				found = true
				break
			}
		}
		if found {
			mChats, _ := c.svcCtx.Dao.ChatClient.Client().ChatGetChatListByIdList(c.ctx, &chatpb.TLChatGetChatListByIdList{
				SelfId: c.MD.UserId,
				IdList: commonChats,
			})
			chats = mChats.GetChatListByIdList(c.MD.UserId, commonChats...)
		}
	} else {
		chats = []*mtproto.Chat{}
	}

	// TODO: channel

	if count > 0 {
		messagesChats = mtproto.MakeTLMessagesChats(&mtproto.Messages_Chats{
			Chats: chats,
		}).To_Messages_Chats()
	} else {
		messagesChats = mtproto.MakeTLMessagesChatsSlice(&mtproto.Messages_Chats{
			Count: count,
			Chats: []*mtproto.Chat{},
		}).To_Messages_Chats()
	}

	return messagesChats, nil
}
