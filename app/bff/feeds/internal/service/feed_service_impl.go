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
// req chats: [ { chat_id: int64, peer_type: int32 } ], resp: nil
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
// req: [ { chat_id: int64, max_id: int32, peer_type: int32 } ]
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

// GetHistory
// If you need read previous messages mast be above zero
// req: GetHistoryReq { before: int32, limit: int32 }
/* resp: GetHistoryResp
{
	messages: [{

    }],
    chats: [{

    }],
    users: [{

    }],
    chat_id_info: [{
        chat_id:          int64,
        inexact:          bool,
        count:            int32,
        next_rate:        int32,
        offset_id_offset: int32,
        pts:              int32
    }]
}
*/
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
// req: nil, resp: { count: int32 }
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
