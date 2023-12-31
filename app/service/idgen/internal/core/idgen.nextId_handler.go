package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/idgen/idgen"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

// IdgenNextId
// idgen.nextId = Int64;
func (c *IdgenCore) IdgenNextId(in *idgen.TLIdgenNextId) (*mtproto.Int64, error) {
	_ = in

	return &mtproto.Int64{
		V: c.svcCtx.Dao.Node.Generate().Int64(),
	}, nil
}
