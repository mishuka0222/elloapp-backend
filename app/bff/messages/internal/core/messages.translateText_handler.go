package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

// MessagesTranslateText
// messages.translateText#24ce6dee flags:# peer:flags.0?InputPeer msg_id:flags.0?int text:flags.1?string from_lang:flags.2?string to_lang:string = messages.TranslatedText;
func (c *MessagesCore) MessagesTranslateText(in *mtproto.TLMessagesTranslateText) (*mtproto.Messages_TranslatedText, error) {
	// TODO: not impl
	c.Logger.Errorf("messages.translateText blocked, License key from https://elloapp.com required to unlock enterprise features.")

	return nil, mtproto.ErrEnterpriseIsBlocked
}
