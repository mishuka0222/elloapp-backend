package dataobject

type UserPresencesDO struct {
	Id         int64 `db:"id"`
	UserId     int64 `db:"user_id"`
	LastSeenAt int64 `db:"last_seen_at"`
	Expires    int32 `db:"expires"`
}
