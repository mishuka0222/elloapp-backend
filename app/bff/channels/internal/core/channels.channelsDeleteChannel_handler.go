package core

import (
	msgpb "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/messenger/msg/msg/msg"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/messenger/sync/sync"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/channels/channels"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/dialog/dialog"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/username/username"

	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

func (c *ChannelsCore) ChannelsDeleteChannel(in *mtproto.TLChannelsDeleteChannel) (upd *mtproto.Updates, err error) {
	var (
		channel *channels.ChannelData
		peer    = mtproto.MakeChannelPeerUtil(in.Channel.GetChannelId())
	)
	channel, err = c.svcCtx.Dao.ChannelsClient.DeleteChannel(c.ctx, &channels.DeleteChannelReq{
		ChannelId:  peer.PeerId,
		OperatorId: c.MD.UserId,
	})
	if err != nil {
		c.Logger.Errorf("channels.deleteChannel - error: %v", err)
		return nil, err
	}

	_, err = c.svcCtx.Dao.UsernameClient.UsernameDeleteUsername(c.ctx, &username.TLUsernameDeleteUsername{
		Username: channel.Channel.Username,
	})
	if err != nil {
		c.Logger.Errorf("channels.deleteChannel - error: %v", err)
		err = nil
	}

	pushUpdates := mtproto.MakeUpdatesByUpdatesChats(
		[]*mtproto.Chat{channel.ToChannelForbidden()},
		mtproto.MakeUpdateChannel(peer.PeerId))

	// 1. kicked all
	for _, cp := range channel.Participants {
		c.svcCtx.Dao.DialogClient.DialogDeleteDialog(c.ctx, &dialog.TLDialogDeleteDialog{
			UserId:   cp.UserId,
			PeerType: peer.PeerType,
			PeerId:   peer.PeerId,
		})

		if cp.UserId == c.MD.UserId || cp.IsChatMemberStateNormal() {
			c.svcCtx.Dao.SyncClient.SyncPushUpdates(c.ctx, &sync.TLSyncPushUpdates{
				UserId:  cp.UserId,
				Updates: pushUpdates,
			})
		}

		c.svcCtx.Dao.MsgClient.MsgDeleteHistory(c.ctx, &msgpb.TLMsgDeleteHistory{
			UserId:    cp.UserId,
			AuthKeyId: 0,
			PeerType:  peer.PeerType,
			PeerId:    peer.PeerId,
			JustClear: false,
			Revoke:    false,
		})

	}

	upd = mtproto.MakeEmptyUpdates()

	return
}
