package core

import (
	"encoding/json"

	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/account/account"
)

func (c *AccountCore) AccountDelete(_ json.RawMessage) (*account.DeleteAccountResp, error) {
	res, err := c.svcCtx.Dao.AccountClient.AccountDelete(c.ctx, &account.DeleteAccountReq{
		UserId: c.MD.GetUserId(),
	})

	return res, err
}
