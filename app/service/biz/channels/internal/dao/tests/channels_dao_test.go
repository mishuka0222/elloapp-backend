package tests

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/channels/internal/dao"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/pkg2/stores/sqlx"
	"testing"
)

func TestChannelsSelectDAO(t *testing.T) {
	d := dao.New(newConfDao()).ChannelsDAO

	ch, err := d.Select(context.Background(), 1073741826)
	if err != nil {
		t.Fatal(err)
	}
	logx.Debug(ch)
}

func TestChannelsSelectByIdListDAO(t *testing.T) {
	d := dao.New(newConfDao()).ChannelsDAO

	chs, err := d.SelectByIdList(context.Background(), []int64{1073741826})
	if err != nil {
		t.Fatal(err)
	}
	logx.Debug(chs)
}

func TestChannelsCheckUsernameDAO(t *testing.T) {
	d := dao.New(newConfDao()).ChannelsDAO

	chs, err := d.CheckUsername(context.Background(), "own2")
	if err == sqlx.ErrNotFound {
		return
	}
	if err != nil {
		t.Fatal(err)
	}
	logx.Debug(chs)
}
