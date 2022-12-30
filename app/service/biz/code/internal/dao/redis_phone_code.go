package dao

import (
	"context"
	"encoding/json"
	"fmt"

	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/code/code"

	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logx"
)

const (
	phoneCodeTimeout int64 = 90 // salt timeout
	cachePhonePrefix       = "phone_codes"
)

func genCachePhoneCodeKey(authKeyId int64, phoneNumber string) string {
	return fmt.Sprintf("%s_%d_%s", cachePhonePrefix, authKeyId, phoneNumber)
}

func (d *Dao) GetCachePhoneCode(ctx context.Context, authKeyId int64, phoneNumber string) (*code.PhoneCodeTransaction, error) {
	cacheKey := genCachePhoneCodeKey(authKeyId, phoneNumber)

	v, err := d.kv.Get(cacheKey)
	if err != nil {
		logx.WithContext(ctx).Errorf("conn.GET(%s) error(%v)", cacheKey, err)
		return nil, err
	}

	codeData := &code.PhoneCodeTransaction{}
	err = jsonx.UnmarshalFromString(v, codeData)
	return codeData, err
}

func (d *Dao) PutCachePhoneCode(ctx context.Context, authKeyId int64, phoneNumber string, codeData *code.PhoneCodeTransaction) (err error) {
	cacheKey := genCachePhoneCodeKey(authKeyId, phoneNumber)
	b, _ := json.Marshal(codeData)

	if err = d.kv.Setex(cacheKey, string(b), int(phoneCodeTimeout)); err != nil {
		logx.WithContext(ctx).Errorf("conn.SETEX(%s) error(%v)", cacheKey, err)
	}
	return
}

func (d *Dao) DeleteCachePhoneCode(ctx context.Context, authKeyId int64, phoneNumber string) (err error) {
	cacheKey := genCachePhoneCodeKey(authKeyId, phoneNumber)

	if _, err = d.kv.Del(cacheKey); err != nil {
		logx.WithContext(ctx).Errorf("conn.DEL(%s) error(%v)", cacheKey, err)
	}

	return
}
