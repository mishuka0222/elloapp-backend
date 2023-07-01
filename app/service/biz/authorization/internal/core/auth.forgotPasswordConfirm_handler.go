package core

import (
	"errors"
	"fmt"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/authorization/authorization"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

func (c *AuthorizationCore) ForgotPasswordConfirm(in *authorization.ForgotPasswordConfirmReq) (*authorization.ForgotPasswordConfirmRsp, error) {
	var userEllo models.UsersEllo

	if err := c.svcCtx.Gorm.Where("email = ?", in.Email).Or("email = ?", in.Email).First(&userEllo).Error; err != nil && err != gorm.ErrRecordNotFound {
		err = fmt.Errorf("can not get userEllo record (%v)", err)
		c.Logger.Error(err)
		return &authorization.ForgotPasswordConfirmRsp{
			Status:  false,
			Message: "Can not get userEllo record",
		}, err
	} else if err == gorm.ErrRecordNotFound {
		err = fmt.Errorf("no such userEllo exists with email = %s", in.Email)
		c.Logger.Error(err)
		return &authorization.ForgotPasswordConfirmRsp{
			Status:  false,
			Message: "No such userEllo exists with email",
		}, err
	}

	var confirmationCode models.ConfirmationCodes
	if err := c.svcCtx.Gorm.Where("user_id = ?", userEllo.UserID).First(&confirmationCode).Error; err != nil && err != gorm.ErrRecordNotFound {
		err = fmt.Errorf("can not get user record (%v)", err)
		c.Logger.Error(err)
		return &authorization.ForgotPasswordConfirmRsp{
			Status:  false,
			Message: "Can not get user record",
		}, err
	} else if err == gorm.ErrRecordNotFound {
		err = fmt.Errorf("no such userEllo exists with email = %s", userEllo.Email)
		c.Logger.Error(err)
		return &authorization.ForgotPasswordConfirmRsp{
			Status:  false,
			Message: "No such userEllo exists with email",
		}, err
	}

	if in.Code != confirmationCode.Code {
		err := errors.New("incorrect code")
		c.Logger.Error(err)
		return &authorization.ForgotPasswordConfirmRsp{
			Status:  false,
			Message: "Incorrect code",
		}, err
	} else if confirmationCode.ExpiredAt.Before(time.Now()) {
		err := errors.New("code has been expired")
		c.Logger.Error(err)
		return &authorization.ForgotPasswordConfirmRsp{
			Status:  false,
			Message: "Code has been expired",
		}, err
	}

	userEllo.EmailConfirmed = true
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(in.NewPass), bcrypt.DefaultCost)
	hashedPass := string(hashPassword[:])
	if err != nil {
		err = errors.New("can not generate hash for password")
		c.Logger.Error(err)
	}
	userEllo.Password = hashedPass

	if err := c.svcCtx.Gorm.Save(&userEllo).Error; err != nil {
		c.Logger.Error(err)
		return nil, err
	}

	return &authorization.ForgotPasswordConfirmRsp{
		Status:  true,
		Message: "Password Changed Successfully",
	}, nil
}
