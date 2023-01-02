package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
)

// HelpSaveAppLog
// help.saveAppLog#6f02f748 events:Vector<InputAppEvent> = Bool;
func (c *MiscellaneousCore) HelpSaveAppLog(in *mtproto.TLHelpSaveAppLog) (*mtproto.Bool, error) {
	// TODO: not impl
	c.Logger.Errorf("help.saveAppLog blocked, License key from https://elloapp.com required to unlock enterprise features.")

	return nil, mtproto.ErrEnterpriseIsBlocked
}
