#!/usr/bin/env bash

PWD=`pwd`
ELLOAPP=${PWD}"/app"
INSTALL=${PWD}"/elloappd"

echo "build idgen ..."
cd ${ELLOAPP}/service/idgen/cmd/idgen
go build -o ${INSTALL}/bin/idgen

echo "build status ..."
cd ${ELLOAPP}/service/status/cmd/status
go build -o ${INSTALL}/bin/status

echo "build dfs ..."
cd ${ELLOAPP}/service/dfs/cmd/dfs
go build -o ${INSTALL}/bin/dfs

echo "build media ..."
cd ${ELLOAPP}/service/media/cmd/media
go build -o ${INSTALL}/bin/media

echo "build authsession ..."
cd ${ELLOAPP}/service/authsession/cmd/authsession
go build -o ${INSTALL}/bin/authsession

echo "build biz ..."
cd ${ELLOAPP}/service/biz/biz/cmd/biz
go build -o ${INSTALL}/bin/biz

echo "build msg ..."
cd ${ELLOAPP}/messenger/msg/cmd/msg
go build -o ${INSTALL}/bin/msg

echo "build sync ..."
cd ${ELLOAPP}/messenger/sync/cmd/sync
go build -o ${INSTALL}/bin/sync

echo "build bff ..."
cd ${ELLOAPP}/bff/bff/cmd/bff
go build -o ${INSTALL}/bin/bff

echo "build session ..."
cd ${ELLOAPP}/interface/session/cmd/session
go build -o ${INSTALL}/bin/session

echo "build gateway ..."
cd ${ELLOAPP}/interface/gateway/cmd/gateway
go build -o ${INSTALL}/bin/gateway
