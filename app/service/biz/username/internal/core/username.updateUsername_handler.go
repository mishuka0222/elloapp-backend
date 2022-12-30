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
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/username/internal/dal/dataobject"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/username/username"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/pkg2/stores/sqlx"
)

// UsernameUpdateUsername
// username.updateUsername peer_type:int peer_id:int username:string = Bool;
func (c *UsernameCore) UsernameUpdateUsername(in *username.TLUsernameUpdateUsername) (*mtproto.Bool, error) {
	_, _, err := c.svcCtx.Dao.UsernameDAO.Insert(c.ctx, &dataobject.UsernameDO{
		Username: in.Username,
		PeerType: in.PeerType,
		PeerId:   in.PeerId,
	})
	if err != nil {
		if sqlx.IsDuplicate(err) {
			return mtproto.BoolFalse, nil
		} else {
			c.Logger.Errorf("username.updateUsername - error: %v", err)
			return nil, err
		}
	}

	return mtproto.BoolTrue, nil
}
