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
	"context"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/messenger/msg/inbox/inbox"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/messenger/sync/sync"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
)

// InboxDeleteMessagesToInbox
// inbox.deleteMessagesToInbox from_id:long id:Vector<int> = Void;
func (c *InboxCore) InboxDeleteMessagesToInbox(in *inbox.TLInboxDeleteMessagesToInbox) (*mtproto.Void, error) {
	c.svcCtx.Dao.DeleteInboxMessages(
		c.ctx,
		in.FromId,
		mtproto.MakePeerUtil(in.PeerType, in.PeerId),
		in.Id,
		func(ctx context.Context, userId int64, idList []int32) {
			c.svcCtx.Dao.SyncClient.SyncPushUpdates(ctx, &sync.TLSyncPushUpdates{
				UserId: userId,
				Updates: mtproto.MakeUpdatesByUpdates(mtproto.MakeTLUpdateDeleteMessages(&mtproto.Update{
					Messages:  idList,
					Pts_INT32: c.svcCtx.Dao.IDGenClient2.NextNPtsId(ctx, userId, len(idList)),
					PtsCount:  int32(len(idList)),
				}).To_Update()),
			})
		})
	return mtproto.EmptyVoid, nil
}
