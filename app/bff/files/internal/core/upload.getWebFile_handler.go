package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
	"google.golang.org/grpc/status"
)

var (
	ErrWebfileNotAvailable = status.Error(mtproto.ErrBadRequest, "WEBFILE_NOT_AVAILABLE")
)

// UploadGetWebFile
// upload.getWebFile#24e6818d location:InputWebFileLocation offset:int limit:int = upload.WebFile;
func (c *FilesCore) UploadGetWebFile(in *mtproto.TLUploadGetWebFile) (*mtproto.Upload_WebFile, error) {
	switch in.GetLocation().GetPredicateName() {
	case mtproto.Predicate_inputWebFileAudioAlbumThumbLocation:
		err := ErrWebfileNotAvailable
		c.Logger.Errorf("upload.getWebFile - error: %v", err)

		return nil, err
	default:
		c.Logger.Errorf("upload.getWebFile blocked, License key from https://elloapp.com required to unlock enterprise features.")

		return nil, mtproto.ErrEnterpriseIsBlocked
	}
}
