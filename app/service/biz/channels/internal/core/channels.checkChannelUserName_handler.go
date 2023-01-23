package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/channels/channels"
)

func (c *ChannelsCore) CheckChannelUserName(in *channels.CheckChannelUserNameReq) (res *channels.CheckChannelUserNameResp, err error) {
	var (
		id     int64
		status bool
	)
	id, err = c.svcCtx.Dao.UsersDAO.CheckExistsUsername(c.ctx, in.UserName)
	if err != nil {
		return
	}

	status, err = c.svcCtx.Dao.ChannelParticipantsDAO.CheckExists(c.ctx, in.ChannelId, id)

	res = &channels.CheckChannelUserNameResp{Status: status}

	return
}
