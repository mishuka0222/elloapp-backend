package core

import (
	"context"
	"fmt"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/bff/authorization/internal/logic"
	"golang.org/x/crypto/bcrypt"
	"log"

	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto/crypto"
	//"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/bff/authorization/internal/model"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/authsession/authsession"
	userpb "gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/user/user"
	usernamepb "gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/username/username"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/pkg/env2"
)

/*
  Android client auth.signUp#1b067634, handler error
	if (error.text.contains("PHONE_NUMBER_INVALID")) {
		needShowAlert(LocaleController.getString("AppName", R.string.AppName), LocaleController.getString("InvalidPhoneNumber", R.string.InvalidPhoneNumber));
	} else if (error.text.contains("PHONE_CODE_EMPTY") || error.text.contains("PHONE_CODE_INVALID")) {
		needShowAlert(LocaleController.getString("AppName", R.string.AppName), LocaleController.getString("InvalidCode", R.string.InvalidCode));
	} else if (error.text.contains("PHONE_CODE_EXPIRED")) {
		needShowAlert(LocaleController.getString("AppName", R.string.AppName), LocaleController.getString("CodeExpired", R.string.CodeExpired));
	} else if (error.text.contains("FIRSTNAME_INVALID")) {
		needShowAlert(LocaleController.getString("AppName", R.string.AppName), LocaleController.getString("InvalidFirstName", R.string.InvalidFirstName));
	} else if (error.text.contains("LASTNAME_INVALID")) {
		needShowAlert(LocaleController.getString("AppName", R.string.AppName), LocaleController.getString("InvalidLastName", R.string.InvalidLastName));
	} else {
		needShowAlert(LocaleController.getString("AppName", R.string.AppName), error.text);
	}

*/

