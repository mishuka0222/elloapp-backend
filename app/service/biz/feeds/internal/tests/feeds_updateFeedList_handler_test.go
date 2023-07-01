package tests

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/feeds/feeds"
	"testing"
)

func TestUpdateFeedList(t *testing.T) {
	c := NewRPCClient()
	/*
		empty request
	*/

	data, err := c.FeedsUpdateFeedList(context.Background(), &feeds.UpdateFeedListReq{
		UserId: 777062,
		Data: []*feeds.FeedInsertItemDO{
			{ChatId: 14, PeerType: 2},
			{ChatId: 15, PeerType: 2},
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	logx.Debugf("%+v", data)
}
