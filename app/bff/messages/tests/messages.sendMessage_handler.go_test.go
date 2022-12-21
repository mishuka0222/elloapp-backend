package tests

import (
	"context"
	"fmt"
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/proto/mtproto/rpc/metadata"
	"testing"
)

func TestMessagesSendMessage(t *testing.T) {
	c := NewRPCClient()
	/*
		{"server_id":"127.0.0.1:20120",
		"client_addr":"192.168.1.5",
		"auth_id":-6993468262964525027,
		"session_id":-6066989938548526142,
		"receive_time":1671460067,
		"user_id":777062,
		"client_msg_id":7178866321461636096,
		"layer":147,
		"client":"android",
		"langpack":"android",
		"perm_auth_key_id":4596884163960997514}
	*/
	ctx, err := metadata.RpcMetadataToOutgoing(context.Background(), &metadata.RpcMetadata{
		ServerId:      "127.0.0.1:20120",
		ClientAddr:    "192.168.1.5",
		AuthId:        -6993468262964525027,
		SessionId:     -6066989938548526142,
		ReceiveTime:   1671460067,
		UserId:        777062,
		ClientMsgId:   7178866321461636096,
		IsBot:         false,
		Layer:         147,
		Client:        "android",
		IsAdmin:       false,
		Langpack:      "android",
		PermAuthKeyId: 4596884163960997514,
	})
	if err != nil {
		panic(err)
	}
	/*
		{"constructor":"CRC32_messages_sendMessage",
		"clear_draft":true,
		"peer":
			{"predicate_name":"inputPeerChat",
			"constructor":"CRC32_inputPeerChat",
			"chat_id":"14"},
		"message":"Ret",
		"random_id":"608675164444996151"}
	*/
	updates, err := c.MessagesSendMessage(ctx, &mtproto.TLMessagesSendMessage{
		Constructor: mtproto.CRC32_messages_sendMessage,
		ClearDraft:  true,
		Peer: (&mtproto.InputPeer{
			Constructor: mtproto.CRC32_inputPeerChat,
			ChannelId:   14,
		}).To_InputPeerChat().To_InputPeer(),
		Message:  "My console msg",
		RandomId: 608675164444996153,
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", updates)
}
