package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/dfs/dfs"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

// UploadGetFile
// upload.getFile#b15a9afc flags:# precise:flags.0?true cdn_supported:flags.1?true location:InputFileLocation offset:int limit:int = upload.File;
func (c *FilesCore) UploadGetFile(in *mtproto.TLUploadGetFile) (*mtproto.Upload_File, error) {
	var (
		location = in.GetLocation()
		offset   = in.GetOffset_INT64()
		limit    = in.GetLimit()
	)

	if offset == 0 {
		offset = int64(in.GetOffset_INT32())
	}

	switch location.GetPredicateName() {
	case mtproto.Predicate_inputFileLocation:
		// inputFileLocation#dfdaabe1
		//	volume_id:long
		//	local_id:int
		//	secret:long
		//	file_reference:bytes = InputFileLocation;

		err := mtproto.ErrInputRequestInvalid
		c.Logger.Errorf("upload.getFile - error: %v inputFileLocation", err)
		return nil, err
	case mtproto.Predicate_inputEncryptedFileLocation:
		// inputEncryptedFileLocation#f5235d55
		//	id:long
		//	access_hash:long = InputFileLocation;

	case mtproto.Predicate_inputDocumentFileLocation:
		// inputDocumentFileLocation#bad07584
		//	id:long
		//	access_hash:long
		//	file_reference:bytes
		//	thumb_size:string = InputFileLocation;
		//
	case mtproto.Predicate_inputSecureFileLocation:
		// inputSecureFileLocation#cbc7ee28
		//	id:long
		//	access_hash:long = InputFileLocation;
		//
	case mtproto.Predicate_inputTakeoutFileLocation:
		// inputTakeoutFileLocation#29be5899 = InputFileLocation;
		//
	case mtproto.Predicate_inputPhotoFileLocation:
		// inputPhotoFileLocation#40181ffe
		//	id:long
		//	access_hash:long
		//	file_reference:bytes
		//	thumb_size:string = InputFileLocation;
		//
	case mtproto.Predicate_inputPeerPhotoFileLocation:
		// inputPeerPhotoFileLocation#37257e99 flags:#
		//	big:flags.0?true
		//	peer:InputPeer
		//	photo_id:long = InputFileLocation;
		//
	case mtproto.Predicate_inputStickerSetThumb:
		// inputStickerSetThumb#9d84f3db
		//	stickerset:InputStickerSet
		//	thumb_version:int = InputFileLocation;
		if c.svcCtx.Plugin != nil {
			location2, err := c.svcCtx.Plugin.GetStickerSetThumbFileLocation(c.ctx,
				c.MD.UserId,
				location.GetStickerset(),
				location.GetThumbVersion())
			if err != nil {
				c.Logger.Errorf("upload.getFile - error: %v inputFileLocation", err)
				return nil, mtproto.ErrStickerIdInvalid
			}
			location = location2
		} else {
			c.Logger.Errorf("upload.getFile blocked, License key from https://elloapp.com required to unlock enterprise features.")
			return nil, mtproto.ErrEnterpriseIsBlocked
		}
	case mtproto.Predicate_inputGroupCallStream:
		// inputGroupCallStream#bba51639
		//	call:InputGroupCall
		//	time_ms:long
		//	scale:int = InputFileLocation;
		//
		if c.svcCtx.Plugin != nil {
			uploadFile, err := c.svcCtx.Plugin.GetGroupCallStreamFile(
				c.ctx,
				c.MD.UserId,
				in.GetLocation())
			if err != nil {
				c.Logger.Errorf("upload.getFile - error: %v inputFileLocation", err)
				return nil, mtproto.ErrGroupcallForbidden
			}
			return uploadFile, nil
		}
	default:
		c.Logger.Errorf("upload.getFile - error: invalid location")
		return nil, mtproto.ErrLocationInvalid
	}

	uploadFile, err := c.svcCtx.Dao.DfsClient.DfsDownloadFile(c.ctx, &dfs.TLDfsDownloadFile{
		Location: location,
		Offset:   offset,
		Limit:    limit,
	})
	if err != nil {
		c.Logger.Errorf("upload.getFile - error: %v", err.Error())
		// err = mtproto.ErrOffsetInvalid
		return nil, err
	}
	if uploadFile.GetType().GetPredicateName() == mtproto.Predicate_storage_fileUnknown {
		uploadFile.Type.PredicateName = mtproto.Predicate_storage_filePartial
	}

	return uploadFile, nil
}
