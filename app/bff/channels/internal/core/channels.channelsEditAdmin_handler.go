package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/messenger/sync/sync"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/channels/channels"
	userpb "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/user/user"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

func (c *ChannelsCore) ChannelsEditAdmin(in *mtproto.TLChannelsEditAdmin) (upd *mtproto.Updates, err error) {

	// todo: add your logic here and delete this line
	//return nil, errors.New("Unimplemented")
	// return &mtproto.Updates{}, nil

	var (
		peer                   = mtproto.MakeChannelPeerUtil(in.Channel.GetChannelId())
		adminUser              = mtproto.FromInputUser(c.MD.UserId, in.UserId)
		channel                *channels.ChannelData
		participants           *mtproto.ChatParticipants
		updateChatParticipants *mtproto.Update
		adminRights            = in.GetAdminRights()
		mUsers                 *userpb.Vector_ImmutableUser
		idList                 []int64
	)

	if adminUser.PeerType != mtproto.PEER_USER ||
		adminUser.PeerId == c.MD.UserId {
		err = mtproto.ErrUserIdInvalid
		c.Logger.Errorf("messages.editChatAdmin - invalid user_id, err: %v", err)
		return
	}

	if adminRights != nil {
		if !adminRights.ChangeInfo &&
			!adminRights.PostMessages &&
			!adminRights.EditMessages &&
			!adminRights.DeleteMessages &&
			!adminRights.BanUsers &&
			!adminRights.InviteUsers &&
			!adminRights.PinMessages &&
			!adminRights.AddAdmins &&
			!adminRights.Anonymous &&
			!adminRights.ManageCall &&
			!adminRights.Other {
			adminRights = nil
		}
	}

	channel, err = c.svcCtx.Dao.ChannelsClient.EditChannelAdmin(c.ctx, &channels.EditChannelAdminReq{
		ChannelId:   peer.PeerId,
		OperatorId:  c.MD.UserId,
		EditUserId:  in.UserId.GetUserId(),
		AdminRights: adminRights,
	})
	if err != nil {
		return
	}
	participants, idList = channel.GetNormalParticipants()
	updateChatParticipants = mtproto.MakeTLUpdateChatParticipants(&mtproto.Update{
		Participants_CHATPARTICIPANTS: participants,
	}).To_Update()

	mUsers, err = c.svcCtx.Dao.UserClient.UserGetMutableUsers(c.ctx, &userpb.TLUserGetMutableUsers{
		Id: idList,
	})
	if err != nil {
		c.Logger.Errorf("messages.getFullChat - error: not found dialog")
	}

	for _, cp := range channel.Participants {
		rChat, _ := c.svcCtx.Dao.ChannelsClient.ToChat(c.ctx, &channels.ToChatReq{Channel: channel, SelfUserId: cp.ChannelId})

		c.svcCtx.Dao.SyncClient.SyncPushUpdates(c.ctx,
			&sync.TLSyncPushUpdates{
				UserId: cp.UserId,
				Updates: mtproto.MakeUpdatesByUpdatesUsersChats(
					mUsers.GetUserListByIdList(cp.UserId, idList...),
					[]*mtproto.Chat{rChat.Chat},
					updateChatParticipants),
			})
	}

	upd = mtproto.MakeEmptyUpdates()

	return
}
