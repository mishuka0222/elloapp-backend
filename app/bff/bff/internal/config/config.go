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

package config

import (
	"github.com/zeromicro/go-zero/core/stores/kv"
	"github.com/zeromicro/go-zero/zrpc"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/pkg/code/conf"
	kafka "gitlab.com/merehead/elloapp/backend/elloapp_backend/pkg2/mq"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/pkg2/stores/sqlx"
)

type Config struct {
	zrpc.RpcServerConf
	KV                kv.KvConf
	Code              *conf.SmsVerifyCodeConfig
	BizServiceClient  zrpc.RpcClientConf
	AuthSessionClient zrpc.RpcClientConf
	MediaClient       zrpc.RpcClientConf
	IdgenClient       zrpc.RpcClientConf
	MsgClient         zrpc.RpcClientConf
	SyncClient        *kafka.KafkaProducerConf
	DfsClient         zrpc.RpcClientConf
	StatusClient      zrpc.RpcClientConf
	Mysql             sqlx.Config
}
