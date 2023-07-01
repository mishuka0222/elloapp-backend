package core

import (
	"encoding/json"
	"errors"
	"fmt"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/authorization/authorization"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/pkg/date"
)

type AuthSingUPReq struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	Gender      string `json:"gender"`
	DateOfBirth string `json:"date_of_birth,omitempty"`
	Type        string `json:"type"`
	Kind        string `json:"kind"`
	Email       string `json:"email"`
	Phone       string `json:"phone,omitempty"`
	CountryCode string `json:"country_code"`
	Avatar      string `json:"avatar"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
}

type AuthSingUPResp struct {
	Email              string `json:"email"`
	ConfirmationExpire int64  `json:"confirmation_expire"`
}

type AuthSignUpErrorResp struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (c *AuthorizationCore) AuthSingUP(in json.RawMessage) (interface{}, error) {
	var req AuthSingUPReq
	if err := json.Unmarshal(in, &req); err != nil {
		return nil, err
	} else if errResp := req.validate(); errResp != nil {
		return errResp, errors.New("validation fail")
	}

	respCustom, err := c.svcCtx.Dao.AuthorizationClient.AuthSignUp(c.ctx, &authorization.AuthSignUpRequest{
		Password:    req.Password,
		Gender:      req.Gender,
		DateOfBirth: req.DateOfBirth,
		Type:        req.Type,
		Kind:        req.Kind,
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

func (r *AuthSingUPReq) validate() []*AuthSignUpErrorResp {
	var resp []*AuthSignUpErrorResp

	if r.Kind != "private" && r.Kind != "public" {
		resp = append(resp, &AuthSignUpErrorResp{
			Field:   "kind",
			Message: "account kind should be private or public",
		})
	}

	switch r.Type {
	case "business":
		r.DateOfBirth = ""
	case "personal":
		if _, err := date.FormatDateIso8601(r.DateOfBirth); err != nil {
			resp = append(resp, &AuthSignUpErrorResp{
				Field:   "date_of_birth",
				Message: fmt.Sprintf("date of birth should be in ISO 8601, current \"%v\"", r.DateOfBirth),
			})
		}
	default:
		resp = append(resp, &AuthSignUpErrorResp{
			Field:   "type",
			Message: "account type should be business or personal",
		})

	}
	return resp
}
