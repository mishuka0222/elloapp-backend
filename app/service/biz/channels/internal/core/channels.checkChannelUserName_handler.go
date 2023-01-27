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

	switch true {
	case err == sqlx.ErrNotFound:
		break
	case id != 0:
		res = &channels.CheckChannelUserNameResp{Status: true}
		return
	case err != nil:
		return
	default:
	}

	status, err = c.svcCtx.Dao.ChannelsDAO.CheckUsername(c.ctx, in.UserName)
	switch true {
	case err == sqlx.ErrNotFound:
		res = &channels.CheckChannelUserNameResp{Status: false}
		return
	case err != nil:
		return
	default:
	}

	res = &channels.CheckChannelUserNameResp{Status: status}

	return
}
