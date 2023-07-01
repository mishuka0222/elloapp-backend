package tests

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/feeds/feeds"
	"testing"
)

func TestFeedsGetFeedsOffsetList(t *testing.T) {
	c := NewRPCClient()
	/*
		empty request
	*/
	var tLimit int32 = 35
	data, err := c.FeedsGetFeedsOffsetList(context.Background(), &feeds.GetFeedsOffsetListReq{
		UserId: 777062,
		Limit:  tLimit,
		Before: 0,
	})
	if err != nil {
		t.Fatal(err)
	}
	logx.Debugf("%+v", data.Data)
}
