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
	"strconv"
)

// IdgenSetCurrentSeqId
// idgen.setCurrentSeqId key:string id:long = Bool;
func (c *IdgenCore) IdgenSetCurrentSeqId(in *idgen.TLIdgenSetCurrentSeqId) (*mtproto.Bool, error) {
	err := c.svcCtx.Dao.Store.Set(in.GetKey(), strconv.FormatInt(in.GetId(), 10))
	if err != nil {
		c.Logger.Errorf("idgen.setCurrentSeqId(%s, %d) error: %v", in.GetKey(), in.GetId(), err)
		return nil, err
	}

	return mtproto.BoolTrue, nil
}
