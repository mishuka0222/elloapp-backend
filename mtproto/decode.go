package mtproto

import (
	"encoding/binary"
	"errors"
	"fmt"
	"math"
	"math/big"
)

type DecodeBuf struct {
	buf  []byte
	off  int
	size int
	err  error
}

func NewDecodeBuf(b []byte) *DecodeBuf {
	return &DecodeBuf{b, 0, len(b), nil}
}

func (m *DecodeBuf) GetOffset() int {
	return m.off
}

func (m *DecodeBuf) GetSize() int {
	return m.size
}

func (m *DecodeBuf) GetError() error {
	return m.err
}

func (m *DecodeBuf) SetError(err error) {
	m.err = err
}

func (m *DecodeBuf) Long() int64 {
	if m.err != nil {
		return 0
	}
	if m.off+8 > m.size {
		m.err = errors.New("DecodeLong")
		return 0
	}
	x := int64(binary.LittleEndian.Uint64(m.buf[m.off : m.off+8]))
	m.off += 8
	return x
}

func (m *DecodeBuf) Double() float64 {
	if m.err != nil {
		return 0
	}
	if m.off+8 > m.size {
		m.err = errors.New("DecodeDouble")
		return 0
	}
	x := math.Float64frombits(binary.LittleEndian.Uint64(m.buf[m.off : m.off+8]))
	m.off += 8
	return x
}

func (m *DecodeBuf) Int() int32 {
	if m.err != nil {
		return 0
	}
	if m.off+4 > m.size {
		m.err = errors.New("DecodeInt")
		return 0
	}
	x := binary.LittleEndian.Uint32(m.buf[m.off : m.off+4])
	m.off += 4
	return int32(x)
}

func (m *DecodeBuf) UInt() uint32 {
	if m.err != nil {
		return 0
	}
	if m.off+4 > m.size {
		m.err = errors.New("DecodeUInt")
		return 0
	}
	x := binary.LittleEndian.Uint32(m.buf[m.off : m.off+4])
	m.off += 4
	return x
}

func (m *DecodeBuf) Bytes(size int) []byte {
	if m.err != nil {
		return nil
	}
	if m.off+size > m.size {
		m.err = errors.New("DecodeBytes")
		return nil
	}
	x := make([]byte, size)
	copy(x, m.buf[m.off:m.off+size])
	m.off += size
	return x
}

// StringBytes
/*
   public String readString(boolean exception) {
       int startReadPosition = getPosition();
       try {
           int sl = 1;
           int l = getIntFromByte(buffer.get());
           if (l >= 254) {
               l = getIntFromByte(buffer.get()) | (getIntFromByte(buffer.get()) << 8) | (getIntFromByte(buffer.get()) << 16);
               sl = 4;
           }
           byte[] b = new byte[l];
           buffer.get(b);
           int i = sl;
           while ((l + i) % 4 != 0) {
               buffer.get();
               i++;
           }
           return new String(b, "UTF-8");
       } catch (Exception e) {
           if (exception) {
               throw new RuntimeException("read string error", e);
           } else {
               if (BuildVars.LOGS_ENABLED) {
                   FileLog.e("read string error");
               }
           }
           position(startReadPosition);
       }
       return "";
   }
*/
func (m *DecodeBuf) StringBytes() []byte {
	if m.err != nil {
		return nil
	}
	var size, padding int

	if m.off+1 > m.size {
		m.err = errors.New("DecodeStringBytes")
		return nil
	}
	size = int(m.buf[m.off])
	m.off++
	padding = (4 - ((size + 1) % 4)) & 3
	if size == 254 {
		if m.off+3 > m.size {
			m.err = errors.New("DecodeStringBytes")
			return nil
		}
		size = int(m.buf[m.off]) | int(m.buf[m.off+1])<<8 | int(m.buf[m.off+2])<<16
		m.off += 3
		padding = (4 - size%4) & 3
	}

	if m.off+size > m.size {
		m.err = fmt.Errorf("DecodeStringBytes - Wrong size, (%d, %d, %d)", m.off, size, m.size)
		return nil
	}
	x := make([]byte, size)
	copy(x, m.buf[m.off:m.off+size])
	m.off += size

	if m.off+padding > m.size {
		m.err = errors.New("DecodeStringBytes: Wrong padding")
		return nil
	}
	m.off += padding

	return x
}

func (m *DecodeBuf) String() string {
	b := m.StringBytes()
	if m.err != nil {
		return ""
	}
	x := string(b)
	return x
}

func (m *DecodeBuf) BigInt() *big.Int {
	b := m.StringBytes()
	if m.err != nil {
		return nil
	}
	y := make([]byte, len(b)+1)
	y[0] = 0
	copy(y[1:], b)
	x := new(big.Int).SetBytes(y)
	return x
}

