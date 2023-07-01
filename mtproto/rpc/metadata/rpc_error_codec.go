package metadata

import (
	"encoding/base64"

	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"

	"github.com/gogo/protobuf/proto"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

var (
	headerRpcError = "rpc_error"
)

// RpcErrorFromMD Server To Client
func RpcErrorFromMD(md metadata.MD) (rpcErr *mtproto.TLRpcError) {
	val := NiceMD(md).Get(headerRpcError)
	if val == "" {
		// TODO(@benqi): 未设置rpc_error
		rpcErr = mtproto.NewRpcError(status.Convert(mtproto.ErrInternelServerError))
		// log.Errorf("%v", rpcErr)
		return
	}

	// proto.Marshal()
	buf, err := base64.StdEncoding.DecodeString(val)
	if err != nil {
		rpcErr = mtproto.NewRpcError(status.Convert(mtproto.ErrInternelServerError))
		// log.Errorf("%v", rpcErr)
		return
	}

	rpcErr = &mtproto.TLRpcError{}
	err = proto.Unmarshal(buf, rpcErr)
	if err != nil {
		rpcErr = mtproto.NewRpcError(status.Convert(mtproto.ErrInternelServerError))
		// log.Errorf("%v", rpcErr)
		return
	}

	// log.Errorf("%v", rpcErr)
	return rpcErr
}

func RpcErrorToMD(md *mtproto.TLRpcError) (metadata.MD, error) {
	buf, err := proto.Marshal(md)
	if err != nil {
		// log.Errorf("Marshal rpc_metadata error: %v", err)
		return nil, err
	}

	return metadata.Pairs(headerRpcError, base64.StdEncoding.EncodeToString(buf)), nil
}
