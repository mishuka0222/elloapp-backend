package core

import (
	"github.com/zeromicro/go-zero/core/mr"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/user/internal/dal/dataobject"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/user/user"
)

// UserGetLastSeens
// user.getLastSeens id:Vector<long> = Vector<LastSeenData>;
func (c *UserCore) UserGetLastSeens(in *user.TLUserGetLastSeens) (*user.Vector_LastSeenData, error) {
	type idxId struct {
		idx int
		id  int64
	}

	doList := make([]*dataobject.UserPresencesDO, len(in.Id))
	mr.ForEach(
		func(source chan<- interface{}) {
			for i, v := range in.GetId() {
				source <- idxId{i, v}
			}
		},
		func(item interface{}) {
			id := item.(idxId)
			do, _ := c.svcCtx.Dao.GetLastSeenAt(c.ctx, id.id)
			if do != nil {
				doList[id.idx] = do
			}
		})

	rValues := &user.Vector_LastSeenData{
		Datas: make([]*user.LastSeenData, 0, len(doList)),
	}

	for i := 0; i < len(doList); i++ {
		if doList[i] == nil {
			continue
		}

		rValues.Datas = append(rValues.Datas, user.MakeTLLastSeenData(&user.LastSeenData{
			UserId:     doList[i].UserId,
			LastSeenAt: doList[i].LastSeenAt,
			Expires:    doList[i].Expires,
		}).To_LastSeenData())
	}

	return rValues, nil
}
