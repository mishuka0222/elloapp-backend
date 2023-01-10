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
	/*
		{"server_id":"127.0.0.1:20120",
		"client_addr":"192.168.1.5",
		"auth_id":7338124102345237054,
		"session_id":-6066913564088792595,
		"receive_time":1671767460,
		"user_id":777062,
		"client_msg_id":7180186570927840256,
		"layer":147,
		"client":"android",
		"langpack":"android",
		"perm_auth_key_id":9167378892795598833}
	*/
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
			Balance:   0,
			UserName:  "test",
			FirstName: "test",
			LastName:  "test",
			Password:  "test",
			Email:     "test@ff.fd",
			Type:      "Test",
			Profile:   "@test",
			Gender:    "male",
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
