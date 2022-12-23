package tests

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/teamgram/teamgram-server/app/bff/feeds/internal/dao"
	"log"
	"testing"
)

func TestInsertListDAO(t *testing.T) {

	d := dao.New(newConfDao())
	l, err := d.UserFeedsDAO.InsertList(context.Background(), 777062, []int64{12, 11, 1, 14})
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

func TestSelectReadOutboxListDAO(t *testing.T) {

	d := dao.New(newConfDao())
	l, err := d.UserFeedsDAO.SelectReadOutboxList(context.Background(), 777062)
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
	l, err := d.UserFeedsDAO.DeleteFromListElseValue(context.Background(), 777062, []int64{12, 11, 1})
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
