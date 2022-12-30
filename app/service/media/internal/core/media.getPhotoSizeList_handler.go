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

// MediaGetPhotoSizeList
// media.getPhotoSizeList size_id:long = PhotoSizeList;
func (c *MediaCore) MediaGetPhotoSizeList(in *media.TLMediaGetPhotoSizeList) (*media.PhotoSizeList, error) {
	szList := c.svcCtx.Dao.GetPhotoSizeListV2(c.ctx, in.GetSizeId())
	if szList == nil {
		szList = []*mtproto.PhotoSize{}
	}

	return &media.PhotoSizeList{
		SizeId: in.GetSizeId(),
		Sizes:  szList,
		DcId:   1,
	}, nil
}
