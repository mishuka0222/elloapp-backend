package core

import (
	"errors"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/channels/channels"

	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

func (c *ChannelsCore) ChannelsEditAdmin(in *mtproto.TLChannelsEditAdmin) (upd *mtproto.Updates, err error) {

	// todo: add your logic here and delete this line
	return nil, errors.New("Unimplemented")
	// return &mtproto.Updates{}, nil

	var (
		peer = mtproto.MakeChannelPeerUtil(in.Channel.GetChannelId())
	)

	c.svcCtx.Dao.ChannelsClient.EditChannelAdmin(c.ctx, &channels.EditChannelAdminReq{
		ChannelId:          peer.PeerId,
		OperatorId:         c.MD.UserId,
		EditChannelAdminId: 0,
		IsAdmin:            false,
	})

	return
}
