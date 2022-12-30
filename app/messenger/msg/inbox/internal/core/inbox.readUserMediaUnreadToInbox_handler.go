// Copyright 2022 Teamgram Authors
//  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Author: teamgramio (teamgram.io@gmail.com)
//

package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/messenger/msg/inbox/inbox"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/messenger/msg/internal/dal/dataobject"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/messenger/sync/sync"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
)

// InboxReadUserMediaUnreadToInbox
// inbox.readUserMediaUnreadToInbox from_id:long id:Vector<int> = Void;
func (c *InboxCore) InboxReadUserMediaUnreadToInbox(in *inbox.TLInboxReadUserMediaUnreadToInbox) (*mtproto.Void, error) {
	idList := make([]int64, 0, len(in.GetId()))
	for _, id := range in.GetId() {
		idList = append(idList, id.DialogMessageId)
	}

	c.svcCtx.Dao.MessagesDAO.SelectByMessageDataIdListWithCB(
		c.ctx,
		c.svcCtx.Dao.MessagesDAO.CalcTableName(in.PeerUserId),
		idList,
		func(i int, v *dataobject.MessagesDO) {
			c.svcCtx.Dao.MessagesDAO.UpdateMediaUnread(c.ctx, v.UserId, v.UserMessageBoxId)

			// TODO: batch handle
			pts := c.svcCtx.Dao.IDGenClient2.NextPtsId(c.ctx, v.UserId)
			c.svcCtx.Dao.SyncClient.SyncPushUpdates(c.ctx, &sync.TLSyncPushUpdates{
				UserId: v.UserId,
				Updates: mtproto.MakeUpdatesByUpdates(mtproto.MakeTLUpdateReadMessagesContents(&mtproto.Update{
					Messages:  []int32{v.UserMessageBoxId},
					Pts_INT32: pts,
					PtsCount:  1,
				}).To_Update()),
			})
		},
	)

	return mtproto.EmptyVoid, nil
}
