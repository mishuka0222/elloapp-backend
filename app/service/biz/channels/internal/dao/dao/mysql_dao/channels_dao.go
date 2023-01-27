package mysql_dao

import (
	"context"
	"database/sql"
	"github.com/zeromicro/go-zero/core/logx"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/channels/internal/dao/dataobject"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/pkg2/stores/sqlx"
)

var _ *sql.Result

type ChannelsDAO struct {
	db *sqlx.DB
}

func NewChannelsDAO(db *sqlx.DB) *ChannelsDAO {
	return &ChannelsDAO{db}
}

// Insert
// insert into channels(creator_user_id, access_hash, random_id, participant_count, title, about, `date`) values (:creator_user_id, :access_hash, :random_id, :participant_count, :title, :about, :date)
func (dao *ChannelsDAO) Insert(ctx context.Context, do *dataobject.ChannelDO) (id int64, err error) {
	var (
		query = "insert into channels(creator_user_id, access_hash, random_id, participant_count, title, about, `date`) values (:creator_user_id, :access_hash, :random_id, :participant_count, :title, :about, :date)"
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

// Select
// select id, creator_user_id, access_hash, participant_count, title, about, photo_id, admins_enabled, deactivated, version, `date` from channels where id = :id
func (dao *ChannelsDAO) Select(ctx context.Context, id int64) (rValue *dataobject.ChannelDO, err error) {
	var (
		query = "select id, creator_user_id, access_hash, participant_count, title, about, photo_id, link, username, admins_enabled, deactivated, version, `date` from channels where id = ?"
		row   dataobject.ChannelDO
	)
	err = dao.db.QueryRowPartial(ctx, &row, query, id)

	if err != nil {
		logx.WithContext(ctx).Errorf("Queryx in Select(_), error: %v", err)
		return
	}

	rValue = &row

	return
}

// UpdateTitle
// update channels set title = :title, `date` = :date, version = version + 1 where id = :id
func (dao *ChannelsDAO) UpdateTitle(ctx context.Context, title string, date int32, id int64) (i int64, err error) {
	var (
		query = "update channels set title = ?, `date` = ?, version = version + 1 where id = ?"
		res   sql.Result
	)
	res, err = dao.db.Exec(ctx, query, title, date, id)

	if err != nil {
		logx.WithContext(ctx).Errorf("Exec in UpdateTitle(_), error: %v", err)
		return
	}

	i, err = res.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("RowsAffected in UpdateTitle(_), error: %v", err)
		return
	}

	return
}

// UpdateAbout
// update channels set about = :about, `date` = :date, version = version + 1 where id = :id
func (dao *ChannelsDAO) UpdateAbout(ctx context.Context, about string, date int32, id int64) (i int64, err error) {
	var (
		query = "update channels set about = ?, `date` = ?, version = version + 1 where id = ?"
		res   sql.Result
	)
	res, err = dao.db.Exec(ctx, query, about, date, id)

	if err != nil {
		logx.WithContext(ctx).Errorf("Exec in UpdateAbout(_), error: %v", err)
		return
	}

	i, err = res.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("RowsAffected in UpdateAbout(_), error: %v", err)
		return
	}

	return
}

// UpdateLink
// update channels set link = :link, `date` = :date, version = version + 1 where id = :id
func (dao *ChannelsDAO) UpdateLink(ctx context.Context, link string, date int32, id int64) (i int64, err error) {
	var (
		query = "update channels set link = ?, `date` = ?, version = version + 1 where id = ?"
		res   sql.Result
	)
	res, err = dao.db.Exec(ctx, query, link, date, id)

	if err != nil {
		logx.WithContext(ctx).Errorf("Exec in UpdateLink(_), error: %v", err)
		return
	}

	i, err = res.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("RowsAffected in UpdateLink(_), error: %v", err)
		return
	}

	return
}

// UpdateUsername
// update channels set username = :link, `date` = :date, version = version + 1 where id = :id
func (dao *ChannelsDAO) UpdateUsername(ctx context.Context, username string, date int32, id int64) (i int64, err error) {
	var (
		query = "update channels set username = ?, `date` = ?, version = version + 1 where id = ?"
		res   sql.Result
	)
	res, err = dao.db.Exec(ctx, query, username, date, id)

	if err != nil {
		logx.WithContext(ctx).Errorf("Exec in UpdateUsername(_), error: %v", err)
		return
	}

	i, err = res.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("RowsAffected in UpdateUsername(_), error: %v", err)
		return
	}

	return
}

