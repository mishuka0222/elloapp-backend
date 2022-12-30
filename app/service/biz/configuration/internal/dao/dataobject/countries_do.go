package dataobject

type CountriesDO struct {
	Id   int64  `db:"id"`
	Code string `db:"code"`
	Name string `db:"name"`
	Flag string `db:"flag"`
}
