package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/media/media"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

// MediaGetPhoto
// media.getPhoto photo_id:long = Photo;
func (c *MediaCore) MediaGetPhoto(in *media.TLMediaGetPhoto) (*mtproto.Photo, error) {
	photo, err := c.svcCtx.Dao.GetPhotoV2(c.ctx, in.GetPhotoId())
	if err != nil {
		c.Logger.Errorf("media.getPhotoFileData(%d) - error: %v", in.GetPhotoId(), err)
		return nil, err
	}

	return photo, nil
}
