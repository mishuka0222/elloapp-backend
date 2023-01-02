package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/dialog/dialog"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/dialog/internal/dal/dataobject"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"

	"github.com/zeromicro/go-zero/core/jsonx"
)

// DialogGetDialogFilters
// dialog.getDialogFilters user_id:long = Vector<DialogFilterExt>;
func (c *DialogCore) DialogGetDialogFilters(in *dialog.TLDialogGetDialogFilters) (*dialog.Vector_DialogFilterExt, error) {
	var (
		dialogFilterExtList []*dialog.DialogFilterExt
	)

	c.svcCtx.Dao.DialogFiltersDAO.SelectListWithCB(
		c.ctx,
		in.UserId,
		func(i int, v *dataobject.DialogFiltersDO) {
			dialogFilter := &dialog.DialogFilterExt{
				Id:           v.DialogFilterId,
				DialogFilter: nil,
				Order:        v.OrderValue,
			}

			if err := jsonx.UnmarshalFromString(v.DialogFilter, &dialogFilter.DialogFilter); err != nil {
				c.Logger.Errorf("jsonx.UnmarshalFromString(%v) - error: %v", v, err)
				// continue
				return
			}

			if dialogFilter.DialogFilter == nil {
				dialogFilter.DialogFilter = mtproto.MakeTLDialogFilter(nil).To_DialogFilter()
			}

			dialogFilterExtList = append(dialogFilterExtList, dialogFilter)
		})

	return &dialog.Vector_DialogFilterExt{
		Datas: dialogFilterExtList,
	}, nil
}
