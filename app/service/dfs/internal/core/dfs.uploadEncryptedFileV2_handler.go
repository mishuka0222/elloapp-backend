package core

import (
	"fmt"
	"math/rand"

	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/dfs/dfs"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
)

// DfsUploadEncryptedFileV2
// dfs.uploadEncryptedFileV2 creator:long file:InputEncryptedFile = EncryptedFile;
func (c *DfsCore) DfsUploadEncryptedFileV2(in *dfs.TLDfsUploadEncryptedFileV2) (*mtproto.EncryptedFile, error) {
	var (
		file            = in.GetFile()
		creatorId       = in.GetCreator()
		encryptedFileId = c.svcCtx.Dao.IDGenClient2.NextId(c.ctx)
		accessHash      = int64(mtproto.CRC32_storage_filePartial)<<32 | int64(rand.Uint32())
	)

	fileInfo, err := c.svcCtx.Dao.GetFileInfo(c.ctx, creatorId, file.Id)
	if err != nil {
		c.Logger.Errorf("dfs.uploadDocumentFile - error: %v", err)
		return nil, err
	}
	c.svcCtx.Dao.SetCacheFileInfo(c.ctx, encryptedFileId, fileInfo)
	path := fmt.Sprintf("%d.dat", encryptedFileId)

	go func() {
		_, err2 := c.svcCtx.Dao.PutEncryptedFile(c.ctx, path, c.svcCtx.Dao.NewSSDBReader(fileInfo))
		if err2 != nil {
			c.Logger.Errorf("dfs.uploadEncryptedFile - error: %v", err)
		}
	}()

	encryptedFile := mtproto.MakeTLEncryptedFile(&mtproto.EncryptedFile{
		Id:             encryptedFileId,
		AccessHash:     accessHash,
		Size2_INT32:    int32(fileInfo.GetFileSize()),
		Size2_INT64:    fileInfo.GetFileSize(),
		DcId:           1,
		KeyFingerprint: 0,
	}).To_EncryptedFile()

	return encryptedFile, nil
}
