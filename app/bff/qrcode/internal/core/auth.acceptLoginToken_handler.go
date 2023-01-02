package core

import (
	"encoding/binary"
	"strconv"
	"time"

	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/bff/qrcode/internal/model"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/messenger/sync/sync"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/authsession/authsession"
	userpb "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/user/user"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

// AuthAcceptLoginToken
// auth.acceptLoginToken#e894ad4d token:bytes = Authorization;
func (c *QrCodeCore) AuthAcceptLoginToken(in *mtproto.TLAuthAcceptLoginToken) (*mtproto.Authorization, error) {
	// 8 + 16
	if len(in.Token) != 24 {
		err := mtproto.ErrAuthTokenInvalid
		c.Logger.Errorf("auth.acceptLoginToken - error: %v", err)
		return nil, err
	}

	var (
		keyId = int64(binary.BigEndian.Uint64(in.Token))
	)

	qrCode, err := c.svcCtx.Dao.GetCacheQRLoginCode(c.ctx, keyId)
	if err != nil || qrCode == nil {
		err := mtproto.ErrAuthTokenExpired
		c.Logger.Errorf("auth.acceptLoginToken - error: %v", err)
		return nil, err
	}

	c.Logger.Infof("auth.acceptLoginToken - qrCode: %#v", qrCode)

	if !qrCode.CheckByToken(in.Token) {
		err := mtproto.ErrAuthTokenInvalid
		c.Logger.Errorf("auth.acceptLoginToken - error: %v", err)
		return nil, err
	}

	if qrCode.ExpireAt >= time.Now().Unix() {
		//s.AuthCore.DeleteQRCode(ctx, keyId)
		//err := mtproto.ErrAuthTokenExpired
		//log.Errorf("auth.acceptLoginToken - error: %v", err)
		//return nil, err
	}

	switch qrCode.State {
	case model.QRCodeStateNew:
		// ok
	case model.QRCodeStateAccepted, model.QRCodeStateSuccess:
		err := mtproto.ErrAuthTokenAccepted
		c.Logger.Errorf("auth.acceptLoginToken - error: %v", err)
		return nil, err
	default:
		err := mtproto.ErrAuthTokenInvalid
		c.Logger.Errorf("auth.acceptLoginToken - error: %v", err)
		return nil, err
	}

	c.svcCtx.Dao.UpdateCacheQRLoginCode(c.ctx, keyId, map[string]string{
		"user_id": strconv.FormatInt(c.MD.UserId, 10),
		"state":   strconv.Itoa(model.QRCodeStateAccepted),
	})

	user, err := c.svcCtx.Dao.UserClient.UserGetImmutableUser(c.ctx, &userpb.TLUserGetImmutableUser{
		Id: c.MD.UserId,
	})
	if err != nil {
		c.Logger.Errorf("auth.acceptLoginToken - error: %v", err)
		return nil, err
	}

	// Bind authKeyId and userId
	hash, err := c.svcCtx.Dao.AuthsessionClient.AuthsessionBindAuthKeyUser(c.ctx, &authsession.TLAuthsessionBindAuthKeyUser{
		AuthKeyId: qrCode.AuthKeyId,
		UserId:    user.Id(),
	})
	if err != nil {
		c.Logger.Errorf("auth.acceptLoginToken - error: %v", err)
		return nil, err
	}
	authorization, err := c.svcCtx.Dao.AuthsessionClient.AuthsessionGetAuthorization(c.ctx, &authsession.TLAuthsessionGetAuthorization{
		AuthKeyId: qrCode.AuthKeyId,
	})
	if err != nil {
		c.Logger.Errorf("auth.acceptLoginToken - error: %v", err)
		return nil, err
	}

	authorization.DateCreated = int32(time.Now().Unix())
	authorization.DateActive = authorization.DateCreated
	authorization.Hash = hash.V

	c.svcCtx.Dao.SyncClient.SyncUpdatesMe(
		c.ctx,
		&sync.TLSyncUpdatesMe{
			UserId:    user.Id(),
			AuthKeyId: qrCode.AuthKeyId,
			ServerId:  qrCode.ServerId,
			SessionId: mtproto.MakeFlagsInt64(qrCode.SessionId),
			Updates: mtproto.MakeTLUpdateShort(&mtproto.Updates{
				Update: mtproto.MakeTLUpdateLoginToken(nil).To_Update(),
				Date:   int32(time.Now().Unix()),
			}).To_Updates(),
		})

	return authorization, nil
}
