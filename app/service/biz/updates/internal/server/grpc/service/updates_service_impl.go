package service

import (
	"context"

	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/updates/internal/core"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/updates/updates"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

// UpdatesGetState
// updates.getState auth_key_id:long user_id:long = updates.State;
func (s *Service) UpdatesGetState(ctx context.Context, request *updates.TLUpdatesGetState) (*mtproto.Updates_State, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("updates.getState - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.UpdatesGetState(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("updates.getState - reply: %s", r.DebugString())
	return r, err
}

// UpdatesGetDifferenceV2
// updates.getDifferenceV2 flags:# auth_key_id:long user_id:long pts:int pts_total_limit:flags.0?int date:long = Difference;
func (s *Service) UpdatesGetDifferenceV2(ctx context.Context, request *updates.TLUpdatesGetDifferenceV2) (*updates.Difference, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("updates.getDifferenceV2 - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.UpdatesGetDifferenceV2(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("updates.getDifferenceV2 - reply: %s", r.DebugString())
	return r, err
}

// UpdatesGetChannelDifferenceV2
// updates.getChannelDifferenceV2 auth_key_id:long user_id:long channel_id:long pts:int limit:int = ChannelDifference;
func (s *Service) UpdatesGetChannelDifferenceV2(ctx context.Context, request *updates.TLUpdatesGetChannelDifferenceV2) (*updates.ChannelDifference, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("updates.getChannelDifferenceV2 - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.UpdatesGetChannelDifferenceV2(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("updates.getChannelDifferenceV2 - reply: %s", r.DebugString())
	return r, err
}
