package mysql_dao

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/teamgram-server/app/bff/feeds/internal/dao/dataobject"
	"github.com/zeromicro/go-zero/core/logx"
	"strings"
)

var _ *sql.Result

type UserFeedsDAO struct {
	db *sqlx.DB
}

func NewUserFeedsDAO(db *sqlx.DB) *UserFeedsDAO {
	return &UserFeedsDAO{db}
}

// SelectFeedList - select only active chats and where user is participant
// select c.id from chat_participants cp, chats c where cp.user_id in (?) and cp.chat_id = c.id and cp.state = 0 and c.deactivated = 0 and c.id in (select uf.chat_id from user_feeds uf where uf.user_id in (?))
func (dao *UserFeedsDAO) SelectFeedList(ctx context.Context, user_id int64) (rList []int64, err error) {
	var (
		query  = `select c.id from chat_participants cp, chats c where cp.user_id in (?) and cp.chat_id = c.id and cp.state = 0 and c.deactivated = 0 and c.id in (select uf.chat_id from user_feeds uf where uf.user_id in (?))`
		values []int64
	)
	err = dao.db.QueryRowsPartial(ctx, &values, query, user_id, user_id)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectFeedList(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectChatList
// select cp.chat_id, c.photo_id, c.title from chat_participants cp, chats c where cp.state = 0 and cp.user_id in (?) and c.id = cp.chat_id and c.deactivated = 0
func (dao *UserFeedsDAO) SelectChatList(ctx context.Context, user_id int64) (rList []dataobject.UserChatDO, err error) {
	var (
		query  = `select cp.chat_id, c.photo_id, c.title from chat_participants cp, chats c where cp.state = 0 and cp.user_id in (?) and c.id = cp.chat_id and c.deactivated = 0`
		values []dataobject.UserChatDO
	)
	err = dao.db.QueryRowsPartial(ctx, &values, query, user_id)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectChatList(_), error: %v", err)
		return
	}

	rList = values

	return
}

// DeleteFromListElseValue - delete all user chats except selected
// elete from user_feeds where user_id = ? and chat_id not in (?,?,?)
func (dao *UserFeedsDAO) DeleteFromListElseValue(ctx context.Context, user_id int64, chats []int64) (rowsAffected int64, err error) {
	var (
		query   = `delete from user_feeds where user_id = ?`
		rResult sql.Result
	)
	if len(chats) != 0 {
		query += " and chat_id not in"
	} else {
		chats = make([]int64, 0)
	}
	var args []interface{}
	args = append(args, user_id)

	var cArr []string
	for i := range chats {
		cArr = append(cArr, "?")
		args = append(args, chats[i])
	}
	q := strings.Join(cArr, ",")
	if q != "" {
		query = fmt.Sprintf("%s (%s) ", query, q)
	}

	rResult, err = dao.db.Exec(ctx, query, args...)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in DeleteFromListElseValue(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in DeleteFromListElseValue(_), error: %v", err)
	}

	return
}

// InsertList
// insert into user_feeds (user_id, chat_id) values (?,?),(?,?),(?,?) on duplicate key update id = last_insert_id(id)
func (dao *UserFeedsDAO) InsertList(ctx context.Context, user_id int64, chats []int64) (rowsAffected int64, err error) {
	var (
		query   = `insert into user_feeds (user_id, chat_id) values `
		rResult sql.Result
	)
	if len(chats) == 0 {
		return 0, nil
	}
	var args []interface{}
	var cArr []string
	for i := range chats {
		cArr = append(cArr, "(?,?)")
		args = append(args, user_id, chats[i])
	}
	q := strings.Join(cArr, ",")
	query += q + " on duplicate key update id = last_insert_id(id)"
	rResult, err = dao.db.Exec(ctx, query, args...)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in InsertList(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in InsertList(_), error: %v", err)
	}

	return
}
