package updates

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
	"TLUpdatesGetState":               RPCContextTuple{"/mtproto.RPCUpdates/updates_getState", func() interface{} { return new(mtproto.Updates_State) }},
	"TLUpdatesGetDifferenceV2":        RPCContextTuple{"/mtproto.RPCUpdates/updates_getDifferenceV2", func() interface{} { return new(Difference) }},
	"TLUpdatesGetChannelDifferenceV2": RPCContextTuple{"/mtproto.RPCUpdates/updates_getChannelDifferenceV2", func() interface{} { return new(ChannelDifference) }},
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
