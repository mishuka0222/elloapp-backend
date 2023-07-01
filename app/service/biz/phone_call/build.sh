SRC_DIR=.
DST_DIR=.

GOPATH=$HOME/go
GOGOPROTO_PATH=$GOPATH/pkg/mod/github.com/gogo/protobuf@v1.3.2/protobuf

protoc --proto_path=$GOGOPROTO_PATH:./ \
    --gogo_out=plugins=grpc\,Mgoogle/protobuf/wrappers.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/api.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/empty.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/descriptor.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/field_mask.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/source_context.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/struct.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/type.proto=github.com/gogo/protobuf/types,:$DST_DIR \
        $SRC_DIR/*.proto
