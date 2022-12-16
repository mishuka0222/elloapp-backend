package tests

import (
	"context"
	"fmt"
	"github.com/teamgram/teamgram-server/app/service/biz/dialog/dialog"
	"log"
	"testing"
)

func TestDialogGetDialogsByIdList(t *testing.T) {
	c := NewDialogClient()

	idList, err := c.DialogGetDialogsByIdList(context.Background(), &dialog.TLDialogGetDialogsByIdList{
		UserId: 777062,
		IdList: []int64{14, 777063},
	})
	if err != nil {
		t.Error(err)
		return
	}
	if idList != nil {
		log.Printf("TestDialogGetDialogsByIdList: \nid list count: %d \ndata: \n", len(idList.Datas))
		for _, it := range idList.Datas {
			fmt.Printf("%+s\n", it.Dialog)
			fmt.Printf("peer:\n%+s\n", it.Dialog.Peer)
		}
	}

}
