package core

import (
	"errors"
	"github.com/zeromicro/go-zero/core/jsonx"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/channels/channels"
	userpb "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/user/user"

	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

func (c *ChannelsCore) ChannelsGetFullChannel(in *mtproto.TLChannelsGetFullChannel) (res *mtproto.Messages_ChatFull, err error) {

	if in.Channel.Constructor == mtproto.CRC32_inputChannelEmpty {
		err = errors.New("BAD_REQUEST")
		return
	}

	var (
		inputChannel = in.GetChannel().To_InputChannel()
		channel      *channels.ChannelData
		resFull      *channels.GetChannelFullBySelfIdResp
		resChat      *channels.ToChatResp
	)

	channel, err = c.svcCtx.Dao.ChannelsClient.GetChannelDataById(c.ctx, &channels.ChannelDataByIdReq{ChannelId: inputChannel.GetChannelId()})
	if err != nil {
		return
	}

	resFull, err = c.svcCtx.Dao.ChannelsClient.GetChannelFullBySelfId(c.ctx, &channels.GetChannelFullBySelfIdReq{
		ChannelData: channel,
		SelfUserId:  c.MD.UserId,
	})
	if err != nil {
		return
	}

	resChat, err = c.svcCtx.Dao.ChannelsClient.ToChat(c.ctx, &channels.ToChatReq{
		Channel:    channel,
		SelfUserId: c.MD.UserId,
	})
	if err != nil {
		return
	}
	var idList []int64
	for i := range channel.Participants {
		idList = append(idList, channel.Participants[i].UserId)
	}
	mUsers, err2 := c.svcCtx.Dao.UserClient.UserGetMutableUsers(c.ctx, &userpb.TLUserGetMutableUsers{
		Id: idList,
	})
	if err2 != nil {
		c.Logger.Errorf("channels.getFullChannel - error: not found dialog")
	}

	res = mtproto.MakeTLMessagesChatFull(&mtproto.Messages_ChatFull{
		FullChat: resFull.Channel.To_ChatFull(),
		Chats:    []*mtproto.Chat{resChat.Chat},
		Users:    mUsers.GetUserListByIdList(c.MD.UserId, idList...),
	}).To_Messages_ChatFull()

	v, _ := jsonx.MarshalToString(res)
	c.Logger.Debug("MY_DEBUG", v)

	return
}
