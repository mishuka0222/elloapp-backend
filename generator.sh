PWD=`pwd`
GOPATH=$HOME/go
GO_CTR_PATH=$PWD/generator/macapp

GOGOPROTO_PATH=$GOPATH/pkg/mod/github.com/gogo/protobuf@v1.3.2/protobuf

WORKDIR=$PWD/app/service/biz/authorization2
SRC_PATH=authorization.proto
DST_DIR=.

cd $WORKDIR

$GO_CTR_PATH rpc protoc $SRC_PATH \
  -I=$PWD \
  -I=$GOGOPROTO_PATH \
  --gogo_dst=authorization \
  --gogo_out="plugins=grpc\,Mgoogle/protobuf/wrappers.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/api.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/empty.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/descriptor.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/field_mask.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/source_context.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/struct.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/type.proto=github.com/gogo/protobuf/types"\
  --zrpc_out=$DST_DIR \
  --commands_pkg="gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/pkg2/commands" \
  --mtproto_pkg="gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto/rpc/metadata" \
  --verbose
