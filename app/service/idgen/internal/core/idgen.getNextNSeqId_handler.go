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

// IdgenGetNextNSeqId
// idgen.getNextNSeqId key:string n:int = Int64;
func (c *IdgenCore) IdgenGetNextNSeqId(in *idgen.TLIdgenGetNextNSeqId) (*mtproto.Int64, error) {
	id, err := c.svcCtx.Dao.Store.Incrby(in.GetKey(), int64(in.GetN()))
	if err != nil {
		c.Logger.Errorf("dgen.getNextNSeqId(%s, %d) error: %v", in.GetKey(), in.GetN(), err)
		return nil, err
	}

	return &mtproto.Int64{
		V: id,
	}, nil
}
