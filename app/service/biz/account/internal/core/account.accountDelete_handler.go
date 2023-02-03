package core

import (
	"errors"
	"fmt"

	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/account/account"
	"gorm.io/gorm"
)

func (c *AccountCore) AccountDelete(in *account.DeleteAccountReq) (*account.DeleteAccountResp, error) {
	// todo: add your logic here and delete this line
	if err := c.svcCtx.Gorm.Where("user_id = ?", in.UserId).Debug().Error; err != nil && err != gorm.ErrRecordNotFound {
		err := fmt.Errorf("can not get users record (%v)", err)
		c.Logger.Error(err)
		return &account.DeleteAccountResp{
			Status:  false,
			Message: "Can not get users record",
		}, err
	} else if err == gorm.ErrRecordNotFound {
		err = errors.New("there is no account in users")
		c.Logger.Error(err)
		return &account.DeleteAccountResp{
			Status:  false,
			Message: "There is not account in users",
		}, err
	}

	return &account.DeleteAccountResp{
		Status:  true,
		Message: "Deleted Successfully",
	}, nil
}
