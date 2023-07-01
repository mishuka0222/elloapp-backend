package tests

import (
	"context"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/configuration/internal/dao"
	"log"
	"testing"
)

func TestCountriesDAO(t *testing.T) {

	d := dao.New(newConfDao())
	l, err := d.CountriesDAO.SelectList(context.Background())
	if err != nil {
		t.Error(err)
	}
	if len(l) <= 0 {
		t.Errorf("they should be not empty")
	} else {
		log.Printf("count row are selected %d", len(l))
	}
}
