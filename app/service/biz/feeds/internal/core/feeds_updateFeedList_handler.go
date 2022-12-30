package core

import (
	"github.com/teamgram/teamgram-server/app/service/biz/feeds/feeds"
)

// FeedsUpdateFeedList
// update user_feeds list
// req chats: [ { chat_id: int64, peer_type: int32 } ], resp: nil
func (c *FeedsCore) FeedsUpdateFeedList(in *feeds.UpdateFeedListReq) (*feeds.UpdateFeedListResp, error) {
	if _, err := c.svcCtx.Dao.UserFeedsDAO.DeleteFromListElseValue(c.ctx, in.GetUserId(), in.Data); err != nil {
		return nil, err
	}
	if _, err := c.svcCtx.Dao.UserFeedsDAO.InsertList(c.ctx, in.GetUserId(), in.Data); err != nil {
		return nil, err
	}
	return &feeds.UpdateFeedListResp{}, nil
}
