// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: gateway.tl.proto

package gateway

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	mtproto "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type TLConstructor int32

const (
	CRC32_UNKNOWN                   TLConstructor = 0
	CRC32_gateway_sendDataToGateway TLConstructor = 645953552
)

var TLConstructor_name = map[int32]string{
	0:         "CRC32_UNKNOWN",
	645953552: "CRC32_gateway_sendDataToGateway",
}

var TLConstructor_value = map[string]int32{
	"CRC32_UNKNOWN":                   0,
	"CRC32_gateway_sendDataToGateway": 645953552,
}

func (x TLConstructor) String() string {
	return proto.EnumName(TLConstructor_name, int32(x))
}

func (TLConstructor) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c10c1e4729b66838, []int{0}
}

// --------------------------------------------------------------------------------------------
// gateway.sendDataToGateway auth_key_id:long session_id:long payload:bytes = Bool;
type TLGatewaySendDataToGateway struct {
	Constructor          TLConstructor `protobuf:"varint,1,opt,name=constructor,proto3,enum=gateway.TLConstructor" json:"constructor,omitempty"`
	AuthKeyId            int64         `protobuf:"varint,3,opt,name=auth_key_id,json=authKeyId,proto3" json:"auth_key_id,omitempty"`
	SessionId            int64         `protobuf:"varint,4,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	Payload              []byte        `protobuf:"bytes,5,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *TLGatewaySendDataToGateway) Reset()         { *m = TLGatewaySendDataToGateway{} }
func (m *TLGatewaySendDataToGateway) String() string { return proto.CompactTextString(m) }
func (*TLGatewaySendDataToGateway) ProtoMessage()    {}
func (*TLGatewaySendDataToGateway) Descriptor() ([]byte, []int) {
	return fileDescriptor_c10c1e4729b66838, []int{0}
}
func (m *TLGatewaySendDataToGateway) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TLGatewaySendDataToGateway) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TLGatewaySendDataToGateway.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TLGatewaySendDataToGateway) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TLGatewaySendDataToGateway.Merge(m, src)
}
func (m *TLGatewaySendDataToGateway) XXX_Size() int {
	return m.Size()
}
func (m *TLGatewaySendDataToGateway) XXX_DiscardUnknown() {
	xxx_messageInfo_TLGatewaySendDataToGateway.DiscardUnknown(m)
}

var xxx_messageInfo_TLGatewaySendDataToGateway proto.InternalMessageInfo

func (m *TLGatewaySendDataToGateway) GetConstructor() TLConstructor {
	if m != nil {
		return m.Constructor
	}
	return CRC32_UNKNOWN
}

func (m *TLGatewaySendDataToGateway) GetAuthKeyId() int64 {
	if m != nil {
		return m.AuthKeyId
	}
	return 0
}

func (m *TLGatewaySendDataToGateway) GetSessionId() int64 {
	if m != nil {
		return m.SessionId
	}
	return 0
}

func (m *TLGatewaySendDataToGateway) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func init() {
	proto.RegisterEnum("gateway.TLConstructor", TLConstructor_name, TLConstructor_value)
	proto.RegisterType((*TLGatewaySendDataToGateway)(nil), "gateway.TL_gateway_sendDataToGateway")
}

func init() { proto.RegisterFile("gateway.tl.proto", fileDescriptor_c10c1e4729b66838) }

