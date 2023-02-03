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
	// todo: add your logic here and delete this line
	user := &models.Users{
		FirstName:   in.FirstName,
		LastName:    in.LastName,
		About:       in.Bio,
		CountryCode: sql.NullString{String: in.CountryCode, Valid: true},
	}

	dob, _ := date.FormatDateIso8601(in.Birthday)
	userEllo := &models.UsersEllo{
		Username:    in.UserName,
		DateOfBirth: &dob,
		Gender:      in.Gender,
	}

	if err := c.svcCtx.Gorm.Where("user_id = ?", in.UserId).Updates(user).Error; err != nil && err != gorm.ErrRecordNotFound {
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
	}

	if err := c.svcCtx.Gorm.Where("user_id = ?", in.UserId).Updates(userEllo).Error; err != nil && err != gorm.ErrRecordNotFound {
		err := fmt.Errorf("can not get users record (%v)", err)
		c.Logger.Error(err)
		return &account.ChangeAccountInfoResp{
			Status:  false,
			Message: "Can not get usersEllo record",
		}, err
	} else if err == gorm.ErrRecordNotFound {
		err = errors.New("there is no account in usersEllo")
		c.Logger.Error(err)
		return &account.ChangeAccountInfoResp{
			Status:  false,
			Message: "There is not account in usersEllo",
		}, err
	}

	return &account.ChangeAccountInfoResp{
		Status:  true,
		Message: "Updated successfully",
	}, nil
}
