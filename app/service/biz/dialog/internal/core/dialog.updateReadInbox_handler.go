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
	"github.com/gogo/protobuf/types"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/dialog/dialog"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
)

// DialogUpdateReadInbox
// dialog.updateReadInbox user_id:long peer_type:int peer_id:long read_inbox_id:int = Bool;
func (c *DialogCore) DialogUpdateReadInbox(in *dialog.TLDialogUpdateReadInbox) (*mtproto.Bool, error) {
	c.DialogInsertOrUpdateDialog(&dialog.TLDialogInsertOrUpdateDialog{
		UserId:          in.UserId,
		PeerType:        in.PeerType,
		PeerId:          in.PeerId,
		TopMessage:      nil,
		ReadOutboxMaxId: nil,
		ReadInboxMaxId:  &types.Int32Value{Value: in.ReadInboxId},
		UnreadCount:     nil,
		UnreadMark:      false,
	})

	return mtproto.BoolTrue, nil
}
