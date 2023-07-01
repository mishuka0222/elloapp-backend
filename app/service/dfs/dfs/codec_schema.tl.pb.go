// ConstructorList
// RequestList

package dfs

import (
	"fmt"

	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/gogo/protobuf/types"
)

//////////////////////////////////////////////////////////////////////////////////////////

var _ *types.Int32Value
var _ *mtproto.Bool
var _ fmt.GoStringer

var clazzIdRegisters2 = map[int32]func() mtproto.TLObject{
	// Constructor

	// Method
	440942855: func() mtproto.TLObject { // 0x1a484107
		return &TLDfsWriteFilePartData{
			Constructor: 440942855,
		}
	},
	605082018: func() mtproto.TLObject { // 0x2410d1a2
		return &TLDfsUploadPhotoFileV2{
			Constructor: 605082018,
		}
	},
	-870473038: func() mtproto.TLObject { // 0xcc1da2b2
		return &TLDfsUploadProfilePhotoFileV2{
			Constructor: -870473038,
		}
	},
	2043921699: func() mtproto.TLObject { // 0x79d3c523
		return &TLDfsUploadEncryptedFileV2{
			Constructor: 2043921699,
		}
	},
	-692064706: func() mtproto.TLObject { // 0xd6bfee3e
		return &TLDfsDownloadFile{
			Constructor: -692064706,
		}
	},
	1983081911: func() mtproto.TLObject { // 0x76336db7
		return &TLDfsUploadDocumentFileV2{
			Constructor: 1983081911,
		}
	},
	1103416576: func() mtproto.TLObject { // 0x41c4cd00
		return &TLDfsUploadGifDocumentMedia{
			Constructor: 1103416576,
		}
	},
	-1566246888: func() mtproto.TLObject { // 0xa2a4f818
		return &TLDfsUploadMp4DocumentMedia{
			Constructor: -1566246888,
		}
	},
	-1046081450: func() mtproto.TLObject { // 0xc1a61056
		return &TLDfsUploadWallPaperFile{
			Constructor: -1046081450,
		}
	},
	-559525993: func() mtproto.TLObject { // 0xdea64f97
		return &TLDfsUploadThemeFile{
			Constructor: -559525993,
		}
	},
	45335985: func() mtproto.TLObject { // 0x2b3c5b1
		return &TLDfsUploadRingtoneFile{
			Constructor: 45335985,
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

//----------------------------------------------------------------------------------------------------------------
// TLDfsWriteFilePartData
///////////////////////////////////////////////////////////////////////////////

func (m *TLDfsWriteFilePartData) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_dfs_writeFilePartData))

	switch uint32(m.Constructor) {
	case 0x1a484107:
		x.UInt(0x1a484107)

		// set flags
		var flags uint32 = 0

		if m.GetBig() == true {
			flags |= 1 << 0
		}
		if m.GetFileTotalParts() != nil {
			flags |= 1 << 1
		}

		x.UInt(flags)

		// flags Debug by @benqi
		x.Long(m.GetCreator())
		x.Long(m.GetFileId())
		x.Int(m.GetFilePart())
		x.StringBytes(m.GetBytes())
		if m.GetFileTotalParts() != nil {
			x.Int(m.GetFileTotalParts().Value)
		}

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLDfsWriteFilePartData) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLDfsWriteFilePartData) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x1a484107:

		flags := dBuf.UInt()
		_ = flags

		// flags Debug by @benqi
		m.Creator = dBuf.Long()
		m.FileId = dBuf.Long()
		m.FilePart = dBuf.Int()
		m.Bytes = dBuf.StringBytes()
		if (flags & (1 << 0)) != 0 {
			m.Big = true
		}
		if (flags & (1 << 1)) != 0 {
			m.FileTotalParts = &types.Int32Value{Value: dBuf.Int()}
		}

		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLDfsWriteFilePartData) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLDfsUploadPhotoFileV2
///////////////////////////////////////////////////////////////////////////////

func (m *TLDfsUploadPhotoFileV2) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_dfs_uploadPhotoFileV2))

	switch uint32(m.Constructor) {
	case 0x2410d1a2:
		x.UInt(0x2410d1a2)

		// no flags

		x.Long(m.GetCreator())
		x.Bytes(m.GetFile().Encode(layer))

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLDfsUploadPhotoFileV2) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLDfsUploadPhotoFileV2) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x2410d1a2:

		// not has flags

		m.Creator = dBuf.Long()

		m2 := &mtproto.InputFile{}
		m2.Decode(dBuf)
		m.File = m2

		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLDfsUploadPhotoFileV2) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLDfsUploadProfilePhotoFileV2
