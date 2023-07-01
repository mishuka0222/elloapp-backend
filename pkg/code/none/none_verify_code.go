package none

import (
	"context"

	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/pkg/code/conf"
)

func New(c *conf.SmsVerifyCodeConfig) *noneVerifyCode {
	return &noneVerifyCode{
		code: c,
	}
}

type noneVerifyCode struct {
	code *conf.SmsVerifyCodeConfig
}

func (m *noneVerifyCode) SendSmsVerifyCode(ctx context.Context, phoneNumber, code, codeHash string) (string, error) {
	return code, nil
}

func (m *noneVerifyCode) VerifySmsCode(ctx context.Context, codeHash, code, extraData string) error {
	if code != "12345" {
		return mtproto.ErrPhoneCodeInvalid
	}
	return nil
}
