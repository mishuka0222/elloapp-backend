package core

import (
	"encoding/json"
	"github.com/teamgram/proto/mtproto"
)

// send array with { chat_id: int64, peer_type: int32, state: bool }
func (c *FeedCore) UpdateFeedList(in json.RawMessage) (*mtproto.BizDataRaw, error) {
	// todo: add your logic here and delete this line

	return &mtproto.BizDataRaw{}, nil
}
