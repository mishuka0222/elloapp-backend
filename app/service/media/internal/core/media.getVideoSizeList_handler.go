package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/media/media"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

// MediaGetVideoSizeList
// media.getVideoSizeList size_id = VideoSizeList;
func (c *MediaCore) MediaGetVideoSizeList(in *media.TLMediaGetVideoSizeList) (*media.VideoSizeList, error) {
	szList := c.svcCtx.Dao.GetVideoSizeList(c.ctx, in.GetSizeId())
	if szList == nil {
		szList = []*mtproto.VideoSize{}
	}
	videoSizeList := &media.VideoSizeList{
		SizeId: in.GetSizeId(),
		Sizes:  szList,
		DcId:   1,
	}

	return videoSizeList, nil
}
