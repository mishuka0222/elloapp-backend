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

package svc

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/bff/authorization/internal/config"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/bff/authorization/internal/dao"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/bff/authorization/internal/logic"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/bff/authorization/plugin"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/pkg/code"
)

type ServiceContext struct {
	Config config.Config
	*dao.Dao
	*logic.AuthLogic
	Plugin plugin.AuthorizationPlugin
}

func NewServiceContext(c config.Config, code2 code.VerifyCodeInterface, plugin plugin.AuthorizationPlugin) *ServiceContext {
	d := dao.New(c)
	if code2 == nil {
		code2 = code.NewVerifyCode(c.Code)
	}
	return &ServiceContext{
		Config:    c,
		Dao:       d,
		AuthLogic: logic.NewAuthSignLogic(d, code2),
		Plugin:    plugin,
	}
}
