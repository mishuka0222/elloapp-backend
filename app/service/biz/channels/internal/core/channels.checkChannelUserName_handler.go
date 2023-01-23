package core

import (
	"errors"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/channels/channels"
)

func (c *ChannelsCore) CheckChannelUserName(in *channels.CheckChannelUserNameReq) (res *channels.CheckChannelUserNameResp, err error) {
	// todo: add your logic here and delete this line
	return nil, errors.New("Unimplemented")

	//do, err = c.svcCtx.Dao.UsersDAO.SelectByUsername(in.UserName)
	//if err != nil {
	//	return
	//}
	//
	//params := map[string]interface{}{
	//	"channel_id": in.Id,
	//	"user_id": do.Id,
	//}
	//status, err = c.svcCtx.Dao.CommonDAO.CheckExists("channel_participants", params)

	res = &channels.CheckChannelUserNameResp{}

	return
}
