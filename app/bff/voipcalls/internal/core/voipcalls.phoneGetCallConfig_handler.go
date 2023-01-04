package core

import (
	"errors"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

func (c *VoipcallsCore) PhoneGetCallConfig(in *mtproto.TLPhoneGetCallConfig) (*mtproto.DataJSON, error) {
	// todo: add your logic here and delete this line
	return nil, errors.New("Unimplemented")
	return &mtproto.DataJSON{}, nil
}
