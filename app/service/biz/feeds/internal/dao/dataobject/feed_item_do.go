package dataobject

type FeedItemDO struct {
	UserID   int64 `json:"user_id"`
	ChatID   int64 `json:"chat_id"`
	PeerType int32 `json:"peer_type"`
}
