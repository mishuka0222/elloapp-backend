/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2022 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package authorization_helper

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/bff/authorization/internal/config"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/bff/authorization/internal/server/grpc/service"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/bff/authorization/internal/svc"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/bff/authorization/plugin"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/pkg/code"
)

type (
	Config               = config.Config
	AuthorizationService = service.Service
)

func New(c Config, code2 code.VerifyCodeInterface, plugin plugin.AuthorizationPlugin) *service.Service {
	return service.New(svc.NewServiceContext(c, code2, plugin))
}
