package core

import (
	"encoding/json"
	"github.com/teamgram/teamgram-server/app/bff/feeds/internal/dao/dataobject"
)

// UpdateFeedList
// update user_feeds list
// req chats: [] dataobject.FeedInsertItemDO { chat_id: int64, peer_type: int32 }, resp: nil
func (c *FeedCore) UpdateFeedList(in json.RawMessage) (interface{}, error) {
	var chats []dataobject.FeedInsertItemDO
	if err := json.Unmarshal(in, &chats); err != nil {
		return nil, err
	}
	if _, err := c.svcCtx.Dao.UserFeedsDAO.DeleteFromListElseValue(c.ctx, c.MD.UserId, chats); err != nil {
		return nil, err
	}
	if _, err := c.svcCtx.Dao.UserFeedsDAO.InsertList(c.ctx, c.MD.UserId, chats); err != nil {
		return nil, err
	}
	return nil, nil
}
