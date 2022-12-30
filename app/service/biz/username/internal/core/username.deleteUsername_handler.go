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
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/username/username"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
)

// UsernameDeleteUsername
// username.deleteUsername username:string = Bool;
func (c *UsernameCore) UsernameDeleteUsername(in *username.TLUsernameDeleteUsername) (*mtproto.Bool, error) {
	_, err := c.svcCtx.Dao.UsernameDAO.Delete(c.ctx, in.Username)
	if err != nil {
		return mtproto.BoolFalse, nil
	}

	return mtproto.BoolTrue, nil
}
