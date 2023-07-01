package dataobject

type AuthKeysDO struct {
	Id        int64  `db:"id"`
	AuthKeyId int64  `db:"auth_key_id"`
	Body      string `db:"body"`
	Deleted   bool   `db:"deleted"`
}
