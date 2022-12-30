package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
)

// HelpGetCdnConfig
// help.getCdnConfig#52029342 = CdnConfig;
func (c *FilesCore) HelpGetCdnConfig(in *mtproto.TLHelpGetCdnConfig) (*mtproto.CdnConfig, error) {
	// TODO: not impl
	c.Logger.Errorf("help.getCdnConfig blocked, License key from https://elloapp.com required to unlock enterprise features.")

	rValue := mtproto.MakeTLCdnConfig(&mtproto.CdnConfig{
		PublicKeys: []*mtproto.CdnPublicKey{},
	}).To_CdnConfig()

	return rValue, nil
}
