package core

import (
	"database/sql"
	"fmt"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/authorization/authorization"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/models"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/pkg/date"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/pkg/mail"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/pkg/numbers"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

// AuthSingUP
// TODO: need to write logic
func (c *AuthorizationCore) AuthSingUP(in *authorization.AuthSignUpRequest) (*authorization.AuthSignUpRsp, error) {
	usersEllo := &models.UsersEllo{}
	result := c.svcCtx.Gorm.Where("email = ?", in.Email).First(usersEllo)
	if result.Error != gorm.ErrRecordNotFound {
		c.Logger.Errorf("sign up query AuthSingUP method (%v)", result.Error)
		return nil, result.Error
	} else if result.RowsAffected > 0 {
		c.Logger.Info("this email was been registered")
		return nil, fmt.Errorf("this email was been registered")
	}

	dob, err := date.FormatDateIso8601(in.DateOfBirth)
	if err != nil {
		err = fmt.Errorf("date of birth should be in ISO 8601, current \"%v\"", in.DateOfBirth)
		c.Logger.Error(err)
		return nil, err
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	password := string(hashPassword[:])
	if err != nil {
		c.Logger.Errorf("cannot bcrypt password in sign up (%v)", err)
		return nil, err
	}

	confirmationCodes := &models.ConfirmationCodes{}

	if err = c.svcCtx.Gorm.Transaction(func(tx *gorm.DB) error {
		usersEllo = &models.UsersEllo{
			UserID:   uint(in.UserId),
			Password: password,
			Email:    in.Email,
			Gender:   in.Gender,
			Avatar: sql.NullString{
				String: in.Avatar,
			},
			DateOfBirth: &dob,
		}
		if err = tx.Create(usersEllo).Error; err != nil {
			err = fmt.Errorf("can not create users_ello record (%v)", err)
			return err
		}

		code, err := numbers.ConfirmationCode(6)
		if err != nil {
			err = fmt.Errorf("can not generate confirmation code (%v)", err)
			return err
		}

		exp := time.Now().Add(time.Minute * 10)
		confirmationCodes = &models.ConfirmationCodes{
			UserID:    uint(in.UserId),
			Code:      code,
			ExpiredAt: &exp,
		}
		if err = tx.Create(confirmationCodes).Error; err != nil {
			err = fmt.Errorf("can not create confirmation_codes record (%v)", err)
			return err
		}

		mailReq := &mail.SendMailRequest{
			Email:    in.Email,
			Username: in.Username,
			Message:  "Confirmation code: " + code,
			Subject:  "Confirmation code from ElloApp",
		}

		if _, err := mail.SendMail(c.ctx, mailReq); err != nil {
			err = fmt.Errorf("can not send code to mail (%v)", err)
			return err
		}

		return nil
	}); err != nil {
		c.Logger.Error(err)
		return nil, err
	}

	return &authorization.AuthSignUpRsp{
		Email:              in.Email,
		ConfirmationExpire: confirmationCodes.ExpiredAt.Unix(),
	}, nil
}
