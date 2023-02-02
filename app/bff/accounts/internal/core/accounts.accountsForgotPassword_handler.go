package core

import (
	"encoding/json"

	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/accounts/accounts"
)

type ForgotPasswordReq struct {
	NewPass string `json:"new_pass"`
}

func (c *AccountsCore) AccountsForgotPassword(in json.RawMessage) (*accounts.ForgotAccountPassResp, error) {
	var forgotPasswordReq ForgotPasswordReq
	if err := json.Unmarshal(in, &forgotPasswordReq); err != nil {
		return nil, err
	}

	res, err := c.svcCtx.Dao.AccountsClient.AccountsForgotPassword(c.ctx, &accounts.ForgotAccountPassReq{
		UserId:  c.MD.GetUserId(),
		NewPass: forgotPasswordReq.NewPass,
	})

	if err != nil {
		return nil, err
	}
	// todo: add your logic here and delete this line
	return res, nil
}
