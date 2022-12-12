package dataobject

type CountriesDO struct {
	Id   int64  `db:"id"`
	Code string `db:"country_code"`
	Name string `db:"country_name"`
}
