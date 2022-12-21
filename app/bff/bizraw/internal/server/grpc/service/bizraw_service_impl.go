package service

import (
	"context"
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/bff/bizraw/internal/core"
)

func (s Service) BizInvokeBizDataRaw(ctx context.Context, request *mtproto.TLBizInvokeBizDataRaw) (*mtproto.BizDataRaw, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("bizraw.bizInvokeBizDataRaw - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())
	r, err := c.BizInvokeBizDataRaw(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("bizraw.bizInvokeBizDataRaw - reply: %s", r.DebugString())
	return r, err
}
