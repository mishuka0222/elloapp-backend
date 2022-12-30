// Copyright 2022 Teamgram Authors
//  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Author: teamgramio (teamgram.io@gmail.com)
//

package dao

import (
	"context"
	"encoding/base64"
	"fmt"
	"strconv"

	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/authsession/internal/dal/dataobject"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/pkg2/stores/sqlc"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/pkg2/stores/sqlx"

	"github.com/zeromicro/go-zero/core/contextx"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/mr"
	"github.com/zeromicro/go-zero/core/threading"
)

const (
	cacheAuthKeyPrefix   = "auth_keys"
	cacheAuthKeyV2Prefix = "auth_keys2"
)

func genCacheAuthKeyKey(id int64) string {
	return fmt.Sprintf("%s_%d", cacheAuthKeyPrefix, id)
}

func genCacheAuthKeyV2Key(id int64) string {
	return fmt.Sprintf("%s_%d", cacheAuthKeyV2Prefix, id)
}

//
//func (d *Dao) PutAuthKey(ctx context.Context, keyId int64, keyData *mtproto.AuthKeyInfo, expiredIn int32) (err error) {
//	var (
//		key = genCacheAuthKeyKey(keyId)
//
//		args = map[string]string{
//			"auth_key_type":          strconv.Itoa(int(keyData.AuthKeyType)),
//			"auth_key_id":            strconv.FormatInt(keyData.AuthKeyId, 10),
//			"auth_key":               string(keyData.AuthKey),
//			"perm_auth_key_id":       strconv.FormatInt(keyData.PermAuthKeyId, 10),
//			"temp_auth_key_id":       strconv.FormatInt(keyData.TempAuthKeyId, 10),
//			"media_temp_auth_key_id": strconv.FormatInt(keyData.MediaTempAuthKeyId, 10),
//		}
//	)
//
//	// TODO(@benqi): args error??
//	if err = d.kv.Hmset(key, args); err != nil {
//		logx.WithContext(ctx).Errorf("conn.Send(HMSET %s,%v) error(%v)", key, args, err)
//		return
//	}
//
//	if expiredIn > 0 {
//		if err = d.kv.Expire(key, int(expiredIn)); err != nil {
//			logx.WithContext(ctx).Errorf("conn.Send(EXPIRE %d,%d) error(%v)", key, expiredIn, err)
//		}
//	}
//
//	return
//}
//
//func (d *Dao) UnsafeBindKeyId(ctx context.Context, keyId int64, bindType int32, bindKeyId int64) (err error) {
//	var (
//		key = genCacheAuthKeyKey(keyId)
//	)
//
//	switch bindType {
//	case mtproto.AuthKeyTypePerm:
//		if err = d.kv.Hset(key, "perm_auth_key_id", strconv.FormatInt(bindKeyId, 10)); err != nil {
//			logx.WithContext(ctx).Errorf("conn.Do(HSET %s,perm_auth_key_id,%d) error(%v)", key, bindKeyId, err)
//		}
//	case mtproto.AuthKeyTypeTemp:
//		if err = d.kv.Hset(key, "temp_auth_key_id", strconv.FormatInt(bindKeyId, 10)); err != nil {
//			logx.WithContext(ctx).Errorf("conn.Do(HSET %s,temp_auth_key_id,%d) error(%v)", key, bindKeyId, err)
//		}
//	case mtproto.AuthKeyTypeMediaTemp:
//		if err = d.kv.Hset(key, "media_temp_auth_key_id", strconv.FormatInt(bindKeyId, 10)); err != nil {
//			logx.WithContext(ctx).Errorf("conn.Do(HSET %s,media_temp_auth_key_id,%d) error(%v)", key, bindKeyId, err)
//		}
//	default:
//		return
//	}
//
//	return
//}

func (d *Dao) getAuthKey(ctx context.Context, keyId int64) (keyData *mtproto.AuthKeyInfo, err error) {
	var (
		key    = genCacheAuthKeyKey(keyId)
		values map[string]string
	)

	values, err = d.kv.Hgetall(key)
	if err != nil {
		logx.WithContext(ctx).Errorf("conn.Do(HGETALL %s) error(%v)", key, err)
		return
	} else if len(values) == 0 {
		err = sqlc.ErrNotFound
		return
	}

	// TODO(@benqi): check len(values)

	keyData = mtproto.MakeTLAuthKeyInfo(&mtproto.AuthKeyInfo{}).To_AuthKeyInfo()
	for k, v := range values {
		switch k {
		case "auth_key_type":
			authKeyType, _ := strconv.Atoi(v)
			keyData.AuthKeyType = int32(authKeyType)
		case "auth_key_id":
			keyData.AuthKeyId, _ = strconv.ParseInt(v, 10, 64)
		case "auth_key":
			keyData.AuthKey = []byte(v)
		case "perm_auth_key_id":
			keyData.PermAuthKeyId, _ = strconv.ParseInt(v, 10, 64)
		case "temp_auth_key_id":
			keyData.TempAuthKeyId, _ = strconv.ParseInt(v, 10, 64)
		case "media_temp_auth_key_id":
			keyData.MediaTempAuthKeyId, _ = strconv.ParseInt(v, 10, 64)
		}
	}

	return
}

