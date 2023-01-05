package core

import (
	"encoding/hex"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/phone_call/internal/dao/dataobject"
	"math/rand"
	"time"

	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/phone_call/phonecall"
)

func (c *PhonecallCore) MakePhoneCallSession(in *phonecall.TLMakePhoneCallSession) (*phonecall.PhoneCallSession, error) {
	session := &phonecall.PhoneCallSession{
		Id:                    10101010101010, // TODO: IdGen Client => NextId()
		AdminId:               in.AdminId,
		AdminAccessHash:       rand.Int63(),
		ParticipantId:         in.ParticipantId,
		ParticipantAccessHash: rand.Int63(),
		UdpP2P:                in.Protocol.GetUdpP2P(),
		UdpReflector:          in.Protocol.GetUdpReflector(),
		MinLayer:              in.Protocol.GetMinLayer(),
		MaxLayer:              in.Protocol.GetMaxLayer(),
		GA:                    in.Ga,
		State:                 0,
		Date:                  time.Now().Unix(),
	}

	do := &dataobject.PhoneCallSessionsDO{
		CallSessionId:         session.Id,
		AdminId:               session.AdminId,
		AdminAccessHash:       session.AdminAccessHash,
		ParticipantId:         session.ParticipantId,
		ParticipantAccessHash: session.ParticipantAccessHash,
		UdpP2P:                session.UdpP2P,
		UdpReflector:          session.UdpReflector,
		MinLayer:              session.MinLayer,
		MaxLayer:              session.MaxLayer,
		GA:                    hex.EncodeToString(session.GA),
		Date:                  int32(session.Date),
	}
	// TODO: write logic
	//dao.GetPhoneCallSessionsDAO(dao.DB_MASTER).Insert(do)
	_ = do
	K_SESSION[session.Id] = session
	return session, nil
}
