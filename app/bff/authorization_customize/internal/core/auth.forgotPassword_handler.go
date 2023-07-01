package core

import (
	"encoding/json"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/authorization/authorization"
)

type ForgotPasswordReq struct {
	Email string `json:"email"`
}

func (c *AuthorizationCore) ForgotPassword(in json.RawMessage) (*authorization.ForgotPasswordRsp, error) {
	var req ForgotPasswordReq
	if err := json.Unmarshal(in, &req); err != nil {
		return nil, err
	}

	res, err := c.svcCtx.Dao.AuthorizationClient.ForgotPassword(c.ctx, &authorization.ForgotPasswordReq{
		Email: req.Email,
	})

	return res, err
}
