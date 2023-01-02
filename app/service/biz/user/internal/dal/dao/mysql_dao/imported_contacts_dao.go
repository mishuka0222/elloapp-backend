package mysql_dao

import (
	"context"
	"database/sql"

	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/user/internal/dal/dataobject"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/pkg2/stores/sqlx"

	"github.com/zeromicro/go-zero/core/logx"
)

var _ *sql.Result

type ImportedContactsDAO struct {
	db *sqlx.DB
}

func NewImportedContactsDAO(db *sqlx.DB) *ImportedContactsDAO {
	return &ImportedContactsDAO{db}
}

// InsertOrUpdate
// insert into imported_contacts(user_id, imported_user_id) values (:user_id, :imported_user_id) on duplicate key update deleted = 0
// TODO(@benqi): sqlmap
func (dao *ImportedContactsDAO) InsertOrUpdate(ctx context.Context, do *dataobject.ImportedContactsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into imported_contacts(user_id, imported_user_id) values (:user_id, :imported_user_id) on duplicate key update deleted = 0"
		r     sql.Result
	)

	r, err = dao.db.NamedExec(ctx, query, do)
	if err != nil {
		logx.WithContext(ctx).Errorf("namedExec in InsertOrUpdate(%v), error: %v", do, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		logx.WithContext(ctx).Errorf("lastInsertId in InsertOrUpdate(%v)_error: %v", do, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in InsertOrUpdate(%v)_error: %v", do, err)
	}

	return
}

// InsertOrUpdateTx
// insert into imported_contacts(user_id, imported_user_id) values (:user_id, :imported_user_id) on duplicate key update deleted = 0
// TODO(@benqi): sqlmap
func (dao *ImportedContactsDAO) InsertOrUpdateTx(tx *sqlx.Tx, do *dataobject.ImportedContactsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into imported_contacts(user_id, imported_user_id) values (:user_id, :imported_user_id) on duplicate key update deleted = 0"
		r     sql.Result
	)

	r, err = tx.NamedExec(query, do)
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("namedExec in InsertOrUpdate(%v), error: %v", do, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("lastInsertId in InsertOrUpdate(%v)_error: %v", do, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in InsertOrUpdate(%v)_error: %v", do, err)
	}

	return
}

// SelectList
// select id, user_id, imported_user_id from imported_contacts where user_id = :user_id and deleted = 0
// TODO(@benqi): sqlmap
func (dao *ImportedContactsDAO) SelectList(ctx context.Context, user_id int64) (rList []dataobject.ImportedContactsDO, err error) {
	var (
		query  = "select id, user_id, imported_user_id from imported_contacts where user_id = ? and deleted = 0"
		values []dataobject.ImportedContactsDO
	)
	err = dao.db.QueryRowsPartial(ctx, &values, query, user_id)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectList(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectListWithCB
// select id, user_id, imported_user_id from imported_contacts where user_id = :user_id and deleted = 0
// TODO(@benqi): sqlmap
func (dao *ImportedContactsDAO) SelectListWithCB(ctx context.Context, user_id int64, cb func(i int, v *dataobject.ImportedContactsDO)) (rList []dataobject.ImportedContactsDO, err error) {
	var (
		query  = "select id, user_id, imported_user_id from imported_contacts where user_id = ? and deleted = 0"
		values []dataobject.ImportedContactsDO
	)
	err = dao.db.QueryRowsPartial(ctx, &values, query, user_id)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectList(_), error: %v", err)
		return
	}

	rList = values

	if cb != nil {
		for i := 0; i < len(rList); i++ {
			cb(i, &rList[i])
		}
	}

	return
}

// SelectListByImportedList
// select id, user_id, imported_user_id from imported_contacts where user_id = :user_id and deleted = 0 and imported_user_id in (:idList)
// TODO(@benqi): sqlmap
func (dao *ImportedContactsDAO) SelectListByImportedList(ctx context.Context, user_id int64, idList []int64) (rList []dataobject.ImportedContactsDO, err error) {
	var (
		query  = "select id, user_id, imported_user_id from imported_contacts where user_id = ? and deleted = 0 and imported_user_id in (?)"
		a      []interface{}
		values []dataobject.ImportedContactsDO
	)

	if len(idList) == 0 {
		rList = []dataobject.ImportedContactsDO{}
		return
	}

	query, a, err = sqlx.In(query, user_id, idList)
	if err != nil {
		// r sql.Result
		logx.WithContext(ctx).Errorf("sqlx.In in SelectListByImportedList(_), error: %v", err)
		return
	}
	err = dao.db.QueryRowsPartial(ctx, &values, query, a...)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectListByImportedList(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectListByImportedListWithCB
// select id, user_id, imported_user_id from imported_contacts where user_id = :user_id and deleted = 0 and imported_user_id in (:idList)
// TODO(@benqi): sqlmap
func (dao *ImportedContactsDAO) SelectListByImportedListWithCB(ctx context.Context, user_id int64, idList []int64, cb func(i int, v *dataobject.ImportedContactsDO)) (rList []dataobject.ImportedContactsDO, err error) {
	var (
		query  = "select id, user_id, imported_user_id from imported_contacts where user_id = ? and deleted = 0 and imported_user_id in (?)"
		a      []interface{}
		values []dataobject.ImportedContactsDO
	)

	if len(idList) == 0 {
		rList = []dataobject.ImportedContactsDO{}
		return
	}

	query, a, err = sqlx.In(query, user_id, idList)
	if err != nil {
		// r sql.Result
		logx.WithContext(ctx).Errorf("sqlx.In in SelectListByImportedList(_), error: %v", err)
		return
	}
	err = dao.db.QueryRowsPartial(ctx, &values, query, a...)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectListByImportedList(_), error: %v", err)
		return
	}

	rList = values

	if cb != nil {
		for i := 0; i < len(rList); i++ {
			cb(i, &rList[i])
		}
	}

	return
}

// SelectAllList
// select id, user_id, imported_user_id from imported_contacts where user_id = :user_id
// TODO(@benqi): sqlmap
func (dao *ImportedContactsDAO) SelectAllList(ctx context.Context, user_id int64) (rList []dataobject.ImportedContactsDO, err error) {
	var (
		query  = "select id, user_id, imported_user_id from imported_contacts where user_id = ?"
		values []dataobject.ImportedContactsDO
	)
	err = dao.db.QueryRowsPartial(ctx, &values, query, user_id)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectAllList(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectAllListWithCB
// select id, user_id, imported_user_id from imported_contacts where user_id = :user_id
// TODO(@benqi): sqlmap
func (dao *ImportedContactsDAO) SelectAllListWithCB(ctx context.Context, user_id int64, cb func(i int, v *dataobject.ImportedContactsDO)) (rList []dataobject.ImportedContactsDO, err error) {
	var (
		query  = "select id, user_id, imported_user_id from imported_contacts where user_id = ?"
		values []dataobject.ImportedContactsDO
	)
	err = dao.db.QueryRowsPartial(ctx, &values, query, user_id)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectAllList(_), error: %v", err)
		return
	}

	rList = values

	if cb != nil {
		for i := 0; i < len(rList); i++ {
			cb(i, &rList[i])
		}
	}

	return
}

// Delete
// update imported_contacts set deleted = 1 where user_id = :user_id and imported_user_id = :imported_user_id
// TODO(@benqi): sqlmap
func (dao *ImportedContactsDAO) Delete(ctx context.Context, user_id int64, imported_user_id int64) (rowsAffected int64, err error) {
	var (
		query   = "update imported_contacts set deleted = 1 where user_id = ? and imported_user_id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, user_id, imported_user_id)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in Delete(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in Delete(_), error: %v", err)
	}

	return
}

// update imported_contacts set deleted = 1 where user_id = :user_id and imported_user_id = :imported_user_id
// DeleteTx
// TODO(@benqi): sqlmap
func (dao *ImportedContactsDAO) DeleteTx(tx *sqlx.Tx, user_id int64, imported_user_id int64) (rowsAffected int64, err error) {
	var (
		query   = "update imported_contacts set deleted = 1 where user_id = ? and imported_user_id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, user_id, imported_user_id)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in Delete(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in Delete(_), error: %v", err)
	}

	return
}