///////////////////////////////////////////////////////////////////////////////

func (m *TLDfsUploadProfilePhotoFileV2) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_dfs_uploadProfilePhotoFileV2))

	switch uint32(m.Constructor) {
	case 0xcc1da2b2:
		x.UInt(0xcc1da2b2)

		// set flags
		var flags uint32 = 0

		if m.GetFile() != nil {
			flags |= 1 << 0
		}
		if m.GetVideo() != nil {
			flags |= 1 << 1
		}
		if m.GetVideoStartTs() != nil {
			flags |= 1 << 2
		}

		x.UInt(flags)

		// flags Debug by @benqi
		x.Long(m.GetCreator())
		if m.GetFile() != nil {
			x.Bytes(m.GetFile().Encode(layer))
		}

		if m.GetVideo() != nil {
			x.Bytes(m.GetVideo().Encode(layer))
		}

		if m.GetVideoStartTs() != nil {
			x.Double(m.GetVideoStartTs().Value)
		}

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLDfsUploadProfilePhotoFileV2) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLDfsUploadProfilePhotoFileV2) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xcc1da2b2:

		flags := dBuf.UInt()
		_ = flags

		// flags Debug by @benqi
		m.Creator = dBuf.Long()
		if (flags & (1 << 0)) != 0 {
			m3 := &mtproto.InputFile{}
			m3.Decode(dBuf)
			m.File = m3
		}
		if (flags & (1 << 1)) != 0 {
			m4 := &mtproto.InputFile{}
			m4.Decode(dBuf)
			m.Video = m4
		}
		if (flags & (1 << 2)) != 0 {
			m.VideoStartTs = &types.DoubleValue{Value: dBuf.Double()}
		}

		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLDfsUploadProfilePhotoFileV2) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLDfsUploadEncryptedFileV2
///////////////////////////////////////////////////////////////////////////////

func (m *TLDfsUploadEncryptedFileV2) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_dfs_uploadEncryptedFileV2))

	switch uint32(m.Constructor) {
	case 0x79d3c523:
		x.UInt(0x79d3c523)

		// no flags

		x.Long(m.GetCreator())
		x.Bytes(m.GetFile().Encode(layer))

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLDfsUploadEncryptedFileV2) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLDfsUploadEncryptedFileV2) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x79d3c523:

		// not has flags

		m.Creator = dBuf.Long()

		m2 := &mtproto.InputEncryptedFile{}
		m2.Decode(dBuf)
		m.File = m2

		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLDfsUploadEncryptedFileV2) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLDfsDownloadFile
///////////////////////////////////////////////////////////////////////////////

func (m *TLDfsDownloadFile) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_dfs_downloadFile))

	switch uint32(m.Constructor) {
	case 0xd6bfee3e:
		x.UInt(0xd6bfee3e)

		// no flags

		x.Bytes(m.GetLocation().Encode(layer))
		x.Long(m.GetOffset())
		x.Int(m.GetLimit())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLDfsDownloadFile) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLDfsDownloadFile) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xd6bfee3e:

		// not has flags

		m1 := &mtproto.InputFileLocation{}
		m1.Decode(dBuf)
		m.Location = m1

		m.Offset = dBuf.Long()
		m.Limit = dBuf.Int()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLDfsDownloadFile) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLDfsUploadDocumentFileV2
///////////////////////////////////////////////////////////////////////////////

func (m *TLDfsUploadDocumentFileV2) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_dfs_uploadDocumentFileV2))

	switch uint32(m.Constructor) {
	case 0x76336db7:
		x.UInt(0x76336db7)

		// no flags

		x.Long(m.GetCreator())
		x.Bytes(m.GetMedia().Encode(layer))

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLDfsUploadDocumentFileV2) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLDfsUploadDocumentFileV2) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x76336db7:

		// not has flags

		m.Creator = dBuf.Long()

		m2 := &mtproto.InputMedia{}
		m2.Decode(dBuf)
		m.Media = m2

		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLDfsUploadDocumentFileV2) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLDfsUploadGifDocumentMedia
///////////////////////////////////////////////////////////////////////////////

