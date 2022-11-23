package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/teamgram/proto/mtproto"
	authorization_client "github.com/teamgram/teamgram-server/app/bff/authorization/client"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/zrpc"
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

	in := &mtproto.TLAuthSignUp{
		Constructor:   0,
		PhoneNumber:   "+86 111 1111 1111",
		PhoneCodeHash: "1",
		FirstName:     "Ramazan",
		LastName:      "Poyraz",
	}
	signIn, err := client.AuthSignUp(context.Background(), in)
	if err != nil {
		fmt.Println("ERROR 2 : ", err.Error())
		return
	}

	fmt.Println("RESULT: ", signIn)
}
