package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/feeds/feeds"
)

// FeedsGetFeedsOffsetList
// If you need read previous messages mast be above zero
// req: { before: int32, limit: int32 }
// resp: { offset_min: int32, offset_max: int32, peer_id: int32, count: int32 }
func (c *FeedsCore) FeedsGetFeedsOffsetList(in *feeds.GetFeedsOffsetListReq) (*feeds.GetFeedsOffsetListResp, error) {
	feedsList, err := c.svcCtx.Dao.SelectFeedList(c.ctx, in.GetUserId())
	if err != nil {
		return nil, err
	}
	if len(feedsList) == 0 {
		return nil, nil
	}
	chatIdList := make([]int64, len(feedsList))
	for _, it := range feedsList {
		chatIdList = append(chatIdList, it.ChatID)
	}

	var offsetList []*feeds.OffsetItemDo
	if in.GetBefore() == 0 {
		offsetList, err = c.svcCtx.Dao.UserFeedsDAO.SelectOffsetMaxList(c.ctx, in.GetUserId(), chatIdList, in.GetLimit())
		if err != nil {
			return nil, err
		}
	} else {
		offsetList, err = c.svcCtx.Dao.UserFeedsDAO.SelectOffsetMinList(c.ctx, in.GetUserId(), chatIdList, in.GetLimit(), in.GetBefore())
		if err != nil {
			return nil, err
		}
	}
	return &feeds.GetFeedsOffsetListResp{Data: offsetList}, nil
}
