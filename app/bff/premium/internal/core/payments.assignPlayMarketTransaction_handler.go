package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

// PaymentsAssignPlayMarketTransaction
// payments.assignPlayMarketTransaction#dffd50d3 receipt:DataJSON purpose:InputStorePaymentPurpose = Updates;
func (c *PremiumCore) PaymentsAssignPlayMarketTransaction(in *mtproto.TLPaymentsAssignPlayMarketTransaction) (*mtproto.Updates, error) {
	// TODO: not impl
	c.Logger.Errorf("payments.assignPlayMarketTransaction blocked, License key from https://elloapp.com required to unlock enterprise features.")

	return nil, mtproto.ErrEnterpriseIsBlocked
}
