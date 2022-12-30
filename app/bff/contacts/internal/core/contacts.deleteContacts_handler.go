package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/messenger/sync/sync"
	userpb "gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/user/user"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
)

// ContactsDeleteContacts
// contacts.deleteContacts#96a0e00 id:Vector<InputUser> = Updates;
func (c *ContactsCore) ContactsDeleteContacts(in *mtproto.TLContactsDeleteContacts) (*mtproto.Updates, error) {
	var (
		rUpdates = mtproto.MakeEmptyUpdates()
		idHelper = mtproto.NewIDListHelper(c.MD.UserId)
	)

	for _, id := range in.GetId() {
		switch id.GetPredicateName() {
		case mtproto.Predicate_inputUser:
		default:
			// TODO:
			c.Logger.Errorf("contacts.deleteContacts - error: invalid id %v", id)
			continue
		}
		idHelper.AppendUsers(id.GetUserId())
	}

	deleteUsers, err := c.svcCtx.Dao.UserGetMutableUsers(c.ctx, &userpb.TLUserGetMutableUsers{
		Id: idHelper.UserIdList,
	})
	if err != nil {
		c.Logger.Errorf("contacts.deleteContacts - error: %v", err)
		return nil, err
	}

	deleteUsers.VisitByMe(c.MD.UserId, func(me, it *userpb.ImmutableUser) {
		if me.Id() != it.Id() {
			// TODO: mutual
			c.svcCtx.Dao.UserClient.UserDeleteContact(c.ctx, &userpb.TLUserDeleteContact{
				UserId: c.MD.UserId,
				Id:     it.Id(),
			})

			rUpdates.PushBackUpdate(
				mtproto.MakeTLUpdatePeerSettings(&mtproto.Update{
					Peer_PEER: mtproto.MakePeerUser(it.Id()),
					Settings: mtproto.MakeTLPeerSettings(&mtproto.PeerSettings{
						ReportSpam:            false,
						AddContact:            false,
						BlockContact:          false,
						ShareContact:          false,
						NeedContactsException: false,
						ReportGeo:             false,
						Autoarchived:          false,
						GeoDistance:           nil,
					}).To_PeerSettings(),
				}).To_Update())

			cUser := it.ToUnsafeUser(me)
			cUser.Contact = false
			cUser.MutualContact = false
			rUpdates.PushUser(cUser)
		}
	})

	c.svcCtx.Dao.SyncClient.SyncUpdatesNotMe(c.ctx, &sync.TLSyncUpdatesNotMe{
		UserId:    c.MD.UserId,
		AuthKeyId: c.MD.AuthId,
		Updates:   rUpdates,
	})

	return rUpdates, nil
}
