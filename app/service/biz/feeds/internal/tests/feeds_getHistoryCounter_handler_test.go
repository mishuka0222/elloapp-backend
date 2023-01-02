package tests

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/feeds/feeds"
	"testing"
)

func TestFeedsGetHistoryCounter(t *testing.T) {
	c := NewRPCClient()
	/*
		empty request
	*/
	data, err := c.FeedsGetHistoryCounter(context.Background(), &feeds.GetHistoryCounterReq{
		UserId: 777062,
	})
	if err != nil {
		t.Fatal(err)
	}
	logx.Debugf("%+v", data.GetCount())
}