func (d *Dao) QueryAuthKeyV2(ctx context.Context, authKeyId int64) (*mtproto.AuthKeyInfo, error) {
	var (
		keyInfo = mtproto.MakeTLAuthKeyInfo(&mtproto.AuthKeyInfo{
			AuthKeyId:          authKeyId,
			AuthKey:            nil,
			AuthKeyType:        0,
			PermAuthKeyId:      0,
			TempAuthKeyId:      0,
			MediaTempAuthKeyId: 0,
		}).To_AuthKeyInfo()
	)

	err := d.CachedConn.QueryRow(
		ctx,
		keyInfo,
		genCacheAuthKeyV2Key(authKeyId),
		func(ctx context.Context, conn *sqlx.DB, v interface{}) error {
			kInfo := v.(*mtproto.AuthKeyInfo)
			err := mr.Finish(
				func() error {
					do1, err2 := d.AuthKeyInfosDAO.SelectByAuthKeyId(ctx, authKeyId)
					if err2 != nil {
						return err2
					} else if do1 == nil {
						return sqlc.ErrNotFound
					}
					kInfo.AuthKeyType = do1.AuthKeyType
					kInfo.PermAuthKeyId = do1.PermAuthKeyId
					kInfo.TempAuthKeyId = do1.TempAuthKeyId
					kInfo.MediaTempAuthKeyId = do1.MediaTempAuthKeyId

					return nil
				},
				func() error {
					do2, err2 := d.AuthKeysDAO.SelectByAuthKeyId(ctx, authKeyId)
					if err2 != nil {
						return err2
					} else if do2 == nil {
						return sqlc.ErrNotFound
					}

					kInfo.AuthKey, err2 = base64.RawStdEncoding.DecodeString(do2.Body)
					if err2 != nil {
						return err2
					}

					return nil
				})
			if err == sqlc.ErrNotFound {
				kInfo2, _ := d.getAuthKey(ctx, authKeyId)
				if kInfo2 != nil {
					kInfo.AuthKeyType = kInfo2.AuthKeyType
					kInfo.AuthKey = kInfo2.AuthKey
					kInfo.TempAuthKeyId = kInfo2.TempAuthKeyId
					kInfo.PermAuthKeyId = kInfo2.PermAuthKeyId
					kInfo.MediaTempAuthKeyId = kInfo2.MediaTempAuthKeyId
					threading.GoSafe(func() {
						d.AuthKeyInfosDAO.Insert(
							contextx.ValueOnlyFrom(ctx),
							&dataobject.AuthKeyInfosDO{
								AuthKeyId:          keyInfo.AuthKeyId,
								AuthKeyType:        keyInfo.AuthKeyType,
								PermAuthKeyId:      keyInfo.PermAuthKeyId,
								TempAuthKeyId:      keyInfo.TempAuthKeyId,
								MediaTempAuthKeyId: keyInfo.MediaTempAuthKeyId,
								Deleted:            false,
							})
					})
					err = nil
				}
			}
			return err
		})
	if err != nil {
		if err == sqlc.ErrNotFound {
			err = mtproto.ErrAuthKeyInvalid
		} else {
			err = mtproto.ErrInternelServerError
		}
		return nil, err
	}

	return keyInfo, nil
}

func (d *Dao) SetAuthKeyV2(ctx context.Context, authKey *mtproto.AuthKeyInfo, expiredIn int32) (err error) {
	// TODO(@benqi): expiredIn
	_ = expiredIn

	tR := sqlx.TxWrapper(ctx, d.DB, func(tx *sqlx.Tx, result *sqlx.StoreResult) {
		_, _, err = d.AuthKeysDAO.InsertTx(tx, &dataobject.AuthKeysDO{
			AuthKeyId: authKey.AuthKeyId,
			Body:      base64.RawStdEncoding.EncodeToString(authKey.AuthKey),
		})
		if err != nil {
			result.Err = err
			return
		}
		_, _, err = d.AuthKeyInfosDAO.InsertTx(tx, &dataobject.AuthKeyInfosDO{
			AuthKeyId:          authKey.AuthKeyId,
			AuthKeyType:        authKey.AuthKeyType,
			PermAuthKeyId:      authKey.PermAuthKeyId,
			TempAuthKeyId:      authKey.TempAuthKeyId,
			MediaTempAuthKeyId: authKey.MediaTempAuthKeyId,
			Deleted:            false,
		})
		if err != nil {
			result.Err = err
		}
	})

	return tR.Err
}

func (d *Dao) UnsafeBindKeyIdV2(ctx context.Context, keyId int64, bindType int32, bindKeyId int64) (err error) {
	var (
		key = genCacheAuthKeyV2Key(keyId)
	)

	_, _, err = d.CachedConn.Exec(
		ctx,
		func(ctx context.Context, conn *sqlx.DB) (int64, int64, error) {
			//
			var (
				err2 error
			)

			switch bindType {
			case mtproto.AuthKeyTypePerm:
				_, err2 = d.AuthKeyInfosDAO.UpdateCustomMap(
					ctx,
					map[string]interface{}{
						"perm_auth_key_id": bindKeyId,
					},
					keyId)
			case mtproto.AuthKeyTypeTemp:
				_, err2 = d.AuthKeyInfosDAO.UpdateCustomMap(
					ctx,
					map[string]interface{}{
						"temp_auth_key_id": bindKeyId,
					},
					keyId)
			case mtproto.AuthKeyTypeMediaTemp:
				_, err2 = d.AuthKeyInfosDAO.UpdateCustomMap(
					ctx,
					map[string]interface{}{
						"media_temp_auth_key_id": bindKeyId,
					},
					keyId)
			}

			return 0, 0, err2
		},
		key)

	return
}

func (d *Dao) GetPermAuthKeyId(ctx context.Context, authKeyId int64) int64 {
	if k, err := d.QueryAuthKeyV2(ctx, authKeyId); err != nil || k == nil {
		return 0
	} else {
		return k.PermAuthKeyId
	}
}