// AuthSignUp
// auth.signUp#80eee427 phone_number:string phone_code_hash:string first_name:string last_name:string = auth.Authorization;
func (c *AuthorizationCore) AuthSignUp(in *mtproto.TLAuthSignUp) (*mtproto.Auth_Authorization, error) {
	if c.svcCtx.Plugin != nil {
		c.svcCtx.Plugin.OnAuthAction(c.ctx,
			c.MD.AuthId,
			c.MD.ClientMsgId,
			c.MD.ClientAddr,
			in.PhoneNumber,
			logic.GetActionType(in),
			"auth.signUp")
	}

	// IMPORTANT! in.PhoneNumber currently uses as username
	// IMPORTANT! in.PhoneCodeHash currently uses as password

	// 1. check phone_code empty
	var (
		err error
	)

	// 3. check number
	// 3.1. empty
	if in.PhoneNumber == "" {
		c.Logger.Errorf("check phone_number (username) error - empty")
		err = ErrPhoneNumberUsernameInvalid
		return nil, err
	}

	//// 3.2. check phone_number
	//// 客户端发送的手机号格式为: "+86 111 1111 1111"，归一化
	//// We need getRegionCode from phone_number
	//pNumber, err := phonenumber.MakePhoneNumberHelper(in.PhoneNumber, "")
	//if err != nil {
	//	c.Logger.Errorf("check phone_number error - %v", err)
	//	err = mtproto.ErrPhoneNumberInvalid
	//	return nil, err
	//}
	//phoneNumber := pNumber.GetNormalizeDigits()

	if len(in.PhoneNumber) < 3 {
		c.Logger.Errorf("check phone_number (username) error - empty")
		err = ErrPhoneNumberUsernameInvalid
		return nil, err
	}
	phoneNumber := in.PhoneNumber

	if in.PhoneCodeHash == "" {
		c.Logger.Errorf("check phone_code_hash error - empty")
		err = mtproto.ErrPhoneCodeHashEmpty
		return nil, err
	}

	// TODO(@benqi): register name ruler
	// check first name invalid
	if in.FirstName == "" {
		c.Logger.Errorf("check first_name error - empty")
		err = mtproto.ErrFirstNameInvalid
		return nil, err
	}

	// TODO(@benqi): PHONE_NUMBER_FLOOD
	// <string name="PhoneNumberFlood">Sorry, you have deleted and re-created your account too many times recently.
	//    Please wait for a few days before signing up again.</string>
	//

	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	//var (
	//	codeData *model.PhoneCodeTransaction
	//)
	// phoneRegistered := auth.CheckPhoneNumberExist(phoneNumber)

	//codeData, err = c.svcCtx.AuthLogic.DoAuthSignUpV2(c.ctx, c.MD.AuthId, phoneNumber, phoneCode, in.PhoneCodeHash)
	//if err != nil {
	//	c.Logger.Errorf(err.Error())
	//	return nil, err
	//}

	var (
		user *userpb.ImmutableUser
	)

	key := crypto.CreateAuthKey()
	_, err = c.svcCtx.Dao.AuthsessionClient.AuthsessionSetAuthKey(c.ctx, &authsession.TLAuthsessionSetAuthKey{
		AuthKey: &mtproto.AuthKeyInfo{
			AuthKeyId:          key.AuthKeyId(),
			AuthKey:            key.AuthKey(),
			AuthKeyType:        mtproto.AuthKeyTypePerm,
			PermAuthKeyId:      key.AuthKeyId(),
			TempAuthKeyId:      0,
			MediaTempAuthKeyId: 0,
		},
		FutureSalt: nil,
	})

	if err != nil {
		c.Logger.Errorf("create user secret key error")
		return nil, err
	}

	var (
		firstName      = in.FirstName
		lastName       = in.LastName
		username       = in.PhoneNumber
		predefinedUser *mtproto.PredefinedUser
	)

	if env2.PredefinedUser {
		predefinedUser, _ = c.svcCtx.Dao.UserClient.UserGetPredefinedUser(c.ctx, &userpb.TLUserGetPredefinedUser{
			Phone: phoneNumber,
		})
		if predefinedUser == nil {
			c.Logger.Errorf("check predefinedUsers error - %v", err)
			err = mtproto.ErrPhoneNumberInvalid
			return nil, err
		}
		firstName = predefinedUser.GetFirstName().GetValue()
		lastName = predefinedUser.GetLastName().GetValue()
		username = predefinedUser.GetUsername().GetValue()
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(in.PhoneCodeHash), bcrypt.DefaultCost)
	password := string(hashPassword[:])
	if err != nil {
		log.Println(err)
	}

	// Create new user
	fmt.Println("========== (1) ==========")
	if user, err = c.svcCtx.UserClient.UserCreateNewUser(c.ctx, &userpb.TLUserCreateNewUser{
		SecretKeyId: key.AuthKeyId(),
		Phone:       phoneNumber,
		CountryCode: password,
		FirstName:   firstName,
		LastName:    lastName,
	}); err != nil {
		c.Logger.Errorf("createNewUser error: %v", err)
		return nil, err
	}

	_, err_un := c.svcCtx.Dao.UserUpdateUsername(c.ctx, &userpb.TLUserUpdateUsername{
		UserId:   user.Id(),
		Username: user.FirstName(),
	})
	if err_un != nil {
		c.Logger.Errorf("username error: %v", err_un)
		// return nil, err
	}
	_, err_un = c.svcCtx.Dao.UsernameUpdateUsername(c.ctx, &usernamepb.TLUsernameUpdateUsername{
		PeerType: mtproto.PEER_USER,
		PeerId:   user.Id(),
		Username: user.Phone(),
	})
	if err_un != nil {
		c.Logger.Errorf("username error: %v", err_un)
		// return nil, err
	}

	fmt.Println("========== (2) ==========")
	if env2.PredefinedUser {
		c.svcCtx.Dao.UserClient.UserPredefinedBindRegisteredUserId(c.ctx, &userpb.TLUserPredefinedBindRegisteredUserId{
			Phone:            phoneNumber,
			RegisteredUserId: user.User.Id,
		})
		if username != "" {
			// TODO(@benqi): setUsername
			//s.UserFacade.UpdateUsername(ctx, user.Id, username)
			//s.UsernameFacade.UpdateUsername(ctx, model2.PEER_USER, user.Id, username)
			//s.UserFacade.UpdateVerified(ctx, user.Id, predefinedUser.Verified)
		}
	}

	// TODO(@benqi): remove to createNewUser
	// user.Self = true

	// bind auth_key and user_id
	fmt.Println("========== (3) ==========")
	_, err = c.svcCtx.Dao.AuthsessionClient.AuthsessionBindAuthKeyUser(c.ctx, &authsession.TLAuthsessionBindAuthKeyUser{
		AuthKeyId: c.MD.AuthId,
		UserId:    user.User.Id,
	})
	if err != nil {
		c.Logger.Errorf("bindAuthKeyUser error: %v", err)
		err = mtproto.ErrInternelServerError
		return nil, err
	}

	// on event
	fmt.Println("========== (4) ==========")
	c.svcCtx.AuthLogic.DeletePhoneCode(c.ctx, c.MD.AuthId, phoneNumber, in.PhoneCodeHash)
	//fmt.Println(codeData.PhoneCode)
	//c.pushSignInMessage(c.ctx, user.Id(), codeData.PhoneCode)
	//c.onContactSignUp(c.ctx, c.MD.AuthId, user.Id(), phoneNumber)

	return mtproto.MakeTLAuthAuthorization(&mtproto.Auth_Authorization{
		User: user.ToSelfUser(),
	}).To_Auth_Authorization(), nil
}

func (c *AuthorizationCore) onContactSignUp(ctx context.Context, authKeyId, userId int64, phone string) {
	// log.Debugc(ctx, "onContactSignUp - {phone: %s}")
	//importers := s.UserFacade.GetImportersByPhone(ctx, phone)
	//for _, c := range importers {
	//	log.Debugc(ctx, "importer: %v", c)
	//	v := s.AccountFacade.GetSettingValue(ctx, int32(c.ClientId), "contactSignUpNotification")
	//	if v == "true" {
	//		s.MsgFacade.PushUserMessage(ctx,
	//			1,
	//			userId,
	//			int32(c.ClientId),
	//			rand.Int63(),
	//			model2.MakeContactSignUpMessage(userId, int32(c.ClientId)))
	//	} else {
	//		log.Debugc(ctx, "not setting contactSignUpNotification")
	//	}
	//}
	//s.UserFacade.DeleteImportersByPhone(ctx, phone)
}
