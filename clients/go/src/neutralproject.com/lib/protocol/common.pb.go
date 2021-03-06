// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common.proto

package neutralservice

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Common Definitions
type UserID struct {
	SourcePubkeyHash     *PubkeyHash `protobuf:"bytes,1,opt,name=sourcePubkeyHash,proto3" json:"sourcePubkeyHash,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *UserID) Reset()         { *m = UserID{} }
func (m *UserID) String() string { return proto.CompactTextString(m) }
func (*UserID) ProtoMessage()    {}
func (*UserID) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{0}
}

func (m *UserID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserID.Unmarshal(m, b)
}
func (m *UserID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserID.Marshal(b, m, deterministic)
}
func (m *UserID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserID.Merge(m, src)
}
func (m *UserID) XXX_Size() int {
	return xxx_messageInfo_UserID.Size(m)
}
func (m *UserID) XXX_DiscardUnknown() {
	xxx_messageInfo_UserID.DiscardUnknown(m)
}

var xxx_messageInfo_UserID proto.InternalMessageInfo

func (m *UserID) GetSourcePubkeyHash() *PubkeyHash {
	if m != nil {
		return m.SourcePubkeyHash
	}
	return nil
}

type SessionToken struct {
	Token                string        `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Duration             *UtcMicroTime `protobuf:"bytes,2,opt,name=duration,proto3" json:"duration,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *SessionToken) Reset()         { *m = SessionToken{} }
func (m *SessionToken) String() string { return proto.CompactTextString(m) }
func (*SessionToken) ProtoMessage()    {}
func (*SessionToken) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{1}
}

func (m *SessionToken) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SessionToken.Unmarshal(m, b)
}
func (m *SessionToken) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SessionToken.Marshal(b, m, deterministic)
}
func (m *SessionToken) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SessionToken.Merge(m, src)
}
func (m *SessionToken) XXX_Size() int {
	return xxx_messageInfo_SessionToken.Size(m)
}
func (m *SessionToken) XXX_DiscardUnknown() {
	xxx_messageInfo_SessionToken.DiscardUnknown(m)
}

var xxx_messageInfo_SessionToken proto.InternalMessageInfo

func (m *SessionToken) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *SessionToken) GetDuration() *UtcMicroTime {
	if m != nil {
		return m.Duration
	}
	return nil
}

type SystemStatus struct {
	Status               uint32   `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SystemStatus) Reset()         { *m = SystemStatus{} }
func (m *SystemStatus) String() string { return proto.CompactTextString(m) }
func (*SystemStatus) ProtoMessage()    {}
func (*SystemStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{2}
}

func (m *SystemStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SystemStatus.Unmarshal(m, b)
}
func (m *SystemStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SystemStatus.Marshal(b, m, deterministic)
}
func (m *SystemStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SystemStatus.Merge(m, src)
}
func (m *SystemStatus) XXX_Size() int {
	return xxx_messageInfo_SystemStatus.Size(m)
}
func (m *SystemStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_SystemStatus.DiscardUnknown(m)
}

var xxx_messageInfo_SystemStatus proto.InternalMessageInfo

func (m *SystemStatus) GetStatus() uint32 {
	if m != nil {
		return m.Status
	}
	return 0
}

