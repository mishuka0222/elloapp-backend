/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2022 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package autodownload_helper

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/bff/autodownload/internal/config"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/bff/autodownload/internal/server/grpc/service"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/bff/autodownload/internal/svc"
)

type (
	Config = config.Config
)

func New(c Config) *service.Service {
	return service.New(svc.NewServiceContext(c))
}
