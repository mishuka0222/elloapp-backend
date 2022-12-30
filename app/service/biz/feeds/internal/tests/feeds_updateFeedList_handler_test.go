package tests

import (
	"context"
	"github.com/teamgram/teamgram-server/app/service/biz/feeds/feeds"
	"github.com/zeromicro/go-zero/core/logx"
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
