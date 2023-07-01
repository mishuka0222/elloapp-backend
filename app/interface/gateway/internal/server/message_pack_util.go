package server

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

func parseFromIncomingMessage(b []byte) (msgId int64, obj mtproto.TLObject, err error) {
	dBuf := mtproto.NewDecodeBuf(b)

	msgId = dBuf.Long()
	_ = dBuf.Int()
	obj = dBuf.Object()
	err = dBuf.GetError()

	return
}

func serializeToBuffer(msgId int64, obj mtproto.TLObject) []byte {
	oBuf := obj.Encode(0)
	x := mtproto.NewEncodeBuf(8 + 4 + len(oBuf))
	x.Long(0)
	x.Long(msgId)
	x.Int(int32(len(oBuf)))
	x.Bytes(oBuf)

	return x.GetBuf()
}
