package tests

import (
	"context"
	"fmt"
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/proto/mtproto/rpc/metadata"
	"testing"
)

func TestMessagesGetPeerDialogs(t *testing.T) {
	c := NewRPCClient()
	/*
		{"server_id":"127.0.0.1:20120",
		"client_addr":"192.168.1.2",
		"auth_id":8496467996535307930,
		"session_id":-6067082837832915424,
		"receive_time":1671422104,
		"user_id":777063,
		"client_msg_id":7178703274379841536,
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
		ClientMsgId:   7178703274379841536,
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
		{"constructor":"CRC32_messages_getPeerDialogs",
		"peers":[
			{"predicate_name":"inputDialogPeer",
			"constructor":"CRC32_inputDialogPeer",
			"peer":
				{"predicate_name":"inputPeerUser",
				"constructor":"CRC32_inputPeerUser",
				"user_id":"777063",
				"access_hash":"1649049434428990075"}
			}
		]}
	*/
	dialogs, err := c.MessagesGetPeerDialogs(ctx, &mtproto.TLMessagesGetPeerDialogs{
		Constructor: mtproto.CRC32_messages_getDialogs,
		Peers: []*mtproto.InputDialogPeer{
			(&mtproto.InputDialogPeer{
				Constructor: mtproto.CRC32_inputPeerUser,
				Peer: (&mtproto.InputPeer{
					PredicateName: "",
					Constructor:   mtproto.CRC32_inputPeerUser,
					UserId:        777063,
					AccessHash:    1649049434428990075,
				}).To_InputPeerUser().To_InputPeer(),
			}).To_InputDialogPeer().To_InputDialogPeer(),
		},
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", dialogs)
}