// SelectByIdList
// select id, access_hash, participant_count, title, about, photo_id, admins_enabled, deactivated, version, `date` from channels where id in (:idList)
func (dao *ChannelsDAO) SelectByIdList(ctx context.Context, idList []int64) (rValues []dataobject.ChannelDO, err error) {
	var (
		q    = "select id, access_hash, participant_count, title, about, photo_id, link, username, admins_enabled, deactivated, version, `date` from channels where id in (?)"
		rows []dataobject.ChannelDO
	)
	query, a, err := sqlx.In(q, idList)
	err = dao.db.QueryRowsPartial(ctx, &rows, query, a...)

	if err != nil {
		logx.WithContext(ctx).Errorf("Queryx in SelectByIdList(_), error: %v", err)
		return
	}

	rValues = rows

	return
}

// UpdateParticipantCount
// update channels set participant_count = :participant_count, `date` = :date, version = version + 1 where id = :id
func (dao *ChannelsDAO) UpdateParticipantCount(ctx context.Context, participant_count int32, date int32, id int64) (i int64, err error) {
	var (
		query = "update channels set participant_count = ?, `date` = ?, version = version + 1 where id = ?"
		res   sql.Result
	)
	res, err = dao.db.Exec(ctx, query, participant_count, date, id)

	if err != nil {
		logx.WithContext(ctx).Errorf("Exec in UpdateParticipantCount(_), error: %v", err)
		return
	}

	i, err = res.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("RowsAffected in UpdateParticipantCount(_), error: %v", err)
		return
	}

	return
}

// UpdatePhotoId
// update channels set photo_id = :photo_id, `date` = :date, version = version + 1 where id = :id
func (dao *ChannelsDAO) UpdatePhotoId(ctx context.Context, photo_id int64, date int32, id int64) (i int64, err error) {
	var (
		query = "update channels set photo_id = ?, `date` = ?, version = version + 1 where id = ?"
		res   sql.Result
	)
	res, err = dao.db.Exec(ctx, query, photo_id, date, id)

	if err != nil {
		logx.WithContext(ctx).Errorf("Exec in UpdatePhotoId(_), error: %v", err)
		return
	}

	i, err = res.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("RowsAffected in UpdatePhotoId(_), error: %v", err)
		return
	}

	return
}

// UpdateAdminsEnabled
// update channels set admins_enabled = :admins_enabled, `date` = :date, version = version + 1 where id = :id
func (dao *ChannelsDAO) UpdateAdminsEnabled(ctx context.Context, admins_enabled bool, date int32, id int64) (i int64, err error) {
	var (
		query = "update channels set admins_enabled = ?, `date` = ?, version = version + 1 where id = ?"
		res   sql.Result
	)
	res, err = dao.db.Exec(ctx, query, admins_enabled, date, id)

	if err != nil {
		logx.WithContext(ctx).Errorf("Exec in UpdateAdminsEnabled(_), error: %v", err)
		return
	}

	i, err = res.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("RowsAffected in UpdateAdminsEnabled(_), error: %v", err)
		return
	}

	return
}

// UpdateVersion
// update channels set `date` = :date, version = version + 1 where id = :id
func (dao *ChannelsDAO) UpdateVersion(ctx context.Context, date int32, id int64) (i int64, err error) {
	var (
		query = "update channels set `date` = ?, version = version + 1 where id = ?"
		res   sql.Result
	)
	res, err = dao.db.Exec(ctx, query, date, id)

	if err != nil {
		logx.WithContext(ctx).Errorf("Exec in UpdateVersion(_), error: %v", err)
		return
	}

	i, err = res.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("RowsAffected in UpdateVersion(_), error: %v", err)
		return
	}

	return
}

func (dao *ChannelsDAO) CheckUsername(ctx context.Context, username string) (rValue bool, err error) {
	var (
		query = "select id from channels where link = ? limit 1"
		row   int64
	)
	err = dao.db.QueryRow(ctx, &row, query, username)

	if err != nil {
		logx.WithContext(ctx).Errorf("Queryx in CheckUsername(_), error: %v", err)
		return
	}

	rValue = row != 0

	return
}
