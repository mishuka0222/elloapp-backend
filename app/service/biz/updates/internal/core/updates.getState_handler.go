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
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/updates/updates"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
	"time"
)

// UpdatesGetState
// updates.getState auth_key_id:long user_id:long = updates.State;
func (c *UpdatesCore) UpdatesGetState(in *updates.TLUpdatesGetState) (*mtproto.Updates_State, error) {
	pts := int32(c.svcCtx.Dao.IDGenClient2.CurrentPtsId(c.ctx, in.UserId))
	if pts == 0 {
		pts = int32(c.svcCtx.Dao.IDGenClient2.NextPtsId(c.ctx, in.UserId))
	}

	seq := c.svcCtx.Dao.IDGenClient2.CurrentSeqId(c.ctx, in.AuthKeyId)
	if seq == 0 {
		seq = -1
	}
	return mtproto.MakeTLUpdatesState(&mtproto.Updates_State{
		Pts:         pts,
		Qts:         0,
		Seq:         seq,
		Date:        int32(time.Now().Unix()), // TODO(@benqi): do.Date2???
		UnreadCount: 0,
	}).To_Updates_State(), nil
}
