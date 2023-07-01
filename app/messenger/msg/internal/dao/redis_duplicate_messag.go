package dao

import (
	"context"
	"fmt"

	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"

	"github.com/gogo/protobuf/proto"
	"github.com/zeromicro/go-zero/core/logx"
)

const (
	duplicateMessageId   = "duplicate_message_id"
	duplicateMessageData = "duplicate_message_data"
	expireTimeout        = 60 // 60s
)

func makeDuplicateMessageKey(prefix string, senderUserId, clientRandomId int64) string {
	return fmt.Sprintf("%s_%d_%d", prefix, senderUserId, clientRandomId)
}

func (d *Dao) HasDuplicateMessage(ctx context.Context, senderUserId, clientRandomId int64) (bool, error) {
	//conn := d.redis.Redis.Get(ctx)
	//defer conn.Close()

	k := makeDuplicateMessageKey(duplicateMessageId, senderUserId, clientRandomId)

	seq, err := d.KV.Incr(k)
	if err != nil {
		logx.WithContext(ctx).Errorf("checkDuplicateMessage - INCR {%s}, error: {%v}", k, err)
		return false, err
	}

	if err = d.KV.Expire(k, expireTimeout); err != nil {
		logx.WithContext(ctx).Errorf("expire DuplicateMessage - EXPIRE {%s, %d}, error: %s", k, expireTimeout, err)
		return false, err
	}

	return seq > 1, nil
}

func (d *Dao) PutDuplicateMessage(ctx context.Context, senderUserId, clientRandomId int64, upd *mtproto.Updates) error {
	k := makeDuplicateMessageKey(duplicateMessageData, senderUserId, clientRandomId)
	cacheData, _ := proto.Marshal(upd)

	if err := d.KV.Setex(k, string(cacheData), expireTimeout); err != nil {
		logx.WithContext(ctx).Errorf("putDuplicateMessage - SET {%s, %s, %d}, error: %s", k, cacheData, expireTimeout, err)
		return err
	}

	return nil
}

func (d *Dao) GetDuplicateMessage(ctx context.Context, senderUserId, clientRandomId int64) (*mtproto.Updates, error) {
	k := makeDuplicateMessageKey(duplicateMessageData, senderUserId, clientRandomId)

	if cacheData, err := d.KV.Get(k); err != nil {
		if err.Error() == "redigo: nil returned" {
			return nil, nil
		}

		logx.WithContext(ctx).Errorf("getDuplicateMessage - GET {%s}, error: %s", k, err)
		return nil, err
	} else {
		upd := &mtproto.Updates{}
		proto.Unmarshal([]byte(cacheData), upd)

		return upd, nil
	}
}
