package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/interface/session/session"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/messenger/sync/sync"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
)

// SyncPushRpcResult
// sync.pushRpcResult server_id:long auth_key_id:long req_msg_id:long result:bytes = PushUpdates;
func (c *SyncCore) SyncPushRpcResult(in *sync.TLSyncPushRpcResult) (*mtproto.Void, error) {
	c.svcCtx.Dao.PushRpcResultToSession(c.ctx, in.ServerId, &session.TLSessionPushRpcResultData{
		AuthKeyId:      in.AuthKeyId,
		SessionId:      in.SessionId,
		ClientReqMsgId: in.ClientReqMsgId,
		RpcResultData:  in.RpcResult,
	})

	return mtproto.EmptyVoid, nil
}
