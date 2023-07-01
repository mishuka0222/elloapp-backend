package core

import (
	"errors"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/phone_call/phonecall"
)

func (c *PhonecallCore) SetGB(in *phonecall.TLSetGB) (*phonecall.Void, error) {
	// TODO: write logic with Cache Client
	if sess, ok := K_SESSION[in.SessionId]; !ok {
		return nil, errors.New(" Session not found")
	} else {
		sess.GB = in.Gb
		K_SESSION[in.SessionId] = sess
	}
	return &phonecall.Void{}, nil
}
