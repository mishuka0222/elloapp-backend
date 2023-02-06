package core

import (
	"encoding/json"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/authorization/authorization"
)

type ForgotPasswordConfirmReq struct {
	Email   string `json:"email"`
	Code    string `json:"code"`
	NewPass string `json:"new_pass"`
}

func (c *AuthorizationCore) ForgotPasswordConfirm(in json.RawMessage) (*authorization.ForgotPasswordConfirmRsp, error) {
	var req ForgotPasswordConfirmReq
	if err := json.Unmarshal(in, &req); err != nil {
		return nil, err
	}

	res, err := c.svcCtx.Dao.AuthorizationClient.ForgotPasswordConfirm(c.ctx, &authorization.ForgotPasswordConfirmReq{
		Email:   req.Email,
		Code:    req.Code,
		NewPass: req.NewPass,
	})

	return res, err
}
