package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/username/username"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/pkg2/strings2"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/pkg2/utils"
)

// AccountCheckUsername
// account.checkUsername#2714d86c username:string = Bool;
func (c *UsernamesCore) AccountCheckUsername(in *mtproto.TLAccountCheckUsername) (*mtproto.Bool, error) {
	// Check username format
	// You can choose a username on Telegram.
	// If you do, other people will be able to find
	// you by this username and contact you
	// without knowing your phone number.
	//
	// You can use a-z, 0-9 and underscores.
	// Minimum length is 5 characters.";
	//
	if len(in.Username) < username.MinUsernameLen ||
		!strings2.IsAlNumString(in.Username) ||
		utils.IsNumber(in.Username[0]) {
		err := mtproto.ErrUsernameInvalid
		c.Logger.Errorf("account.checkUsername#2714d86c - format error: %v", err)
		return nil, err
	} else {
		existed, err := c.svcCtx.Dao.UsernameClient.UsernameCheckAccountUsername(c.ctx, &username.TLUsernameCheckAccountUsername{
			UserId:   c.MD.UserId,
			Username: in.GetUsername(),
		})
		if err != nil {
			return nil, err
		}

		switch existed.GetPredicateName() {
		case username.Predicate_usernameExistedNotMe:
			err = mtproto.ErrUsernameOccupied
			c.Logger.Errorf("account.checkUsername#2714d86c - exists username: %v", err)
			return mtproto.BoolFalse, nil
		default:
			break
		}
	}

	return mtproto.BoolTrue, nil
}
