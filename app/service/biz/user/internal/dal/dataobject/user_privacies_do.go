package dataobject

type UserPrivaciesDO struct {
	Id      int64  `db:"id"`
	UserId  int64  `db:"user_id"`
	KeyType int32  `db:"key_type"`
	Rules   string `db:"rules"`
}
