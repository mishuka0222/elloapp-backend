package core

import (
	"errors"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

func (c *VoipcallsCore) PhoneSaveCallLog(in *mtproto.TLPhoneSaveCallLog) (*mtproto.Bool, error) {
	// todo: add your logic here and delete this line
	return nil, errors.New("Unimplemented")
	return &mtproto.Bool{}, nil
}
