package service

import (
	"context"
	"encoding/json"
	"github.com/teamgram/teamgram-server/app/bff/feeds/internal/core"
)

// return all chats with bool for user { chat_id: int64, peer_type: int32, state: bool } req: { user_id: int64 }
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

// send array with { chat_id: int64, peer_type: int32, state: bool }
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

// for user req: { user_id: int64 }
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

// for user req: { user_id: int64 }
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
