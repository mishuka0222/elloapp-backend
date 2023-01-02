package core

import (
	userpb "gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/user/user"
	mediapb "gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/media/media"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
)

// PhotosGetUserPhotos
// photos.getUserPhotos#91cd32a8 user_id:InputUser offset:int max_id:long limit:int = photos.Photos;
func (c *PhotosCore) PhotosGetUserPhotos(in *mtproto.TLPhotosGetUserPhotos) (*mtproto.Photos_Photos, error) {
	userId := mtproto.FromInputUser(c.MD.UserId, in.UserId)
	switch userId.PeerType {
	case mtproto.PEER_SELF, mtproto.PEER_USER:
	default:
		err := mtproto.ErrUserIdInvalid
		c.Logger.Errorf("photos.getUserPhotos - error: %v", err)
		return nil, err
	}

	cachePhotos, err := c.svcCtx.Dao.UserClient.UserGetProfilePhotos(c.ctx, &userpb.TLUserGetProfilePhotos{
		UserId: userId.PeerId,
	})
	if err != nil {
		c.Logger.Errorf("photos.getUserPhotos - error: %v", err)
		return nil, err
	}

	photos := mtproto.MakeTLPhotosPhotos(&mtproto.Photos_Photos{
		Photos: make([]*mtproto.Photo, 0, len(cachePhotos.GetDatas())),
		Users:  []*mtproto.User{},
	}).To_Photos_Photos()

	for _, id := range cachePhotos.GetDatas() {
		if photo, err := c.svcCtx.Dao.MediaClient.MediaGetPhoto(c.ctx,
			&mediapb.TLMediaGetPhoto{
				PhotoId: id,
			}); err != nil {
			c.Logger.Errorf("photos.getUserPhotos - error: %v", err)
		} else if photo != nil {
			photos.Photos = append(photos.Photos, photo)
		}
	}

	return photos, nil
}
