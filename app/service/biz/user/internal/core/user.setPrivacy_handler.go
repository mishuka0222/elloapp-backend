package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/user/user"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

// UserSetPrivacy
// user.setPrivacy user_id:int key_type:int rules:Vector<PrivacyRule> = Bool;
func (c *UserCore) UserSetPrivacy(in *user.TLUserSetPrivacy) (*mtproto.Bool, error) {
	c.svcCtx.Dao.SetUserPrivacyRules(
		c.ctx,
		in.GetUserId(),
		in.GetKeyType(),
		in.GetRules())

	return mtproto.BoolTrue, nil
}
