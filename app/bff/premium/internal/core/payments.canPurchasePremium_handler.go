package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

// PaymentsCanPurchasePremium
// payments.canPurchasePremium#9fc19eb6 purpose:InputStorePaymentPurpose = Bool;
func (c *PremiumCore) PaymentsCanPurchasePremium(in *mtproto.TLPaymentsCanPurchasePremium) (*mtproto.Bool, error) {
	// TODO: not impl
	// c.Logger.Errorf("payments.canPurchasePremium blocked, License key from https://elloapp.com required to unlock enterprise features.")

	return mtproto.BoolFalse, nil
}
