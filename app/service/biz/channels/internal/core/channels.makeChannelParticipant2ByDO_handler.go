package core

import (
	"errors"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"

	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/channels/channels"
)

// channel_data
func (c *ChannelsCore) MakeChannelParticipant2ByDO(in *channels.MakeChannelParticipant2ByDOReq) (res *channels.ChannelParticipantRes, err error) {
	var (
		participant *mtproto.ChannelParticipant
	)

	participant = &mtproto.ChannelParticipant{
		UserId:          in.Participant.UserId,
		InviterId_INT64: in.Participant.InviterUserId,
		Date:            in.Participant.JoinedAt,
	}

	switch in.Participant.ParticipantType {
	case kChannelParticipant:
		if in.Participant.UserId == in.SelfId {
			participant.Constructor = mtproto.CRC32_channelParticipantSelf
		} else {
			participant.Constructor = mtproto.CRC32_channelParticipant
		}
	case kChannelParticipantCreator:
		participant.Constructor = mtproto.CRC32_channelParticipantCreator
	case kChannelParticipantAdmin:
		participant.Constructor = mtproto.CRC32_channelParticipantAdmin
	case kChannelParticipantBanned:
		participant.Constructor = mtproto.CRC32_channelParticipantBanned
	default:
		err = errors.New(" channelParticipant type error.")
		return
	}
	res = &channels.ChannelParticipantRes{Participant: participant}

	return
}
