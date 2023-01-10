package tests

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/bff/authorization_customize/internal/core"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/bff/authorization_customize/internal/service"
	op_srv "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/bff/bizraw/service"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto/rpc/metadata"
	"testing"
	"time"
)

func TestAuthSingUp(t *testing.T) {
	c := NewRPCClient()
	ctx, err := metadata.RpcMetadataToOutgoing(context.Background(), &metadata.RpcMetadata{
		ServerId:      "127.0.0.1:20120",
		ClientAddr:    "92.38.127.109",
		AuthId:        -792622960526419741,
		SessionId:     -6067082085191987035,
		ReceiveTime:   time.Now().Unix(),
		UserId:        777099,
		ClientMsgId:   7186682945929270272,
		IsBot:         false,
		Layer:         147,
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
		Service: op_srv.AuthorizationCustomize,
		Method:  service.AuthSingUP,
		Data: core.AuthSingUPReq{
			Username:    "nmz",
			Password:    "password123P",
			Gender:      "Male",
			DateOfBirth: "1988-03-14T00:00:00+0000",
			Email:       "lalaland@gmail.com",
			Phone:       "+12345678",
			CountryCode: "UZB",
			Avatar:      "",
			FirstName:   "firstname",
			LastName:    "lastname",
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	data, err := c.BizInvokeBizDataRaw(ctx, op)
	if err != nil {
		t.Fatal(err)
	}
	var resp core.AuthSingUPResp
	if err := json.Unmarshal(data.Data, &resp); err != nil {
		t.Fatal(err)
	}
	logx.Debugf("%+v", resp)
}
