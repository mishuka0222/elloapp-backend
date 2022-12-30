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

package dao

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/bff/messages/internal/config"
	msg_client "gitlab.com/merehead/elloapp/backend/elloapp_backend/app/messenger/msg/msg/client"
	sync_client "gitlab.com/merehead/elloapp/backend/elloapp_backend/app/messenger/sync/client"
	chat_client "gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/chat/client"
	dialog_client "gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/dialog/client"
	message_client "gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/message/client"
	user_client "gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/user/client"
	username_client "gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/username/client"
	idgen_client "gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/idgen/client"
	media_client "gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/media/client"
	kafka "gitlab.com/merehead/elloapp/backend/elloapp_backend/pkg2/mq"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/pkg2/net/rpcx"
)

type Dao struct {
	msg_client.MsgClient
	user_client.UserClient
	ChatClient *chat_client.ChatClientHelper
	media_client.MediaClient
	username_client.UsernameClient
	message_client.MessageClient
	idgen_client.IDGenClient2
	dialog_client.DialogClient
	sync_client.SyncClient
}

func New(c config.Config) *Dao {
	return &Dao{
		MsgClient:      msg_client.NewMsgClient(rpcx.GetCachedRpcClient(c.MsgClient)),
		UserClient:     user_client.NewUserClient(rpcx.GetCachedRpcClient(c.UserClient)),
		ChatClient:     chat_client.NewChatClientHelper(rpcx.GetCachedRpcClient(c.ChatClient)),
		MediaClient:    media_client.NewMediaClient(rpcx.GetCachedRpcClient(c.MediaClient)),
		DialogClient:   dialog_client.NewDialogClient(rpcx.GetCachedRpcClient(c.DialogClient)),
		IDGenClient2:   idgen_client.NewIDGenClient2(rpcx.GetCachedRpcClient(c.IdgenClient)),
		MessageClient:  message_client.NewMessageClient(rpcx.GetCachedRpcClient(c.MessageClient)),
		UsernameClient: username_client.NewUsernameClient(rpcx.GetCachedRpcClient(c.UsernameClient)),
		SyncClient:     sync_client.NewSyncMqClient(kafka.MustKafkaProducer(c.SyncClient)),
	}
}
