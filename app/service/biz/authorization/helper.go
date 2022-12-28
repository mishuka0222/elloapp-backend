/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package chat_helper

import (
	"github.com/teamgram/teamgram-server/app/service/biz/authorization/internal/config"
	"github.com/teamgram/teamgram-server/app/service/biz/authorization/internal/server/grpc/service"
	"github.com/teamgram/teamgram-server/app/service/biz/authorization/internal/svc"
)

type (
	Config = config.Config
)

func New(c Config) *service.Service {
	return service.New(svc.NewServiceContext(c))
}
