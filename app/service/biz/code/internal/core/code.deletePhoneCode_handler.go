package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/code/code"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
)

// CodeDeletePhoneCode
// code.deletePhoneCode auth_key_id:long phone:string phone_code_hash:string = Bool;
func (c *CodeCore) CodeDeletePhoneCode(in *code.TLCodeDeletePhoneCode) (*mtproto.Bool, error) {
	if err := c.svcCtx.Dao.DeleteCachePhoneCode(c.ctx, in.AuthKeyId, in.Phone); err != nil {
		c.Logger.Errorf("code.deletePhoneCode - error: %v", err)
	}

	return mtproto.BoolTrue, nil
}
