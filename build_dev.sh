#!/usr/bin/env bash

PWD=`pwd`
ELLOAPP=${PWD}"/app"
INSTALL=${PWD}"/elloappd"




echo "build idgen ..."
cd ${ELLOAPP}/service/idgen/cmd/idgen
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags='-w -s' -o ${INSTALL}/bin_dev/idgen

echo "build status ..."
cd ${ELLOAPP}/service/status/cmd/status
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags='-w -s' -o ${INSTALL}/bin_dev/status

echo "build media ..."
cd ${ELLOAPP}/service/media/cmd/media
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags='-w -s' -o ${INSTALL}/bin_dev/media

echo "build authsession ..."
cd ${ELLOAPP}/service/authsession/cmd/authsession
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags='-w -s' -o ${INSTALL}/bin_dev/authsession

echo "build biz ..."
cd ${ELLOAPP}/service/biz/biz/cmd/biz
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags='-w -s' -o ${INSTALL}/bin_dev/biz

echo "build msg ..."
cd ${ELLOAPP}/messenger/msg/cmd/msg
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags='-w -s' -o ${INSTALL}/bin_dev/msg

echo "build sync ..."
cd ${ELLOAPP}/messenger/sync/cmd/sync
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags='-w -s' -o ${INSTALL}/bin_dev/sync

echo "build bff ..."
cd ${ELLOAPP}/bff/bff/cmd/bff
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags='-w -s' -o ${INSTALL}/bin_dev/bff

echo "build session ..."
cd ${ELLOAPP}/interface/session/cmd/session
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags='-w -s' -o ${INSTALL}/bin_dev/session

echo "build gateway ..."
cd ${ELLOAPP}/interface/gateway/cmd/gateway
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags='-w -s' -o ${INSTALL}/bin_dev/gateway

echo "build dfs ..."
cd ${ELLOAPP}/service/dfs/cmd/dfs
go build -o ${INSTALL}/bin_dev/dfs


