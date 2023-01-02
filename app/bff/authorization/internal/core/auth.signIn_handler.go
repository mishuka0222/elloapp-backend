package core

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/status"
	"time"

	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/bff/authorization/internal/logic"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/messenger/sync/sync"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/authsession/authsession"
	userpb "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/user/user"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/pkg/code/conf"
)

/*
 1. PHONE_NUMBER_UNOCCUPIED ==> setPage(5, true, params, false);
 2. SESSION_PASSWORD_NEEDED ==> invoke rpc: TL_account_getPassword
 3. error:
	if (error.text.contains("PHONE_NUMBER_INVALID")) {
		needShowAlert(LocaleController.getString("AppName", R.string.AppName), LocaleController.getString("InvalidPhoneNumber", R.string.InvalidPhoneNumber));
	} else if (error.text.contains("PHONE_CODE_EMPTY") || error.text.contains("PHONE_CODE_INVALID")) {
		needShowAlert(LocaleController.getString("AppName", R.string.AppName), LocaleController.getString("InvalidCode", R.string.InvalidCode));
	} else if (error.text.contains("PHONE_CODE_EXPIRED")) {
		onBackPressed();
		setPage(0, true, null, true);
		needShowAlert(LocaleController.getString("AppName", R.string.AppName), LocaleController.getString("CodeExpired", R.string.CodeExpired));
	} else if (error.text.startsWith("FLOOD_WAIT")) {
		needShowAlert(LocaleController.getString("AppName", R.string.AppName), LocaleController.getString("FloodWait", R.string.FloodWait));
	} else {
		needShowAlert(LocaleController.getString("AppName", R.string.AppName), LocaleController.getString("ErrorOccurred", R.string.ErrorOccurred) + "\n" + error.text);
	}
*/

