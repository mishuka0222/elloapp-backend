package mysql_dao

import (
	"context"
	"database/sql"
	"github.com/zeromicro/go-zero/core/logx"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/bff/channels/internal/dao/dataobject"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/pkg2/stores/sqlx"
)

var _ *sql.Result

type ChannelPtsUpdatesDAO struct {
	db *sqlx.DB
}

func NewChannelPtsUpdatesDAO(db *sqlx.DB) *ChannelPtsUpdatesDAO {
	return &ChannelPtsUpdatesDAO{db}
}

// Insert
// insert into channel_pts_updates(channel_id, pts, pts_count, update_type, update_data, date2) values (:channel_id, :pts, :pts_count, :update_type, :update_data, :date2)
func (dao *ChannelPtsUpdatesDAO) Insert(ctx context.Context, do *dataobject.ChannelPtsUpdateDO) (id int64, err error) {
	var (
		query = "insert into channel_pts_updates(channel_id, pts, pts_count, update_type, update_data, date2) values (:channel_id, :pts, :pts_count, :update_type, :update_data, :date2)"
		res   sql.Result
	)
	res, err = dao.db.NamedExec(ctx, query, do)
	if err != nil {
		logx.WithContext(ctx).Errorf("NamedExec in Insert(%v), error: %v", do, err)
		return
	}

	id, err = res.LastInsertId()
	if err != nil {
		logx.WithContext(ctx).Errorf("LastInsertId in Insert(%v)_error: %v", do, err)
		return
	}
	return
}

// SelectLastPts
// select pts from channel_pts_updates where channel_id = :channel_id order by pts desc limit 1
func (dao *ChannelPtsUpdatesDAO) SelectLastPts(ctx context.Context, channel_id int32) (rValue *dataobject.ChannelPtsUpdateDO, err error) {
	var (
		query = "select pts from channel_pts_updates where channel_id = ? order by pts desc limit 1"
		row   *dataobject.ChannelPtsUpdateDO
	)
	err = dao.db.QueryRow(ctx, &row, query, channel_id)

	if err != nil {
		logx.WithContext(ctx).Errorf("Queryx in SelectLastPts(_), error: %v", err)
		return
	}

	rValue = row

	return
}

// SelectByGtPts
// select channel_id, pts, pts_count, update_type, update_data from channel_pts_updates where channel_id = :channel_id and pts > :pts order by pts asc
func (dao *ChannelPtsUpdatesDAO) SelectByGtPts(ctx context.Context, channel_id int32, pts int32) (rValues []dataobject.ChannelPtsUpdateDO, err error) {
	var (
		query = "select channel_id, pts, pts_count, update_type, update_data from channel_pts_updates where channel_id = ? and pts > ? order by pts asc"
		rows  []dataobject.ChannelPtsUpdateDO
	)
	err = dao.db.QueryRows(ctx, &rows, query, channel_id, pts)

	if err != nil {
		logx.WithContext(ctx).Errorf("Queryx in SelectByGtPts(_), error: %v", err)
		return
	}

	rValues = rows

	return
}
