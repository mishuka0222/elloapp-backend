/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2022-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teagramio (teagram.io@gmail.com)
 */

package webpage

import (
	"reflect"

	"github.com/teamgram/proto/mtproto"
)

var _ *mtproto.Bool

type newRPCReplyFunc func() interface{}

type RPCContextTuple struct {
	Method       string
	NewReplyFunc newRPCReplyFunc
}

var rpcContextRegisters = map[string]RPCContextTuple{
	"TLWebpageGetPendingWebPagePreview": RPCContextTuple{"/mtproto.RPCWebpage/webpage_getPendingWebPagePreview", func() interface{} { return new(mtproto.WebPage) }},
	"TLWebpageGetWebPagePreview":        RPCContextTuple{"/mtproto.RPCWebpage/webpage_getWebPagePreview", func() interface{} { return new(mtproto.WebPage) }},
	"TLWebpageGetWebPage":               RPCContextTuple{"/mtproto.RPCWebpage/webpage_getWebPage", func() interface{} { return new(mtproto.WebPage) }},
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
