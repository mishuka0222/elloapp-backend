package svc

import (
	"github.com/teamgram/teamgram-server/app/service/biz/configuration/internal/config"
	"github.com/teamgram/teamgram-server/app/service/biz/configuration/internal/dao"
)

type ServiceContext struct {
	Config config.Config
	*dao.Dao
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Dao:    dao.New(c),
	}
}
