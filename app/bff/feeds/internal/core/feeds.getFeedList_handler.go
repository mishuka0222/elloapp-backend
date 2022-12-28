package core

import (
	"encoding/json"
	"github.com/teamgram/teamgram-server/app/bff/feeds/internal/dao/dataobject"
)

// GetFeedList
// return all chats with bool for user
// req: nil, resp: { chat_id: int64, photo_id: int64, title: string, status: bool, peer_type: int32 }
func (c *FeedCore) GetFeedList(_ json.RawMessage) ([]dataobject.UserChatDO, error) {

	feedsData, err := c.svcCtx.Dao.UserFeedsDAO.SelectFeedList(c.ctx, c.MD.UserId)
	if err != nil {
		return nil, err
	}
	feedsMap := make(map[int64]int32, len(feedsData))
	for _, i := range feedsData {
		feedsMap[i.ChatID] = i.PeerType
	}

	chatsData, err := c.svcCtx.Dao.UserFeedsDAO.SelectChatList(c.ctx, c.MD.UserId)
	if err != nil {
		return nil, err
	}

	for i := range chatsData {
		if peerType, ok := feedsMap[chatsData[i].ChatID]; ok {
			chatsData[i].Status = true
			chatsData[i].PeerType = peerType
		}
	}

	return chatsData, nil
}
