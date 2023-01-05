/*
 *  Copyright (c) 2018, https://github.com/airwide-code
 *  All rights reserved.
 *
 *
 *
 */

package phonecall

import (
	"github.com/gogo/protobuf/types"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/phone_call/internal/dao"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

func (p *PhoneCallSession) SetGB(dao dao.Dao, gb []byte) {
	p.GB = gb
	// TODO: write logic
	//dao.GetPhoneCallSessionsDAO(dao.DB_MASTER).UpdateGB(hex.EncodeToString(gb), p.Id)
}

func (p *PhoneCallSession) SetAdminDebugData(dao dao.Dao, dataJson string) {
	// TODO: write logic
	//dao.GetPhoneCallSessionsDAO(dao.DB_MASTER).UpdateAdminDebugData(dataJson, p.Id)
}

func (p *PhoneCallSession) SetParticipantDebugData(dataJson string) {
	// TODO: write logic
	//dao.GetPhoneCallSessionsDAO(dao.DB_MASTER).UpdateParticipantDebugData(dataJson, p.Id)
}

func (p *PhoneCallSession) toPhoneCallProtocol() *mtproto.PhoneCallProtocol {
	return &mtproto.PhoneCallProtocol{
		Constructor:  mtproto.CRC32_phoneCallProtocol,
		UdpP2P:       p.UdpP2P,
		UdpReflector: p.UdpReflector,
		MinLayer:     p.MinLayer,
		MaxLayer:     p.MaxLayer,
	}
}

func (p *PhoneCallSession) ToPhoneCallProtocol() *mtproto.PhoneCallProtocol {
	return p.toPhoneCallProtocol()
}

// phoneCallRequested#83761ce4 id:long access_hash:long date:int admin_id:int participant_id:int g_a_hash:bytes protocol:PhoneCallProtocol = PhoneCall;
func (p *PhoneCallSession) ToPhoneCallRequested() *mtproto.TLPhoneCallRequested {
	return &mtproto.TLPhoneCallRequested{Data2: &mtproto.PhoneCall{
		Id:            p.Id,
		AccessHash:    p.ParticipantAccessHash,
		Date:          int32(p.Date),
		AdminId:       p.AdminId,
		ParticipantId: p.ParticipantId,
		GAHash:        p.GA,
		Protocol:      p.toPhoneCallProtocol(),
	}}
}

// phoneCallWaiting#1b8f4ad1 flags:# id:long access_hash:long date:int admin_id:int participant_id:int protocol:PhoneCallProtocol receive_date:flags.0?int = PhoneCall;
func (p *PhoneCallSession) ToPhoneCallWaiting(selfId int64, receiveDate int32) *mtproto.TLPhoneCallWaiting {
	var (
		accessHash int64
	)

	if selfId == p.AdminId {
		accessHash = p.AdminAccessHash
	} else {
		accessHash = p.ParticipantAccessHash
	}

	return &mtproto.TLPhoneCallWaiting{Data2: &mtproto.PhoneCall{
		Id:            p.Id,
		AccessHash:    accessHash,
		Date:          int32(p.Date),
		AdminId:       p.AdminId,
		ParticipantId: p.ParticipantId,
		GAHash:        p.GA,
		Protocol:      p.toPhoneCallProtocol(),
		ReceiveDate:   &types.Int32Value{Value: receiveDate},
	}}
}

// phoneCallAccepted#6d003d3f id:long access_hash:long date:int admin_id:int participant_id:int g_b:bytes protocol:PhoneCallProtocol = PhoneCall;
func (p *PhoneCallSession) ToPhoneCallAccepted() *mtproto.TLPhoneCallAccepted {
	return &mtproto.TLPhoneCallAccepted{Data2: &mtproto.PhoneCall{
		Id:            p.Id,
		AccessHash:    p.AdminAccessHash,
		Date:          int32(p.Date),
		AdminId:       p.AdminId,
		ParticipantId: p.ParticipantId,
		GB:            p.GB,
		Protocol:      p.toPhoneCallProtocol(),
	}}
}

// phoneConnection#9d4c17c0 id:long ip:string ipv6:string port:int peer_tag:bytes = PhoneConnection;
func makeConnection() *mtproto.PhoneConnection {
	// TODO: write logic
	return &mtproto.PhoneConnection{
		Constructor: mtproto.CRC32_phoneConnection_9d4c17c0,
		Id:          50003,
		// Ip:      "192.168.4.32",
		Ip:      "192.168.1.104",
		Ipv6:    "",
		Port:    50001,
		PeerTag: []byte("24ffcbeb7980d28b"),
	}
}

// phoneCall#ffe6ab67 id:long access_hash:long date:int admin_id:int participant_id:int g_a_or_b:bytes key_fingerprint:long protocol:PhoneCallProtocol connection:PhoneConnection alternative_connections:Vector<PhoneConnection> start_date:int = PhoneCall;
func (p *PhoneCallSession) ToPhoneCall(selfId int64, keyFingerprint int64) *mtproto.TLPhoneCall {
	var (
		accessHash int64
		gaOrGb     []byte
	)

	if selfId == p.AdminId {
		accessHash = p.AdminAccessHash
		gaOrGb = p.GB
	} else {
		accessHash = p.ParticipantAccessHash
		gaOrGb = p.GA
	}

	return &mtproto.TLPhoneCall{Data2: &mtproto.PhoneCall{
		Id:             p.Id,
		AccessHash:     accessHash,
		Date:           int32(p.Date),
		AdminId:        p.AdminId,
		ParticipantId:  p.ParticipantId,
		GAOrB:          gaOrGb,
		KeyFingerprint: keyFingerprint,
		Protocol:       p.toPhoneCallProtocol(),
		// TODO: write logic
		Connections: []*mtproto.PhoneConnection{makeConnection()},
		//Connections: []*mtproto.PhoneConnection{},
		StartDate: 0,
	}}
}
