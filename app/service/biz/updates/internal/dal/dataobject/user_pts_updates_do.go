package dataobject

type UserPtsUpdatesDO struct {
	Id         int64  `db:"id"`
	UserId     int64  `db:"user_id"`
	Pts        int32  `db:"pts"`
	PtsCount   int32  `db:"pts_count"`
	UpdateType int32  `db:"update_type"`
	UpdateData string `db:"update_data"`
	Date2      int64  `db:"date2"`
}
