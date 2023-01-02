package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/authsession/authsession"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
)

// AuthsessionResetAuthorization
// authsession.resetAuthorization user_id:long auth_key_id:long hash:long = Vector<long>;
func (c *AuthsessionCore) AuthsessionResetAuthorization(in *authsession.TLAuthsessionResetAuthorization) (*authsession.Vector_Long, error) {
	var (
		excludeKeyId = in.AuthKeyId
	)

	if excludeKeyId != 0 {
		myKeyData, err := c.svcCtx.Dao.QueryAuthKeyV2(c.ctx, in.AuthKeyId)
		if err != nil {
			c.Logger.Errorf("session.getAuthorizations - error: %v", err)
			return nil, err
		} else if myKeyData == nil {
			c.Logger.Errorf("session.getAuthorizations - error: %v", err)
			err = mtproto.ErrAuthKeyInvalid
			return nil, err
		} else {
			excludeKeyId = myKeyData.PermAuthKeyId
		}
	}

	keyIdList := c.svcCtx.Dao.ResetAuthorization(c.ctx, in.UserId, excludeKeyId, in.Hash)
	// log.Debugf("keyIdList: %v", keyIdList)

	keyIdL2ist := make([]int64, 0, len(keyIdList))
	for _, keyId := range keyIdList {
		keyData, _ := c.svcCtx.Dao.QueryAuthKeyV2(c.ctx, keyId)
		if keyData != nil {
			if keyData.TempAuthKeyId != 0 {
				keyIdL2ist = append(keyIdL2ist, keyData.TempAuthKeyId)
			} else {
				keyIdL2ist = append(keyIdL2ist, keyId)
			}
		}
	}

	return &authsession.Vector_Long{
		Datas: keyIdList,
	}, nil
}
