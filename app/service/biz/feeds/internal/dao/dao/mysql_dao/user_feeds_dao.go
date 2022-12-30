package mysql_dao

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/zeromicro/go-zero/core/logx"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/feeds/feeds"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/feeds/internal/dao/dataobject"
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
// select uf.user_id, uf.chat_id, uf.peer_type from user_feeds uf, chat_participants cp, chats c where uf.chat_id = cp.chat_id = c.id and cp.user_id = uf.user_id and cp.state = 0 and c.deactivated = 0 and uf.user_id = ?
func (dao *UserFeedsDAO) SelectFeedList(ctx context.Context, user_id int64) (rList []dataobject.FeedItemDO, err error) {
	var (
		query = `select uf.user_id, uf.chat_id, uf.peer_type from user_feeds uf, chat_participants cp, chats c 
        where uf.chat_id = cp.chat_id = c.id and cp.user_id = uf.user_id and cp.state = 0 and c.deactivated = 0 
        and uf.user_id = ?`
		values []dataobject.FeedItemDO
	)
	err = dao.db.QueryRowsPartial(ctx, &values, query, user_id)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectFeedList(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectOffsetMinList - generate offset for getHistory func
// select min(user_message_box_id) as offset_min, max(user_message_box_id) as offset_max, peer_id, count(*) as count from (
// select m.user_message_box_id, d.peer_id, d.created_at from dialogs d, messages m where d.peer_id in ( %s )
// and d.user_id = ? and user_message_box_id < (read_inbox_max_id - ?) and d.peer_id = m.peer_id
// and d.user_id = m.user_id order by m.id desc limit ?) t group by peer_id
func (dao *UserFeedsDAO) SelectOffsetMinList(ctx context.Context, user_id int64, chats []int64, limit, offset int32) (rList []*feeds.OffsetItemDo, err error) {
	var (
		query = `select min(user_message_box_id) as offset_min, max(user_message_box_id) as offset_max, peer_id, count(*) as count from (
        select m.user_message_box_id, d.peer_id, d.created_at from dialogs d, messages m
        where d.peer_id in ( %s ) and d.user_id = ?
        and user_message_box_id < (read_inbox_max_id - ?)
		and d.peer_id = m.peer_id and d.user_id = m.user_id and d.peer_type = m.peer_type
        order by m.id desc limit ?) t group by peer_id`
		values []*feeds.OffsetItemDo
		args   []interface{}
		cArr   []string
	)
	if len(chats) == 0 {
		return nil, errors.New(" Empty chat_id list not allowed")
	}
	for i := range chats {
		cArr = append(cArr, "?")
		args = append(args, chats[i])
	}
	query = fmt.Sprintf(query, strings.Join(cArr, ","))
	args = append(args, user_id, offset, limit)

	err = dao.db.QueryRowsPartial(ctx, &values, query, args...)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectOffsetMinList(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectOffsetMaxList - generate offset for getHistory func
// select min(user_message_box_id) as offset_min, max(user_message_box_id) as offset_max, peer_id, count(*) as count from (
// select m.user_message_box_id, d.peer_id, d.created_at from dialogs d, messages m where d.peer_id in ( %s )
// and d.user_id = ? and user_message_box_id >= read_inbox_max_id and d.peer_id = m.peer_id
// and d.user_id = m.user_id order by m.id asc limit ?) t group by peer_id
func (dao *UserFeedsDAO) SelectOffsetMaxList(ctx context.Context, user_id int64, chats []int64, limit int32) (rList []*feeds.OffsetItemDo, err error) {
	var (
		query = `select min(user_message_box_id) as offset_min, max(user_message_box_id) as offset_max, peer_id, count(*) as count from (
    	select m.user_message_box_id, d.peer_id, d.created_at from dialogs d, messages m
        where d.peer_id in ( %s ) and d.user_id = ?
        and user_message_box_id >= read_inbox_max_id
  		and d.peer_id = m.peer_id and d.user_id = m.user_id and d.peer_type = m.peer_type
        order by m.id asc limit ?) t group by peer_id`
		values []*feeds.OffsetItemDo
		args   []interface{}
		cArr   []string
	)
	if len(chats) == 0 {
		return nil, errors.New(" Empty chat_id list not allowed")
	}
	for i := range chats {
		cArr = append(cArr, "?")
		args = append(args, chats[i])
	}
	query = fmt.Sprintf(query, strings.Join(cArr, ","))
	args = append(args, user_id, limit)

	err = dao.db.QueryRowsPartial(ctx, &values, query, args...)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectOffsetMaxList(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectUnreadCountList - select count unread messages for user feeds chat_id
// select uf.chat_id, top_message - read_inbox_max_id as unread from dialogs d, (
// select c.id as chat_id, cp.user_id  from chat_participants cp, chats c where cp.user_id in (?)
// and cp.chat_id = c.id  and cp.state = 0 and c.deactivated = 0 and c.id in (
// select uf.chat_id
// from user_feeds uf where uf.user_id in (?))) uf where d.user_id = uf.user_id
// and d.peer_id = uf.chat_id and d.user_id in (?)
func (dao *UserFeedsDAO) SelectUnreadCountList(ctx context.Context, user_id int64) (rList []dataobject.UnreadCountDo, err error) {
	var (
		query = `select uf.chat_id, top_message - read_inbox_max_id as unread from dialogs d, 
        (select c.id as chat_id, cp.user_id from chat_participants cp, chats c where cp.user_id in (?) 
        and cp.chat_id = c.id and cp.state = 0 and c.deactivated = 0 and c.id in 
        (select uf.chat_id from user_feeds uf where uf.user_id in (?))) uf where d.user_id = uf.user_id 
        and d.peer_id = uf.chat_id and d.user_id in (?)`
		values []dataobject.UnreadCountDo
	)
	err = dao.db.QueryRowsPartial(ctx, &values, query, user_id, user_id, user_id)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectUnreadCountList(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectChatList
// select cp.chat_id, c.photo_id, c.title from chat_participants cp, chats c where cp.state = 0
// and cp.user_id in (?) and c.id = cp.chat_id and c.deactivated = 0
func (dao *UserFeedsDAO) SelectChatList(ctx context.Context, user_id int64) (rList []*feeds.UserChatDO, err error) {
	var (
		query = `select cp.chat_id, c.photo_id, c.title from chat_participants cp, chats c where cp.state = 0 
        and cp.user_id in (?) and c.id = cp.chat_id and c.deactivated = 0`
		values []*feeds.UserChatDO
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
func (dao *UserFeedsDAO) DeleteFromListElseValue(ctx context.Context, user_id int64, chats []*feeds.FeedInsertItemDO) (rowsAffected int64, err error) {
	var (
		query   = `delete from user_feeds where user_id = ?`
		rResult sql.Result
	)
	if len(chats) != 0 {
		query += " and chat_id not in"
	} else {
		chats = make([]*feeds.FeedInsertItemDO, 0)
	}
	var args []interface{}
	args = append(args, user_id)

	var cArr []string
	for i := range chats {
		cArr = append(cArr, "?")
		args = append(args, chats[i].GetChatId())
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
func (dao *UserFeedsDAO) InsertList(ctx context.Context, user_id int64, chats []*feeds.FeedInsertItemDO) (rowsAffected int64, err error) {
	var (
		query   = `insert into user_feeds (user_id, chat_id, peer_type) values `
		rResult sql.Result
	)
	if len(chats) == 0 {
		return 0, nil
	}
	var args []interface{}
	var cArr []string
	for i := range chats {
		cArr = append(cArr, "(?,?,?)")
		args = append(args, user_id, chats[i].GetChatId(), chats[i].GetPeerType())
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
