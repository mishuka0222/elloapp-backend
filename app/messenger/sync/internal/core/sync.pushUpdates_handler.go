package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/messenger/sync/sync"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
)

// SyncPushUpdates
// sync.pushUpdates user_id:long updates:Updates = Void;
func (c *SyncCore) SyncPushUpdates(in *sync.TLSyncPushUpdates) (*mtproto.Void, error) {
	var (
		userId  = in.GetUserId()
		updates = in.GetUpdates()
	)

	notification, err := c.processUpdates(syncTypeUser, userId, false, updates)
	if err != nil {
		c.Logger.Errorf("sync.updatesNotMe - error: %v", err)
		return nil, err
	}

	c.pushUpdatesToSession(syncTypeUser, userId, 0, 0, updates, "", notification)

	return mtproto.EmptyVoid, nil
}
