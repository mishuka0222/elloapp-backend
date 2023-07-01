package inbox_helper

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/messenger/msg/inbox/internal/config"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/messenger/msg/inbox/internal/server/mq"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/messenger/msg/inbox/internal/svc"
	kafka "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/pkg2/mq"
)

type (
	Config = config.Config
)

func New(c Config) *kafka.ConsumerGroup {
	return mq.New(svc.NewServiceContext(c), c.InboxConsumer)
}
