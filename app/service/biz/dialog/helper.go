package dialog_helper

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/dialog/internal/config"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/dialog/internal/dal/dao/mysql_dao"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/dialog/internal/dal/dataobject"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/dialog/internal/server/grpc/service"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/dialog/internal/svc"
)

type (
	Config = config.Config
)

func New(c Config) *service.Service {
	return service.New(svc.NewServiceContext(c))
}

type (
	DialogsDAO = mysql_dao.DialogsDAO
	DialogsDO  = dataobject.DialogsDO
)

var (
	NewDialogsDAO = mysql_dao.NewDialogsDAO
)
