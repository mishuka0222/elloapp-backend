package service

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

/*
	 By android source code - RequestFlagWithoutLogin:
		TL_account_getAuthorizationForm
		TL_account_getPasswordSettings
		TL_account_confirmPasswordEmail
		TL_account_updatePasswordSettings
		TL_account_getTmpPassword
		TL_account_getPassword
		TL_account_deleteAccount

		TL_auth_resendCode
		TL_auth_signIn
		TL_auth_cancelCode
		TL_auth_requestPasswordRecovery
		TL_auth_checkPassword
		TL_auth_checkRecoveryPassword
		TL_auth_recoverPassword
		TL_auth_signUp
		TL_auth_exportAuthorization
		TL_auth_importAuthorization
		TL_auth_bindTempAuthKey
		TL_auth_cancelCode

		TL_rpc_drop_answer
		TL_get_future_salts
		TL_ping

		TL_help_getNearestDc
		TL_help_getConfig
		TL_help_getCdnConfig

		TL_langpack_getLanguages
		TL_langpack_getLanguages
		TL_langpack_getDifference
		TL_langpack_getLangPack
		TL_langpack_getStrings
*/
func checkRpcWithoutLogin(tl mtproto.TLObject) bool {
	switch tl.(type) {
	// account
	// case *mtproto.TLAccountGetPassword:
	//	return true

	// auth
	case *mtproto.TLAuthSendCode,
		*mtproto.TLAuthResendCode,
		*mtproto.TLAuthSignUp,
		*mtproto.TLAuthSignIn,
		*mtproto.TLAuthImportLoginToken,
		*mtproto.TLAuthExportedAuthorization,
		*mtproto.TLAuthExportAuthorization,
		*mtproto.TLAuthImportAuthorization,
		*mtproto.TLAuthCancelCode,
		// *mtproto.TLAuthRequestPasswordRecovery,	// TODO: before process, try fetch usrId
		// *mtproto.TLAuthRecoverPassword,			// TODO: before process, try fetch usrId
		*mtproto.TLAuthExportLoginToken,
		*mtproto.TLAuthAcceptLoginToken,
		*mtproto.TLAuthLogOut, // TODO: before process, try fetch usrId
		*mtproto.TLAuthBindTempAuthKey:
		return true

	// help
	case *mtproto.TLHelpGetConfig,
		*mtproto.TLHelpGetNearestDc,
		*mtproto.TLHelpGetAppUpdate,
		*mtproto.TLHelpGetCdnConfig,
		*mtproto.TLHelpGetAppConfig:
		return true

	// langpack
	case *mtproto.TLLangpackGetLangPack,
		*mtproto.TLLangpackGetStrings,
		*mtproto.TLLangpackGetDifference,
		*mtproto.TLLangpackGetLanguages,
		*mtproto.TLLangpackGetLanguage:
		return true

	// TODO(@benqi): debug.
	case *mtproto.TLUploadGetWebFile,
		*mtproto.TLUploadGetFile:
		return true

	// country
	case *mtproto.TLHelpGetCountriesList:
		return true

	case *mtproto.TLJsonObject:
		return true

	default:
		return false
	}
}