type UtcMicroTime struct {
	UtcMicroTime         uint64   `protobuf:"varint,1,opt,name=utcMicroTime,proto3" json:"utcMicroTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UtcMicroTime) Reset()         { *m = UtcMicroTime{} }
func (m *UtcMicroTime) String() string { return proto.CompactTextString(m) }
func (*UtcMicroTime) ProtoMessage()    {}
func (*UtcMicroTime) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{3}
}

func (m *UtcMicroTime) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UtcMicroTime.Unmarshal(m, b)
}
func (m *UtcMicroTime) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UtcMicroTime.Marshal(b, m, deterministic)
}
func (m *UtcMicroTime) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UtcMicroTime.Merge(m, src)
}
func (m *UtcMicroTime) XXX_Size() int {
	return xxx_messageInfo_UtcMicroTime.Size(m)
}
func (m *UtcMicroTime) XXX_DiscardUnknown() {
	xxx_messageInfo_UtcMicroTime.DiscardUnknown(m)
}

var xxx_messageInfo_UtcMicroTime proto.InternalMessageInfo

func (m *UtcMicroTime) GetUtcMicroTime() uint64 {
	if m != nil {
		return m.UtcMicroTime
	}
	return 0
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{4}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type Signature struct {
	Method               string   `protobuf:"bytes,1,opt,name=method,proto3" json:"method,omitempty"`
	Signature            []byte   `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
	PubKeyFingerprint    []byte   `protobuf:"bytes,3,opt,name=pubKeyFingerprint,proto3" json:"pubKeyFingerprint,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Signature) Reset()         { *m = Signature{} }
func (m *Signature) String() string { return proto.CompactTextString(m) }
func (*Signature) ProtoMessage()    {}
func (*Signature) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{5}
}

func (m *Signature) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Signature.Unmarshal(m, b)
}
func (m *Signature) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Signature.Marshal(b, m, deterministic)
}
func (m *Signature) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Signature.Merge(m, src)
}
func (m *Signature) XXX_Size() int {
	return xxx_messageInfo_Signature.Size(m)
}
func (m *Signature) XXX_DiscardUnknown() {
	xxx_messageInfo_Signature.DiscardUnknown(m)
}

var xxx_messageInfo_Signature proto.InternalMessageInfo

func (m *Signature) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *Signature) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *Signature) GetPubKeyFingerprint() []byte {
	if m != nil {
		return m.PubKeyFingerprint
	}
	return nil
}

type PubkeyHash struct {
	Hash                 []byte   `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PubkeyHash) Reset()         { *m = PubkeyHash{} }
func (m *PubkeyHash) String() string { return proto.CompactTextString(m) }
func (*PubkeyHash) ProtoMessage()    {}
func (*PubkeyHash) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{6}
}

func (m *PubkeyHash) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PubkeyHash.Unmarshal(m, b)
}
func (m *PubkeyHash) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PubkeyHash.Marshal(b, m, deterministic)
}
func (m *PubkeyHash) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PubkeyHash.Merge(m, src)
}
func (m *PubkeyHash) XXX_Size() int {
	return xxx_messageInfo_PubkeyHash.Size(m)
}
func (m *PubkeyHash) XXX_DiscardUnknown() {
	xxx_messageInfo_PubkeyHash.DiscardUnknown(m)
}

var xxx_messageInfo_PubkeyHash proto.InternalMessageInfo

func (m *PubkeyHash) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

type Pubkey struct {
	Pubkey               []byte   `protobuf:"bytes,1,opt,name=pubkey,proto3" json:"pubkey,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Pubkey) Reset()         { *m = Pubkey{} }
func (m *Pubkey) String() string { return proto.CompactTextString(m) }
func (*Pubkey) ProtoMessage()    {}
func (*Pubkey) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{7}
}

func (m *Pubkey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pubkey.Unmarshal(m, b)
}
func (m *Pubkey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pubkey.Marshal(b, m, deterministic)
}
func (m *Pubkey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pubkey.Merge(m, src)
}
func (m *Pubkey) XXX_Size() int {
	return xxx_messageInfo_Pubkey.Size(m)
}
func (m *Pubkey) XXX_DiscardUnknown() {
	xxx_messageInfo_Pubkey.DiscardUnknown(m)
}

var xxx_messageInfo_Pubkey proto.InternalMessageInfo

func (m *Pubkey) GetPubkey() []byte {
	if m != nil {
		return m.Pubkey
	}
	return nil
}

type Price struct {
	Value                float64  `protobuf:"fixed64,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Price) Reset()         { *m = Price{} }
func (m *Price) String() string { return proto.CompactTextString(m) }
func (*Price) ProtoMessage()    {}
func (*Price) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{8}
}

