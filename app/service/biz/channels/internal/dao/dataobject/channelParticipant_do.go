package dataobject

import (
	"github.com/zeromicro/go-zero/core/jsonx"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/channels/channels"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

type ChannelParticipantDO struct {
	Id              int64  `db:"id"`
	ChannelId       int64  `db:"channel_id"`
	UserId          int64  `db:"user_id"`
	ParticipantType int32  `db:"participant_type"`
	AdminRights     string `db:"admin_rights"`
	BannedRights    string `db:"banned_rights"`
	InviterUserId   int64  `db:"inviter_user_id"`
	KickedBy        int64  `db:"kicked_by"`
	InvitedAt       int32  `db:"invited_at"`
	JoinedAt        int32  `db:"joined_at"`
	LeftAt          int32  `db:"left_at"`
	State           int32  `db:"state"`
	CreatedAt       string `db:"created_at"`
	UpdatedAt       string `db:"updated_at"`
}

func (pnt *ChannelParticipantDO) ToChannelParticipant() (channel *channels.ChannelParticipant) {
	channel = &channels.ChannelParticipant{
		Id:              pnt.Id,
		ChannelId:       pnt.ChannelId,
		UserId:          pnt.UserId,
		ParticipantType: pnt.ParticipantType,
		InviterUserId:   pnt.InviterUserId,
		KickedBy:        pnt.KickedBy,
		InvitedAt:       pnt.InvitedAt,
		JoinedAt:        pnt.JoinedAt,
		LeftAt:          pnt.LeftAt,
		State:           pnt.State,
		CreatedAt:       pnt.CreatedAt,
		UpdatedAt:       pnt.UpdatedAt,
	}
	if pnt.AdminRights != "" {
		var (
			adminRights mtproto.ChatAdminRights
			err         error
		)
		if err = jsonx.UnmarshalFromString(pnt.AdminRights, &adminRights); err != nil {
			return
		}
		channel.AdminRights = &adminRights
	}
	if pnt.BannedRights != "" {
		var (
			bannedRights mtproto.ChatBannedRights
			err          error
		)
		if err = jsonx.UnmarshalFromString(pnt.BannedRights, &bannedRights); err != nil {
			return
		}
		channel.BannedRights = &bannedRights
	}

	return
}

func (pnt *ChannelParticipantDO) GetAdminRights() (res *mtproto.ChatAdminRights) {
	if pnt.AdminRights != "" {
		var (
			adminRights mtproto.ChatAdminRights
			err         error
		)
		if err = jsonx.UnmarshalFromString(pnt.AdminRights, &adminRights); err != nil {
			return
		}
		res = &adminRights
	}
	return
}

func (pnt *ChannelParticipantDO) GetBannedRights() (res *mtproto.ChatBannedRights) {
	if pnt.BannedRights != "" {
		var (
			bannedRights mtproto.ChatBannedRights
			err          error
		)
		if err = jsonx.UnmarshalFromString(pnt.BannedRights, &bannedRights); err != nil {
			return
		}
		res = &bannedRights
	}
	return
}
