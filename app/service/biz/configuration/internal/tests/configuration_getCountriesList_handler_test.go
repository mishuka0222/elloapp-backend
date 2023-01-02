package tests

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/configuration/configuration"
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
