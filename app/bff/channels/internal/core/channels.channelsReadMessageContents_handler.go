package core

import (
	"errors"

	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

func (c *ChannelsCore) ChannelsReadMessageContents(in *mtproto.TLChannelsReadMessageContents) (*mtproto.Bool, error) {

	return mtproto.BoolTrue, nil
	// todo: add your logic here and delete this line
	return nil, errors.New("Unimplemented")
	// return &mtproto.Bool{}, nil
}
