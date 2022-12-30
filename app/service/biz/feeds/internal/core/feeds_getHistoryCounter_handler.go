package core

import (
	"github.com/teamgram/teamgram-server/app/service/biz/feeds/feeds"
)

// FeedsGetHistoryCounter
// count unread messages in feeds
// req: nil, resp: { count: int32 }
func (c *FeedsCore) FeedsGetHistoryCounter(in *feeds.GetHistoryCounterReq) (*feeds.GetHistoryCounterResp, error) {
	readInbox, err := c.svcCtx.Dao.UserFeedsDAO.SelectUnreadCountList(c.ctx, in.GetUserId())
	if err != nil {
		return nil, err
	}
	var count int32
	for _, it := range readInbox {
		if it.Unread > 0 {
			count += it.Unread
		}
	}

	return &feeds.GetHistoryCounterResp{Count: count}, nil
}
