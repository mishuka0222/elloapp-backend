package core

import (
	"github.com/gogo/protobuf/proto"
	"github.com/zeromicro/go-zero/core/jsonx"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
	"os"
	"time"
)

const (

	// date = 1509066502,    2017/10/27 09:08:22
	// expires = 1509070295, 2017/10/27 10:11:35
	expiresTimeout = 3600 // 超时时间设置为3600秒

	// support user: @benqi
	// SUPPORT_USER_ID = 2
)

var configFile = "./config.json"

var config mtproto.TLConfig

func init() {
	if path := os.Getenv("ELLO_APP_BFF_CONFIGURATION_CONFIG_JSON_PATH"); path != "" {
		configFile = path
	}
	configData, err := os.ReadFile(configFile)
	if err != nil {
		panic(err)
		return
	}

	err = jsonx.Unmarshal(configData, &config)
	if err != nil {
		panic(err)
		return
	}
}

// HelpGetConfig
// help.getConfig#c4f9186b = Config;
func (c *ConfigurationCore) HelpGetConfig(in *mtproto.TLHelpGetConfig) (*mtproto.Config, error) {
	_ = in

	rValue, _ := proto.Clone(&config).(*mtproto.TLConfig)
	now := int32(time.Now().Unix())
	rValue.SetDate(now)
	rValue.SetExpires(now + expiresTimeout)

	return rValue.To_Config(), nil
}
