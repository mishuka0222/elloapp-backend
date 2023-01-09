package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/user/user"
)

// UserGetContactIdList
// user.getContactIdList user_id:long = Vector<long>;
func (c *UserCore) UserGetContactIdList(in *user.TLUserGetContactIdList) (*user.Vector_Long, error) {
	rValList := &user.Vector_Long{
		Datas: []int64{},
	}

	cacheUserData := c.svcCtx.Dao.GetCacheUserData(c.ctx, in.GetUserId())
	if len(cacheUserData.GetContactIdList()) > 0 {
		rValList.Datas = cacheUserData.GetContactIdList()
	}

	return rValList, nil
}
