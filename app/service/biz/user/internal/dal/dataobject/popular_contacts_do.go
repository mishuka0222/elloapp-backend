package dataobject

type PopularContactsDO struct {
	Id        int64  `db:"id"`
	Phone     string `db:"phone"`
	Importers int32  `db:"importers"`
	Deleted   bool   `db:"deleted"`
	UpdateAt  string `db:"update_at"`
}
