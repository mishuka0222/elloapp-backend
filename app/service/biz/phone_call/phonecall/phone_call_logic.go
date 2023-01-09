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

func (p *PhoneCallSession) SetGB(gb []byte) {
	p.GB = gb
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
		PredicateName:   mtproto.Predicate_phoneCallProtocol,
		Constructor:     mtproto.CRC32_phoneCallProtocol,
		UdpP2P:          p.UdpP2P,
		UdpReflector:    p.UdpReflector,
		MinLayer:        p.MinLayer,
		MaxLayer:        p.MaxLayer,
		LibraryVersions: []string{},
	}
}

func (p *PhoneCallSession) ToPhoneCallProtocol() *mtproto.PhoneCallProtocol {
	return p.toPhoneCallProtocol()
}

// phoneCallRequested#83761ce4 id:long access_hash:long date:int admin_id:int participant_id:int g_a_hash:bytes protocol:PhoneCallProtocol = PhoneCall;
func (p *PhoneCallSession) ToPhoneCallRequested(versions []string) *mtproto.TLPhoneCallRequested {
	protocol := p.toPhoneCallProtocol()
	// MARK: "library_versions":["4.1.2","4.0.2","4.0.1","4.0.0","3.0.0","2.7.7","2.4.4"]
	protocol.LibraryVersions = versions
	return mtproto.MakeTLPhoneCallRequested(&mtproto.PhoneCall{
		Constructor:   mtproto.CRC32_phoneCallRequested,
		Id:            p.Id,
		AccessHash:    p.ParticipantAccessHash,
		Date:          int32(p.Date),
		AdminId:       p.AdminId,
		ParticipantId: p.ParticipantId,
		GAHash:        p.GA,
		Protocol:      protocol,
	})
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

	return mtproto.MakeTLPhoneCallWaiting(&mtproto.PhoneCall{
		Constructor:   mtproto.CRC32_phoneCallWaiting,
		Id:            p.Id,
		AccessHash:    accessHash,
		Date:          int32(p.Date),
		AdminId:       p.AdminId,
		ParticipantId: p.ParticipantId,
		GAHash:        p.GA,
		Protocol:      p.toPhoneCallProtocol(),
		ReceiveDate:   &types.Int32Value{Value: receiveDate},
	})
}

// phoneCallAccepted#6d003d3f id:long access_hash:long date:int admin_id:int participant_id:int g_b:bytes protocol:PhoneCallProtocol = PhoneCall;
func (p *PhoneCallSession) ToPhoneCallAccepted(versions []string) *mtproto.TLPhoneCallAccepted {
	protocol := p.toPhoneCallProtocol()
	// MARK: "library_versions":["4.1.2","4.0.2","4.0.1","4.0.0","3.0.0","2.7.7","2.4.4"]
	protocol.LibraryVersions = versions
	return mtproto.MakeTLPhoneCallAccepted(&mtproto.PhoneCall{
		Constructor:   mtproto.CRC32_phoneCallAccepted,
		Id:            p.Id,
		AccessHash:    p.AdminAccessHash,
		Date:          int32(p.Date),
		AdminId:       p.AdminId,
		ParticipantId: p.ParticipantId,
		GB:            p.GB,
		Protocol:      protocol,
	})
}

// phoneConnection#9d4c17c0 id:long ip:string ipv6:string port:int peer_tag:bytes = PhoneConnection;
func makeConnection(addr string) *mtproto.PhoneConnection {
	// TODO: write logic
	return &mtproto.PhoneConnection{
		PredicateName: mtproto.Predicate_phoneConnection,
		Constructor:   mtproto.CRC32_phoneConnection_9d4c17c0,
		Id:            50003,
		// Ip:      "192.168.4.32",
		Ip:      addr,
		Ipv6:    "",
		Port:    50001,
		PeerTag: []byte("24ffcbeb7980d28b"),
	}
}

// phoneCall#ffe6ab67 id:long access_hash:long date:int admin_id:int participant_id:int g_a_or_b:bytes key_fingerprint:long protocol:PhoneCallProtocol connection:PhoneConnection alternative_connections:Vector<PhoneConnection> start_date:int = PhoneCall;
func (p *PhoneCallSession) ToPhoneCall(selfId int64, keyFingerprint int64, addr string, versions []string) *mtproto.TLPhoneCall {
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
	protocol := p.toPhoneCallProtocol()
	// MARK: "library_versions":["4.1.2","4.0.2","4.0.1","4.0.0","3.0.0","2.7.7","2.4.4"]
	protocol.LibraryVersions = versions

	var video bool
	if p.State == 1 {
		video = true
	}

	return mtproto.MakeTLPhoneCall(&mtproto.PhoneCall{
		Constructor:    mtproto.CRC32_phoneCall,
		Id:             p.Id,
		Video:          video,
		AccessHash:     accessHash,
		Date:           int32(p.Date),
		AdminId:        p.AdminId,
		ParticipantId:  p.ParticipantId,
		GAOrB:          gaOrGb,
		KeyFingerprint: keyFingerprint,
		Protocol:       protocol,
		// TODO: write logic
		//Connections: []*mtproto.PhoneConnection{makeConnection(addr)},
		Connections: []*mtproto.PhoneConnection{},
		StartDate:   0,
	})
}
