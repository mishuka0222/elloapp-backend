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
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/messenger/msg/inbox/internal/config"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/messenger/msg/inbox/internal/server/mq"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/messenger/msg/inbox/internal/svc"
	kafka "gitlab.com/merehead/elloapp/backend/elloapp_backend/pkg2/mq"
)

type (
	Config = config.Config
)

func New(c Config) *kafka.ConsumerGroup {
	return mq.New(svc.NewServiceContext(c), c.InboxConsumer)
}
