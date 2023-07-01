package dataobject

type UserSettingsDO struct {
	Id      int64  `db:"id"`
	UserId  int64  `db:"user_id"`
	Key2    string `db:"key2"`
	Value   string `db:"value"`
	Deleted bool   `db:"deleted"`
}
