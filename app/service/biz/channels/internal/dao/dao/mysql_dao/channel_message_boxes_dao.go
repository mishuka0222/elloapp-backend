package mysql_dao

import (
	"context"
	"database/sql"
	"github.com/zeromicro/go-zero/core/logx"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/channels/internal/dao/dataobject"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/pkg2/stores/sqlx"
)

var _ *sql.Result

type ChannelMessageBoxesDAO struct {
	db *sqlx.DB
}

func NewChannelMessageBoxesDAO(db *sqlx.DB) *ChannelMessageBoxesDAO {
	return &ChannelMessageBoxesDAO{db}
}

// Insert
// insert into channel_message_boxes(sender_user_id, channel_id, channel_message_box_id, message_id, `date`) values (:sender_user_id, :channel_id, :channel_message_box_id, :message_id, :date)
func (dao *ChannelMessageBoxesDAO) Insert(ctx context.Context, do *dataobject.ChannelMessageBoxDO) (id int64, err error) {
	var (
		query = "insert into channel_message_boxes(sender_user_id, channel_id, channel_message_box_id, message_id, `date`) values (:sender_user_id, :channel_id, :channel_message_box_id, :message_id, :date)"
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

// SelectByMessageIdList
// select sender_user_id, channel_id, channel_message_box_id, message_id, `date` from channel_message_boxes where channel_id = :channel_id and deleted = 0 and channel_message_box_id in (:idList) order by channel_message_box_id desc
func (dao *ChannelMessageBoxesDAO) SelectByMessageIdList(ctx context.Context, channel_id int64, idList []int64) (rValues []dataobject.ChannelMessageBoxDO, err error) {
	var (
		q    = "select sender_user_id, channel_id, channel_message_box_id, message_id, `date` from channel_message_boxes where channel_id = ? and deleted = 0 and channel_message_box_id in (?) order by channel_message_box_id desc"
		rows []dataobject.ChannelMessageBoxDO
	)
	query, a, err := sqlx.In(q, channel_id, idList)
	err = dao.db.QueryRows(ctx, &rows, query, a...)

	if err != nil {
		logx.WithContext(ctx).Errorf("Queryx in SelectByMessageIdList(_), error: %v", err)
		return
	}

	rValues = rows

	return
}

// SelectByMessageId
// select sender_user_id, channel_id, channel_message_box_id, message_id, `date` from channel_message_boxes where channel_id = :channel_id and channel_message_box_id = :channel_message_box_id and deleted = 0 limit 1
func (dao *ChannelMessageBoxesDAO) SelectByMessageId(ctx context.Context, channel_id int64, channel_message_box_id int64) (rValue *dataobject.ChannelMessageBoxDO, err error) {
	var (
		query = "select sender_user_id, channel_id, channel_message_box_id, message_id, `date` from channel_message_boxes where channel_id = ? and channel_message_box_id = ? and deleted = 0 limit 1"
		row   *dataobject.ChannelMessageBoxDO
	)
	err = dao.db.QueryRow(ctx, &row, query, channel_id, channel_message_box_id)

	if err != nil {
		logx.WithContext(ctx).Errorf("Queryx in SelectByMessageId(_), error: %v", err)
		return
	}

	rValue = row

	return
}

// SelectBackwardByOffsetLimit
// select sender_user_id, channel_id, channel_message_box_id, message_id, `date` from channel_message_boxes where channel_id = :channel_id and channel_message_box_id < :channel_message_box_id and deleted = 0 order by channel_message_box_id desc limit :limit
func (dao *ChannelMessageBoxesDAO) SelectBackwardByOffsetLimit(ctx context.Context, channel_id int64, channel_message_box_id int64, limit int32) (rValues []dataobject.ChannelMessageBoxDO, err error) {
	var (
		query = "select sender_user_id, channel_id, channel_message_box_id, message_id, `date` from channel_message_boxes where channel_id = ? and channel_message_box_id < ? and deleted = 0 order by channel_message_box_id desc limit ?"
		rows  []dataobject.ChannelMessageBoxDO
	)
	err = dao.db.QueryRows(ctx, &rows, query, channel_id, channel_message_box_id, limit)

	if err != nil {
		logx.WithContext(ctx).Errorf("Queryx in SelectBackwardByOffsetLimit(_), error: %v", err)
		return
	}

	rValues = rows

	return
}

// SelectForwardByOffsetLimit
// select sender_user_id, channel_id, channel_message_box_id, message_id, `date` from channel_message_boxes where channel_id = :channel_id and channel_message_box_id >= :channel_message_box_id and deleted = 0 order by channel_message_box_id asc limit :limit
func (dao *ChannelMessageBoxesDAO) SelectForwardByOffsetLimit(ctx context.Context, channel_id int64, channel_message_box_id int64, limit int32) (rValues []dataobject.ChannelMessageBoxDO, err error) {
	var (
		query = "select sender_user_id, channel_id, channel_message_box_id, message_id, `date` from channel_message_boxes where channel_id = ? and channel_message_box_id >= ? and deleted = 0 order by channel_message_box_id asc limit ?"
		rows  []dataobject.ChannelMessageBoxDO
	)

	err = dao.db.QueryRows(ctx, &rows, query, channel_id, channel_message_box_id, limit)

	if err != nil {
		logx.WithContext(ctx).Errorf("Queryx in SelectForwardByOffsetLimit(_), error: %v", err)
		return
	}

	rValues = rows

	return
}
