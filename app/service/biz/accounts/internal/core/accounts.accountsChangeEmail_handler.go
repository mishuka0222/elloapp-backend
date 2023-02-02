package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/accounts/accounts"
)

func (c *AccountsCore) AccountsChangeEmail(in *accounts.ChangeAccountEmailReq) (*accounts.ChangeAccountEmailResp, error) {
	_, err := c.svcCtx.Dao.UserAccountsDAO.ChangeAccountEmail(c.ctx, c.MD.GetUserId(), in.GetNewEmail())
	// todo: add your logic here and delete this line
	if err != nil {
		return &accounts.ChangeAccountEmailResp{
			Status:  false,
			Message: "There are some errors to change your email",
		}, nil
	}

	return &accounts.ChangeAccountEmailResp{
		Status:  true,
		Message: "Changed Successfully!",
	}, nil
}
