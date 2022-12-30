package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
)

// MessagesGetDocumentByHash
// messages.getDocumentByHash#338e2464 sha256:bytes size:int mime_type:string = Document;
func (c *FilesCore) MessagesGetDocumentByHash(in *mtproto.TLMessagesGetDocumentByHash) (*mtproto.Document, error) {
	// TODO: not impl
	c.Logger.Errorf("messages.getDocumentByHash blocked, License key from https://elloapp.com required to unlock enterprise features.")

	return mtproto.MakeTLDocumentEmpty(nil).To_Document(), nil
}
