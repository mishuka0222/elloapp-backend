package core

import (
	"github.com/gogo/protobuf/types"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

// AuthImportAuthorization
// auth.importAuthorization#a57a7dad id:long bytes:bytes = auth.Authorization;
func (c *AuthorizationCore) AuthImportAuthorization(in *mtproto.TLAuthImportAuthorization) (*mtproto.Auth_Authorization, error) {
	// TODO: make tmp_session ????
	rValue := mtproto.MakeTLAuthAuthorization(&mtproto.Auth_Authorization{
		TmpSessions: &types.Int32Value{Value: int32(in.GetId())},
		User:        mtproto.MakeTLUserEmpty(nil).To_User(),
	}).To_Auth_Authorization()

	return rValue, nil
}
