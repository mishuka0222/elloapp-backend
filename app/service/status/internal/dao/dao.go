package dao

type Dao struct {
}

func New() *Dao {
	return new(Dao)
}
