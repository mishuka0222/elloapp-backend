package tests

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/teamgram/teamgram-server/app/service/biz/feeds/feeds"
	"github.com/teamgram/teamgram-server/app/service/biz/feeds/internal/dao"
	"log"
	"testing"
)

func TestInsertListDAO(t *testing.T) {

	d := dao.New(newConfDao())
	l, err := d.UserFeedsDAO.InsertList(context.Background(), 777062, []*feeds.FeedInsertItemDO{
		{ChatId: 12, PeerType: 2},
		{ChatId: 11, PeerType: 2},
		{ChatId: 1, PeerType: 2},
		{ChatId: 14, PeerType: 2},
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, int64(4), l)
	log.Printf("count row are inserted %d", l)

}

func TestSelectFeedListDAO(t *testing.T) {

	d := dao.New(newConfDao())
	l, err := d.UserFeedsDAO.SelectFeedList(context.Background(), 777062)
	if err != nil {
		t.Fatal(err)
	}
	if len(l) <= 0 {
		t.Fatalf("they should be not empty, count %d", len(l))
	} else {
		log.Printf("count row are selected %d", len(l))
	}
}

func TestSelectChatListDAO(t *testing.T) {

	d := dao.New(newConfDao())
	l, err := d.UserFeedsDAO.SelectChatList(context.Background(), 777062)
	if err != nil {
		t.Fatal(err)
	}
	if len(l) <= 0 {
		t.Fatalf("they should be not empty, count %d", len(l))
	} else {
		log.Printf("count row are selected %d", len(l))
	}
}

func TestSelectUnreadCountDAO(t *testing.T) {

	d := dao.New(newConfDao())
	l, err := d.UserFeedsDAO.SelectUnreadCountList(context.Background(), 777062)
	if err != nil {
		t.Fatal(err)
	}
	if len(l) <= 0 {
		t.Fatalf("they should be not empty, count %d", len(l))
	} else {
		log.Printf("count row are selected %d", len(l))
	}
}

func TestSelectOffsetMinListDAO(t *testing.T) {

	d := dao.New(newConfDao())
	l, err := d.UserFeedsDAO.SelectOffsetMinList(context.Background(),
		777015, []int64{2, 3, 4, 5, 6, 7, 8, 11, 13}, 50, 20)
	if err != nil {
		t.Fatal(err)
	}
	if len(l) <= 0 {
		t.Fatalf("they should be not empty, count %d", len(l))
	} else {
		log.Printf("count row are selected %d", len(l))
	}
}

func TestSelectOffsetMaxListDAO(t *testing.T) {

	d := dao.New(newConfDao())
	l, err := d.UserFeedsDAO.SelectOffsetMaxList(context.Background(),
		777015, []int64{2, 3, 4, 5, 6, 7, 8, 11, 13}, 50)
	if err != nil {
		t.Fatal(err)
	}
	if len(l) <= 0 {
		t.Fatalf("they should be not empty, count %d", len(l))
	} else {
		log.Printf("count row are selected %d", len(l))
	}
}

func TestDeleteFromListElseValueDAO(t *testing.T) {

	d := dao.New(newConfDao())
	l, err := d.UserFeedsDAO.DeleteFromListElseValue(context.Background(), 777062, []*feeds.FeedInsertItemDO{
		{ChatId: 12, PeerType: 2},
		{ChatId: 11, PeerType: 2},
		{ChatId: 1, PeerType: 2},
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, int64(1), l)
	log.Printf("count row are deleted %d", l)

}

func TestClear(t *testing.T) {
	d := dao.New(newConfDao())
	if _, err := d.UserFeedsDAO.DeleteFromListElseValue(context.Background(), 777062, nil); err != nil {
		panic(err)
	}
}
