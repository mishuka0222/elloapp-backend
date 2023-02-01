package tests

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/channels/channels"
	"testing"
)

func TestCheckChannelUsername(t *testing.T) {
	c := NewRPCClient()

	/*
		empty request
	*/

	data, err := c.CheckChannelUserName(context.Background(), &channels.CheckChannelUserNameReq{
		ChannelId: 1073741830,
		UserName:  "myLinkErr",
	})
	if err != nil {
		t.Fatal(err)
	}
	logx.Debugf("%+v", data.Status)
}
