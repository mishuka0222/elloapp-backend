package dataobject

type PredefinedUsersDO struct {
	Id               int64  `db:"id"`
	Phone            string `db:"phone"`
	FirstName        string `db:"first_name"`
	LastName         string `db:"last_name"`
	Username         string `db:"username"`
	Code             string `db:"code"`
	Verified         bool   `db:"verified"`
	RegisteredUserId int64  `db:"registered_user_id"`
	Deleted          bool   `db:"deleted"`
}
