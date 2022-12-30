package core

import (
	"encoding/json"
	"github.com/gogo/protobuf/types"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/feeds/feeds"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
	"sort"
)

type GetHistoryResp struct {
	Messages   []*mtproto.Message `json:"messages,omitempty"`
	Chats      []*mtproto.Chat    `json:"chats,omitempty"`
	Users      []*mtproto.User    `json:"users,omitempty"`
	ChatIdInfo []ChatIdInfo       `json:"chat_id_info,omitempty"`
}

type ChatIdInfo struct {
	ChatID         int64             `json:"chat_id"`
	Inexact        bool              `json:"inexact,omitempty"`
	Count          int32             `json:"count,omitempty"`
	NextRate       *types.Int32Value `json:"next_rate,omitempty"`
	OffsetIdOffset *types.Int32Value `json:"offset_id_offset,omitempty"`
	Pts            int32             `json:"pts,omitempty"`
}

type GetHistoryReq struct {
	Before int32 `json:"before,omitempty"`
	Limit  int32 `json:"limit,omitempty"`
}

// GetHistory
// If you need read previous messages mast be above zero
// req: GetHistoryReq { before: int32, limit: int32 }
// resp: GetHistoryResp { }
func (c *FeedCore) GetHistory(in json.RawMessage) (*GetHistoryResp, error) {
	var req GetHistoryReq
	if err := json.Unmarshal(in, &req); err != nil {
		return nil, err
	}
	var limit = req.Limit
	if limit <= 0 {
		limit = 45
	}
	data, err := c.svcCtx.FeedsClient.FeedsGetFeedsOffsetList(c.ctx,
		&feeds.GetFeedsOffsetListReq{UserId: c.MD.GetUserId(), Limit: limit, Before: req.Before})
	if err != nil {
		return nil, err
	}
	offsetList := data.GetData()
	if len(offsetList) == 0 {
		return nil, nil
	}

	resp := GetHistoryResp{
		Messages:   make([]*mtproto.Message, 0),
		Chats:      make([]*mtproto.Chat, 0),
		Users:      make([]*mtproto.User, 0),
		ChatIdInfo: make([]ChatIdInfo, 0),
	}
	for _, it := range offsetList {
		in := &mtproto.TLMessagesGetHistory{
			Constructor: mtproto.CRC32_messages_getHistory,
			Peer: (&mtproto.InputPeer{
				Constructor: mtproto.CRC32_inputPeerChat,
				ChatId:      it.GetPeerId(),
			}).To_InputPeerChat().To_InputPeer(),
			OffsetId: it.GetOffsetMin(),
			Limit:    it.GetCount(),
			MinId:    it.GetOffsetMin(),
			MaxId:    it.GetOffsetMax(),
		}

		history, err := c.svcCtx.MessagesCore.MessagesGetHistory(c.ctx, in)
		if err != nil {
			return nil, err
		}

		resp.Messages = append(resp.Messages, history.Messages...)
		resp.Chats = append(resp.Chats, history.Chats...)
		resp.Users = append(resp.Users, history.Users...)
		resp.ChatIdInfo = append(resp.ChatIdInfo, ChatIdInfo{
			ChatID:         it.GetPeerId(),
			Inexact:        history.Inexact,
			Count:          history.Count,
			NextRate:       history.NextRate,
			OffsetIdOffset: history.OffsetIdOffset,
			Pts:            history.Pts,
		})
	}

	if len(resp.Messages) > 0 {
		sort.Slice(resp.Messages, func(i, j int) bool {
			return resp.Messages[i].Date < resp.Messages[j].Date
		})
	}

	return &resp, nil
}
