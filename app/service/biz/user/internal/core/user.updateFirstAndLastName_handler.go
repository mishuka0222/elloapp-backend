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

// UserUpdateFirstAndLastName
// user.updateFirstAndLastName user_id:long first_name:string last_name:string = Bool;
func (c *UserCore) UserUpdateFirstAndLastName(in *user.TLUserUpdateFirstAndLastName) (*mtproto.Bool, error) {
	rB := c.svcCtx.Dao.UpdateUserFirstAndLastName(
		c.ctx,
		in.GetUserId(),
		in.GetFirstName(),
		in.GetLastName())

	return mtproto.ToBool(rB), nil
}
