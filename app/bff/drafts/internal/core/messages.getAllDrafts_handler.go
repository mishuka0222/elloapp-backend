package core

import (
	chatpb "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/chat/chat"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/dialog/dialog"
	userpb "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/user/user"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

// MessagesGetAllDrafts
// messages.getAllDrafts#6a3f8d65 = Updates;
func (c *DraftsCore) MessagesGetAllDrafts(in *mtproto.TLMessagesGetAllDrafts) (*mtproto.Updates, error) {
	_ = in

	drafts, err := c.svcCtx.Dao.DialogClient.DialogGetAllDrafts(c.ctx, &dialog.TLDialogGetAllDrafts{
		UserId: c.MD.UserId,
	})
	if err != nil {
		c.Logger.Errorf("messages.getAllDrafts - error: %v", err)
		return nil, err
	}

	var (
		rUpdates = mtproto.MakeEmptyUpdates()
		idHelper = mtproto.NewIDListHelper(c.MD.UserId)
	)

	for _, v := range drafts.Datas {
		rUpdates.PushBackUpdate(mtproto.MakeTLUpdateDraftMessage(&mtproto.Update{
			Peer_PEER: v.Peer,
			Draft:     v.Draft,
		}).To_Update())

		idHelper.PickByPeer(v.Peer)
	}

	idHelper.Visit(
		func(userIdList []int64) {
			users, _ := c.svcCtx.Dao.UserClient.UserGetMutableUsers(c.ctx,
				&userpb.TLUserGetMutableUsers{
					Id: userIdList,
				})
			rUpdates.PushUser(users.GetUserListByIdList(c.MD.UserId, userIdList...)...)
		},
		func(chatIdList []int64) {
			chats, _ := c.svcCtx.Dao.ChatClient.ChatGetChatListByIdList(c.ctx,
				&chatpb.TLChatGetChatListByIdList{
					IdList: chatIdList,
				})
			rUpdates.PushChat(chats.GetChatListByIdList(c.MD.UserId, chatIdList...)...)
		},
		func(channelIdList []int64) {
			if c.svcCtx.Plugin != nil {
				chats := c.svcCtx.Plugin.GetChannelListByIdList(c.ctx, c.MD.UserId, channelIdList...)
				rUpdates.PushChat(chats...)
			} else {
				c.Logger.Errorf("messages.getAllDrafts blocked, License key from https://elloapp.com required to unlock enterprise features.")
			}
		})

	return rUpdates, nil
}
