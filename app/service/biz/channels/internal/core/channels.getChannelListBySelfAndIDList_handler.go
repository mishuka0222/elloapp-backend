package core

import (
	"github.com/golang/glog"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"

	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/channels/channels"
)

// channel_util
func (c *ChannelsCore) GetChannelListBySelfAndIDList(in *channels.GetChannelListBySelfAndIDListReq) (res *channels.GetChannelListBySelfAndIDListResp, err error) {
	res = &channels.GetChannelListBySelfAndIDListResp{Chats: make([]*mtproto.Chat, 0)}

	if len(in.IdList) == 0 {
		return
	}

	res.Chats = make([]*mtproto.Chat, 0, len(in.IdList))

	var (
		r *channels.ChannelCoreData
	)

	for _, id := range in.IdList {
		r, err = c.NewChannelCoreById(&channels.ChannelCoreByIdReq{ChannelId: id})
		if err != nil {
			glog.Error("getChatListBySelfIDList - not find chat_id: ", id)
			chatEmpty := (&mtproto.Chat{Id: id}).To_ChatEmpty()
			res.Chats = append(res.Chats, chatEmpty.To_Chat())
		} else {
			res.Chats = append(res.Chats, r.ToChannel(in.SelfUserId))
		}
	}

	return
}
