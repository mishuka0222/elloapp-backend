package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/media/media"
)

// MediaGetPhotoSizeListList
// media.getPhotoSizeListList id_list:Vector<long> = Vector<PhotoSizeList>;
func (c *MediaCore) MediaGetPhotoSizeListList(in *media.TLMediaGetPhotoSizeListList) (*media.Vector_PhotoSizeList, error) {
	szListList := c.svcCtx.Dao.GetPhotoSizeListList(c.ctx, in.GetIdList())
	photoSizeListList := &media.Vector_PhotoSizeList{
		Datas: make([]*media.PhotoSizeList, 0, len(szListList)),
	}

	for szId, szList := range szListList {
		photoSizeListList.Datas = append(photoSizeListList.Datas, &media.PhotoSizeList{
			SizeId: szId,
			Sizes:  szList,
			DcId:   1,
		})
	}

	return photoSizeListList, nil
}
