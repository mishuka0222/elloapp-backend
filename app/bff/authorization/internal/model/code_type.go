package model

import (
	"fmt"

	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"

	"github.com/zeromicro/go-zero/core/logx"
)

/**
  auth.codeTypeSms#72a3158c = auth.CodeType;
  auth.codeTypeCall#741cd3e3 = auth.CodeType;
  auth.codeTypeFlashCall#226ccefb = auth.CodeType;

  auth.sentCodeTypeApp#3dbb5986 length:int = auth.SentCodeType;
  auth.sentCodeTypeSms#c000bba2 length:int = auth.SentCodeType;
  auth.sentCodeTypeCall#5353e5a7 length:int = auth.SentCodeType;
  auth.sentCodeTypeFlashCall#ab03c6d9 pattern:string = auth.SentCodeType;
*/

const (
	CodeTypeNone      = 0
	CodeTypeApp       = 1
	CodeTypeSms       = 2
	CodeTypeCall      = 3
	CodeTypeFlashCall = 4
)

const (
	CodeStateOk      = 1
	CodeStateSend    = 2
	CodeStateSent    = 3
	CodeStateReSent  = 6
	CodeStateSignIn  = 4
	CodeStateSignUp  = 5
	CodeStateDeleted = -1
)

/////////////////////////////////////////////////////////////////////////////////////////////////////////
// by params(phoneRegistered, allowFlashCall, currentNumber) ==> sentType and nextType

// FIXME(@benqi): ignore it.
func MakeCodeType(phoneRegistered, allowFlashCall, currentNumber bool) (int, int) {
	//if phoneRegistered {
	//	// TODO(@benqi): check other session online
	//	authSentCodeType := &mtproto.TLAuthSentCodeTypeApp{Data2: &mtproto.Auth_SentCodeType_Data{
	//		Length: code.GetPhoneCodeLength(),
	//	}}
	//	authSentCode.SetType(authSentCodeType.To_Auth_SentCodeType())
	//} else {
	//	// TODO(@benqi): sentCodeTypeFlashCall and sentCodeTypeCall, nextType
	//	// telegramd, we only use sms
	//	authSentCodeType := &mtproto.TLAuthSentCodeTypeSms{Data2: &mtproto.Auth_SentCodeType_Data{
	//		Length: code.GetPhoneCodeLength(),
	//	}}
	//	authSentCode.SetType(authSentCodeType.To_Auth_SentCodeType())
	//
	//	// TODO(@benqi): nextType
	//	// authSentCode.SetNextType()
	//}
	_ = phoneRegistered
	_ = allowFlashCall
	_ = currentNumber

	sentCodeType := CodeTypeApp
	nextCodeType := CodeTypeNone
	return sentCodeType, nextCodeType
}

func makeAuthCodeType(codeType int) *mtproto.Auth_CodeType {
	switch codeType {
	case CodeTypeSms:
		return mtproto.MakeTLAuthCodeTypeSms(nil).To_Auth_CodeType()
	case CodeTypeCall:
		return mtproto.MakeTLAuthCodeTypeCall(nil).To_Auth_CodeType()
	case CodeTypeFlashCall:
		return mtproto.MakeTLAuthCodeTypeFlashCall(nil).To_Auth_CodeType()
	default:
		return nil
	}
}

func makeAuthSentCodeType(codeType, codeLength int, pattern string) (authSentCodeType *mtproto.Auth_SentCodeType) {
	switch codeType {
	case CodeTypeApp:
		authSentCodeType = mtproto.MakeTLAuthSentCodeTypeApp(&mtproto.Auth_SentCodeType{
			Length: int32(codeLength),
		}).To_Auth_SentCodeType()
	case CodeTypeSms:
		authSentCodeType = mtproto.MakeTLAuthSentCodeTypeSms(&mtproto.Auth_SentCodeType{
			Length: int32(codeLength),
		}).To_Auth_SentCodeType()
	case CodeTypeCall:
		authSentCodeType = mtproto.MakeTLAuthSentCodeTypeCall(&mtproto.Auth_SentCodeType{
			Length: int32(codeLength),
		}).To_Auth_SentCodeType()
	case CodeTypeFlashCall:
		authSentCodeType = mtproto.MakeTLAuthSentCodeTypeFlashCall(&mtproto.Auth_SentCodeType{
			Length:  int32(codeLength),
			Pattern: pattern,
		}).To_Auth_SentCodeType()
	default:
		// code bug.
		err := fmt.Errorf("invalid sentCodeType: %d", codeType)
		logx.Errorf("makeAuthSentCodeType - %v", err)
		panic(err)
	}

	return
}
