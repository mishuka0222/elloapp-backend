// Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
//  All rights reserved.
//
// Author: teamgramio (teamgram.io@gmail.com)
//

package mq

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/messenger/sync/internal/core"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/messenger/sync/internal/svc"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/messenger/sync/sync"
	kafka "gitlab.com/merehead/elloapp/backend/elloapp_backend/pkg2/mq"

	"github.com/gogo/protobuf/proto"
	"github.com/zeromicro/go-zero/core/logx"
)

// New new a grpc server.
func New(svcCtx *svc.ServiceContext, conf kafka.KafkaConsumerConf) *kafka.ConsumerGroup {
	s := kafka.MustKafkaConsumer(&conf)
	s.RegisterHandlers(
		conf.Topics[0],
		func(ctx context.Context, key string, value []byte) {
			logx.WithContext(ctx).Debugf("key: %s, value: %s", key, value)

			switch strings.Split(key, "#")[0] {
			case proto.MessageName((*sync.TLSyncUpdatesMe)(nil)):
				c := core.New(ctx, svcCtx)

				r := new(sync.TLSyncUpdatesMe)
				if err := json.Unmarshal(value, r); err != nil {
					c.Logger.Error(err.Error())
					return
				}
				c.Logger.Debugf("sync.updatesMe - request: %s", r.DebugString())

				c.SyncUpdatesMe(r)
			case proto.MessageName((*sync.TLSyncUpdatesNotMe)(nil)):
				c := core.New(ctx, svcCtx)

				r := new(sync.TLSyncUpdatesNotMe)
				if err := json.Unmarshal(value, r); err != nil {
					c.Logger.Error(err.Error())
					return
				}
				c.Logger.Debugf("sync.updatesNotMe - request: %s", r.DebugString())

				c.SyncUpdatesNotMe(r)
			case proto.MessageName((*sync.TLSyncPushUpdates)(nil)):
				c := core.New(ctx, svcCtx)

				r := new(sync.TLSyncPushUpdates)
				if err := json.Unmarshal(value, r); err != nil {
					c.Logger.Error(err.Error())
					return
				}
				c.Logger.Debugf("sync.pushUpdates - request: %s", r.DebugString())

				c.SyncPushUpdates(r)
			case proto.MessageName((*sync.TLSyncPushRpcResult)(nil)):
				c := core.New(ctx, svcCtx)

				r := new(sync.TLSyncPushRpcResult)
				if err := json.Unmarshal(value, r); err != nil {
					c.Logger.Error(err.Error())
					return
				}
				c.Logger.Debugf("sync.pushRpcResult - request: %s", r.DebugString())

				c.SyncPushRpcResult(r)
			case proto.MessageName((*sync.TLSyncBroadcastUpdates)(nil)):
				c := core.New(ctx, svcCtx)

				r := new(sync.TLSyncBroadcastUpdates)
				if err := json.Unmarshal(value, r); err != nil {
					c.Logger.Error(err.Error())
					return
				}
				c.Logger.Debugf("sync.broadcastUpdates - request: %s", r.DebugString())

				c.SyncBroadcastUpdates(r)
			default:
				err := fmt.Errorf("invalid key: %s", key)
				logx.Error(err.Error())
			}
		})
	return s
}
