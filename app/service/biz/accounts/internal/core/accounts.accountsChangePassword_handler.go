package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/accounts/accounts"
)

func (c *AccountsCore) AccountsChangePassword(in *accounts.ChangeAccountPasswordReq) (*accounts.ChangeAccountPasswordResp, error) {
	status, err := c.svcCtx.Dao.UserAccountsDAO.ChangeAccountPassword(c.ctx, c.MD.GetUserId(), in.GetPrevPass(), in.GetNewPass())
	// todo: add your logic here and delete this line
	if err != nil {
		return &accounts.ChangeAccountPasswordResp{
			Status:  false,
			Message: "There are some errors to change your password",
		}, nil
	}

	if !status {
		return &accounts.ChangeAccountPasswordResp{
			Status:  false,
			Message: "Incorrect Password!",
		}, nil
	}

	return &accounts.ChangeAccountPasswordResp{
		Status:  true,
		Message: "Changed Successfully!",
	}, nil
}
