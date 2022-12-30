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

// UserUpdatePredefinedFirstAndLastName
// user.updatePredefinedFirstAndLastName flags:# phone:string first_name:string last_name:flags.0?string = PredefinedUser;
func (c *UserCore) UserUpdatePredefinedFirstAndLastName(in *user.TLUserUpdatePredefinedFirstAndLastName) (*mtproto.PredefinedUser, error) {
	// TODO: not impl
	c.Logger.Errorf("user.updatePredefinedFirstAndLastName - error: method UserUpdatePredefinedFirstAndLastName not impl")

	return nil, mtproto.ErrMethodNotImpl
}
