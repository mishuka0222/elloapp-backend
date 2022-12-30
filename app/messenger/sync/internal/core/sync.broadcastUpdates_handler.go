/*
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package core

import (
	// "gitlab.com/merehead/elloapp/backend/elloapp_backend/pkg2/container2"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/messenger/sync/sync"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
	// channelpb "gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/channel/channel"
	chatpb "gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/chat/chat"
)

// SyncBroadcastUpdates
// sync.broadcastUpdates broadcast_type:int chat_id:long exclude_id_list:Vector<long> updates:Updates = Void;
func (c *SyncCore) SyncBroadcastUpdates(in *sync.TLSyncBroadcastUpdates) (*mtproto.Void, error) {
	pushUpdates := &sync.TLSyncPushUpdates{
		UserId:  0,
		Updates: in.Updates,
	}

	if in.BroadcastType != sync.BroadcastTypeChat {
		c.Logger.Errorf("invalid broadcast_type: %s", in.DebugString())
		return mtproto.EmptyVoid, nil
	}

	idList, _ := c.svcCtx.Dao.ChatClient.ChatGetChatParticipantIdList(c.ctx, &chatpb.TLChatGetChatParticipantIdList{
		ChatId: in.ChatId,
	})

	for _, id := range idList.GetDatas() {
		pushUpdates.UserId = id
		c.SyncPushUpdates(pushUpdates)
	}

	return mtproto.EmptyVoid, nil
}
