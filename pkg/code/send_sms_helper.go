package code

import (
	"context"

	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/pkg/code/conf"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/pkg/code/none"
)

type VerifyCodeInterface interface {
	SendSmsVerifyCode(ctx context.Context, phoneNumber, code, codeHash string) (string, error)
	VerifySmsCode(ctx context.Context, codeHash, code, extraData string) error
}

func NewVerifyCode(c *conf.SmsVerifyCodeConfig) VerifyCodeInterface {
	if c == nil {
		c = new(conf.SmsVerifyCodeConfig)
	}

	switch c.Name {
	// case "predefined":
	// 	return predefined.New(c)
	case "none":
		return none.New(c)
	}
	return none.New(c)
}
