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
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/user/user"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
)

// UserUpdateUsername
// user.updateUsername user_id:long username:string = Bool;
func (c *UserCore) UserUpdateUsername(in *user.TLUserUpdateUsername) (*mtproto.Bool, error) {
	rB := c.svcCtx.Dao.UpdateUserUsername(
		c.ctx,
		in.GetUserId(),
		in.GetUsername())

	return mtproto.ToBool(rB), nil
}
