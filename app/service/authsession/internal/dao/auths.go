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
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/teamgram/marmota/pkg/stores/sqlc"
	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/authsession/authsession"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/authsession/internal/dal/dataobject"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/mr"
)

const (
	authDataPrefix = "auth_data"
	// authUsersTablePrefix = "auth_users"
)

func genAuthDataCacheKey(id int64) string {
	return fmt.Sprintf("%s_%d", authDataPrefix, id)
}

//
//func genAuthUserCacheKey(id int64) string {
//	return fmt.Sprintf("%s_%d", authUsersTablePrefix, id)
//}

type BindUser struct {
	UserId        int64
	Hash          int64
	DateCreated   int64
	DateActivated int64
}

type CacheAuthData struct {
	Client   *authsession.ClientSession `json:"client"`
	BindUser *BindUser                  `json:"bind_user,omitempty"`
}

//func (c *CacheAuthData) ClientSession() *authsession.ClientSession {
//	if c == nil {
//		return nil
//	}
//	return c.Client
//}

func (c *CacheAuthData) AuthKeyId() int64 {
	return c.Client.GetAuthKeyId()
}

func (c *CacheAuthData) Layer() int32 {
	return c.Client.GetLayer()
}

func (c *CacheAuthData) ApiId() int32 {
	return c.Client.GetApiId()
}

func (c *CacheAuthData) DeviceModel() string {
	return c.Client.GetDeviceModel()
}

func (c *CacheAuthData) SystemVersion() string {
	return c.Client.GetSystemVersion()
}

func (c *CacheAuthData) AppVersion() string {
	return c.Client.GetAppVersion()
}

func (c *CacheAuthData) SystemLangCode() string {
	return c.Client.GetSystemLangCode()
}

func (c *CacheAuthData) LangPack() string {
	return c.Client.GetLangPack()
}

func (c *CacheAuthData) LangCode() string {
	return c.Client.GetLangCode()
}

func (c *CacheAuthData) ClientIp() string {
	return c.Client.GetIp()
}

func (c *CacheAuthData) Proxy() string {
	return c.Client.GetProxy()
}

func (c *CacheAuthData) Params() string {
	return c.Client.GetParams()
}

func (c *CacheAuthData) UserId() int64 {
	if c.BindUser == nil {
		return 0
	}
	return c.BindUser.UserId
}

func (c *CacheAuthData) DateCreated() int64 {
	if c.BindUser == nil {
		return 0
	}
	return c.BindUser.DateCreated
}

func (c *CacheAuthData) DateActivated() int64 {
	if c.BindUser == nil {
		return 0
	}
	return c.BindUser.DateActivated
}

func (c *CacheAuthData) Hash() int64 {
	if c.BindUser == nil {
		return 0
	}
	return c.BindUser.Hash
}

func (d *Dao) GetApiLayer(ctx context.Context, authKeyId int64) int32 {
	cData, err := d.GetCacheAuthData(ctx, authKeyId)
	if err != nil {
		logx.WithContext(ctx).Errorf("not find layer - keyId = %d", authKeyId)
		return 0
	}
	return cData.Layer()
}

func (d *Dao) GetLangCode(ctx context.Context, authKeyId int64) string {
	cData, err := d.GetCacheAuthData(ctx, authKeyId)
	if err != nil {
		logx.WithContext(ctx).Errorf("not find lang_code - keyId = %d", authKeyId)
		return "en"
	}
	return cData.LangCode()
}

func (d *Dao) GetLangPack(ctx context.Context, authKeyId int64) string {
	cData, err := d.GetCacheAuthData(ctx, authKeyId)
	if err != nil {
		logx.WithContext(ctx).Errorf("not find lang_pack - keyId = %d", authKeyId)
		return ""
	}
	return cData.LangPack()
}

func (d *Dao) GetClient(ctx context.Context, authKeyId int64) string {
	cData, err := d.GetCacheAuthData(ctx, authKeyId)
	if err != nil {
		logx.WithContext(ctx).Errorf("not find client - keyId = %d", authKeyId)
		return ""
	}
	c := cData.LangPack()
	if c == "android" {
		if strings.Index(cData.AppVersion(), "TDLib") >= 0 {
			c = "react"
		}
	} else if c == "" {
		if strings.HasSuffix(cData.AppVersion(), " Z") {
			c = "webz"
		}
	}
	return c
}

func (d *Dao) GetAuthKeyUserId(ctx context.Context, authKeyId int64) int64 {
	cData, _ := d.GetCacheAuthData(ctx, authKeyId)
	// do, _ := d.AuthUsersDAO.Select(ctx, authKeyId)
	if cData == nil {
		logx.WithContext(ctx).Errorf("not find user - keyId = %d", authKeyId)
		return 0
	}

	return cData.UserId()
}

func (d *Dao) GetPushSessionId(ctx context.Context, userId int64, authKeyId int64, tokenType int32) int64 {
	do, _ := d.DevicesDAO.Select(ctx, authKeyId, userId, tokenType)
	if do == nil {
		logx.WithContext(ctx).Errorf("not find token - keyId = %d", authKeyId)
		return 0
	}
	sessionId, _ := strconv.ParseInt(do.Token, 10, 64)
	return sessionId
}

