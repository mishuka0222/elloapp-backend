package dataobject

type UserProfilePhotosDO struct {
	Id      int64 `db:"id"`
	UserId  int64 `db:"user_id"`
	PhotoId int64 `db:"photo_id"`
	Date2   int64 `db:"date2"`
	Deleted bool  `db:"deleted"`
}
