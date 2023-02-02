package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/accounts/accounts"
)

func (c *AccountsCore) AccountsDelete(in *accounts.DeleteAccountReq) (*accounts.DeleteAccountResp, error) {
	_, err := c.svcCtx.Dao.UserAccountsDAO.DeleteUserAccount(c.ctx, in.UserId)
	// todo: add your logic here and delete this line
	if err != nil {
		return &accounts.DeleteAccountResp{
			Status:  false,
			Message: "There are some errors to delete your account",
		}, nil
	}
	return &accounts.DeleteAccountResp{
		Status:  true,
		Message: "Deleted Successfully!",
	}, nil
}
