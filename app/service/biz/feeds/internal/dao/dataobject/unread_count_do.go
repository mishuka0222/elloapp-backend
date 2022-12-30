package dataobject

type UnreadCountDo struct {
	ChatID int64 `db:"chat_id" json:"chat_id"`
	Unread int32 `db:"unread" json:"unread"`
}
