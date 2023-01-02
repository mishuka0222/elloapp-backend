package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
)

// PaymentsRequestRecurringPayment
// payments.requestRecurringPayment#146e958d user_id:InputUser recurring_init_charge:string invoice_media:InputMedia = Updates;
func (c *PremiumCore) PaymentsRequestRecurringPayment(in *mtproto.TLPaymentsRequestRecurringPayment) (*mtproto.Updates, error) {
	// TODO: not impl
	c.Logger.Errorf("payments.requestRecurringPayment blocked, License key from https://elloapp.com required to unlock enterprise features.")

	return nil, mtproto.ErrEnterpriseIsBlocked
}
