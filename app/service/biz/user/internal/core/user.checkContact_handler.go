package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/user/user"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/pkg2/container2"
)

// UserCheckContact
// user.checkContact user_id:long id:long = Bool;
func (c *UserCore) UserCheckContact(in *user.TLUserCheckContact) (*mtproto.Bool, error) {
	cacheUserData := c.svcCtx.Dao.GetCacheUserData(c.ctx, in.GetUserId())
	//_, idList := c.svcCtx.Dao.GetUserContactIdList(c.ctx, in.GetUserId())
	isContact, _ := container2.Contains(in.GetId(), cacheUserData.GetContactIdList())

	return mtproto.ToBool(isContact), nil
}
