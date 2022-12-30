package core

import (
	"encoding/json"
	"github.com/teamgram/teamgram-server/app/service/biz/feeds/feeds"
)

// GetFeedList
// return all chats with bool for user
// req: nil, resp: { chat_id: int64, photo_id: int64, title: string, status: bool, peer_type: int32 }
func (c *FeedCore) GetFeedList(_ json.RawMessage) ([]*feeds.UserChatDO, error) {
	data, err := c.svcCtx.Dao.FeedsClient.FeedsGetFeedList(c.ctx, &feeds.GetFeedListReq{UserId: c.MD.GetUserId()})
	if err != nil {
		return nil, err
	}
	return data.GetData(), nil
}
