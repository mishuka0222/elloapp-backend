package core

import (
	"encoding/json"
)

type GetHistoryCounterResp struct {
	Count int32 `json:"count"`
}

// GetHistoryCounter
// count unread messages in feeds
// req: nil, resp: GetHistoryCounterResp { count: int32 }
func (c *FeedCore) GetHistoryCounter(_ json.RawMessage) (*GetHistoryCounterResp, error) {

	/*l, err := c.svcCtx.Dao.SelectFeedList(c.ctx, c.MD.UserId)
	if err != nil {
		return nil, err
	}
	var count int32
	for _, it := range l {
		mCnt, err := c.svcCtx.Dao.MessageClient.MessageGetHistoryMessagesCount(c.ctx, &message.TLMessageGetHistoryMessagesCount{
			Constructor: message.CRC32_message_getHistoryMessagesCount,
			UserId:      c.MD.UserId,
			PeerType:    3,
			PeerId:      it,
		})
		if err != nil {
			return nil, err
		}
		if mCnt.V > 0 {
			count += mCnt.V
		}
	}*/

	readInbox, err := c.svcCtx.Dao.UserFeedsDAO.SelectUnreadCountList(c.ctx, c.MD.UserId)
	if err != nil {
		return nil, err
	}
	var count int32
	for _, it := range readInbox {
		if it.Unread > 0 {
			count += it.Unread
		}
	}
	return &GetHistoryCounterResp{count}, nil
}
