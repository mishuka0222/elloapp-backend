package dataobject

type ReadOutboxDO struct {
	ChatID          int64 `db:"chat_id" json:"chat_id"`
	TopMessage      int64 `db:"top_message" json:"top_message"`
	ReadOutboxMaxID int64 `db:"read_outbox_max_id" json:"read_outbox_max_id"`
}
