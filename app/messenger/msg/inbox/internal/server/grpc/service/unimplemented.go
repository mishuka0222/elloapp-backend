package service

import (
	"context"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/messenger/msg/inbox/inbox"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
)

func (s *Service) InboxSendChannelMessageToInbox(ctx context.Context, inbox *inbox.TLInboxSendChannelMessageToInbox) (*mtproto.Void, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) InboxSendChannelMultiMessageToInbox(ctx context.Context, inbox *inbox.TLInboxSendChannelMultiMessageToInbox) (*mtproto.Void, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) InboxEditChannelMessageToInbox(ctx context.Context, inbox *inbox.TLInboxEditChannelMessageToInbox) (*mtproto.Void, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) InboxDeleteChannelMessagesToInbox(ctx context.Context, inbox *inbox.TLInboxDeleteChannelMessagesToInbox) (*mtproto.Void, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) InboxReadChannelMediaUnreadToInbox(ctx context.Context, inbox *inbox.TLInboxReadChannelMediaUnreadToInbox) (*mtproto.Void, error) {
	//TODO implement me
	panic("implement me")
}
