package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
)

// HelpDismissSuggestion
// help.dismissSuggestion#f50dbaa1 peer:InputPeer suggestion:string = Bool;
func (c *ConfigurationCore) HelpDismissSuggestion(in *mtproto.TLHelpDismissSuggestion) (*mtproto.Bool, error) {
	// TODO: not impl
	c.Logger.Errorf("help.dismissSuggestion blocked, License key from https://elloapp.com required to unlock enterprise features.")

	return mtproto.BoolTrue, nil
}
