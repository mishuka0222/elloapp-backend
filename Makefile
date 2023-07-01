

#LOCAL
PWD=`pwd`
ELLOAPP=${PWD}/app
INSTALL=${PWD}/elloappd
I_FOLDER=bin_dev
ELLOAPP_BIN_PATH=${INSTALL}/${I_FOLDER}
ELLOAPP_CONFIG_PATH=${INSTALL}/etc_dev
ELLOAPP_SQL_PATH=${INSTALL}/sql

#SSH
SSH_USER=root
SSH_PORT=65.109.139.165
SSH_PROJECT_PATH=proj/elloapp_tg_backend/elloappd
SSH_PATH=${SSH_USER}@${SSH_PORT}:/${SSH_USER}/${SSH_PROJECT_PATH}
#SSH_PATH="${SSH_USER}@${SSH_PORT}:/home/${SSH_USER}/${SSH_PROJECT_PATH}"


copy_sql_migrations:
	@echo "copy etc configs ..."
	scp -r ${ELLOAPP_SQL_PATH}/  ${SSH_PATH}/

copy_etc_configs:
	@echo "copy etc configs ..."
	scp -r ${ELLOAPP_CONFIG_PATH}/  ${SSH_PATH}/

copy_all_bin_files:
	@echo "copy all bin files ..."
	scp -r ${ELLOAPP_BIN_PATH}/  ${SSH_PATH}/

copy_bin_config_files:
	@echo "copy bin config files ..."
	scp -r ${ELLOAPP_BIN_PATH}/config.json  ${SSH_PATH}/bin_dev/
	scp -r ${ELLOAPP_BIN_PATH}/server_pkcs1.key  ${SSH_PATH}/bin_dev/

#idgen
build_idgen:
	@echo "build idgen ..."
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags='-w -s' -o ${ELLOAPP_BIN_PATH}/idgen ${ELLOAPP}/service/idgen/cmd/idgen

copy_idgen:
	@echo "copy idgen ..."
	scp -r ${ELLOAPP_BIN_PATH}/idgen  ${SSH_PATH}/bin_dev/

do_idgen: build_idgen copy_idgen

#status
build_status:
	@echo "build status ..."
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags='-w -s' -o ${ELLOAPP_BIN_PATH}/status ${ELLOAPP}/service/status/cmd/status

copy_status:
	@echo "copy status ..."
	scp -r ${ELLOAPP_BIN_PATH}/status  ${SSH_PATH}/bin_dev/

do_status: build_status copy_status

#media
build_media:
	@echo "build media ..."
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags='-w -s' -o ${ELLOAPP_BIN_PATH}/media ${ELLOAPP}/service/media/cmd/media

copy_media:
	@echo "copy media ..."
	scp -r ${ELLOAPP_BIN_PATH}/media  ${SSH_PATH}/bin_dev/

do_media: build_media copy_media

#authsession
build_authsession:
	@echo "build authsession ..."
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags='-w -s' -o ${ELLOAPP_BIN_PATH}/authsession ${ELLOAPP}/service/authsession/cmd/authsession

copy_authsession:
	@echo "copy authsession ..."
	scp -r ${ELLOAPP_BIN_PATH}/authsession  ${SSH_PATH}/bin_dev/

do_authsession: build_authsession copy_authsession

#biz
build_biz:
	@echo "build biz ..."
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags='-w -s' -o ${ELLOAPP_BIN_PATH}/biz ${ELLOAPP}/service/biz/biz/cmd/biz

copy_biz:
	@echo "copy biz ..."
	scp -r ${ELLOAPP_BIN_PATH}/biz  ${SSH_PATH}/bin_dev/

do_biz: build_biz copy_biz

#msg
build_msg:
	@echo "build msg ..."
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags='-w -s' -o ${ELLOAPP_BIN_PATH}/msg ${ELLOAPP}/messenger/msg/cmd/msg

copy_msg:
	@echo "copy msg ..."
	scp -r ${ELLOAPP_BIN_PATH}/msg  ${SSH_PATH}/bin_dev/

do_msg: build_msg copy_msg

#sync
build_sync:
	@echo "build sync ..."
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags='-w -s' -o ${ELLOAPP_BIN_PATH}/sync ${ELLOAPP}/messenger/sync/cmd/sync

copy_sync:
	@echo "copy sync ..."
	scp -r ${ELLOAPP_BIN_PATH}/sync  ${SSH_PATH}/bin_dev/

do_sync: build_sync copy_sync

#bff
build_bff:
	@echo "build bff ..."
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags='-w -s' -o ${ELLOAPP_BIN_PATH}/bff ${ELLOAPP}/bff/bff/cmd/bff

