// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: inbound/events.proto

package types

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/router-protocol/sdk-go/routerchain/multichain/types"
	io "io"
	math "math"
	math_bits "math/bits"
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

type EventIncomingTxCreated struct {
	AttestationId        []byte           `protobuf:"bytes,1,opt,name=attestation_id,json=attestationId,proto3" json:"attestation_id,omitempty"`
	ChainType            types.ChainType  `protobuf:"varint,2,opt,name=chainType,proto3,enum=routerprotocol.routerchain.multichain.ChainType" json:"chainType,omitempty"`
	ChainId              string           `protobuf:"bytes,3,opt,name=chainId,proto3" json:"chainId,omitempty"`
	EventNonce           uint64           `protobuf:"varint,4,opt,name=eventNonce,proto3" json:"eventNonce,omitempty"`
	BlockHeight          uint64           `protobuf:"varint,5,opt,name=blockHeight,proto3" json:"blockHeight,omitempty"`
	SourceTxHash         string           `protobuf:"bytes,6,opt,name=sourceTxHash,proto3" json:"sourceTxHash,omitempty"`
	SourceSender         string           `protobuf:"bytes,7,opt,name=sourceSender,proto3" json:"sourceSender,omitempty"`
	RouterBridgeContract string           `protobuf:"bytes,8,opt,name=routerBridgeContract,proto3" json:"routerBridgeContract,omitempty"`
	Payload              []byte           `protobuf:"bytes,9,opt,name=payload,proto3" json:"payload,omitempty"`
	Status               IncomingTxStatus `protobuf:"varint,10,opt,name=status,proto3,enum=routerprotocol.routerchain.inbound.IncomingTxStatus" json:"status,omitempty"`
}

func (m *EventIncomingTxCreated) Reset()         { *m = EventIncomingTxCreated{} }
func (m *EventIncomingTxCreated) String() string { return proto.CompactTextString(m) }
func (*EventIncomingTxCreated) ProtoMessage()    {}
func (*EventIncomingTxCreated) Descriptor() ([]byte, []int) {
	return fileDescriptor_c68545bf45595e04, []int{0}
}
func (m *EventIncomingTxCreated) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventIncomingTxCreated) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventIncomingTxCreated.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventIncomingTxCreated) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventIncomingTxCreated.Merge(m, src)
}
func (m *EventIncomingTxCreated) XXX_Size() int {
	return m.Size()
}
func (m *EventIncomingTxCreated) XXX_DiscardUnknown() {
	xxx_messageInfo_EventIncomingTxCreated.DiscardUnknown(m)
}

var xxx_messageInfo_EventIncomingTxCreated proto.InternalMessageInfo

func (m *EventIncomingTxCreated) GetAttestationId() []byte {
	if m != nil {
		return m.AttestationId
	}
	return nil
}

func (m *EventIncomingTxCreated) GetChainType() types.ChainType {
	if m != nil {
		return m.ChainType
	}
	return types.CHAIN_TYPE_EVM
}

func (m *EventIncomingTxCreated) GetChainId() string {
	if m != nil {
		return m.ChainId
	}
	return ""
}

func (m *EventIncomingTxCreated) GetEventNonce() uint64 {
	if m != nil {
		return m.EventNonce
	}
	return 0
}

func (m *EventIncomingTxCreated) GetBlockHeight() uint64 {
	if m != nil {
		return m.BlockHeight
	}
	return 0
}

func (m *EventIncomingTxCreated) GetSourceTxHash() string {
	if m != nil {
		return m.SourceTxHash
	}
	return ""
}

func (m *EventIncomingTxCreated) GetSourceSender() string {
	if m != nil {
		return m.SourceSender
	}
	return ""
}

func (m *EventIncomingTxCreated) GetRouterBridgeContract() string {
	if m != nil {
		return m.RouterBridgeContract
	}
	return ""
}

