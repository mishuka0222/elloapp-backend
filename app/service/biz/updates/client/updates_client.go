package updates_client

import (
	"context"

	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/updates/updates"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"

	"github.com/zeromicro/go-zero/zrpc"
)

var _ *mtproto.Bool

type UpdatesClient interface {
	UpdatesGetState(ctx context.Context, in *updates.TLUpdatesGetState) (*mtproto.Updates_State, error)
	UpdatesGetDifferenceV2(ctx context.Context, in *updates.TLUpdatesGetDifferenceV2) (*updates.Difference, error)
	UpdatesGetChannelDifferenceV2(ctx context.Context, in *updates.TLUpdatesGetChannelDifferenceV2) (*updates.ChannelDifference, error)
}

type defaultUpdatesClient struct {
	cli zrpc.Client
}

func NewUpdatesClient(cli zrpc.Client) UpdatesClient {
	return &defaultUpdatesClient{
		cli: cli,
	}
}

// UpdatesGetState
// updates.getState auth_key_id:long user_id:long = updates.State;
func (m *defaultUpdatesClient) UpdatesGetState(ctx context.Context, in *updates.TLUpdatesGetState) (*mtproto.Updates_State, error) {
	client := updates.NewRPCUpdatesClient(m.cli.Conn())
	return client.UpdatesGetState(ctx, in)
}

// UpdatesGetDifferenceV2
// updates.getDifferenceV2 flags:# auth_key_id:long user_id:long pts:int pts_total_limit:flags.0?int date:long = Difference;
func (m *defaultUpdatesClient) UpdatesGetDifferenceV2(ctx context.Context, in *updates.TLUpdatesGetDifferenceV2) (*updates.Difference, error) {
	client := updates.NewRPCUpdatesClient(m.cli.Conn())
	return client.UpdatesGetDifferenceV2(ctx, in)
}

// UpdatesGetChannelDifferenceV2
// updates.getChannelDifferenceV2 auth_key_id:long user_id:long channel_id:long pts:int limit:int = ChannelDifference;
func (m *defaultUpdatesClient) UpdatesGetChannelDifferenceV2(ctx context.Context, in *updates.TLUpdatesGetChannelDifferenceV2) (*updates.ChannelDifference, error) {
	client := updates.NewRPCUpdatesClient(m.cli.Conn())
	return client.UpdatesGetChannelDifferenceV2(ctx, in)
}
