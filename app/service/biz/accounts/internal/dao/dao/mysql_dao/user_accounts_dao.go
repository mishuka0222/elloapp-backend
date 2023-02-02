package mysql_dao

import (
	"context"
	"database/sql"

	"github.com/zeromicro/go-zero/core/logx"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/pkg2/stores/sqlx"
	"golang.org/x/crypto/bcrypt"
)

var _ *sql.Result

type UserAccountsDAO struct {
	db *sqlx.DB
}

func NewUserAccountsDAO(db *sqlx.DB) *UserAccountsDAO {
	return &UserAccountsDAO{db}
}

func (dao *UserAccountsDAO) DeleteUserAccount(ctx context.Context, user_id int64) (status bool, err error) {
	var (
		query = `update users set deleted = 1 where user_id = ?`
	)
	_, err = dao.db.Exec(ctx, query, user_id)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in DeleteAccount(_), error: %v", err)
		return false, err
	}

	return true, nil
}

func (dao *UserAccountsDAO) UpdateAccountProfile(ctx context.Context, user_id int64, firstName string, lastName string, username string, bio string, birthday string, gender string, country string) (status bool, err error) {
	var (
		query  = `update users set first_name = ?, last_name = ?, about = ?, country_code = ? where user_id = ?`
		query1 = `update users_ello set username = ?, date_of_birth = ?, gender = ? where user_id = ?`
	)
	_, err = dao.db.Exec(ctx, query, firstName, lastName, bio, country, user_id)
	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in UpdateProfile in users, error: %v", err)
		return false, err
	}

	_, err = dao.db.Exec(ctx, query1, username, birthday, gender, user_id)
	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in UpdateProfile in users_ello, error: %v", err)
		return false, err
	}

	return true, nil
}

func (dao *UserAccountsDAO) ChangeAccountEmail(ctx context.Context, user_id int64, new_email string) (status bool, err error) {
	var (
		query = `update users_ello set email = ? where user_id = ?`
	)
	_, err = dao.db.Exec(ctx, query, new_email, user_id)
	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in UpdateEmail in users, error: %v", err)
		return false, err
	}

	return true, nil
}

func (dao *UserAccountsDAO) ForgotAccountPassword(ctx context.Context, user_id int64, new_pass string) (status bool, err error) {
	var (
		query = `update users_ello set password = ? where user_id = ?`
	)

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(new_pass), bcrypt.DefaultCost)
	newPass := string(hashPassword[:])
	if err != nil {
		logx.WithContext(ctx).Errorf("can not generate hash for password")
		return false, err
	}
	_, err = dao.db.Exec(ctx, query, newPass, user_id)
	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in ChangePassword in users_ello, error: %v", err)
		return false, err
	}

	return true, nil
}

func (dao *UserAccountsDAO) ChangeAccountPassword(ctx context.Context, user_id int64, prev_pass string, new_pass string) (status bool, err error) {
	var (
		query    = `select password from users_ello where user_id = ?`
		query1   = `update users_ello set password = ? where user_id = ?`
		password string
	)

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(prev_pass), bcrypt.DefaultCost)
	prevPass := string(hashPassword[:])
	if err != nil {
		logx.WithContext(ctx).Errorf("can not generate hash for password")
		return false, err
	}

	hashPassword, err = bcrypt.GenerateFromPassword([]byte(new_pass), bcrypt.DefaultCost)
	newPass := string(hashPassword[:])
	if err != nil {
		logx.WithContext(ctx).Errorf("can not generate hash for password")
		return false, err
	}

	err = dao.db.QueryRow(ctx, &password, query, user_id)
	if err != nil {
		logx.WithContext(ctx).Errorf("Queryx in GetPassword(_), error: %v", err)
		return
	}

	if password != prevPass {
		return false, nil
	}

	_, err = dao.db.Exec(ctx, query1, newPass, user_id)
	if err != nil {
		logx.WithContext(ctx).Errorf("Queryx in UpdatePassword(_), error: %v", err)
		return
	}

	return true, nil
}
