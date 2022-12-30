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
	chatpb "gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/chat/chat"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
)

// MessagesDeleteExportedChatInvite
// messages.deleteExportedChatInvite#d464a42b peer:InputPeer link:string = Bool;
func (c *ChatInvitesCore) MessagesDeleteExportedChatInvite(in *mtproto.TLMessagesDeleteExportedChatInvite) (*mtproto.Bool, error) {
	var (
		peer = mtproto.FromInputPeer2(c.MD.UserId, in.Peer)
		err  error
	)

	if !peer.IsChat() {
		err = mtproto.ErrPeerIdInvalid
		c.Logger.Errorf("messages.deleteExportedChatInvite - error: ", err)
		return nil, err
	}

	_, err = c.svcCtx.Dao.ChatClient.ChatDeleteExportedChatInvite(c.ctx, &chatpb.TLChatDeleteExportedChatInvite{
		SelfId: c.MD.UserId,
		ChatId: peer.PeerId,
		Link:   in.GetLink(),
	})
	if err != nil {
		c.Logger.Errorf("messages.deleteExportedChatInvite - error: %v", err)
		return nil, err
	}

	return mtproto.BoolTrue, nil
}
