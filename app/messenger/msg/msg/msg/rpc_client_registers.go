package msg

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
	"TLMsgSendMessage":            RPCContextTuple{"/mtproto.RPCMsg/msg_sendMessage", func() interface{} { return new(mtproto.Updates) }},
	"TLMsgSendMultiMessage":       RPCContextTuple{"/mtproto.RPCMsg/msg_sendMultiMessage", func() interface{} { return new(mtproto.Updates) }},
	"TLMsgPushUserMessage":        RPCContextTuple{"/mtproto.RPCMsg/msg_pushUserMessage", func() interface{} { return new(mtproto.Bool) }},
	"TLMsgReadMessageContents":    RPCContextTuple{"/mtproto.RPCMsg/msg_readMessageContents", func() interface{} { return new(mtproto.Messages_AffectedMessages) }},
	"TLMsgSendMessageV2":          RPCContextTuple{"/mtproto.RPCMsg/msg_sendMessageV2", func() interface{} { return new(mtproto.UpdateList) }},
	"TLMsgEditMessage":            RPCContextTuple{"/mtproto.RPCMsg/msg_editMessage", func() interface{} { return new(mtproto.Updates) }},
	"TLMsgDeleteMessages":         RPCContextTuple{"/mtproto.RPCMsg/msg_deleteMessages", func() interface{} { return new(mtproto.Messages_AffectedMessages) }},
	"TLMsgDeleteHistory":          RPCContextTuple{"/mtproto.RPCMsg/msg_deleteHistory", func() interface{} { return new(mtproto.Messages_AffectedHistory) }},
	"TLMsgDeletePhoneCallHistory": RPCContextTuple{"/mtproto.RPCMsg/msg_deletePhoneCallHistory", func() interface{} { return new(mtproto.Messages_AffectedFoundMessages) }},
	"TLMsgDeleteChatHistory":      RPCContextTuple{"/mtproto.RPCMsg/msg_deleteChatHistory", func() interface{} { return new(mtproto.Bool) }},
	"TLMsgReadHistory":            RPCContextTuple{"/mtproto.RPCMsg/msg_readHistory", func() interface{} { return new(mtproto.Messages_AffectedMessages) }},
	"TLMsgUpdatePinnedMessage":    RPCContextTuple{"/mtproto.RPCMsg/msg_updatePinnedMessage", func() interface{} { return new(mtproto.Updates) }},
	"TLMsgUnpinAllMessages":       RPCContextTuple{"/mtproto.RPCMsg/msg_unpinAllMessages", func() interface{} { return new(mtproto.Messages_AffectedHistory) }},
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
