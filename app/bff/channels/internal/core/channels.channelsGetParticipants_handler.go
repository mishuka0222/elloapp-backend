package core

import (
	"errors"
	"github.com/zeromicro/go-zero/core/jsonx"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/channels/channels"
	userpb "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/user/user"

	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

func (c *ChannelsCore) ChannelsGetParticipants(in *mtproto.TLChannelsGetParticipants) (res *mtproto.Channels_ChannelParticipants, err error) {

	if in.Channel.Constructor == mtproto.CRC32_inputChannelEmpty {
		err = errors.New("BAD_REQUEST")
		return
	}

	var (
		inputChannel = in.GetChannel().To_InputChannel()
		channel      *channels.ChannelData
		participants *channels.GetChannelParticipantsResp
	)

	channel, err = c.svcCtx.Dao.ChannelsClient.GetChannelDataById(c.ctx, &channels.ChannelDataByIdReq{ChannelId: inputChannel.GetChannelId()})
	if err != nil {
		return
	}

	participants, err = c.svcCtx.Dao.ChannelsClient.GetChannelParticipants(c.ctx,
		&channels.ChannelParticipantsReq{ChannelId: inputChannel.GetChannelId()})
	if err != nil {
		return
	}

	r, _ := c.svcCtx.Dao.ChannelsClient.ToChat(c.ctx, &channels.ToChatReq{Channel: channel, SelfUserId: c.MD.UserId})

	participants.Participants.SetChats([]*mtproto.Chat{r.Chat})

	var idList []int64
	for i := range channel.Participants {
		idList = append(idList, channel.Participants[i].UserId)
	}
	mUsers, err := c.svcCtx.Dao.UserClient.UserGetMutableUsers(c.ctx, &userpb.TLUserGetMutableUsers{
		Id: idList,
	})
	if err != nil {
		c.Logger.Errorf("messages.getFullChat - error: not found dialog")
	}
	participants.Participants.SetUsers(mUsers.GetUserListByIdList(c.MD.UserId, idList...))

	res = participants.Participants.To_Channels_ChannelParticipants()

	v, _ := jsonx.MarshalToString(res)
	c.Logger.Debug("MY_DEBUG", v)

	return
}
