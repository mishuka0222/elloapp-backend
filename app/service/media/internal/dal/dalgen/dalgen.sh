#!/bin/bash

dalgen3 --xml=$1 --db=elloapp --go2=gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/media/internal/dal/dataobject

gofmt -w ../dao/mysql_dao/*.go
gofmt -w ../dataobject/*.go
