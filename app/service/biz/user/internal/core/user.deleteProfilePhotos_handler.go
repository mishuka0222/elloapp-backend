package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/user/user"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/pkg2/container2"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/pkg2/stores/sqlx"
)

// UserDeleteProfilePhotos
// user.deleteProfilePhotos user_id:long id:Vector<long> = Int64;
func (c *UserCore) UserDeleteProfilePhotos(in *user.TLUserDeleteProfilePhotos) (*mtproto.Int64, error) {
	var (
		meId = in.GetUserId()
	)

	mainPhotoId, err := c.svcCtx.Dao.UsersDAO.SelectProfilePhoto(c.ctx, meId)
	if err != nil {
		c.Logger.Errorf("user.deleteProfilePhotos - error: %v", err)
	}

	if mainPhotoId > 0 {
		if b, _ := container2.Contains(mainPhotoId, in.GetId()); b {
			c.svcCtx.Dao.UserProfilePhotosDAO.Delete(c.ctx, meId, in.GetId())
		} else {
			nextMainId, _ := c.svcCtx.Dao.UserProfilePhotosDAO.SelectNext(c.ctx, meId, in.Id)
			sqlx.TxWrapper(c.ctx, c.svcCtx.Dao.DB, func(tx *sqlx.Tx, result *sqlx.StoreResult) {
				c.svcCtx.Dao.UserProfilePhotosDAO.DeleteTx(tx, meId, in.Id)
				c.svcCtx.Dao.UsersDAO.UpdateProfilePhotoTx(tx, nextMainId, meId)
				//
				mainPhotoId = nextMainId
			})
		}
	} else {
		c.Logger.Errorf("user.deleteProfilePhotos - error: mainPhotoId invalid")
	}

	return &mtproto.Int64{
		V: mainPhotoId,
	}, nil
}
