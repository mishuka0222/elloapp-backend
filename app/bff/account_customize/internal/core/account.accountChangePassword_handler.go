package core

import (
	"encoding/json"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/account/account"
)

type ChangePassword struct {
	PrevPass string `json:"prev_pass"`
	NewPass  string `json:"new_pass"`
}

func (c *AccountCore) AccountChangePassword(in json.RawMessage) (*account.ChangeAccountPasswordResp, error) {
	var req ChangePassword
	if err := json.Unmarshal(in, &req); err != nil {
		return &account.ChangeAccountPasswordResp{
			Status:  false,
			Message: "can not unmarshal",
		}, err
	}

	res, err := c.svcCtx.Dao.AccountClient.AccountChangePassword(c.ctx, &account.ChangeAccountPasswordReq{
		UserId:   c.MD.GetUserId(),
		PrevPass: req.PrevPass,
		NewPass:  req.NewPass,
	})
	return res, err
}
