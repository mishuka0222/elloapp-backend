package core

import (
	"errors"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

func (c *VoipcallsCore) PhoneRequestCall(in *mtproto.TLPhoneRequestCall) (*mtproto.Phone_PhoneCall, error) {
	// todo: add your logic here and delete this line
	return nil, errors.New("Unimplemented")
	return &mtproto.Phone_PhoneCall{}, nil
}
