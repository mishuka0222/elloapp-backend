// ConstructorList
// RequestList

package status

import (
	"fmt"

	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/gogo/protobuf/types"
)

//////////////////////////////////////////////////////////////////////////////////////////

var _ *types.Int32Value
var _ *mtproto.Bool
var _ fmt.GoStringer

var clazzIdRegisters2 = map[int32]func() mtproto.TLObject{
	// Constructor
	392473649: func() mtproto.TLObject { // 0x1764ac31
		o := MakeTLSessionEntry(nil)
		o.Data2.Constructor = 392473649
		return o
	},
	-269700200: func() mtproto.TLObject { // 0xefecb398
		o := MakeTLUserSessionEntryList(nil)
		o.Data2.Constructor = -269700200
		return o
	},

	// Method
	1381075919: func() mtproto.TLObject { // 0x52518bcf
		return &TLStatusSetSessionOnline{
			Constructor: 1381075919,
		}
	},
	631663196: func() mtproto.TLObject { // 0x25a66a5c
		return &TLStatusSetSessionOffline{
			Constructor: 631663196,
		}
	},
	-406788659: func() mtproto.TLObject { // 0xe7c0e5cd
		return &TLStatusGetUserOnlineSessions{
			Constructor: -406788659,
		}
	},
	-2009385532: func() mtproto.TLObject { // 0x883b35c4
		return &TLStatusGetUsersOnlineSessionsList{
			Constructor: -2009385532,
		}
	},
}

func NewTLObjectByClassID(classId int32) mtproto.TLObject {
	f, ok := clazzIdRegisters2[classId]
	if !ok {
		return nil
	}
	return f()
}

func CheckClassID(classId int32) (ok bool) {
	_, ok = clazzIdRegisters2[classId]
	return
}

//----------------------------------------------------------------------------------------------------------------

///////////////////////////////////////////////////////////////////////////////
// SessionEntry <--
//  + TL_SessionEntry
//

func (m *SessionEntry) Encode(layer int32) []byte {
	predicateName := m.PredicateName
	if predicateName == "" {
		if n, ok := clazzIdNameRegisters2[int32(m.Constructor)]; ok {
			predicateName = n
		}
	}

	var (
		xBuf []byte
	)

	switch predicateName {
	case Predicate_sessionEntry:
		t := m.To_SessionEntry()
		xBuf = t.Encode(layer)

	default:
		// logx.Errorf("invalid predicate error: %s",  m.PredicateName)
		return []byte{}
	}

	return xBuf
}

func (m *SessionEntry) CalcByteSize(layer int32) int {
	return 0
}

func (m *SessionEntry) Decode(dBuf *mtproto.DecodeBuf) error {
	m.Constructor = TLConstructor(dBuf.Int())
	switch uint32(m.Constructor) {
	case 0x1764ac31:
		m2 := MakeTLSessionEntry(m)
		m2.Decode(dBuf)

	default:
		return fmt.Errorf("invalid constructorId: 0x%x", uint32(m.Constructor))
	}
	return dBuf.GetError()
}

func (m *SessionEntry) DebugString() string {
	switch m.PredicateName {
	case Predicate_sessionEntry:
		t := m.To_SessionEntry()
		return t.DebugString()

	default:
		return "{}"
	}
}

// To_SessionEntry
func (m *SessionEntry) To_SessionEntry() *TLSessionEntry {
	m.PredicateName = Predicate_sessionEntry
	return &TLSessionEntry{
		Data2: m,
	}
}

// MakeTLSessionEntry
func MakeTLSessionEntry(data2 *SessionEntry) *TLSessionEntry {
	if data2 == nil {
		return &TLSessionEntry{Data2: &SessionEntry{
			PredicateName: Predicate_sessionEntry,
		}}
	} else {
		data2.PredicateName = Predicate_sessionEntry
		return &TLSessionEntry{Data2: data2}
	}
}

func (m *TLSessionEntry) To_SessionEntry() *SessionEntry {
	m.Data2.PredicateName = Predicate_sessionEntry
	return m.Data2
}

func (m *TLSessionEntry) SetUserId(v int64) { m.Data2.UserId = v }
func (m *TLSessionEntry) GetUserId() int64  { return m.Data2.UserId }

func (m *TLSessionEntry) SetAuthKeyId(v int64) { m.Data2.AuthKeyId = v }
func (m *TLSessionEntry) GetAuthKeyId() int64  { return m.Data2.AuthKeyId }

func (m *TLSessionEntry) SetGateway(v string) { m.Data2.Gateway = v }
func (m *TLSessionEntry) GetGateway() string  { return m.Data2.Gateway }

func (m *TLSessionEntry) SetExpired(v int64) { m.Data2.Expired = v }
func (m *TLSessionEntry) GetExpired() int64  { return m.Data2.Expired }

func (m *TLSessionEntry) SetLayer(v int32) { m.Data2.Layer = v }
func (m *TLSessionEntry) GetLayer() int32  { return m.Data2.Layer }

func (m *TLSessionEntry) SetPermAuthKeyId(v int64) { m.Data2.PermAuthKeyId = v }
func (m *TLSessionEntry) GetPermAuthKeyId() int64  { return m.Data2.PermAuthKeyId }

