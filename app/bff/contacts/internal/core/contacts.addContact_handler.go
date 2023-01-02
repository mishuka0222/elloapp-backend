package core

import (
	userpb "gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/user/user"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
)

// ContactsAddContact
// contacts.addContact#e8f463d0 flags:# add_phone_privacy_exception:flags.0?true id:InputUser first_name:string last_name:string phone:string = Updates;
func (c *ContactsCore) ContactsAddContact(in *mtproto.TLContactsAddContact) (*mtproto.Updates, error) {
	// 400	CONTACT_NAME_EMPTY	Contact name empty.
	if in.FirstName == "" && in.LastName == "" {
		err := mtproto.ErrContactNameEmpty
		c.Logger.Errorf("contacts.addContact - error: %v", err)
		return nil, err
	}

	id := mtproto.FromInputUser(c.MD.UserId, in.Id)

	// TODO: check inputUserFromMessage
	if !id.IsUser() || id.IsSelf() || id.PeerId == c.MD.UserId {
		err := mtproto.ErrContactIdInvalid
		c.Logger.Errorf("contacts.addContact - error: %v", err)
		return nil, err
	}

	users, err := c.svcCtx.Dao.UserClient.UserGetMutableUsers(c.ctx, &userpb.TLUserGetMutableUsers{
		Id: []int64{c.MD.UserId, id.PeerId},
	})
	if err != nil {
		c.Logger.Errorf("contacts.addContact - error: %v", err)
		err = mtproto.ErrContactIdInvalid
		return nil, err
	}

	if !users.CheckExistUser(c.MD.UserId, id.PeerId) {
		err = mtproto.ErrContactIdInvalid
		c.Logger.Errorf("contacts.addContact - error: %v", err)
		return nil, err
	}

	changeMutual, err := c.svcCtx.Dao.UserClient.UserAddContact(c.ctx, &userpb.TLUserAddContact{
		UserId:                   c.MD.UserId,
		AddPhonePrivacyException: mtproto.ToBool(in.AddPhonePrivacyException),
		Id:                       id.PeerId,
		FirstName:                in.FirstName,
		LastName:                 in.LastName,
		Phone:                    in.Phone,
	})
	if err != nil {
		c.Logger.Errorf("contacts.addContact - error: %v", err)
		err = mtproto.ErrContactIdInvalid
		return nil, err
	}

	cUser, _ := users.GetUnsafeUser(c.MD.UserId, id.PeerId)
	cUser.Contact = true
	cUser.MutualContact = mtproto.FromBool(changeMutual)
	me, _ := users.GetUnsafeUserSelf(c.MD.UserId)

	// TODO(@benqi): 性能优化，复用users
	rUpdates := mtproto.MakeUpdatesByUpdatesUsers(
		[]*mtproto.User{me, cUser},
		mtproto.MakeTLUpdatePeerSettings(&mtproto.Update{
			Peer_PEER: id.ToPeer(),
			Settings: mtproto.MakeTLPeerSettings(&mtproto.PeerSettings{
				ReportSpam:            false,
				AddContact:            false,
				BlockContact:          false,
				ShareContact:          false,
				NeedContactsException: false,
				ReportGeo:             false,
			}).To_PeerSettings(),
		}).To_Update())

	return rUpdates, nil
}
