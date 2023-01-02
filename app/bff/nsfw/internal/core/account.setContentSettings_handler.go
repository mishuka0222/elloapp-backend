package core

import (
	userpb "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/user/user"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

// AccountSetContentSettings
// account.setContentSettings#b574b16b flags:# sensitive_enabled:flags.0?true = Bool;
func (c *NsfwCore) AccountSetContentSettings(in *mtproto.TLAccountSetContentSettings) (*mtproto.Bool, error) {
	rValue, err := c.svcCtx.Dao.UserClient.UserSetContentSettings(c.ctx, &userpb.TLUserSetContentSettings{
		UserId:           c.MD.UserId,
		SensitiveEnabled: in.SensitiveEnabled,
	})
	if err != nil {
		c.Logger.Errorf("account.setGlobalPrivacySettings - error: %v", err)
		return nil, err
	}

	return rValue, nil
}
