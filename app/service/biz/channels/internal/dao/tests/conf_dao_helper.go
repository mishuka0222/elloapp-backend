package tests

import (
	"flag"
	"github.com/zeromicro/go-zero/core/conf"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/channels/internal/config"
)

var configFile = flag.String("f", "../../../etc/channels.yaml", "the config file")

func newConfDao() config.Config {
	var c config.Config
	conf.MustLoad(*configFile, &c)
	return c
}