func (m *Price) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Price.Unmarshal(m, b)
}
func (m *Price) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Price.Marshal(b, m, deterministic)
}
func (m *Price) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Price.Merge(m, src)
}
func (m *Price) XXX_Size() int {
	return xxx_messageInfo_Price.Size(m)
}
func (m *Price) XXX_DiscardUnknown() {
	xxx_messageInfo_Price.DiscardUnknown(m)
}

var xxx_messageInfo_Price proto.InternalMessageInfo

func (m *Price) GetValue() float64 {
	if m != nil {
		return m.Value
	}
	return 0
}

type Address struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	AddressType          string   `protobuf:"bytes,2,opt,name=addressType,proto3" json:"addressType,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Address) Reset()         { *m = Address{} }
func (m *Address) String() string { return proto.CompactTextString(m) }
func (*Address) ProtoMessage()    {}
func (*Address) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{9}
}

func (m *Address) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Address.Unmarshal(m, b)
}
func (m *Address) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Address.Marshal(b, m, deterministic)
}
func (m *Address) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Address.Merge(m, src)
}
func (m *Address) XXX_Size() int {
	return xxx_messageInfo_Address.Size(m)
}
func (m *Address) XXX_DiscardUnknown() {
	xxx_messageInfo_Address.DiscardUnknown(m)
}

var xxx_messageInfo_Address proto.InternalMessageInfo

func (m *Address) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *Address) GetAddressType() string {
	if m != nil {
		return m.AddressType
	}
	return ""
}

type InstrumentID struct {
	Symbol               string   `protobuf:"bytes,1,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Address              *Address `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InstrumentID) Reset()         { *m = InstrumentID{} }
func (m *InstrumentID) String() string { return proto.CompactTextString(m) }
func (*InstrumentID) ProtoMessage()    {}
func (*InstrumentID) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{10}
}

func (m *InstrumentID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InstrumentID.Unmarshal(m, b)
}
func (m *InstrumentID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InstrumentID.Marshal(b, m, deterministic)
}
func (m *InstrumentID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InstrumentID.Merge(m, src)
}
func (m *InstrumentID) XXX_Size() int {
	return xxx_messageInfo_InstrumentID.Size(m)
}
func (m *InstrumentID) XXX_DiscardUnknown() {
	xxx_messageInfo_InstrumentID.DiscardUnknown(m)
}

var xxx_messageInfo_InstrumentID proto.InternalMessageInfo

func (m *InstrumentID) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *InstrumentID) GetAddress() *Address {
	if m != nil {
		return m.Address
	}
	return nil
}

type HeartbeatRequest struct {
	Time                 *UtcMicroTime `protobuf:"bytes,1,opt,name=time,proto3" json:"time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *HeartbeatRequest) Reset()         { *m = HeartbeatRequest{} }
func (m *HeartbeatRequest) String() string { return proto.CompactTextString(m) }
func (*HeartbeatRequest) ProtoMessage()    {}
func (*HeartbeatRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{11}
}

func (m *HeartbeatRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HeartbeatRequest.Unmarshal(m, b)
}
func (m *HeartbeatRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HeartbeatRequest.Marshal(b, m, deterministic)
}
func (m *HeartbeatRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeartbeatRequest.Merge(m, src)
}
func (m *HeartbeatRequest) XXX_Size() int {
	return xxx_messageInfo_HeartbeatRequest.Size(m)
}
func (m *HeartbeatRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HeartbeatRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HeartbeatRequest proto.InternalMessageInfo

func (m *HeartbeatRequest) GetTime() *UtcMicroTime {
	if m != nil {
		return m.Time
	}
	return nil
}

type Heartbeat struct {
	Time                 *UtcMicroTime `protobuf:"bytes,1,opt,name=time,proto3" json:"time,omitempty"`
	Status               *SystemStatus `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Heartbeat) Reset()         { *m = Heartbeat{} }
