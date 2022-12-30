package core

import (
	"encoding/json"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/feeds/feeds"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto/rpc/metadata"
)

// UpdateFeedList
// update user_feeds list
// req chats: [ { chat_id: int64, peer_type: int32 } ], resp: nil
func (c *FeedCore) UpdateFeedList(in json.RawMessage) (interface{}, error) {
	var chats []*feeds.FeedInsertItemDO
	if err := json.Unmarshal(in, &chats); err != nil {
		return nil, err
	}
	ctx, err := metadata.RpcMetadataToOutgoing(c.ctx, &metadata.RpcMetadata{
		UserId: c.MD.UserId,
	})
	if err != nil {
		panic(err)
	}
	return c.svcCtx.Dao.FeedsClient.FeedsUpdateFeedList(ctx,
		&feeds.UpdateFeedListReq{UserId: c.MD.GetUserId(), Data: chats})
}
