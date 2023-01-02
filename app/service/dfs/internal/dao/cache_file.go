package dao

import (
	"context"
	"fmt"
	"io"
	"strconv"
	"strings"

	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/dfs/internal/model"

	"github.com/zeromicro/go-zero/core/logx"
)

const (
	_cacheFileInfoKeyPrefix = "cache_file_info_%d"
)

func getCacheFileInfoKey(id int64) string {
	return fmt.Sprintf(_cacheFileInfoKeyPrefix, id)
}

func (d *Dao) SetCacheFileInfo(ctx context.Context, id int64, dfsFileInfo *model.DfsFileInfo) (err error) {
	var (
		key = getCacheFileInfoKey(id)
	)

	if err = d.ssdb.Setex(key, fmt.Sprintf("%d_%d", dfsFileInfo.Creator, dfsFileInfo.FileId), 2*60*60); err != nil {
		logx.WithContext(ctx).Errorf("conn.Do(SETEX %s,%v) error(%v)", key, dfsFileInfo, err)
	}

	return
}

func (d *Dao) GetCacheDfsFileInfo(ctx context.Context, id int64) (*model.DfsFileInfo, error) {
	ownerId, fileId, err := d.getCacheFileInfo(ctx, id)
	if err != nil {
		logx.WithContext(ctx).Errorf("getCacheFileInfo (%d, %d) error(%v)", id, err)
		return nil, err
	}

	return d.GetFileInfo(ctx, ownerId, fileId)
}

func (d *Dao) getCacheFileInfo(ctx context.Context, id int64) (ownerId, fileId int64, err error) {
	var (
		key = getCacheFileInfoKey(id)
		s   string
	)

	s, err = d.ssdb.Get(key)
	if err != nil {
		logx.WithContext(ctx).Errorf("getCacheFileInfo(%s) error(%v)", key, err)
		return
	}

	v := strings.Split(s, "_")
	if len(v) != 2 {
		logx.WithContext(ctx).Errorf("split error(len(%v)!=2)", s)
		return
	}

	ownerId, err = strconv.ParseInt(v[0], 10, 64)
	if err != nil {
		logx.WithContext(ctx).Errorf("getCacheFileInfo(%s) error(%v)", key, err)
		return
	}
	fileId, err = strconv.ParseInt(v[1], 10, 64)
	if err != nil {
		logx.WithContext(ctx).Errorf("getCacheFileInfo(%s) error(%v)", key, err)
		return
	}

	return
}

func (d *Dao) GetCacheFile(ctx context.Context, bucket string, id int64, offset int64, limit int32) (bytes []byte, err error) {
	var (
		cacheFile *model.DfsFileInfo
		n         int
	)
	cacheFile, _ = d.GetCacheDfsFileInfo(ctx, id)
	if cacheFile != nil {
		r := d.NewSSDBReader(cacheFile)
		r.Seek(int64(offset), io.SeekStart)
		// TODO(@benqi: check limit)
		bytes = make([]byte, limit)
		n, err = r.Read(bytes)
		if err != nil {
			logx.WithContext(ctx).Errorf("getCacheFile(bucket: %s, id: %d, offset: %d, limit: %d) error :%v",
				bucket,
				id,
				offset,
				limit,
				err)
			return
		}
		bytes = bytes[:n]
	} else {
		path := fmt.Sprintf("%d.dat", id)
		bytes, err = d.GetFile(ctx, bucket, path, offset, limit)
		if err != nil {
			logx.WithContext(ctx).Errorf("getCacheFile(bucket: %s, id: %d: %d, offset: %d, limit: %d) error :%v",
				bucket,
				id,
				offset,
				limit,
				err)
			return
		}
	}

	return
}