func (m *Heartbeat) String() string { return proto.CompactTextString(m) }
func (*Heartbeat) ProtoMessage()    {}
func (*Heartbeat) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{12}
}

func (m *Heartbeat) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Heartbeat.Unmarshal(m, b)
}
func (m *Heartbeat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Heartbeat.Marshal(b, m, deterministic)
}
func (m *Heartbeat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Heartbeat.Merge(m, src)
}
func (m *Heartbeat) XXX_Size() int {
	return xxx_messageInfo_Heartbeat.Size(m)
}
func (m *Heartbeat) XXX_DiscardUnknown() {
	xxx_messageInfo_Heartbeat.DiscardUnknown(m)
}

var xxx_messageInfo_Heartbeat proto.InternalMessageInfo

func (m *Heartbeat) GetTime() *UtcMicroTime {
	if m != nil {
		return m.Time
	}
	return nil
}

func (m *Heartbeat) GetStatus() *SystemStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

type CanonicalResponse struct {
	Canonical            string   `protobuf:"bytes,1,opt,name=canonical,proto3" json:"canonical,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CanonicalResponse) Reset()         { *m = CanonicalResponse{} }
func (m *CanonicalResponse) String() string { return proto.CompactTextString(m) }
func (*CanonicalResponse) ProtoMessage()    {}
func (*CanonicalResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{13}
}

func (m *CanonicalResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CanonicalResponse.Unmarshal(m, b)
}
func (m *CanonicalResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CanonicalResponse.Marshal(b, m, deterministic)
}
func (m *CanonicalResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CanonicalResponse.Merge(m, src)
}
func (m *CanonicalResponse) XXX_Size() int {
	return xxx_messageInfo_CanonicalResponse.Size(m)
}
func (m *CanonicalResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CanonicalResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CanonicalResponse proto.InternalMessageInfo

func (m *CanonicalResponse) GetCanonical() string {
	if m != nil {
		return m.Canonical
	}
	return ""
}

type Anything struct {
	Message              *any.Any `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Anything) Reset()         { *m = Anything{} }
func (m *Anything) String() string { return proto.CompactTextString(m) }
func (*Anything) ProtoMessage()    {}
func (*Anything) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{14}
}

func (m *Anything) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Anything.Unmarshal(m, b)
}
func (m *Anything) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Anything.Marshal(b, m, deterministic)
}
func (m *Anything) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Anything.Merge(m, src)
}
func (m *Anything) XXX_Size() int {
	return xxx_messageInfo_Anything.Size(m)
}
func (m *Anything) XXX_DiscardUnknown() {
	xxx_messageInfo_Anything.DiscardUnknown(m)
}

var xxx_messageInfo_Anything proto.InternalMessageInfo

func (m *Anything) GetMessage() *any.Any {
	if m != nil {
		return m.Message
	}
	return nil
}

func init() {
	proto.RegisterType((*UserID)(nil), "neutralservice.UserID")
	proto.RegisterType((*SessionToken)(nil), "neutralservice.SessionToken")
	proto.RegisterType((*SystemStatus)(nil), "neutralservice.SystemStatus")
	proto.RegisterType((*UtcMicroTime)(nil), "neutralservice.UtcMicroTime")
	proto.RegisterType((*Empty)(nil), "neutralservice.Empty")
	proto.RegisterType((*Signature)(nil), "neutralservice.Signature")
	proto.RegisterType((*PubkeyHash)(nil), "neutralservice.PubkeyHash")
	proto.RegisterType((*Pubkey)(nil), "neutralservice.Pubkey")
	proto.RegisterType((*Price)(nil), "neutralservice.Price")
	proto.RegisterType((*Address)(nil), "neutralservice.Address")
	proto.RegisterType((*InstrumentID)(nil), "neutralservice.InstrumentID")
	proto.RegisterType((*HeartbeatRequest)(nil), "neutralservice.HeartbeatRequest")
	proto.RegisterType((*Heartbeat)(nil), "neutralservice.Heartbeat")
	proto.RegisterType((*CanonicalResponse)(nil), "neutralservice.CanonicalResponse")
	proto.RegisterType((*Anything)(nil), "neutralservice.Anything")
}

