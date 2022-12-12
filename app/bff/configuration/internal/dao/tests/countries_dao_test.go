package tests

import (
	"context"
	"github.com/teamgram/teamgram-server/app/bff/configuration/internal/dao"
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
	}
}
