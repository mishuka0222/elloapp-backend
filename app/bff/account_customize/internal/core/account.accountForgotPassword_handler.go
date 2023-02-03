package core

import (
	"encoding/json"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/account/account"
)

type ForgotPassword struct {
	NewPass string `json:"new_pass"`
}

func (c *AccountCore) AccountForgotPassword(in json.RawMessage) (*account.ForgotAccountPassResp, error) {
	var req ForgotPassword
	if err := json.Unmarshal(in, &req); err != nil {
		return &account.ForgotAccountPassResp{
			Status:  false,
			Message: "can not unmarshal",
		}, err
	}

	res, err := c.svcCtx.Dao.AccountClient.AccountForgotPassword(c.ctx, &account.ForgotAccountPassReq{
		UserId:  c.MD.GetUserId(),
		NewPass: req.NewPass,
	})

	return res, err
}
