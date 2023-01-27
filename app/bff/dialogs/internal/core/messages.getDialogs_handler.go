package core

import (
	"context"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/channels/channels"
	"sort"

	chatpb "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/chat/chat"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/dialog/dialog"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/message/message"
	userpb "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/user/user"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"

	"github.com/zeromicro/go-zero/core/mr"
)

// MessagesGetDialogs
// messages.getDialogs#a0f4cb4f flags:# exclude_pinned:flags.0?true folder_id:flags.1?int offset_date:int offset_id:int offset_peer:InputPeer limit:int hash:long = messages.Dialogs;
func (c *DialogsCore) MessagesGetDialogs(in *mtproto.TLMessagesGetDialogs) (*mtproto.Messages_Dialogs, error) {
	var (
		offsetPeer         = mtproto.FromInputPeer2(c.MD.UserId, in.OffsetPeer)
		folderId           = in.GetFolderId().GetValue()
		limit              = in.Limit
		dialogExtList      dialog.DialogExtList
		notifySettingsList []*userpb.PeerPeerNotifySettings
	)

	if limit > 500 {
		limit = 500
	}

	mr.FinishVoid(
		func() {
			dialogs, err := c.svcCtx.Dao.DialogClient.DialogGetDialogs(c.ctx, &dialog.TLDialogGetDialogs{
				UserId:        c.MD.UserId,
				ExcludePinned: mtproto.ToBool(in.ExcludePinned),
				FolderId:      folderId,
			})
			if err != nil {
				c.Logger.Errorf("messages.getDialogs - error: %v", err)
			} else {
				dialogExtList = dialogs.GetDatas()
			}
		},
		func() {
			settingsList, err := c.svcCtx.Dao.UserClient.UserGetAllNotifySettings(c.ctx, &userpb.TLUserGetAllNotifySettings{
				UserId: c.MD.UserId,
			})
			if err != nil {
				c.Logger.Errorf("messages.getDialogs - error: %v", err)
			} else {
				notifySettingsList = settingsList.GetDatas()
			}
		})

	if len(dialogExtList) == 0 {
		return mtproto.MakeTLMessagesDialogsSlice(&mtproto.Messages_Dialogs{
			Dialogs:  []*mtproto.Dialog{},
			Messages: []*mtproto.Message{},
			Chats:    []*mtproto.Chat{},
			Users:    []*mtproto.User{},
			Count:    0,
		}).To_Messages_Dialogs(), nil
	}

	var (
		dialogCount = int32(dialogExtList.Len())
	)

	for _, dialogEx := range dialogExtList {
		peer2 := mtproto.FromPeer(dialogEx.GetDialog().GetPeer())

		//if peer2.IsChannel() {
		//	c.Logger.Errorf("messages.getDialogs blocked, License key from https://elloapp.com required to unlock enterprise features.")
		//}

		dialogEx.Dialog.NotifySettings = userpb.FindPeerPeerNotifySettings(notifySettingsList, peer2)
	}

	sort.Sort(sort.Reverse(dialogExtList))

	dialogExtList = dialogExtList.GetDialogsByOffsetLimit(
		in.OffsetDate,
		in.OffsetId,
		offsetPeer,
		in.Limit)

	messageDialogs := dialogExtList.DoGetMessagesDialogs(
		c.ctx,
		c.MD.UserId,
		func(ctx context.Context, selfUserId int64, id ...dialog.TopMessageId) []*mtproto.Message {
			var (
				msgList   = make([]*mtproto.Message, 0, len(id))
				msgIdList = make([]int32, 0, len(id))
			)
			for _, id2 := range id {
				//if !id2.Peer.IsChannel() {
				msgIdList = append(msgIdList, id2.TopMessage)
				//} else {
				//	c.Logger.Errorf("blocked, License key from https://elloapp.com required to unlock enterprise features.")
				//}
			}
			if len(msgIdList) > 0 {
				boxList, _ := c.svcCtx.Dao.MessageClient.MessageGetUserMessageList(c.ctx, &message.TLMessageGetUserMessageList{
					UserId: c.MD.UserId,
					IdList: msgIdList,
				})
				boxList.Walk(func(idx int, v *mtproto.MessageBox) {
					msgList = append(msgList, v.ToMessage(c.MD.UserId))
				})
			}

			return msgList
		},
		func(ctx context.Context, selfUserId int64, id ...int64) []*mtproto.User {
			users, _ := c.svcCtx.Dao.UserClient.UserGetMutableUsers(c.ctx,
				&userpb.TLUserGetMutableUsers{
					Id: id,
				})

			return users.GetUserListByIdList(c.MD.UserId, id...)
		},
		func(ctx context.Context, selfUserId int64, id ...int64) []*mtproto.Chat {
			chats, _ := c.svcCtx.Dao.ChatClient.ChatGetChatListByIdList(c.ctx,
				&chatpb.TLChatGetChatListByIdList{
					IdList: id,
				})

			return chats.GetChatListByIdList(c.MD.UserId, id...)
		},
		func(ctx context.Context, selfUserId int64, id ...int64) []*mtproto.Chat {
			res, _ := c.svcCtx.Dao.ChannelsClient.GetChatsListBySelfAndIDList(c.ctx, &channels.GetChatsListBySelfAndIDListReq{
				SelfUserId: selfUserId,
				IdList:     id,
			})
			return res.Chats
		})

	return messageDialogs.ToMessagesDialogs(dialogCount), nil
}
