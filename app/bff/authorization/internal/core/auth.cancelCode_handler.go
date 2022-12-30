package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/pkg/phonenumber"
)

// AuthCancelCode
// auth.cancelCode#1f040578 phone_number:string phone_code_hash:string = Bool;
func (c *AuthorizationCore) AuthCancelCode(in *mtproto.TLAuthCancelCode) (*mtproto.Bool, error) {
	//// 1. check phone code
	//if request.PhoneCodeHash == "" {
	//	err := mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_PHONE_CODE_EMPTY)
	//	return nil, err
	//}
	//
	// 2. check number
	// 客户端发送的手机号格式为: "+86 111 1111 1111"，归一化
	_, err := phonenumber.CheckAndGetPhoneNumber(in.GetPhoneNumber())
	if err != nil {
		c.Logger.Errorf("check phone_number error - %v", err)
		err = mtproto.ErrPhoneNumberInvalid
		return nil, err
	}

	// code := logic.NewAuthSignLogic(s.AuthFacade)
	// canceled := mtproto.ToBool(code.DoAuthCancelCode(md.AuthId, phoneNumber, request.PhoneCodeHash))

	return mtproto.BoolTrue, nil
}
