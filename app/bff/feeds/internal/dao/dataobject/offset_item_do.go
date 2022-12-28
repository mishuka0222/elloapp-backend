package dataobject

type OffsetItemDo struct {
	OffsetMin int32 `json:"offset_min"`
	OffsetMax int32 `json:"offset_max"`
	PeerID    int64 `json:"peer_id"`
	Count     int32 `json:"count"`
}

type OffsetItemList []OffsetItemDo

func (l OffsetItemList) CalcCount() (count int32) {
	for i := range l {
		count += l[i].Count
	}
	return
}
