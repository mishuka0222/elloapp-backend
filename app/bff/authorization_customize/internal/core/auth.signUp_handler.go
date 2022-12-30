package core

import (
	"encoding/json"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/authorization/authorization"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
)

type AuthSingUPReq struct {
	Balance   float64 `json:"balance"`
	UserName  string  `json:"user_name"`
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Password  string  `json:"password"`
	Email     string  `json:"email"`
	Type      string  `json:"type"`
	Profile   string  `json:"profile"`
	Gender    string  `json:"gender"`
}

type AuthSingUPResp struct {
	// TODO: need to write logic
}

// AuthSingUP
// TODO: need to write logic
func (c *AuthorizationCore) AuthSingUP(in json.RawMessage) (*AuthSingUPResp, error) {
	var req AuthSingUPReq
	if err := json.Unmarshal(in, &req); err != nil {
		return nil, err
	}

	respOrigin, err := c.svcCtx.AuthorizationService.AuthSignUp(c.ctx, &mtproto.TLAuthSignUp{
		Constructor:   mtproto.CRC32_auth_signUp,
		PhoneNumber:   req.UserName,
		PhoneCodeHash: req.Password,
		FirstName:     req.FirstName,
		LastName:      req.LastName,
	})
	if err != nil {
		return nil, err
	}

	// TODO: need to write logic
	respCustom, err := c.svcCtx.Dao.AuthorizationClient.AuthSingUP(c.ctx, &authorization.AuthSingUpRequest{
		Balance:   req.Balance,
		UserName:  req.UserName,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Password:  req.Password,
		Email:     req.Email,
		Type:      req.Type,
		Profile:   req.Profile,
		Gender:    req.Gender,
	})
	if err != nil {
		return nil, err
	}

	_, _ = respOrigin, respCustom

	return &AuthSingUPResp{}, nil
}