// AuthSignIn
// auth.signIn#bcd51581 phone_number:string phone_code_hash:string phone_code:string = auth.Authorization;
func (c *AuthorizationCore) AuthSignIn(in *mtproto.TLAuthSignIn) (*mtproto.Auth_Authorization, error) {
	var (
		phoneCode = in.PhoneCodeHash
	)

	if phoneCode == "" {
		err := mtproto.ErrPasswordEmpty
		c.Logger.Errorf("auth.signIn - error: %v", err)
		return nil, err
	}

	// 3. check number
	// client phone number format: "+86 111 1111 1111"
	//phoneNumber, err := checkPhoneNumberInvalid(in.PhoneNumber)
	//if err != nil {
	//	c.Logger.Errorf("check phone_number(%s) error - %v", in.PhoneNumber, err)
	//	err = mtproto.ErrPhoneNumberInvalid
	//	return nil, err
	//}
	phoneNumber := in.PhoneNumber
	var (
		err error
	)

	// 6. check can do action
	actionType := logic.GetActionType(in)
	if err = c.svcCtx.Dao.CheckCanDoAction(c.ctx, c.MD.AuthId, phoneNumber, actionType); err != nil {
		c.Logger.Errorf("check can do action - %s: %v", phoneNumber, err)
		return nil, err
	}

	//codeData, err2 := c.svcCtx.AuthLogic.DoAuthSignIn(c.ctx,
	//	c.MD.AuthId,
	//	phoneNumber,
	//	"",
	//	phoneCodeHash,
	//	func(codeData2 *model.PhoneCodeTransaction) error {
	//		return c.svcCtx.AuthLogic.VerifyCodeInterface.VerifySmsCode(c.ctx,
	//			codeData2.PhoneCodeHash,
	//			phoneCode,
	//			codeData2.PhoneCodeExtraData)
	//
	//		//log.Debugf("111")
	//		//if s.VerifyCodeInterface == nil {
	//		//	log.Debugf("222")
	//		//	if env2.PredefinedUser {
	//		//		log.Debugf("333")
	//		//		predefinedUser, _ := s.UserFacade.GetPredefinedUser(ctx, phoneNumber)
	//		//		if predefinedUser == nil || predefinedUser.Code != request.PhoneCode {
	//		//			log.Debugf("invalid code: %s", request.PhoneCode)
	//		//			return mtproto.ErrPhoneCodeInvalid
	//		//		} else {
	//		//			return nil
	//		//		}
	//		//	} else {
	//		//		if request.PhoneCode != "12345" {
	//		//			return mtproto.ErrPhoneCodeInvalid
	//		//		} else {
	//		//			return nil
	//		//		}
	//		//	}
	//		//} else {
	//		//	log.Debugf("444")
	//		//	if codeData2.SentCodeType == model.CodeTypeSms {
	//		//		return s.VerifyCodeInterface.VerifySmsCode(ctx, codeData2.PhoneCodeHash, request.PhoneCode, codeData2.PhoneCode)
	//		//	} else {
	//		//		if request.PhoneCode == codeData2.PhoneCode {
	//		//			return nil
	//		//		} else {
	//		//			return mtproto.ErrPhoneCodeInvalid
	//		//		}
	//		//	}
	//		//}
	//	})
	//
	//if err2 != nil {
	//	c.Logger.Error(err2.Error())
	//	err = err2
	//	return nil, err
	//}

	//if c.svcCtx.Plugin != nil {
	//	c.svcCtx.Plugin.OnAuthAction(c.ctx,
	//		c.MD.AuthId,
	//		c.MD.ClientMsgId,
	//		c.MD.ClientAddr,
	//		in.PhoneNumber,
	//		logic.GetActionType(in),
	//		"auth.signIn")
	//}

	// signIn successful, check phoneRegistered.
	//if !codeData.PhoneNumberRegistered {
	//	if !env2.PredefinedUser2 {
	//		if c.MD.Layer >= 104 {
	//			//  not register, next step: auth.singIn
	//			return mtproto.MakeTLAuthAuthorizationSignUpRequired(&mtproto.Auth_Authorization{
	//				// TermsOfService: model.MakeTermOfService(),
	//			}).To_Auth_Authorization(), nil
	//		} else {
	//			c.Logger.Errorf("auth.signIn - not registered, next step auth.signIn, %v", err)
	//			err = mtproto.ErrPhoneNumberUnoccupied
	//			return nil, err
	//		}
	//	} else {
	//		predefinedUser, err3 := c.svcCtx.Dao.UserClient.UserGetPredefinedUser(c.ctx, &userpb.TLUserGetPredefinedUser{
	//			Phone: phoneNumber,
	//		})
	//		if err3 != nil {
	//			c.Logger.Errorf("auth.signIn - not registered, next step auth.signIn, %v", err3)
	//			err = mtproto.ErrPhoneNumberUnoccupied
	//			return nil, err
	//		}
	//
	//		key := crypto.CreateAuthKey()
	//		_, err = c.svcCtx.Dao.AuthsessionClient.AuthsessionSetAuthKey(c.ctx, &authsession.TLAuthsessionSetAuthKey{
	//			AuthKey: &mtproto.AuthKeyInfo{
	//				AuthKeyId:          key.AuthKeyId(),
	//				AuthKey:            key.AuthKey(),
	//				AuthKeyType:        mtproto.AuthKeyTypePerm,
	//				PermAuthKeyId:      key.AuthKeyId(),
	//				TempAuthKeyId:      0,
	//				MediaTempAuthKeyId: 0,
	//			},
	//		})
	//		if err != nil {
	//			c.Logger.Errorf("create user secret key error")
	//			err = mtproto.ErrPhoneNumberUnoccupied
	//			return nil, err
	//		}
	//
	//		// 3.2. check phone_number
	//		// 客户端发送的手机号格式为: "+86 111 1111 1111"，归一化
	//		// We need getRegionCode from phone_number
	//		pNumber, _ := phonenumber.MakePhoneNumberHelper(phoneNumber, "")
	//
	//		// TODO: check
	//		_, err = c.svcCtx.UserClient.UserCreateNewUser(c.ctx, &userpb.TLUserCreateNewUser{
	//			SecretKeyId: key.AuthKeyId(),
	//			Phone:       phoneNumber,
	//			CountryCode: pNumber.GetRegionCode(),
	//			FirstName:   predefinedUser.GetFirstName().GetValue(),
	//			LastName:    predefinedUser.GetLastName().GetValue(),
	//		})
	//		codeData.PhoneNumberRegistered = true
	//	}
	//}

	// TODO(@benqi): err handle
	// do signIn...
	var (
		user *userpb.ImmutableUser
	)

	user, err = c.svcCtx.Dao.UserClient.UserGetImmutableUserByPhone(c.ctx, &userpb.TLUserGetImmutableUserByPhone{
		Phone: phoneNumber,
	})
	if err != nil {
		c.Logger.Errorf("user(%s) is err - %v", phoneNumber, err)
		return nil, err
	} else if user == nil {
		c.Logger.Errorf("user(%s) is nil", phoneNumber)
		err = mtproto.ErrInternelServerError
		return nil, err
	}

	// Bind authKeyId and userId
	c.svcCtx.Dao.AuthsessionClient.AuthsessionBindAuthKeyUser(c.ctx, &authsession.TLAuthsessionBindAuthKeyUser{
		AuthKeyId: c.MD.AuthId,
		UserId:    user.User.Id,
	})

	// Check SESSION_PASSWORD_NEEDED
	if c.svcCtx.Plugin != nil {
		if c.svcCtx.Plugin.CheckSessionPasswordNeeded(c.ctx, user.User.Id) {
			err = mtproto.ErrSessionPasswordNeeded
			c.Logger.Infof("auth.signIn - registered, next step auth.checkPassword: %v", err)
			return nil, err
		}
	}

	code, err := c.svcCtx.Dao.UserClient.UserGetCountryCode(c.ctx, &userpb.TLUserGetCountryCode{
		UserId: user.Id(),
	})
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(code.V), []byte(phoneCode))
	if err != nil {
		err := status.Error(mtproto.ErrBadRequest, "PASSWORD_INVALID")
		c.Logger.Errorf("auth.signIn - error: %v", err)
		return nil, err
	}

	selfUser := user.ToSelfUser()

	//c.svcCtx.AuthLogic.DeletePhoneCode(c.ctx, c.MD.AuthId, in.PhoneNumber, phoneCodeHash)
	region, _ := c.svcCtx.Dao.GetCountryAndRegionByIp(c.MD.ClientAddr)

	var (
		now     = time.Now()
		signInN *mtproto.Update
	)

	if len(c.svcCtx.Config.SignInServiceNotification) == 0 {
		signInN = mtproto.MakeSignInServiceNotification(selfUser, c.MD.AuthId, c.MD.Client, region, c.MD.ClientAddr)
	} else {
		signInN = mtproto.MakeTLUpdateServiceNotification(&mtproto.Update{
			Popup:          false,
			InboxDate:      mtproto.MakeFlagsInt32(int32(now.Unix())),
			Type:           fmt.Sprintf("auth%d_%d", c.MD.AuthId, now.Unix()),
			Message_STRING: "",
			Media:          mtproto.MakeTLMessageMediaEmpty(nil).To_MessageMedia(),
			Entities:       nil,
		}).To_Update()

		builder := conf.ToMessageBuildHelper(
			c.svcCtx.Config.SignInServiceNotification,
			map[string]interface{}{
				"name":     mtproto.GetUserName(selfUser),
				"now":      now.UTC(),
				"client":   c.MD.Client,
				"region":   region,
				"clientIp": c.MD.ClientAddr,
			})
		signInN.Message_STRING, signInN.Entities = mtproto.MakeTextAndMessageEntities(builder)
	}

	c.svcCtx.Dao.SyncClient.SyncUpdatesNotMe(
		c.ctx,
		&sync.TLSyncUpdatesNotMe{
			UserId:    user.Id(),
			AuthKeyId: c.MD.AuthId,
			Updates:   mtproto.MakeUpdatesByUpdates(signInN),
		})

	return mtproto.MakeTLAuthAuthorization(&mtproto.Auth_Authorization{
		User: selfUser,
	}).To_Auth_Authorization(), nil
}
