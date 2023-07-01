package core

import (
	"database/sql"
	"errors"
	"fmt"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/account/account"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/models"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/pkg/date"
	"gorm.io/gorm"
)

func (c *AccountCore) AccountChangeProfile(in *account.ChangeAccountInfoReq) (*account.ChangeAccountInfoResp, error) {
	var (
		userEllo models.UsersEllo
		user     models.Users
		username models.Username
		err      error
		resp     = &account.ChangeAccountInfoResp{
			Status:  true,
			Message: "Updated successfully",
		}
	)

	if err := c.svcCtx.Gorm.Where("id = ?", in.UserId).First(&user).Error; err != nil && err != gorm.ErrRecordNotFound {
		err := fmt.Errorf("can not get users record (%v)", err)
		c.Logger.Error(err)
		return &account.ChangeAccountInfoResp{
			Status:  false,
			Message: "Can not get users record",
		}, err
	} else if err == gorm.ErrRecordNotFound {
		err = errors.New("there is no account in users")
		c.Logger.Error(err)
		return &account.ChangeAccountInfoResp{
			Status:  false,
			Message: "There is not account in users",
		}, err
	} else if user.Username != in.UserName {
		if err := c.svcCtx.Gorm.Where("username = ? and id != ?", in.UserName, in.UserId).First(&models.Users{}).Error; err != nil && err != gorm.ErrRecordNotFound {
			err := fmt.Errorf("can not get users record (%v)", err)
			c.Logger.Error(err)
			return &account.ChangeAccountInfoResp{
				Status:  false,
				Message: "Can not get users record",
			}, err
		} else if err != gorm.ErrRecordNotFound {
			err = errors.New("username already exists")
			c.Logger.Info(err)
			return &account.ChangeAccountInfoResp{
				Status:  false,
				Message: "Username already exists",
			}, err
		}
	}

	if err := c.svcCtx.Gorm.Where("user_id = ?", in.UserId).First(&userEllo).Error; err != nil && err != gorm.ErrRecordNotFound {
		err := fmt.Errorf("can not get users_ello record (%v)", err)
		c.Logger.Error(err)
		return &account.ChangeAccountInfoResp{
			Status:  false,
			Message: "Can not get users_ello record",
		}, err
	} else if err == gorm.ErrRecordNotFound {
		err = errors.New("there is no account in users_ello")
		c.Logger.Error(err)
		return &account.ChangeAccountInfoResp{
			Status:  false,
			Message: "There is not account in users",
		}, err
	} else if !userEllo.EmailConfirmed {
		err = errors.New("email not confirmed")
		c.Logger.Error(err)
		return &account.ChangeAccountInfoResp{
			Status:  false,
			Message: "Email not confirmed",
		}, err
	}

	var updateUsername bool

	if user.Username != in.UserName {
		if err := c.svcCtx.Gorm.Where("peer_id = ?", in.UserId).First(&username).Error; err != nil && err != gorm.ErrRecordNotFound {
			err := fmt.Errorf("can not get username record (%v)", err)
			c.Logger.Error(err)
			return &account.ChangeAccountInfoResp{
				Status:  false,
				Message: "Can not get username record",
			}, err
		} else if err == gorm.ErrRecordNotFound {
			err = errors.New("there is no account in username")
			c.Logger.Error(err)
			return &account.ChangeAccountInfoResp{
				Status:  false,
				Message: "There is not account in users",
			}, err
		}
		user.Username = in.UserName
		userEllo.Username = in.UserName
		username.Username = in.UserName

		updateUsername = true
	}

	user.FirstName = in.FirstName
	user.LastName = in.LastName
	user.About = in.Bio
	user.CountryCode = sql.NullString{String: in.CountryCode, Valid: true}
	userEllo.Gender = in.Gender

	if userEllo.Type == "personal" {
		dob, _ := date.FormatDateIso8601(in.Birthday)
		userEllo.DateOfBirth = sql.NullTime{
			Time:  dob,
			Valid: true,
		}
	}

	err = c.svcCtx.Gorm.Transaction(func(tx *gorm.DB) (err error) {
		errResp := &account.ChangeAccountInfoResp{
			Status:  false,
			Message: "Can not update users record",
		}
		if updateUsername {
			if err = c.svcCtx.Gorm.Save(&username).Error; err != nil {
				c.Logger.Errorf("cannot update username record (%v)", err)
				resp = errResp
				return err
			}
		}
		if err = c.svcCtx.Gorm.Save(&user).Error; err != nil {
			c.Logger.Errorf("cannot update users record (%v)", err)
			resp = errResp
			return err
		} else if err = c.svcCtx.Gorm.Save(&userEllo).Error; err != nil {
			c.Logger.Errorf("cannot update users_ello record (%v)", err)
			resp = errResp
			return err
		}
		return
	})

	return resp, err
}
