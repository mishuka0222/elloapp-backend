package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/dialog/dialog"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/pkg2/stores/sqlx"
)

// DialogSetHistoryTTL
// dialog.setHistoryTTL user_id:long peer_type:int peer_id:long ttl_period:int = Bool;
func (c *DialogCore) DialogSetHistoryTTL(in *dialog.TLDialogSetHistoryTTL) (*mtproto.Bool, error) {
	sqlx.TxWrapper(c.ctx, c.svcCtx.Dao.DB, func(tx *sqlx.Tx, result *sqlx.StoreResult) {
		c.svcCtx.Dao.DialogsDAO.UpdateCustomMapTx(
			tx,
			map[string]interface{}{
				"ttl_period": in.TtlPeriod,
			},
			in.UserId,
			in.PeerType,
			in.PeerId)
		c.svcCtx.Dao.DialogsDAO.UpdateCustomMapTx(
			tx,
			map[string]interface{}{
				"ttl_period": in.TtlPeriod,
			},
			in.PeerId,
			in.PeerType,
			in.UserId)
	})

	return mtproto.BoolTrue, nil
}
