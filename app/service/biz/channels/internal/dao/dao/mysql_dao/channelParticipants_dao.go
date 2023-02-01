package mysql_dao

import (
	"context"
	"database/sql"
	"github.com/zeromicro/go-zero/core/logx"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/channels/internal/dao/dataobject"
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
func (dao *ChannelParticipantsDAO) Insert(ctx context.Context, do *dataobject.ChannelParticipantDO) (id int64, err error) {
	var (
		query = "insert into channel_participants(channel_id, user_id, participant_type, admin_rights, banned_rights, inviter_user_id, invited_at, joined_at, state) values (:channel_id, :user_id, :participant_type, :admin_rights, :banned_rights, :inviter_user_id, :invited_at, :joined_at, :state) on duplicate key update id = last_insert_id(id)"
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
func (dao *ChannelParticipantsDAO) SelectByChannelId(ctx context.Context, channel_id int64) (rValues []dataobject.ChannelParticipantDO, err error) {
	var (
		query = "select id, channel_id, user_id, participant_type, admin_rights, inviter_user_id, invited_at, joined_at, left_at, state from channel_participants where channel_id = ?"
		rows  []dataobject.ChannelParticipantDO
	)
	err = dao.db.QueryRowsPartial(ctx, &rows, query, channel_id)

	if err != nil {
		logx.WithContext(ctx).Errorf("Queryx in SelectByChannelId(_), error: %v", err)
		return
	}

	rValues = rows

	return
}

// SelectByUserIdList
// select id, channel_id, user_id, participant_type, inviter_user_id, invited_at, joined_at, state from channel_participants where channel_id = :channel_id and user_id in (:idList)
func (dao *ChannelParticipantsDAO) SelectByUserIdList(ctx context.Context, channel_id int64, idList []int64) (rValues []dataobject.ChannelParticipantDO, err error) {
	var (
		q    = "select id, channel_id, user_id, participant_type, admin_rights, banned_rights, inviter_user_id, invited_at, joined_at, left_at, state from channel_participants where channel_id = ? and user_id in (?)"
		rows []dataobject.ChannelParticipantDO
	)
	query, a, err := sqlx.In(q, channel_id, idList)
	err = dao.db.QueryRowsPartial(ctx, &rows, query, a...)

	if err != nil {
		logx.WithContext(ctx).Errorf("Queryx in SelectByUserIdList(_), error: %v", err)
		return
	}

	rValues = rows

	return
}

// SelectByUserId
// select id, channel_id, user_id, participant_type, inviter_user_id, invited_at, joined_at, state from channel_participants where channel_id = :channel_id and user_id = :user_id
func (dao *ChannelParticipantsDAO) SelectByUserId(ctx context.Context, channel_id int64, user_id int64) (rValue *dataobject.ChannelParticipantDO, err error) {
	var (
		query = "select id, channel_id, user_id, participant_type, admin_rights, banned_rights, inviter_user_id, invited_at, joined_at, left_at, state from channel_participants where channel_id = ? and user_id = ?"
		row   *dataobject.ChannelParticipantDO
	)

	err = dao.db.QueryRowPartial(ctx, &row, query, channel_id, user_id)

	if err != nil {
		logx.WithContext(ctx).Errorf("Queryx in SelectByUserId(_), error: %v", err)
		return
	}

	rValue = row

	return
}

// DeleteChannelUser
// update channel_participants set state = 1 where id = :id
func (dao *ChannelParticipantsDAO) DeleteChannelUser(ctx context.Context, reason int32, id int64, left_at int32, bannedRights string, kiked_by int64) (i int64, err error) {
	var (
		query = "update channel_participants set state = 1, participant_type = ?, left_at = ?, banned_rights = ?, kicked_by = ? where id = ?"
		res   sql.Result
	)
	res, err = dao.db.Exec(ctx, query, reason, left_at, bannedRights, kiked_by, id)

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
func (dao *ChannelParticipantsDAO) Update(ctx context.Context, inviter_user_id int64, invited_at int32, joined_at int32, id int64) (i int64, err error) {
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
func (dao *ChannelParticipantsDAO) UpdateParticipantType(ctx context.Context, participant_type, state int8, id int64) (i int64, err error) {
	var (
		query = "update channel_participants set participant_type = ?, state = ? where id = ?"
		res   sql.Result
	)
	res, err = dao.db.Exec(ctx, query, participant_type, state, id)

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

// UpdateParticipantState
// update channel_participants set participant_type = :participant_type where id = :id
func (dao *ChannelParticipantsDAO) UpdateParticipantState(ctx context.Context, state int32, id int64) (i int64, err error) {
	var (
		query = "update channel_participants set state = ? where id = ?"
		res   sql.Result
	)
	res, err = dao.db.Exec(ctx, query, state, id)

	if err != nil {
		logx.WithContext(ctx).Errorf("Exec in UpdateParticipantState(_), error: %v", err)
		return
	}

	i, err = res.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("RowsAffected in UpdateParticipantState(_), error: %v", err)
		return
	}

	return
}

// UpdateAdminRights
// update channel_participants set participant_type = :participant_type where id = :id
func (dao *ChannelParticipantsDAO) UpdateAdminRights(ctx context.Context, admin_rights string, id int64) (i int64, err error) {
	var (
		query = "update channel_participants set admin_rights = ? where id = ?"
		res   sql.Result
	)
	res, err = dao.db.Exec(ctx, query, admin_rights, id)

	if err != nil {
		logx.WithContext(ctx).Errorf("Exec in UpdateAdminRights(_), error: %v", err)
		return
	}

	i, err = res.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("RowsAffected in UpdateAdminRights(_), error: %v", err)
		return
	}

	return
}

// UpdateBannedRights
// update channel_participants set participant_type = :participant_type where id = :id
func (dao *ChannelParticipantsDAO) UpdateBannedRights(ctx context.Context, banned_rights string, id int64) (i int64, err error) {
	var (
		query = "update channel_participants set banned_rights = ? where id = ?"
		res   sql.Result
	)
	res, err = dao.db.Exec(ctx, query, banned_rights, id)

	if err != nil {
		logx.WithContext(ctx).Errorf("Exec in UpdateBannedRights(_), error: %v", err)
		return
	}

	i, err = res.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("RowsAffected in UpdateBannedRights(_), error: %v", err)
		return
	}

	return
}

func (dao *ChannelParticipantsDAO) CheckExists(ctx context.Context, channelId, userId int64) (rValue bool, err error) {
	var (
		query = "select id from channel_participants where channel_id = ? and user_id = ? limit 1"
		id    int64
	)

	err = dao.db.QueryRowPartial(ctx, &id, query, channelId, userId)

	if err != nil {
		logx.WithContext(ctx).Errorf("Queryx in CheckExists(_), error: %v", err)
		return
	}

	rValue = id != 0

	return
}
