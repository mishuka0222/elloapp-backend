package mysql_dao

import (
	"context"
	"database/sql"
	"github.com/zeromicro/go-zero/core/logx"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/bff/channels/internal/dao/dataobject"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/pkg2/stores/sqlx"
)

var _ *sql.Result

type ChannelParticipantsDAO struct {
	db *sqlx.DB
}

func NewChannelParticipantsDAO(db *sqlx.DB) *ChannelParticipantsDAO {
	return &ChannelParticipantsDAO{db}
}

// Insert
// insert into channel_participants(channel_id, user_id, participant_type, inviter_user_id, invited_at, joined_at, state) values (:channel_id, :user_id, :participant_type, :inviter_user_id, :invited_at, :joined_at, :state)
func (dao *ChannelParticipantsDAO) Insert(ctx context.Context, do *dataobject.ChannelParticipantsDO) (id int64, err error) {
	var (
		query = "insert into channel_participants(channel_id, user_id, participant_type, inviter_user_id, invited_at, joined_at, state) values (:channel_id, :user_id, :participant_type, :inviter_user_id, :invited_at, :joined_at, :state)"
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

// SelectByChannelId
// select id, channel_id, user_id, participant_type, inviter_user_id, invited_at, joined_at, state from channel_participants where channel_id = :channel_id
func (dao *ChannelParticipantsDAO) SelectByChannelId(ctx context.Context, channel_id int32) (rValues []dataobject.ChannelParticipantsDO, err error) {
	var (
		query = "select id, channel_id, user_id, participant_type, inviter_user_id, invited_at, joined_at, state from channel_participants where channel_id = ?"
		rows  []dataobject.ChannelParticipantsDO
	)
	err = dao.db.QueryRows(ctx, &rows, query, channel_id)

	if err != nil {
		logx.WithContext(ctx).Errorf("Queryx in SelectByChannelId(_), error: %v", err)
		return
	}

	rValues = rows

	return
}

// SelectByUserIdList
// select id, channel_id, user_id, participant_type, inviter_user_id, invited_at, joined_at, state from channel_participants where channel_id = :channel_id and user_id in (:idList)
func (dao *ChannelParticipantsDAO) SelectByUserIdList(ctx context.Context, channel_id int32, idList []int32) (rValues []dataobject.ChannelParticipantsDO, err error) {
	var (
		q    = "select id, channel_id, user_id, participant_type, inviter_user_id, invited_at, joined_at, state from channel_participants where channel_id = ? and user_id in (?)"
		rows []dataobject.ChannelParticipantsDO
	)
	query, a, err := sqlx.In(q, channel_id, idList)
	err = dao.db.QueryRows(ctx, &rows, query, a...)

	if err != nil {
		logx.WithContext(ctx).Errorf("Queryx in SelectByUserIdList(_), error: %v", err)
		return
	}

	rValues = rows

	return
}

// SelectByUserId
// select id, channel_id, user_id, participant_type, inviter_user_id, invited_at, joined_at, state from channel_participants where channel_id = :channel_id and user_id = :user_id
func (dao *ChannelParticipantsDAO) SelectByUserId(ctx context.Context, channel_id int32, user_id int32) (rValue *dataobject.ChannelParticipantsDO, err error) {
	var (
		query = "select id, channel_id, user_id, participant_type, inviter_user_id, invited_at, joined_at, state from channel_participants where channel_id = ? and user_id = ?"
		row   *dataobject.ChannelParticipantsDO
	)

	err = dao.db.QueryRow(ctx, &row, query, channel_id, user_id)

	if err != nil {
		logx.WithContext(ctx).Errorf("Queryx in SelectByUserId(_), error: %v", err)
		return
	}

	rValue = row

	return
}

// DeleteChannelUser
// update channel_participants set state = 1 where id = :id
func (dao *ChannelParticipantsDAO) DeleteChannelUser(ctx context.Context, id int32) (i int64, err error) {
	var (
		query = "update channel_participants set state = 1 where id = ?"
		res   sql.Result
	)
	res, err = dao.db.Exec(ctx, query, id)

	if err != nil {
		logx.WithContext(ctx).Errorf("Exec in DeleteChannelUser(_), error: %v", err)
		return
	}

	i, err = res.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("RowsAffected in DeleteChannelUser(_), error: %v", err)
		return
	}

	return
}

// Update
// update channel_participants set inviter_user_id = :inviter_user_id, invited_at = :invited_at, joined_at = :joined_at, state = 0 where id = :id
func (dao *ChannelParticipantsDAO) Update(ctx context.Context, inviter_user_id int32, invited_at int32, joined_at int32, id int32) (i int64, err error) {
	var (
		query = "update channel_participants set inviter_user_id = ?, invited_at = ?, joined_at = ?, state = 0 where id = ?"
		res   sql.Result
	)
	res, err = dao.db.Exec(ctx, query, inviter_user_id, invited_at, joined_at, id)

	if err != nil {
		logx.WithContext(ctx).Errorf("Exec in Update(_), error: %v", err)
		return
	}

	i, err = res.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("RowsAffected in Update(_), error: %v", err)
		return
	}

	return
}

// UpdateParticipantType
// update channel_participants set participant_type = :participant_type where id = :id
func (dao *ChannelParticipantsDAO) UpdateParticipantType(ctx context.Context, participant_type int8, id int32) (i int64, err error) {
	var (
		query = "update channel_participants set participant_type = ? where id = ?"
		res   sql.Result
	)
	res, err = dao.db.Exec(ctx, query, participant_type, id)

	if err != nil {
		logx.WithContext(ctx).Errorf("Exec in UpdateParticipantType(_), error: %v", err)
		return
	}

	i, err = res.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("RowsAffected in UpdateParticipantType(_), error: %v", err)
		return
	}

	return
}