func init() { proto.RegisterFile("common.proto", fileDescriptor_555bd8c177793206) }

var fileDescriptor_555bd8c177793206 = []byte{
	// 529 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0x5b, 0x6b, 0x13, 0x41,
	0x14, 0xc7, 0x89, 0xb6, 0x49, 0xf7, 0x74, 0x95, 0x76, 0x28, 0x1a, 0x4b, 0x85, 0x30, 0x82, 0xf8,
	0x20, 0x1b, 0x5b, 0x7d, 0x10, 0xdf, 0x52, 0x63, 0x69, 0x10, 0x21, 0x4c, 0xd2, 0x07, 0x5f, 0x84,
	0xd9, 0xcd, 0x71, 0x33, 0x76, 0x77, 0x66, 0x9d, 0x4b, 0x61, 0xbf, 0xbd, 0x64, 0x76, 0x36, 0x97,
	0x16, 0x11, 0xdf, 0xce, 0xff, 0xcc, 0x6f, 0xce, 0x6d, 0xce, 0x2e, 0xc4, 0x99, 0x2a, 0x4b, 0x25,
	0x93, 0x4a, 0x2b, 0xab, 0xc8, 0x53, 0x89, 0xce, 0x6a, 0x5e, 0x18, 0xd4, 0x77, 0x22, 0xc3, 0xd3,
	0x17, 0xb9, 0x52, 0x79, 0x81, 0x43, 0x7f, 0x9a, 0xba, 0x9f, 0x43, 0x2e, 0xeb, 0x06, 0xa5, 0x53,
	0xe8, 0xde, 0x18, 0xd4, 0x93, 0x31, 0xb9, 0x82, 0x23, 0xa3, 0x9c, 0xce, 0x70, 0xea, 0xd2, 0x5b,
	0xac, 0xaf, 0xb9, 0x59, 0xf6, 0x3b, 0x83, 0xce, 0x9b, 0xc3, 0x8b, 0xd3, 0x64, 0x37, 0x5e, 0xb2,
	0x21, 0xd8, 0x83, 0x3b, 0xf4, 0x07, 0xc4, 0x33, 0x34, 0x46, 0x28, 0x39, 0x57, 0xb7, 0x28, 0xc9,
	0x09, 0xec, 0xdb, 0x95, 0xe1, 0x83, 0x45, 0xac, 0x11, 0xe4, 0x23, 0x1c, 0x2c, 0x9c, 0xe6, 0x56,
	0x28, 0xd9, 0x7f, 0xe4, 0xb3, 0x9c, 0xdd, 0xcf, 0x72, 0x63, 0xb3, 0x6f, 0x22, 0xd3, 0x6a, 0x2e,
	0x4a, 0x64, 0x6b, 0x9a, 0xbe, 0x86, 0x78, 0x56, 0x1b, 0x8b, 0xe5, 0xcc, 0x72, 0xeb, 0x0c, 0x79,
	0x06, 0x5d, 0xe3, 0x2d, 0x9f, 0xe0, 0x09, 0x0b, 0x8a, 0x5e, 0x40, 0xbc, 0x1d, 0x81, 0x50, 0x88,
	0xdd, 0x96, 0xf6, 0xf4, 0x1e, 0xdb, 0xf1, 0xd1, 0x1e, 0xec, 0x7f, 0x29, 0x2b, 0x5b, 0x53, 0x05,
	0xd1, 0x4c, 0xe4, 0x92, 0x5b, 0xa7, 0x71, 0x95, 0xa1, 0x44, 0xbb, 0x54, 0x8b, 0xd0, 0x42, 0x50,
	0xe4, 0x0c, 0x22, 0xd3, 0x42, 0xbe, 0x89, 0x98, 0x6d, 0x1c, 0xe4, 0x2d, 0x1c, 0x57, 0x2e, 0xfd,
	0x8a, 0xf5, 0x95, 0x90, 0x39, 0xea, 0x4a, 0x0b, 0x69, 0xfb, 0x8f, 0x3d, 0xf5, 0xf0, 0x80, 0x0e,
	0x00, 0x36, 0x33, 0x24, 0x04, 0xf6, 0x96, 0xed, 0xfc, 0x63, 0xe6, 0x6d, 0x3a, 0x80, 0x6e, 0x43,
	0xac, 0xea, 0xa9, 0xbc, 0x15, 0xce, 0x83, 0xa2, 0x2f, 0x61, 0x7f, 0xaa, 0x45, 0x86, 0xab, 0x91,
	0xdf, 0xf1, 0xc2, 0x35, 0x3d, 0x76, 0x58, 0x23, 0xe8, 0x08, 0x7a, 0xa3, 0xc5, 0x42, 0xa3, 0x31,
	0xbb, 0x40, 0x14, 0x00, 0x32, 0x80, 0x43, 0xde, 0x00, 0xf3, 0xba, 0x6a, 0x3a, 0x8a, 0xd8, 0xb6,
	0x8b, 0x7e, 0x87, 0x78, 0x22, 0x8d, 0xd5, 0xae, 0x44, 0x69, 0x27, 0x63, 0x3f, 0xfb, 0xba, 0x4c,
	0x55, 0xd1, 0x4e, 0xa6, 0x51, 0xe4, 0x1c, 0x7a, 0xe1, 0x5a, 0x78, 0xdc, 0xe7, 0xf7, 0x1f, 0x37,
	0x54, 0xc2, 0x5a, 0x8e, 0x8e, 0xe1, 0xe8, 0x1a, 0xb9, 0xb6, 0x29, 0x72, 0xcb, 0xf0, 0xb7, 0x43,
	0x63, 0xc9, 0x3b, 0xd8, 0xb3, 0xed, 0x53, 0xfd, 0x6b, 0x41, 0x3c, 0x49, 0x0d, 0x44, 0xeb, 0x28,
	0xff, 0x7f, 0x9d, 0x7c, 0x58, 0xef, 0xd2, 0x5f, 0x76, 0x72, 0x7b, 0xf3, 0xd6, 0x9b, 0x76, 0x0e,
	0xc7, 0x9f, 0xb9, 0x54, 0x52, 0x64, 0xbc, 0x60, 0x68, 0x2a, 0x25, 0x0d, 0xae, 0x96, 0x23, 0x6b,
	0x9d, 0x61, 0x3a, 0x1b, 0x07, 0xfd, 0x04, 0x07, 0x23, 0x59, 0xdb, 0xa5, 0x90, 0x39, 0x49, 0xa0,
	0x57, 0xa2, 0x31, 0x3c, 0x6f, 0x2b, 0x3d, 0x49, 0x9a, 0xef, 0x35, 0x69, 0xbf, 0xd7, 0x64, 0x24,
	0x6b, 0xd6, 0x42, 0x97, 0x43, 0x78, 0x95, 0xa9, 0xb2, 0xad, 0xac, 0xd2, 0xea, 0x17, 0x66, 0x36,
	0x29, 0x44, 0x9a, 0x84, 0x2a, 0x4d, 0xa2, 0xab, 0xec, 0xb2, 0xdb, 0xfc, 0x12, 0xa6, 0x9d, 0xb4,
	0xeb, 0xe3, 0xbc, 0xff, 0x13, 0x00, 0x00, 0xff, 0xff, 0xc1, 0x64, 0xb2, 0x31, 0x25, 0x04, 0x00,
	0x00,
}
