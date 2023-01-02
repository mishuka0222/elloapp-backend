package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/messenger/sync/sync"
	chatpb "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/chat/chat"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/dialog/dialog"
	userpb "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/user/user"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

// MessagesToggleDialogPin
// messages.toggleDialogPin#a731e257 flags:# pinned:flags.0?true peer:InputDialogPeer = Bool;
func (c *DialogsCore) MessagesToggleDialogPin(in *mtproto.TLMessagesToggleDialogPin) (*mtproto.Bool, error) {
	var (
		peer *mtproto.PeerUtil
	)

	switch in.GetPeer().GetPredicateName() {
	case mtproto.Predicate_inputDialogPeer:
		peer = mtproto.FromInputPeer2(c.MD.UserId, in.GetPeer().GetPeer())
	case mtproto.Predicate_inputDialogPeerFolder:
		// error
		c.Logger.Errorf("messages.toggleDialogPin - error: client not send inputDialogPeerFolder: %v", in.GetPeer())
		return mtproto.BoolFalse, nil
	default:
		err := mtproto.ErrInputConstructorInvalid
		c.Logger.Errorf("messages.toggleDialogPin - error: %v", err)
		return nil, err
	}

	folderId, err := c.svcCtx.Dao.DialogClient.DialogToggleDialogPin(c.ctx, &dialog.TLDialogToggleDialogPin{
		UserId:   c.MD.UserId,
		PeerType: peer.PeerType,
		PeerId:   peer.PeerId,
		Pinned:   mtproto.ToBool(in.Pinned),
	})
	if err != nil {
		c.Logger.Errorf("messages.toggleDialogPin - error: %v", err)
		return mtproto.BoolFalse, nil
	}

	var (
		idHelper    = mtproto.NewIDListHelper(c.MD.UserId)
		syncUpdates = mtproto.MakeUpdatesByUpdates(mtproto.MakeTLUpdateDialogPinned(&mtproto.Update{
			Pinned:   in.GetPinned(),
			FolderId: mtproto.MakeFlagsInt32(folderId.V),
			Peer_DIALOGPEER: mtproto.MakeTLDialogPeer(&mtproto.DialogPeer{
				Peer: peer.ToPeer(),
			}).To_DialogPeer(),
		}).To_Update())
	)

	idHelper.PickByPeerUtil(peer.PeerType, peer.PeerId)
	idHelper.Visit(
		func(userIdList []int64) {
			users, _ := c.svcCtx.Dao.UserClient.UserGetMutableUsers(c.ctx,
				&userpb.TLUserGetMutableUsers{
					Id: userIdList,
				})
			syncUpdates.PushUser(users.GetUserListByIdList(c.MD.UserId, userIdList...)...)
		},
		func(chatIdList []int64) {
			chats, _ := c.svcCtx.Dao.ChatClient.ChatGetChatListByIdList(c.ctx,
				&chatpb.TLChatGetChatListByIdList{
					IdList: chatIdList,
				})
			syncUpdates.PushChat(chats.GetChatListByIdList(c.MD.UserId, chatIdList...)...)
		},
		func(channelIdList []int64) {
			// TODO
		})
	c.svcCtx.Dao.SyncClient.SyncUpdatesNotMe(c.ctx, &sync.TLSyncUpdatesNotMe{
		UserId:    c.MD.UserId,
		AuthKeyId: c.MD.AuthId,
		Updates:   syncUpdates,
	})

	return mtproto.BoolTrue, nil
}
