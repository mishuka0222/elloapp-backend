package core

import (
	"errors"

	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

func (c *ChannelsCore) ChannelsCreateChannel(in *mtproto.TLChannelsCreateChannel) (*mtproto.Updates, error) {
	// todo: add your logic here and delete this line
	return nil, errors.New("unimplemented 123")
	// return &mtproto.Updates{}, nil
}