func (m *TLSessionEntry) SetClient(v string) { m.Data2.Client = v }
func (m *TLSessionEntry) GetClient() string  { return m.Data2.Client }

func (m *TLSessionEntry) GetPredicateName() string {
	return Predicate_sessionEntry
}

func (m *TLSessionEntry) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)

	var encodeF = map[uint32]func() []byte{
		0x1764ac31: func() []byte {
			x.UInt(0x1764ac31)

			x.Long(m.GetUserId())
			x.Long(m.GetAuthKeyId())
			x.String(m.GetGateway())
			x.Long(m.GetExpired())
			x.Int(m.GetLayer())
			x.Long(m.GetPermAuthKeyId())
			x.String(m.GetClient())
			return x.GetBuf()
		},
	}

	clazzId := GetClazzID(Predicate_sessionEntry, int(layer))
	if f, ok := encodeF[uint32(clazzId)]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		// log.Errorf("not found clazzId by (%s, %d)", Predicate_sessionEntry, layer)
		return x.GetBuf()
	}

	return x.GetBuf()
}

func (m *TLSessionEntry) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLSessionEntry) Decode(dBuf *mtproto.DecodeBuf) error {
	var decodeF = map[uint32]func() error{
		0x1764ac31: func() error {
			m.SetUserId(dBuf.Long())
			m.SetAuthKeyId(dBuf.Long())
			m.SetGateway(dBuf.String())
			m.SetExpired(dBuf.Long())
			m.SetLayer(dBuf.Int())
			m.SetPermAuthKeyId(dBuf.Long())
			m.SetClient(dBuf.String())
			return dBuf.GetError()
		},
	}

	if f, ok := decodeF[uint32(m.Data2.Constructor)]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", uint32(m.Data2.Constructor))
	}
}

func (m *TLSessionEntry) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

///////////////////////////////////////////////////////////////////////////////
// UserSessionEntryList <--
//  + TL_UserSessionEntryList
//

func (m *UserSessionEntryList) Encode(layer int32) []byte {
	predicateName := m.PredicateName
	if predicateName == "" {
		if n, ok := clazzIdNameRegisters2[int32(m.Constructor)]; ok {
			predicateName = n
		}
	}

	var (
		xBuf []byte
	)

	switch predicateName {
	case Predicate_userSessionEntryList:
		t := m.To_UserSessionEntryList()
		xBuf = t.Encode(layer)

	default:
		// logx.Errorf("invalid predicate error: %s",  m.PredicateName)
		return []byte{}
	}

	return xBuf
}

func (m *UserSessionEntryList) CalcByteSize(layer int32) int {
	return 0
}

func (m *UserSessionEntryList) Decode(dBuf *mtproto.DecodeBuf) error {
	m.Constructor = TLConstructor(dBuf.Int())
	switch uint32(m.Constructor) {
	case 0xefecb398:
		m2 := MakeTLUserSessionEntryList(m)
		m2.Decode(dBuf)

	default:
		return fmt.Errorf("invalid constructorId: 0x%x", uint32(m.Constructor))
	}
	return dBuf.GetError()
}

func (m *UserSessionEntryList) DebugString() string {
	switch m.PredicateName {
	case Predicate_userSessionEntryList:
		t := m.To_UserSessionEntryList()
		return t.DebugString()

	default:
		return "{}"
	}
}

// To_UserSessionEntryList
func (m *UserSessionEntryList) To_UserSessionEntryList() *TLUserSessionEntryList {
	m.PredicateName = Predicate_userSessionEntryList
	return &TLUserSessionEntryList{
		Data2: m,
	}
}

// MakeTLUserSessionEntryList
func MakeTLUserSessionEntryList(data2 *UserSessionEntryList) *TLUserSessionEntryList {
	if data2 == nil {
		return &TLUserSessionEntryList{Data2: &UserSessionEntryList{
			PredicateName: Predicate_userSessionEntryList,
		}}
	} else {
		data2.PredicateName = Predicate_userSessionEntryList
		return &TLUserSessionEntryList{Data2: data2}
	}
}

func (m *TLUserSessionEntryList) To_UserSessionEntryList() *UserSessionEntryList {
	m.Data2.PredicateName = Predicate_userSessionEntryList
	return m.Data2
}

func (m *TLUserSessionEntryList) SetUserId(v int64) { m.Data2.UserId = v }
func (m *TLUserSessionEntryList) GetUserId() int64  { return m.Data2.UserId }

func (m *TLUserSessionEntryList) SetUserSessions(v []*SessionEntry) { m.Data2.UserSessions = v }
func (m *TLUserSessionEntryList) GetUserSessions() []*SessionEntry  { return m.Data2.UserSessions }

func (m *TLUserSessionEntryList) GetPredicateName() string {
	return Predicate_userSessionEntryList
}

