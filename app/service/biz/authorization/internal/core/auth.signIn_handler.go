package core

import (
	"context"
	"errors"
	"fmt"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/messenger/sync/sync"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/authsession/authsession"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/authorization/authorization"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/user/user"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/models"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto/rpc/metadata"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/pkg/code/conf"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

// AuthSingIN
// TODO: need to write logic
func (c *AuthorizationCore) AuthSingIN(in *authorization.AuthSignInRequest) (*mtproto.Auth_Authorization, error) {
	var (
		userEllo models.UsersEllo
	)
	if err := c.svcCtx.Gorm.Where("username = ?", in.Username).First(&userEllo).Error; err != nil && err != gorm.ErrRecordNotFound {
		err = fmt.Errorf("can not get users_ello record (%v)", err)
		c.Logger.Error(err)
		return nil, err
	} else if err == gorm.ErrRecordNotFound {
		err = errors.New("username or password incorrect")
		c.Logger.Error(err)
		return nil, err
	} else if !userEllo.EmailConfirmed {
		err := fmt.Errorf("email not confirmed for user %s", userEllo.Username)
		c.Logger.Error(err)
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userEllo.Password), []byte(in.Password)); err != nil {
		err = errors.New("username or password are incorrect")
		c.Logger.Error(err)
		return nil, err
	}

	ctx, err := metadata.RpcMetadataToOutgoing(context.Background(), &metadata.RpcMetadata{
		ServerId:      in.MData.ServerId,
		ClientAddr:    in.MData.ClientAddr,
		AuthId:        in.MData.AuthId,
		SessionId:     in.MData.SessionId,
		ReceiveTime:   in.MData.ReceiveTime,
		ClientMsgId:   in.MData.ClientMsgId,
		UserId:        in.MData.UserId,
		IsBot:         in.MData.IsBot,
		Layer:         in.MData.Layer,
		Client:        in.MData.Client,
		IsAdmin:       in.MData.IsAdmin,
		Langpack:      in.MData.Langpack,
		PermAuthKeyId: in.MData.PermAuthKeyId,
	})
	if err != nil {
		c.Logger.Errorf("biz.auth.signUp_handler context %v", err)
		return nil, err
	}

	userRecord, err := c.svcCtx.UserClient.UserGetImmutableUser(ctx, &user.TLUserGetImmutableUser{
		Id: int64(userEllo.UserID),
	})
	if err != nil {
		c.Logger.Errorf("biz.auth.signUp_handler user(%s) is err - %v", in.Username, err)
		return nil, err
	} else if userRecord == nil {
		c.Logger.Errorf("biz.auth.signUp_handler user(%s) is nil", in.Username)
		err = mtproto.ErrInternelServerError
		return nil, err
	}

	if _, err := c.svcCtx.Dao.AuthsessionClient.AuthsessionBindAuthKeyUser(ctx, &authsession.TLAuthsessionBindAuthKeyUser{
		AuthKeyId: in.MData.AuthId,
		UserId:    userRecord.User.Id,
	}); err != nil {
		c.Logger.Errorf("biz.auth.signUp_handler AuthsessionBindAuthKeyUser %v", err)
		return nil, err
	}

	selfUser := userRecord.ToSelfUser()

	//c.svcCtx.AuthLogic.DeletePhoneCode(c.ctx, in.MData.AuthId, in.PhoneNumber, phoneCodeHash)
	region, _ := c.svcCtx.Dao.GetCountryAndRegionByIp(in.MData.ClientAddr)

	var (
		now     = time.Now()
		signInN *mtproto.Update
	)

	if len(c.svcCtx.Config.SignInServiceNotification) == 0 {
		signInN = mtproto.MakeSignInServiceNotification(selfUser, in.MData.AuthId, in.MData.Client, region, in.MData.ClientAddr)
	} else {
		signInN = mtproto.MakeTLUpdateServiceNotification(&mtproto.Update{
			Popup:          false,
			InboxDate:      mtproto.MakeFlagsInt32(int32(now.Unix())),
			Type:           fmt.Sprintf("auth%d_%d", in.MData.AuthId, now.Unix()),
			Message_STRING: "",
			Media:          mtproto.MakeTLMessageMediaEmpty(nil).To_MessageMedia(),
			Entities:       nil,
		}).To_Update()

		builder := conf.ToMessageBuildHelper(
			c.svcCtx.Config.SignInServiceNotification,
			map[string]interface{}{
				"name":     mtproto.GetUserName(selfUser),
				"now":      now.UTC(),
				"client":   in.MData.Client,
				"region":   region,
				"clientIp": in.MData.ClientAddr,
			})
		signInN.Message_STRING, signInN.Entities = mtproto.MakeTextAndMessageEntities(builder)
	}

	c.svcCtx.Dao.SyncClient.SyncUpdatesNotMe(
		c.ctx,
		&sync.TLSyncUpdatesNotMe{
			UserId:    userRecord.Id(),
			AuthKeyId: in.MData.AuthId,
			Updates:   mtproto.MakeUpdatesByUpdates(signInN),
		})

	return mtproto.MakeTLAuthAuthorization(&mtproto.Auth_Authorization{
		User: selfUser,
	}).To_Auth_Authorization(), nil
}
