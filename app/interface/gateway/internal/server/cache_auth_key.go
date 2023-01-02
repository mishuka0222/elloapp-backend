package server

import (
	"strconv"

	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
)

func (s *Server) GetAuthKey(authKeyId int64) *mtproto.AuthKeyInfo {
	var (
		cacheK = strconv.Itoa(int(authKeyId))
		value  *mtproto.AuthKeyInfo
	)

	if v, ok := s.cache.Get(cacheK); ok {
		value = v.(*mtproto.AuthKeyInfo)
	}

	return value
}

func (s *Server) PutAuthKey(keyInfo *mtproto.AuthKeyInfo) {
	var (
		cacheK = strconv.Itoa(int(keyInfo.AuthKeyId))
	)

	// TODO: expires_in
	s.cache.Set(cacheK, keyInfo)
}
