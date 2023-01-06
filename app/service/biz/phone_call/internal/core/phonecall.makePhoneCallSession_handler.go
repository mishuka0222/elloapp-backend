package core

import (
	"math/rand"
	"time"

	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/phone_call/phonecall"
)

func (c *PhonecallCore) MakePhoneCallSession(in *phonecall.TLMakePhoneCallSession) (*phonecall.PhoneCallSession, error) {
	var state int32 = 0
	if in.IsVideo {
		state = 1
	}

	session := &phonecall.PhoneCallSession{
		Id:                    c.svcCtx.Dao.IDGenClient2.NextId(c.ctx),
		AdminId:               in.AdminId,
		AdminAccessHash:       rand.Int63(),
		ParticipantId:         in.ParticipantId,
		ParticipantAccessHash: rand.Int63(),
		UdpP2P:                in.Protocol.GetUdpP2P(),
		UdpReflector:          in.Protocol.GetUdpReflector(),
		MinLayer:              in.Protocol.GetMinLayer(),
		MaxLayer:              in.Protocol.GetMaxLayer(),
		GA:                    in.GaHash,
		State:                 state,
		Date:                  time.Now().Unix(),
	}

	//do := &dataobject.PhoneCallSessionsDO{
	//	CallSessionId:         session.Id,
	//	AdminId:               session.AdminId,
	//	AdminAccessHash:       session.AdminAccessHash,
	//	ParticipantId:         session.ParticipantId,
	//	ParticipantAccessHash: session.ParticipantAccessHash,
	//	UdpP2P:                session.UdpP2P,
	//	UdpReflector:          session.UdpReflector,
	//	MinLayer:              session.MinLayer,
	//	MaxLayer:              session.MaxLayer,
	//	GA:                    hex.EncodeToString(session.GA),
	//	Date:                  int32(session.Date),
	//}
	//dao.GetPhoneCallSessionsDAO(dao.DB_MASTER).Insert(do)
	//_ = do
	// TODO: write logic with Cache Client
	K_SESSION[session.Id] = session
	return session, nil
}
