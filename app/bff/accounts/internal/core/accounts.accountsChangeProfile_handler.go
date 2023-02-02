package core

import (
	"encoding/json"

	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/accounts/accounts"
)

type ChangeProfileReq struct {
	UserId      string `json:"user_id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	UserName    string `json:"user_name"`
	Bio         string `json:"bio"`
	Birthday    string `json:"birthday"`
	Gender      string `json:"gender"`
	CountryCode string `json:"country_code"`
}

func (c *AccountsCore) AccountsChangeProfile(in json.RawMessage) (*accounts.ChangeAccountInfoResp, error) {
	var changeProfileReq ChangeProfileReq
	if err := json.Unmarshal(in, &changeProfileReq); err != nil {
		return nil, err
	}

	res, err := c.svcCtx.Dao.AccountsClient.AccountsChangeProfile(c.ctx, &accounts.ChangeAccountInfoReq{
		UserId:      c.MD.GetUserId(),
		FirstName:   changeProfileReq.FirstName,
		LastName:    changeProfileReq.LastName,
		UserName:    changeProfileReq.UserName,
		Bio:         changeProfileReq.Bio,
		Birthday:    changeProfileReq.Birthday,
		Gender:      changeProfileReq.Gender,
		CountryCode: changeProfileReq.CountryCode,
	})

	if err != nil {
		return nil, err
	}
	// todo: add your logic here and delete this line
	return res, nil
}
