package core

import (
	"errors"

	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

func (c *ChannelsCore) ChannelsGetGroupsForDiscussion(in *mtproto.TLChannelsGetGroupsForDiscussion) (*mtproto.Messages_Chats, error) {
	// todo: add your logic here and delete this line
	return nil, errors.New("Unimplemented")
	// return &mtproto.Messages_Chats{}, nil
}
