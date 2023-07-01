package core

import (
	"errors"

	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

func (c *ChannelsCore) ChannelsDeleteMessages(in *mtproto.TLChannelsDeleteMessages) (*mtproto.Messages_AffectedMessages, error) {
	// todo: add your logic here and delete this line
	return nil, errors.New("Unimplemented")
	// return &mtproto.Messages_AffectedMessages{}, nil
}
