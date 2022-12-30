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
	"time"

	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/dialog/dialog"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/dialog/internal/dal/dataobject"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"

	"github.com/zeromicro/go-zero/core/jsonx"
)

// DialogInsertOrUpdateDialogFilter
// dialog.insertOrUpdateDialogFilter user_id:long id:int dialog_filter:DialogFilter = Bool;
func (c *DialogCore) DialogInsertOrUpdateDialogFilter(in *dialog.TLDialogInsertOrUpdateDialogFilter) (*mtproto.Bool, error) {
	dialogFilterData, err := jsonx.Marshal(in.GetDialogFilter())
	if err != nil {
		c.Logger.Errorf("dialog.insertOrUpdateDialogFilter - error: %v", err)
		return nil, err
	}

	c.svcCtx.Dao.DialogFiltersDAO.InsertOrUpdate(c.ctx, &dataobject.DialogFiltersDO{
		UserId:         in.UserId,
		DialogFilterId: in.Id,
		DialogFilter:   string(dialogFilterData),
		OrderValue:     time.Now().Unix() << 32,
		Deleted:        false,
	})

	return mtproto.BoolTrue, nil
}
