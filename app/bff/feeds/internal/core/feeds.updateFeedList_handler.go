package core

import (
	"encoding/json"
)

// UpdateFeedList
// update user_feeds list
// req chats: []int64, resp: nil
func (c *FeedCore) UpdateFeedList(in json.RawMessage) (interface{}, error) {
	var chats []int64
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
