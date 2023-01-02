package core

import (
	"time"

	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/messenger/sync/sync"
	userpb "gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/user/user"
	mediapb "gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/media/media"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
)

// PhotosUploadProfilePhoto
// photos.uploadProfilePhoto#89f30f69 flags:# file:flags.0?InputFile video:flags.1?InputFile video_start_ts:flags.2?double = photos.Photo;
func (c *PhotosCore) PhotosUploadProfilePhoto(in *mtproto.TLPhotosUploadProfilePhoto) (*mtproto.Photos_Photo, error) {
	photo, err := c.svcCtx.Dao.MediaClient.MediaUploadProfilePhotoFile(c.ctx, &mediapb.TLMediaUploadProfilePhotoFile{
		OwnerId:      c.MD.AuthId,
		File:         in.GetFile(),
		Video:        in.GetVideo(),
		VideoStartTs: in.GetVideoStartTs(),
	})
	if err != nil {
		c.Logger.Errorf("photos.uploadProfilePhoto - error: %v", err)
		return nil, err
	}

	// TODO: ALBUM_PHOTOS_TOO_MANY
	_, err = c.svcCtx.Dao.UserClient.UserUpdateProfilePhoto(c.ctx, &userpb.TLUserUpdateProfilePhoto{
		UserId: c.MD.UserId,
		Id:     photo.GetId(),
	})
	if err != nil {
		c.Logger.Errorf("photos.uploadProfilePhoto - error: %v", err)
		return nil, err
	}

	c.svcCtx.Dao.SyncClient.SyncPushUpdates(c.ctx, &sync.TLSyncPushUpdates{
		UserId: c.MD.UserId,
		Updates: mtproto.MakeUpdatesByUpdates(mtproto.MakeTLUpdateUserPhoto(&mtproto.Update{
			UserId:   c.MD.UserId,
			Date:     int32(time.Now().Unix()),
			Photo:    mtproto.MakeUserProfilePhotoByPhoto(photo),
			Previous: mtproto.BoolFalse,
		}).To_Update()),
	})

	return mtproto.MakeTLPhotosPhoto(&mtproto.Photos_Photo{
		Photo: photo,
		Users: []*mtproto.User{},
	}).To_Photos_Photo(), nil
}
