#!/usr/bin/env bash

PWD=`pwd`
ELLOAPP=${PWD}"/../../app"

cd ${ELLOAPP}/messenger/msg/internal/dal/dalgen
./dalgen_all.sh

cd ${ELLOAPP}/messenger/sync/internal/dal/dalgen
./dalgen_all.sh

cd ${ELLOAPP}/bff/authorization/internal/dal/dalgen
./dalgen_all.sh

cd ${ELLOAPP}/service/biz/chat/internal/dal/dalgen
./dalgen_all.sh

cd ${ELLOAPP}/service/biz/message/internal/dal/dalgen
./dalgen_all.sh

cd ${ELLOAPP}/service/biz/user/internal/dal/dalgen
./dalgen_all.sh

cd ${ELLOAPP}/service/biz/updates/internal/dal/dalgen
./dalgen_all.sh

cd ${ELLOAPP}/service/biz/dialog/internal/dal/dalgen
./dalgen_all.sh

cd ${ELLOAPP}/service/biz/username/internal/dal/dalgen
./dalgen_all.sh

cd ${ELLOAPP}/service/authsession/internal/dal/dalgen
./dalgen_all.sh

cd ${ELLOAPP}/service/media/internal/dal/dalgen
./dalgen_all.sh


