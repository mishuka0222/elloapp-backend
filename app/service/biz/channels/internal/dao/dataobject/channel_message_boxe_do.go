package dataobject

type ChannelMessageBoxDO struct {
	Id                  int64  `db:"id"`
	SenderUserId        int64  `db:"sender_user_id"`
	ChannelId           int64  `db:"channel_id"`
	ChannelMessageBoxId int64  `db:"channel_message_box_id"`
	MessageId           int64  `db:"message_id"`
	Date                int32  `db:"date"`
	Deleted             int8   `db:"deleted"`
	CreatedAt           string `db:"created_at"`
	UpdatedAt           string `db:"updated_at"`
}