func (m *DecodeBuf) VectorInt() []int32 {
	constructor := m.Int()
	if m.err != nil {
		return nil
	}
	if constructor != int32(CRC32_vector) {
		m.err = fmt.Errorf("DecodeVectorInt: Wrong constructor (0x%08x)", constructor)
		return nil
	}
	size := m.Int()
	if m.err != nil {
		return nil
	}
	if size < 0 {
		m.err = errors.New("DecodeVectorInt: Wrong size")
		return nil
	}
	x := make([]int32, size)
	i := int32(0)
	for i < size {
		y := m.Int()
		if m.err != nil {
			return nil
		}
		x[i] = y
		i++
	}
	return x
}

func (m *DecodeBuf) VectorLong() []int64 {
	constructor := m.Int()
	if m.err != nil {
		return nil
	}
	if constructor != int32(CRC32_vector) {
		m.err = fmt.Errorf("DecodeVectorLong: Wrong constructor (0x%08x)", constructor)
		return nil
	}
	size := m.Int()
	if m.err != nil {
		return nil
	}
	if size < 0 {
		m.err = errors.New("DecodeVectorLong: Wrong size")
		return nil
	}
	x := make([]int64, size)
	i := int32(0)
	for i < size {
		y := m.Long()
		if m.err != nil {
			return nil
		}
		x[i] = y
		i++
	}
	return x
}

func (m *DecodeBuf) VectorString() []string {
	constructor := m.Int()
	if m.err != nil {
		return nil
	}
	if constructor != int32(CRC32_vector) {
		m.err = fmt.Errorf("DecodeVectorString: Wrong constructor (0x%08x)", constructor)
		return nil
	}
	size := m.Int()
	if m.err != nil {
		return nil
	}
	if size < 0 {
		m.err = errors.New("DecodeVectorString: Wrong size")
		return nil
	}
	x := make([]string, size)
	i := int32(0)
	for i < size {
		y := m.String()
		if m.err != nil {
			return nil
		}
		x[i] = y
		i++
	}
	return x
}

func (m *DecodeBuf) VectorBytes() [][]byte {
	constructor := m.Int()
	if m.err != nil {
		return nil
	}
	if constructor != int32(CRC32_vector) {
		m.err = fmt.Errorf("DecodeVectorBytes: Wrong constructor (0x%08x)", constructor)
		return nil
	}
	size := m.Int()
	if m.err != nil {
		return nil
	}
	if size < 0 {
		m.err = errors.New("DecodeVectorBytes: Wrong size")
		return nil
	}
	x := make([][]byte, size)
	i := int32(0)
	for i < size {
		y := m.StringBytes()
		if m.err != nil {
			return nil
		}
		x[i] = y
		i++
	}
	return x
}

func (m *DecodeBuf) Bool() bool {
	constructor := m.Int()
	if m.err != nil {
		return false
	}
	switch constructor {
	case int32(CRC32_boolTrue):
		return true
	case int32(CRC32_boolFalse):
		return false
	}
	return false
}

/*
func (m *DecodeBuf) Vector() []TLObject {
	constructor := m.Int()
	if m.err != nil {
		return nil
	}
	if constructor != int32(TLConstructor_CRC32_vector) {
		m.err = fmt.Errorf("DecodeVector: Wrong constructor (0x%08x)", constructor)
		return nil
	}
	size := m.Int()
	if m.err != nil {
		return nil
	}
	if size < 0 {
		m.err = errors.New("DecodeVector: Wrong size")
		return nil
	}
	x := make([]TLObject, size)
	i := int32(0)
	for i < size {
		y := m.Object()
		if m.err != nil {
			return nil
		}
		x[i] = y
		i++
	}
	return x
}
*/

func (m *DecodeBuf) Object() (r TLObject) {
	classID := m.Int()
	if m.err != nil {
		return nil
	}

	r = NewTLObjectByClassID(classID)
	if r == nil {
		m.err = fmt.Errorf("can't find registered classId: %x", uint32(classID))
		return nil
	}

	// log.Infof("newTLObjectByClassID, classID: %x", uint32(classID))
	err := r.(TLObject).Decode(m)
	if err != nil {
		m.err = fmt.Errorf("object(%x) decode error: %v", uint32(classID), err)
	}

	return
}

/*
func (d *DecodeBuf) dump() {
	fmt.Println(hex.Dump(d.buf[d.off:d.size]))
}

/*
func toBool(x TLObject) bool {
	_, ok := x.(TL_boolTrue)
	return ok
}
*/
