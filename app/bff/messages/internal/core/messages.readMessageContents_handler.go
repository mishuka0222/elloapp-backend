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
	msgpb "gitlab.com/merehead/elloapp/backend/elloapp_backend/app/messenger/msg/msg/msg"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/message/message"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
)

// MessagesReadMessageContents
// messages.readMessageContents#36a73f77 id:Vector<int> = messages.AffectedMessages;
func (c *MessagesCore) MessagesReadMessageContents(in *mtproto.TLMessagesReadMessageContents) (*mtproto.Messages_AffectedMessages, error) {
	messages, err := c.svcCtx.Dao.MessageClient.MessageGetUserMessageList(c.ctx, &message.TLMessageGetUserMessageList{
		UserId: c.MD.UserId,
		IdList: in.Id,
	})
	if err != nil {
		c.Logger.Errorf("messages.readMessageContents - error: %v", err)
		return nil, err
	} else if messages.Length() == 0 {
		c.Logger.Errorf("messages.readMessageContents - error: missing messages")
		return mtproto.MakeTLMessagesAffectedMessages(&mtproto.Messages_AffectedMessages{
			Pts:      c.svcCtx.Dao.IDGenClient2.CurrentPtsId(c.ctx, c.MD.UserId),
			PtsCount: 0,
		}).To_Messages_AffectedMessages(), nil
	}

	// TODO(@benqi): check peer??
	var (
		peer *mtproto.PeerUtil
	)
	if messages.Datas[0].PeerType == mtproto.PEER_CHAT {
		peer = mtproto.MakeChatPeerUtil(messages.Datas[0].Message.GetPeerId().GetChatId())
	} else {
		peer = mtproto.MakeUserPeerUtil(messages.Datas[0].SenderUserId)
	}
	contents := make([]*msgpb.ContentMessage, 0, len(messages.GetDatas()))
	for _, m := range messages.GetDatas() {
		// TODO(@benqi): check peer??
		// peer := model.FromPeer(m.Message.ToId)
		if m.Message.GetMentioned() {
			contents = append(contents, &msgpb.ContentMessage{
				Id:              m.MessageId,
				DialogMessageId: m.DialogMessageId,
				Mentioned:       true,
			})
		} else if m.Message.GetMediaUnread() {
			contents = append(contents, &msgpb.ContentMessage{
				Id:              m.MessageId,
				DialogMessageId: m.DialogMessageId,
				MediaUnread:     true,
			})
		} else if m.GetMessage().GetReactions() != nil {
			contents = append(contents, &msgpb.ContentMessage{
				Id:              m.MessageId,
				DialogMessageId: m.DialogMessageId,
				Reaction:        true,
			})
		} else {
			c.Logger.Infof("content has readed")
		}
	}

	affected, err := c.svcCtx.Dao.MsgClient.MsgReadMessageContents(c.ctx, &msgpb.TLMsgReadMessageContents{
		UserId:    c.MD.UserId,
		AuthKeyId: c.MD.AuthId,
		PeerType:  peer.PeerType,
		PeerId:    peer.PeerId,
		Id:        contents,
	})
	if err != nil {
		c.Logger.Errorf("messages.readMessageContents - %v", err)
		return nil, err
	}

	return affected, nil
}
