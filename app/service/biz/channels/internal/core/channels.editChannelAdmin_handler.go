package core

import (
	"errors"
	"github.com/zeromicro/go-zero/core/jsonx"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/channels/channels"
	"time"
)

func (c *ChannelsCore) EditChannelAdmin(in *channels.EditChannelAdminReq) (channel *channels.ChannelData, err error) {
	var (
		founded int
	)

	channel, err = c.GetChannelDataById(&channels.ChannelDataByIdReq{ChannelId: in.ChannelId})
	if err != nil {
		return
	}

	// TODO: remove or update
	// operatorId is creator
	if in.OperatorId != channel.Channel.CreatorUserId {
		err = errors.New("NO_EDIT_CHAT_PERMISSION")
		return
	}

	// editChatAdminId not creator
	if in.EditUserId == in.OperatorId {
		err = errors.New("BAD_REQUEST")
		return
	}

	err = c.checkOrLoadChannelParticipantList(channel)
	if err != nil {
		return
	}

	// check exists
	founded, _ = channel.FindChatParticipant(in.EditUserId)
	if founded == -1 { // add new user to chnnel
		//err = errors.New("PARTICIPANT_NOT_EXISTS")
		channel, err = c.AddChannelParticipant(&channels.AddChannelParticipantReq{
			Channel:     channel,
			InviterId:   in.OperatorId,
			UserId:      in.EditUserId,
			AdminRights: in.AdminRights,
		})
		if err != nil {
			return
		}
	} else { // update user access privileges
		var (
			isAdmin           bool
			updateParticipant bool
			strAdminRights    = "{}"
		)

		if channel.Participants[founded].ParticipantType == channels.K_ChannelParticipantAdmin {
			isAdmin = true
		}

		if !isAdmin && in.AdminRights != nil { // up to admin
			channel.Participants[founded].AdminRights = in.AdminRights
			channel.Participants[founded].ParticipantType = channels.K_ChannelParticipantAdmin
			strAdminRights, _ = jsonx.MarshalToString(in.AdminRights)
			updateParticipant = true
		} else if isAdmin && in.AdminRights == nil { // down to participant
			channel.Participants[founded].AdminRights = nil
			channel.Participants[founded].ParticipantType = channels.K_ChannelParticipant
			updateParticipant = true
		} else if isAdmin && in.AdminRights != nil { // update admin rules
			channel.Participants[founded].AdminRights = in.AdminRights
			strAdminRights, _ = jsonx.MarshalToString(in.AdminRights)
		} else if !isAdmin && in.AdminRights == nil { // not modified
			err = errors.New("CHAT_NOT_MODIFIED")
			return
		}

		if updateParticipant {
			_, err = c.svcCtx.Dao.ChannelParticipantsDAO.UpdateParticipantType(c.ctx, int8(channel.Participants[founded].ParticipantType),
				channels.K_ParticipantActiveState, channel.Participants[founded].Id)
			if err != nil {
				return
			}
		}

		_, err = c.svcCtx.Dao.ChannelParticipantsDAO.UpdateAdminRights(c.ctx, strAdminRights, channel.Participants[founded].Id)
		if err != nil {
			return
		}

		// update version
		channel.Channel.Date = int32(time.Now().Unix())
		channel.Channel.Version += 1
		_, err = c.svcCtx.Dao.ChannelsDAO.UpdateVersion(c.ctx, channel.Channel.Date, channel.GetChannelId())
		if err != nil {
			return
		}
	}

	return
}