func (m *EventIncomingTxCreated) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *EventIncomingTxCreated) GetStatus() IncomingTxStatus {
	if m != nil {
		return m.Status
	}
	return INCOMING_TX_CREATED
}

type EventIncomingTxDelegated struct {
	AttestationId      []byte          `protobuf:"bytes,1,opt,name=attestation_id,json=attestationId,proto3" json:"attestation_id,omitempty"`
	ChainType          types.ChainType `protobuf:"varint,2,opt,name=chainType,proto3,enum=routerprotocol.routerchain.multichain.ChainType" json:"chainType,omitempty"`
	ChainId            string          `protobuf:"bytes,3,opt,name=chainId,proto3" json:"chainId,omitempty"`
	EventNonce         uint64          `protobuf:"varint,4,opt,name=eventNonce,proto3" json:"eventNonce,omitempty"`
	DelegationResponse []byte          `protobuf:"bytes,5,opt,name=delegationResponse,proto3" json:"delegationResponse,omitempty"`
}

func (m *EventIncomingTxDelegated) Reset()         { *m = EventIncomingTxDelegated{} }
func (m *EventIncomingTxDelegated) String() string { return proto.CompactTextString(m) }
func (*EventIncomingTxDelegated) ProtoMessage()    {}
func (*EventIncomingTxDelegated) Descriptor() ([]byte, []int) {
	return fileDescriptor_c68545bf45595e04, []int{1}
}
func (m *EventIncomingTxDelegated) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventIncomingTxDelegated) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventIncomingTxDelegated.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventIncomingTxDelegated) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventIncomingTxDelegated.Merge(m, src)
}
func (m *EventIncomingTxDelegated) XXX_Size() int {
	return m.Size()
}
func (m *EventIncomingTxDelegated) XXX_DiscardUnknown() {
	xxx_messageInfo_EventIncomingTxDelegated.DiscardUnknown(m)
}

var xxx_messageInfo_EventIncomingTxDelegated proto.InternalMessageInfo

func (m *EventIncomingTxDelegated) GetAttestationId() []byte {
	if m != nil {
		return m.AttestationId
	}
	return nil
}

func (m *EventIncomingTxDelegated) GetChainType() types.ChainType {
	if m != nil {
		return m.ChainType
	}
	return types.CHAIN_TYPE_EVM
}

func (m *EventIncomingTxDelegated) GetChainId() string {
	if m != nil {
		return m.ChainId
	}
	return ""
}

func (m *EventIncomingTxDelegated) GetEventNonce() uint64 {
	if m != nil {
		return m.EventNonce
	}
	return 0
}

func (m *EventIncomingTxDelegated) GetDelegationResponse() []byte {
	if m != nil {
		return m.DelegationResponse
	}
	return nil
}

type EventIncomingTxDelegationFailed struct {
	AttestationId           []byte          `protobuf:"bytes,1,opt,name=attestation_id,json=attestationId,proto3" json:"attestation_id,omitempty"`
	ChainType               types.ChainType `protobuf:"varint,2,opt,name=chainType,proto3,enum=routerprotocol.routerchain.multichain.ChainType" json:"chainType,omitempty"`
	ChainId                 string          `protobuf:"bytes,3,opt,name=chainId,proto3" json:"chainId,omitempty"`
	EventNonce              uint64          `protobuf:"varint,4,opt,name=eventNonce,proto3" json:"eventNonce,omitempty"`
	DelegationErrorResponse string          `protobuf:"bytes,5,opt,name=delegationErrorResponse,proto3" json:"delegationErrorResponse,omitempty"`
}

