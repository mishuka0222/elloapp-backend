package tests

import (
	"context"
	"fmt"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/dialog/dialog"
	"log"
	"testing"
)

func TestDialogGetDialogsByIdList(t *testing.T) {
	c := NewRPCClient()

	dialogsList, err := c.DialogGetDialogsByIdList(context.Background(), &dialog.TLDialogGetDialogsByIdList{
		UserId: 777062,
		IdList: []int64{14, 777063, 15, 7, 12, 13, 777015, 777003, 2, 11},
	})
	if err != nil {
		t.Error(err)
		return
	}
	if dialogsList != nil {
		log.Printf("TestDialogGetDialogsByIdList: \nid list count: %d \ndata: \n", len(dialogsList.Datas))
		for _, it := range dialogsList.Datas {
			fmt.Printf("%+s\n", it.Dialog)
			fmt.Printf("peer:\n%+s\n", it.Dialog.Peer)
		}
	}

}
