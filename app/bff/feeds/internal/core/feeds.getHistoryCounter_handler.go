package core

import (
	"encoding/json"
)

// GetHistoryCounter
// count unread messages in feeds req: empty
func (c *FeedCore) GetHistoryCounter(_ json.RawMessage) (interface{}, error) {

	readOutbox, err := c.svcCtx.Dao.UserFeedsDAO.SelectReadOutboxList(c.ctx, c.MD.UserId)
	if err != nil {
		return nil, err
	}
	var count int64
	for _, it := range readOutbox {
		cnt := it.TopMessage - it.ReadOutboxMaxID
		if cnt > 0 {
			count += cnt
		}
	}
	c.Logger.Errorf("here %d", count)
	return map[string]int64{
		"count": count,
	}, nil
}
