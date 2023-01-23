package dataobject

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
