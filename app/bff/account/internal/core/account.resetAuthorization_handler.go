package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/messenger/sync/sync"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/authsession/authsession"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
)

// AccountResetAuthorization
// account.resetAuthorization#df77f3bc hash:long = Bool;
func (c *AccountCore) AccountResetAuthorization(in *mtproto.TLAccountResetAuthorization) (*mtproto.Bool, error) {
	if in.Hash == 0 {
		c.Logger.Errorf("account.resetAuthorization#df77f3bc - hash is 0")
		return mtproto.BoolFalse, nil
	}

	tKeyIdList, err := c.svcCtx.Dao.AuthsessionClient.AuthsessionResetAuthorization(c.ctx, &authsession.TLAuthsessionResetAuthorization{
		UserId:    c.MD.UserId,
		AuthKeyId: c.MD.AuthId,
		Hash:      in.Hash,
	})

	if err != nil {
		c.Logger.Errorf("account.resetAuthorization#df77f3bc - error: %v", err)
		return nil, err
	}

	for _, id := range tKeyIdList.Datas {
		// notify kill session
		c.svcCtx.Dao.SyncClient.SyncUpdatesMe(
			c.ctx,
			&sync.TLSyncUpdatesMe{
				UserId:    c.MD.UserId,
				AuthKeyId: id,
				ServerId:  "",
				SessionId: nil,
				Updates: mtproto.MakeTLUpdateAccountResetAuthorization(&mtproto.Updates{
					UserId:    c.MD.UserId,
					AuthKeyId: id,
				}).To_Updates(),
			})
	}

	return mtproto.BoolTrue, nil
}
