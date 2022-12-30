package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
)

// UploadGetCdnFile
// upload.getCdnFile#2000bcc3 file_token:bytes offset:int limit:int = upload.CdnFile;
func (c *FilesCore) UploadGetCdnFile(in *mtproto.TLUploadGetCdnFile) (*mtproto.Upload_CdnFile, error) {
	// TODO: not impl
	c.Logger.Errorf("upload.getCdnFile blocked, License key from https://elloapp.com required to unlock enterprise features.")

	return mtproto.MakeTLUploadCdnFile(&mtproto.Upload_CdnFile{
		Bytes: []byte{},
	}).To_Upload_CdnFile(), nil
}