copy_bff:
	@echo "copy bff ..."
	scp -r ${ELLOAPP_BIN_PATH}/bff  ${SSH_PATH}/bin_dev/

do_bff: build_bff copy_bff

#session
build_session:
	@echo "build session ..."
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags='-w -s' -o ${ELLOAPP_BIN_PATH}/session ${ELLOAPP}/interface/session/cmd/session

copy_session:
	@echo "copy session ..."
	scp -r ${ELLOAPP_BIN_PATH}/session  ${SSH_PATH}/bin_dev/

do_session: build_session copy_session

#gateway
build_gateway:
	@echo "build gateway ..."
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags='-w -s' -o ${ELLOAPP_BIN_PATH}/gateway ${ELLOAPP}/interface/gateway/cmd/gateway

copy_gateway:
	@echo "copy gateway ..."
	scp -r ${ELLOAPP_BIN_PATH}/gateway  ${SSH_PATH}/bin_dev/

do_gateway: build_gateway copy_gateway


# NEW <=============
#===================
#===================
#===================
#===================
#===================
#===================
#===================
#===================
# OLD =============>

VERSION=v0.87.1-elloapp-server
BUILD=`date +%F`
SHELL := /bin/bash
BASEDIR = $(shell pwd)

# build with verison infos
versionDir="gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/pkg2/version"
gitTag=$(shell if [ "`git describe --tags --abbrev=0 2>/dev/null`" != "" ];then git describe --tags --abbrev=0; else git log --pretty=format:'%h' -n 1; fi)
gitBranch=$(shell git rev-parse --abbrev-ref HEAD)
buildDate=$(shell TZ=Asia/Shanghai date +%FT%T%z)
gitCommit=$(shell git rev-parse --short HEAD)
gitTreeState=$(shell if git status|grep -q 'clean';then echo clean; else echo dirty; fi)

ldflags="-s -w -X ${versionDir}.gitTag=${gitTag} -X ${versionDir}.buildDate=${buildDate} -X ${versionDir}.gitCommit=${gitCommit} -X ${versionDir}.gitTreeState=${gitTreeState} -X ${versionDir}.version=${VERSION} -X ${versionDir}.gitBranch=${gitBranch}"

all: idgen status dfs media authsession biz msg sync bff session gateway

idgen:
	@echo "build idgen..."
	@go build -ldflags ${ldflags} -o elloappd/bin/idgen -tags=jsoniter app/service/idgen/cmd/idgen/*.go

status:
	@echo "build status..."
	@go build -ldflags ${ldflags} -o elloappd/bin/status -tags=jsoniter app/service/status/cmd/status/*.go

dfs:
	@echo "build dfs..."
	@go build -ldflags ${ldflags} -o elloappd/bin/dfs -tags=jsoniter app/service/dfs/cmd/dfs/*.go

media:
	@echo "build media..."
	@go build -ldflags ${ldflags} -o elloappd/bin/media -tags=jsoniter app/service/media/cmd/media/*.go

authsession:
	@echo "build authsession..."
	@go build -ldflags ${ldflags} -o elloappd/bin/authsession -tags=jsoniter app/service/authsession/cmd/authsession/*.go

biz:
	@echo "build biz..."
	@go build -ldflags ${ldflags} -o elloappd/bin/biz -tags=jsoniter app/service/biz/biz/cmd/biz/*.go

msg:
	@echo "build msg..."
	@go build -ldflags ${ldflags} -o elloappd/bin/msg -tags=jsoniter app/messenger/msg/cmd/msg/*.go

sync:
	@echo "build sync..."
	@go build -ldflags ${ldflags} -o elloappd/bin/sync -tags=jsoniter app/messenger/sync/cmd/sync/*.go

bff:
	@echo "build bff..."
	@go build -ldflags ${ldflags} -o elloappd/bin/bff -tags=jsoniter app/bff/bff/cmd/bff/*.go

session:
	@echo "build session..."
	@go build -ldflags ${ldflags} -o elloappd/bin/session -tags=jsoniter app/interface/session/cmd/session/*.go

gateway:
	@echo "build gateway..."
	@go build -ldflags ${ldflags} -o elloappd/bin/gateway -tags=jsoniter app/interface/gateway/cmd/gateway/*.go

clean:
	@rm -rf elloappd/bin/idgen
	@rm -rf elloappd/bin/status
	@rm -rf elloappd/bin/dfs
	@rm -rf elloappd/bin/media
	@rm -rf elloappd/bin/authsession
	@rm -rf elloappd/bin/biz
	@rm -rf elloappd/bin/msg
	@rm -rf elloappd/bin/sync
	@rm -rf elloappd/bin/bff
	@rm -rf elloappd/bin/session
	@rm -rf elloappd/bin/gateway
