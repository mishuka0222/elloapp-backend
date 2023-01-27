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
	AdminRights     string `json:"admin_rights"`
	InviterUserId   int64  `db:"inviter_user_id"`
	InvitedAt       int32  `db:"invited_at"`
	KickedAt        int32  `db:"kicked_at"`
	LeftAt          int32  `db:"left_at"`
	JoinedAt        int32  `db:"joined_at"`
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
		InvitedAt:       pnt.InvitedAt,
		KickedAt:        pnt.KickedAt,
		LeftAt:          pnt.LeftAt,
		JoinedAt:        pnt.JoinedAt,
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
