/*
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/user/user"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
)

// UserGetPrivacy
// user.getPrivacy user_id:int key_type:int = Vector<PrivacyRule>;
func (c *UserCore) UserGetPrivacy(in *user.TLUserGetPrivacy) (*user.Vector_PrivacyRule, error) {
	rulesList := c.svcCtx.Dao.GetUserPrivacyRulesListByKeys(c.ctx, in.GetUserId(), in.GetKeyType())
	if len(rulesList) == 0 {
		err := mtproto.ErrPrivacyKeyInvalid
		c.Logger.Errorf("user.getPrivacy - error: %v", err)
		return nil, err
	}

	return &user.Vector_PrivacyRule{
		Datas: rulesList[0].GetRules(),
	}, nil
}
