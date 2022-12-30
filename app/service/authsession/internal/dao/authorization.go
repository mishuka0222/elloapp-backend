package dao

import (
	"context"
	"net"

	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/authsession/internal/dal/dataobject"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/pkg2/stores/sqlx"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/mr"
)

func (d *Dao) getCountryAndRegionByIp(ip string) (string, string) {
	if d.MMDB == nil {
		return "UNKNOWN", ""
	} else {
		r, err := d.MMDB.City(net.ParseIP(ip))
		if err != nil {
			logx.Errorf("getCountryAndRegionByIp - error: %v", err)
			return "UNKNOWN", ""
		}

		return r.City.Names["en"] + ", " + r.Country.Names["en"], r.Country.IsoCode
	}
}

//func getAppNameByAppId(appId int32) string {
//	return "tdesktop"
//}

func (d *Dao) GetAuthorization(ctx context.Context, authKeyId int64) (*mtproto.Authorization, error) {
	cData, err := d.GetCacheAuthData(ctx, authKeyId)
	if err != nil {
		return nil, err
	}

	country, region := d.getCountryAndRegionByIp(cData.ClientIp())

	// TODO(@benqi): fill plat_form, app_name, (country, region)
	return mtproto.MakeTLAuthorization(&mtproto.Authorization{
		Current:         false,
		OfficialApp:     true,
		Hash:            0,
		PasswordPending: false,
		DeviceModel:     cData.DeviceModel(),
		Platform:        "",
		SystemVersion:   cData.SystemVersion(),
		ApiId:           cData.ApiId(),
		AppName:         cData.LangPack(),
		AppVersion:      cData.AppVersion(),
		DateCreated:     0,
		DateActive:      0,
		Ip:              cData.ClientIp(),
		Country:         country,
		Region:          region,
	}).To_Authorization(), nil
}

type idxId struct {
	idx int
	id  int64
}

func (d *Dao) GetAuthorizations(ctx context.Context, userId int64, excludeAuthKeyId int64) (authorizations []*mtproto.Authorization) {
	doList, _ := d.AuthUsersDAO.SelectListByUserId(ctx, userId)
	if len(doList) == 0 {
		return
	}

	authorizations = make([]*mtproto.Authorization, len(doList)+1)
	mr.ForEach(
		func(source chan<- interface{}) {
			for i := 0; i < len(doList); i++ {
				source <- idxId{i, doList[i].AuthKeyId}
			}
		},
		func(item interface{}) {
			idx := item.(idxId)
			//kId := item.(int64)
			cData, _ := d.GetCacheAuthData(ctx, idx.id)
			if cData != nil {
				country, region := d.getCountryAndRegionByIp(cData.ClientIp())
				// TODO(@benqi): fill plat_form, app_name, (country, region)
				authorization := mtproto.MakeTLAuthorization(&mtproto.Authorization{
					Current:         false,
					OfficialApp:     true,
					PasswordPending: false,
					Hash:            cData.Hash(),
					DeviceModel:     cData.DeviceModel(),
					Platform:        "",
					SystemVersion:   cData.SystemVersion(),
					ApiId:           cData.ApiId(),
					AppName:         cData.LangPack(),
					AppVersion:      cData.AppVersion(),
					DateCreated:     int32(cData.DateCreated()),
					DateActive:      int32(cData.DateActivated()),
					Ip:              cData.ClientIp(),
					Country:         country,
					Region:          region,
				}).To_Authorization()

				if idx.id == excludeAuthKeyId {
					authorization.Current = true
					authorization.Hash = 0
					authorizations[0] = authorization
				} else {
					authorizations[idx.idx+1] = authorization
				}
			}
		})

	return removeAllNil(authorizations)
}

func (d *Dao) ResetAuthorization(ctx context.Context, userId int64, authKeyId, hash int64) []int64 {
	var (
		cacheKeyIdList []string
		hashList       []int64
		keyIdList      []int64
	)

	d.AuthUsersDAO.SelectListByUserIdWithCB(
		ctx,
		userId,
		func(i int, v *dataobject.AuthUsersDO) {
			if hash == 0 {
				cacheKeyIdList = append(cacheKeyIdList, genAuthDataCacheKey(v.AuthKeyId))
				hashList = append(hashList, v.Id)
				keyIdList = append(keyIdList, v.AuthKeyId)
			} else {
				if hash == v.Hash && authKeyId != v.AuthKeyId {
					cacheKeyIdList = append(cacheKeyIdList, genAuthDataCacheKey(v.AuthKeyId))
					hashList = append(hashList, v.Id)
					keyIdList = append(keyIdList, v.AuthKeyId)
				}
			}
		})
	if len(keyIdList) == 0 {
		return keyIdList
	}

	d.CachedConn.Exec(
		ctx,
		func(ctx context.Context, conn *sqlx.DB) (int64, int64, error) {
			d.AuthUsersDAO.DeleteByHashList(ctx, hashList)
			return 0, 0, nil
		},
		cacheKeyIdList...)

	return keyIdList
}
