package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/dialog/dialog"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
)

// DialogDeleteDialogFilter
// dialog.deleteDialogFilter user_id:long id:int = Bool;
func (c *DialogCore) DialogDeleteDialogFilter(in *dialog.TLDialogDeleteDialogFilter) (*mtproto.Bool, error) {
	c.svcCtx.Dao.DialogFiltersDAO.Clear(c.ctx, in.UserId, in.Id)

	return mtproto.BoolTrue, nil
}
