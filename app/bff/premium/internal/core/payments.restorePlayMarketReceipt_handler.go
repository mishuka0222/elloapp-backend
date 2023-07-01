package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

// PaymentsRestorePlayMarketReceipt
// payments.restorePlayMarketReceipt#d164e36a receipt:bytes = Updates;
func (c *PremiumCore) PaymentsRestorePlayMarketReceipt(in *mtproto.TLPaymentsRestorePlayMarketReceipt) (*mtproto.Updates, error) {
	// TODO: not impl
	c.Logger.Errorf("payments.restorePlayMarketReceipt blocked, License key from https://elloapp.com required to unlock enterprise features.")

	return nil, mtproto.ErrEnterpriseIsBlocked
}
