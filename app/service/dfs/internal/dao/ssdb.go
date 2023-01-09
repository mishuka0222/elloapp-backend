package dao

import (
	"context"
	"fmt"
	"strconv"

	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/dfs/internal/model"

	"github.com/zeromicro/go-zero/core/logx"
)

const (
	_fileKeyPrefix     = "file_%d_%d"
	_fileInfoKeyPrefix = "file_info_%d_%d"
)

func getFileKey(ownerId, fileId int64) string {
	return fmt.Sprintf(_fileKeyPrefix, ownerId, fileId)
}

func getFileInfoKey(ownerId, fileId int64) string {
	return fmt.Sprintf(_fileInfoKeyPrefix, ownerId, fileId)
}

func (d *Dao) WriteFilePartData(ctx context.Context, ownerId, fileId int64, filePart int32, bytes []byte) (err error) {
	var (
		k = getFileKey(ownerId, fileId)
	)

	err = d.ssdb.Hset(k, strconv.Itoa(int(filePart)), string(bytes))
	if err != nil {
		logx.WithContext(ctx).Errorf("conn.Send(HSET %d, %d, %s) error(%v)", ownerId, fileId, filePart, err)
		return
	}

	err = d.ssdb.Expire(k, ssdbExpire)
	if err != nil {
		logx.WithContext(ctx).Errorf("conn.Send(EXPIRE %d,%d,%s) error(%v)", ownerId, fileId, filePart, err)
		return
	}

	return
}

func (d *Dao) ReadFile(ctx context.Context, ownerId, fileId int64, parts int32) (partLength int32, bytes []byte, err error) {
	var (
		k = getFileKey(ownerId, fileId)
	)

	for i := int32(0); i < parts; i++ {
		var (
			bBuf string
			b    []byte
		)
		bBuf, err = d.ssdb.Hget(k, strconv.Itoa(int(i)))
		if err != nil {
			logx.WithContext(ctx).Errorf("conn.Send(HGET %s, %d) error(%v)", k, i, err)
			return 0, nil, err
		}
		b = []byte(bBuf)
		if bytes == nil {
			bytes = make([]byte, 0, len(b)*int(parts))
		}
		if i == 0 {
			partLength = int32(len(b))
		}
		bytes = append(bytes, b...)
	}

	return
}

func (d *Dao) ReadFileCB(ctx context.Context, ownerId, fileId int64, parts int32, cb func(part int32, bytes []byte) error) (err error) {
	var (
		k    = getFileKey(ownerId, fileId)
		bBuf string
	)

	for i := int32(0); i < parts; i++ {
		bBuf, err = d.ssdb.Hget(k, strconv.Itoa(int(i)))
		if err != nil {
			logx.WithContext(ctx).Errorf("conn.Do(HGET %s, %d) error(%v)", k, i, err)
			return
		}
		if err = cb(i, []byte(bBuf)); err != nil {
			return
		}
	}

	return
}

func (d *Dao) ReadOffsetLimit(ctx context.Context, fileInfo *model.DfsFileInfo, offset, limit int32) (bytes []byte, err error) {
	if limit == 0 && offset != 0 {
		return
	}

	var (
		k    = getFileKey(fileInfo.Creator, fileInfo.FileId)
		bBuf string
		b    []byte
	)

	if limit == 0 && offset == 0 {
		for i := 0; i < fileInfo.FileTotalParts; i++ {
			bBuf, err = d.ssdb.Hget(k, strconv.Itoa(int(i)))
			if err != nil {
				logx.WithContext(ctx).Errorf("conn.Send(HGET %s, %d) error(%v)", k, i, err)
				return
			}
			b = []byte(bBuf)
			if bytes == nil {
				bytes = make([]byte, 0, len(b)*fileInfo.FileTotalParts)
			}
			bytes = append(bytes, b...)
		}
	} else {
		var (
			bPart = int(offset) / fileInfo.FilePartSize
			ePart = int(offset+limit) / fileInfo.FilePartSize
			bP    = int(offset) % fileInfo.FilePartSize
			eP    = int(offset+limit) % fileInfo.FilePartSize
		)

		bytes = make([]byte, 0, limit)
		for i := bPart; i <= ePart; i++ {
			bBuf, err = d.ssdb.Hget(k, strconv.Itoa(i))
			if err != nil {
				logx.WithContext(ctx).Errorf("conn.Send(HGET %s, %d) error(%v)", k, i, err)
				return
			}
			b = []byte(bBuf)
			if i == bPart {
				if i == ePart {
					bytes = append(bytes, b[bP:eP]...)
				} else {
					bytes = append(bytes, b[bP:]...)
				}
			} else if i == ePart {
				bytes = append(bytes, b[:eP]...)
			} else {
				bytes = append(bytes, b...)
			}
		}
	}

	return
}

func (d *Dao) OpenFile(ctx context.Context, ownerId, fileId int64, parts int32) (*SSDBReader, error) {
	fileInfo, err := d.GetFileInfo(ctx, ownerId, fileId)
	if err != nil {
		return nil, err
	}
	if parts > 0 {
		// TODO(@benqi): check fileInfo.FileTotalParts == parts
	}
	return d.NewSSDBReader(fileInfo), nil
}
