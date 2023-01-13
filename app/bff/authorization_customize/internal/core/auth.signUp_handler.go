package core

import (
	"encoding/json"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/authorization/authorization"
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
	Email              string `json:"email"`
	ConfirmationExpire int64  `json:"confirmation_expire"`
}

// AuthSingUP
// TODO: need to write logic
func (c *AuthorizationCore) AuthSingUP(in json.RawMessage) (*AuthSingUPResp, error) {
	var req AuthSingUPReq
	if err := json.Unmarshal(in, &req); err != nil {
		return nil, err
	}

	// TODO: need to write logic
	respCustom, err := c.svcCtx.Dao.AuthorizationClient.AuthSignUp(c.ctx, &authorization.AuthSignUpRequest{
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
		MData: &authorization.MData{
			ServerId:      c.MD.ServerId,
			ClientAddr:    c.MD.ClientAddr,
			AuthId:        c.MD.AuthId,
			SessionId:     c.MD.SessionId,
			ReceiveTime:   c.MD.ReceiveTime,
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

	return &AuthSingUPResp{
		Email:              respCustom.Email,
		ConfirmationExpire: respCustom.ConfirmationExpire,
	}, nil
}
