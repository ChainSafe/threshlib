// Code generated by protoc-gen-go. DO NOT EDIT.
// source: eddsa-signing.proto

package signing

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

//
// Represents a BROADCAST message sent to all parties during Round 1 of the EDDSA TSS signing protocol.
type SignRound1Message struct {
	Commitment           []byte   `protobuf:"bytes,1,opt,name=commitment,proto3" json:"commitment,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignRound1Message) Reset()         { *m = SignRound1Message{} }
func (m *SignRound1Message) String() string { return proto.CompactTextString(m) }
func (*SignRound1Message) ProtoMessage()    {}
func (*SignRound1Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_cf83f80fc7454980, []int{0}
}

func (m *SignRound1Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignRound1Message.Unmarshal(m, b)
}
func (m *SignRound1Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignRound1Message.Marshal(b, m, deterministic)
}
func (m *SignRound1Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignRound1Message.Merge(m, src)
}
func (m *SignRound1Message) XXX_Size() int {
	return xxx_messageInfo_SignRound1Message.Size(m)
}
func (m *SignRound1Message) XXX_DiscardUnknown() {
	xxx_messageInfo_SignRound1Message.DiscardUnknown(m)
}

var xxx_messageInfo_SignRound1Message proto.InternalMessageInfo

func (m *SignRound1Message) GetCommitment() []byte {
	if m != nil {
		return m.Commitment
	}
	return nil
}

//
// Represents a BROADCAST message sent to all parties during Round 2 of the EDDSA TSS signing protocol.
type SignRound2Message struct {
	DeCommitment         [][]byte `protobuf:"bytes,1,rep,name=de_commitment,json=deCommitment,proto3" json:"de_commitment,omitempty"`
	ProofAlphaX          []byte   `protobuf:"bytes,2,opt,name=proof_alpha_x,json=proofAlphaX,proto3" json:"proof_alpha_x,omitempty"`
	ProofAlphaY          []byte   `protobuf:"bytes,3,opt,name=proof_alpha_y,json=proofAlphaY,proto3" json:"proof_alpha_y,omitempty"`
	ProofT               []byte   `protobuf:"bytes,4,opt,name=proof_t,json=proofT,proto3" json:"proof_t,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignRound2Message) Reset()         { *m = SignRound2Message{} }
func (m *SignRound2Message) String() string { return proto.CompactTextString(m) }
func (*SignRound2Message) ProtoMessage()    {}
func (*SignRound2Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_cf83f80fc7454980, []int{1}
}

func (m *SignRound2Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignRound2Message.Unmarshal(m, b)
}
func (m *SignRound2Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignRound2Message.Marshal(b, m, deterministic)
}
func (m *SignRound2Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignRound2Message.Merge(m, src)
}
func (m *SignRound2Message) XXX_Size() int {
	return xxx_messageInfo_SignRound2Message.Size(m)
}
func (m *SignRound2Message) XXX_DiscardUnknown() {
	xxx_messageInfo_SignRound2Message.DiscardUnknown(m)
}

var xxx_messageInfo_SignRound2Message proto.InternalMessageInfo

func (m *SignRound2Message) GetDeCommitment() [][]byte {
	if m != nil {
		return m.DeCommitment
	}
	return nil
}

func (m *SignRound2Message) GetProofAlphaX() []byte {
	if m != nil {
		return m.ProofAlphaX
	}
	return nil
}

func (m *SignRound2Message) GetProofAlphaY() []byte {
	if m != nil {
		return m.ProofAlphaY
	}
	return nil
}

func (m *SignRound2Message) GetProofT() []byte {
	if m != nil {
		return m.ProofT
	}
	return nil
}

//
// Represents a BROADCAST message sent to all parties during Round 3 of the EDDSA TSS signing protocol.
type SignRound3Message struct {
	S                    []byte   `protobuf:"bytes,1,opt,name=s,proto3" json:"s,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignRound3Message) Reset()         { *m = SignRound3Message{} }
func (m *SignRound3Message) String() string { return proto.CompactTextString(m) }
func (*SignRound3Message) ProtoMessage()    {}
func (*SignRound3Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_cf83f80fc7454980, []int{2}
}

func (m *SignRound3Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignRound3Message.Unmarshal(m, b)
}
func (m *SignRound3Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignRound3Message.Marshal(b, m, deterministic)
}
func (m *SignRound3Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignRound3Message.Merge(m, src)
}
func (m *SignRound3Message) XXX_Size() int {
	return xxx_messageInfo_SignRound3Message.Size(m)
}
func (m *SignRound3Message) XXX_DiscardUnknown() {
	xxx_messageInfo_SignRound3Message.DiscardUnknown(m)
}

var xxx_messageInfo_SignRound3Message proto.InternalMessageInfo

func (m *SignRound3Message) GetS() []byte {
	if m != nil {
		return m.S
	}
	return nil
}

func init() {
	proto.RegisterType((*SignRound1Message)(nil), "binance.tsslib.eddsa.signing.SignRound1Message")
	proto.RegisterType((*SignRound2Message)(nil), "binance.tsslib.eddsa.signing.SignRound2Message")
	proto.RegisterType((*SignRound3Message)(nil), "binance.tsslib.eddsa.signing.SignRound3Message")
}

func init() { proto.RegisterFile("eddsa-signing.proto", fileDescriptor_cf83f80fc7454980) }

var fileDescriptor_cf83f80fc7454980 = []byte{
	// 219 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4e, 0x4d, 0x49, 0x29,
	0x4e, 0xd4, 0x2d, 0xce, 0x4c, 0xcf, 0xcb, 0xcc, 0x4b, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x92, 0x49, 0xca, 0xcc, 0x4b, 0xcc, 0x4b, 0x4e, 0xd5, 0x2b, 0x29, 0x2e, 0xce, 0xc9, 0x4c, 0xd2,
	0x03, 0xab, 0xd1, 0x83, 0xaa, 0x51, 0x32, 0xe6, 0x12, 0x0c, 0xce, 0x4c, 0xcf, 0x0b, 0xca, 0x2f,
	0xcd, 0x4b, 0x31, 0xf4, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0x15, 0x92, 0xe3, 0xe2, 0x4a, 0xce,
	0xcf, 0xcd, 0xcd, 0x2c, 0xc9, 0x4d, 0xcd, 0x2b, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x09, 0x42,
	0x12, 0x51, 0x9a, 0xc9, 0x88, 0xa4, 0xcb, 0x08, 0xa6, 0x4b, 0x99, 0x8b, 0x37, 0x25, 0x35, 0x1e,
	0x45, 0x23, 0xb3, 0x06, 0x4f, 0x10, 0x4f, 0x4a, 0xaa, 0x33, 0x5c, 0x4c, 0x48, 0x89, 0x8b, 0xb7,
	0xa0, 0x28, 0x3f, 0x3f, 0x2d, 0x3e, 0x31, 0xa7, 0x20, 0x23, 0x31, 0xbe, 0x42, 0x82, 0x09, 0x6c,
	0x3a, 0x37, 0x58, 0xd0, 0x11, 0x24, 0x16, 0x81, 0xae, 0xa6, 0x52, 0x82, 0x19, 0x5d, 0x4d, 0xa4,
	0x90, 0x38, 0x17, 0x3b, 0x44, 0x4d, 0x89, 0x04, 0x0b, 0x58, 0x96, 0x0d, 0xcc, 0x0d, 0x51, 0x52,
	0x44, 0x72, 0x9a, 0x31, 0xcc, 0x69, 0x3c, 0x5c, 0x8c, 0xc5, 0x50, 0x7f, 0x30, 0x16, 0x3b, 0xf1,
	0x47, 0xf1, 0x82, 0x03, 0x41, 0x1f, 0x1a, 0x08, 0x49, 0x6c, 0xe0, 0x90, 0x32, 0x06, 0x04, 0x00,
	0x00, 0xff, 0xff, 0x26, 0xfd, 0x9b, 0xd6, 0x40, 0x01, 0x00, 0x00,
}
