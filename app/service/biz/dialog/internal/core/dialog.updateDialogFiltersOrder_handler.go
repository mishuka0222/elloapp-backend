/*
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/dialog/dialog"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/pkg2/stores/sqlx"
	"time"
)

// DialogUpdateDialogFiltersOrder
// dialog.updateDialogFiltersOrder user_id:long order:Vector<long> = Bool;
func (c *DialogCore) DialogUpdateDialogFiltersOrder(in *dialog.TLDialogUpdateDialogFiltersOrder) (*mtproto.Bool, error) {
	var (
		err    error
		orderV = time.Now().Unix() << 32
	)

	sqlx.TxWrapper(c.ctx, c.svcCtx.Dao.DB, func(tx *sqlx.Tx, result *sqlx.StoreResult) {
		for _, id := range in.Order {
			_, err = c.svcCtx.DialogFiltersDAO.UpdateOrder(c.ctx, orderV, in.UserId, id)
			if err != nil {
				result.Err = err
				return
			}
			orderV--
		}
	})

	return mtproto.BoolTrue, nil
}
