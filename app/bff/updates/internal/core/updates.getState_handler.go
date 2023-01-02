package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/updates/updates"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
)

// UpdatesGetState
// updates.getState#edd4882a = updates.State;
func (c *UpdatesCore) UpdatesGetState(in *mtproto.TLUpdatesGetState) (*mtproto.Updates_State, error) {
	rValue, err := c.svcCtx.Dao.UpdatesClient.UpdatesGetState(c.ctx, &updates.TLUpdatesGetState{
		AuthKeyId: c.MD.AuthId,
		UserId:    c.MD.UserId,
	})
	if err != nil {
		c.Logger.Errorf("updates.getState - error: %v", err)
		return nil, err
	}

	return rValue, nil
}
