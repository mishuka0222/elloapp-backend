package tests

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/channels/internal/dao"
	"testing"
)

func TestSelectDAO(t *testing.T) {
	d := dao.New(newConfDao()).ChannelsDAO

	ch, err := d.Select(context.Background(), 1073741826)
	if err != nil {
		t.Fatal(err)
	}
	logx.Debug(ch)
}
func TestSelectByIdListDAO(t *testing.T) {
	d := dao.New(newConfDao()).ChannelsDAO

	chs, err := d.SelectByIdList(context.Background(), []int64{1073741826})
	if err != nil {
		t.Fatal(err)
	}
	logx.Debug(chs)
}
