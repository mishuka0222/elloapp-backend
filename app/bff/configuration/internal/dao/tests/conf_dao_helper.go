package tests

import (
	"flag"
	"github.com/teamgram/teamgram-server/app/bff/configuration/internal/config"
	"github.com/zeromicro/go-zero/core/conf"
)

var configFile = flag.String("f", "../../../etc/configuration.yaml", "the config file")

func newConfDao() config.Config {
	var c config.Config
	conf.MustLoad(*configFile, &c)
	return c
}
