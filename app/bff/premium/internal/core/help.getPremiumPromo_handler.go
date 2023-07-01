package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

// HelpGetPremiumPromo
// help.getPremiumPromo#b81b93d4 = help.PremiumPromo;
func (c *PremiumCore) HelpGetPremiumPromo(in *mtproto.TLHelpGetPremiumPromo) (*mtproto.Help_PremiumPromo, error) {
	// TODO: not impl
	c.Logger.Errorf("help.getPremiumPromo blocked, License key from https://elloapp.com required to unlock enterprise features.")

	return nil, mtproto.ErrEnterpriseIsBlocked
}
