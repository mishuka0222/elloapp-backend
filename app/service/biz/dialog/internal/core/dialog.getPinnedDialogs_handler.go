package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/dialog/dialog"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/dialog/internal/dal/dataobject"
)

// DialogGetPinnedDialogs
// dialog.getPinnedDialogs  user_id:long folder_id:int = Vector<DialogExt>;
func (c *DialogCore) DialogGetPinnedDialogs(in *dialog.TLDialogGetPinnedDialogs) (*dialog.Vector_DialogExt, error) {
	var (
		dList    dialog.DialogExtList
		folderId = in.GetFolderId()
		meId     = in.GetUserId()
	)

	if folderId == 0 {
		c.svcCtx.Dao.DialogsDAO.SelectPinnedDialogsWithCB(
			c.ctx,
			meId,
			func(i int, v *dataobject.DialogsDO) {
				dList = append(dList, makeDialog(v))
			})
	} else {
		c.svcCtx.Dao.DialogsDAO.SelectFolderPinnedDialogsWithCB(
			c.ctx,
			meId,
			func(i int, v *dataobject.DialogsDO) {
				dList = append(dList, makeDialog(v))
			})
	}

	return &dialog.Vector_DialogExt{
		Datas: dList,
	}, nil
}
