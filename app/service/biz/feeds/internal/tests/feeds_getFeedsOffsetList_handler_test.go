package tests

import (
	"context"
	"github.com/teamgram/teamgram-server/app/service/biz/feeds/feeds"
	"github.com/zeromicro/go-zero/core/logx"
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
