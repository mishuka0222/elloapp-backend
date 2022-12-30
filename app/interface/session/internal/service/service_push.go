package service

import (
	"context"

	sessionpb "gitlab.com/merehead/elloapp/backend/elloapp_backend/app/interface/session/session"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"

	"github.com/zeromicro/go-zero/core/logx"
)

// SessionPushUpdatesData
// RPCPushClient is the client API for RPCPush service.
func (s *Service) SessionPushUpdatesData(ctx context.Context, r *sessionpb.TLSessionPushUpdatesData) (*mtproto.Bool, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var (
		sessList *authSessions
		ok       bool
	)
	if sessList, ok = s.sessionsManager[r.AuthKeyId]; !ok {
		logx.WithContext(ctx).Errorf("not found authKeyId")
		return mtproto.ToBool(false), nil
	}

	sessList.syncDataArrived(r.Notification, &messageData{obj: r.Updates})
	return mtproto.ToBool(true), nil
}

// SessionPushSessionUpdatesData
// RPCPushClient is the client API for RPCPush service.
func (s *Service) SessionPushSessionUpdatesData(ctx context.Context, r *sessionpb.TLSessionPushSessionUpdatesData) (*mtproto.Bool, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var (
		sessList *authSessions
		ok       bool
	)
	if sessList, ok = s.sessionsManager[r.AuthKeyId]; !ok {
		logx.WithContext(ctx).Errorf("not found authKeyId")
		return mtproto.ToBool(false), nil
	}

	sessList.syncSessionDataArrived(r.SessionId, &messageData{obj: r.Updates})
	return mtproto.ToBool(true), nil
}

func (s *Service) SessionPushRpcResultData(ctx context.Context, r *sessionpb.TLSessionPushRpcResultData) (*mtproto.Bool, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var (
		sessList *authSessions
		ok       bool
	)
	if sessList, ok = s.sessionsManager[r.AuthKeyId]; !ok {
		logx.WithContext(ctx).Errorf("not found authKeyId")
		return mtproto.ToBool(false), nil
	}

	sessList.syncRpcResultDataArrived(r.SessionId, r.ClientReqMsgId, r.RpcResultData)
	return mtproto.ToBool(true), nil
}
