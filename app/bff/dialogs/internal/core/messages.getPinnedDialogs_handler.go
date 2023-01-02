package core

import (
	"context"
	"sort"

	chatpb "gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/chat/chat"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/dialog/dialog"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/message/message"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/updates/updates"
	userpb "gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/user/user"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"

	"github.com/zeromicro/go-zero/core/mr"
)

// MessagesGetPinnedDialogs
// messages.getPinnedDialogs#d6b94df2 folder_id:int = messages.PeerDialogs;
func (c *DialogsCore) MessagesGetPinnedDialogs(in *mtproto.TLMessagesGetPinnedDialogs) (*mtproto.Messages_PeerDialogs, error) {
	var (
		err      error
		folderId = in.GetFolderId()
	)

	if folderId != 0 && folderId != 1 {
		err = mtproto.ErrFolderIdInvalid
		c.Logger.Errorf("messages.getPinnedDialogs - error: %v", err)
		return nil, err
	}

	var (
		peers              []*mtproto.PeerUtil
		dialogExtList      dialog.DialogExtList
		state              *mtproto.Updates_State
		notifySettingsList []*userpb.PeerPeerNotifySettings
	)

	err = mr.Finish(
		func() error {
			state, err = c.svcCtx.Dao.UpdatesClient.UpdatesGetState(c.ctx, &updates.TLUpdatesGetState{
				AuthKeyId: c.MD.AuthId,
				UserId:    c.MD.UserId,
			})
			if err != nil {
				c.Logger.Errorf("messages.getPinnedDialogs - error: %v", err)
				return mtproto.ErrInternelServerError
			}

			return nil
		},
		func() error {
			if folderId == 0 {
				dialogFolder, err2 := c.svcCtx.Dao.DialogClient.DialogGetDialogFolder(c.ctx, &dialog.TLDialogGetDialogFolder{
					UserId:   c.MD.UserId,
					FolderId: 1,
				})
				if err2 != nil {
					c.Logger.Errorf("messages.getPinnedDialogs - error: %v", err2)
					return err2
				}

				dialogExtList = append(dialogExtList, dialogFolder.GetDatas()...)
			}

			return nil
		},
		func() error {
			dList, err2 := c.svcCtx.Dao.DialogClient.DialogGetPinnedDialogs(c.ctx, &dialog.TLDialogGetPinnedDialogs{
				UserId:   c.MD.UserId,
				FolderId: folderId,
			})
			if err2 != nil {
				c.Logger.Errorf("messages.getPinnedDialogs - error: %v", err2)
				return err2
			} else {
				dialogExtList = append(dialogExtList, dList.GetDatas()...)
			}

			return nil
		},
	)

	for _, dialogEx := range dialogExtList {
		peer2 := mtproto.FromPeer(dialogEx.GetDialog().GetPeer())
		peers = append(peers, peer2)
		if peer2.IsChannel() {
			c.Logger.Errorf("blocked, License key from https://elloapp.com required to unlock enterprise features.")
		}
	}

	if len(peers) > 0 {
		settingsList, err2 := c.svcCtx.Dao.UserClient.UserGetNotifySettingsList(c.ctx, &userpb.TLUserGetNotifySettingsList{
			UserId: c.MD.UserId,
			Peers:  peers,
		})
		if err2 != nil {
			c.Logger.Errorf("messages.getDialogs - error: %v", err2)
			return nil, err2
		}

		notifySettingsList = settingsList.GetDatas()
	}

	for _, dialogEx := range dialogExtList {
		peer2 := mtproto.FromPeer(dialogEx.GetDialog().GetPeer())
		dialogEx.Dialog.NotifySettings = userpb.FindPeerPeerNotifySettings(notifySettingsList, peer2)
		if peer2.IsChannel() {
			c.Logger.Errorf("blocked, License key from https://elloapp.com required to unlock enterprise features.")
		}
	}

	sort.Sort(sort.Reverse(dialogExtList))

	messageDialogs := dialogExtList.DoGetMessagesDialogs(
		c.ctx,
		c.MD.UserId,
		func(ctx context.Context, selfUserId int64, id ...dialog.TopMessageId) []*mtproto.Message {
			var (
				msgList   = make([]*mtproto.Message, 0, len(id))
				msgIdList = make([]int32, 0, len(id))
			)
			for _, id2 := range id {
				if id2.Peer.IsChannel() {
					c.Logger.Errorf("blocked, License key from https://elloapp.com required to unlock enterprise features.")
				} else {
					msgIdList = append(msgIdList, id2.TopMessage)
				}
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
			c.Logger.Errorf("blocked, License key from https://elloapp.com required to unlock enterprise features.")
			return []*mtproto.Chat{}
		})

	return messageDialogs.ToMessagesPeerDialogs(state), nil
}
