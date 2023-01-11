package core

import (
	"errors"
	"fmt"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/authorization/authorization"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/models"
	"gorm.io/gorm"
	"net/mail"
	"time"
)

func (c *AuthorizationCore) Confirmation(in *authorization.ConfirmationRequest) (*authorization.ConfirmationResponse, error) {
	var userId uint
	if _, err := mail.ParseAddress(in.UsernameOrEmail); err != nil {
		var user models.Users
		if err := c.svcCtx.Gorm.Where("username = ?", in.UsernameOrEmail).First(&user).Error; err != nil && err != gorm.ErrRecordNotFound {
			err = fmt.Errorf("can not get user record (%v)", err)
			c.Logger.Error(err)
			return nil, err
		} else if err == gorm.ErrRecordNotFound {
			err = fmt.Errorf("no such user exists with username = %s", in.UsernameOrEmail)
			c.Logger.Error(err)
			return nil, err
		}
		userId = user.ID
	} else {
		var userEllo models.UsersEllo
		if err := c.svcCtx.Gorm.Where("email = ?", in.UsernameOrEmail).First(&userEllo).Error; err != nil && err != gorm.ErrRecordNotFound {
			err = fmt.Errorf("can not get user record (%v)", err)
			c.Logger.Error(err)
			return nil, err
		} else if err == gorm.ErrRecordNotFound {
			err = fmt.Errorf("no such user exists with email = %s", in.UsernameOrEmail)
			c.Logger.Error(err)
			return nil, err
		}
		userId = userEllo.UserID
	}

	var confirmationCode models.ConfirmationCodes
	if err := c.svcCtx.Gorm.Where("user_id = ?", userId).First(&confirmationCode).Error; err != nil && err != gorm.ErrRecordNotFound {
		err = fmt.Errorf("can not get user record (%v)", err)
		c.Logger.Error(err)
		return nil, err
	} else if err == gorm.ErrRecordNotFound {
		err = fmt.Errorf("no such user exists with email = %s", in.UsernameOrEmail)
		c.Logger.Error(err)
		return nil, err
	}

	if in.Code != confirmationCode.Code {
		err := errors.New("incorrect code")
		c.Logger.Error(err)
		return nil, err
	} else if confirmationCode.ExpiredAt.Before(time.Now()) {
		err := errors.New("code has been expired")
		c.Logger.Error(err)
		return nil, err
	}

	if err := c.svcCtx.Gorm.Model(&models.UsersEllo{}).Where("user_id = ?", userId).Update("email_confirmed", true).Error; err != nil {
		c.Logger.Error(err)
		return nil, err
	}

	return &authorization.ConfirmationResponse{
		Message: "success",
	}, nil
}