var fileDescriptor_c10c1e4729b66838 = []byte{
	// 404 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0xcf, 0x8a, 0xda, 0x50,
	0x14, 0xc6, 0xbd, 0xd5, 0x56, 0x7a, 0xad, 0xc5, 0x86, 0x52, 0x62, 0x68, 0x53, 0x11, 0x4a, 0xa5,
	0x60, 0x02, 0xba, 0x69, 0xb7, 0xa6, 0x50, 0x44, 0xb1, 0x6d, 0x6a, 0x29, 0x74, 0x13, 0xae, 0xc9,
	0x35, 0x86, 0x26, 0x39, 0x97, 0x7b, 0x6f, 0x2a, 0x59, 0xce, 0x6e, 0x96, 0xf3, 0x0e, 0xf3, 0x04,
	0x03, 0xf3, 0x10, 0xb3, 0x9c, 0x47, 0x98, 0xf1, 0x09, 0x7c, 0x84, 0xc1, 0xfc, 0x21, 0xce, 0xc2,
	0x55, 0xce, 0xf9, 0x7e, 0x27, 0xe7, 0xbb, 0xdf, 0xc1, 0x1d, 0x9f, 0x48, 0xba, 0x25, 0xa9, 0x21,
	0x43, 0x83, 0x71, 0x90, 0xa0, 0x34, 0x0b, 0x45, 0x1b, 0xfa, 0x81, 0xdc, 0x24, 0x2b, 0xc3, 0x85,
	0xc8, 0xf4, 0xc1, 0x07, 0x33, 0xe3, 0xab, 0x64, 0x9d, 0x75, 0x59, 0x93, 0x55, 0xf9, 0x7f, 0x9a,
	0xee, 0x03, 0xf8, 0x21, 0xad, 0xa6, 0xb6, 0x9c, 0x30, 0x46, 0xb9, 0x28, 0xb8, 0x26, 0xdc, 0x0d,
	0x8d, 0xc8, 0xc1, 0xc8, 0x05, 0x4e, 0x1d, 0x99, 0x32, 0x5a, 0xb2, 0x6e, 0xc5, 0x24, 0x27, 0xb1,
	0x60, 0xc0, 0x65, 0x81, 0x5e, 0x57, 0x48, 0xa4, 0xb1, 0x9b, 0xab, 0xfd, 0x2b, 0x84, 0xdf, 0x2e,
	0xe7, 0x4e, 0xf1, 0x54, 0x47, 0xd0, 0xd8, 0xfb, 0x4a, 0x24, 0x59, 0xc2, 0xb7, 0x5c, 0x51, 0x3e,
	0xe3, 0x96, 0x0b, 0xb1, 0x90, 0x3c, 0x71, 0x25, 0x70, 0x15, 0xf5, 0xd0, 0xe0, 0xe5, 0xe8, 0x8d,
	0x51, 0xa6, 0x5d, 0xce, 0xad, 0x8a, 0xda, 0xc7, 0xa3, 0x8a, 0x8e, 0x5b, 0x24, 0x91, 0x1b, 0xe7,
	0x1f, 0x4d, 0x9d, 0xc0, 0x53, 0xeb, 0x3d, 0x34, 0xa8, 0xdb, 0xcf, 0x0f, 0xd2, 0x8c, 0xa6, 0x53,
	0x4f, 0x79, 0x87, 0xb1, 0xa0, 0x42, 0x04, 0x10, 0x1f, 0x70, 0x23, 0xc7, 0x85, 0x32, 0xf5, 0x14,
	0x15, 0x37, 0x19, 0x49, 0x43, 0x20, 0x9e, 0xfa, 0xb4, 0x87, 0x06, 0x2f, 0xec, 0xb2, 0xfd, 0xf4,
	0x13, 0xb7, 0x1f, 0xd9, 0x2a, 0xaf, 0x70, 0xdb, 0xb2, 0xad, 0xf1, 0xc8, 0xf9, 0xbd, 0x98, 0x2d,
	0xbe, 0xff, 0x59, 0x74, 0x6a, 0xca, 0x47, 0xfc, 0x3e, 0x97, 0x4e, 0x26, 0xeb, 0x5c, 0xec, 0xcf,
	0xae, 0x9f, 0x68, 0x8d, 0xf3, 0x4b, 0xbd, 0x36, 0x22, 0x18, 0xdb, 0x3f, 0xac, 0x32, 0xf3, 0x2f,
	0xdc, 0x3d, 0x7d, 0x90, 0x0f, 0x47, 0xd9, 0x4f, 0x6f, 0xd7, 0xda, 0x46, 0x24, 0xb3, 0x13, 0x1b,
	0x13, 0x80, 0xb0, 0x5f, 0x9b, 0xcc, 0xf6, 0xf7, 0x3a, 0xba, 0xd9, 0xe9, 0xe8, 0x76, 0xa7, 0xa3,
	0xbb, 0x9d, 0x8e, 0xfe, 0x7e, 0x91, 0x94, 0x44, 0x3e, 0x27, 0x91, 0x11, 0x80, 0x59, 0xd6, 0x43,
	0x41, 0xf9, 0x7f, 0xca, 0x4d, 0xc2, 0x98, 0x19, 0xc4, 0x92, 0xf2, 0x35, 0x71, 0xa9, 0x59, 0xf8,
	0x94, 0xdf, 0xd5, 0xb3, 0x6c, 0xf5, 0xf8, 0x21, 0x00, 0x00, 0xff, 0xff, 0xe4, 0x9f, 0xd9, 0x62,
	0x76, 0x02, 0x00, 0x00,
}

