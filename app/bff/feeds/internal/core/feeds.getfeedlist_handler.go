package core

import (
	"encoding/json"
	"github.com/teamgram/proto/mtproto"
)

type GetListResp struct {
	UserID int32 `json:"user_id"`
}

// return all chats with bool for user { chat_id: int64, peer_type: int32, state: bool } req: { user_id: int64 }
func (c *FeedCore) GetFeedList(in json.RawMessage) (*mtproto.BizDataRaw, error) {
	// todo: add your logic here and delete this line
	var it GetListResp
	err := json.Unmarshal(in, &it)
	if err != nil {
		return nil, err
	}
	c.Logger.Debugf("GetFeedList response: %s", in)

	b, err := json.Marshal([]map[string]interface{}{
		{
			"chat_id":   777063,
			"peer_type": 2,
			"state":     false,
		},
		{
			"chat_id":   14,
			"peer_type": 3,
			"state":     true,
		},
	})
	if err != nil {
		return nil, err
	}

	return &mtproto.BizDataRaw{
		Constructor: mtproto.CRC32_bizDataRaw,
		Data:        b,
	}, nil
}
