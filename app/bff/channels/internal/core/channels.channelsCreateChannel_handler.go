package core

import (
	// "errors"
	"math/rand"
	"time"

	msgpb "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/messenger/msg/msg/msg"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/authsession/authsession"
	chatpb "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/chat/chat"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto/crypto"
	// "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/pkg/env2"
	"golang.org/x/crypto/bcrypt"

	userpb "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/user/user"
	usernamepb "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/username/username"

	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

func (c *ChannelsCore) ChannelsCreateChannel(in *mtproto.TLChannelsCreateChannel) (*mtproto.Updates, error) {
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

	channel, err := c.svcCtx.Dao.ChatClient.Client().ChatCreateChat2(c.ctx, &chatpb.TLChatCreateChat2{
		CreatorId: c.MD.UserId,
		// UserIdList: []int64{},
		UserIdList: channelUserIdList,
		Title:      channelTitle,
	})
	if err != nil {
		c.Logger.Errorf("createChat duplicate: %v", err)
		return nil, err
	}

	//Create user
	var userChannel *userpb.ImmutableUser
	hashPhone, err := bcrypt.GenerateFromPassword([]byte(channel.Title()+time.Now().String()), bcrypt.DefaultCost)
	phone := string(hashPhone[:])
	if err != nil {
		c.Logger.Errorf("create phone: %v", err)
	}
	key := crypto.CreateAuthKey()
	_, err = c.svcCtx.Dao.AuthsessionClient.AuthsessionSetAuthKey(c.ctx, &authsession.TLAuthsessionSetAuthKey{
		AuthKey: &mtproto.AuthKeyInfo{
			AuthKeyId:          key.AuthKeyId(),
			AuthKey:            key.AuthKey(),
			AuthKeyType:        mtproto.AuthKeyTypePerm,
			PermAuthKeyId:      key.AuthKeyId(),
			TempAuthKeyId:      0,
			MediaTempAuthKeyId: 0,
		},
		FutureSalt: nil,
	})
	if err != nil {
		c.Logger.Errorf("create user secret key error")
		return nil, err
	}
	var (
		username = in.Title
		firstName = in.Title
		lastName = in.About
	)

	if userChannel, err = c.svcCtx.Dao.UserClient.UserCreateNewUser(c.ctx, &userpb.TLUserCreateNewUser{
		SecretKeyId: key.AuthKeyId(),
		Phone:       phone,
		FirstName:   firstName,
		LastName:    lastName,
	}); err != nil {
		c.Logger.Errorf("createNewUser error: %v", err)
		return nil, err
	}
	_, err_un := c.svcCtx.Dao.UserUpdateUsername(c.ctx, &userpb.TLUserUpdateUsername{
		UserId:   userChannel.Id(),
		Username: username,
	})
	if err_un != nil {
		c.Logger.Errorf("username error: %v", err_un)
	}
	
	_, err_un = c.svcCtx.Dao.UsernameUpdateUsername(c.ctx, &usernamepb.TLUsernameUpdateUsername{
		PeerType: mtproto.PEER_CHANNEL,
		PeerId:   userChannel.Id(),
		Username: username,
	})
	if err_un != nil {
		c.Logger.Errorf("username error: %v", err_un)
	}

	_, err = c.svcCtx.Dao.AuthsessionClient.AuthsessionBindAuthKeyUser(c.ctx, &authsession.TLAuthsessionBindAuthKeyUser{
		AuthKeyId: key.AuthKeyId(),
		UserId:    userChannel.User.Id,
	})
	if err != nil {
		c.Logger.Errorf("bindAuthKeyUser error: %v", err)
		err = mtproto.ErrInternelServerError
		return nil, err
	}


	message := mtproto.MakeTLMessageService(&mtproto.Message{
		Out:         true,
		Mentioned:   false,
		MediaUnread: false,
		Silent:      false,
		Post:        false,
		Legacy:      false,
		Id:          0,
		FromId:      mtproto.MakePeerUser(c.MD.UserId),
		PeerId: mtproto.MakePeerUser(userChannel.Id()),
		ReplyTo:   nil,
		Date:      int32(time.Now().Unix()),
		TtlPeriod: nil,
		// Action:      &mtproto.MessageAction{},
		Action: mtproto.MakeMessageActionChannelCreate(channelTitle),
	})

	rValue, err := c.svcCtx.Dao.MsgClient.MsgSendMessage(c.ctx, &msgpb.TLMsgSendMessage{
		UserId:    c.MD.UserId,
		AuthKeyId: c.MD.AuthId,
		PeerType: mtproto.PEER_USER,
		PeerId:   userChannel.Id(),
		Message: msgpb.MakeTLOutboxMessage(&msgpb.OutboxMessage{
			NoWebpage:  true,
			Background: false,
			RandomId:   rand.Int63(),
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
