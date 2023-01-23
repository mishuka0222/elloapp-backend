package mysql_dao

import (
	"context"
	"database/sql"
	"github.com/zeromicro/go-zero/core/logx"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/pkg2/stores/sqlx"
)

var _ *sql.Result

type UsersDAO struct {
	db *sqlx.DB
}

func NewUsersDAO(db *sqlx.DB) *UsersDAO {
	return &UsersDAO{db}
}

// SelectByUsername
// select id from users where username = :username limit 1
func (dao *UsersDAO) CheckExistsUsername(ctx context.Context, username string) (rValue int64, err error) {
	var (
		query = "select id from users where username = ? limit 1"
		id    int64
	)
	err = dao.db.QueryRowPartial(ctx, &id, query, username)

	if err != nil {
		if err != sqlx.ErrNotFound {
			logx.WithContext(ctx).Errorf("queryx in CheckExistsUsername(_), error: %v", err)
			return
		} else {
			err = nil
		}
	} else {
		rValue = id
	}

	return
}
