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
