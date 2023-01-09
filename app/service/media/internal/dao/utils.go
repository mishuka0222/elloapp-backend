package dao

import (
	"path"
	"strings"

	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

var (
	emptyPhoto = mtproto.MakeTLPhotoEmpty(nil).To_Photo()
)

func getFileExtName(filePath string) string {
	var ext = path.Ext(filePath)
	if ext == "" {
		ext = "partial"
	}
	return strings.ToLower(ext)
}
