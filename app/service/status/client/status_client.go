/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2022 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package status_client

import (
	"context"

	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/status/status"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"

	"github.com/zeromicro/go-zero/zrpc"
)

var _ *mtproto.Bool

type StatusClient interface {
	StatusSetSessionOnline(ctx context.Context, in *status.TLStatusSetSessionOnline) (*mtproto.Bool, error)
	StatusSetSessionOffline(ctx context.Context, in *status.TLStatusSetSessionOffline) (*mtproto.Bool, error)
	StatusGetUserOnlineSessions(ctx context.Context, in *status.TLStatusGetUserOnlineSessions) (*status.UserSessionEntryList, error)
	StatusGetUsersOnlineSessionsList(ctx context.Context, in *status.TLStatusGetUsersOnlineSessionsList) (*status.Vector_UserSessionEntryList, error)
}

type defaultStatusClient struct {
	cli zrpc.Client
}

func NewStatusClient(cli zrpc.Client) StatusClient {
	return &defaultStatusClient{
		cli: cli,
	}
}

// StatusSetSessionOnline
// status.setSessionOnline user_id:long session:SessionEntry = Bool;
func (m *defaultStatusClient) StatusSetSessionOnline(ctx context.Context, in *status.TLStatusSetSessionOnline) (*mtproto.Bool, error) {
	client := status.NewRPCStatusClient(m.cli.Conn())
	return client.StatusSetSessionOnline(ctx, in)
}

// StatusSetSessionOffline
// status.setSessionOffline user_id:long auth_key_id:long = Bool;
func (m *defaultStatusClient) StatusSetSessionOffline(ctx context.Context, in *status.TLStatusSetSessionOffline) (*mtproto.Bool, error) {
	client := status.NewRPCStatusClient(m.cli.Conn())
	return client.StatusSetSessionOffline(ctx, in)
}

// StatusGetUserOnlineSessions
// status.getUserOnlineSessions user_id:long = UserSessionEntryList;
func (m *defaultStatusClient) StatusGetUserOnlineSessions(ctx context.Context, in *status.TLStatusGetUserOnlineSessions) (*status.UserSessionEntryList, error) {
	client := status.NewRPCStatusClient(m.cli.Conn())
	return client.StatusGetUserOnlineSessions(ctx, in)
}

// StatusGetUsersOnlineSessionsList
// status.getUsersOnlineSessionsList users:Vector<long> = Vector<UserSessionEntryList>;
func (m *defaultStatusClient) StatusGetUsersOnlineSessionsList(ctx context.Context, in *status.TLStatusGetUsersOnlineSessionsList) (*status.Vector_UserSessionEntryList, error) {
	client := status.NewRPCStatusClient(m.cli.Conn())
	return client.StatusGetUsersOnlineSessionsList(ctx, in)
}
