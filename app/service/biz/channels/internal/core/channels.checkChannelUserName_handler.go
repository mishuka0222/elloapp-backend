package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/channels/channels"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/pkg2/stores/sqlx"
)

func (c *ChannelsCore) CheckChannelUserName(in *channels.CheckChannelUserNameReq) (res *channels.CheckChannelUserNameResp, err error) {
	var (
		id     int64
		status bool
	)
	id, err = c.svcCtx.Dao.UsersDAO.SelectByUsername(c.ctx, in.UserName)
	if err == sqlx.ErrNotFound {
		err = nil
	} else if err != nil {
		return
	} else if id != 0 {
		res = &channels.CheckChannelUserNameResp{Status: true}
		return
	}

	status, err = c.svcCtx.Dao.ChannelsDAO.CheckUsername(c.ctx, in.UserName)
	if err == sqlx.ErrNotFound {
		err = nil
		res = &channels.CheckChannelUserNameResp{Status: false}
		return
	} else if err != nil {
		return
	}

	res = &channels.CheckChannelUserNameResp{Status: status}

	return
}
