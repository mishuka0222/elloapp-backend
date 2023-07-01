package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

func (c *ChannelsCore) ChannelsReadHistory(in *mtproto.TLChannelsReadHistory) (*mtproto.Bool, error) {

	// TODO: write logic

	return mtproto.BoolTrue, nil
}
