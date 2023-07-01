package tests

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/feeds/feeds"
	"testing"
)

func TestGetFeedList(t *testing.T) {
	c := NewRPCClient()

	/*
		empty request
	*/

	data, err := c.FeedsGetFeedList(context.Background(), &feeds.GetFeedListReq{
		UserId: 777062,
	})
	if err != nil {
		t.Fatal(err)
	}
	logx.Debugf("%+v", data.Data)
}
