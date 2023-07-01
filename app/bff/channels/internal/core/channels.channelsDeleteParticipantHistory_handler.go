package core

import (
	"errors"

	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

func (c *ChannelsCore) ChannelsDeleteParticipantHistory(in *mtproto.TLChannelsDeleteParticipantHistory) (*mtproto.Messages_AffectedHistory, error) {
	// todo: add your logic here and delete this line
	return nil, errors.New("Unimplemented")
	// return &mtproto.Messages_AffectedHistory{}, nil
}
