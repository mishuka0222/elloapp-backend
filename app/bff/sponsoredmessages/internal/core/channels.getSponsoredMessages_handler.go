package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

// ChannelsGetSponsoredMessages
// channels.getSponsoredMessages#ec210fbf channel:InputChannel = messages.SponsoredMessages;
func (c *SponsoredMessagesCore) ChannelsGetSponsoredMessages(in *mtproto.TLChannelsGetSponsoredMessages) (*mtproto.Messages_SponsoredMessages, error) {
	// disable SponsoredMessages
	return mtproto.MakeTLMessagesSponsoredMessages(&mtproto.Messages_SponsoredMessages{
		Messages: []*mtproto.SponsoredMessage{},
		Chats:    []*mtproto.Chat{},
		Users:    []*mtproto.User{},
	}).To_Messages_SponsoredMessages(), nil
}
