package core

import (
	"context"

	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/messenger/msg/inbox/internal/svc"
	chatpb "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/chat/chat"
	userpb "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/user/user"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto/rpc/metadata"

	"github.com/zeromicro/go-zero/core/logx"
)

type InboxCore struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	MD *metadata.RpcMetadata
}

func New(ctx context.Context, svcCtx *svc.ServiceContext) *InboxCore {
	return &InboxCore{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		MD:     metadata.RpcMetadataFromIncoming(ctx),
	}
}

func (c *InboxCore) makeUpdateNewMessageListUpdates(selfUserId int64, boxList ...*mtproto.MessageBox) *mtproto.Updates {
	var (
		messages      = make([]*mtproto.Message, 0, len(boxList))
		updateNewList = make([]*mtproto.Update, 0, len(boxList))
	)

	for _, box := range boxList {
		if box == nil {
			continue
		}
		if box.PeerType == mtproto.PEER_CHANNEL {
			m := box.ToMessage(selfUserId)
			messages = append(messages, m)
			updateNewMessage := mtproto.MakeTLUpdateNewChannelMessage(&mtproto.Update{
				Message_MESSAGE: box.ToMessage(selfUserId),
				Pts_INT32:       box.Pts,
				PtsCount:        box.PtsCount,
			})
			updateNewList = append(updateNewList, updateNewMessage.To_Update())
		} else {
			m := box.ToMessage(selfUserId)
			messages = append(messages, m)
			updateNewMessage := mtproto.MakeTLUpdateNewMessage(&mtproto.Update{
				Message_MESSAGE: box.ToMessage(selfUserId),
				Pts_INT32:       box.Pts,
				PtsCount:        box.PtsCount,
			})
			updateNewList = append(updateNewList, updateNewMessage.To_Update())
		}
	}

	rUpdates := mtproto.MakeUpdatesByUpdates(updateNewList...)

	idHelper := mtproto.NewIDListHelper(selfUserId)
	idHelper.PickByMessages(messages...)
	idHelper.Visit(
		func(userIdList []int64) {
			users, _ := c.svcCtx.Dao.UserClient.UserGetMutableUsers(c.ctx,
				&userpb.TLUserGetMutableUsers{
					Id: userIdList,
				})
			rUpdates.PushUser(users.GetUserListByIdList(selfUserId, userIdList...)...)
		},
		func(chatIdList []int64) {
			chats, _ := c.svcCtx.Dao.ChatClient.ChatGetChatListByIdList(c.ctx,
				&chatpb.TLChatGetChatListByIdList{
					IdList: chatIdList,
				})
			rUpdates.PushChat(chats.GetChatListByIdList(selfUserId, chatIdList...)...)
		},
		func(channelIdList []int64) {
			// TODO
		})

	return rUpdates
}
