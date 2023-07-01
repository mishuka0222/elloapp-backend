package core

import (
	"fmt"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/authsession/authsession"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

// AuthsessionSetClientSessionInfo
// authsession.setClientSessionInfo data:ClientSession = Bool;
func (c *AuthsessionCore) AuthsessionSetClientSessionInfo(in *authsession.TLAuthsessionSetClientSessionInfo) (*mtproto.Bool, error) {
	clientSession := in.GetData()
	if clientSession == nil {
		err := mtproto.ErrInputRequestInvalid
		c.Logger.Errorf("session.setClientSessionInfo - error: %v", err)
		return nil, err
	}

	keyData, err := c.svcCtx.Dao.QueryAuthKeyV2(c.ctx, clientSession.GetAuthKeyId())
	if err != nil {
		c.Logger.Errorf("session.setClientSessionInfo - error: %v", err)
		return nil, err
	} else if keyData == nil || keyData.PermAuthKeyId == 0 {
		return nil, fmt.Errorf("not found keyId")
	}

	clientSession.AuthKeyId = keyData.PermAuthKeyId
	r := c.svcCtx.Dao.SetClientSessionInfo(c.ctx, clientSession)

	//c.svcCtx.Dao.CachedConn.Exec(
	//	c.ctx,
	//	func(ctx context.Context, conn *sqlx.DB) (int64, int64, error) {
	//
	//		return 0, 0, nil
	//	},
	//	"")

	return mtproto.ToBool(r), nil
}
