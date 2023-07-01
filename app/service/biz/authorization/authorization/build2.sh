SRC_DIR=.
DST_DIR=.

GOPATH=$HOME/go
GOGOPROTO_PATH=$GOPATH/pkg/mod/github.com/gogo/protobuf@v1.3.2/protobuf
SCHEMA_PATH=/home/makhmudov/Desktop/Merehead/elloapp_tg_backend/mtproto/example.proto

protoc --proto_path=$GOGOPROTO_PATH:./ \
    --proto_path=$SCHEMA_PATH:./ \
    --gogo_out=plugins=grpc\,Mgoogle/protobuf/wrappers.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/api.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/empty.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/descriptor.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/field_mask.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/source_context.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/struct.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/type.proto=github.com/gogo/protobuf/types,Mrpc_error_codes.proto=gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto,Mschema.tl.core_types.proto=gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto,Mschema.tl.crc32.proto=gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto,Mschema.tl.handshake.proto=gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto,Mschema.tl.handshake_service.proto=gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto,Mschema.tl.sync.proto=gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto,Mschema.tl.sync_service.proto=gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto,Mschema.tl.transport.proto=gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto,Mschema.tl.transport_service.proto=gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto,Mservice.tl.proto=gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto,Mmy_type.proto=gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto,:$DST_DIR \
        $SRC_DIR/*.proto