package dao

import (
	"context"
	"fmt"

	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/user/internal/dal/dataobject"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/pkg2/stores/sqlc"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/pkg2/stores/sqlx"

	"github.com/zeromicro/go-zero/core/mr"
)

const (
	userPresencesKeyPrefix = "user_presences"
)

func genUserPresencesKey(userId int64) string {
	return fmt.Sprintf("%s_%d", userPresencesKeyPrefix, userId)
}

func (d *Dao) GetLastSeenAt(ctx context.Context, id int64) (*dataobject.UserPresencesDO, error) {
	var (
		do = &dataobject.UserPresencesDO{}
	)

	err := d.CachedConn.QueryRow(
		ctx,
		do,
		genUserPresencesKey(id),
		func(ctx context.Context, conn *sqlx.DB, v interface{}) error {
			do2, err := d.UserPresencesDAO.Select(ctx, id)
			if err != nil {
				return err
			}
			if do2 != nil {
				*v.(*dataobject.UserPresencesDO) = *do2
			} else {
				return sqlc.ErrNotFound
			}
			return nil
		})
	if err != nil {
		return nil, err
	}

	return do, nil
}

func (d *Dao) PutLastSeenAt(ctx context.Context, userId int64, lastSeenAt int64, expires int32) {
	do := &dataobject.UserPresencesDO{
		UserId:     userId,
		LastSeenAt: lastSeenAt,
		Expires:    expires,
	}

	mr.FinishVoid(
		func() {
			d.UserPresencesDAO.UpdateLastSeenAt(ctx, lastSeenAt, expires, userId)
		},
		func() {
			d.CachedConn.SetCache(ctx, genUserPresencesKey(userId), do)
		})
}
