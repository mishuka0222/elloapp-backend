package core

import (
	"errors"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/phone_call/phonecall"
)

var K_SESSION = make(map[int64]*phonecall.PhoneCallSession)

func (c *PhonecallCore) MakePhoneCallSessionByLoad(in *phonecall.TLMakePhoneCallSessionByLoad) (*phonecall.PhoneCallSession, error) {
	// TODO: write logic with Cache Client
	if sess, ok := K_SESSION[in.SessionId]; ok {
		return sess, nil
	} else {
		return nil, errors.New(" Session not found")
	}

	//var do *dataobject.PhoneCallSessionsDO
	//do = c.svcCtx.Dao.PhoneCallSessionsDAO.SelectSession(in.SessionId)
	//if do == nil {
	//	err := fmt.Errorf("not found call session: %d", in.SessionId)
	//	return nil, err
	//}

	//session := &phonecall.PhoneCallSession{
	//	Id:                    do.Id,
	//	AdminId:               do.AdminId,
	//	AdminAccessHash:       do.AdminAccessHash,
	//	ParticipantId:         do.ParticipantId,
	//	ParticipantAccessHash: do.ParticipantAccessHash,
	//	UdpP2P:                do.UdpP2P,
	//	UdpReflector:          do.UdpReflector,
	//	MinLayer:              do.MinLayer,
	//	MaxLayer:              do.MaxLayer,
	//	// GA:                    do.GA,
	//	State: 0,
	//	Date:  int64(do.Date),
	//}

	//session.GA, _ = hex.DecodeString(do.GA)
	//return session, nil
}
