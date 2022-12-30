package service

import (
	"context"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/feeds/feeds"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/feeds/internal/core"
)

// FeedsGetFeedList
func (s *Service) FeedsGetFeedList(ctx context.Context, in *feeds.GetFeedListReq) (*feeds.GetFeedListResp, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("feeds.GetFeedList - metadata: %s, request: %s", c.MD.DebugString(), in.String())

	r, err := c.FeedsGetFeedList(in)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("feeds.GetFeedList - reply: %+v", r)
	return r, err
}

// FeedsGetFeedsOffsetList
func (s *Service) FeedsGetFeedsOffsetList(ctx context.Context, in *feeds.GetFeedsOffsetListReq) (*feeds.GetFeedsOffsetListResp, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("feeds.GetFeedsOffsetList - metadata: %s, request: %s", c.MD.DebugString(), in.String())

	r, err := c.FeedsGetFeedsOffsetList(in)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("feeds.GetFeedsOffsetList - reply: %+v", r)
	return r, err
}

// FeedsGetHistoryCounter
func (s *Service) FeedsGetHistoryCounter(ctx context.Context, in *feeds.GetHistoryCounterReq) (*feeds.GetHistoryCounterResp, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("feeds.GetHistoryCounter - metadata: %s, request: %s", c.MD.DebugString(), in.String())

	r, err := c.FeedsGetHistoryCounter(in)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("feeds.GetHistoryCounter - reply: %+v", r)
	return r, err
}

// FeedsUpdateFeedList
func (s *Service) FeedsUpdateFeedList(ctx context.Context, in *feeds.UpdateFeedListReq) (*feeds.UpdateFeedListResp, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("feeds.UpdateFeedList - metadata: %s, request: %s", c.MD.DebugString(), in.String())

	r, err := c.FeedsUpdateFeedList(in)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("feeds.UpdateFeedList - reply: %+v", r)
	return r, err
}
