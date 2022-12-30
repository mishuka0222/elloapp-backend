package dao

import (
	"context"
	"fmt"

	"github.com/gogo/protobuf/types"

	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/media/internal/dal/dataobject"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
)

func (m *Dao) GetVideoSizeListList(ctx context.Context, idList []int64) (sizes map[int64][]*mtproto.VideoSize) {
	sizes = make(map[int64][]*mtproto.VideoSize)
	if len(idList) == 0 {
		return
	}

	sizeDOList, _ := m.VideoSizesDAO.SelectListByVideoSizeIdList(ctx, idList)
	for i := 0; i < len(sizeDOList); i++ {
		szList, ok := sizes[sizeDOList[i].VideoSizeId]
		if !ok {
			szList = []*mtproto.VideoSize{}
		}

		size := &sizeDOList[i]
		videoSize := mtproto.MakeTLVideoSize(&mtproto.VideoSize{
			Type:         size.SizeType,
			W:            size.Width,
			H:            size.Height,
			Size2:        size.FileSize,
			VideoStartTs: nil,
		}).To_VideoSize()
		if size.VideoStartTs > 0 {
			videoSize.VideoStartTs = &types.DoubleValue{Value: size.VideoStartTs}
		}
		szList = append(szList, videoSize)

		sizes[sizeDOList[i].VideoSizeId] = szList
	}
	return
}

func (m *Dao) GetVideoSizeList(ctx context.Context, sizeId int64) (sizes []*mtproto.VideoSize) {
	sizeDOList, _ := m.VideoSizesDAO.SelectListByVideoSizeId(ctx, sizeId)
	if len(sizeDOList) >= 0 {
		sizes = make([]*mtproto.VideoSize, 0, len(sizeDOList))
		for i := 0; i < len(sizeDOList); i++ {
			size := &sizeDOList[i]
			videoSize := mtproto.MakeTLVideoSize(&mtproto.VideoSize{
				Type:         size.SizeType,
				W:            size.Width,
				H:            size.Height,
				Size2:        size.FileSize,
				VideoStartTs: nil,
			}).To_VideoSize()
			if size.VideoStartTs > 0 {
				videoSize.VideoStartTs = &types.DoubleValue{Value: size.VideoStartTs}
			}
			sizes = append(sizes, videoSize)
		}
	}
	return
}

func (m *Dao) SaveVideoSizeV2(ctx context.Context, szId int64, szList []*mtproto.VideoSize) error {
	if len(szList) == 0 {
		return nil
	}

	for _, sz := range szList {
		szDO := &dataobject.VideoSizesDO{
			VideoSizeId:  szId,
			SizeType:     sz.Type,
			Width:        sz.W,
			Height:       sz.H,
			FileSize:     sz.Size2,
			VideoStartTs: sz.GetVideoStartTs().GetValue(),
			FilePath:     fmt.Sprintf("%s/%d.dat", sz.Type, szId),
		}
		if _, _, err := m.VideoSizesDAO.Insert(ctx, szDO); err != nil {
			return err
		}
	}

	return nil
}
