package status

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
	"TLStatusSetSessionOnline":           RPCContextTuple{"/mtproto.RPCStatus/status_setSessionOnline", func() interface{} { return new(mtproto.Bool) }},
	"TLStatusSetSessionOffline":          RPCContextTuple{"/mtproto.RPCStatus/status_setSessionOffline", func() interface{} { return new(mtproto.Bool) }},
	"TLStatusGetUserOnlineSessions":      RPCContextTuple{"/mtproto.RPCStatus/status_getUserOnlineSessions", func() interface{} { return new(UserSessionEntryList) }},
	"TLStatusGetUsersOnlineSessionsList": RPCContextTuple{"/mtproto.RPCStatus/status_getUsersOnlineSessionsList", func() interface{} { return new(Vector_UserSessionEntryList) }},
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