func (this *TLGatewaySendDataToGateway) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 8)
	s = append(s, "&gateway.TLGatewaySendDataToGateway{")
	s = append(s, "Constructor: "+fmt.Sprintf("%#v", this.Constructor)+",\n")
	s = append(s, "AuthKeyId: "+fmt.Sprintf("%#v", this.AuthKeyId)+",\n")
	s = append(s, "SessionId: "+fmt.Sprintf("%#v", this.SessionId)+",\n")
	s = append(s, "Payload: "+fmt.Sprintf("%#v", this.Payload)+",\n")
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringGatewayTl(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RPCGatewayClient is the client API for RPCGateway service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RPCGatewayClient interface {
	// gateway.sendDataToGateway auth_key_id:long session_id:long payload:bytes = Bool;
	GatewaySendDataToGateway(ctx context.Context, in *TLGatewaySendDataToGateway, opts ...grpc.CallOption) (*mtproto.Bool, error)
}

type rPCGatewayClient struct {
	cc *grpc.ClientConn
}

func NewRPCGatewayClient(cc *grpc.ClientConn) RPCGatewayClient {
	return &rPCGatewayClient{cc}
}

func (c *rPCGatewayClient) GatewaySendDataToGateway(ctx context.Context, in *TLGatewaySendDataToGateway, opts ...grpc.CallOption) (*mtproto.Bool, error) {
	out := new(mtproto.Bool)
	err := c.cc.Invoke(ctx, "/gateway.RPCGateway/gateway_sendDataToGateway", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RPCGatewayServer is the server API for RPCGateway service.
type RPCGatewayServer interface {
	// gateway.sendDataToGateway auth_key_id:long session_id:long payload:bytes = Bool;
	GatewaySendDataToGateway(context.Context, *TLGatewaySendDataToGateway) (*mtproto.Bool, error)
}

// UnimplementedRPCGatewayServer can be embedded to have forward compatible implementations.
type UnimplementedRPCGatewayServer struct {
}

func (*UnimplementedRPCGatewayServer) GatewaySendDataToGateway(ctx context.Context, req *TLGatewaySendDataToGateway) (*mtproto.Bool, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GatewaySendDataToGateway not implemented")
}

func RegisterRPCGatewayServer(s *grpc.Server, srv RPCGatewayServer) {
	s.RegisterService(&_RPCGateway_serviceDesc, srv)
}

func _RPCGateway_GatewaySendDataToGateway_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TLGatewaySendDataToGateway)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCGatewayServer).GatewaySendDataToGateway(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gateway.RPCGateway/GatewaySendDataToGateway",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCGatewayServer).GatewaySendDataToGateway(ctx, req.(*TLGatewaySendDataToGateway))
	}
	return interceptor(ctx, in, info, handler)
}

var _RPCGateway_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gateway.RPCGateway",
	HandlerType: (*RPCGatewayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "gateway_sendDataToGateway",
			Handler:    _RPCGateway_GatewaySendDataToGateway_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gateway.tl.proto",
}

func (m *TLGatewaySendDataToGateway) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TLGatewaySendDataToGateway) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TLGatewaySendDataToGateway) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Payload) > 0 {
		i -= len(m.Payload)
		copy(dAtA[i:], m.Payload)
		i = encodeVarintGatewayTl(dAtA, i, uint64(len(m.Payload)))
		i--
		dAtA[i] = 0x2a
	}
	if m.SessionId != 0 {
		i = encodeVarintGatewayTl(dAtA, i, uint64(m.SessionId))
		i--
		dAtA[i] = 0x20
	}
	if m.AuthKeyId != 0 {
		i = encodeVarintGatewayTl(dAtA, i, uint64(m.AuthKeyId))
		i--
		dAtA[i] = 0x18
	}
	if m.Constructor != 0 {
		i = encodeVarintGatewayTl(dAtA, i, uint64(m.Constructor))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintGatewayTl(dAtA []byte, offset int, v uint64) int {
	offset -= sovGatewayTl(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TLGatewaySendDataToGateway) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Constructor != 0 {
		n += 1 + sovGatewayTl(uint64(m.Constructor))
	}
	if m.AuthKeyId != 0 {
		n += 1 + sovGatewayTl(uint64(m.AuthKeyId))
	}
	if m.SessionId != 0 {
		n += 1 + sovGatewayTl(uint64(m.SessionId))
	}
	l = len(m.Payload)
	if l > 0 {
		n += 1 + l + sovGatewayTl(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovGatewayTl(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGatewayTl(x uint64) (n int) {
	return sovGatewayTl(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TLGatewaySendDataToGateway) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGatewayTl
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TL_gateway_sendDataToGateway: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TL_gateway_sendDataToGateway: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Constructor", wireType)
			}
			m.Constructor = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGatewayTl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Constructor |= TLConstructor(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuthKeyId", wireType)
			}
			m.AuthKeyId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGatewayTl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AuthKeyId |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SessionId", wireType)
			}
			m.SessionId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGatewayTl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SessionId |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Payload", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGatewayTl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthGatewayTl
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthGatewayTl
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Payload = append(m.Payload[:0], dAtA[iNdEx:postIndex]...)
			if m.Payload == nil {
				m.Payload = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGatewayTl(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGatewayTl
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipGatewayTl(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGatewayTl
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGatewayTl
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGatewayTl
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthGatewayTl
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGatewayTl
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGatewayTl
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGatewayTl        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGatewayTl          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGatewayTl = fmt.Errorf("proto: unexpected end of group")
)
