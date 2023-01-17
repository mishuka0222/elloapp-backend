package core

import (
	// "errors"
	"time"
	"math/rand"

	msgpb "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/messenger/msg/msg/msg"
	// chatpb "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/chat/chat"

	userpb "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/user/user"

	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

func (c *ChannelsCore) ChannelsCreateChannel(in *mtproto.TLChannelsCreateChannel) (*mtproto.Updates, error) {
	// todo: add your logic here and delete this line
	// return nil, errors.New("unimplemented 123")
	// return &mtproto.Updates{}, nil
	var (
		channelUserIdList []int64
		channelTitle      = in.GetTitle()
		// channelAbout      = in.GetAbout()
	)

	if channelTitle == "" {
		err := mtproto.ErrChatTitleEmpty
		c.Logger.Errorf("channels.createChannel - error: %v", err)
		return nil, err
	}

	channelUserIdList = []int64{c.MD.UserId}
	users, _ := c.svcCtx.Dao.UserClient.UserGetMutableUsers(c.ctx, &userpb.TLUserGetMutableUsers{
		Id: channelUserIdList,
	})

	if me, _ := users.GetImmutableUser(c.MD.UserId); me.Restricted() {
		err := mtproto.ErrUserRestricted
		c.Logger.Errorf("channels.createChannel - error: %v", err)
		return nil, err
	}

	// chat, err := c.svcCtx.Dao.ChatClient.Client().ChatCreateChat2(c.ctx, &chatpb.TLChatCreateChat2{
	// 	CreatorId:  c.MD.UserId,
	// 	UserIdList: channelUserIdList,
	// 	Title:      channelTitle,
	// })
	// if err != nil {
	// 	c.Logger.Errorf("createChat duplicate: %v", err)
	// 	return nil, err
	// }

	message := mtproto.MakeTLMessageService(&mtproto.Message{
		Out:         true,
		Mentioned:   false,
		MediaUnread: false,
		Silent:      false,
		Post:        false,
		Legacy:      false,
		Id:          0,
		FromId:      mtproto.MakePeerUser(c.MD.UserId),
		PeerId:      mtproto.MakePeerUser(c.MD.UserId),
		// FromId:      &mtproto.Peer{},
		// PeerId:      &mtproto.Peer{},
		ReplyTo:     nil,
		Date:        int32(time.Now().Unix()),
		TtlPeriod:   nil,
		// Action:      &mtproto.MessageAction{},
		Action:      mtproto.MakeMessageActionChannelCreate(channelTitle),
	})

	rValue, err := c.svcCtx.Dao.MsgClient.MsgSendMessage(c.ctx, &msgpb.TLMsgSendMessage{
		UserId:    c.MD.UserId,
		AuthKeyId: c.MD.AuthId,
		PeerType:  mtproto.PEER_USER,
		// PeerId:    chat.Chat.Id,
		PeerId:      c.MD.UserId,
		Message: msgpb.MakeTLOutboxMessage(&msgpb.OutboxMessage{
			NoWebpage:  true,
			Background: false,
			RandomId:   rand.Int63(),
			// Message:      chat.MakeMessageService(c.MD.UserId, mtproto.MakeMessageActionChannelCreate(channelTitle)),
			Message:      message.To_Message(),
			ScheduleDate: nil,
		}).To_OutboxMessage(),
	})
	if err != nil {
		c.Logger.Errorf("messages.createChat - error: %v", err)
		return nil, err
	}

	return rValue, nil

}