func (d *Dao) BindAuthKeyUser(ctx context.Context, authKeyId int64, userId int64) int64 {
	now := time.Now().Unix()
	authUsersDO := &dataobject.AuthUsersDO{
		AuthKeyId:   authKeyId,
		UserId:      userId,
		Hash:        rand.Int63(),
		DateCreated: now,
		DateActived: now,
	}

	_, _, err := d.CachedConn.Exec(
		ctx,
		func(ctx context.Context, conn *sqlx.DB) (int64, int64, error) {
			return d.AuthUsersDAO.InsertOrUpdates(ctx, authUsersDO)
		},
		genAuthDataCacheKey(authKeyId))
	if err != nil {
		return 0
	}

	return authUsersDO.Hash
}

func (d *Dao) UnbindAuthUser(ctx context.Context, authKeyId int64, userId int64) bool {
	var (
		err error
	)

	if authKeyId == 0 {
		var (
			idList []string
		)
		d.AuthUsersDAO.SelectAuthKeyIdsWithCB(
			ctx,
			userId,
			func(i int, v *dataobject.AuthUsersDO) {
				idList = append(idList, genAuthDataCacheKey(v.AuthKeyId))
			})
		if len(idList) > 0 {
			_, _, err = d.CachedConn.Exec(
				ctx,
				func(ctx context.Context, conn *sqlx.DB) (int64, int64, error) {
					_, err2 := d.AuthUsersDAO.DeleteUser(ctx, userId)
					return 0, 0, err2
				},
				idList...)
		}
	} else {
		_, _, err = d.CachedConn.Exec(
			ctx,
			func(ctx context.Context, conn *sqlx.DB) (int64, int64, error) {
				_, err2 := d.AuthUsersDAO.Delete(ctx, authKeyId, userId)
				return 0, 0, err2
			},
			genAuthDataCacheKey(authKeyId))
	}

	return err == nil
}

func (d *Dao) SetClientSessionInfo(ctx context.Context, session *authsession.ClientSession) bool {
	_, _, err := d.CachedConn.Exec(
		ctx,
		func(ctx context.Context, conn *sqlx.DB) (int64, int64, error) {
			do := &dataobject.AuthsDO{
				AuthKeyId:      session.GetAuthKeyId(),
				Layer:          session.GetLayer(),
				ApiId:          session.GetApiId(),
				DeviceModel:    session.GetDeviceModel(),
				SystemVersion:  session.GetSystemVersion(),
				AppVersion:     session.GetAppVersion(),
				SystemLangCode: session.GetSystemLangCode(),
				LangPack:       session.GetLangPack(),
				LangCode:       session.GetLangCode(),
				ClientIp:       session.GetIp(),
				Proxy:          session.GetProxy(),
				Params:         session.GetParams(),
				DateActive:     time.Now().Unix(),
			}
			if do.Params == "" {
				do.Params = "null"
			}
			return d.AuthsDAO.InsertOrUpdate(ctx, do)
		},
		genAuthDataCacheKey(session.GetAuthKeyId()))

	return err == nil
}

func (d *Dao) GetCacheAuthData(ctx context.Context, authKeyId int64) (*CacheAuthData, error) {
	var (
		cData *CacheAuthData
	)

	err := d.CachedConn.QueryRow(
		ctx,
		&cData,
		genAuthDataCacheKey(authKeyId),
		func(ctx context.Context, conn *sqlx.DB, v interface{}) error {
			cacheAuthData := &CacheAuthData{
				Client:   nil,
				BindUser: nil,
			}
			err := mr.Finish(
				func() error {
					do, err := d.AuthsDAO.SelectByAuthKeyId(ctx, authKeyId)
					if err != nil {
						return err
					}
					if do == nil {
						return sqlc.ErrNotFound
					}
					cacheAuthData.Client = authsession.MakeTLClientSession(&authsession.ClientSession{
						AuthKeyId:      authKeyId,
						Ip:             do.ClientIp,
						Layer:          do.Layer,
						ApiId:          do.ApiId,
						DeviceModel:    do.DeviceModel,
						SystemVersion:  do.SystemVersion,
						AppVersion:     do.AppVersion,
						SystemLangCode: do.SystemLangCode,
						LangPack:       do.LangPack,
						LangCode:       do.LangCode,
						Proxy:          do.Proxy,
						Params:         do.Params,
					}).To_ClientSession()

					return nil
				},
				func() error {
					do, _ := d.AuthUsersDAO.Select(ctx, authKeyId)
					if do != nil {
						cacheAuthData.BindUser = &BindUser{
							UserId:        do.UserId,
							Hash:          do.Hash,
							DateCreated:   do.DateCreated,
							DateActivated: do.DateActived,
						}
					}
					return nil
				})

			if err != nil {
				return err
			}
			*v.(**CacheAuthData) = cacheAuthData

			return nil
		})
	if err != nil {
		return nil, err
	}

	return cData, nil
}
