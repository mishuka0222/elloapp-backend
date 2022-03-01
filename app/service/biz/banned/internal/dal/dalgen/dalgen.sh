#!/bin/bash

dalgen2 --xml=$1 --db=teamgram --go2=github.com/teamgram/teamgram-server/app/service/biz/banned/internal/dal/dataobject

gofmt -w ../dao/mysql_dao/*.go
gofmt -w ../dataobject/*.go
