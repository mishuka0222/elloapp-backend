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

// IdgenGetNextSeqId
// idgen.getNextSeqId key:string = Int64;
func (c *IdgenCore) IdgenGetNextSeqId(in *idgen.TLIdgenGetNextSeqId) (*mtproto.Int64, error) {
	id, err := c.svcCtx.Dao.Store.Incrby(in.GetKey(), 1)
	if err != nil {
		c.Logger.Errorf("dgen.getNextSeqId(%s) error: %v", in.GetKey(), err)
		return nil, err
	}

	return &mtproto.Int64{
		V: id,
	}, nil
}
