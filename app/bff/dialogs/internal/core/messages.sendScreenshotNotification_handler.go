package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
)

// MessagesSendScreenshotNotification
// messages.sendScreenshotNotification#c97df020 peer:InputPeer reply_to_msg_id:int random_id:long = Updates;
func (c *DialogsCore) MessagesSendScreenshotNotification(in *mtproto.TLMessagesSendScreenshotNotification) (*mtproto.Updates, error) {
	// TODO: not impl
	// c.Logger.Errorf("messages.sendScreenshotNotification blocked, License key from https://elloapp.com required to unlock enterprise features.")

	return mtproto.MakeEmptyUpdates(), nil
}
