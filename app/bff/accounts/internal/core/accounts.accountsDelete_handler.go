package core

import (
	"encoding/json"

	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/accounts/accounts"
)

func (c *AccountsCore) AccountsDelete(_ json.RawMessage) (*accounts.DeleteAccountResp, error) {
	res, err := c.svcCtx.Dao.AccountsClient.AccountsDelete(c.ctx, &accounts.DeleteAccountReq{
		UserId: c.MD.GetUserId(),
	})

	if err != nil {
		return nil, err
	}
	// todo: add your logic here and delete this line
	return res, nil
}
