package core

import (
	"github.com/zeromicro/go-zero/core/mr"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/user/user"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/pkg2/container2"
)

// UserGetMutableUsers
// user.getMutableUsers id:Vector<int> = Vector<ImmutableUser>;
func (c *UserCore) UserGetMutableUsers(in *user.TLUserGetMutableUsers) (*user.Vector_ImmutableUser, error) {
	id := make([]int64, 0, len(in.Id)+len(in.To))
	for _, v := range in.Id {
		if ok, _ := container2.Contains(v, id); !ok {
			id = append(id, v)
		}
	}
	for _, v := range in.To {
		if ok, _ := container2.Contains(v, id); !ok {
			id = append(id, v)
		}
	}

	vUser := &user.Vector_ImmutableUser{
		Datas: make([]*user.ImmutableUser, 0, len(id)),
	}

	if len(id) == 0 {
		return vUser, nil
	} else if len(id) == 1 {
		immutableUser, _ := c.svcCtx.Dao.GetImmutableUser(c.ctx, id[0], false)
		if immutableUser != nil {
			vUser.Datas = append(vUser.Datas, immutableUser)
		}

		return vUser, nil
	}

	mUsers := make([]*user.ImmutableUser, len(id))
	mr.ForEach(
		func(source chan<- interface{}) {
			for idx := 0; idx < len(id); idx++ {
				source <- idx
			}
		},
		func(item interface{}) {
			var (
				idx = item.(int)
				err error
			)

			if ok, _ := container2.Contains(id[idx], in.To); ok {
				mUsers[idx], err = c.svcCtx.Dao.GetImmutableUser(c.ctx, id[idx], true, in.Id...)
				if err != nil {
					c.Logger.Errorf("getImmutableUser - error: %v", err)
				}
			} else {
				if len(in.To) == 0 {
					mUsers[idx], err = c.svcCtx.Dao.GetImmutableUser(c.ctx, id[idx], true, in.Id...)
					if err != nil {
						c.Logger.Errorf("getImmutableUser - error: %v", err)
					}
				} else {
					mUsers[idx], err = c.svcCtx.Dao.GetImmutableUser(c.ctx, id[idx], true, in.To...)
					if err != nil {
						c.Logger.Errorf("getImmutableUser - error: %v", err)
					}
				}
			}
		})

	for _, v := range mUsers {
		if v != nil {
			vUser.Datas = append(vUser.Datas, v)
		}
	}

	return vUser, nil
}
