package core

import (
	"time"

	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

// HelpGetAppChangelog
// help.getAppChangelog#9010ef6f prev_app_version:string = Updates;
func (c *ConfigurationCore) HelpGetAppChangelog(in *mtproto.TLHelpGetAppChangelog) (*mtproto.Updates, error) {
	// TODO: not impl
	c.Logger.Errorf("help.getAppChangelog blocked, License key from https://elloapp.com required to unlock enterprise features.")

	_ = in.GetPrevAppVersion()

	rValue := mtproto.MakeTLUpdates(&mtproto.Updates{
		Updates: []*mtproto.Update{},
		Users:   []*mtproto.User{},
		Chats:   []*mtproto.Chat{},
		Date:    int32(time.Now().Unix()),
		Seq:     0,
	}).To_Updates()

	return rValue, nil
}
