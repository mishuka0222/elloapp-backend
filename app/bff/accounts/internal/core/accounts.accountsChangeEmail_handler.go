package core

import (
	"encoding/json"

	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/accounts/accounts"
)

type ChangeEmailReq struct {
	NewEmail string `json:"new_email"`
}

func (c *AccountsCore) AccountsChangeEmail(in json.RawMessage) (*accounts.ChangeAccountEmailResp, error) {
	var changeEmailReq ChangeEmailReq
	if err := json.Unmarshal(in, &changeEmailReq); err != nil {
		return nil, err
	}

	res, err := c.svcCtx.Dao.AccountsClient.AccountsChangeEmail(c.ctx, &accounts.ChangeAccountEmailReq{
		UserId:   c.MD.GetUserId(),
		NewEmail: changeEmailReq.NewEmail,
	})

	if err != nil {
		return nil, err
	}
	// todo: add your logic here and delete this line
	return res, nil
}
