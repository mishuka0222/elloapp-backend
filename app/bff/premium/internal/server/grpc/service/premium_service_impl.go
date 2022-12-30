package service

import (
	"context"

	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/bff/premium/internal/core"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
)

// HelpGetPremiumPromo
// help.getPremiumPromo#b81b93d4 = help.PremiumPromo;
func (s *Service) HelpGetPremiumPromo(ctx context.Context, request *mtproto.TLHelpGetPremiumPromo) (*mtproto.Help_PremiumPromo, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("help.getPremiumPromo - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.HelpGetPremiumPromo(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("help.getPremiumPromo - reply: %s", r.DebugString())
	return r, err
}

// PaymentsAssignAppStoreTransaction
// payments.assignAppStoreTransaction#80ed747d receipt:bytes purpose:InputStorePaymentPurpose = Updates;
func (s *Service) PaymentsAssignAppStoreTransaction(ctx context.Context, request *mtproto.TLPaymentsAssignAppStoreTransaction) (*mtproto.Updates, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("payments.assignAppStoreTransaction - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.PaymentsAssignAppStoreTransaction(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("payments.assignAppStoreTransaction - reply: %s", r.DebugString())
	return r, err
}

// PaymentsAssignPlayMarketTransaction
// payments.assignPlayMarketTransaction#dffd50d3 receipt:DataJSON purpose:InputStorePaymentPurpose = Updates;
func (s *Service) PaymentsAssignPlayMarketTransaction(ctx context.Context, request *mtproto.TLPaymentsAssignPlayMarketTransaction) (*mtproto.Updates, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("payments.assignPlayMarketTransaction - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.PaymentsAssignPlayMarketTransaction(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("payments.assignPlayMarketTransaction - reply: %s", r.DebugString())
	return r, err
}

// PaymentsCanPurchasePremium
// payments.canPurchasePremium#9fc19eb6 purpose:InputStorePaymentPurpose = Bool;
func (s *Service) PaymentsCanPurchasePremium(ctx context.Context, request *mtproto.TLPaymentsCanPurchasePremium) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("payments.canPurchasePremium - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.PaymentsCanPurchasePremium(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("payments.canPurchasePremium - reply: %s", r.DebugString())
	return r, err
}

// PaymentsRequestRecurringPayment
// payments.requestRecurringPayment#146e958d user_id:InputUser recurring_init_charge:string invoice_media:InputMedia = Updates;
func (s *Service) PaymentsRequestRecurringPayment(ctx context.Context, request *mtproto.TLPaymentsRequestRecurringPayment) (*mtproto.Updates, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("payments.requestRecurringPayment - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.PaymentsRequestRecurringPayment(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("payments.requestRecurringPayment - reply: %s", r.DebugString())
	return r, err
}

// PaymentsRestorePlayMarketReceipt
// payments.restorePlayMarketReceipt#d164e36a receipt:bytes = Updates;
func (s *Service) PaymentsRestorePlayMarketReceipt(ctx context.Context, request *mtproto.TLPaymentsRestorePlayMarketReceipt) (*mtproto.Updates, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("payments.restorePlayMarketReceipt - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.PaymentsRestorePlayMarketReceipt(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("payments.restorePlayMarketReceipt - reply: %s", r.DebugString())
	return r, err
}