func (m *TLUserSessionEntryList) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)

	var encodeF = map[uint32]func() []byte{
		0xefecb398: func() []byte {
			x.UInt(0xefecb398)

			x.Long(m.GetUserId())

			x.Int(int32(mtproto.CRC32_vector))
			x.Int(int32(len(m.GetUserSessions())))
			for _, v := range m.GetUserSessions() {
				x.Bytes((*v).Encode(layer))
			}

			return x.GetBuf()
		},
	}

	clazzId := GetClazzID(Predicate_userSessionEntryList, int(layer))
	if f, ok := encodeF[uint32(clazzId)]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		// log.Errorf("not found clazzId by (%s, %d)", Predicate_userSessionEntryList, layer)
		return x.GetBuf()
	}

	return x.GetBuf()
}

func (m *TLUserSessionEntryList) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserSessionEntryList) Decode(dBuf *mtproto.DecodeBuf) error {
	var decodeF = map[uint32]func() error{
		0xefecb398: func() error {
			m.SetUserId(dBuf.Long())
			c1 := dBuf.Int()
			if c1 != int32(mtproto.CRC32_vector) {
				// dBuf.err = fmt.Errorf("invalid mtproto.CRC32_vector, c%d: %d", 1, c1)
				return fmt.Errorf("invalid mtproto.CRC32_vector, c%d: %d", 1, c1)
			}
			l1 := dBuf.Int()
			v1 := make([]*SessionEntry, l1)
			for i := int32(0); i < l1; i++ {
				v1[i] = &SessionEntry{}
				v1[i].Decode(dBuf)
			}
			m.SetUserSessions(v1)

			return dBuf.GetError()
		},
	}

	if f, ok := decodeF[uint32(m.Data2.Constructor)]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", uint32(m.Data2.Constructor))
	}
}

func (m *TLUserSessionEntryList) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

//----------------------------------------------------------------------------------------------------------------
// TLStatusSetSessionOnline
///////////////////////////////////////////////////////////////////////////////

func (m *TLStatusSetSessionOnline) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_status_setSessionOnline))

	switch uint32(m.Constructor) {
	case 0x52518bcf:
		x.UInt(0x52518bcf)

		// no flags

		x.Long(m.GetUserId())
		x.Bytes(m.GetSession().Encode(layer))

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLStatusSetSessionOnline) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLStatusSetSessionOnline) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x52518bcf:

		// not has flags

		m.UserId = dBuf.Long()

		m2 := &SessionEntry{}
		m2.Decode(dBuf)
		m.Session = m2

		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLStatusSetSessionOnline) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLStatusSetSessionOffline
///////////////////////////////////////////////////////////////////////////////

func (m *TLStatusSetSessionOffline) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_status_setSessionOffline))

	switch uint32(m.Constructor) {
	case 0x25a66a5c:
		x.UInt(0x25a66a5c)

		// no flags

		x.Long(m.GetUserId())
		x.Long(m.GetAuthKeyId())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLStatusSetSessionOffline) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLStatusSetSessionOffline) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x25a66a5c:

		// not has flags

		m.UserId = dBuf.Long()
		m.AuthKeyId = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLStatusSetSessionOffline) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLStatusGetUserOnlineSessions
///////////////////////////////////////////////////////////////////////////////

func (m *TLStatusGetUserOnlineSessions) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_status_getUserOnlineSessions))

	switch uint32(m.Constructor) {
	case 0xe7c0e5cd:
		x.UInt(0xe7c0e5cd)

		// no flags

		x.Long(m.GetUserId())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLStatusGetUserOnlineSessions) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLStatusGetUserOnlineSessions) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xe7c0e5cd:

		// not has flags

		m.UserId = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLStatusGetUserOnlineSessions) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLStatusGetUsersOnlineSessionsList
///////////////////////////////////////////////////////////////////////////////

func (m *TLStatusGetUsersOnlineSessionsList) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_status_getUsersOnlineSessionsList))

	switch uint32(m.Constructor) {
	case 0x883b35c4:
		x.UInt(0x883b35c4)

		// no flags

		x.VectorLong(m.GetUsers())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLStatusGetUsersOnlineSessionsList) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLStatusGetUsersOnlineSessionsList) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x883b35c4:

		// not has flags

		m.Users = dBuf.VectorLong()

		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLStatusGetUsersOnlineSessionsList) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// ----------------------------------------------------------------------------------------------------------------
// Vector_UserSessionEntryList
// /////////////////////////////////////////////////////////////////////////////
func (m *Vector_UserSessionEntryList) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	x.Int(int32(mtproto.CRC32_vector))
	x.Int(int32(len(m.Datas)))
	for _, v := range m.Datas {
		x.Bytes((*v).Encode(layer))
	}

	return x.GetBuf()
}

func (m *Vector_UserSessionEntryList) Decode(dBuf *mtproto.DecodeBuf) error {
	dBuf.Int() // TODO(@benqi): Check crc32 invalid
	l1 := dBuf.Int()
	m.Datas = make([]*UserSessionEntryList, l1)
	for i := int32(0); i < l1; i++ {
		m.Datas[i] = new(UserSessionEntryList)
		(*m.Datas[i]).Decode(dBuf)
	}

	return dBuf.GetError()
}

func (m *Vector_UserSessionEntryList) CalcByteSize(layer int32) int {
	return 0
}

func (m *Vector_UserSessionEntryList) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}
