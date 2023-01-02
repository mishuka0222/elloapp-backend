package inbox_client

import (
	"context"
	"fmt"

	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/messenger/msg/inbox/inbox"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
	kafka "gitlab.com/merehead/elloapp/backend/elloapp_backend/pkg2/mq"

	"github.com/gogo/protobuf/proto"
	"github.com/zeromicro/go-zero/core/jsonx"
)

type defaultInboxMqClient struct {
	cli *kafka.Producer
}

func NewInboxMqClient(cli *kafka.Producer) InboxClient {
	return &defaultInboxMqClient{
		cli: cli,
	}
}

func (m *defaultInboxMqClient) sendMessage(ctx context.Context, k string, in interface{}) (*mtproto.Void, error) {
	var (
		b   []byte
		err error
	)

	b, err = jsonx.Marshal(in)
	if err != nil {
		return nil, err
	}

	_, _, err = m.cli.SendMessage(ctx, k, b)
	if err != nil {
		return nil, err
	}

	return mtproto.EmptyVoid, nil
}

// InboxSendUserMessageToInbox
// inbox.sendUserMessageToInbox from_id:long peer_user_id:long message:InboxMessageData = Void;
func (m *defaultInboxMqClient) InboxSendUserMessageToInbox(ctx context.Context, in *inbox.TLInboxSendUserMessageToInbox) (*mtproto.Void, error) {
	return m.sendMessage(
		ctx,
		fmt.Sprintf("%s#%d", proto.MessageName(in), in.GetPeerUserId()),
		in)
}

// InboxSendChatMessageToInbox
// inbox.sendChatMessageToInbox from_id:long peer_chat_id:long message:InboxMessageData = Void;
func (m *defaultInboxMqClient) InboxSendChatMessageToInbox(ctx context.Context, in *inbox.TLInboxSendChatMessageToInbox) (*mtproto.Void, error) {
	return m.sendMessage(
		ctx,
		fmt.Sprintf("%s#%d", proto.MessageName(in), in.GetPeerChatId()),
		in)
}

// InboxSendUserMultiMessageToInbox
// inbox.sendUserMultiMessageToInbox from_id:long peer_user_id:long message:Vector<InboxMessageData> = Void;
func (m *defaultInboxMqClient) InboxSendUserMultiMessageToInbox(ctx context.Context, in *inbox.TLInboxSendUserMultiMessageToInbox) (*mtproto.Void, error) {
	return m.sendMessage(
		ctx,
		fmt.Sprintf("%s#%d", proto.MessageName(in), in.GetPeerUserId()),
		in)
}

// InboxSendChatMultiMessageToInbox
// inbox.sendChatMultiMessageToInbox from_id:long peer_chat_id:long message:Vector<InboxMessageData> = Void;
func (m *defaultInboxMqClient) InboxSendChatMultiMessageToInbox(ctx context.Context, in *inbox.TLInboxSendChatMultiMessageToInbox) (*mtproto.Void, error) {
	return m.sendMessage(
		ctx,
		fmt.Sprintf("%s#%d", proto.MessageName(in), in.GetPeerChatId()),
		in)
}

// InboxEditUserMessageToInbox
// inbox.editUserMessageToInbox from_id:long peer_user_id:long message:Message = Void;
func (m *defaultInboxMqClient) InboxEditUserMessageToInbox(ctx context.Context, in *inbox.TLInboxEditUserMessageToInbox) (*mtproto.Void, error) {
	return m.sendMessage(
		ctx,
		fmt.Sprintf("%s#%d", proto.MessageName(in), in.GetPeerUserId()),
		in)
}

// InboxEditChatMessageToInbox
// inbox.editChatMessageToInbox from_id:long peer_chat_id:long message:Message = Void;
func (m *defaultInboxMqClient) InboxEditChatMessageToInbox(ctx context.Context, in *inbox.TLInboxEditChatMessageToInbox) (*mtproto.Void, error) {
	return m.sendMessage(
		ctx,
		fmt.Sprintf("%s#%d", proto.MessageName(in), in.GetPeerChatId()),
		in)
}

