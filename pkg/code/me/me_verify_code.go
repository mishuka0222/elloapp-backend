package me

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"gitlab.com/merehead/elloapp/backend/elloapp_backend/pkg/code/conf"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/pkg2/hack"

	"github.com/zeromicro/go-zero/core/logx"
)

var (
	_smsURL = "http://127.0.0.1:8181/code?phone=%s&code=%s"
)

func New(c *conf.SmsVerifyCodeConfig) *meVerifyCode {
	return &meVerifyCode{
		code: c,
	}
}

type meVerifyCode struct {
	code *conf.SmsVerifyCodeConfig
}

func (m *meVerifyCode) SendSmsVerifyCode(ctx context.Context, phoneNumber, code, codeHash string) (string, error) {
	urlV := m.code.SendCodeUrl + fmt.Sprintf("?phone=%s&code=%s", phoneNumber, code)
	logx.Infof("send me sms: %s", urlV)
	resp, err := http.Get(urlV)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logx.Errorf("request verify code error: %v", err)
		return "", err
	} else {
		logx.Infof("result: %s", hack.String(body))
	}
	_ = body
	return "", nil
}

func (m *meVerifyCode) VerifySmsCode(ctx context.Context, codeHash, code, extraData string) error {
	if len(code) != 5 {
		return fmt.Errorf("code invalid")
	}

	//
	if code != extraData {
		return fmt.Errorf("code invalid")
	}

	// ...
	return nil
}
