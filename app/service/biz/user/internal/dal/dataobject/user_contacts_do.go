package dataobject

type UserContactsDO struct {
	Id               int64  `db:"id"`
	OwnerUserId      int64  `db:"owner_user_id"`
	ContactUserId    int64  `db:"contact_user_id"`
	ContactPhone     string `db:"contact_phone"`
	ContactFirstName string `db:"contact_first_name"`
	ContactLastName  string `db:"contact_last_name"`
	Mutual           bool   `db:"mutual"`
	IsDeleted        bool   `db:"is_deleted"`
	Date2            int64  `db:"date2"`
}
