package sqlx

import (
	"context"
	"database/sql"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/pkg2/stores/sqlx/reflectx"
)

type conn struct {
	conn sqlx.SqlConn
	*reflectx.Mapper
}

func newConn(c sqlx.SqlConn) *conn {
	return &conn{
		conn:   c,
		Mapper: mapper(),
	}
}

func (db *conn) Exec(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return db.conn.ExecCtx(ctx, query, args...)
}

// NamedExec using this DB.
// Any named placeholder parameters are replaced with fields from arg.
func (db *conn) NamedExec(c context.Context, query string, arg interface{}) (sql.Result, error) {
	q, args, err := bindNamedMapper(QUESTION, query, arg, db.Mapper)
	if err != nil {
		return nil, err
	}

	return db.conn.ExecCtx(c, q, args...)
}

func (db *conn) Prepare(ctx context.Context, query string) (sqlx.StmtSession, error) {
	return db.conn.PrepareCtx(ctx, query)
}

// TODO:
//func (db *conn) NamedPrepare(ctx context.Context, query string) (sqlx.StmtSession, error) {
//	q, args, err := bindNamedMapper(QUESTION, query, arg, db.Mapper)
//	if err != nil {
//		return nil, err
//	}
//
//	return db.conn.PrepareCtx(ctx, query)
//}

func (db *conn) QueryRow(ctx context.Context, v interface{}, query string, args ...interface{}) error {
	return db.conn.QueryRowCtx(ctx, v, query, args...)
}

// NamedQueryRow using this DB.
// Any named placeholder parameters are replaced with fields from arg.
func (db *conn) NamedQueryRow(c context.Context, v interface{}, query string, arg interface{}) error {
	q, args, err := bindNamedMapper(QUESTION, query, arg, db.Mapper)
	if err != nil {
		return err
	}
	return db.QueryRow(c, v, q, args...)
}

func (db *conn) QueryRowPartial(ctx context.Context, v interface{}, query string, args ...interface{}) error {
	return db.conn.QueryRowPartialCtx(ctx, v, query, args...)
}

func (db *conn) NamedQueryRowPartial(c context.Context, v interface{}, query string, arg interface{}) error {
	q, args, err := bindNamedMapper(QUESTION, query, arg, db.Mapper)
	if err != nil {
		return err
	}
	return db.QueryRowPartial(c, v, q, args...)
}

func (db *conn) QueryRows(ctx context.Context, v interface{}, query string, args ...interface{}) error {
	return db.conn.QueryRowsCtx(ctx, v, query, args...)
}

func (db *conn) NamedQueryRows(c context.Context, v interface{}, query string, arg interface{}) error {
	q, args, err := bindNamedMapper(QUESTION, query, arg, db.Mapper)
	if err != nil {
		return err
	}
	return db.QueryRows(c, v, q, args...)
}

func (db *conn) QueryRowsPartial(ctx context.Context, v interface{}, query string, args ...interface{}) error {
	return db.conn.QueryRowsPartialCtx(ctx, v, query, args...)
}

func (db *conn) NamedQueryRowsPartial(c context.Context, v interface{}, query string, arg interface{}) error {
	q, args, err := bindNamedMapper(QUESTION, query, arg, db.Mapper)
	if err != nil {
		return err
	}
	return db.QueryRowsPartial(c, v, q, args...)
}

func (db *conn) RawDB() (*sql.DB, error) {
	return db.conn.RawDB()
}

func (db *conn) Transact(ctx context.Context, fn func(*Tx) error) error {
	return db.conn.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		tx := newTx(ctx, session)
		return fn(tx)
	})
}
