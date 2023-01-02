package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
)

// MessagesUploadEncryptedFile
// messages.uploadEncryptedFile#5057c497 peer:InputEncryptedChat file:InputEncryptedFile = EncryptedFile;
func (c *FilesCore) MessagesUploadEncryptedFile(in *mtproto.TLMessagesUploadEncryptedFile) (*mtproto.EncryptedFile, error) {
	// TODO: not impl
	c.Logger.Errorf("messages.uploadEncryptedFile blocked, License key from https://elloapp.com required to unlock enterprise features.")

	return nil, mtproto.ErrEnterpriseIsBlocked
}
