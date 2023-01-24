package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

// AccountDeleteAccount
// account.deleteAccount#418d4e0b reason:string = Bool;
func (c *AccountCore) AccountDeleteAccount(in *mtproto.TLAccountDeleteAccount) (*mtproto.Bool, error) {
	// TODO: not impl
	success, err := c.svcCtx.Dao.AccountClient.AccountDeleteAccount(c.ctx, in)

	if err != nil {
		c.Logger.Errorf("account.deleteAccount - error: %v", err)
		return nil, err
	}

	return success, nil
	// c.Logger.Errorf("account.deleteAccount blocked, License key from https://elloapp.com required to unlock enterprise features.")
	// return nil, mtproto.ErrEnterpriseIsBlocked
}
