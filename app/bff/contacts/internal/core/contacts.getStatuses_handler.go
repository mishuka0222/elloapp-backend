package core

import (
	userpb "gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/user/user"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
)

// ContactsGetStatuses
// contacts.getStatuses#c4a353ee = Vector<ContactStatus>;
func (c *ContactsCore) ContactsGetStatuses(in *mtproto.TLContactsGetStatuses) (*mtproto.Vector_ContactStatus, error) {
	cList, _ := c.svcCtx.Dao.UserClient.UserGetContactList(c.ctx, &userpb.TLUserGetContactList{
		UserId: c.MD.UserId,
	})

	rList := &mtproto.Vector_ContactStatus{
		Datas: make([]*mtproto.ContactStatus, 0, len(cList.GetDatas())),
	}

	idList := make([]int64, 0, len(cList.GetDatas()))
	for _, id := range cList.GetDatas() {
		idList = append(idList, id.ContactUserId)
	}

	lastSeenList, _ := c.svcCtx.Dao.UserGetLastSeens(c.ctx, &userpb.TLUserGetLastSeens{
		Id: idList,
	})

	for _, v := range lastSeenList.GetDatas() {
		rList.Datas = append(rList.Datas, mtproto.MakeTLContactStatus(&mtproto.ContactStatus{
			UserId: v.UserId,
			Status: userpb.MakeUserStatus(v.LastSeenAt, true),
		}).To_ContactStatus())
	}

	return rList, nil
}
