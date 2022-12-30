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

// UserGetGlobalPrivacySettings
// user.getGlobalPrivacySettings user_id:int = GlobalPrivacySettings;
func (c *UserCore) UserGetGlobalPrivacySettings(in *user.TLUserGetGlobalPrivacySettings) (*mtproto.GlobalPrivacySettings, error) {
	var (
		archiveAndMuteNewNoncontactPeers = false
	)

	do, _ := c.svcCtx.Dao.UserGlobalPrivacySettingsDAO.Select(c.ctx, in.UserId)
	if do != nil {
		archiveAndMuteNewNoncontactPeers = do.ArchiveAndMuteNewNoncontactPeers
	}

	return mtproto.MakeTLGlobalPrivacySettings(&mtproto.GlobalPrivacySettings{
		ArchiveAndMuteNewNoncontactPeers: mtproto.ToBool(archiveAndMuteNewNoncontactPeers),
	}).To_GlobalPrivacySettings(), nil
}
