package core

import (
	"errors"
	"fmt"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/account/account"
)

func (c *AccountCore) AccountChangePassword(in *account.ChangeAccountPasswordReq) (*account.ChangeAccountPasswordResp, error) {
	// todo: add your logic here and delete this line
	var userEllo models.UsersEllo

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(in.PrevPass), bcrypt.DefaultCost)
	prevPass := string(hashPassword[:])
	if err != nil {
		err = errors.New("can not generate hash for password")
		c.Logger.Error(err)
	}

	hashPassword, err = bcrypt.GenerateFromPassword([]byte(in.NewPass), bcrypt.DefaultCost)
	newPass := string(hashPassword[:])
	if err != nil {
		err = errors.New("can not generate hash for password")
		c.Logger.Error(err)
	}

	if err := c.svcCtx.Gorm.Where("user_id = ?", in.UserId).First(&userEllo).Error; err != nil && err != gorm.ErrRecordNotFound {
		err := fmt.Errorf("can not get users record (%v)", err)
		c.Logger.Error(err)
		return &account.ChangeAccountPasswordResp{
			Status:  false,
			Message: "Can not get usersEllo record",
		}, err
	} else if err == gorm.ErrRecordNotFound {
		err = errors.New("there is no account in usersEllo")
		c.Logger.Error(err)
		return &account.ChangeAccountPasswordResp{
			Status:  false,
			Message: "There is not account in usersEllo",
		}, err
	}

	if userEllo.Password != prevPass {
		err = errors.New("password is not matched")
		c.Logger.Error(err)
		return &account.ChangeAccountPasswordResp{
			Status:  false,
			Message: "Incorrect Password",
		}, err
	}

	if err := c.svcCtx.Gorm.Where("user_id = ?", in.UserId).Updates(&models.UsersEllo{Password: newPass}).Error; err != nil && err != gorm.ErrRecordNotFound {
		err := fmt.Errorf("can not get users record (%v)", err)
		c.Logger.Error(err)
		return &account.ChangeAccountPasswordResp{
			Status:  false,
			Message: "Can not get usersEllo record",
		}, err
	} else if err == gorm.ErrRecordNotFound {
		err = errors.New("there is no account in usersEllo")
		c.Logger.Error(err)
		return &account.ChangeAccountPasswordResp{
			Status:  false,
			Message: "There is not account in usersEllo",
		}, err
	}
	return &account.ChangeAccountPasswordResp{
		Status:  true,
		Message: "Password updated successfully",
	}, nil
}
