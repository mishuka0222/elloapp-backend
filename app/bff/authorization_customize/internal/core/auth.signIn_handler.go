package core

import (
	"encoding/json"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/authorization/authorization"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

type AuthSingINReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// AuthSingIN
// TODO: need to write logic
func (c *AuthorizationCore) AuthSingIN(in json.RawMessage) (*mtproto.Auth_Authorization, error) {
	var req AuthSingINReq
	if err := json.Unmarshal(in, &req); err != nil {
		return nil, err
	}

	// TODO: need to write logic
	resp, err := c.svcCtx.Dao.AuthorizationClient.AuthSignIn(c.ctx, &authorization.AuthSignInRequest{
		Username: req.Username,
		Password: req.Password,
		MData: &authorization.MData{
			ServerId:      c.MD.ServerId,
			ClientAddr:    c.MD.ClientAddr,
			AuthId:        c.MD.AuthId,
			SessionId:     c.MD.SessionId,
			ReceiveTime:   c.MD.ReceiveTime,
			UserId:        c.MD.UserId,
			ClientMsgId:   c.MD.ClientMsgId,
			IsBot:         c.MD.IsBot,
			Layer:         c.MD.Layer,
			Client:        c.MD.Client,
			IsAdmin:       c.MD.IsAdmin,
			Langpack:      c.MD.Langpack,
			PermAuthKeyId: c.MD.PermAuthKeyId,
		},
	})
	if err != nil {
		return nil, err
	}

	return resp, nil
}