func (m *EventIncomingTxDelegationFailed) Reset()         { *m = EventIncomingTxDelegationFailed{} }
func (m *EventIncomingTxDelegationFailed) String() string { return proto.CompactTextString(m) }
func (*EventIncomingTxDelegationFailed) ProtoMessage()    {}
func (*EventIncomingTxDelegationFailed) Descriptor() ([]byte, []int) {
	return fileDescriptor_c68545bf45595e04, []int{2}
}
func (m *EventIncomingTxDelegationFailed) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventIncomingTxDelegationFailed) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventIncomingTxDelegationFailed.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventIncomingTxDelegationFailed) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventIncomingTxDelegationFailed.Merge(m, src)
}
func (m *EventIncomingTxDelegationFailed) XXX_Size() int {
	return m.Size()
}
func (m *EventIncomingTxDelegationFailed) XXX_DiscardUnknown() {
	xxx_messageInfo_EventIncomingTxDelegationFailed.DiscardUnknown(m)
}

var xxx_messageInfo_EventIncomingTxDelegationFailed proto.InternalMessageInfo

func (m *EventIncomingTxDelegationFailed) GetAttestationId() []byte {
	if m != nil {
		return m.AttestationId
	}
	return nil
}

func (m *EventIncomingTxDelegationFailed) GetChainType() types.ChainType {
	if m != nil {
		return m.ChainType
	}
	return types.CHAIN_TYPE_EVM
}

func (m *EventIncomingTxDelegationFailed) GetChainId() string {
	if m != nil {
		return m.ChainId
	}
	return ""
}

func (m *EventIncomingTxDelegationFailed) GetEventNonce() uint64 {
	if m != nil {
		return m.EventNonce
	}
	return 0
}

func (m *EventIncomingTxDelegationFailed) GetDelegationErrorResponse() string {
	if m != nil {
		return m.DelegationErrorResponse
	}
	return ""
}

func init() {
	proto.RegisterType((*EventIncomingTxCreated)(nil), "routerprotocol.routerchain.inbound.EventIncomingTxCreated")
	proto.RegisterType((*EventIncomingTxDelegated)(nil), "routerprotocol.routerchain.inbound.EventIncomingTxDelegated")
	proto.RegisterType((*EventIncomingTxDelegationFailed)(nil), "routerprotocol.routerchain.inbound.EventIncomingTxDelegationFailed")
}

func init() { proto.RegisterFile("inbound/events.proto", fileDescriptor_c68545bf45595e04) }

