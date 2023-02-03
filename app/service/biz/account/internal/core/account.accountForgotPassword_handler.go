package core

import (
	"errors"
	"fmt"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/account/account"
)

func (c *AccountCore) AccountForgotPassword(in *account.ForgotAccountPassReq) (*account.ForgotAccountPassResp, error) {
	// todo: add your logic here and delete this line
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(in.NewPass), bcrypt.DefaultCost)
	hashedPass := string(hashPassword[:])
	if err != nil {
		err = errors.New("can not generate hash for password")
		c.Logger.Error(err)
	}

	if err := c.svcCtx.Gorm.Where("user_id = ?", in.UserId).Updates(&models.UsersEllo{Password: hashedPass}).Error; err != nil && err != gorm.ErrRecordNotFound {
		err := fmt.Errorf("can not get users record (%v)", err)
		c.Logger.Error(err)
		return &account.ForgotAccountPassResp{
			Status:  false,
			Message: "Can not get usersEllo record",
		}, err
	} else if err == gorm.ErrRecordNotFound {
		err = errors.New("there is no account in usersEllo")
		c.Logger.Error(err)
		return &account.ForgotAccountPassResp{
			Status:  false,
			Message: "There is not account in usersEllo",
		}, err
	}
	return &account.ForgotAccountPassResp{}, nil
}
