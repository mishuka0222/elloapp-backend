package core

import (
	"encoding/json"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/feeds/feeds"
)

// GetHistoryCounter
// count unread messages in feeds
// req: nil, resp: { count: int32 }
func (c *FeedCore) GetHistoryCounter(_ json.RawMessage) (*feeds.GetHistoryCounterResp, error) {

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
	return c.svcCtx.Dao.FeedsClient.FeedsGetHistoryCounter(c.ctx, &feeds.GetHistoryCounterReq{UserId: c.MD.GetUserId()})
}
