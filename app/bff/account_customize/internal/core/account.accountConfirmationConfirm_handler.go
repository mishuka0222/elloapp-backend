package core

import (
	"encoding/json"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/account/account"
)

type ConfirmationConfirm struct {
	Email string `json:"email"`
	Code  string `json:"code"`
}

func (c *AccountCore) AccountConfirmationConfirm(in json.RawMessage) (*account.ConfirmationConfirmResp, error) {
	var req ConfirmationConfirm
	if err := json.Unmarshal(in, &req); err != nil {
		return &account.ConfirmationConfirmResp{
			Status:  false,
			Message: "can not unmarshal",
		}, err
	}

	res, err := c.svcCtx.Dao.AccountClient.AccountConfirmationConfirm(c.ctx, &account.ConfirmationConfirmReq{
		Email: req.Email,
		Code:  req.Code,
	})
	
	return res, err
}
