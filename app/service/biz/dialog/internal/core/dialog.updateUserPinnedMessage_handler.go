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
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/dialog/dialog"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
)

// DialogUpdateUserPinnedMessage
// dialog.updateUserPinnedMessage user_id:long peer_type:int peer_id:long pinned_msg_id:int = Bool;
func (c *DialogCore) DialogUpdateUserPinnedMessage(in *dialog.TLDialogUpdateUserPinnedMessage) (*mtproto.Bool, error) {
	// TODO: not impl
	c.Logger.Errorf("dialog.updateUserPinnedMessage - error: method DialogUpdateUserPinnedMessage not impl")

	return mtproto.BoolTrue, nil
}
