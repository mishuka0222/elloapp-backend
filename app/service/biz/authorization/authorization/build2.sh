#!/bin/sh

SRC_DIR=.
DST_DIR=.

GOPATH=$HOME/go

GOGOPROTO_PATH=$GOPATH/pkg/mod/github.com/gogo/protobuf@v1.3.2/protobuf
#MTPROTO_PATH=/Volumes/EXTERNAL_DISK/GolandProjects/elloapp_tg_backend/mtproto
MTPROTO_PATH=/home/arthur/MereHead/Projects/elloapp_tg_backend/mtproto

protoc -I=$SRC_DIR:$MTPROTO_PATH --proto_path=$GOPATH/pkg/mod:$GOGOPROTO_PATH:./ \
    --gogo_out=plugins=grpc,Mgoogle/protobuf/wrappers.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/empty.proto=github.com/gogo/protobuf/types,:$DST_DIR \
    $SRC_DIR/*.proto

