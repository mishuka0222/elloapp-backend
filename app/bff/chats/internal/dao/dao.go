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
	kafka "github.com/teamgram/marmota/pkg/mq"
	"github.com/teamgram/marmota/pkg/net/rpcx"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/bff/chats/internal/config"
	msg_client "gitlab.com/merehead/elloapp/backend/elloapp_backend/app/messenger/msg/msg/client"
	sync_client "gitlab.com/merehead/elloapp/backend/elloapp_backend/app/messenger/sync/client"
	authsession_client "gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/authsession/client"
	chat_client "gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/chat/client"
	dialog_client "gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/dialog/client"
	message_client "gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/message/client"
	user_client "gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/user/client"
	idgen_client "gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/idgen/client"
	media_client "gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/media/client"
)

type Dao struct {
	user_client.UserClient
	ChatClient *chat_client.ChatClientHelper
	msg_client.MsgClient
	sync_client.SyncClient
	media_client.MediaClient
	dialog_client.DialogClient
	authsession_client.AuthsessionClient
	idgen_client.IDGenClient2
	message_client.MessageClient
}

func New(c config.Config) *Dao {
	return &Dao{
		UserClient:        user_client.NewUserClient(rpcx.GetCachedRpcClient(c.UserClient)),
		ChatClient:        chat_client.NewChatClientHelper(rpcx.GetCachedRpcClient(c.ChatClient)),
		MsgClient:         msg_client.NewMsgClient(rpcx.GetCachedRpcClient(c.MsgClient)),
		DialogClient:      dialog_client.NewDialogClient(rpcx.GetCachedRpcClient(c.DialogClient)),
		SyncClient:        sync_client.NewSyncMqClient(kafka.MustKafkaProducer(c.SyncClient)),
		MediaClient:       media_client.NewMediaClient(rpcx.GetCachedRpcClient(c.MediaClient)),
		AuthsessionClient: authsession_client.NewAuthsessionClient(rpcx.GetCachedRpcClient(c.AuthsessionClient)),
		IDGenClient2:      idgen_client.NewIDGenClient2(rpcx.GetCachedRpcClient(c.IdgenClient)),
		MessageClient:     message_client.NewMessageClient(rpcx.GetCachedRpcClient(c.MessageClient)),
	}
}