func (m *TLDfsUploadGifDocumentMedia) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_dfs_uploadGifDocumentMedia))

	switch uint32(m.Constructor) {
	case 0x41c4cd00:
		x.UInt(0x41c4cd00)

		// no flags

		x.Long(m.GetCreator())
		x.Bytes(m.GetMedia().Encode(layer))

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLDfsUploadGifDocumentMedia) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLDfsUploadGifDocumentMedia) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x41c4cd00:

		// not has flags

		m.Creator = dBuf.Long()

		m2 := &mtproto.InputMedia{}
		m2.Decode(dBuf)
		m.Media = m2

		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLDfsUploadGifDocumentMedia) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLDfsUploadMp4DocumentMedia
///////////////////////////////////////////////////////////////////////////////

func (m *TLDfsUploadMp4DocumentMedia) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_dfs_uploadMp4DocumentMedia))

	switch uint32(m.Constructor) {
	case 0xa2a4f818:
		x.UInt(0xa2a4f818)

		// no flags

		x.Long(m.GetCreator())
		x.Bytes(m.GetMedia().Encode(layer))

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLDfsUploadMp4DocumentMedia) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLDfsUploadMp4DocumentMedia) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xa2a4f818:

		// not has flags

		m.Creator = dBuf.Long()

		m2 := &mtproto.InputMedia{}
		m2.Decode(dBuf)
		m.Media = m2

		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLDfsUploadMp4DocumentMedia) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLDfsUploadWallPaperFile
///////////////////////////////////////////////////////////////////////////////

func (m *TLDfsUploadWallPaperFile) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_dfs_uploadWallPaperFile))

	switch uint32(m.Constructor) {
	case 0xc1a61056:
		x.UInt(0xc1a61056)

		// no flags

		x.Long(m.GetCreator())
		x.Bytes(m.GetFile().Encode(layer))
		x.String(m.GetMimeType())
		x.Bytes(m.GetAdmin().Encode(layer))

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLDfsUploadWallPaperFile) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLDfsUploadWallPaperFile) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xc1a61056:

		// not has flags

		m.Creator = dBuf.Long()

		m2 := &mtproto.InputFile{}
		m2.Decode(dBuf)
		m.File = m2

		m.MimeType = dBuf.String()

		m4 := &mtproto.Bool{}
		m4.Decode(dBuf)
		m.Admin = m4

		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLDfsUploadWallPaperFile) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLDfsUploadThemeFile
///////////////////////////////////////////////////////////////////////////////

func (m *TLDfsUploadThemeFile) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_dfs_uploadThemeFile))

	switch uint32(m.Constructor) {
	case 0xdea64f97:
		x.UInt(0xdea64f97)

		// set flags
		var flags uint32 = 0

		if m.GetThumb() != nil {
			flags |= 1 << 0
		}

		x.UInt(flags)

		// flags Debug by @benqi
		x.Long(m.GetCreator())
		x.Bytes(m.GetFile().Encode(layer))
		if m.GetThumb() != nil {
			x.Bytes(m.GetThumb().Encode(layer))
		}

		x.String(m.GetMimeType())
		x.String(m.GetFileName())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLDfsUploadThemeFile) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLDfsUploadThemeFile) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xdea64f97:

		flags := dBuf.UInt()
		_ = flags

		// flags Debug by @benqi
		m.Creator = dBuf.Long()

		m3 := &mtproto.InputFile{}
		m3.Decode(dBuf)
		m.File = m3

		if (flags & (1 << 0)) != 0 {
			m4 := &mtproto.InputFile{}
			m4.Decode(dBuf)
			m.Thumb = m4
		}
		m.MimeType = dBuf.String()
		m.FileName = dBuf.String()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLDfsUploadThemeFile) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLDfsUploadRingtoneFile
///////////////////////////////////////////////////////////////////////////////

func (m *TLDfsUploadRingtoneFile) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_dfs_uploadRingtoneFile))

	switch uint32(m.Constructor) {
	case 0x2b3c5b1:
		x.UInt(0x2b3c5b1)

		// no flags

		x.Long(m.GetCreator())
		x.Bytes(m.GetFile().Encode(layer))
		x.String(m.GetMimeType())
		x.String(m.GetFileName())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLDfsUploadRingtoneFile) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLDfsUploadRingtoneFile) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x2b3c5b1:

		// not has flags

		m.Creator = dBuf.Long()

		m2 := &mtproto.InputFile{}
		m2.Decode(dBuf)
		m.File = m2

		m.MimeType = dBuf.String()
		m.FileName = dBuf.String()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLDfsUploadRingtoneFile) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

//----------------------------------------------------------------------------------------------------------------