var fileDescriptor_c68545bf45595e04 = []byte{
	// 463 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x54, 0xc1, 0x8a, 0x13, 0x41,
	0x10, 0x4d, 0xef, 0xae, 0x59, 0xd3, 0xc6, 0x3d, 0x34, 0x8b, 0x36, 0x2b, 0x8c, 0xc3, 0x80, 0x90,
	0x8b, 0x1d, 0x59, 0x05, 0x3d, 0x6f, 0x5c, 0xd9, 0x80, 0x2c, 0x38, 0x9b, 0x93, 0x97, 0xd0, 0x99,
	0x2e, 0x26, 0x8d, 0x93, 0xae, 0xa1, 0xa7, 0x47, 0x92, 0x7f, 0xf0, 0xe0, 0x67, 0x79, 0x73, 0x8f,
	0x1e, 0x25, 0xf9, 0x0c, 0x2f, 0x92, 0x9e, 0x0c, 0x93, 0x95, 0xa8, 0xe7, 0x5c, 0x1a, 0x5e, 0x55,
	0xbd, 0xea, 0x7a, 0xf5, 0xa0, 0xe8, 0xa9, 0x36, 0x13, 0x2c, 0x8d, 0xea, 0xc3, 0x67, 0x30, 0xae,
	0x10, 0xb9, 0x45, 0x87, 0x2c, 0xb2, 0x58, 0x3a, 0xb0, 0x1e, 0x24, 0x98, 0x89, 0x0a, 0x26, 0x53,
	0xa9, 0x8d, 0xd8, 0x10, 0xce, 0xc2, 0x9a, 0xa9, 0x4d, 0x82, 0x33, 0x6d, 0xd2, 0xb1, 0x9b, 0x8f,
	0x0b, 0x27, 0x5d, 0xb9, 0xe9, 0x72, 0xf6, 0x64, 0x56, 0x66, 0x4e, 0x7b, 0x56, 0xdf, 0xbf, 0x63,
	0xb7, 0xc8, 0xa1, 0x4a, 0x46, 0xdf, 0x0f, 0xe9, 0xa3, 0xcb, 0xf5, 0x9f, 0xc3, 0x0d, 0x7d, 0x34,
	0x1f, 0x58, 0x90, 0x0e, 0x14, 0x7b, 0x46, 0x4f, 0xa4, 0x73, 0xb0, 0xee, 0xa5, 0xd1, 0x8c, 0xb5,
	0xe2, 0x24, 0x24, 0xbd, 0x6e, 0xfc, 0x70, 0x2b, 0x3a, 0x54, 0xec, 0x9a, 0x76, 0x7c, 0xd7, 0xd1,
	0x22, 0x07, 0x7e, 0x10, 0x92, 0xde, 0xc9, 0xf9, 0x0b, 0xf1, 0x8f, 0xc1, 0x9b, 0x69, 0xc4, 0xa0,
	0xe6, 0xc5, 0x4d, 0x0b, 0xc6, 0xe9, 0xb1, 0x07, 0x43, 0xc5, 0x0f, 0x43, 0xd2, 0xeb, 0xc4, 0x35,
	0x64, 0x01, 0xa5, 0x7e, 0x3d, 0xd7, 0x68, 0x12, 0xe0, 0x47, 0x21, 0xe9, 0x1d, 0xc5, 0x5b, 0x11,
	0x16, 0xd2, 0x07, 0x93, 0x0c, 0x93, 0x4f, 0x57, 0xa0, 0xd3, 0xa9, 0xe3, 0xf7, 0x7c, 0xc1, 0x76,
	0x88, 0x45, 0xb4, 0x5b, 0x60, 0x69, 0x13, 0x18, 0xcd, 0xaf, 0x64, 0x31, 0xe5, 0x6d, 0xff, 0xc1,
	0x9d, 0x58, 0x53, 0x73, 0x03, 0x46, 0x81, 0xe5, 0xc7, 0xdb, 0x35, 0x55, 0x8c, 0x9d, 0xd3, 0xd3,
	0x4a, 0xd2, 0x85, 0xd5, 0x2a, 0x85, 0x01, 0x1a, 0x67, 0x65, 0xe2, 0xf8, 0x7d, 0x5f, 0xbb, 0x33,
	0xb7, 0xd6, 0x95, 0xcb, 0x45, 0x86, 0x52, 0xf1, 0x8e, 0xdf, 0x63, 0x0d, 0xd9, 0x7b, 0xda, 0xae,
	0x0c, 0xe3, 0xd4, 0xaf, 0xef, 0x95, 0xf8, 0xbf, 0xef, 0xa2, 0xf1, 0xeb, 0xc6, 0x73, 0xe3, 0x4d,
	0x8f, 0xe8, 0x17, 0xa1, 0xfc, 0x0f, 0x47, 0xdf, 0x42, 0x06, 0xe9, 0x7e, 0x7a, 0x2a, 0x28, 0x53,
	0xd5, 0xf4, 0x1a, 0x4d, 0x0c, 0x45, 0x8e, 0xa6, 0x00, 0x6f, 0x6d, 0x37, 0xde, 0x91, 0x89, 0xbe,
	0x1c, 0xd0, 0xa7, 0xbb, 0xd5, 0x6b, 0x34, 0xef, 0xa4, 0xce, 0xf6, 0x71, 0x09, 0x6f, 0xe8, 0xe3,
	0x46, 0xea, 0xa5, 0xb5, 0x68, 0xef, 0x6c, 0xa2, 0x13, 0xff, 0x2d, 0x7d, 0xf1, 0xe1, 0xdb, 0x32,
	0x20, 0xb7, 0xcb, 0x80, 0xfc, 0x5c, 0x06, 0xe4, 0xeb, 0x2a, 0x68, 0xdd, 0xae, 0x82, 0xd6, 0x8f,
	0x55, 0xd0, 0xfa, 0xf8, 0x3a, 0xd5, 0x6e, 0x5a, 0x4e, 0x44, 0x82, 0xb3, 0x7e, 0xa5, 0xe2, 0x79,
	0xad, 0xaa, 0xc6, 0xd5, 0xc9, 0x98, 0xf7, 0xeb, 0x0b, 0xb3, 0x3e, 0x1b, 0xc5, 0xa4, 0xed, 0xcb,
	0x5e, 0xfe, 0x0e, 0x00, 0x00, 0xff, 0xff, 0xfb, 0x9b, 0x38, 0x3d, 0xb3, 0x04, 0x00, 0x00,
}

