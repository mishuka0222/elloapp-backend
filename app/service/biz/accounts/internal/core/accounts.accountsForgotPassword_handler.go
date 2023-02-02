package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/accounts/accounts"
)

func (c *AccountsCore) AccountsForgotPassword(in *accounts.ForgotAccountPassReq) (*accounts.ForgotAccountPassResp, error) {
	_, err := c.svcCtx.Dao.UserAccountsDAO.ForgotAccountPassword(c.ctx, c.MD.GetUserId(), in.GetNewPass())
	// todo: add your logic here and delete this line
	if err != nil {
		return &accounts.ForgotAccountPassResp{
			Status:  false,
			Message: "There are some errors to change your email",
		}, nil
	}

	return &accounts.ForgotAccountPassResp{
		Status:  true,
		Message: "Changed Successfully!",
	}, nil
}