// InboxDeleteMessagesToInbox
// inbox.deleteMessagesToInbox from_id:long id:Vector<int> = Void;
func (m *defaultInboxMqClient) InboxDeleteMessagesToInbox(ctx context.Context, in *inbox.TLInboxDeleteMessagesToInbox) (*mtproto.Void, error) {
	return m.sendMessage(
		ctx,
		fmt.Sprintf("%s#%d", proto.MessageName(in), in.GetPeerId()),
		in)
}

// InboxDeleteUserHistoryToInbox
// inbox.deleteUserHistoryToInbox flags:# from_id:long peer_user_id:long just_clear:flags.1?true max_id:int = Void;
func (m *defaultInboxMqClient) InboxDeleteUserHistoryToInbox(ctx context.Context, in *inbox.TLInboxDeleteUserHistoryToInbox) (*mtproto.Void, error) {
	return m.sendMessage(
		ctx,
		fmt.Sprintf("%s#%d", proto.MessageName(in), in.GetPeerUserId()),
		in)
}

// InboxDeleteChatHistoryToInbox
// inbox.deleteChatHistoryToInbox from_id:long peer_chat_id:long max_id:int = Void;
func (m *defaultInboxMqClient) InboxDeleteChatHistoryToInbox(ctx context.Context, in *inbox.TLInboxDeleteChatHistoryToInbox) (*mtproto.Void, error) {
	return m.sendMessage(
		ctx,
		fmt.Sprintf("%s#%d", proto.MessageName(in), in.GetPeerChatId()),
		in)
}

// InboxReadUserMediaUnreadToInbox
// inbox.readUserMediaUnreadToInbox from_id:long id:Vector<int> = Void;
func (m *defaultInboxMqClient) InboxReadUserMediaUnreadToInbox(ctx context.Context, in *inbox.TLInboxReadUserMediaUnreadToInbox) (*mtproto.Void, error) {
	return m.sendMessage(
		ctx,
		fmt.Sprintf("%s#%d", proto.MessageName(in), in.GetPeerUserId()),
		in)
}

// InboxReadChatMediaUnreadToInbox
// inbox.readChatMediaUnreadToInbox from_id:long peer_chat_id:long id:Vector<int> = Void;
func (m *defaultInboxMqClient) InboxReadChatMediaUnreadToInbox(ctx context.Context, in *inbox.TLInboxReadChatMediaUnreadToInbox) (*mtproto.Void, error) {
	return m.sendMessage(
		ctx,
		fmt.Sprintf("%s#%d", proto.MessageName(in), in.GetPeerChatId()),
		in)
}

// InboxUpdateHistoryReaded
// inbox.updateHistoryReaded from_id:long peer_type:int peer_id:long max_id:int = Void;
func (m *defaultInboxMqClient) InboxUpdateHistoryReaded(ctx context.Context, in *inbox.TLInboxUpdateHistoryReaded) (*mtproto.Void, error) {
	return m.sendMessage(
		ctx,
		fmt.Sprintf("%s#%d", proto.MessageName(in), in.GetPeerId()),
		in)
}

// InboxUpdatePinnedMessage
// inbox.updatePinnedMessage flags:# user_id:long auth_key_id:long silent:flags.0?true unpin:flags.1?true pm_oneside:flags.2?true peer_type:int peer_id:long id:int = Updates;
func (m *defaultInboxMqClient) InboxUpdatePinnedMessage(ctx context.Context, in *inbox.TLInboxUpdatePinnedMessage) (*mtproto.Void, error) {
	return m.sendMessage(
		ctx,
		fmt.Sprintf("%s#%d", proto.MessageName(in), in.GetPeerId()),
		in)
}

// InboxUnpinAllMessages
// inbox.unpinAllMessages user_id:long auth_key_id:long peer_type:int peer_id:long = messages.AffectedHistory;
func (m *defaultInboxMqClient) InboxUnpinAllMessages(ctx context.Context, in *inbox.TLInboxUnpinAllMessages) (*mtproto.Void, error) {
	return m.sendMessage(
		ctx,
		fmt.Sprintf("%s#%d", proto.MessageName(in), in.GetPeerId()),
		in)
}
