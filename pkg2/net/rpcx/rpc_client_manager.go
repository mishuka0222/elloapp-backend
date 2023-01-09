package rpcx

import (
	"github.com/zeromicro/go-zero/core/logx"
	"io"
	"strings"

	"github.com/zeromicro/go-zero/core/syncx"
	"github.com/zeromicro/go-zero/zrpc"
)

var (
	clientManager = syncx.NewResourceManager()
)

type client struct {
	zrpc.Client
}

func (c *client) Close() error {
	return nil
}

func GetCachedRpcClient(c zrpc.RpcClientConf) zrpc.Client {
	var (
		val io.Closer
		err error
	)
	if c.Etcd.Key == "" && len(c.Endpoints) == 0 {
		panic(c)
	}
	logx.Infof("client: %v", c)
	if len(c.Endpoints) > 0 {
		val, err = clientManager.GetResource(strings.Join(c.Endpoints, "/"), func() (io.Closer, error) {
			cli := zrpc.MustNewClient(c)
			return &client{
				Client: cli,
			}, nil
		})
		if err != nil {
			panic(err)
		}
	} else {
		val, err = clientManager.GetResource(c.Etcd.Key, func() (io.Closer, error) {
			cli := zrpc.MustNewClient(c)
			return &client{
				Client: cli,
			}, nil
		})
		if err != nil {
			panic(err)
		}
	}

	return val.(zrpc.Client)
}
