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
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/idgen/idgen"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
)

// IdgenNextId
// idgen.nextId = Int64;
func (c *IdgenCore) IdgenNextId(in *idgen.TLIdgenNextId) (*mtproto.Int64, error) {
	_ = in

	return &mtproto.Int64{
		V: c.svcCtx.Dao.Node.Generate().Int64(),
	}, nil
}
