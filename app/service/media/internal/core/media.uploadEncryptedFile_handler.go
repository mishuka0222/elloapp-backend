package core

import (
	"fmt"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/dfs/dfs"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/media/media"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
)

// MediaUploadEncryptedFile
// media.uploadEncryptedFile owner_id:long file:InputEncryptedFile = EncryptedFile;
func (c *MediaCore) MediaUploadEncryptedFile(in *media.TLMediaUploadEncryptedFile) (*mtproto.EncryptedFile, error) {
	inputFile := in.File
	if inputFile == nil {
		return nil, fmt.Errorf("bad request")
	}

	encryptedFile, err := c.svcCtx.Dao.DfsClient.DfsUploadEncryptedFileV2(c.ctx, &dfs.TLDfsUploadEncryptedFileV2{
		Constructor:          0,
		Creator:              in.OwnerId,
		File:                 inputFile,
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	})
	if err != nil {
		c.Logger.Errorf("media.uploadEncryptedFile - error: %v", err)
		return nil, err
	}
	c.svcCtx.Dao.SaveEncryptedFileV2(c.ctx, encryptedFile)

	return encryptedFile, nil
}
