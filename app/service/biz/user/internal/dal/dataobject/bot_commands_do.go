package dataobject

type BotCommandsDO struct {
	Id          int64  `db:"id"`
	BotId       int64  `db:"bot_id"`
	Command     string `db:"command"`
	Description string `db:"description"`
}
