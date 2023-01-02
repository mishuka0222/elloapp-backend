package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
)

// HelpAcceptTermsOfService
// help.acceptTermsOfService#ee72f79a id:DataJSON = Bool;
func (c *TosCore) HelpAcceptTermsOfService(in *mtproto.TLHelpAcceptTermsOfService) (*mtproto.Bool, error) {
	// TODO: not impl
	c.Logger.Errorf("help.acceptTermsOfService blocked, License key from https://elloapp.com required to unlock enterprise features.")

	return mtproto.BoolTrue, nil
}
