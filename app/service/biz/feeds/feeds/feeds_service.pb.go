// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: feeds_service.proto

package feeds

import (
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
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

type GetFeedListReq struct {
	UserId               int64    `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetFeedListReq) Reset()         { *m = GetFeedListReq{} }
func (m *GetFeedListReq) String() string { return proto.CompactTextString(m) }
func (*GetFeedListReq) ProtoMessage()    {}
func (*GetFeedListReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_31efcc759eeb2667, []int{0}
}
func (m *GetFeedListReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetFeedListReq.Unmarshal(m, b)
}
func (m *GetFeedListReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetFeedListReq.Marshal(b, m, deterministic)
}
func (m *GetFeedListReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetFeedListReq.Merge(m, src)
}
func (m *GetFeedListReq) XXX_Size() int {
	return xxx_messageInfo_GetFeedListReq.Size(m)
}
func (m *GetFeedListReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetFeedListReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetFeedListReq proto.InternalMessageInfo

func (m *GetFeedListReq) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

type UserChatDO struct {
	ChatId               int64    `protobuf:"varint,1,opt,name=chat_id,json=chatId,proto3" json:"chat_id,omitempty"`
	PhotoId              int64    `protobuf:"varint,2,opt,name=photo_id,json=photoId,proto3" json:"photo_id,omitempty"`
	Title                string   `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Status               bool     `protobuf:"varint,4,opt,name=status,proto3" json:"status,omitempty"`
	PeerType             int32    `protobuf:"varint,5,opt,name=peer_type,json=peerType,proto3" json:"peer_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserChatDO) Reset()         { *m = UserChatDO{} }
func (m *UserChatDO) String() string { return proto.CompactTextString(m) }
func (*UserChatDO) ProtoMessage()    {}
func (*UserChatDO) Descriptor() ([]byte, []int) {
	return fileDescriptor_31efcc759eeb2667, []int{1}
}
func (m *UserChatDO) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserChatDO.Unmarshal(m, b)
}
func (m *UserChatDO) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserChatDO.Marshal(b, m, deterministic)
}
func (m *UserChatDO) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserChatDO.Merge(m, src)
}
func (m *UserChatDO) XXX_Size() int {
	return xxx_messageInfo_UserChatDO.Size(m)
}
func (m *UserChatDO) XXX_DiscardUnknown() {
	xxx_messageInfo_UserChatDO.DiscardUnknown(m)
}

var xxx_messageInfo_UserChatDO proto.InternalMessageInfo

func (m *UserChatDO) GetChatId() int64 {
	if m != nil {
		return m.ChatId
	}
	return 0
}

func (m *UserChatDO) GetPhotoId() int64 {
	if m != nil {
		return m.PhotoId
	}
	return 0
}

func (m *UserChatDO) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *UserChatDO) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

func (m *UserChatDO) GetPeerType() int32 {
	if m != nil {
		return m.PeerType
	}
	return 0
}

type GetFeedListResp struct {
	Data                 []*UserChatDO `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *GetFeedListResp) Reset()         { *m = GetFeedListResp{} }
func (m *GetFeedListResp) String() string { return proto.CompactTextString(m) }
func (*GetFeedListResp) ProtoMessage()    {}
func (*GetFeedListResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_31efcc759eeb2667, []int{2}
}
func (m *GetFeedListResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetFeedListResp.Unmarshal(m, b)
}
func (m *GetFeedListResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetFeedListResp.Marshal(b, m, deterministic)
}
func (m *GetFeedListResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetFeedListResp.Merge(m, src)
}
func (m *GetFeedListResp) XXX_Size() int {
	return xxx_messageInfo_GetFeedListResp.Size(m)
}
func (m *GetFeedListResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GetFeedListResp.DiscardUnknown(m)
}

var xxx_messageInfo_GetFeedListResp proto.InternalMessageInfo

func (m *GetFeedListResp) GetData() []*UserChatDO {
	if m != nil {
		return m.Data
	}
	return nil
}

type FeedInsertItemDO struct {
	ChatId               int64    `protobuf:"varint,1,opt,name=chat_id,json=chatId,proto3" json:"chat_id,omitempty"`
	PeerType             int32    `protobuf:"varint,2,opt,name=peer_type,json=peerType,proto3" json:"peer_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FeedInsertItemDO) Reset()         { *m = FeedInsertItemDO{} }
func (m *FeedInsertItemDO) String() string { return proto.CompactTextString(m) }
func (*FeedInsertItemDO) ProtoMessage()    {}
func (*FeedInsertItemDO) Descriptor() ([]byte, []int) {
	return fileDescriptor_31efcc759eeb2667, []int{3}
}
func (m *FeedInsertItemDO) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FeedInsertItemDO.Unmarshal(m, b)
}
func (m *FeedInsertItemDO) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FeedInsertItemDO.Marshal(b, m, deterministic)
}
func (m *FeedInsertItemDO) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeedInsertItemDO.Merge(m, src)
}
func (m *FeedInsertItemDO) XXX_Size() int {
	return xxx_messageInfo_FeedInsertItemDO.Size(m)
}
func (m *FeedInsertItemDO) XXX_DiscardUnknown() {
	xxx_messageInfo_FeedInsertItemDO.DiscardUnknown(m)
}

var xxx_messageInfo_FeedInsertItemDO proto.InternalMessageInfo

func (m *FeedInsertItemDO) GetChatId() int64 {
	if m != nil {
		return m.ChatId
	}
	return 0
}

func (m *FeedInsertItemDO) GetPeerType() int32 {
	if m != nil {
		return m.PeerType
	}
	return 0
}

type UpdateFeedListReq struct {
	UserId               int64               `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Data                 []*FeedInsertItemDO `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *UpdateFeedListReq) Reset()         { *m = UpdateFeedListReq{} }
func (m *UpdateFeedListReq) String() string { return proto.CompactTextString(m) }
func (*UpdateFeedListReq) ProtoMessage()    {}
func (*UpdateFeedListReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_31efcc759eeb2667, []int{4}
}
func (m *UpdateFeedListReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateFeedListReq.Unmarshal(m, b)
}
func (m *UpdateFeedListReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateFeedListReq.Marshal(b, m, deterministic)
}
func (m *UpdateFeedListReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateFeedListReq.Merge(m, src)
}
func (m *UpdateFeedListReq) XXX_Size() int {
	return xxx_messageInfo_UpdateFeedListReq.Size(m)
}
func (m *UpdateFeedListReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateFeedListReq.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateFeedListReq proto.InternalMessageInfo

func (m *UpdateFeedListReq) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *UpdateFeedListReq) GetData() []*FeedInsertItemDO {
	if m != nil {
		return m.Data
	}
	return nil
}

type UpdateFeedListResp struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateFeedListResp) Reset()         { *m = UpdateFeedListResp{} }
func (m *UpdateFeedListResp) String() string { return proto.CompactTextString(m) }
func (*UpdateFeedListResp) ProtoMessage()    {}
func (*UpdateFeedListResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_31efcc759eeb2667, []int{5}
}
func (m *UpdateFeedListResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateFeedListResp.Unmarshal(m, b)
}
func (m *UpdateFeedListResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateFeedListResp.Marshal(b, m, deterministic)
}
func (m *UpdateFeedListResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateFeedListResp.Merge(m, src)
}
func (m *UpdateFeedListResp) XXX_Size() int {
	return xxx_messageInfo_UpdateFeedListResp.Size(m)
}
func (m *UpdateFeedListResp) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateFeedListResp.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateFeedListResp proto.InternalMessageInfo

type GetHistoryCounterReq struct {
	UserId               int64    `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetHistoryCounterReq) Reset()         { *m = GetHistoryCounterReq{} }
func (m *GetHistoryCounterReq) String() string { return proto.CompactTextString(m) }
func (*GetHistoryCounterReq) ProtoMessage()    {}
func (*GetHistoryCounterReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_31efcc759eeb2667, []int{6}
}
func (m *GetHistoryCounterReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetHistoryCounterReq.Unmarshal(m, b)
}
func (m *GetHistoryCounterReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetHistoryCounterReq.Marshal(b, m, deterministic)
}
func (m *GetHistoryCounterReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetHistoryCounterReq.Merge(m, src)
}
func (m *GetHistoryCounterReq) XXX_Size() int {
	return xxx_messageInfo_GetHistoryCounterReq.Size(m)
}
func (m *GetHistoryCounterReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetHistoryCounterReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetHistoryCounterReq proto.InternalMessageInfo

func (m *GetHistoryCounterReq) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

type GetHistoryCounterResp struct {
	Count                int32    `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetHistoryCounterResp) Reset()         { *m = GetHistoryCounterResp{} }
func (m *GetHistoryCounterResp) String() string { return proto.CompactTextString(m) }
func (*GetHistoryCounterResp) ProtoMessage()    {}
func (*GetHistoryCounterResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_31efcc759eeb2667, []int{7}
}
func (m *GetHistoryCounterResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetHistoryCounterResp.Unmarshal(m, b)
}
func (m *GetHistoryCounterResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetHistoryCounterResp.Marshal(b, m, deterministic)
}
func (m *GetHistoryCounterResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetHistoryCounterResp.Merge(m, src)
}
func (m *GetHistoryCounterResp) XXX_Size() int {
	return xxx_messageInfo_GetHistoryCounterResp.Size(m)
}
func (m *GetHistoryCounterResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GetHistoryCounterResp.DiscardUnknown(m)
}

var xxx_messageInfo_GetHistoryCounterResp proto.InternalMessageInfo

func (m *GetHistoryCounterResp) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

type GetFeedsOffsetListReq struct {
	UserId               int64    `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Limit                int32    `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Before               int32    `protobuf:"varint,3,opt,name=before,proto3" json:"before,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetFeedsOffsetListReq) Reset()         { *m = GetFeedsOffsetListReq{} }
func (m *GetFeedsOffsetListReq) String() string { return proto.CompactTextString(m) }
func (*GetFeedsOffsetListReq) ProtoMessage()    {}
func (*GetFeedsOffsetListReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_31efcc759eeb2667, []int{8}
}
func (m *GetFeedsOffsetListReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetFeedsOffsetListReq.Unmarshal(m, b)
}
func (m *GetFeedsOffsetListReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetFeedsOffsetListReq.Marshal(b, m, deterministic)
}
func (m *GetFeedsOffsetListReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetFeedsOffsetListReq.Merge(m, src)
}
func (m *GetFeedsOffsetListReq) XXX_Size() int {
	return xxx_messageInfo_GetFeedsOffsetListReq.Size(m)
}
func (m *GetFeedsOffsetListReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetFeedsOffsetListReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetFeedsOffsetListReq proto.InternalMessageInfo

func (m *GetFeedsOffsetListReq) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *GetFeedsOffsetListReq) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *GetFeedsOffsetListReq) GetBefore() int32 {
	if m != nil {
		return m.Before
	}
	return 0
}

type OffsetItemDo struct {
	OffsetMin            int32    `protobuf:"varint,1,opt,name=offset_min,json=offsetMin,proto3" json:"offset_min,omitempty"`
	OffsetMax            int32    `protobuf:"varint,2,opt,name=offset_max,json=offsetMax,proto3" json:"offset_max,omitempty"`
	PeerId               int64    `protobuf:"varint,3,opt,name=peer_id,json=peerId,proto3" json:"peer_id,omitempty"`
	Count                int32    `protobuf:"varint,4,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OffsetItemDo) Reset()         { *m = OffsetItemDo{} }
func (m *OffsetItemDo) String() string { return proto.CompactTextString(m) }
func (*OffsetItemDo) ProtoMessage()    {}
func (*OffsetItemDo) Descriptor() ([]byte, []int) {
	return fileDescriptor_31efcc759eeb2667, []int{9}
}
func (m *OffsetItemDo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OffsetItemDo.Unmarshal(m, b)
}
func (m *OffsetItemDo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OffsetItemDo.Marshal(b, m, deterministic)
}
func (m *OffsetItemDo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OffsetItemDo.Merge(m, src)
}
func (m *OffsetItemDo) XXX_Size() int {
	return xxx_messageInfo_OffsetItemDo.Size(m)
}
func (m *OffsetItemDo) XXX_DiscardUnknown() {
	xxx_messageInfo_OffsetItemDo.DiscardUnknown(m)
}

var xxx_messageInfo_OffsetItemDo proto.InternalMessageInfo

func (m *OffsetItemDo) GetOffsetMin() int32 {
	if m != nil {
		return m.OffsetMin
	}
	return 0
}

func (m *OffsetItemDo) GetOffsetMax() int32 {
	if m != nil {
		return m.OffsetMax
	}
	return 0
}

func (m *OffsetItemDo) GetPeerId() int64 {
	if m != nil {
		return m.PeerId
	}
	return 0
}

func (m *OffsetItemDo) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

type GetFeedsOffsetListResp struct {
	Data                 []*OffsetItemDo `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *GetFeedsOffsetListResp) Reset()         { *m = GetFeedsOffsetListResp{} }
func (m *GetFeedsOffsetListResp) String() string { return proto.CompactTextString(m) }
func (*GetFeedsOffsetListResp) ProtoMessage()    {}
func (*GetFeedsOffsetListResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_31efcc759eeb2667, []int{10}
}
func (m *GetFeedsOffsetListResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetFeedsOffsetListResp.Unmarshal(m, b)
}
func (m *GetFeedsOffsetListResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetFeedsOffsetListResp.Marshal(b, m, deterministic)
}
func (m *GetFeedsOffsetListResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetFeedsOffsetListResp.Merge(m, src)
}
func (m *GetFeedsOffsetListResp) XXX_Size() int {
	return xxx_messageInfo_GetFeedsOffsetListResp.Size(m)
}
func (m *GetFeedsOffsetListResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GetFeedsOffsetListResp.DiscardUnknown(m)
}

var xxx_messageInfo_GetFeedsOffsetListResp proto.InternalMessageInfo

func (m *GetFeedsOffsetListResp) GetData() []*OffsetItemDo {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*GetFeedListReq)(nil), "mtproto.GetFeedListReq")
	proto.RegisterType((*UserChatDO)(nil), "mtproto.UserChatDO")
	proto.RegisterType((*GetFeedListResp)(nil), "mtproto.GetFeedListResp")
	proto.RegisterType((*FeedInsertItemDO)(nil), "mtproto.FeedInsertItemDO")
	proto.RegisterType((*UpdateFeedListReq)(nil), "mtproto.UpdateFeedListReq")
	proto.RegisterType((*UpdateFeedListResp)(nil), "mtproto.UpdateFeedListResp")
	proto.RegisterType((*GetHistoryCounterReq)(nil), "mtproto.GetHistoryCounterReq")
	proto.RegisterType((*GetHistoryCounterResp)(nil), "mtproto.GetHistoryCounterResp")
	proto.RegisterType((*GetFeedsOffsetListReq)(nil), "mtproto.GetFeedsOffsetListReq")
	proto.RegisterType((*OffsetItemDo)(nil), "mtproto.OffsetItemDo")
	proto.RegisterType((*GetFeedsOffsetListResp)(nil), "mtproto.GetFeedsOffsetListResp")
}

