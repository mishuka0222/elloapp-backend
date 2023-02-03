package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/account/account"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

// AccountDeleteAccount
// account.deleteAccount#418d4e0b reason:string = Bool;
func (c *AccountCore) AccountDeleteAccount(_ *mtproto.TLAccountDeleteAccount) (*mtproto.Bool, error) {
	_, err := c.svcCtx.Dao.AccountClient.AccountDelete(c.ctx, &account.DeleteAccountReq{
		UserId: c.MD.GetUserId(),
	})

	if err != nil {
		c.Logger.Errorf("account.deleteAccount - error: %v", err)
		return mtproto.BoolFalse, err
	}

	return mtproto.BoolTrue, nil
}
