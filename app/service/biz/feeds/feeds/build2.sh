#!/bin/sh

SRC_DIR=.
DST_DIR=.

GOGOPROTO_PATH=$GOPATH/pkg/mod/github.com/gogo/protobuf@v1.3.2/protobuf
MTPROTO_PATH=$GOPATH/pkg/mod/github.com/teamgram/proto@v0.150.0/mtproto

protoc -I=$SRC_DIR:$MTPROTO_PATH --proto_path=$GOPATH/pkg/mod:$GOGOPROTO_PATH:./ \
    --gogo_out=plugins=grpc,Mgoogle/protobuf/wrappers.proto=github.com/gogo/protobuf/types,:$DST_DIR \
    $SRC_DIR/*.proto
#protoc -I=$SRC_DIR --proto_path=$GOPATH/src:$GOPATH/src/nebula.chat/vendor:$GOGOPROTO_PATH:./ \
#    --gogo_out=plugins=grpc,Mgoogle/protobuf/wrappers.proto=github.com/gogo/protobuf/types,:$DST_DIR \
#    $SRC_DIR/rpc_error_codes.proto
