package core

import (
	"encoding/json"
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/biz/authorization/authorization"
)

type AuthSingINReq struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

type AuthSingINResp struct {
	// TODO: need to write logic
}

// AuthSingIN
// TODO: need to write logic
func (c *AuthorizationCore) AuthSingIN(in json.RawMessage) (*AuthSingINResp, error) {
	var req AuthSingINReq
	if err := json.Unmarshal(in, &req); err != nil {
		return nil, err
	}

	respOrigin, err := c.svcCtx.AuthorizationService.AuthSignIn(c.ctx, &mtproto.TLAuthSignIn{
		Constructor:   mtproto.CRC32_auth_signIn_8d52a951,
		PhoneNumber:   req.UserName,
		PhoneCodeHash: req.Password,
		// ....
	})
	if err != nil {
		return nil, err
	}

	// TODO: need to write logic
	respCustom, err := c.svcCtx.Dao.AuthorizationClient.AuthSingIN(c.ctx, &authorization.AuthSingInRequest{
		UserName: req.UserName,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}

	_, _ = respOrigin, respCustom

	return &AuthSingINResp{}, nil
}
