package service

import (
	"context"
	"encoding/json"
	"github.com/teamgram/teamgram-server/app/bff/feeds/internal/core"
)

// GetFeedList
// return all chats with bool for user
// req: nil, resp: { chat_id: int64, photo_id: int64, title: string, status: bool, peer_type: int32 }
func (s *Service) GetFeedList(ctx context.Context, in json.RawMessage) (interface{}, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("feeds.GetFeedList - metadata: %s, request: %s", c.MD.DebugString(), string(in))

	r, err := c.GetFeedList(in)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("feeds.GetFeedList - reply: %+v", r)
	return r, err
}

// UpdateFeedList
// update user_feeds list
// req chats: []int64, resp: nil
func (s *Service) UpdateFeedList(ctx context.Context, in json.RawMessage) (interface{}, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("feeds.UpdateFeedList - metadata: %s, request: %s", c.MD.DebugString(), string(in))

	r, err := c.UpdateFeedList(in)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("feeds.UpdateFeedList - reply: %+v", r)
	return r, err
}

// ReadHistory
// need for updating unread messages count
// req: [] ChatMaxID { chat_id: int64, max_id: int32, peer_type: int32 }
// resp: ReadHistoryResp { pts: int32, pts_count: int32, chat_id: int64, peer_type: int32 }
func (s *Service) ReadHistory(ctx context.Context, in json.RawMessage) (interface{}, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("feeds.ReadHistory - metadata: %s, request: %s", c.MD.DebugString(), string(in))

	r, err := c.ReadHistory(in)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("feeds.ReadHistory - reply: %+v", r)
	return r, err
}

// for user req: { uread_state: int64 }
func (s *Service) GetHistory(ctx context.Context, in json.RawMessage) (interface{}, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("feeds.ReadHistory - metadata: %s, request: %s", c.MD.DebugString(), string(in))

	r, err := c.GetHistory(in)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("feeds.ReadHistory - reply: %+v", r)
	return r, err
}

// GetHistoryCounter
// count unread messages in feeds
// req: nil, resp: GetHistoryCounterResp { count: int32 }
func (s *Service) GetHistoryCounter(ctx context.Context, in json.RawMessage) (interface{}, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("feeds.GetHistoryCounter - metadata: %s, request: %s", c.MD.DebugString(), string(in))

	r, err := c.GetHistoryCounter(in)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("feeds.GetHistoryCounter - reply: %+v", r)
	return r, err
}
