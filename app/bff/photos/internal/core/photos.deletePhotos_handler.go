package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/messenger/sync/sync"
	userpb "gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/user/user"
	mediapb "gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/media/media"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
	"time"
)

// PhotosDeletePhotos
// photos.deletePhotos#87cf7f2f id:Vector<InputPhoto> = Vector<long>;
func (c *PhotosCore) PhotosDeletePhotos(in *mtproto.TLPhotosDeletePhotos) (*mtproto.Vector_Long, error) {
	var (
		photo        *mtproto.Photo
		deleteIdList = make([]int64, 0, len(in.GetId()))
	)

	for _, id := range in.GetId() {
		deleteIdList = append(deleteIdList, id.GetId())
	}

	// TODO: ALBUM_PHOTOS_TOO_MANY
	mainId, err := c.svcCtx.Dao.UserClient.UserDeleteProfilePhotos(c.ctx, &userpb.TLUserDeleteProfilePhotos{
		UserId: c.MD.UserId,
		Id:     deleteIdList,
	})
	if err != nil {
		c.Logger.Errorf("photos.deletePhotos - error: %v", err)
		return nil, err
	}

	if mainId.V != 0 {
		photo, err = c.svcCtx.Dao.MediaClient.MediaGetPhoto(c.ctx, &mediapb.TLMediaGetPhoto{
			PhotoId: mainId.V,
		})
		if err != nil {
			c.Logger.Errorf("photos.deletePhotos - error: %v", err)
			return nil, err
		}
	}

	if photo == nil {
		photo = mtproto.MakeTLPhotoEmpty(nil).To_Photo()
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

	return &mtproto.Vector_Long{
		Datas: deleteIdList,
	}, nil
}