func (m *EventIncomingTxCreated) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventIncomingTxCreated) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventIncomingTxCreated) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Status != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x50
	}
	if len(m.Payload) > 0 {
		i -= len(m.Payload)
		copy(dAtA[i:], m.Payload)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Payload)))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.RouterBridgeContract) > 0 {
		i -= len(m.RouterBridgeContract)
		copy(dAtA[i:], m.RouterBridgeContract)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.RouterBridgeContract)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.SourceSender) > 0 {
		i -= len(m.SourceSender)
		copy(dAtA[i:], m.SourceSender)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.SourceSender)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.SourceTxHash) > 0 {
		i -= len(m.SourceTxHash)
		copy(dAtA[i:], m.SourceTxHash)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.SourceTxHash)))
		i--
		dAtA[i] = 0x32
	}
	if m.BlockHeight != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.BlockHeight))
		i--
		dAtA[i] = 0x28
	}
	if m.EventNonce != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.EventNonce))
		i--
		dAtA[i] = 0x20
	}
	if len(m.ChainId) > 0 {
		i -= len(m.ChainId)
		copy(dAtA[i:], m.ChainId)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.ChainId)))
		i--
		dAtA[i] = 0x1a
	}
	if m.ChainType != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.ChainType))
		i--
		dAtA[i] = 0x10
	}
	if len(m.AttestationId) > 0 {
		i -= len(m.AttestationId)
		copy(dAtA[i:], m.AttestationId)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.AttestationId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventIncomingTxDelegated) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventIncomingTxDelegated) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventIncomingTxDelegated) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.DelegationResponse) > 0 {
		i -= len(m.DelegationResponse)
		copy(dAtA[i:], m.DelegationResponse)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.DelegationResponse)))
		i--
		dAtA[i] = 0x2a
	}
	if m.EventNonce != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.EventNonce))
		i--
		dAtA[i] = 0x20
	}
	if len(m.ChainId) > 0 {
		i -= len(m.ChainId)
		copy(dAtA[i:], m.ChainId)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.ChainId)))
		i--
		dAtA[i] = 0x1a
	}
	if m.ChainType != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.ChainType))
		i--
		dAtA[i] = 0x10
	}
	if len(m.AttestationId) > 0 {
		i -= len(m.AttestationId)
		copy(dAtA[i:], m.AttestationId)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.AttestationId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventIncomingTxDelegationFailed) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventIncomingTxDelegationFailed) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventIncomingTxDelegationFailed) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.DelegationErrorResponse) > 0 {
		i -= len(m.DelegationErrorResponse)
		copy(dAtA[i:], m.DelegationErrorResponse)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.DelegationErrorResponse)))
		i--
		dAtA[i] = 0x2a
	}
	if m.EventNonce != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.EventNonce))
		i--
		dAtA[i] = 0x20
	}
	if len(m.ChainId) > 0 {
		i -= len(m.ChainId)
		copy(dAtA[i:], m.ChainId)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.ChainId)))
		i--
		dAtA[i] = 0x1a
	}
	if m.ChainType != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.ChainType))
		i--
		dAtA[i] = 0x10
	}
	if len(m.AttestationId) > 0 {
		i -= len(m.AttestationId)
		copy(dAtA[i:], m.AttestationId)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.AttestationId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintEvents(dAtA []byte, offset int, v uint64) int {
	offset -= sovEvents(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EventIncomingTxCreated) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.AttestationId)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	if m.ChainType != 0 {
		n += 1 + sovEvents(uint64(m.ChainType))
	}
	l = len(m.ChainId)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	if m.EventNonce != 0 {
		n += 1 + sovEvents(uint64(m.EventNonce))
	}
	if m.BlockHeight != 0 {
		n += 1 + sovEvents(uint64(m.BlockHeight))
	}
	l = len(m.SourceTxHash)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.SourceSender)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.RouterBridgeContract)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.Payload)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	if m.Status != 0 {
		n += 1 + sovEvents(uint64(m.Status))
	}
	return n
}

