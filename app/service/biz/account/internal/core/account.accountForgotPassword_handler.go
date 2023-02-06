package core

import (
	"errors"
	"fmt"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/models"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/pkg/mail"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/pkg/numbers"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"

	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/account/account"
)

func (c *AccountCore) AccountForgotPassword(in *account.ForgotAccountPassReq) (*account.ForgotAccountPassResp, error) {
	var userEllo models.UsersEllo

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
			Message: "There is no account in usersEllo",
		}, err
	}

	//send email verification code
	code, err := numbers.ConfirmationCode(6)
	if err != nil {
		c.Logger.Errorf("can not generate confirmation code (%v)", err)
		return &account.ForgotAccountPassResp{
			Status:             false,
			Email:              userEllo.Email,
			ConfirmationExpire: 0,
			Message:            "Can not generate confirmation code",
		}, err
	}

	if err := c.svcCtx.Gorm.Where("user_id", in.UserId).First(&userEllo).Error; err != nil && err != gorm.ErrRecordNotFound {
		err := fmt.Errorf("can not get users record (%v)", err)
		c.Logger.Error(err)
		return &account.ForgotAccountPassResp{
			Status:             false,
			Email:              userEllo.Email,
			ConfirmationExpire: 0,
			Message:            "Can not get usersEllo record",
		}, err
	} else if err == gorm.ErrRecordNotFound {
		err = errors.New("there is no account in usersEllo")
		c.Logger.Error(err)
		return &account.ForgotAccountPassResp{
			Status:             false,
			Email:              userEllo.Email,
			ConfirmationExpire: 0,
			Message:            "There is no account in usersEllo",
		}, err
	}

	exp := time.Now().Add(time.Minute * 10)
	confirmationCodes := &models.ConfirmationCodes{
		UserID:    userEllo.UserID,
		Code:      code,
		ExpiredAt: &exp,
	}

	if err := c.svcCtx.Gorm.Transaction(func(tx *gorm.DB) (err error) {
		if err = tx.Create(confirmationCodes).Error; err != nil {
			c.Logger.Errorf("can not create confirmation_codes record (%v)", err)
			return err
		}

		mailReq := &mail.SendMailRequest{
			Email:    userEllo.Email,
			Username: userEllo.Username,
			Message:  "Confirmation code: " + code,
			Subject:  "Confirmation code from ElloApp",
		}

		if _, err = mail.SendMail(c.ctx, mailReq); err != nil {
			c.Logger.Errorf("can not send code to mail (%v)", err)
			return err
		}

		return
	}); err != nil {
		return nil, err
	}

	return &account.ForgotAccountPassResp{
		Status:             true,
		Email:              userEllo.Email,
		ConfirmationExpire: confirmationCodes.ExpiredAt.Unix(),
		Message:            "Sent Successfully",
	}, nil
}
