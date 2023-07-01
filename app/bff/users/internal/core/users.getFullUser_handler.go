package core

import (
	chatpb "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/chat/chat"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/dialog/dialog"
	userpb "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/user/user"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/pkg2/utils"

	"github.com/zeromicro/go-zero/core/mr"
)

// UsersGetFullUser
// users.getFullUser#b60f5918 id:InputUser = users.UserFull;
func (c *UsersCore) UsersGetFullUser(in *mtproto.TLUsersGetFullUser) (*mtproto.Users_UserFull, error) {
	var (
		peerId int64
		id     = mtproto.FromInputUser(c.MD.UserId, in.Id)
		me     *userpb.ImmutableUser
		user   *userpb.ImmutableUser
	)

	switch id.PeerType {
	case mtproto.PEER_SELF, mtproto.PEER_USER:
		peerId = id.PeerId
	default:
		err := mtproto.ErrUserIdInvalid
		c.Logger.Errorf("users.getFullUser - error: %v", err)
		return nil, err
	}

	mutableUsers, err := c.svcCtx.Dao.UserClient.UserGetMutableUsers(c.ctx, &userpb.TLUserGetMutableUsers{
		Id: []int64{c.MD.UserId, peerId},
	})
	if err != nil {
		c.Logger.Errorf("users.getFullUser - error: %v", err)
		return nil, err
	} else if len(mutableUsers.Datas) == 0 {
		err = mtproto.ErrInternelServerError
		c.Logger.Errorf("users.getFullUser - error: %v", err)
		return nil, err
	}

	me, _ = mutableUsers.GetImmutableUser(c.MD.UserId)
	user, _ = mutableUsers.GetImmutableUser(peerId)

	if user == nil {
		err = mtproto.ErrInternelServerError
		c.Logger.Errorf("users.getFullUser - error: %v", err)
		return nil, err
	}

	// Layer135
	// userFull#cf366521 flags:#
	//	blocked:flags.0?true
	//	phone_calls_available:flags.4?true
	//	phone_calls_private:flags.5?true
	//	can_pin_message:flags.7?true
	//	has_scheduled:flags.12?true
	//	video_calls_available:flags.13?true
	//	id:long
	//	about:flags.1?string
	//	settings:PeerSettings
	//	profile_photo:flags.2?Photo
	//	notify_settings:PeerNotifySettings
	//	bot_info:flags.3?BotInfo
	//	pinned_msg_id:flags.6?int
	//	common_chats_count:int
	//	folder_id:flags.11?int
	//	ttl_period:flags.14?int
	//	theme_emoticon:flags.15?string
	//	private_forward_name:flags.16?string = UserFull;
	userFull := mtproto.MakeTLUserFull(&mtproto.UserFull{
		Blocked:             false,
		PhoneCallsAvailable: true,
		PhoneCallsPrivate:   false,
		CanPinMessage:       true,
		HasScheduled:        false,
		VideoCallsAvailable: true,
		Id:                  peerId,
		About:               user.GetUser().GetAbout(),
		Settings:            nil,
		ProfilePhoto:        user.GetUser().GetProfilePhoto(),
		NotifySettings:      nil,
		BotInfo:             nil,
		PinnedMsgId:         nil,
		CommonChatsCount:    0,
		FolderId:            nil,
		TtlPeriod:           nil,
		ThemeEmoticon:       nil,
		PrivateForwardName:  nil,
	}).To_UserFull()

	mr.FinishVoid(
		func() {
			// blocked
			if c.MD.UserId != peerId {
				blocked, _ := c.svcCtx.Dao.UserClient.UserBlockedByUser(
					c.ctx,
					&userpb.TLUserBlockedByUser{
						UserId:     c.MD.UserId,
						PeerUserId: peerId,
					})
				userFull.Blocked = mtproto.FromBool(blocked)
			}
		},
		func() {
			userFull.Settings, _ = c.svcCtx.Dao.UserClient.UserGetPeerSettings(c.ctx, &userpb.TLUserGetPeerSettings{
				UserId:   c.MD.UserId,
				PeerType: mtproto.PEER_USER,
				PeerId:   peerId,
			})
		},
		func() {
			userFull.NotifySettings, _ = c.svcCtx.Dao.UserClient.UserGetNotifySettings(c.ctx, &userpb.TLUserGetNotifySettings{
				UserId:   c.MD.UserId,
				PeerType: mtproto.PEER_USER,
				PeerId:   peerId,
			})
		},
		func() {
			if user.GetUser().GetBot() != nil {
				userFull.PhoneCallsAvailable = false
				userFull.PhoneCallsPrivate = false
				userFull.BotInfo, _ = c.svcCtx.Dao.UserClient.UserGetBotInfo(c.ctx, &userpb.TLUserGetBotInfo{
					BotId: peerId,
				})
			}
		},
		func() {
			// TODO: PinnedMsgId:         nil,
			if c.MD.UserId != peerId {
				usersChatIdList, _ := c.svcCtx.Dao.ChatClient.ChatGetUsersChatIdList(c.ctx, &chatpb.TLChatGetUsersChatIdList{
					Id: []int64{c.MD.UserId, peerId},
				})
				if len(usersChatIdList.Datas) == 2 {
					commonChats := utils.Int64Intersect(
						usersChatIdList.Datas[0].ChatIdList,
						usersChatIdList.Datas[1].ChatIdList)
					userFull.CommonChatsCount = int32(len(commonChats))
				}

				// TODO: Fetch CommonChannelsCount
			}
		},
		func() {
			if peerId != c.MD.UserId {
				// theme_emoticon
				dialogExt, _ := c.svcCtx.Dao.DialogClient.DialogGetDialogById(c.ctx, &dialog.TLDialogGetDialogById{
					UserId:   c.MD.UserId,
					PeerType: mtproto.PEER_USER,
					PeerId:   peerId,
				})
				if dialogExt != nil {
					userFull.ThemeEmoticon = mtproto.MakeFlagsString(dialogExt.ThemeEmoticon)
					userFull.TtlPeriod = mtproto.MakeFlagsInt32(dialogExt.TtlPeriod)
				}
			}
		})

	// TODO: FolderId:    0,

	return mtproto.MakeTLUsersUserFull(&mtproto.Users_UserFull{
		FullUser: userFull,
		Chats:    []*mtproto.Chat{},
		Users:    []*mtproto.User{user.ToUnsafeUser(me)},
	}).To_Users_UserFull(), nil
}
