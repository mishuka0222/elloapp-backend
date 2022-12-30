// Copyright 2022 Teamgram Authors
//  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Author: teamgramio (teamgram.io@gmail.com)
//

package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/dialog/dialog"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
)

// MessagesMarkDialogUnread
// messages.markDialogUnread#c286d98f flags:# unread:flags.0?true peer:InputDialogPeer = Bool;
func (c *DialogsCore) MessagesMarkDialogUnread(in *mtproto.TLMessagesMarkDialogUnread) (*mtproto.Bool, error) {
	var (
		peer *mtproto.PeerUtil
	)

	if c.MD.IsBot {
		err := mtproto.ErrBotMethodInvalid
		c.Logger.Errorf("messages.markDialogUnread - error: %v", err)
		return nil, err
	}

	switch in.Peer.GetPredicateName() {
	case mtproto.Predicate_inputDialogPeer:
		peer = mtproto.FromInputPeer2(c.MD.UserId, in.Peer.Peer)
	case mtproto.Predicate_inputDialogPeerFolder:
		// error
		c.Logger.Errorf("client not send inputDialogPeerFolder: %v", in.Peer)
		return mtproto.BoolFalse, nil
	default:
		err := mtproto.ErrInputConstructorInvalid
		c.Logger.Errorf("messages.markDialogUnread - error: %v", err)
		return nil, err
	}

	rValue, err := c.svcCtx.Dao.DialogClient.DialogMarkDialogUnread(c.ctx, &dialog.TLDialogMarkDialogUnread{
		UserId:     c.MD.UserId,
		PeerType:   peer.PeerType,
		PeerId:     peer.PeerId,
		UnreadMark: mtproto.ToBool(in.Unread),
	})
	if err != nil {
		c.Logger.Errorf("messages.markDialogUnread - error: %v", err)
		return nil, err
	}

	// TODO: updates ...

	return rValue, nil
}
