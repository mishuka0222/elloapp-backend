package mtproto

import (
	"encoding/json"
)

const (
	MTPROTO_VERSION = 2
)

type TLObject interface {
	Encode(layer int32) []byte
	Decode(dBuf *DecodeBuf) error
	String() string
	DebugString() string
}

func TLObjectToJson(object TLObject) (b []byte) {
	b, _ = json.Marshal(object)
	return
}
