package service

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

////////////////////////////////////////////////////////////////////////////////////////////////////////////
/*
bool SecureRequest::needAck() const {
	if (_data->size() <= kMessageBodyPosition) {
		return false;
	}
	const auto type = mtpTypeId((*_data)[kMessageBodyPosition]);
	switch (type) {
	case mtpc_msg_container:
	case mtpc_msgs_ack:
	case mtpc_http_wait:
	case mtpc_bad_msg_notification:
	case mtpc_msgs_all_info:
	case mtpc_msgs_state_info:
	case mtpc_msg_detailed_info:
	case mtpc_msg_new_detailed_info:
		return false;
	}
	return true;
}
*/

func checkMessageConfirm(msg mtproto.TLObject) bool {
	switch msg.(type) {
	case *mtproto.TLMsgContainer,
		*mtproto.TLMsgsAck,
		*mtproto.TLHttpWait,
		*mtproto.TLBadMsgNotification,
		*mtproto.TLMsgsAllInfo,
		*mtproto.TLMsgsStateInfo,
		*mtproto.TLMsgDetailedInfo,
		*mtproto.TLMsgNewDetailedInfo:

		return false
	default:
		return true
	}
}
