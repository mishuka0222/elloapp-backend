package core

import (
	msgpb "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/messenger/msg/msg/msg"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/messenger/sync/sync"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/channels/channels"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/dialog/dialog"
	"math/rand"
	"time"

	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

func (c *ChannelsCore) ChannelsEditBanned(in *mtproto.TLChannelsEditBanned) (upd *mtproto.Updates, err error) {
	var (
		peer = mtproto.MakeChannelPeerUtil(in.Channel.GetChannelId())
	)

	_, err = c.svcCtx.Dao.ChannelsClient.DeleteChannelUser(c.ctx, &channels.DeleteChannelUserReq{
		ChannelId:    in.Channel.GetChannelId(),
		OperatorId:   c.MD.UserId,
		DeleteUserId: in.Participant.UserId,
		Reason:       channels.K_ChannelParticipantBanned,
		BannedRights: in.BannedRights,
	})
	if err != nil {
		return
	}

	_, err = c.svcCtx.Dao.DialogClient.DialogDeleteDialog(c.ctx, &dialog.TLDialogDeleteDialog{
		UserId:   c.MD.UserId,
		PeerType: peer.PeerType,
		PeerId:   peer.PeerId,
	})
	if err != nil {
		return
	}

	upd, err = c.svcCtx.Dao.MsgClient.MsgSendMessage(c.ctx, &msgpb.TLMsgSendMessage{
		UserId:    c.MD.UserId,
		AuthKeyId: c.MD.AuthId,
		PeerType:  peer.PeerType,
		PeerId:    peer.PeerId,
		Message: msgpb.MakeTLOutboxMessage(&msgpb.OutboxMessage{
			NoWebpage:    true,
			Background:   false,
			RandomId:     rand.Int63(),
			Message:      channels.MakeMessageService(peer.PeerId, c.MD.UserId, int32(time.Now().Unix()), mtproto.MakeMessageActionChatDeleteUser(c.MD.UserId)),
			ScheduleDate: nil,
		}).To_OutboxMessage(),
	})
	if err != nil {
		c.Logger.Errorf("messages.deleteChannelUser - error: %v", err)
		return
	}
	// Delete chat for banned user
	{
		pushUpdates := mtproto.MakeUpdatesByUpdatesChats(
			[]*mtproto.Chat{
				mtproto.MakeTLChannelForbidden(&mtproto.Chat{
					Id:    in.Channel.GetChannelId(),
					Title: "",
				}).To_Chat(),
			},
			mtproto.MakeUpdateChannel(peer.PeerId))
		c.svcCtx.Dao.DialogClient.DialogDeleteDialog(c.ctx, &dialog.TLDialogDeleteDialog{
			UserId:   in.Participant.UserId,
			PeerType: peer.PeerType,
			PeerId:   peer.PeerId,
		})

		c.svcCtx.Dao.SyncClient.SyncPushUpdates(c.ctx, &sync.TLSyncPushUpdates{
			UserId:  in.Participant.UserId,
			Updates: pushUpdates,
		})

		c.svcCtx.Dao.MsgClient.MsgDeleteHistory(c.ctx, &msgpb.TLMsgDeleteHistory{
			UserId:    in.Participant.UserId,
			AuthKeyId: 0,
			PeerType:  peer.PeerType,
			PeerId:    peer.PeerId,
			JustClear: false,
			Revoke:    false,
		})
	}

	return
}
