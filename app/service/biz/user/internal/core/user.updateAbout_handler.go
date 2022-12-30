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

// UserUpdateAbout
// user.updateAbout user_id:long about:string = Bool;
func (c *UserCore) UserUpdateAbout(in *user.TLUserUpdateAbout) (*mtproto.Bool, error) {
	rB := c.svcCtx.Dao.UpdateUserAbout(
		c.ctx,
		in.GetUserId(),
		in.GetAbout())

	return mtproto.ToBool(rB), nil
}
