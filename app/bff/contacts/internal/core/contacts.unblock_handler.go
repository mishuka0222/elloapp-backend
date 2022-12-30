package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/messenger/sync/sync"
	userpb "gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/user/user"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
)

// ContactsUnblock
// contacts.unblock#bea65d50 id:InputPeer = Bool;
func (c *ContactsCore) ContactsUnblock(in *mtproto.TLContactsUnblock) (*mtproto.Bool, error) {
	var (
		mUsers    *userpb.Vector_ImmutableUser
		err       error
		unblockId = mtproto.FromInputPeer2(c.MD.UserId, in.GetId())
		idHelper  = mtproto.NewIDListHelper(c.MD.UserId)
	)

	if !unblockId.IsUser() || unblockId.IsSelf() {
		err = mtproto.ErrPeerIdInvalid
		c.Logger.Errorf("contacts.block - error: %v", err)
		return nil, err
	}

	if unblockId.PeerId == c.MD.UserId {
		err = mtproto.ErrContactIdInvalid
		c.Logger.Errorf("contacts.block - error: %v", err)
		return nil, err
	}

	// TODO
	/*
		auto BlockPeerBoxController::createRow(not_null<History*> history)
		-> std::unique_ptr<BlockPeerBoxController::Row> {
			if (!history->peer->isUser()
				|| history->peer->isServiceUser()
				|| history->peer->isSelf()
				|| history->peer->isRepliesChat()) {
				return nullptr;
			}
			auto row = std::make_unique<Row>(history);
			updateIsBlocked(row.get(), history->peer);
			return row;
		}
	*/

	mUsers, err = c.svcCtx.Dao.UserClient.UserGetMutableUsers(c.ctx, &userpb.TLUserGetMutableUsers{
		Id: []int64{c.MD.UserId, unblockId.PeerId},
	})

	// me, _ := users.GetImmutableUser(c.MD.UserId)
	blocked, _ := mUsers.GetImmutableUser(unblockId.PeerId)

	if blocked == nil {
		err = mtproto.ErrContactIdInvalid
		c.Logger.Errorf("contacts.unblock - error: %v", err)
		return nil, err
	} else if blocked.GetUser().GetDeleted() {
		err = mtproto.ErrInputUserDeactivated
		c.Logger.Errorf("contacts.unblock - error: %v", err)
		return nil, err
	}
	c.svcCtx.Dao.UserClient.UserUnBlockPeer(c.ctx, &userpb.TLUserUnBlockPeer{
		UserId:   c.MD.UserId,
		PeerType: unblockId.PeerType,
		PeerId:   unblockId.PeerId,
	})
	idHelper.AppendUsers(unblockId.PeerId)

	syncUpdates := mtproto.MakeUpdatesByUpdates(mtproto.MakeTLUpdatePeerBlocked(&mtproto.Update{
		PeerId:  unblockId.ToPeer(),
		Blocked: mtproto.BoolFalse,
	}).To_Update())

	idHelper.Visit(
		func(userIdList []int64) {
			syncUpdates.PushUser(mUsers.GetUserListByIdList(c.MD.UserId, userIdList...)...)
		},
		func(chatIdList []int64) {
		},
		func(channelIdList []int64) {
		})

	c.svcCtx.Dao.SyncClient.SyncUpdatesNotMe(c.ctx, &sync.TLSyncUpdatesNotMe{
		UserId:    c.MD.UserId,
		AuthKeyId: c.MD.AuthId,
		Updates:   syncUpdates,
	})

	return mtproto.BoolTrue, nil
}
