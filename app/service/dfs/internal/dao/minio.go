package dao

import (
	"bytes"
	"context"
	"io"
	"path/filepath"

	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/dfs/internal/model"

	"github.com/minio/minio-go"
	"github.com/minio/minio-go/pkg/encrypt"
	"github.com/zeromicro/go-zero/core/logx"
)

//func (d *Dao) Read() {
//
//}

func s3PutOptions(encrypted bool, contentType string) minio.PutObjectOptions {
	options := minio.PutObjectOptions{}
	if encrypted {
		options.ServerSideEncryption = encrypt.NewSSE()
	}
	options.ContentType = contentType

	return options
}

func (d *Dao) GetFile(ctx context.Context, bucket, path string, offset int64, limit int32) (bytes []byte, err error) {
	_ = ctx

	var (
		object *minio.Object
		n      int
	)

	object, err = d.minio.Client.GetObject(bucket, path, minio.GetObjectOptions{})
	if err != nil {
		logx.WithContext(ctx).Errorf("GetFile error: %v")
		return
	}

	bytes = make([]byte, limit)
	n, err = object.ReadAt(bytes, offset)
	//if err != nil {
	//	// return
	//}
	bytes = bytes[:n]
	if n > 0 {
		err = nil
	} else {
		logx.WithContext(ctx).Errorf("GetFile (%s) error: %v", path, err)
	}
	return
}

func (d *Dao) PutPhotoFile(ctx context.Context, path string, buf []byte) (n int64, err error) {
	_ = ctx

	var contentType string
	if ext := filepath.Ext(path); model.IsFileExtImage(ext) {
		contentType = model.GetImageMimeType(ext)
	} else {
		contentType = "binary/octet-stream"
	}

	options := s3PutOptions(false, contentType)
	n, err = d.minio.Client.PutObject("photos", path, bytes.NewReader(buf), int64(len(buf)), options)
	if err != nil {
		logx.WithContext(ctx).Errorf("PutPhotoFile (%s) error: %v", path, err)
	}
	return
}

func (d *Dao) PutPhotoFileV2(ctx context.Context, path string, r io.Reader) (n int64, err error) {
	var (
		contentType string
	)

	if ext := filepath.Ext(path); model.IsFileExtImage(ext) {
		contentType = model.GetImageMimeType(ext)
	} else {
		contentType = "binary/octet-stream"
	}

	options := s3PutOptions(false, contentType)
	n, err = d.minio.Client.PutObject("photos", path, r, -1, options)
	if err != nil {
		logx.Errorf("PutPhotoFile (%s) error: %v", path, err)
	}
	return
}

func (d *Dao) PutVideoFile(ctx context.Context, path string, buf []byte) (n int64, err error) {
	_ = ctx

	var contentType string
	if ext := filepath.Ext(path); model.IsFileExtImage(ext) {
		contentType = model.GetImageMimeType(ext)
	} else {
		contentType = "binary/octet-stream"
	}

	options := s3PutOptions(false, contentType)
	n, err = d.minio.Client.PutObject("videos", path, bytes.NewReader(buf), int64(len(buf)), options)
	if err != nil {
		logx.WithContext(ctx).Errorf("PutPhotoFile (%s) error: %v", path, err)
	}
	return
}

func (d *Dao) PutDocumentFile(ctx context.Context, path string, r io.Reader) (n int64, err error) {
	_ = ctx

	var contentType string
	if ext := filepath.Ext(path); model.IsFileExtImage(ext) {
		contentType = model.GetImageMimeType(ext)
	} else {
		contentType = "binary/octet-stream"
	}

	options := s3PutOptions(false, contentType)
	n, err = d.minio.Client.PutObject("documents", path, r, -1, options)
	if err != nil {
		logx.WithContext(ctx).Errorf("PutDocumentFile (%s) error: %v", path, err)
	}
	return
}

func (d *Dao) FPutDocumentFile(ctx context.Context, path string, r string) (n int64, err error) {
	_ = ctx

	var contentType string
	if ext := filepath.Ext(path); model.IsFileExtImage(ext) {
		contentType = model.GetImageMimeType(ext)
	} else {
		contentType = "binary/octet-stream"
	}

	options := s3PutOptions(false, contentType)
	n, err = d.minio.Client.FPutObject("documents", path, r, options)
	if err != nil {
		logx.WithContext(ctx).Errorf("PutDocumentFile (%s) error: %v", path, err)
	}
	return
}

func (d *Dao) PutEncryptedFile(ctx context.Context, path string, r io.Reader) (n int64, err error) {
	_ = ctx

	options := s3PutOptions(false, "binary/octet-stream")
	n, err = d.minio.Client.PutObject("encryptedfiles", path, r, -1, options)
	if err != nil {
		logx.WithContext(ctx).Errorf("PutEncryptedFile (%s) error: %v", path, err)
	}
	return
}
