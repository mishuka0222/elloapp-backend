package core

import (
	"errors"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

func (c *SecretchatsCore) MessagesSendEncryptedService(in *mtproto.TLMessagesSendEncryptedService) (*mtproto.Messages_SentEncryptedMessage, error) {
	// todo: add your logic here and delete this line
	return nil, errors.New("Unimplemented")
	return &mtproto.Messages_SentEncryptedMessage{}, nil
}
