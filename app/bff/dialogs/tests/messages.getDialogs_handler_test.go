package tests

import (
	"context"
	"fmt"
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/proto/mtproto/rpc/metadata"
	"testing"
)

func TestMessagesGetDialogs(t *testing.T) {
	c := NewRPCClient()
	/*
		{"server_id":"127.0.0.1:20120",
		"client_addr":"192.168.1.2",
		"auth_id":8496467996535307930,
		"session_id":-6067082837832915424,
		"receive_time":1671422104,
		"user_id":777063,
		"client_msg_id":7178703274366956544,
		"layer":147,
		"client":"android",
		"langpack":"android",
		"perm_auth_key_id":-4551200845352882821}
	*/
	ctx, err := metadata.RpcMetadataToOutgoing(context.Background(), &metadata.RpcMetadata{
		ServerId:      "127.0.0.1:20120",
		ClientAddr:    "192.168.1.2",
		AuthId:        8496467996535307930,
		SessionId:     -6067082837832915424,
		ReceiveTime:   1671422104,
		UserId:        777063,
		ClientMsgId:   7178703274366956544,
		IsBot:         false,
		Layer:         147,
		Client:        "android",
		IsAdmin:       false,
		Langpack:      "android",
		PermAuthKeyId: -4551200845352882821,
	})
	if err != nil {
		panic(err)
	}
	/*
		{"constructor":"CRC32_messages_getDialogs",
		"exclude_pinned":true,
		"offset_peer":{
			"predicate_name":"inputPeerEmpty",
			"constructor":"CRC32_inputPeerEmpty"},
		"limit":100}
	*/
	dialogs, err := c.MessagesGetDialogs(ctx, &mtproto.TLMessagesGetDialogs{
		Constructor:   mtproto.CRC32_messages_getDialogs,
		ExcludePinned: true,
		OffsetPeer: (&mtproto.InputPeer{
			Constructor: mtproto.CRC32_inputPeerEmpty,
		}).To_InputPeerEmpty().To_InputPeer(),
		Limit: 100,
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", dialogs)
}
