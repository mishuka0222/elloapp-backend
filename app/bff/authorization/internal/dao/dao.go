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
	"flag"
	"github.com/oschwald/geoip2-golang"
	"github.com/zeromicro/go-zero/core/stores/kv"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/bff/authorization/internal/config"
	msg_client "gitlab.com/merehead/elloapp/backend/elloapp_backend/app/messenger/msg/msg/client"
	sync_client "gitlab.com/merehead/elloapp/backend/elloapp_backend/app/messenger/sync/client"
	authsession_client "gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/authsession/client"
	chat_client "gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/chat/client"
	user_client "gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/user/client"
	username_client "gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/username/client"
	status_client "gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/status/client"
	kafka "gitlab.com/merehead/elloapp/backend/elloapp_backend/pkg2/mq"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/pkg2/net/rpcx"
)

var (
	mmdb string
)

func init() {
	flag.StringVar(&mmdb, "mmdb", "./GeoLite2-City.mmdb", "mmdb")
}

type Dao struct {
	kv   kv.Store
	MMDB *geoip2.Reader
	authsession_client.AuthsessionClient
	user_client.UserClient
	username_client.UsernameClient
	sync_client.SyncClient
	chat_client.ChatClient
	status_client.StatusClient
	msg_client.MsgClient
}

func New(c config.Config) *Dao {
	MMDB, err := geoip2.Open(mmdb)
	if err != nil {
		// panic(err)
	}
	return &Dao{
		kv:                kv.NewStore(c.KV),
		MMDB:              MMDB,
		UserClient:        user_client.NewUserClient(rpcx.GetCachedRpcClient(c.UserClient)),
		UsernameClient:    username_client.NewUsernameClient(rpcx.GetCachedRpcClient(c.UsernameClient)),
		AuthsessionClient: authsession_client.NewAuthsessionClient(rpcx.GetCachedRpcClient(c.AuthsessionClient)),
		ChatClient:        chat_client.NewChatClient(rpcx.GetCachedRpcClient(c.ChatClient)),
		SyncClient:        sync_client.NewSyncMqClient(kafka.MustKafkaProducer(c.SyncClient)),
		StatusClient:      status_client.NewStatusClient(rpcx.GetCachedRpcClient(c.StatusClient)),
		MsgClient:         msg_client.NewMsgClient(rpcx.GetCachedRpcClient(c.MsgClient)),
	}
}
