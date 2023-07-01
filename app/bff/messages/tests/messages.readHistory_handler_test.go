package tests

import (
	"context"
	"fmt"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto/rpc/metadata"
	"testing"
)

func TestMessagesReadHistoryUser(t *testing.T) {
	c := NewRPCClient()
	/*
		{"server_id":"127.0.0.1:20120",
		"client_addr":"192.168.1.102",
		"auth_id":-2045420147104573983,
		"session_id":-6066981631837197752,
		"receive_time":1671531921,
		"user_id":777062,
		"client_msg_id":7179174937126964224,
		"layer":147,
		"client":"android",
		"langpack":"android",
		"perm_auth_key_id":7709644538661168808}
	*/
	ctx, err := metadata.RpcMetadataToOutgoing(context.Background(), &metadata.RpcMetadata{
		ServerId:      "127.0.0.1:20120",
		ClientAddr:    "192.168.1.102",
		AuthId:        -2045420147104573983,
		SessionId:     -6066981631837197752,
		ReceiveTime:   1671531921,
		UserId:        777062,
		ClientMsgId:   7179174937126964224,
		IsBot:         false,
		Layer:         147,
		Client:        "android",
		IsAdmin:       false,
		Langpack:      "android",
		PermAuthKeyId: 7709644538661168808,
	})
	if err != nil {
		panic(err)
	}
	/*
		{"constructor":"CRC32_messages_readHistory",
		"peer":
			{"predicate_name":"inputPeerUser",
			"constructor":"CRC32_inputPeerUser",
			"user_id":"777063",
			"access_hash":"1649049434428990075"},
		"max_id":338}
	*/
	dialogs, err := c.MessagesReadHistory(ctx, &mtproto.TLMessagesReadHistory{
		Constructor: mtproto.CRC32_messages_readHistory,
		Peer: (&mtproto.InputPeer{
			Constructor: mtproto.CRC32_inputPeerUser,
			UserId:      777063,
			AccessHash:  1649049434428990075,
		}).To_InputPeerUser().To_InputPeer(),
		MaxId: 388,
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", dialogs)
}

func TestMessagesReadHistoryGroup(t *testing.T) {
	c := NewRPCClient()
	/*
		{"server_id":"127.0.0.1:20120",
		"client_addr":"192.168.1.102",
		"auth_id":358524631330184663,
		"session_id":-6066926370668461958,
		"receive_time":1671531615,
		"user_id":777063,
		"client_msg_id":7179173624816886784,
		"layer":147,
		"client":"android",
		"langpack":"android",
		"perm_auth_key_id":327701596064343245}
	*/
	ctx, err := metadata.RpcMetadataToOutgoing(context.Background(), &metadata.RpcMetadata{
		ServerId:      "127.0.0.1:20120",
		ClientAddr:    "192.168.1.102",
		AuthId:        358524631330184663,
		SessionId:     -6066926370668461958,
		ReceiveTime:   1671531615,
		UserId:        777063,
		ClientMsgId:   7179173624816886784,
		IsBot:         false,
		Layer:         147,
		Client:        "android",
		IsAdmin:       false,
		Langpack:      "android",
		PermAuthKeyId: 327701596064343245,
	})
	if err != nil {
		panic(err)
	}
	/*
		{"constructor":"CRC32_messages_readHistory",
		"peer":
			{"predicate_name":"inputPeerChat",
			"constructor":"CRC32_inputPeerChat",
			"chat_id":"14"},
		"max_id":300}
	*/
	history, err := c.MessagesReadHistory(ctx, &mtproto.TLMessagesReadHistory{
		Constructor: mtproto.CRC32_messages_readHistory,
		Peer: (&mtproto.InputPeer{
			Constructor: mtproto.CRC32_inputPeerChat,
			ChannelId:   14,
		}).To_InputPeerUser().To_InputPeer(),
		MaxId: 300,
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", history)
}
