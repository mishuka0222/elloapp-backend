package mysql_dao

import (
	"context"
	"database/sql"

	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/updates/internal/dal/dataobject"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/pkg2/stores/sqlx"

	"github.com/zeromicro/go-zero/core/logx"
)

var _ *sql.Result

type UserPtsUpdatesDAO struct {
	db *sqlx.DB
}

func NewUserPtsUpdatesDAO(db *sqlx.DB) *UserPtsUpdatesDAO {
	return &UserPtsUpdatesDAO{db}
}

// Insert
// insert into user_pts_updates(user_id, pts, pts_count, update_type, update_data, date2) values (:user_id, :pts, :pts_count, :update_type, :update_data, :date2)
// TODO(@benqi): sqlmap
func (dao *UserPtsUpdatesDAO) Insert(ctx context.Context, do *dataobject.UserPtsUpdatesDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into user_pts_updates(user_id, pts, pts_count, update_type, update_data, date2) values (:user_id, :pts, :pts_count, :update_type, :update_data, :date2)"
		r     sql.Result
	)

	r, err = dao.db.NamedExec(ctx, query, do)
	if err != nil {
		logx.WithContext(ctx).Errorf("namedExec in Insert(%v), error: %v", do, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		logx.WithContext(ctx).Errorf("lastInsertId in Insert(%v)_error: %v", do, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in Insert(%v)_error: %v", do, err)
	}

	return
}

// InsertTx
// insert into user_pts_updates(user_id, pts, pts_count, update_type, update_data, date2) values (:user_id, :pts, :pts_count, :update_type, :update_data, :date2)
// TODO(@benqi): sqlmap
func (dao *UserPtsUpdatesDAO) InsertTx(tx *sqlx.Tx, do *dataobject.UserPtsUpdatesDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into user_pts_updates(user_id, pts, pts_count, update_type, update_data, date2) values (:user_id, :pts, :pts_count, :update_type, :update_data, :date2)"
		r     sql.Result
	)

	r, err = tx.NamedExec(query, do)
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("namedExec in Insert(%v), error: %v", do, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("lastInsertId in Insert(%v)_error: %v", do, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in Insert(%v)_error: %v", do, err)
	}

	return
}

// SelectLastPts
// select pts from user_pts_updates where user_id = :user_id order by pts desc limit 1
// TODO(@benqi): sqlmap
func (dao *UserPtsUpdatesDAO) SelectLastPts(ctx context.Context, user_id int64) (rValue *dataobject.UserPtsUpdatesDO, err error) {
	var (
		query = "select pts from user_pts_updates where user_id = ? order by pts desc limit 1"
		do    = &dataobject.UserPtsUpdatesDO{}
	)
	err = dao.db.QueryRowPartial(ctx, do, query, user_id)

	if err != nil {
		if err != sqlx.ErrNotFound {
			logx.WithContext(ctx).Errorf("queryx in SelectLastPts(_), error: %v", err)
			return
		} else {
			err = nil
		}
	} else {
		rValue = do
	}

	return
}

// SelectByGtPts
// select user_id, pts, pts_count, update_type, update_data from user_pts_updates where user_id = :user_id and pts > :pts order by pts asc limit :limit
// TODO(@benqi): sqlmap
func (dao *UserPtsUpdatesDAO) SelectByGtPts(ctx context.Context, user_id int64, pts int32, limit int32) (rList []dataobject.UserPtsUpdatesDO, err error) {
	var (
		query  = "select user_id, pts, pts_count, update_type, update_data from user_pts_updates where user_id = ? and pts > ? order by pts asc limit ?"
		values []dataobject.UserPtsUpdatesDO
	)
	err = dao.db.QueryRowsPartial(ctx, &values, query, user_id, pts, limit)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectByGtPts(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectByGtPtsWithCB
// select user_id, pts, pts_count, update_type, update_data from user_pts_updates where user_id = :user_id and pts > :pts order by pts asc limit :limit
// TODO(@benqi): sqlmap
func (dao *UserPtsUpdatesDAO) SelectByGtPtsWithCB(ctx context.Context, user_id int64, pts int32, limit int32, cb func(i int, v *dataobject.UserPtsUpdatesDO)) (rList []dataobject.UserPtsUpdatesDO, err error) {
	var (
		query  = "select user_id, pts, pts_count, update_type, update_data from user_pts_updates where user_id = ? and pts > ? order by pts asc limit ?"
		values []dataobject.UserPtsUpdatesDO
	)
	err = dao.db.QueryRowsPartial(ctx, &values, query, user_id, pts, limit)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectByGtPts(_), error: %v", err)
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
