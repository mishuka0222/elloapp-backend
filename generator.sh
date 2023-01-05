PWD=`pwd`
GOPATH=$HOME/go
GO_CTR_PATH=$PWD/generator/macapp

GOGOPROTO_PATH=$GOPATH/pkg/mod/github.com/gogo/protobuf@v1.3.2/protobuf
SCHEMA_PATH=$PWD/mtproto/example.proto

WD_PATH=app/service/biz/phone_call
WORKDIR=$PWD/$WD_PATH
SRC_PATH=phone_call.proto
DST_DIR=.
DST_DIR2=phone_call

cd $WORKDIR

$GO_CTR_PATH rpc protoc $SRC_PATH \
  -I=$PWD \
  -I=$SCHEMA_PATH \
  -I=$GOGOPROTO_PATH \
  --gogo_dst=$DST_DIR2 \
  --gogo_out="plugins=grpc\,Mgoogle/protobuf/wrappers.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/api.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/empty.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/descriptor.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/field_mask.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/source_context.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/struct.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/type.proto=github.com/gogo/protobuf/types,Mrpc_error_codes.proto=gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto,Mschema.tl.core_types.proto=gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto,Mschema.tl.crc32.proto=gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto,Mschema.tl.handshake.proto=gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto,Mschema.tl.handshake_service.proto=gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto,Mschema.tl.sync.proto=gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto,Mschema.tl.sync_service.proto=gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto,Mschema.tl.transport.proto=gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto,Mschema.tl.transport_service.proto=gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto,Mservice.tl.proto=gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"\
  --zrpc_out=$DST_DIR \
  --commands_pkg="gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/pkg2/commands" \
  --mtproto_pkg="gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto/rpc/metadata" \
#  --verbose


#import "rpc_error_codes.proto";
#import "schema.tl.core_types.proto";
#import "schema.tl.crc32.proto";
#import "schema.tl.handshake.proto";
#import "schema.tl.handshake_service.proto";
#import "schema.tl.sync.proto";
#import "schema.tl.sync_service.proto";
#import "schema.tl.transport.proto";
#import "schema.tl.transport_service.proto";
#import "service.tl.proto";
#import "my_type.proto";
