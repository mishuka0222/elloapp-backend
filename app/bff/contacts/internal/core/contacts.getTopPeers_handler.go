package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

// ContactsGetTopPeers
// contacts.getTopPeers#973478b6 flags:# correspondents:flags.0?true bots_pm:flags.1?true bots_inline:flags.2?true phone_calls:flags.3?true forward_users:flags.4?true forward_chats:flags.5?true groups:flags.10?true channels:flags.15?true offset:int limit:int hash:long = contacts.TopPeers;
func (c *ContactsCore) ContactsGetTopPeers(in *mtproto.TLContactsGetTopPeers) (*mtproto.Contacts_TopPeers, error) {
	// TODO: not impl
	c.Logger.Errorf("contacts.getTopPeers blocked, License key from https://elloapp.com required to unlock enterprise features.")

	topPeers := mtproto.MakeTLContactsTopPeers(&mtproto.Contacts_TopPeers{
		Categories: []*mtproto.TopPeerCategoryPeers{},
		Chats:      []*mtproto.Chat{},
		Users:      []*mtproto.User{},
	}).To_Contacts_TopPeers()

	return topPeers, nil
}
