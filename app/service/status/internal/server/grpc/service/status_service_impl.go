package service

import (
	"context"

	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/status/internal/core"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/status/status"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
)

// StatusSetSessionOnline
// status.setSessionOnline user_id:long session:SessionEntry = Bool;
func (s *Service) StatusSetSessionOnline(ctx context.Context, request *status.TLStatusSetSessionOnline) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("status.setSessionOnline - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.StatusSetSessionOnline(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("status.setSessionOnline - reply: %s", r.DebugString())
	return r, err
}

// StatusSetSessionOffline
// status.setSessionOffline user_id:long auth_key_id:long = Bool;
func (s *Service) StatusSetSessionOffline(ctx context.Context, request *status.TLStatusSetSessionOffline) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("status.setSessionOffline - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.StatusSetSessionOffline(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("status.setSessionOffline - reply: %s", r.DebugString())
	return r, err
}

// StatusGetUserOnlineSessions
// status.getUserOnlineSessions user_id:long = UserSessionEntryList;
func (s *Service) StatusGetUserOnlineSessions(ctx context.Context, request *status.TLStatusGetUserOnlineSessions) (*status.UserSessionEntryList, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("status.getUserOnlineSessions - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.StatusGetUserOnlineSessions(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("status.getUserOnlineSessions - reply: %s", r.DebugString())
	return r, err
}

// StatusGetUsersOnlineSessionsList
// status.getUsersOnlineSessionsList users:Vector<long> = Vector<UserSessionEntryList>;
func (s *Service) StatusGetUsersOnlineSessionsList(ctx context.Context, request *status.TLStatusGetUsersOnlineSessionsList) (*status.Vector_UserSessionEntryList, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("status.getUsersOnlineSessionsList - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.StatusGetUsersOnlineSessionsList(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("status.getUsersOnlineSessionsList - reply: %s", r.DebugString())
	return r, err
}
