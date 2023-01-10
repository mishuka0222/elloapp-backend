package core

import (
	"encoding/json"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/authorization/authorization"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

type AuthSingUPReq struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	Gender      string `json:"gender"`
	DateOfBirth string `json:"date_of_birth"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	CountryCode string `json:"country_code"`
	Avatar      string `json:"avatar"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
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
		PhoneNumber:   req.Phone,
		PhoneCodeHash: req.Password,
		FirstName:     req.FirstName,
		LastName:      req.LastName,
	})
	if err != nil {
		return nil, err
	}

	// TODO: need to write logic
	respCustom, err := c.svcCtx.Dao.AuthorizationClient.AuthSignUp(c.ctx, &authorization.AuthSignUpRequest{
		UserId:      respOrigin.User.Id,
		Password:    req.Password,
		Gender:      req.Gender,
		DateOfBirth: req.DateOfBirth,
		Email:       req.Email,
		Phone:       req.Phone,
		CountryCode: req.CountryCode,
		Avatar:      req.Avatar,
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		Username:    req.Username,
	})
	if err != nil {
		return nil, err
	}

	_, _ = respOrigin, respCustom

	return &AuthSingUPResp{}, nil
}
