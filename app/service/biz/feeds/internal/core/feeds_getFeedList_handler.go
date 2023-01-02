package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/feeds/feeds"
)

// FeedsGetFeedList
// return all chats with bool for user
// req: nil, resp: { chat_id: int64, photo_id: int64, title: string, status: bool, peer_type: int32 }
func (c *FeedsCore) FeedsGetFeedList(in *feeds.GetFeedListReq) (*feeds.GetFeedListResp, error) {
	feedsData, err := c.svcCtx.Dao.UserFeedsDAO.SelectFeedList(c.ctx, in.GetUserId())
	if err != nil {
		return nil, err
	}
	feedsMap := make(map[int64]int32, len(feedsData))
	for _, i := range feedsData {
		feedsMap[i.ChatID] = i.PeerType
	}

	chatsData, err := c.svcCtx.Dao.UserFeedsDAO.SelectChatList(c.ctx, in.GetUserId())
	if err != nil {
		return nil, err
	}

	for i := range chatsData {
		if peerType, ok := feedsMap[chatsData[i].GetChatId()]; ok {
			chatsData[i].Status = true
			chatsData[i].PeerType = peerType
		}
	}

	return &feeds.GetFeedListResp{Data: chatsData}, nil
}
