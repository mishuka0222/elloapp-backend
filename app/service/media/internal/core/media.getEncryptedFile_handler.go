package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/media/media"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
)

// MediaGetEncryptedFile
// media.getEncryptedFile id:long access_hash:long = EncryptedFile;
func (c *MediaCore) MediaGetEncryptedFile(in *media.TLMediaGetEncryptedFile) (*mtproto.EncryptedFile, error) {
	encryptedFile, err := c.svcCtx.Dao.GetEncryptedFile(c.ctx, in.Id, in.AccessHash)
	if err != nil {
		c.Logger.Errorf("media.getEncryptedFile - error: %v", err)
		return nil, err
	}

	return encryptedFile, nil
}
