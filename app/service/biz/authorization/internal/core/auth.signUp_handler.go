package core

import (
	"database/sql"
	"fmt"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/authsession/authsession"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/authorization/authorization"
	userpb "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/user/user"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/models"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto/crypto"
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
	var (
		user *userpb.ImmutableUser
	)

	var findUser models.Users
	result := c.svcCtx.Gorm.Where("username = ?", in.Username).First(&findUser)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		c.Logger.Errorf("sign up query AuthSingUP method (%v)", result.Error)
		return nil, result.Error
	} else if result.RowsAffected > 0 {
		c.Logger.Info("this username was been registered")
		return nil, fmt.Errorf("this email was been registered")
	}

	var findUserEllo models.UsersEllo
	result = c.svcCtx.Gorm.Where("email = ?", in.Email).First(&findUserEllo)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
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

	confirmationCodes := &models.ConfirmationCodes{}

	if err := c.svcCtx.Gorm.Transaction(func(tx *gorm.DB) error {
		/*
			Old Authorization method start
		*/
		key := crypto.CreateAuthKey()
		if _, err := c.svcCtx.Dao.AuthsessionClient.AuthsessionSetAuthKey(c.ctx, &authsession.TLAuthsessionSetAuthKey{
			AuthKey: &mtproto.AuthKeyInfo{
				AuthKeyId:          key.AuthKeyId(),
				AuthKey:            key.AuthKey(),
				AuthKeyType:        mtproto.AuthKeyTypePerm,
				PermAuthKeyId:      key.AuthKeyId(),
				TempAuthKeyId:      0,
				MediaTempAuthKeyId: 0,
			},
			FutureSalt: nil,
		}); err != nil {
			c.Logger.Errorf("create user secret key error")
			return err
		}

		if user, err = c.svcCtx.UserClient.UserCreateNewUser(c.ctx, &userpb.TLUserCreateNewUser{
			SecretKeyId: key.AuthKeyId(),
			Phone:       in.Phone,
			CountryCode: in.CountryCode,
			FirstName:   in.FirstName,
			LastName:    in.LastName,
		}); err != nil {
			c.Logger.Errorf("createNewUser error: %v", err)
			return err
		}

		if _, err = c.svcCtx.Dao.AuthsessionClient.AuthsessionBindAuthKeyUser(c.ctx, &authsession.TLAuthsessionBindAuthKeyUser{
			AuthKeyId: in.MData.AuthId,
			UserId:    user.User.Id,
		}); err != nil {
			c.Logger.Errorf("bindAuthKeyUser error: %v", err)
			err = mtproto.ErrInternelServerError
			return err
		}
		/*
			Old Authorization method end
		*/

		/*
			New Authorization_customize method start
		*/

		hashPassword, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
		password := string(hashPassword[:])
		if err != nil {
			c.Logger.Error("can not generate hash for password")
			return err
		}

		usersEllo := &models.UsersEllo{
			UserID:   uint(user.Id()),
			Username: in.Username,
			Password: password,
			Email:    in.Email,
			Gender:   in.Gender,
			Avatar: sql.NullString{
				String: in.Avatar,
			},
			DateOfBirth: &dob,
		}
		if err = tx.Create(usersEllo).Error; err != nil {
			c.Logger.Errorf("can not create users_ello record (%v)", err)
			return err
		}

		code, err := numbers.ConfirmationCode(6)
		if err != nil {
			c.Logger.Errorf("can not generate confirmation code (%v)", err)
			return err
		}

		exp := time.Now().Add(time.Minute * 10)
		confirmationCodes = &models.ConfirmationCodes{
			UserID:    uint(user.Id()),
			Code:      code,
			ExpiredAt: &exp,
		}
		if err = tx.Create(confirmationCodes).Error; err != nil {
			c.Logger.Errorf("can not create confirmation_codes record (%v)", err)
			return err
		}

		mailReq := &mail.SendMailRequest{
			Email:    in.Email,
			Username: in.Username,
			Message:  "Confirmation code: " + code,
			Subject:  "Confirmation code from ElloApp",
		}

		if _, err := mail.SendMail(c.ctx, mailReq); err != nil {
			c.Logger.Errorf("can not send code to mail (%v)", err)
			return err
		}

		/*
			New Authorization_customize method end
		*/

		return nil
	}); err != nil {
		return nil, err
	}

	return &authorization.AuthSignUpRsp{
		Email:              in.Email,
		ConfirmationExpire: confirmationCodes.ExpiredAt.Unix(),
	}, nil
}
