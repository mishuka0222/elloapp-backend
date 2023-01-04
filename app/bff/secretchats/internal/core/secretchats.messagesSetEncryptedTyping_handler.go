package core

import (
	"errors"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

func (c *SecretchatsCore) MessagesSetEncryptedTyping(in *mtproto.TLMessagesSetEncryptedTyping) (*mtproto.Bool, error) {
	// todo: add your logic here and delete this line
	return nil, errors.New("Unimplemented")
	return &mtproto.Bool{}, nil
}
