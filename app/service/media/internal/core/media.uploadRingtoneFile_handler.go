package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/dfs/dfs"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/media/media"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
)

// MediaUploadRingtoneFile
// media.uploadRingtoneFile flags:# owner_id:long file:InputFile mime_type:string file_name:string = Document;
func (c *MediaCore) MediaUploadRingtoneFile(in *media.TLMediaUploadRingtoneFile) (*mtproto.Document, error) {
	var (
		err      error
		document *mtproto.Document
		file     = in.GetFile()
	)

	if file == nil {
		c.Logger.Errorf("media.uploadRingtoneFile - error: file is nil")
		return nil, mtproto.ErrInputRequestInvalid
	}

	document, err = c.svcCtx.Dao.DfsClient.DfsUploadRingtoneFile(c.ctx, &dfs.TLDfsUploadRingtoneFile{
		Creator:  in.OwnerId,
		File:     file,
		FileName: in.GetFileName(),
		MimeType: in.GetMimeType(),
	})
	if err != nil {
		c.Logger.Errorf("media.uploadRingtoneFile - error: %v", err)
		// err = mtproto.ErrThemeFileInvalid
		return nil, err
	}

	if len(document.GetThumbs()) > 0 {
		c.svcCtx.Dao.SavePhotoSizeV2(c.ctx, document.GetId(), document.GetThumbs())
	}
	c.svcCtx.Dao.SaveDocumentV2(c.ctx, file.GetName(), document)

	return document, nil
}
