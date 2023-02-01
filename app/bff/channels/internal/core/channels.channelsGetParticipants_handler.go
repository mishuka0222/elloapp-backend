package core

import (
	"errors"
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
		inputChannel     = in.GetChannel().To_InputChannel()
		channel          *channels.ChannelData
		participants     *channels.GetChannelParticipantsResp
		participantTypes []int32
	)

	channel, err = c.svcCtx.Dao.ChannelsClient.GetChannelDataById(c.ctx, &channels.ChannelDataByIdReq{ChannelId: inputChannel.GetChannelId()})
	if err != nil {
		return
	}

	if in.GetFilter() == nil {
		participantTypes = append(participantTypes, channels.K_ChannelParticipant)
	} else {
		//  Predicate_channelParticipantsRecent                          = "channelParticipantsRecent"
		//	Predicate_channelParticipantsAdmins                          = "channelParticipantsAdmins"
		//	Predicate_channelParticipantsKicked                          = "channelParticipantsKicked"
		//	Predicate_channelParticipantsBots                            = "channelParticipantsBots"
		//	Predicate_channelParticipantsBanned                          = "channelParticipantsBanned"
		//	Predicate_channelParticipantsSearch                          = "channelParticipantsSearch"
		//	Predicate_channelParticipantsContacts                        = "channelParticipantsContacts"
		//	Predicate_channelParticipantsMentions                        = "channelParticipantsMentions"
		//	Predicate_channels_channelParticipants    					 = "channels_channelParticipants"
		//	Predicate_channels_channelParticipantsNotModified             = "channels_channelParticipantsNotModified"
		switch in.GetFilter().PredicateName {
		case mtproto.Predicate_channelParticipantsAdmins:
			participantTypes = append(participantTypes, channels.K_ChannelParticipantAdmin, channels.K_ChannelParticipantCreator)
		case mtproto.Predicate_channelParticipantsKicked:
			participantTypes = append(participantTypes, channels.K_ChannelParticipantBanned, channels.K_ChannelParticipantKicked, channels.K_ChannelParticipantLeft)
		case mtproto.Predicate_channelParticipantsRecent:
			participantTypes = append(participantTypes, channels.K_ChannelParticipant)
		default:
			res = (&mtproto.Channels_ChannelParticipants{
				Count:        0,
				Participants: []*mtproto.ChannelParticipant{},
				Chats:        []*mtproto.Chat{},
				Users:        []*mtproto.User{},
			}).To_ChannelsChannelParticipants().To_Channels_ChannelParticipants()
			return
		}
	}

	participants, err = c.svcCtx.Dao.ChannelsClient.GetChannelParticipants(c.ctx,
		&channels.ChannelParticipantsReq{
			ChannelId: inputChannel.GetChannelId(),
			Types:     participantTypes,
		})
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

	return
}
