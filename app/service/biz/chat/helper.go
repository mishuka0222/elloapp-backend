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
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/chat/internal/config"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/chat/internal/dal/dao/mysql_dao"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/chat/internal/dal/dataobject"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/chat/internal/server/grpc/service"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/chat/internal/svc"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/chat/plugin"
)

type (
	Config = config.Config
)

type (
	ChatParticipantsDAO = mysql_dao.ChatParticipantsDAO
	ChatParticipantsDO  = dataobject.ChatParticipantsDO
)

var (
	NewChatParticipantsDAO = mysql_dao.NewChatParticipantsDAO
)

func New(c Config, plugin plugin.ChatPlugin) *service.Service {
	return service.New(svc.NewServiceContext(c, plugin))
}
