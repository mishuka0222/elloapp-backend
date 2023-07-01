package core

import (
	"encoding/json"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

type ChatMaxID struct {
	ChatID int64 `json:"chat_id"`
	MaxID  int32 `json:"max_id"`
	// chat, channel
	PeerType int32 `json:"peer_type"`
}

type ReadHistoryResp struct {
	Pts      int32 `json:"pts"`
	PtsCount int32 `json:"pts_count"`
	ChatID   int64 `json:"chat_id"`
	PeerType int32 `json:"peer_type"`
}

// ReadHistory
// need for updating unread messages count
// req: [] ChatMaxID { chat_id: int64, max_id: int32, peer_type: int32 }
// resp: [] ReadHistoryResp { pts: int32, pts_count: int32, chat_id: int64, peer_type: int32 }
func (c *FeedCore) ReadHistory(in json.RawMessage) ([]ReadHistoryResp, error) {
	var maxIdList []ChatMaxID
	if err := json.Unmarshal(in, &maxIdList); err != nil {
		return nil, err
	}
	var histList []ReadHistoryResp
	for _, it := range maxIdList {
		switch it.PeerType {
		case 2:
			history, err := c.svcCtx.MessagesCore.MessagesReadHistory(c.ctx, &mtproto.TLMessagesReadHistory{
				Constructor: mtproto.CRC32_messages_readHistory,
				Peer: (&mtproto.InputPeer{Constructor: mtproto.CRC32_inputPeerChat, ChatId: it.ChatID}).
					To_InputPeerChat().To_InputPeer(),
				MaxId: it.MaxID,
			})
			if err != nil {
				return nil, err
			}
			histList = append(histList, ReadHistoryResp{
				Pts:      history.Pts,
				PtsCount: history.PtsCount,
				ChatID:   it.ChatID,
				PeerType: it.PeerType,
			})
		}

	}
	return histList, nil
}
