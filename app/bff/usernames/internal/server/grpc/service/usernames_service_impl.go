package service

import (
	"context"

	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/bff/usernames/internal/core"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

// AccountCheckUsername
// account.checkUsername#2714d86c username:string = Bool;
func (s *Service) AccountCheckUsername(ctx context.Context, request *mtproto.TLAccountCheckUsername) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("account.checkUsername - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.AccountCheckUsername(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("account.checkUsername - reply: %s", r.DebugString())
	return r, err
}

// AccountUpdateUsername
// account.updateUsername#3e0bdd7c username:string = User;
func (s *Service) AccountUpdateUsername(ctx context.Context, request *mtproto.TLAccountUpdateUsername) (*mtproto.User, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("account.updateUsername - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.AccountUpdateUsername(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("account.updateUsername - reply: %s", r.DebugString())
	return r, err
}

// ContactsResolveUsername
// contacts.resolveUsername#f93ccba3 username:string = contacts.ResolvedPeer;
func (s *Service) ContactsResolveUsername(ctx context.Context, request *mtproto.TLContactsResolveUsername) (*mtproto.Contacts_ResolvedPeer, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("contacts.resolveUsername - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.ContactsResolveUsername(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("contacts.resolveUsername - reply: %s", r.DebugString())
	return r, err
}

// ChannelsCheckUsername
// channels.checkUsername#10e6bd2c channel:InputChannel username:string = Bool;
func (s *Service) ChannelsCheckUsername(ctx context.Context, request *mtproto.TLChannelsCheckUsername) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("channels.checkUsername - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.ChannelsCheckUsername(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("channels.checkUsername - reply: %s", r.DebugString())
	return r, err
}

// ChannelsUpdateUsername
// channels.updateUsername#3514b3de channel:InputChannel username:string = Bool;
func (s *Service) ChannelsUpdateUsername(ctx context.Context, request *mtproto.TLChannelsUpdateUsername) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("channels.updateUsername - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.ChannelsUpdateUsername(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("channels.updateUsername - reply: %s", r.DebugString())
	return r, err
}
