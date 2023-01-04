package core

import (
	"errors"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

func (c *SecretchatsCore) MessagesReceivedQueue(in *mtproto.TLMessagesReceivedQueue) (*mtproto.Vector_Long, error) {
	// todo: add your logic here and delete this line
	return nil, errors.New("Unimplemented")
	return &mtproto.Vector_Long{}, nil
}
