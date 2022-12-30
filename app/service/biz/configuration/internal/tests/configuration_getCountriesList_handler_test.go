package tests

import (
	"context"
	"github.com/teamgram/teamgram-server/app/service/biz/configuration/configuration"
	"github.com/zeromicro/go-zero/core/logx"
	"testing"
)

func TestConfigurationGetCountryList(t *testing.T) {
	c := NewRPCClient()

	data, err := c.ConfigurationGetCountryList(context.Background(), &configuration.GetCountryListReq{
		LangCode: "",
	})
	if err != nil {
		t.Fatal(err)
	}
	logx.Debugf("%+v", data.Data)
}
