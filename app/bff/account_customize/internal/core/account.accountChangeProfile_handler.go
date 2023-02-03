package core

import (
	"encoding/json"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/account/account"
)

type ChangeProfile struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Username    string `json:"username"`
	Bio         string `json:"bio"`
	Gender      string `json:"gender"`
	Birthday    string `json:"birthday"`
	CountryCode string `json:"country_code"`
}

func (c *AccountCore) AccountChangeProfile(in json.RawMessage) (*account.ChangeAccountInfoResp, error) {
	var req ChangeProfile
	if err := json.Unmarshal(in, &req); err != nil {
		return &account.ChangeAccountInfoResp{
			Status:  false,
			Message: "can not unmarshal",
		}, err
	}

	res, err := c.svcCtx.Dao.AccountClient.AccountChangeProfile(c.ctx, &account.ChangeAccountInfoReq{
		UserId:      c.MD.UserId,
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		UserName:    req.Username,
		Bio:         req.Bio,
		Birthday:    req.Birthday,
		CountryCode: req.CountryCode,
	})

	return res, err
}
