package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

// PaymentsAssignAppStoreTransaction
// payments.assignAppStoreTransaction#80ed747d receipt:bytes purpose:InputStorePaymentPurpose = Updates;
func (c *PremiumCore) PaymentsAssignAppStoreTransaction(in *mtproto.TLPaymentsAssignAppStoreTransaction) (*mtproto.Updates, error) {
	// TODO: not impl
	c.Logger.Errorf("payments.assignAppStoreTransaction blocked, License key from https://elloapp.com required to unlock enterprise features.")

	return nil, mtproto.ErrEnterpriseIsBlocked
}
