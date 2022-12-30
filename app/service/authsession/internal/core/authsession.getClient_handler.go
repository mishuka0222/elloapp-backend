/*
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package core

import (
	"fmt"

	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/authsession/authsession"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
)

// AuthsessionGetClient
// authsession.getClient auth_key_id:long = String;
func (c *AuthsessionCore) AuthsessionGetClient(in *authsession.TLAuthsessionGetClient) (*mtproto.String, error) {
	keyData, err := c.svcCtx.Dao.QueryAuthKeyV2(c.ctx, in.GetAuthKeyId())
	if err != nil {
		c.Logger.Errorf("session.getClient - error: %v", err)
		return nil, err
	} else if keyData == nil || keyData.PermAuthKeyId == 0 {
		return nil, fmt.Errorf("not found keyId")
	}

	client := c.svcCtx.Dao.GetClient(c.ctx, keyData.PermAuthKeyId)

	return mtproto.MakeTLString(&mtproto.String{
		V: client,
	}).To_String(), nil
}
