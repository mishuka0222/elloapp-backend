/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package inbox_helper

import (
	kafka "github.com/teamgram/marmota/pkg/mq"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/messenger/msg/inbox/internal/config"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/messenger/msg/inbox/internal/server/mq"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/messenger/msg/inbox/internal/svc"
)

type (
	Config = config.Config
)

func New(c Config) *kafka.ConsumerGroup {
	return mq.New(svc.NewServiceContext(c), c.InboxConsumer)
}
