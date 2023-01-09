package inbox

import (
	"reflect"

	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

var _ *mtproto.Bool

type newRPCReplyFunc func() interface{}

type RPCContextTuple struct {
	Method       string
	NewReplyFunc newRPCReplyFunc
}

var rpcContextRegisters = map[string]RPCContextTuple{
	"TLInboxSendUserMessageToInbox":         RPCContextTuple{"/mtproto.RPCInbox/inbox_sendUserMessageToInbox", func() interface{} { return new(mtproto.Void) }},
	"TLInboxSendChatMessageToInbox":         RPCContextTuple{"/mtproto.RPCInbox/inbox_sendChatMessageToInbox", func() interface{} { return new(mtproto.Void) }},
	"TLInboxSendChannelMessageToInbox":      RPCContextTuple{"/mtproto.RPCInbox/inbox_sendChannelMessageToInbox", func() interface{} { return new(mtproto.Void) }},
	"TLInboxSendUserMultiMessageToInbox":    RPCContextTuple{"/mtproto.RPCInbox/inbox_sendUserMultiMessageToInbox", func() interface{} { return new(mtproto.Void) }},
	"TLInboxSendChatMultiMessageToInbox":    RPCContextTuple{"/mtproto.RPCInbox/inbox_sendChatMultiMessageToInbox", func() interface{} { return new(mtproto.Void) }},
	"TLInboxSendChannelMultiMessageToInbox": RPCContextTuple{"/mtproto.RPCInbox/inbox_sendChannelMultiMessageToInbox", func() interface{} { return new(mtproto.Void) }},
	"TLInboxEditUserMessageToInbox":         RPCContextTuple{"/mtproto.RPCInbox/inbox_editUserMessageToInbox", func() interface{} { return new(mtproto.Void) }},
	"TLInboxEditChatMessageToInbox":         RPCContextTuple{"/mtproto.RPCInbox/inbox_editChatMessageToInbox", func() interface{} { return new(mtproto.Void) }},
	"TLInboxEditChannelMessageToInbox":      RPCContextTuple{"/mtproto.RPCInbox/inbox_editChannelMessageToInbox", func() interface{} { return new(mtproto.Void) }},
	"TLInboxDeleteMessagesToInbox":          RPCContextTuple{"/mtproto.RPCInbox/inbox_deleteMessagesToInbox", func() interface{} { return new(mtproto.Void) }},
	"TLInboxDeleteChannelMessagesToInbox":   RPCContextTuple{"/mtproto.RPCInbox/inbox_deleteChannelMessagesToInbox", func() interface{} { return new(mtproto.Void) }},
	"TLInboxDeleteUserHistoryToInbox":       RPCContextTuple{"/mtproto.RPCInbox/inbox_deleteUserHistoryToInbox", func() interface{} { return new(mtproto.Void) }},
	"TLInboxDeleteChatHistoryToInbox":       RPCContextTuple{"/mtproto.RPCInbox/inbox_deleteChatHistoryToInbox", func() interface{} { return new(mtproto.Void) }},
	"TLInboxReadUserMediaUnreadToInbox":     RPCContextTuple{"/mtproto.RPCInbox/inbox_readUserMediaUnreadToInbox", func() interface{} { return new(mtproto.Void) }},
	"TLInboxReadChatMediaUnreadToInbox":     RPCContextTuple{"/mtproto.RPCInbox/inbox_readChatMediaUnreadToInbox", func() interface{} { return new(mtproto.Void) }},
	"TLInboxReadChannelMediaUnreadToInbox":  RPCContextTuple{"/mtproto.RPCInbox/inbox_readChannelMediaUnreadToInbox", func() interface{} { return new(mtproto.Void) }},
	"TLInboxUpdateHistoryReaded":            RPCContextTuple{"/mtproto.RPCInbox/inbox_updateHistoryReaded", func() interface{} { return new(mtproto.Void) }},
	"TLInboxUpdatePinnedMessage":            RPCContextTuple{"/mtproto.RPCInbox/inbox_updatePinnedMessage", func() interface{} { return new(mtproto.Void) }},
	"TLInboxUnpinAllMessages":               RPCContextTuple{"/mtproto.RPCInbox/inbox_unpinAllMessages", func() interface{} { return new(mtproto.Void) }},
}

func FindRPCContextTuple(t interface{}) *RPCContextTuple {
	rt := reflect.TypeOf(t)
	if rt.Kind() == reflect.Ptr {
		rt = rt.Elem()
	}

	m, ok := rpcContextRegisters[rt.Name()]
	if !ok {
		// log.Errorf("Can't find name: %s", rt.Name())
		return nil
	}
	return &m
}

func GetRPCContextRegisters() map[string]RPCContextTuple {
	return rpcContextRegisters
}
