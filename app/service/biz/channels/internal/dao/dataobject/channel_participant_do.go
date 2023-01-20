package dataobject

type ChannelParticipantDO struct {
	Id              int32  `db:"id"`
	ChannelId       int32  `db:"channel_id"`
	UserId          int32  `db:"user_id"`
	ParticipantType int8   `db:"participant_type"`
	InviterUserId   int32  `db:"inviter_user_id"`
	InvitedAt       int32  `db:"invited_at"`
	JoinedAt        int32  `db:"joined_at"`
	State           int8   `db:"state"`
	CreatedAt       string `db:"created_at"`
	UpdatedAt       string `db:"updated_at"`
}
