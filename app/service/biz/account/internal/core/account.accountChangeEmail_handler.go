package core

import (
	"errors"
	"fmt"
	"gorm.io/gorm"

	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/account/account"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/models"
)

func (c *AccountCore) AccountChangeEmail(in *account.ChangeAccountEmailReq) (*account.ChangeAccountEmailResp, error) {
	// todo: add your logic here and delete this line
	if err := c.svcCtx.Gorm.Where("user_id = ?", in.UserId).Updates(&models.UsersEllo{Email: in.NewEmail}).Error; err != nil && err != gorm.ErrRecordNotFound {
		err := fmt.Errorf("can not get users record (%v)", err)
		c.Logger.Error(err)
		return &account.ChangeAccountEmailResp{
			Status:  false,
			Message: "Can not get usersEllo record",
		}, err
	} else if err == gorm.ErrRecordNotFound {
		err = errors.New("there is no account in usersEllo")
		c.Logger.Error(err)
		return &account.ChangeAccountEmailResp{
			Status:  false,
			Message: "There is not account in usersEllo",
		}, err
	}
	return &account.ChangeAccountEmailResp{
		Status:  true,
		Message: "Changed Successfully",
	}, nil
}
