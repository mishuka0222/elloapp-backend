#!/usr/bin/env bash

PWD=`pwd`
TEAMGRAMAPP=${PWD}"/app"
INSTALL=${PWD}"/teamgramd"



#
#echo "build idgen ..."
#cd ${TEAMGRAMAPP}/service/idgen/cmd/idgen
#CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags='-w -s' -o ${INSTALL}/bin_dev/idgen
#
#echo "build status ..."
#cd ${TEAMGRAMAPP}/service/status/cmd/status
#CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags='-w -s' -o ${INSTALL}/bin_dev/status
#
#echo "build dfs ..."
#cd ${TEAMGRAMAPP}/service/dfs/cmd/dfs
#go build -o ${INSTALL}/bin_dev/dfs
#
#echo "build media ..."
#cd ${TEAMGRAMAPP}/service/media/cmd/media
#CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags='-w -s' -o ${INSTALL}/bin_dev/media
#
#echo "build authsession ..."
#cd ${TEAMGRAMAPP}/service/authsession/cmd/authsession
#CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags='-w -s' -o ${INSTALL}/bin_dev/authsession
#
#echo "build biz ..."
#cd ${TEAMGRAMAPP}/service/biz/biz/cmd/biz
#CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags='-w -s' -o ${INSTALL}/bin_dev/biz

#echo "build msg ..."
#cd ${TEAMGRAMAPP}/messenger/msg/cmd/msg
#CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags='-w -s' -o ${INSTALL}/bin_dev/msg

#echo "build sync ..."
#cd ${TEAMGRAMAPP}/messenger/sync/cmd/sync
#CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags='-w -s' -o ${INSTALL}/bin_dev/sync

echo "build bff ..."
cd ${TEAMGRAMAPP}/bff/bff/cmd/bff
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags='-w -s' -o ${INSTALL}/bin_dev/bff

echo "build session ..."
cd ${TEAMGRAMAPP}/interface/session/cmd/session
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags='-w -s' -o ${INSTALL}/bin_dev/session

#echo "bui
# ld gateway ..."
#cd ${TEAMGRAMAPP}/interface/gateway/cmd/gateway
#CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags='-w -s' -o ${INSTALL}/bin_dev/gateway
