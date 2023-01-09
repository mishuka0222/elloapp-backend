package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/authsession/authsession"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

const (
	kDefaultSaltNum = 32
)

// AuthsessionGetFutureSalts
// authsession.getFutureSalts auth_key_id:long num:int = FutureSalts;
func (c *AuthsessionCore) AuthsessionGetFutureSalts(in *authsession.TLAuthsessionGetFutureSalts) (*mtproto.FutureSalts, error) {
	num := in.GetNum()
	if num == 0 {
		num = kDefaultSaltNum
	}
	futureSalts, err := c.svcCtx.Dao.GetFutureSalts(c.ctx, in.GetAuthKeyId(), num)
	if err != nil {
		c.Logger.Errorf("session.getFutureSalts - %v", err)
		return nil, err
	}

	return futureSalts.To_FutureSalts(), nil
}
