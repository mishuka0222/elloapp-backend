package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

// UploadReuploadCdnFile
// upload.reuploadCdnFile#9b2754a8 file_token:bytes request_token:bytes = Vector<FileHash>;
func (c *FilesCore) UploadReuploadCdnFile(in *mtproto.TLUploadReuploadCdnFile) (*mtproto.Vector_FileHash, error) {
	// TODO: not impl
	c.Logger.Errorf("upload.reuploadCdnFile blocked, License key from https://elloapp.com required to unlock enterprise features.")

	return &mtproto.Vector_FileHash{
		Datas: []*mtproto.FileHash{},
	}, nil
}
