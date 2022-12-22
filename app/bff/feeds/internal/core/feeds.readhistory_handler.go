package core

import (
	"encoding/json"
	"github.com/teamgram/proto/mtproto"
)

// for user req: { user_id: int64 }
func (c *FeedCore) ReadHistory(in json.RawMessage) (*mtproto.BizDataRaw, error) {
	// todo: add your logic here and delete this line

	return &mtproto.BizDataRaw{}, nil
}
