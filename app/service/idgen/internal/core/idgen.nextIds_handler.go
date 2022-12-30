package core

import (
	"fmt"

	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/idgen/idgen"
)

const (
	maxNextIdsNum = 100
)

// IdgenNextIds
// idgen.nextIds num:int = Vector<long>;
func (c *IdgenCore) IdgenNextIds(in *idgen.TLIdgenNextIds) (*idgen.Vector_Long, error) {
	if in.GetNum() > maxNextIdsNum || in.GetNum() < 0 {
		c.Logger.Errorf("NextIds num can't be greater than %d or less than 0", maxNextIdsNum)
		return nil, fmt.Errorf("NextIds num: %d error", in.GetNum())
	}

	ids := make([]int64, in.GetNum())
	for i := int32(0); i < in.GetNum(); i++ {
		// TODO: 库里提供ids方法，以减少Lock次数
		ids[i] = c.svcCtx.Node.Generate().Int64()
	}

	return &idgen.Vector_Long{
		Datas: ids,
	}, nil
}
