package dataobject

import "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/channels/channels"

type ChannelParticipantDO struct {
	Id              int64  `db:"id"`
	ChannelId       int64  `db:"channel_id"`
	UserId          int64  `db:"user_id"`
	ParticipantType int8   `db:"participant_type"`
	InviterUserId   int64  `db:"inviter_user_id"`
	InvitedAt       int32  `db:"invited_at"`
	JoinedAt        int32  `db:"joined_at"`
	State           int8   `db:"state"`
	CreatedAt       string `db:"created_at"`
	UpdatedAt       string `db:"updated_at"`
}

func (pnt *ChannelParticipantDO) ToChannelParticipant() *channels.ChannelParticipant {
	return &channels.ChannelParticipant{
		Id:              pnt.Id,
		ChannelId:       pnt.ChannelId,
		UserId:          pnt.UserId,
		ParticipantType: int32(pnt.ParticipantType),
		InviterUserId:   pnt.InviterUserId,
		InvitedAt:       pnt.InvitedAt,
		JoinedAt:        pnt.JoinedAt,
		State:           int32(pnt.State),
		CreatedAt:       pnt.CreatedAt,
		UpdatedAt:       pnt.UpdatedAt,
	}
}
