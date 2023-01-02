package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
)

// MessagesMigrateChat
// messages.migrateChat#a2875319 chat_id:long = Updates;
func (c *ChatsCore) MessagesMigrateChat(in *mtproto.TLMessagesMigrateChat) (*mtproto.Updates, error) {
	// TODO: not impl
	c.Logger.Errorf("messages.migrateChat blocked, License key from https://elloapp.com required to unlock enterprise features.")

	return nil, mtproto.ErrEnterpriseIsBlocked
}
