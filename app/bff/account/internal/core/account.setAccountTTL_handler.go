package core

import (
	userpb "gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/user/user"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
)

// AccountSetAccountTTL
// account.setAccountTTL#2442485e ttl:AccountDaysTTL = Bool;
func (c *AccountCore) AccountSetAccountTTL(in *mtproto.TLAccountSetAccountTTL) (*mtproto.Bool, error) {
	// TODO(@benqi): Check ttl
	ttl := in.GetTtl().GetDays()
	switch ttl {
	case 30:
	case 90:
	case 182:
	case 365:
	default:
		err := mtproto.ErrTtlDaysInvalid
		c.Logger.Errorf("account.setAccountTTL - error: %v", err)
		return nil, err
	}

	if _, err := c.svcCtx.Dao.UserClient.UserSetAccountDaysTTL(c.ctx, &userpb.TLUserSetAccountDaysTTL{
		UserId: c.MD.UserId,
		Ttl:    ttl,
	}); err != nil {
		c.Logger.Errorf("account.setAccountTTL - error: %v", err)
	}

	return mtproto.BoolTrue, nil
}
