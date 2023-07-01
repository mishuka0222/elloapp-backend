package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

// HelpGetAppConfig
// help.getAppConfig#98914110 = JSONValue;
func (c *ConfigurationCore) HelpGetAppConfig(in *mtproto.TLHelpGetAppConfig) (*mtproto.JSONValue, error) {
	// TODO: not impl
	c.Logger.Errorf("help.getAppConfig blocked, License key from https://elloapp.com required to unlock enterprise features.")

	return mtproto.MakeTLJsonObject(&mtproto.JSONValue{
		Value_VECTORJSONOBJECTVALUE: []*mtproto.JSONObjectValue{},
	}).To_JSONValue(), nil
}
