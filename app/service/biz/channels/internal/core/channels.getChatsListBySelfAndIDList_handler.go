package core

import (
	"github.com/zeromicro/go-zero/core/logx"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"

	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/channels/channels"
)

func (c *ChannelsCore) GetChatsListBySelfAndIDList(in *channels.GetChatsListBySelfAndIDListReq) (res *channels.GetChatsListBySelfAndIDListResp, err error) {
	res = &channels.GetChatsListBySelfAndIDListResp{Chats: make([]*mtproto.Chat, 0)}

	if len(in.IdList) == 0 {
		return
	}

	res.Chats = make([]*mtproto.Chat, 0, len(in.IdList))

	var (
		channelData *channels.ChannelData
		r           *channels.ToChatResp
	)

	for _, id := range in.IdList {
		channelData, err = c.GetChannelDataById(&channels.ChannelDataByIdReq{ChannelId: id})
		if err != nil {
			logx.Error("getChatListBySelfIDList - not find chat_id: ", id)
			chatEmpty := (&mtproto.Chat{Id: id}).To_ChatEmpty()
			res.Chats = append(res.Chats, chatEmpty.To_Chat())
		} else {
			r, _ = c.ToChat(&channels.ToChatReq{Channel: channelData, SelfUserId: in.SelfUserId})
			res.Chats = append(res.Chats, r.Chat)
		}
	}

	return
}
