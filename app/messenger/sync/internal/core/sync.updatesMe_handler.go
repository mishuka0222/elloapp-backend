package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/messenger/sync/sync"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
)

// SyncUpdatesMe
// sync.updatesMe flags:# user_id:long auth_key_id:long server_id:string session_id:flags.0?long updates:Updates = Void;
func (c *SyncCore) SyncUpdatesMe(in *sync.TLSyncUpdatesMe) (*mtproto.Void, error) {
	c.pushUpdatesToSession(syncTypeUserMe,
		in.GetUserId(),
		in.GetAuthKeyId(),
		in.GetSessionId().GetValue(),
		in.GetUpdates(),
		in.GetServerId(),
		false)

	return mtproto.EmptyVoid, nil
}
