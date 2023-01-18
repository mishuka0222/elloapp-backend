package core

import (
	"errors"

	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

func (c *ChannelsCore) ChannelsGetMessages(in *mtproto.TLChannelsGetMessages) (*mtproto.Messages_Messages, error) {
	// todo: add your logic here and delete this line
	return nil, errors.New("Unimplemented")
	// return &mtproto.Messages_Messages{}, nil
}
