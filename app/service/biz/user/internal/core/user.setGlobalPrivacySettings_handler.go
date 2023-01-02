package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/user/internal/dal/dataobject"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/user/user"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
)

// UserSetGlobalPrivacySettings
// user.setGlobalPrivacySettings user_id:int settings:GlobalPrivacySettings = Bool;
func (c *UserCore) UserSetGlobalPrivacySettings(in *user.TLUserSetGlobalPrivacySettings) (*mtproto.Bool, error) {
	var (
		archiveAndMuteNewNoncontactPeers = mtproto.FromBool(in.GetSettings().GetArchiveAndMuteNewNoncontactPeers())
	)

	c.svcCtx.Dao.UserGlobalPrivacySettingsDAO.InsertOrUpdate(c.ctx, &dataobject.UserGlobalPrivacySettingsDO{
		UserId:                           in.UserId,
		ArchiveAndMuteNewNoncontactPeers: archiveAndMuteNewNoncontactPeers,
	})

	return mtproto.BoolTrue, nil
}
