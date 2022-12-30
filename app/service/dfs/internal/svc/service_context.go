/*
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package svc

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/dfs/internal/config"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/dfs/internal/dao"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/dfs/internal/ffmpegutil"
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
