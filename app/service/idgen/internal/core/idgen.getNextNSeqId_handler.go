package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/idgen/idgen"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
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
