/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2022 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package dfs_helper

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/dfs/internal/dao"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/dfs/internal/imaging"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/dfs/internal/minio_util"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/dfs/internal/server"

	"github.com/zeromicro/go-zero/core/stores/kv"
	"github.com/zeromicro/go-zero/zrpc"
)

var (
	New = server.New
)

type (
	MinioConfig = minio_util.MinioConfig
	DFSHelper   = dao.Dao
)

var (
	OpenWebp   = imaging.OpenWebp
	DecodeWebp = imaging.DecodeWebp
	EncodeWebp = imaging.EncodeWebp

	Open       = imaging.Open
	Decode     = imaging.Decode
	Resize     = imaging.Resize
	EncodeJpeg = imaging.EncodeJpeg

	EncodeStripped = imaging.EncodeStripped
)

func init() {
	zrpc.DontLogContentForMethod("/dfs.RPCDfs/DfsWriteFilePartData")
}

func NewDFSHelper(minio *MinioConfig, idgen zrpc.RpcClientConf, ssdb kv.KvConf) *DFSHelper {
	return dao.NewDFSHelper(minio, idgen, ssdb)
}
