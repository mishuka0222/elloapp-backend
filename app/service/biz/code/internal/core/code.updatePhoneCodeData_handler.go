package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/code/code"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
)

// CodeUpdatePhoneCodeData
// code.updatePhoneCodeData auth_key_id:long phone:string phone_code_hash:string code_data:PhoneCodeTransaction = Bool;
func (c *CodeCore) CodeUpdatePhoneCodeData(in *code.TLCodeUpdatePhoneCodeData) (*mtproto.Bool, error) {
	if err := c.svcCtx.Dao.PutCachePhoneCode(c.ctx, in.AuthKeyId, in.PhoneCodeHash, in.CodeData); err != nil {
		c.Logger.Errorf("code.updatePhoneCodeData - error: %v", err)
		return nil, mtproto.ErrInternelServerError
	}

	return mtproto.BoolTrue, nil
}
