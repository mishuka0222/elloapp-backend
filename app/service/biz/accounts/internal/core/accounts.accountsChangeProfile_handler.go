package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/accounts/accounts"
)

func (c *AccountsCore) AccountsChangeProfile(in *accounts.ChangeAccountInfoReq) (*accounts.ChangeAccountInfoResp, error) {
	_, err := c.svcCtx.Dao.UserAccountsDAO.UpdateAccountProfile(c.ctx, c.MD.GetUserId(), in.GetFirstName(), in.GetLastName(), in.GetUserName(), in.GetBio(), in.GetBirthday(), in.GetGender(), in.GetCountryCode())
	// todo: add your logic here and delete this line
	if err != nil {
		return &accounts.ChangeAccountInfoResp{
			Status:  false,
			Message: "There are some errors to update your profile",
		}, nil
	}

	return &accounts.ChangeAccountInfoResp{
		Status:  true,
		Message: "Updated Successfully!",
	}, nil
}
