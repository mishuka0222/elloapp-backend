package core

import (
	"encoding/json"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/account/account"
)

type ConfirmationSend struct {
	Email string `json:"email"`
}

func (c *AccountCore) AccountConfirmationSend(in json.RawMessage) (*account.ConfirmationSendResp, error) {
	var req ConfirmationSend
	if err := json.Unmarshal(in, &req); err != nil {
		return &account.ConfirmationSendResp{
			Status:  false,
			Message: "can not unmarshal",
		}, err
	}

	res, err := c.svcCtx.Dao.AccountClient.AccountConfirmationSend(c.ctx, &account.ConfirmationSendReq{
		Email: req.Email,
	})
	
	return res, err
}
