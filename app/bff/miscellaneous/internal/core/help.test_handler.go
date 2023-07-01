package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

// HelpTest
// help.test#c0e202f7 = Bool;
func (c *MiscellaneousCore) HelpTest(in *mtproto.TLHelpTest) (*mtproto.Bool, error) {
	// TODO: not impl
	c.Logger.Errorf("help.test blocked, License key from https://elloapp.com required to unlock enterprise features.")

	return nil, mtproto.ErrEnterpriseIsBlocked
}
