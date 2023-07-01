package svc

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/dfs/internal/config"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/dfs/internal/dao"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/dfs/internal/ffmpegutil"
)

type ServiceContext struct {
	Config config.Config
	*dao.Dao
	*ffmpegutil.FFmpegUtil
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		Dao:        dao.New(c),
		FFmpegUtil: ffmpegutil.NewFFmpegUtil(),
	}
}
