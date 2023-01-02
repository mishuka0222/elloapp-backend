package service

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
)

func ParseFromIncomingMessage(b []byte) (salt, sessionId int64, msg2 *mtproto.TLMessage2, err error) {
	dbuf := mtproto.NewDecodeBuf(b)
	msg2 = &mtproto.TLMessage2{}

	salt = dbuf.Long()      // salt
	sessionId = dbuf.Long() // session_id
	err = msg2.Decode(dbuf)

	return
}

func SerializeToBuffer(salt, sessionId int64, msg2 *mtproto.TLMessage2, layer int32) []byte {
	oBuf := msg2.Encode(layer)

	x := mtproto.NewEncodeBuf(16 + len(oBuf))
	x.Long(salt)
	x.Long(sessionId)
	x.Bytes(oBuf)

	return x.GetBuf()
}

func SerializeToBuffer2(salt, sessionId int64, msg2 *mtproto.TLMessageRawData) []byte {
	x := mtproto.NewEncodeBuf(32 + len(msg2.Body))

	x.Long(salt)
	x.Long(sessionId)
	x.Long(msg2.MsgId)
	x.Int(msg2.Seqno)
	x.Int(msg2.Bytes)
	x.Bytes(msg2.Body)

	return x.GetBuf()
}
