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

func (c *ChannelsCore) checkOrLoadChannelParticipantList(in *channels.ChannelData) (err error) {
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

func (c *ChannelsCore) loadChannelParticipantDoList(channelId int64) (participants []dataobject.ChannelParticipantDO, err error) {

	participants, err = c.svcCtx.Dao.ChannelParticipantsDAO.SelectByChannelId(c.ctx, channelId)
	if err != nil {
		return
	} else if len(participants) == 0 {
		err = errors.New(fmt.Sprintf("can`t load participants list for this channel: %d", channelId))
	}

	return
}

func makeChannelParticipantByDO(do *dataobject.ChannelParticipantDO) (participant *mtproto.ChannelParticipant, err error) {
	participant = &mtproto.ChannelParticipant{
		UserId: do.UserId,
		//InviterId_INT64: do.InviterUserId,
		Date: do.JoinedAt,
	}

	switch do.ParticipantType {
	case channels.K_ChannelParticipant:
		participant.Constructor = mtproto.CRC32_channelParticipant
	case channels.K_ChannelParticipantCreator:
		participant.Constructor = mtproto.CRC32_channelParticipantCreator
		participant.AdminRights = MakeTLChatAdminRights()
	case channels.K_ChannelParticipantAdmin:
		participant.Constructor = mtproto.CRC32_channelParticipantAdmin
		participant.PromotedBy = do.InviterUserId
		participant.AdminRights = do.GetAdminRights()
	default:
		err = errors.New(" channelParticipant type error.")
	}

	return
}

//var defaultAdminRights = mtproto.MakeTLChatAdminRights(
//	&mtproto.ChatAdminRights{
//		ChangeInfo:     true,
//		PostMessages:   false,
//		EditMessages:   false,
//		DeleteMessages: true,
//		BanUsers:       true,
//		InviteUsers:    true,
//		PinMessages:    true,
//		AddAdmins:      false,
//		Anonymous:      false,
//		ManageCall:     true,
//		Other:          true,
//	}).To_ChatAdminRights()

func MakeTLChatAdminRights() *mtproto.ChatAdminRights {
	return mtproto.MakeTLChatAdminRights(&mtproto.ChatAdminRights{
		ChangeInfo:     true,
		PostMessages:   true, // default false
		EditMessages:   true,
		DeleteMessages: true,
		BanUsers:       true,
		InviteUsers:    true,
		PinMessages:    true,
		AddAdmins:      true,
		Anonymous:      true,
		ManageCall:     true,
		Other:          true,
	}).To_ChatAdminRights()
}

func ToChannelDO(ch *channels.Channel) *dataobject.ChannelDO {
	return &dataobject.ChannelDO{
		Id:               ch.Id,
		CreatorUserId:    ch.CreatorUserId,
		AccessHash:       ch.AccessHash,
		RandomId:         ch.RandomId,
		ParticipantCount: ch.ParticipantCount,
		Title:            ch.Title,
		About:            ch.About,
		PhotoId:          ch.PhotoId,
		Link:             ch.Link,
		Username:         ch.Username,
		AdminsEnabled:    ch.AdminsEnabled,
		Deactivated:      ch.Deactivated,
		Version:          ch.Version,
		Date:             ch.Date,
		CreatedAt:        ch.CreatedAt,
		UpdatedAt:        ch.UpdatedAt,
	}
}

func ToChannelParticipantDO(pnt *channels.ChannelParticipant) *dataobject.ChannelParticipantDO {
	return &dataobject.ChannelParticipantDO{
		Id:              pnt.Id,
		ChannelId:       pnt.ChannelId,
		UserId:          pnt.UserId,
		ParticipantType: pnt.ParticipantType,
		AdminRights:     pnt.AdminRightsToStr(),
		InviterUserId:   pnt.InviterUserId,
		InvitedAt:       pnt.InvitedAt,
		JoinedAt:        pnt.JoinedAt,
		LeftAt:          pnt.LeftAt,
		State:           pnt.State,
		CreatedAt:       pnt.CreatedAt,
		UpdatedAt:       pnt.UpdatedAt,
	}
}
