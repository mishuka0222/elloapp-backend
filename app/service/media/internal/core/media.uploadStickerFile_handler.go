package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/media/media"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

// MediaUploadStickerFile
// media.uploadStickerFile flags:# owner_id:long file:InputFile thumb:flags.0?InputFile mime_type:string file_name:string document_attribute_sticker:DocumentAttribute = Document;
func (c *MediaCore) MediaUploadStickerFile(in *media.TLMediaUploadStickerFile) (*mtproto.Document, error) {
	// TODO: not impl
	c.Logger.Errorf("media.uploadStickerFile - error: method MediaUploadStickerFile not impl")

	return nil, mtproto.ErrMethodNotImpl
}
