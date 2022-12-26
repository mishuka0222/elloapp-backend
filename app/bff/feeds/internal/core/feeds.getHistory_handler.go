package core

import (
	"encoding/json"
	"github.com/gogo/protobuf/types"
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/bff/feeds/internal/dao/dataobject"
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

	feedsList, err := c.svcCtx.Dao.SelectFeedList(c.ctx, c.MD.UserId)
	if err != nil {
		return nil, err
	}
	if len(feedsList) == 0 {
		return nil, nil
	}
	chatIdList := make([]int64, len(feedsList))
	for _, it := range feedsList {
		chatIdList = append(chatIdList, it.ChatID)
	}

	var offsetList dataobject.OffsetItemList
	if req.Before == 0 {
		offsetList, err = c.svcCtx.Dao.UserFeedsDAO.SelectOffsetMaxList(c.ctx, c.MD.UserId, chatIdList, limit)
		if err != nil {
			return nil, err
		}
	} else {
		offsetList, err = c.svcCtx.Dao.UserFeedsDAO.SelectOffsetMinList(c.ctx, c.MD.UserId, chatIdList, limit, req.Before)
		if err != nil {
			return nil, err
		}
	}
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
				ChatId:      it.PeerID,
			}).To_InputPeerChat().To_InputPeer(),
			OffsetId: it.OffsetMin,
			Limit:    it.Count,
			MinId:    it.OffsetMin,
			MaxId:    it.OffsetMax,
		}

		history, err := c.svcCtx.MessagesCore.MessagesGetHistory(c.ctx, in)
		if err != nil {
			return nil, err
		}

		resp.Messages = append(resp.Messages, history.Messages...)
		resp.Chats = append(resp.Chats, history.Chats...)
		resp.Users = append(resp.Users, history.Users...)
		resp.ChatIdInfo = append(resp.ChatIdInfo, ChatIdInfo{
			ChatID:         it.PeerID,
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
