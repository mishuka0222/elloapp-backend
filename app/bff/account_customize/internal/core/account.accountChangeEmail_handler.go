package core

import (
	"encoding/json"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/account/account"
)

type ChangeEmail struct {
	NewEmail string `json:"new_email"`
}

func (c *AccountCore) AccountChangeEmail(in json.RawMessage) (*account.ChangeAccountEmailResp, error) {
	var req ChangeEmail
	if err := json.Unmarshal(in, &req); err != nil {
		return &account.ChangeAccountEmailResp{
			Status:  false,
			Message: "can not unmarshal",
		}, err
	}

	res, err := c.svcCtx.Dao.AccountClient.AccountChangeEmail(c.ctx, &account.ChangeAccountEmailReq{
		UserId:   c.MD.GetUserId(),
		NewEmail: req.NewEmail,
	})

	return res, err
}
