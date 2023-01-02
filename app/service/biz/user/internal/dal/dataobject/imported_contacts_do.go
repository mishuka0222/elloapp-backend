package dataobject

type ImportedContactsDO struct {
	Id             int64 `db:"id"`
	UserId         int64 `db:"user_id"`
	ImportedUserId int64 `db:"imported_user_id"`
	Deleted        bool  `db:"deleted"`
}
