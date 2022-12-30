package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/bff/qrcode/internal/model"
	userpb "gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/user/user"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto/crypto"
	"time"
)

const (
	qrCodeTimeout = 60 // salt timeout
)

// AuthExportLoginToken
// auth.exportLoginToken#b7e085fe api_id:int api_hash:string except_ids:Vector<long> = auth.LoginToken;
func (c *QrCodeCore) AuthExportLoginToken(in *mtproto.TLAuthExportLoginToken) (*mtproto.Auth_LoginToken, error) {
	qrCode, err := c.svcCtx.Dao.GetCacheQRLoginCode(c.ctx, c.MD.AuthId)
	if err != nil {
		c.Logger.Errorf("getQRCode - error: %v", err)
		return nil, err
	} else if qrCode == nil {
		qrCode = &model.QRCodeTransaction{
			AuthKeyId: c.MD.AuthId,
			SessionId: c.MD.SessionId,
			ServerId:  c.MD.ServerId,
			ApiId:     in.ApiId,
			ApiHash:   in.ApiHash,
			CodeHash:  crypto.GenerateStringNonce(16),
			ExpireAt:  time.Now().Unix() + qrCodeTimeout,
			UserId:    0,
			State:     model.QRCodeStateNew,
		}
		c.Logger.Infof("putQRCode - %#v", qrCode)
		if err = c.svcCtx.Dao.PutCacheQRLoginCode(c.ctx, c.MD.AuthId, qrCode, qrCodeTimeout+2); err != nil {
			c.Logger.Errorf("putQRCode - error: %v", err)
			return nil, err
		}
	} else {
		c.Logger.Infof("putQRCode - %#v", qrCode)
	}

	var (
		rQRLoginToken *mtproto.Auth_LoginToken
	)

	switch qrCode.State {
	case model.QRCodeStateAccepted, model.QRCodeStateSuccess:
		//// Check SESSION_PASSWORD_NEEDED
		//if sessionPasswordNeeded, _ := c.svcCtx.Dao.TwofaClient.TwofaCheckSessionPasswordNeeded(c.ctx, &twofa.TLTwofaCheckSessionPasswordNeeded{
		//	UserId: qrCode.UserId,
		//}); mtproto.FromBool(sessionPasswordNeeded) {
		//	err = mtproto.ErrSessionPasswordNeeded
		//	c.Logger.Infof("auth.exportLoginToken - registered, next step auth.checkPassword: %v", err)
		//	return nil, err
		//}

		user, err := c.svcCtx.Dao.UserClient.UserGetImmutableUser(c.ctx, &userpb.TLUserGetImmutableUser{
			Id: qrCode.UserId,
		})
		if err != nil {
			c.Logger.Errorf("auth.exportLoginToken - error: %v", err)
			return nil, err
		}
		rQRLoginToken = mtproto.MakeTLAuthLoginTokenSuccess(&mtproto.Auth_LoginToken{
			Authorization: mtproto.MakeTLAuthAuthorization(&mtproto.Auth_Authorization{
				TmpSessions: nil,
				User:        user.ToSelfUser(),
			}).To_Auth_Authorization(),
		}).To_Auth_LoginToken()
	default:
		rQRLoginToken = mtproto.MakeTLAuthLoginToken(&mtproto.Auth_LoginToken{
			Expires: int32(qrCode.ExpireAt),
			Token:   qrCode.Token(),
		}).To_Auth_LoginToken()
	}

	return rQRLoginToken, nil
}
