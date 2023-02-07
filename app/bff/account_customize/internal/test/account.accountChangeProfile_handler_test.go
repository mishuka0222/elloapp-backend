package tests

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/bff/account_customize/internal/core"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/bff/account_customize/internal/service"
	op_srv "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/bff/bizraw/service"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/account/account"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto/rpc/metadata"
	"testing"
	"time"
)

func TestAccountChangeProfile(t *testing.T) {
	c := NewRPCClient()
	ctx, err := metadata.RpcMetadataToOutgoing(context.Background(), &metadata.RpcMetadata{
		ServerId:      "127.0.0.1:20120",
		ClientAddr:    "92.38.127.109",
		AuthId:        1937090286237253747,
		SessionId:     -6067022183794192689,
		ReceiveTime:   time.Now().Unix(),
		ClientMsgId:   7186682945929270272,
		IsBot:         false,
		Layer:         147,
		UserId:        777109,
		Client:        "android",
		IsAdmin:       false,
		Langpack:      "android",
		PermAuthKeyId: 7681015971433756336,
	})
	if err != nil {
		panic(err)
	}
	/*
		empty request
	*/
	op, err := op_srv.NewOperation(op_srv.Operation{
		Service: op_srv.AccountCustomize,
		Method:  service.AccountChangeProfile,
		Data: core.ChangeProfile{
			FirstName:   "Makhmudov",
			LastName:    "Aziz",
			Username:    "makhmudov0",
			Bio:         "my new bio",
			Gender:      "Male",
			Birthday:    "1998-07-10T00:00:00+0000",
			CountryCode: "AFG",
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	data, err := c.BizInvokeBizDataRaw(ctx, op)
	if err != nil {
		fmt.Println("data", data)
		t.Fatal(err)
	}
	var resp account.ChangeAccountInfoResp
	if err := json.Unmarshal(data.Data, &resp); err != nil {
		t.Fatal(err)
	}
	logx.Debugf("%+v", resp)
}
