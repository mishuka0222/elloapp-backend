package feeds_client

import (
	"context"
	"github.com/zeromicro/go-zero/zrpc"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/feeds/feeds"
)

type FeedsClient interface {
	FeedsGetFeedList(ctx context.Context, in *feeds.GetFeedListReq) (*feeds.GetFeedListResp, error)
	FeedsUpdateFeedList(ctx context.Context, in *feeds.UpdateFeedListReq) (*feeds.UpdateFeedListResp, error)
	FeedsGetHistoryCounter(ctx context.Context, in *feeds.GetHistoryCounterReq) (*feeds.GetHistoryCounterResp, error)
	FeedsGetFeedsOffsetList(ctx context.Context, in *feeds.GetFeedsOffsetListReq) (*feeds.GetFeedsOffsetListResp, error)
}

type defaultFeedsClient struct {
	cli zrpc.Client
}

func NewFeedsClient(cli zrpc.Client) FeedsClient {
	return &defaultFeedsClient{
		cli: cli,
	}
}

func (c *defaultFeedsClient) FeedsGetFeedList(ctx context.Context, in *feeds.GetFeedListReq) (*feeds.GetFeedListResp, error) {
	client := feeds.NewRPCFeedsClient(c.cli.Conn())
	return client.FeedsGetFeedList(ctx, in)
}

func (c *defaultFeedsClient) FeedsUpdateFeedList(ctx context.Context, in *feeds.UpdateFeedListReq) (*feeds.UpdateFeedListResp, error) {
	client := feeds.NewRPCFeedsClient(c.cli.Conn())
	return client.FeedsUpdateFeedList(ctx, in)
}

func (c *defaultFeedsClient) FeedsGetHistoryCounter(ctx context.Context, in *feeds.GetHistoryCounterReq) (*feeds.GetHistoryCounterResp, error) {
	client := feeds.NewRPCFeedsClient(c.cli.Conn())
	return client.FeedsGetHistoryCounter(ctx, in)
}

func (c *defaultFeedsClient) FeedsGetFeedsOffsetList(ctx context.Context, in *feeds.GetFeedsOffsetListReq) (*feeds.GetFeedsOffsetListResp, error) {
	client := feeds.NewRPCFeedsClient(c.cli.Conn())
	return client.FeedsGetFeedsOffsetList(ctx, in)
}
