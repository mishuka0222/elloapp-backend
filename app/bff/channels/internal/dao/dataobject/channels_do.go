package dataobject

type ChannelsDO struct {
	Id               int32  `db:"id"`
	CreatorUserId    int32  `db:"creator_user_id"`
	AccessHash       int64  `db:"access_hash"`
	RandomId         int64  `db:"random_id"`
	ParticipantCount int32  `db:"participant_count"`
	Title            string `db:"title"`
	About            string `db:"about"`
	PhotoId          int64  `db:"photo_id"`
	Link             string `db:"link"`
	AdminsEnabled    int8   `db:"admins_enabled"`
	Deactivated      int8   `db:"deactivated"`
	Version          int32  `db:"version"`
	Date             int32  `db:"date"`
	CreatedAt        string `db:"created_at"`
	UpdatedAt        string `db:"updated_at"`
}
