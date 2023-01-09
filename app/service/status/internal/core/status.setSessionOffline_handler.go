package core

import (
	"strconv"

	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/status/status"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

// StatusSetSessionOffline
// status.setSessionOffline online:SessionEntry = Bool;
func (c *StatusCore) StatusSetSessionOffline(in *status.TLStatusSetSessionOffline) (*mtproto.Bool, error) {
	var (
		userK = getUserKey(in.GetUserId())
		authK = getAuthKeyIdKey(in.GetAuthKeyId())
	)

	if _, err := c.svcCtx.KV.Hdel(userK, strconv.FormatInt(in.GetAuthKeyId(), 10)); err != nil {
		c.Logger.Errorf("status.setSessionOffline(%s) error(%v)", in.DebugString(), err)
		return nil, err
	}

	if _, err := c.svcCtx.KV.Del(authK); err != nil {
		c.Logger.Errorf("status.setSessionOffline(%s) error(%v)", in.DebugString(), err)
		return nil, err
	}

	return mtproto.BoolTrue, nil
}
