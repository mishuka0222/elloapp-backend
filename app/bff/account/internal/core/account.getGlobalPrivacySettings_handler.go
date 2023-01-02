package core

import (
	userpb "gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/user/user"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
)

// AccountGetGlobalPrivacySettings
// account.getGlobalPrivacySettings#eb2b4cf6 = GlobalPrivacySettings;
func (c *AccountCore) AccountGetGlobalPrivacySettings(in *mtproto.TLAccountGetGlobalPrivacySettings) (*mtproto.GlobalPrivacySettings, error) {
	_ = in
	globalPrivacySettings, err := c.svcCtx.Dao.UserGetGlobalPrivacySettings(c.ctx, &userpb.TLUserGetGlobalPrivacySettings{
		UserId: c.MD.UserId,
	})
	if err != nil {
		c.Logger.Errorf("account.getGlobalPrivacySettings - error: %v", err)
		return nil, err
	}

	return globalPrivacySettings, nil
}
