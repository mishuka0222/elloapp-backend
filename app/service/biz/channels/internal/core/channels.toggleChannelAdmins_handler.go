package core

import (
	"errors"
	"time"

	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/channels/channels"
)

func (c *ChannelsCore) ToggleChannelAdmins(in *channels.ToggleChannelAdminsReq) (res *channels.Void, err error) {
	// check is creator
	if in.UserId != in.Channel.Channel.CreatorUserId {
		err = errors.New("NO_EDIT_CHAT_PERMISSION")
		return
	}

	var (
		channelAdminsEnabled = in.Channel.Channel.AdminsEnabled
	)

	// Check modified
	if channelAdminsEnabled == in.AdminsEnabled {
		err = errors.New("CHAT_NOT_MODIFIED")
		return
	}

	in.Channel.Channel.AdminsEnabled = in.AdminsEnabled
	in.Channel.Channel.Date = int32(time.Now().Unix())
	in.Channel.Channel.Version += 1

	_, err = c.svcCtx.Dao.ChannelsDAO.UpdateAdminsEnabled(c.ctx, in.Channel.Channel.AdminsEnabled, in.Channel.Channel.Date, in.Channel.Channel.Id)

	res = &channels.Void{}

	return
}