func (m *EventIncomingTxDelegated) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.AttestationId)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	if m.ChainType != 0 {
		n += 1 + sovEvents(uint64(m.ChainType))
	}
	l = len(m.ChainId)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	if m.EventNonce != 0 {
		n += 1 + sovEvents(uint64(m.EventNonce))
	}
	l = len(m.DelegationResponse)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *EventIncomingTxDelegationFailed) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.AttestationId)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	if m.ChainType != 0 {
		n += 1 + sovEvents(uint64(m.ChainType))
	}
	l = len(m.ChainId)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	if m.EventNonce != 0 {
		n += 1 + sovEvents(uint64(m.EventNonce))
	}
	l = len(m.DelegationErrorResponse)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func sovEvents(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvents(x uint64) (n int) {
	return sovEvents(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EventIncomingTxCreated) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventIncomingTxCreated: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventIncomingTxCreated: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AttestationId", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AttestationId = append(m.AttestationId[:0], dAtA[iNdEx:postIndex]...)
			if m.AttestationId == nil {
				m.AttestationId = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainType", wireType)
			}
			m.ChainType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ChainType |= types.ChainType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EventNonce", wireType)
			}
			m.EventNonce = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EventNonce |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockHeight", wireType)
			}
			m.BlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlockHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SourceTxHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SourceTxHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SourceSender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SourceSender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RouterBridgeContract", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RouterBridgeContract = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Payload", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Payload = append(m.Payload[:0], dAtA[iNdEx:postIndex]...)
			if m.Payload == nil {
				m.Payload = []byte{}
			}
			iNdEx = postIndex
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= IncomingTxStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *EventIncomingTxDelegated) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventIncomingTxDelegated: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventIncomingTxDelegated: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AttestationId", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AttestationId = append(m.AttestationId[:0], dAtA[iNdEx:postIndex]...)
			if m.AttestationId == nil {
				m.AttestationId = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainType", wireType)
			}
			m.ChainType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ChainType |= types.ChainType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EventNonce", wireType)
			}
			m.EventNonce = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EventNonce |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DelegationResponse", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DelegationResponse = append(m.DelegationResponse[:0], dAtA[iNdEx:postIndex]...)
			if m.DelegationResponse == nil {
				m.DelegationResponse = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *EventIncomingTxDelegationFailed) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventIncomingTxDelegationFailed: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventIncomingTxDelegationFailed: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AttestationId", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AttestationId = append(m.AttestationId[:0], dAtA[iNdEx:postIndex]...)
			if m.AttestationId == nil {
				m.AttestationId = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainType", wireType)
			}
			m.ChainType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ChainType |= types.ChainType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EventNonce", wireType)
			}
			m.EventNonce = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EventNonce |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DelegationErrorResponse", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DelegationErrorResponse = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipEvents(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
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
				return 0, ErrInvalidLengthEvents
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEvents
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEvents
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEvents        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEvents          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEvents = fmt.Errorf("proto: unexpected end of group")
)
