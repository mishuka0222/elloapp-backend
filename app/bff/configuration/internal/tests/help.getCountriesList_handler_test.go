package tests

import (
	"context"
	"github.com/teamgram/proto/mtproto"
	"github.com/zeromicro/go-zero/core/logx"
	"testing"
)

func TestGetFeedList(t *testing.T) {
	c := NewRPCClient()

	/*
		empty request
	*/
	data, err := c.HelpGetCountriesList(context.Background(), &mtproto.TLHelpGetCountriesList{
		Constructor: mtproto.CRC32_help_getCountriesList,
		LangCode:    "",
		Hash:        0,
	})
	if err != nil {
		t.Fatal(err)
	}
	logx.Debugf("%+v", data.GetCountries())
}
