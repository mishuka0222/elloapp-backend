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

// UserUpdatePredefinedUsername
// user.updatePredefinedUsername flags:# phone:string username:flags.1?string = PredefinedUser;
func (c *UserCore) UserUpdatePredefinedUsername(in *user.TLUserUpdatePredefinedUsername) (*mtproto.PredefinedUser, error) {
	// TODO: not impl
	c.Logger.Errorf("user.updatePredefinedUsername - error: method UserUpdatePredefinedUsername not impl")

	return nil, mtproto.ErrMethodNotImpl
}
