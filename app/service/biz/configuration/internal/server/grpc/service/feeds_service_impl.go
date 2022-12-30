package service

import (
	"context"
	"github.com/teamgram/teamgram-server/app/service/biz/configuration/configuration"
	"github.com/teamgram/teamgram-server/app/service/biz/configuration/internal/core"
)

// ConfigurationGetCountryList
func (s *Service) ConfigurationGetCountryList(ctx context.Context, in *configuration.GetCountryListReq) (*configuration.GetCountryListResp, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("feeds.GetFeedList - metadata: %s, request: %s", c.MD.DebugString(), in.String())

	r, err := c.ConfigurationGetCountryList(in)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("feeds.GetFeedList - reply: %+v", r)
	return r, err
}
