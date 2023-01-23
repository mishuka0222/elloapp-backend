package core

import (
	"context"
	"errors"
	"fmt"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/channels/channels"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/channels/internal/dao/dataobject"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/channels/internal/svc"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto/rpc/metadata"

	"github.com/zeromicro/go-zero/core/logx"
)

const (
	kChannelParticipant        = 0
	kChannelParticipantSelf    = 1
	kChannelParticipantCreator = 2
	kChannelParticipantAdmin   = 3
	kChannelParticipantBanned  = 4
)

type ChannelsCore struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	MD *metadata.RpcMetadata
}

func NewChannels(ctx context.Context, svcCtx *svc.ServiceContext) *ChannelsCore {
	return &ChannelsCore{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		MD:     metadata.RpcMetadataFromIncoming(ctx),
	}
}

func (c *ChannelsCore) checkOrLoadChannelParticipantList(in *channels.ChannelCoreData) (err error) {
	if len(in.Participants) == 0 {
		var (
			participants []dataobject.ChannelParticipantDO
		)
		participants, err = c.svcCtx.Dao.ChannelParticipantsDAO.SelectByChannelId(c.ctx, in.Channel.Id)
		if err != nil {
			return
		} else if len(participants) == 0 {
			return errors.New(fmt.Sprintf("can`t load participants list for this channel: %d", in.Channel.Id))
		}
		in.Participants = make([]*channels.ChannelParticipant, len(participants))
		for i := range participants {
			in.Participants[i] = participants[i].ToChannelParticipant()
		}
	}
	return
}

func makeChannelParticipantByDO(do *dataobject.ChannelParticipantDO) (participant *mtproto.ChannelParticipant, err error) {
	participant = &mtproto.ChannelParticipant{
		UserId:          do.UserId,
		InviterId_INT64: do.InviterUserId,
		Date:            do.JoinedAt,
	}

	switch do.ParticipantType {
	case kChannelParticipant:
		participant.Constructor = mtproto.CRC32_channelParticipant
	case kChannelParticipantCreator:
		participant.Constructor = mtproto.CRC32_channelParticipantCreator
	case kChannelParticipantAdmin:
		participant.Constructor = mtproto.CRC32_channelParticipantAdmin
	default:
		err = errors.New(" channelParticipant type error.")
	}

	return
}
