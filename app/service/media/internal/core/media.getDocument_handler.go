/*
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/media/media"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
)

// MediaGetDocument
// media.getDocument id:long = Document;
func (c *MediaCore) MediaGetDocument(in *media.TLMediaGetDocument) (*mtproto.Document, error) {
	document := c.svcCtx.Dao.GetDocumentById(c.ctx, in.GetId())

	return document, nil
}
