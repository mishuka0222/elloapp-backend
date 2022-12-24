package core

import (
	"encoding/json"
	"github.com/teamgram/teamgram-server/app/bff/feeds/internal/dao/dataobject"
)

// GetFeedList
// return all chats with bool for user [{ title: string, chat_id: int64, peer_type: int32, state: bool }] req: empty
func (c *FeedCore) GetFeedList(_ json.RawMessage) ([]dataobject.UserChatDO, error) {

	feedsData, err := c.svcCtx.Dao.UserFeedsDAO.SelectFeedList(c.ctx, c.MD.UserId)
	if err != nil {
		return nil, err
	}
	feedsMap := make(map[int64]struct{}, len(feedsData))
	for _, i := range feedsData {
		feedsMap[i.ChatID] = struct{}{}
	}

	chatsData, err := c.svcCtx.Dao.UserFeedsDAO.SelectChatList(c.ctx, c.MD.UserId)
	if err != nil {
		return nil, err
	}

	for i := range chatsData {
		if _, ok := feedsMap[chatsData[i].ChatID]; ok {
			chatsData[i].Status = true
		}
	}

	return chatsData, nil
}
