package mysql_dao

import (
	"context"
	"database/sql"

	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/media/internal/dal/dataobject"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/pkg2/stores/sqlx"

	"github.com/zeromicro/go-zero/core/logx"
)

var _ *sql.Result

type VideoSizesDAO struct {
	db *sqlx.DB
}

func NewVideoSizesDAO(db *sqlx.DB) *VideoSizesDAO {
	return &VideoSizesDAO{db}
}

// Insert
// insert into video_sizes(video_size_id, size_type, width, height, file_size, video_start_ts, file_path) values (:video_size_id, :size_type, :width, :height, :file_size, :video_start_ts, :file_path)
// TODO(@benqi): sqlmap
func (dao *VideoSizesDAO) Insert(ctx context.Context, do *dataobject.VideoSizesDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into video_sizes(video_size_id, size_type, width, height, file_size, video_start_ts, file_path) values (:video_size_id, :size_type, :width, :height, :file_size, :video_start_ts, :file_path)"
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
// insert into video_sizes(video_size_id, size_type, width, height, file_size, video_start_ts, file_path) values (:video_size_id, :size_type, :width, :height, :file_size, :video_start_ts, :file_path)
// TODO(@benqi): sqlmap
func (dao *VideoSizesDAO) InsertTx(tx *sqlx.Tx, do *dataobject.VideoSizesDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into video_sizes(video_size_id, size_type, width, height, file_size, video_start_ts, file_path) values (:video_size_id, :size_type, :width, :height, :file_size, :video_start_ts, :file_path)"
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

// SelectListByVideoSizeId
// select id, video_size_id, size_type, width, height, file_size, video_start_ts, file_path from video_sizes where video_size_id = :video_size_id order by id asc
// TODO(@benqi): sqlmap
func (dao *VideoSizesDAO) SelectListByVideoSizeId(ctx context.Context, video_size_id int64) (rList []dataobject.VideoSizesDO, err error) {
	var (
		query  = "select id, video_size_id, size_type, width, height, file_size, video_start_ts, file_path from video_sizes where video_size_id = ? order by id asc"
		values []dataobject.VideoSizesDO
	)
	err = dao.db.QueryRowsPartial(ctx, &values, query, video_size_id)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectListByVideoSizeId(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectListByVideoSizeIdWithCB
// select id, video_size_id, size_type, width, height, file_size, video_start_ts, file_path from video_sizes where video_size_id = :video_size_id order by id asc
// TODO(@benqi): sqlmap
func (dao *VideoSizesDAO) SelectListByVideoSizeIdWithCB(ctx context.Context, video_size_id int64, cb func(i int, v *dataobject.VideoSizesDO)) (rList []dataobject.VideoSizesDO, err error) {
	var (
		query  = "select id, video_size_id, size_type, width, height, file_size, video_start_ts, file_path from video_sizes where video_size_id = ? order by id asc"
		values []dataobject.VideoSizesDO
	)
	err = dao.db.QueryRowsPartial(ctx, &values, query, video_size_id)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectListByVideoSizeId(_), error: %v", err)
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

// SelectListByVideoSizeIdList
// select id, video_size_id, size_type, width, height, file_size, video_start_ts, file_path from video_sizes where video_size_id in (:idList) order by id asc
// TODO(@benqi): sqlmap
func (dao *VideoSizesDAO) SelectListByVideoSizeIdList(ctx context.Context, idList []int64) (rList []dataobject.VideoSizesDO, err error) {
	var (
		query  = "select id, video_size_id, size_type, width, height, file_size, video_start_ts, file_path from video_sizes where video_size_id in (?) order by id asc"
		a      []interface{}
		values []dataobject.VideoSizesDO
	)
	if len(idList) == 0 {
		rList = []dataobject.VideoSizesDO{}
		return
	}

	query, a, err = sqlx.In(query, idList)
	if err != nil {
		// r sql.Result
		logx.WithContext(ctx).Errorf("sqlx.In in SelectListByVideoSizeIdList(_), error: %v", err)
		return
	}
	err = dao.db.QueryRowsPartial(ctx, &values, query, a...)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectListByVideoSizeIdList(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectListByVideoSizeIdListWithCB
// select id, video_size_id, size_type, width, height, file_size, video_start_ts, file_path from video_sizes where video_size_id in (:idList) order by id asc
// TODO(@benqi): sqlmap
func (dao *VideoSizesDAO) SelectListByVideoSizeIdListWithCB(ctx context.Context, idList []int64, cb func(i int, v *dataobject.VideoSizesDO)) (rList []dataobject.VideoSizesDO, err error) {
	var (
		query  = "select id, video_size_id, size_type, width, height, file_size, video_start_ts, file_path from video_sizes where video_size_id in (?) order by id asc"
		a      []interface{}
		values []dataobject.VideoSizesDO
	)
	if len(idList) == 0 {
		rList = []dataobject.VideoSizesDO{}
		return
	}

	query, a, err = sqlx.In(query, idList)
	if err != nil {
		// r sql.Result
		logx.WithContext(ctx).Errorf("sqlx.In in SelectListByVideoSizeIdList(_), error: %v", err)
		return
	}
	err = dao.db.QueryRowsPartial(ctx, &values, query, a...)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectListByVideoSizeIdList(_), error: %v", err)
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
