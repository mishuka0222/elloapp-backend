package sqlx

import (
	"context"
	"database/sql"

	"github.com/go-sql-driver/mysql"
	"github.com/zeromicro/go-zero/core/logx"
)

var (
	ErrNotFound = sql.ErrNoRows
)

// Config mysql config.
type Config struct {
	DSN         string   // write data source name.
	ReadDSN     []string `json:",optional"` // read data source name.
	Active      int      `json:",optional"` // pool
	Idle        int      `json:",optional"` // pool
	IdleTimeout string   `json:",optional"` // connect max life time.
}

// NewMySQL new db and retry connection when has error.
func NewMySQL(c *Config) (db *DB) {
	db, err := Open(c)
	if err != nil {
		logx.Error("open mysql error(%v)", err)
		panic(err)
	}
	return
}

//// TxWrapper TxWrapper
//func TxWrapper(ctx context.Context, db *DB, fn func(*Tx) error) error {
//	return db.write.Transact(ctx, fn)
//}

// TxWrapper TxWrapper
func TxWrapper(ctx context.Context, db *DB, txF func(*Tx, *StoreResult)) *StoreResult {
	result := &StoreResult{}

	result.Err = db.write.Transact(ctx, func(tx *Tx) error {
		txF(tx, result)
		return result.Err
	})

	return result
}

// IsDuplicate
// Check if MySQL error is a Error Code: 1062. Duplicate entry ... for key ...
func IsDuplicate(err error) bool {
	if err == nil {
		return false
	}

	myerr, ok := err.(*mysql.MySQLError)
	return ok && myerr.Number == 1062
}

// IsMissingDb
// IsMissingDb
func IsMissingDb(err error) bool {
	if err == nil {
		return false
	}

	myerr, ok := err.(*mysql.MySQLError)
	return ok && myerr.Number == 1049
}
