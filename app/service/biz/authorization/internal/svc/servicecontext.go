package svc

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/authorization/internal/config"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/authorization/internal/dao"
	user_helper "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/user"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	*dao.Dao
	Gorm       *gorm.DB
	UserClient *user_helper.UserService
}

func NewServiceContext(c config.Config, userSerivce *user_helper.UserService) *ServiceContext {
	db, err := gorm.Open(mysql.Open(c.Mysql.DSN), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		panic(err)
	}

	return &ServiceContext{
		Config:     c,
		Dao:        dao.New(c),
		Gorm:       db,
		UserClient: userSerivce,
	}
}
