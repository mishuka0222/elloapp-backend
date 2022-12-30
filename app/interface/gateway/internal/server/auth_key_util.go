package server

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto/crypto"
)

type authKeyUtil struct {
	keyData *mtproto.AuthKeyInfo
	key     *crypto.AuthKey
}

func newAuthKeyUtil(k *mtproto.AuthKeyInfo) *authKeyUtil {
	return &authKeyUtil{
		keyData: k,
		key:     crypto.NewAuthKey(k.AuthKeyId, k.AuthKey),
	}
}

func (k *authKeyUtil) Equal(o *authKeyUtil) bool {
	return k.keyData.AuthKeyId == o.keyData.AuthKeyId
}

func (k *authKeyUtil) AuthKeyId() int64 {
	return k.keyData.AuthKeyId
}

func (k *authKeyUtil) AuthKeyType() int {
	return int(k.keyData.AuthKeyType)
}

func (k *authKeyUtil) PermAuthKeyId() int64 {
	return k.keyData.PermAuthKeyId
}

func (k *authKeyUtil) TempAuthKeyId() int64 {
	return k.keyData.TempAuthKeyId
}

func (k *authKeyUtil) MediaTempAuthKeyId() int64 {
	return k.keyData.MediaTempAuthKeyId
}

func (k *authKeyUtil) AesIgeEncrypt(rawData []byte) ([]byte, []byte, error) {
	return k.key.AesIgeEncrypt(rawData)
}

func (k *authKeyUtil) AesIgeDecrypt(msgKey, rawData []byte) ([]byte, error) {
	return k.key.AesIgeDecrypt(msgKey, rawData)
}
