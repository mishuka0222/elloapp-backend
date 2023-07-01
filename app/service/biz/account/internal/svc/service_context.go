package svc

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/account/internal/config"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/account/internal/dao"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	Dao    *dao.Dao
	Gorm   *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, err := gorm.Open(mysql.Open(c.Mysql.DSN), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		panic(err)
	}

	return &ServiceContext{
		Config: c,
		Dao:    dao.New(c),
		Gorm:   db,
	}
}
