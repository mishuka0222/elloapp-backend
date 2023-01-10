package core

import (
	"encoding/json"
	"github.com/gogo/protobuf/types"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/authorization/authorization"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

type AuthSingINReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthSingINResp struct {
	PredicateName         string                       `json:"predicate_name,omitempty"`
	Constructor           mtproto.TLConstructor        `json:"constructor,omitempty"`
	OtherwiseReloginDays  *types.Int32Value            `json:"otherwise_relogin_days,omitempty"`
	TmpSessions           *types.Int32Value            `json:"tmp_sessions,omitempty"`
	TermsOfService        *mtproto.Help_TermsOfService `json:"terms_of_service,omitempty"`
	User                  *mtproto.User                `json:"user,omitempty"`
	SetupPasswordRequired bool                         `json:"setup_password_required,omitempty"`
}

// AuthSingIN
// TODO: need to write logic
func (c *AuthorizationCore) AuthSingIN(in json.RawMessage) (*AuthSingINResp, error) {
	var req AuthSingINReq
	if err := json.Unmarshal(in, &req); err != nil {
		return nil, err
	}

	// TODO: need to write logic
	if _, err := c.svcCtx.Dao.AuthorizationClient.AuthSignIn(c.ctx, &authorization.AuthSignInRequest{
		Username: req.Username,
		Password: req.Password,
	}); err != nil {
		return nil, err
	}

	respOrigin, err := c.svcCtx.AuthorizationService.AuthSignIn(c.ctx, &mtproto.TLAuthSignIn{
		Constructor:   mtproto.CRC32_auth_signIn_8d52a951,
		PhoneNumber:   req.Username,
		PhoneCodeHash: req.Password,
	})
	if err != nil {
		return nil, err
	}

	return &AuthSingINResp{
		PredicateName:         respOrigin.PredicateName,
		Constructor:           respOrigin.Constructor,
		OtherwiseReloginDays:  respOrigin.OtherwiseReloginDays,
		TmpSessions:           respOrigin.TmpSessions,
		TermsOfService:        respOrigin.TermsOfService,
		User:                  respOrigin.User,
		SetupPasswordRequired: respOrigin.SetupPasswordRequired,
	}, nil
}
