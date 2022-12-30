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
	"math/rand"

	"github.com/teamgram/proto/mtproto"
	msgpb "gitlab.com/merehead/elloapp/backend/elloapp_backend/app/messenger/msg/msg/msg"
	chatpb "gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/chat/chat"
)

// MessagesDeleteChatUser
// messages.deleteChatUser#a2185cab flags:# revoke_history:flags.0?true chat_id:long user_id:InputUser = Updates;
func (c *ChatsCore) MessagesDeleteChatUser(in *mtproto.TLMessagesDeleteChatUser) (*mtproto.Updates, error) {
	deleteUser := mtproto.FromInputUser(c.MD.UserId, in.UserId)

	switch deleteUser.PeerType {
	case mtproto.PEER_USER:
	case mtproto.PEER_SELF:
	default:
		err := mtproto.ErrPeerIdInvalid
		c.Logger.Errorf("messages.deleteChatUser - invalid peer", err)
		return nil, err
	}

	chat, err := c.svcCtx.Dao.ChatClient.Client().ChatDeleteChatUser(c.ctx, &chatpb.TLChatDeleteChatUser{
		ChatId:       in.ChatId,
		OperatorId:   c.MD.UserId,
		DeleteUserId: deleteUser.PeerId,
	})
	if err != nil {
		c.Logger.Errorf("messages.deleteChatUser - error: %v", err)
		return nil, err
	}
	c.svcCtx.Dao.DialogClient.DialogDeleteDialog(c.ctx, &dialog.TLDialogDeleteDialog{
		UserId:   deleteUser.PeerId,
		PeerType: mtproto.PEER_CHAT,
		PeerId:   in.ChatId,
	})

	replyUpdates, err := c.svcCtx.Dao.MsgClient.MsgSendMessage(c.ctx, &msgpb.TLMsgSendMessage{
		UserId:    c.MD.UserId,
		AuthKeyId: c.MD.AuthId,
		PeerType:  mtproto.PEER_CHAT,
		PeerId:    in.ChatId,
		Message: msgpb.MakeTLOutboxMessage(&msgpb.OutboxMessage{
			NoWebpage:    true,
			Background:   false,
			RandomId:     rand.Int63(),
			Message:      chat.MakeMessageService(c.MD.UserId, mtproto.MakeMessageActionChatDeleteUser(deleteUser.PeerId)),
			ScheduleDate: nil,
		}).To_OutboxMessage(),
	})
	if err != nil {
		c.Logger.Errorf("messages.deleteChatUser - error: %v", err)
		return nil, err
	}

	updateChatParticipants := mtproto.MakeTLUpdateChatParticipants(&mtproto.Update{
		Participants_CHATPARTICIPANTS: chat.ToChatParticipants(0),
	}).To_Update()
	if deleteUser.PeerType == mtproto.PEER_USER {
		replyUpdates.Updates = append(replyUpdates.Updates, updateChatParticipants)
	}

	return replyUpdates, nil
}
