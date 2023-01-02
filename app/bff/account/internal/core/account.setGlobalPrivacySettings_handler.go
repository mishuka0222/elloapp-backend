package core

import (
	userpb "gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/user/user"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
)

// AccountSetGlobalPrivacySettings
// account.setGlobalPrivacySettings#1edaaac2 settings:GlobalPrivacySettings = GlobalPrivacySettings;
func (c *AccountCore) AccountSetGlobalPrivacySettings(in *mtproto.TLAccountSetGlobalPrivacySettings) (*mtproto.GlobalPrivacySettings, error) {
	rSettings := in.GetSettings()

	_, err := c.svcCtx.Dao.UserClient.UserSetGlobalPrivacySettings(c.ctx, &userpb.TLUserSetGlobalPrivacySettings{
		UserId:   c.MD.UserId,
		Settings: rSettings,
	})
	if err != nil {
		c.Logger.Errorf("account.setGlobalPrivacySettings - error: %v", err)
		return nil, err
	}

	return rSettings, nil
}
