package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/media/media"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

// MediaGetDocument
// media.getDocument id:long = Document;
func (c *MediaCore) MediaGetDocument(in *media.TLMediaGetDocument) (*mtproto.Document, error) {
	document := c.svcCtx.Dao.GetDocumentById(c.ctx, in.GetId())

	return document, nil
}
