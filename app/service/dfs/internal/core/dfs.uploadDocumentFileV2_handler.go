package core

import (
	"bytes"
	"fmt"
	"image"
	"math/rand"

	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/dfs/dfs"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/dfs/internal/imaging"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/dfs/internal/model"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/pkg2/bytes2"
)

// DfsUploadDocumentFileV2
// dfs.uploadDocumentFileV2 creator:long media:InputMedia = Document;
func (c *DfsCore) DfsUploadDocumentFileV2(in *dfs.TLDfsUploadDocumentFileV2) (*mtproto.Document, error) {
	var (
		documentId = c.svcCtx.Dao.IDGenClient2.NextId(c.ctx)
		//idgen.GetUUID()
		file      = in.GetMedia().GetFile()
		cacheData []byte
	)

	// 有点难理解，主要是为了不在这里引入snowflake
	ext := model.GetFileExtName(file.GetName())
	extType := model.GetStorageFileTypeConstructor(ext)
	accessHash := int64(extType)<<32 | int64(rand.Uint32())

	r, err := c.svcCtx.Dao.OpenFile(c.ctx, in.GetCreator(), file.Id, file.Parts)
	if err != nil {
		c.Logger.Errorf("dfs.uploadDocumentFile - %v", err)
		return nil, mtproto.ErrMediaInvalid
	}

	//fileInfo, err := s.Dao.GetFileInfo(ctx, creatorId, file.Id)
	//if err != nil {
	//	log.Errorf("dfs.uploadDocumentFile - error: %v", err)
	//	return nil, err
	//}

	c.svcCtx.Dao.SetCacheFileInfo(c.ctx, documentId, r.DfsFileInfo)

	//go func() {
	//	_, err2 := s.Dao.PutDocumentFile(ctx,
	//		fmt.Sprintf("%d.dat", documentId),
	//		s.Dao.NewSSDBReader(r.DfsFileInfo))
	//	if err2 != nil {
	//		log.Errorf("dfs.uploadDocumentFile - error: %v", err2)
	//	}
	//}()

	attributes := make([]*mtproto.DocumentAttribute, 0, len(in.GetMedia().Attributes))
	for _, attr := range in.GetMedia().GetAttributes() {
		if attr.GetPredicateName() != mtproto.Predicate_documentAttributeAnimated {
			attributes = append(attributes, attr)
		}
	}

	// document#1e87342b flags:#
	//	id:long
	//	access_hash:long
	//	file_reference:bytes
	//	date:int
	//	mime_type:string
	//	size:int
	//	thumbs:flags.0?Vector<PhotoSize>
	//	video_thumbs:flags.1?Vector<VideoSize>
	//	dc_id:int
	//	attributes:Vector<DocumentAttribute> = Document;
	document := mtproto.MakeTLDocument(&mtproto.Document{
		Id:            documentId,
		AccessHash:    accessHash,
		FileReference: nil,
		Date:          int32(r.DfsFileInfo.Mtime),
		MimeType:      in.GetMedia().GetMimeType(),
		Size2_INT32:   int32(r.DfsFileInfo.GetFileSize()),
		Size2_INT64:   r.DfsFileInfo.GetFileSize(),
		Thumbs:        nil,
		VideoThumbs:   nil,
		DcId:          1,
		Attributes:    attributes,
	}).To_Document()

	if mtproto.IsMimeAcceptedForPhotoVideoAlbum(document.MimeType) && model.IsFileExtImage(ext) {
		var (
			thumb image.Image
			// photoId = idgen.GetUUID()
			// ext2     = ".jpg"
			// extType2 = model.GetStorageFileTypeConstructor(ext2)
			// secretId = int64(extType2)<<32 | int64(rand.Uint32())
		)

		cacheData, err = r.ReadAll(c.ctx)
		if err != nil {
			c.Logger.Errorf("dfs.uploadDocumentFile - %v", err)
			return nil, mtproto.ErrWallpaperFileInvalid
		} else {
			// log.Debugf("cacheData: %s", hex.EncodeToString(cacheData))
		}

		// build photoStrippedSize
		thumb, err = imaging.Decode(bytes.NewReader(cacheData))
		if err != nil {
			// ioutil.WriteFile("./t.jpg", cacheData, 0644)
			c.Logger.Errorf("dfs.uploadDocumentFile - error: %v", err)
			return nil, err

		}
		stripped := bytes2.NewBuffer(make([]byte, 0, 4096))
		if thumb.Bounds().Dx() >= thumb.Bounds().Dy() {
			err = imaging.EncodeStripped(stripped, imaging.Resize(thumb, 40, 0), 30)
		} else {
			err = imaging.EncodeStripped(stripped, imaging.Resize(thumb, 0, 40), 30)
		}
		if err != nil {
			c.Logger.Errorf("dfs.uploadDocumentFile - error: %v", err)
			return nil, err
		}

		// upload thumb
		var (
			mThumbData = bytes2.NewBuffer(make([]byte, 0, len(cacheData)))
			mThumb     image.Image
		)
		if thumb.Bounds().Dx() >= thumb.Bounds().Dy() {
			mThumb = imaging.Resize(thumb, 320, 0)
			// err = imaging.Encode(mThumbData, mThumb, 80)
		} else {
			mThumb = imaging.Resize(thumb, 0, 320)
			// err = imaging.Encode(mThumbData, imaging.Resize(thumb, 0, 320), 80)
		}

		err = imaging.EncodeJpeg(mThumbData, mThumb)
		if err != nil {
			c.Logger.Errorf("dfs.uploadDocumentFile - error: %v", err)
			return nil, err
		}

		// upload thumb
		path := fmt.Sprintf("%s/%d.dat", mtproto.PhotoSZMediumType, documentId)
		// upload
		c.svcCtx.Dao.PutPhotoFile(c.ctx, path, mThumbData.Bytes())

		document.Thumbs = []*mtproto.PhotoSize{
			mtproto.MakeTLPhotoStrippedSize(&mtproto.PhotoSize{
				Type:  mtproto.PhotoSZStrippedType,
				Bytes: stripped.Bytes(),
			}).To_PhotoSize(),
			mtproto.MakeTLPhotoSize(&mtproto.PhotoSize{
				Type:  mtproto.PhotoSZMediumType,
				W:     int32(mThumb.Bounds().Dx()),
				H:     int32(mThumb.Bounds().Dy()),
				Size2: int32(len(mThumbData.Bytes())),
			}).To_PhotoSize(),
		}
	}

	defer func() {
		go func() {
			if mtproto.IsMimeAcceptedForPhotoVideoAlbum(document.MimeType) && model.IsFileExtImage(ext) {
				_, err2 := c.svcCtx.Dao.PutDocumentFile(c.ctx,
					fmt.Sprintf("%d.dat", documentId),
					bytes.NewReader(cacheData))
				if err2 != nil {
					c.Logger.Errorf("dfs.uploadDocumentFile - error: %v", err2)
				}
			} else {
				_, err2 := c.svcCtx.Dao.PutDocumentFile(c.ctx,
					fmt.Sprintf("%d.dat", documentId),
					c.svcCtx.Dao.NewSSDBReader(r.DfsFileInfo))
				if err2 != nil {
					c.Logger.Errorf("dfs.uploadDocumentFile - error: %v", err2)
				}
			}
		}()
	}()

	return document, nil
}
