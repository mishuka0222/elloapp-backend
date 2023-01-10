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
)

func TestConfirmation(t *testing.T) {
	c := NewRPCClient()
	ctx, err := metadata.RpcMetadataToOutgoing(context.Background(), &metadata.RpcMetadata{
		ServerId:      "127.0.0.1:20120",
		ClientAddr:    "192.168.1.105",
		AuthId:        7338124102345237054,
		SessionId:     -6066913564088792595,
		ReceiveTime:   1671767460,
		UserId:        777062,
		ClientMsgId:   7180186570927840256,
		IsBot:         false,
		Layer:         147,
		Client:        "android",
		IsAdmin:       false,
		Langpack:      "android",
		PermAuthKeyId: 9167378892795598833,
	})
	if err != nil {
		panic(err)
	}
	/*
		empty request
	*/
	op, err := op_srv.NewOperation(op_srv.Operation{
		Service: op_srv.AuthorizationCustomize,
		Method:  service.Confirmation,
		Data: core.ConfirmationReq{
			UsernameOrEmail: "nmz",
			Code:            "",
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	data, err := c.BizInvokeBizDataRaw(ctx, op)
	if err != nil {
		t.Fatal(err)
	}
	var resp core.ConfirmationResp
	if err := json.Unmarshal(data.Data, &resp); err != nil {
		t.Fatal(err)
	}
	logx.Debugf("%+v", resp)
}
