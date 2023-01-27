package tests

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/channels/internal/dao"
	"testing"
)

func TestParticipantsSelectDAO(t *testing.T) {
	d := dao.New(newConfDao()).ChannelParticipantsDAO

	ch, err := d.SelectByChannelId(context.Background(), 1073741837)
	if err != nil {
		t.Fatal(err)
	}
	logx.Debug(ch)
}
