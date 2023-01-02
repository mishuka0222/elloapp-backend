package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/messenger/sync/sync"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
)

// SyncUpdatesNotMe
// sync.updatesNotMe user_id:long auth_key_id:long updates:Updates = Void;
func (c *SyncCore) SyncUpdatesNotMe(in *sync.TLSyncUpdatesNotMe) (*mtproto.Void, error) {
	var (
		userId    = in.GetUserId()
		authKeyId = in.GetAuthKeyId()
		updates   = in.GetUpdates()
	)

	notification, err := c.processUpdates(syncTypeUserNotMe, userId, false, updates)
	if err != nil {
		c.Logger.Errorf("sync.updatesNotMe - error: %v", err)
		return nil, err
	}

	c.pushUpdatesToSession(syncTypeUserNotMe, userId, authKeyId, 0, updates, "", notification)

	return mtproto.EmptyVoid, nil
}
