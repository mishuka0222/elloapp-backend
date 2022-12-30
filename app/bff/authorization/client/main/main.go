package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/proto/mtproto/rpc/metadata"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/zrpc"
	authorization_client "gitlab.com/merehead/elloapp/backend/elloapp_backend/app/bff/authorization/client"
)

var (
	configFile = flag.String("f", "../../etc/authorization.yaml", "the config file")
)

func main() {
	var c zrpc.RpcClientConf
	conf.MustLoad(*configFile, &c)

	cli, err := zrpc.NewClient(c)
	if err != nil {
		fmt.Println("ERROR 1 : ", err.Error())
		return
	}
	client := authorization_client.NewAuthorizationClient(cli)

	ctx := context.Background()
	md := metadata.RpcMetadata{
		AuthId: 3989972485446262080,
	}
	ctx2, err := metadata.RpcMetadataToOutgoing(ctx, &md)
	if err != nil {
		fmt.Println("ERROR METADATA CONTEXT: ", err.Error())
		return
	}

	in := mtproto.TLAuthSignUp{
		Constructor:   1,
		PhoneNumber:   "ramazan",
		PhoneCodeHash: "12345",
		FirstName:     "Ramazan",
		LastName:      "Poyraz",
	}
	signIn, err := client.AuthSignUp(ctx2, &in)
	if err != nil {
		fmt.Println("ERROR 2 : ", err.Error())
		return
	}

	fmt.Println("RESULT: ", signIn)

	//data := mtproto.TLAuthSignIn{
	//	Constructor:          1,
	//	PhoneNumber:          "ramazan",
	//	PhoneCodeHash:        "12345",
	//}
	//signININ, err := client.AuthSignIn(ctx2, &data)
	//if err != nil {
	//	fmt.Println("ERROR 3 : ", err.Error())
	//	return
	//}
	//
	//fmt.Println("RESULT2: ", signININ)
}
