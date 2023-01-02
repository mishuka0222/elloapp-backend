package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/messenger/sync/sync"
	userpb "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/user/user"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/username/username"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/pkg2/strings2"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/pkg2/utils"
)

// AccountUpdateUsername
// account.updateUsername#3e0bdd7c username:string = User;
func (c *UsernamesCore) AccountUpdateUsername(in *mtproto.TLAccountUpdateUsername) (*mtproto.User, error) {
	me, err := c.svcCtx.Dao.UserClient.UserGetImmutableUser(c.ctx, &userpb.TLUserGetImmutableUser{
		Id: c.MD.UserId,
	})
	if err != nil {
		c.Logger.Errorf("account.updateUsername - error: %v", err)
		return nil, err
	}

	var (
		username2 = in.GetUsername()
	)

	if username2 != me.Username() {
		// TODO: 分布式事物
		if err = c.updateUsername(c.MD.UserId, me.Username(), username2); err != nil {
			c.Logger.Errorf("account.updateUsername - error: %v", err)
		} else if _, err = c.svcCtx.Dao.UserClient.UserUpdateUsername(c.ctx, &userpb.TLUserUpdateUsername{
			UserId:   c.MD.UserId,
			Username: username2,
		}); err != nil {
			c.Logger.Errorf("account.updateUsername - error: %v", err)
		} else {
			me.SetUsername(username2)

			c.svcCtx.Dao.SyncClient.SyncUpdatesNotMe(c.ctx, &sync.TLSyncUpdatesNotMe{
				UserId:    c.MD.UserId,
				AuthKeyId: c.MD.AuthId,
				Updates: mtproto.MakeUpdatesByUpdates(mtproto.MakeTLUpdateUserName(&mtproto.Update{
					UserId:    c.MD.UserId,
					FirstName: me.FirstName(),
					LastName:  me.LastName(),
					Username:  username2,
				}).To_Update()),
			})
		}
	}

	return me.ToSelfUser(), nil
}

func (c *UsernamesCore) updateUsername(userId int64, from, username2 string) error {
	if username2 != "" {
		if len(username2) < username.MinUsernameLen ||
			!strings2.IsAlNumString(username2) ||
			utils.IsNumber(username2[0]) {
			err := mtproto.ErrUsernameInvalid // mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_USERNAME_INVALID)
			c.Logger.Errorf("account.updateUsername - format error: %v", err)
			return err
		}

		ok, err := c.svcCtx.Dao.UsernameClient.UsernameUpdateUsername(c.ctx, &username.TLUsernameUpdateUsername{
			PeerType: mtproto.PEER_USER,
			PeerId:   userId,
			Username: username2,
		})
		// log.Debugf("ok: %v, err: %v", ok, err)
		if err != nil {
			c.Logger.Errorf("account.updateUsername - format error: %v", err)
			return err
		} else {
			if !mtproto.FromBool(ok) {
				err = mtproto.ErrUsernameOccupied // mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_USERNAME_OCCUPIED)
				c.Logger.Errorf("account.updateUsername - format error: %v", err)
				return err
			}
		}
	}

	if from != "" {
		// delete username
		_, err := c.svcCtx.Dao.UsernameClient.UsernameDeleteUsername(c.ctx, &username.TLUsernameDeleteUsername{
			Username: from,
		})
		if err != nil {
			c.Logger.Errorf("account.updateUsername - format error: %v", err)
			return err
		}
	}

	return nil
}
