package core

import (
	"encoding/json"

	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/accounts/accounts"
)

type ChangePassReq struct {
	PrevPass string `json:"prev_pass"`
	NewPass  string `json:"new_pass"`
}

func (c *AccountsCore) AccountsChangePassword(in json.RawMessage) (*accounts.ChangeAccountPasswordResp, error) {
	var changePassReq ChangePassReq
	if err := json.Unmarshal(in, &changePassReq); err != nil {
		return nil, err
	}

	res, err := c.svcCtx.Dao.AccountsClient.AccountsChangePassword(c.ctx, &accounts.ChangeAccountPasswordReq{
		UserId:   c.MD.GetUserId(),
		PrevPass: changePassReq.PrevPass,
		NewPass:  changePassReq.NewPass,
	})

	if err != nil {
		return nil, err
	}
	// todo: add your logic here and delete this line
	return res, nil
}
