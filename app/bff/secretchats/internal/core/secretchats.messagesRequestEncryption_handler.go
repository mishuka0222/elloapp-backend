package core

import (
	"errors"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

func (c *SecretchatsCore) MessagesRequestEncryption(in *mtproto.TLMessagesRequestEncryption) (*mtproto.EncryptedChat, error) {
	// todo: add your logic here and delete this line
	return nil, errors.New("Unimplemented")
	return &mtproto.EncryptedChat{}, nil
}
