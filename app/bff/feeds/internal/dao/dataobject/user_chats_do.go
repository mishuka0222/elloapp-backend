package dataobject

type UserChatDO struct {
	ChatID  int64  `db:"chat_id" json:"chat_id"`
	PhotoID int64  `db:"photo_id" json:"photo_id"`
	Title   string `db:"title" json:"title"`
	Status  bool   `json:"status"`
}
