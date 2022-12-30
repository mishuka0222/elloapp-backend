package dao

import (
	"context"
	"time"

	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/bff/authorization/internal/model"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto/crypto"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/pkg2/random2"

	"github.com/zeromicro/go-zero/core/logx"
)

// CreatePhoneCode
// for sendCode
func (d *Dao) CreatePhoneCode(ctx context.Context,
	authKeyId int64,
	sessionId int64,
	phoneNumber string,
	phoneNumberRegistered bool,
	sendCodeType, nextCodeType, state int) (codeData *model.PhoneCodeTransaction, err error) {
	_ = state
	newCodeData := func() *model.PhoneCodeTransaction {
		return &model.PhoneCodeTransaction{
			AuthKeyId:             authKeyId,
			PhoneNumber:           phoneNumber,
			SessionId:             sessionId,
			PhoneNumberRegistered: phoneNumberRegistered,
			PhoneCode:             random2.RandomNumeric(5),
			PhoneCodeHash:         crypto.GenerateStringNonce(16),
			PhoneCodeExpired:      int32(time.Now().Unix() + 3*60),
			SentCodeType:          sendCodeType,
			FlashCallPattern:      "*",
			NextCodeType:          nextCodeType,
			State:                 model.CodeStateSend, // model.CodeStateSent
		}
	}

	if codeData, err = d.GetCachePhoneCode(ctx, authKeyId, phoneNumber); err != nil {
		logx.WithContext(ctx).Errorf("getCachePhoneCode - error: %v", err)
		err = mtproto.ErrInternelServerError
		return
	}
	if codeData == nil {
		codeData = newCodeData()
	} else if sessionId != codeData.SessionId {
		codeData.State = model.CodeStateSend
		codeData.SessionId = sessionId
	}

	//switch codeData.State {
	//case model.CodeStateSend:
	//	codeData.State = model.CodeStateSent
	//case model.CodeStateSent:
	//	codeData.State = model.CodeStateSent
	//default:
	//	// codeData = newCodeData()
	//}
	//
	//if err = c.Dao.PutCachePhoneCode(ctx, authKeyId, phoneNumber, codeData); err != nil {
	//	log.Errorf("putCachePhoneCode - error: %v", err)
	//	err = mtproto.ErrInternelServerError
	//	return
	//}
	return
}

func (d *Dao) GetPhoneCode(ctx context.Context,
	authKeyId int64,
	phoneNumber, phoneCodeHash string) (codeData *model.PhoneCodeTransaction, err error) {

	if codeData, err = d.GetCachePhoneCode(ctx, authKeyId, phoneNumber); err != nil {
		logx.WithContext(ctx).Errorf("getPhoneCode by {authKeyId: %d, phoneNumber: %s} error - %v", authKeyId, phoneNumber, err)
		err = mtproto.ErrPhoneCodeExpired
		return
	} else if codeData == nil {
		logx.WithContext(ctx).Errorf("getPhoneCode by {authKeyId: %d, phoneNumber: %s} error - %v", authKeyId, phoneNumber, err)
		err = mtproto.ErrPhoneCodeExpired
		return
	} else if codeData.PhoneCodeHash != phoneCodeHash {
		logx.WithContext(ctx).Errorf("getPhoneCode by {authKeyId: %d, phoneNumber: %s, phoneCodeHash: %s} error - invalid phoneCodeHash",
			authKeyId,
			phoneNumber,
			phoneCodeHash)
		err = mtproto.ErrPhoneCodeInvalid
	}
	return
}

func (d *Dao) DeletePhoneCode(ctx context.Context, authKeyId int64, phoneNumber, phoneCodeHash string) error {
	return d.DeleteCachePhoneCode(ctx, authKeyId, phoneNumber)
}

func (d *Dao) UpdatePhoneCodeData(ctx context.Context,
	authKeyId int64,
	phoneNumber, phoneCodeHash string,
	codeData *model.PhoneCodeTransaction) error {
	// TODO(@benqi): check state??
	return d.PutCachePhoneCode(ctx, authKeyId, phoneNumber, codeData)
}

func (d *Dao) CheckCanDoAction(ctx context.Context,
	authKeyId int64,
	phoneNumber string,
	actionType int) error {
	// TODO(@benqi): check can do action

	_ = authKeyId
	_ = phoneNumber
	_ = actionType
	return nil
}

//func (d *Dao) LogAuthAction(ctx context.Context,
//	authKeyId, msgId int64,
//	clientIp string,
//	phoneNumber string,
//	actionType int, log string) {
//	_ = phoneNumber
//	do := &dataobject.AuthOpLogsDO{
//		AuthKeyId: authKeyId,
//		Ip:        clientIp,
//		OpType:    int32(actionType),
//		LogText:   log,
//	}
//	d.AuthOpLogsDAO.Insert(ctx, do)
//}