func init() { proto.RegisterFile("feeds_service.proto", fileDescriptor_31efcc759eeb2667) }

var fileDescriptor_31efcc759eeb2667 = []byte{
	// 538 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0x4d, 0x6f, 0xda, 0x40,
	0x10, 0x95, 0x21, 0xe6, 0x63, 0xfa, 0x99, 0x09, 0x09, 0x0e, 0x51, 0x52, 0xe4, 0x4b, 0xc9, 0x21,
	0x54, 0x4a, 0x6f, 0x3d, 0x86, 0x2a, 0x01, 0xa9, 0x15, 0xc8, 0x6d, 0xa4, 0x2a, 0x91, 0x8a, 0x1c,
	0x3c, 0x94, 0x95, 0x02, 0xde, 0x7a, 0x87, 0x2a, 0xa8, 0x3f, 0xa1, 0xff, 0xb0, 0xbf, 0xa6, 0xda,
	0x5d, 0x17, 0x6c, 0x4a, 0x69, 0x6f, 0x7e, 0x3b, 0x6f, 0x67, 0xdf, 0xdb, 0x37, 0x6b, 0xd8, 0x1b,
	0x13, 0x45, 0x6a, 0xa8, 0x28, 0xf9, 0x26, 0x46, 0xd4, 0x96, 0x49, 0xcc, 0x31, 0x96, 0xa7, 0x6c,
	0x3e, 0xfc, 0x53, 0x78, 0x7a, 0x45, 0x7c, 0x49, 0x14, 0xbd, 0x13, 0x8a, 0x03, 0xfa, 0x8a, 0x75,
	0x28, 0xcf, 0x15, 0x25, 0x43, 0x11, 0x79, 0x4e, 0xd3, 0x69, 0x15, 0x83, 0x92, 0x86, 0xbd, 0xc8,
	0xff, 0xe1, 0x00, 0x5c, 0x2b, 0x4a, 0x3a, 0x93, 0x90, 0xdf, 0xf6, 0x35, 0x6f, 0x34, 0x09, 0x39,
	0xc3, 0xd3, 0xb0, 0x17, 0xe1, 0x21, 0x54, 0xe4, 0x24, 0xe6, 0x58, 0x57, 0x0a, 0xa6, 0x52, 0x36,
	0xb8, 0x17, 0x61, 0x0d, 0x5c, 0x16, 0x7c, 0x4f, 0x5e, 0xb1, 0xe9, 0xb4, 0xaa, 0x81, 0x05, 0x78,
	0x00, 0x25, 0xc5, 0x21, 0xcf, 0x95, 0xb7, 0xd3, 0x74, 0x5a, 0x95, 0x20, 0x45, 0x78, 0x04, 0x55,
	0x49, 0x94, 0x0c, 0x79, 0x21, 0xc9, 0x73, 0x9b, 0x4e, 0xcb, 0x0d, 0x2a, 0x7a, 0xe1, 0xe3, 0x42,
	0x92, 0xff, 0x06, 0x9e, 0xe5, 0x84, 0x2b, 0x89, 0x2f, 0x61, 0x27, 0x0a, 0x39, 0xf4, 0x9c, 0x66,
	0xb1, 0xf5, 0xe8, 0x7c, 0xaf, 0x9d, 0x7a, 0x6c, 0xaf, 0x44, 0x07, 0x86, 0xe0, 0x77, 0xe1, 0xb9,
	0xde, 0xd8, 0x9b, 0x29, 0x4a, 0xb8, 0xc7, 0x34, 0xdd, 0x66, 0x27, 0xa7, 0xa2, 0xb0, 0xa6, 0xe2,
	0x16, 0x76, 0xaf, 0x65, 0x14, 0x32, 0xfd, 0xcf, 0x0d, 0xe2, 0x59, 0x2a, 0xb0, 0x60, 0x04, 0x1e,
	0x2e, 0x05, 0xae, 0x8b, 0x49, 0x65, 0xd6, 0x00, 0xd7, 0x9b, 0x2b, 0xe9, 0xbf, 0x82, 0xda, 0x15,
	0x71, 0x57, 0x28, 0x8e, 0x93, 0x45, 0x27, 0x9e, 0xcf, 0x98, 0x92, 0xad, 0xb9, 0x9d, 0xc1, 0xfe,
	0x86, 0x0d, 0x4a, 0xea, 0x34, 0x46, 0x1a, 0x1a, 0xbe, 0x1b, 0x58, 0xe0, 0x7f, 0x36, 0x74, 0x7d,
	0xa4, 0xea, 0x8f, 0xc7, 0x8a, 0xf8, 0x9f, 0xb6, 0x6a, 0xe0, 0xde, 0x8b, 0xa9, 0xe0, 0xf4, 0x76,
	0x2c, 0xd0, 0xa9, 0xde, 0xd1, 0x38, 0x4e, 0x6c, 0xd8, 0x6e, 0x90, 0x22, 0xff, 0x3b, 0x3c, 0xb6,
	0x7d, 0x8d, 0xd7, 0x18, 0x8f, 0x01, 0x62, 0x83, 0x87, 0x53, 0x31, 0x4b, 0xa5, 0x54, 0xed, 0xca,
	0x7b, 0x31, 0xcb, 0x96, 0xc3, 0x87, 0xf4, 0x84, 0xdf, 0xe5, 0xf0, 0x41, 0x8b, 0x32, 0xe9, 0x88,
	0xc8, 0x1c, 0x53, 0x0c, 0x4a, 0x1a, 0x5a, 0x51, 0xd6, 0xdc, 0x4e, 0xd6, 0x5c, 0x07, 0x0e, 0x36,
	0x99, 0x53, 0x12, 0x4f, 0x73, 0xc3, 0xb3, 0xbf, 0xcc, 0x26, 0xab, 0xd5, 0xe6, 0x72, 0xfe, 0xb3,
	0x00, 0x95, 0x60, 0xd0, 0x31, 0x5d, 0xf0, 0x12, 0x76, 0xed, 0x03, 0xfb, 0xb2, 0x9a, 0x46, 0xac,
	0x2f, 0xb7, 0xe7, 0x1f, 0x57, 0xc3, 0xdb, 0x5c, 0x50, 0x12, 0xfb, 0x50, 0xb3, 0x7d, 0xe6, 0xb9,
	0xc8, 0xb1, 0xb1, 0x1a, 0xe3, 0xf5, 0x41, 0x6b, 0x1c, 0xfd, 0xb5, 0xa6, 0x24, 0x7e, 0x82, 0xfa,
	0x52, 0x58, 0x3e, 0x7c, 0x3c, 0xce, 0xaa, 0xf8, 0x63, 0x92, 0x1a, 0x27, 0xdb, 0xca, 0x4a, 0xe2,
	0x2d, 0x78, 0x39, 0xcb, 0x99, 0xab, 0xc4, 0x93, 0x75, 0x83, 0xf9, 0x21, 0x6a, 0xbc, 0xd8, 0x5a,
	0x57, 0xf2, 0xa2, 0x0d, 0x4f, 0x72, 0x3f, 0xac, 0x0b, 0xbc, 0x19, 0x68, 0xbe, 0xe1, 0x7e, 0xb0,
	0x6b, 0xdd, 0xc2, 0xc0, 0xb9, 0x71, 0x0d, 0xed, 0xae, 0x64, 0xba, 0xbd, 0xfe, 0x15, 0x00, 0x00,
	0xff, 0xff, 0xc3, 0x19, 0x49, 0xc0, 0xe7, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RPCFeedsClient is the client API for RPCFeeds service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RPCFeedsClient interface {
	FeedsGetFeedList(ctx context.Context, in *GetFeedListReq, opts ...grpc.CallOption) (*GetFeedListResp, error)
	FeedsUpdateFeedList(ctx context.Context, in *UpdateFeedListReq, opts ...grpc.CallOption) (*UpdateFeedListResp, error)
	FeedsGetHistoryCounter(ctx context.Context, in *GetHistoryCounterReq, opts ...grpc.CallOption) (*GetHistoryCounterResp, error)
	FeedsGetFeedsOffsetList(ctx context.Context, in *GetFeedsOffsetListReq, opts ...grpc.CallOption) (*GetFeedsOffsetListResp, error)
}

type rPCFeedsClient struct {
	cc *grpc.ClientConn
}

func NewRPCFeedsClient(cc *grpc.ClientConn) RPCFeedsClient {
	return &rPCFeedsClient{cc}
}

func (c *rPCFeedsClient) FeedsGetFeedList(ctx context.Context, in *GetFeedListReq, opts ...grpc.CallOption) (*GetFeedListResp, error) {
	out := new(GetFeedListResp)
	err := c.cc.Invoke(ctx, "/mtproto.RPCFeeds/feeds_getFeedList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCFeedsClient) FeedsUpdateFeedList(ctx context.Context, in *UpdateFeedListReq, opts ...grpc.CallOption) (*UpdateFeedListResp, error) {
	out := new(UpdateFeedListResp)
	err := c.cc.Invoke(ctx, "/mtproto.RPCFeeds/feeds_updateFeedList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCFeedsClient) FeedsGetHistoryCounter(ctx context.Context, in *GetHistoryCounterReq, opts ...grpc.CallOption) (*GetHistoryCounterResp, error) {
	out := new(GetHistoryCounterResp)
	err := c.cc.Invoke(ctx, "/mtproto.RPCFeeds/feeds_getHistoryCounter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCFeedsClient) FeedsGetFeedsOffsetList(ctx context.Context, in *GetFeedsOffsetListReq, opts ...grpc.CallOption) (*GetFeedsOffsetListResp, error) {
	out := new(GetFeedsOffsetListResp)
	err := c.cc.Invoke(ctx, "/mtproto.RPCFeeds/feeds_getFeedsOffsetList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RPCFeedsServer is the server API for RPCFeeds service.
type RPCFeedsServer interface {
	FeedsGetFeedList(context.Context, *GetFeedListReq) (*GetFeedListResp, error)
	FeedsUpdateFeedList(context.Context, *UpdateFeedListReq) (*UpdateFeedListResp, error)
	FeedsGetHistoryCounter(context.Context, *GetHistoryCounterReq) (*GetHistoryCounterResp, error)
	FeedsGetFeedsOffsetList(context.Context, *GetFeedsOffsetListReq) (*GetFeedsOffsetListResp, error)
}

// UnimplementedRPCFeedsServer can be embedded to have forward compatible implementations.
type UnimplementedRPCFeedsServer struct {
}

func (*UnimplementedRPCFeedsServer) FeedsGetFeedList(ctx context.Context, req *GetFeedListReq) (*GetFeedListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FeedsGetFeedList not implemented")
}
func (*UnimplementedRPCFeedsServer) FeedsUpdateFeedList(ctx context.Context, req *UpdateFeedListReq) (*UpdateFeedListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FeedsUpdateFeedList not implemented")
}
func (*UnimplementedRPCFeedsServer) FeedsGetHistoryCounter(ctx context.Context, req *GetHistoryCounterReq) (*GetHistoryCounterResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FeedsGetHistoryCounter not implemented")
}
func (*UnimplementedRPCFeedsServer) FeedsGetFeedsOffsetList(ctx context.Context, req *GetFeedsOffsetListReq) (*GetFeedsOffsetListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FeedsGetFeedsOffsetList not implemented")
}

func RegisterRPCFeedsServer(s *grpc.Server, srv RPCFeedsServer) {
	s.RegisterService(&_RPCFeeds_serviceDesc, srv)
}

func _RPCFeeds_FeedsGetFeedList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFeedListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCFeedsServer).FeedsGetFeedList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mtproto.RPCFeeds/FeedsGetFeedList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCFeedsServer).FeedsGetFeedList(ctx, req.(*GetFeedListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCFeeds_FeedsUpdateFeedList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateFeedListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCFeedsServer).FeedsUpdateFeedList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mtproto.RPCFeeds/FeedsUpdateFeedList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCFeedsServer).FeedsUpdateFeedList(ctx, req.(*UpdateFeedListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCFeeds_FeedsGetHistoryCounter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHistoryCounterReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCFeedsServer).FeedsGetHistoryCounter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mtproto.RPCFeeds/FeedsGetHistoryCounter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCFeedsServer).FeedsGetHistoryCounter(ctx, req.(*GetHistoryCounterReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCFeeds_FeedsGetFeedsOffsetList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFeedsOffsetListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCFeedsServer).FeedsGetFeedsOffsetList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mtproto.RPCFeeds/FeedsGetFeedsOffsetList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCFeedsServer).FeedsGetFeedsOffsetList(ctx, req.(*GetFeedsOffsetListReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _RPCFeeds_serviceDesc = grpc.ServiceDesc{
	ServiceName: "mtproto.RPCFeeds",
	HandlerType: (*RPCFeedsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "feeds_getFeedList",
			Handler:    _RPCFeeds_FeedsGetFeedList_Handler,
		},
		{
			MethodName: "feeds_updateFeedList",
			Handler:    _RPCFeeds_FeedsUpdateFeedList_Handler,
		},
		{
			MethodName: "feeds_getHistoryCounter",
			Handler:    _RPCFeeds_FeedsGetHistoryCounter_Handler,
		},
		{
			MethodName: "feeds_getFeedsOffsetList",
			Handler:    _RPCFeeds_FeedsGetFeedsOffsetList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "feeds_service.proto",
}