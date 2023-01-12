package core

import (
	"encoding/json"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/authorization/authorization"
)

type AuthSingINReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthSingINResp struct {
	PredicateName         string `json:"predicate_name,omitempty"`
	Constructor           any    `json:"constructor,omitempty"`
	OtherwiseReloginDays  any    `json:"otherwise_relogin_days,omitempty"`
	TmpSessions           any    `json:"tmp_sessions,omitempty"`
	TermsOfService        any    `json:"terms_of_service,omitempty"`
	User                  any    `json:"user,omitempty"`
	SetupPasswordRequired bool   `json:"setup_password_required,omitempty"`
}

// AuthSingIN
// TODO: need to write logic
func (c *AuthorizationCore) AuthSingIN(in json.RawMessage) (*AuthSingINResp, error) {
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

	return &AuthSingINResp{
		PredicateName:         resp.PredicateName,
		Constructor:           resp.Constructor,
		OtherwiseReloginDays:  resp.OtherwiseReloginDays,
		TmpSessions:           resp.TmpSessions,
		TermsOfService:        resp.TermsOfService,
		User:                  resp.User,
		SetupPasswordRequired: resp.SetupPasswordRequired,
	}, nil
}
