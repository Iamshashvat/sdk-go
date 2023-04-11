// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: routerchain/crosschain/crosschain_request.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
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

type CrosschainRequest struct {
	SrcChainId      string                                 `protobuf:"bytes,1,opt,name=src_chain_id,json=srcChainId,proto3" json:"src_chain_id,omitempty"`
	EventNonce      uint64                                 `protobuf:"varint,2,opt,name=event_nonce,json=eventNonce,proto3" json:"event_nonce,omitempty"`
	BlockHeight     uint64                                 `protobuf:"varint,3,opt,name=block_height,json=blockHeight,proto3" json:"block_height,omitempty"`
	SourceTxHash    string                                 `protobuf:"bytes,4,opt,name=source_tx_hash,json=sourceTxHash,proto3" json:"source_tx_hash,omitempty"`
	SrcTimestamp    uint64                                 `protobuf:"varint,5,opt,name=src_timestamp,json=srcTimestamp,proto3" json:"src_timestamp,omitempty"`
	SrcTxOrigin     string                                 `protobuf:"bytes,6,opt,name=src_tx_origin,json=srcTxOrigin,proto3" json:"src_tx_origin,omitempty"`
	RouteAmount     github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,7,opt,name=route_amount,json=routeAmount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"route_amount"`
	RouteRecipient  []byte                                 `protobuf:"bytes,8,opt,name=route_recipient,json=routeRecipient,proto3" json:"route_recipient,omitempty"`
	DestChainId     string                                 `protobuf:"bytes,9,opt,name=dest_chain_id,json=destChainId,proto3" json:"dest_chain_id,omitempty"`
	RequestSender   []byte                                 `protobuf:"bytes,10,opt,name=request_sender,json=requestSender,proto3" json:"request_sender,omitempty"`
	RequestMetadata []byte                                 `protobuf:"bytes,11,opt,name=request_metadata,json=requestMetadata,proto3" json:"request_metadata,omitempty"`
	RequestPacket   []byte                                 `protobuf:"bytes,12,opt,name=request_packet,json=requestPacket,proto3" json:"request_packet,omitempty"`
	SrcChainType    types.ChainType                        `protobuf:"varint,13,opt,name=src_chain_type,json=srcChainType,proto3,enum=routerprotocol.routerchain.multichain.ChainType" json:"src_chain_type,omitempty"`
	DestChainType   types.ChainType                        `protobuf:"varint,14,opt,name=dest_chain_type,json=destChainType,proto3,enum=routerprotocol.routerchain.multichain.ChainType" json:"dest_chain_type,omitempty"`
	Status          string                                 `protobuf:"bytes,15,opt,name=status,proto3" json:"status,omitempty"`
}

func (m *CrosschainRequest) Reset()         { *m = CrosschainRequest{} }
func (m *CrosschainRequest) String() string { return proto.CompactTextString(m) }
func (*CrosschainRequest) ProtoMessage()    {}
func (*CrosschainRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e9890b6bb70b2c0, []int{0}
}
func (m *CrosschainRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CrosschainRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CrosschainRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CrosschainRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CrosschainRequest.Merge(m, src)
}
func (m *CrosschainRequest) XXX_Size() int {
	return m.Size()
}
func (m *CrosschainRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CrosschainRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CrosschainRequest proto.InternalMessageInfo

func (m *CrosschainRequest) GetSrcChainId() string {
	if m != nil {
		return m.SrcChainId
	}
	return ""
}

func (m *CrosschainRequest) GetEventNonce() uint64 {
	if m != nil {
		return m.EventNonce
	}
	return 0
}

func (m *CrosschainRequest) GetBlockHeight() uint64 {
	if m != nil {
		return m.BlockHeight
	}
	return 0
}

func (m *CrosschainRequest) GetSourceTxHash() string {
	if m != nil {
		return m.SourceTxHash
	}
	return ""
}

func (m *CrosschainRequest) GetSrcTimestamp() uint64 {
	if m != nil {
		return m.SrcTimestamp
	}
	return 0
}

func (m *CrosschainRequest) GetSrcTxOrigin() string {
	if m != nil {
		return m.SrcTxOrigin
	}
	return ""
}

func (m *CrosschainRequest) GetRouteRecipient() []byte {
	if m != nil {
		return m.RouteRecipient
	}
	return nil
}

func (m *CrosschainRequest) GetDestChainId() string {
	if m != nil {
		return m.DestChainId
	}
	return ""
}

func (m *CrosschainRequest) GetRequestSender() []byte {
	if m != nil {
		return m.RequestSender
	}
	return nil
}

func (m *CrosschainRequest) GetRequestMetadata() []byte {
	if m != nil {
		return m.RequestMetadata
	}
	return nil
}

func (m *CrosschainRequest) GetRequestPacket() []byte {
	if m != nil {
		return m.RequestPacket
	}
	return nil
}

func (m *CrosschainRequest) GetSrcChainType() types.ChainType {
	if m != nil {
		return m.SrcChainType
	}
	return types.CHAIN_TYPE_ROUTER
}

func (m *CrosschainRequest) GetDestChainType() types.ChainType {
	if m != nil {
		return m.DestChainType
	}
	return types.CHAIN_TYPE_ROUTER
}

func (m *CrosschainRequest) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func init() {
	proto.RegisterType((*CrosschainRequest)(nil), "routerprotocol.routerchain.crosschain.CrosschainRequest")
}

func init() {
	proto.RegisterFile("routerchain/crosschain/crosschain_request.proto", fileDescriptor_0e9890b6bb70b2c0)
}

var fileDescriptor_0e9890b6bb70b2c0 = []byte{
	// 542 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0xcf, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0x1b, 0xd8, 0x06, 0x73, 0xd3, 0x14, 0x2c, 0x84, 0xac, 0x21, 0x65, 0x65, 0x30, 0x28,
	0x87, 0x25, 0x0c, 0x6e, 0xdc, 0xd8, 0x2e, 0xdb, 0x81, 0x5f, 0xa1, 0x42, 0x88, 0x4b, 0xe4, 0x3a,
	0x56, 0x63, 0xad, 0x89, 0x83, 0xed, 0x4c, 0xdd, 0x7f, 0xc1, 0x9f, 0x35, 0x89, 0xcb, 0x8e, 0x88,
	0xc3, 0x84, 0xda, 0x7f, 0x04, 0xf9, 0x39, 0x69, 0x0b, 0x07, 0x0e, 0x9c, 0xe2, 0xf7, 0xc9, 0xf3,
	0xf7, 0x3d, 0x7f, 0x9f, 0x8d, 0x62, 0x25, 0x6b, 0xc3, 0x15, 0xcb, 0xa9, 0x28, 0x63, 0xa6, 0xa4,
	0xd6, 0x7f, 0x2f, 0x53, 0xc5, 0xbf, 0xd6, 0x5c, 0x9b, 0xa8, 0x52, 0xd2, 0x48, 0xbc, 0xef, 0x36,
	0x40, 0xc0, 0xe4, 0x34, 0x5a, 0xdb, 0x1f, 0xad, 0x36, 0xed, 0x3c, 0x28, 0xea, 0xa9, 0x11, 0x8d,
	0x16, 0xc8, 0x98, 0x8b, 0x8a, 0x3b, 0x8d, 0x9d, 0x90, 0x49, 0x5d, 0x48, 0x1d, 0x8f, 0xa9, 0xe6,
	0xf1, 0xf9, 0xe1, 0x98, 0x1b, 0x7a, 0x18, 0x33, 0x29, 0xca, 0xe6, 0xff, 0xbd, 0x89, 0x9c, 0x48,
	0x58, 0xc6, 0x76, 0xe5, 0xe8, 0xde, 0xf7, 0x4d, 0x74, 0xf7, 0x78, 0x59, 0x21, 0x71, 0x5d, 0xe1,
	0x01, 0xf2, 0xb5, 0x62, 0xa9, 0xab, 0x21, 0x32, 0xe2, 0x0d, 0xbc, 0xe1, 0x76, 0x82, 0xb4, 0x62,
	0xc7, 0x16, 0x9d, 0x66, 0x78, 0x17, 0x75, 0xf9, 0x39, 0x2f, 0x4d, 0x5a, 0xca, 0x92, 0x71, 0x72,
	0x63, 0xe0, 0x0d, 0x37, 0x12, 0x04, 0xe8, 0xad, 0x25, 0xf8, 0x21, 0xf2, 0xc7, 0x53, 0xc9, 0xce,
	0xd2, 0x9c, 0x8b, 0x49, 0x6e, 0xc8, 0x4d, 0xc8, 0xe8, 0x02, 0x3b, 0x01, 0x84, 0x1f, 0xa3, 0x40,
	0xcb, 0x5a, 0x31, 0x9e, 0x9a, 0x59, 0x9a, 0x53, 0x9d, 0x93, 0x0d, 0xa8, 0xe3, 0x3b, 0x3a, 0x9a,
	0x9d, 0x50, 0x9d, 0xe3, 0x47, 0xa8, 0x67, 0x7b, 0x31, 0xa2, 0xe0, 0xda, 0xd0, 0xa2, 0x22, 0x9b,
	0xa0, 0x64, 0x1b, 0x1c, 0xb5, 0x0c, 0xef, 0x35, 0x49, 0xb3, 0x54, 0x2a, 0x31, 0x11, 0x25, 0xd9,
	0x02, 0xa5, 0xae, 0x4d, 0x9a, 0xbd, 0x03, 0x84, 0x3f, 0x20, 0x1f, 0x7c, 0x4d, 0x69, 0x21, 0xeb,
	0xd2, 0x90, 0x5b, 0x36, 0xe5, 0x28, 0xba, 0xbc, 0xde, 0xed, 0xfc, 0xbc, 0xde, 0x7d, 0x32, 0x11,
	0x26, 0xaf, 0xc7, 0x11, 0x93, 0x45, 0xdc, 0x38, 0xe9, 0x3e, 0x07, 0x3a, 0x3b, 0x8b, 0xad, 0xcf,
	0x3a, 0x3a, 0x2d, 0x4d, 0xd2, 0x05, 0x8d, 0xd7, 0x20, 0x81, 0x9f, 0xa2, 0xbe, 0x93, 0x54, 0x9c,
	0x89, 0x4a, 0xf0, 0xd2, 0x90, 0xdb, 0x03, 0x6f, 0xe8, 0x27, 0x01, 0xe0, 0xa4, 0xa5, 0xb6, 0xbf,
	0x8c, 0x6b, 0xb3, 0x72, 0x74, 0xdb, 0xf5, 0x67, 0x61, 0x6b, 0xe9, 0x3e, 0x0a, 0x9a, 0x5b, 0x91,
	0x6a, 0x5e, 0x66, 0x5c, 0x11, 0x04, 0x5a, 0xbd, 0x86, 0x7e, 0x04, 0x88, 0x9f, 0xa1, 0x3b, 0x6d,
	0x5a, 0xc1, 0x0d, 0xcd, 0xa8, 0xa1, 0xa4, 0x0b, 0x89, 0xfd, 0x86, 0xbf, 0x69, 0xf0, 0xba, 0x62,
	0x45, 0xd9, 0x19, 0x37, 0xc4, 0xff, 0x43, 0xf1, 0x3d, 0x40, 0xfc, 0x09, 0x05, 0xab, 0x69, 0xdb,
	0x93, 0x92, 0xde, 0xc0, 0x1b, 0x06, 0x2f, 0x9e, 0x47, 0xff, 0xb8, 0x96, 0xab, 0xab, 0x18, 0xc1,
	0x01, 0x46, 0x17, 0x15, 0x87, 0xa1, 0x2c, 0x23, 0xfc, 0x19, 0xf5, 0xd7, 0x0e, 0x0d, 0xc2, 0xc1,
	0x7f, 0x0a, 0xf7, 0x96, 0x46, 0x81, 0xf2, 0x7d, 0xb4, 0xa5, 0x0d, 0x35, 0xb5, 0x26, 0x7d, 0xf0,
	0xb1, 0x89, 0x8e, 0x46, 0x97, 0xf3, 0xd0, 0xbb, 0x9a, 0x87, 0xde, 0xaf, 0x79, 0xe8, 0x7d, 0x5b,
	0x84, 0x9d, 0xab, 0x45, 0xd8, 0xf9, 0xb1, 0x08, 0x3b, 0x5f, 0x5e, 0xad, 0x8d, 0xd7, 0x55, 0x3b,
	0x68, 0xab, 0xb7, 0xb1, 0x7b, 0x57, 0xb3, 0xf5, 0x07, 0x0b, 0x63, 0x1f, 0x6f, 0x41, 0xe6, 0xcb,
	0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xe2, 0x60, 0x5d, 0x15, 0xd7, 0x03, 0x00, 0x00,
}

func (m *CrosschainRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CrosschainRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CrosschainRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Status) > 0 {
		i -= len(m.Status)
		copy(dAtA[i:], m.Status)
		i = encodeVarintCrosschainRequest(dAtA, i, uint64(len(m.Status)))
		i--
		dAtA[i] = 0x7a
	}
	if m.DestChainType != 0 {
		i = encodeVarintCrosschainRequest(dAtA, i, uint64(m.DestChainType))
		i--
		dAtA[i] = 0x70
	}
	if m.SrcChainType != 0 {
		i = encodeVarintCrosschainRequest(dAtA, i, uint64(m.SrcChainType))
		i--
		dAtA[i] = 0x68
	}
	if len(m.RequestPacket) > 0 {
		i -= len(m.RequestPacket)
		copy(dAtA[i:], m.RequestPacket)
		i = encodeVarintCrosschainRequest(dAtA, i, uint64(len(m.RequestPacket)))
		i--
		dAtA[i] = 0x62
	}
	if len(m.RequestMetadata) > 0 {
		i -= len(m.RequestMetadata)
		copy(dAtA[i:], m.RequestMetadata)
		i = encodeVarintCrosschainRequest(dAtA, i, uint64(len(m.RequestMetadata)))
		i--
		dAtA[i] = 0x5a
	}
	if len(m.RequestSender) > 0 {
		i -= len(m.RequestSender)
		copy(dAtA[i:], m.RequestSender)
		i = encodeVarintCrosschainRequest(dAtA, i, uint64(len(m.RequestSender)))
		i--
		dAtA[i] = 0x52
	}
	if len(m.DestChainId) > 0 {
		i -= len(m.DestChainId)
		copy(dAtA[i:], m.DestChainId)
		i = encodeVarintCrosschainRequest(dAtA, i, uint64(len(m.DestChainId)))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.RouteRecipient) > 0 {
		i -= len(m.RouteRecipient)
		copy(dAtA[i:], m.RouteRecipient)
		i = encodeVarintCrosschainRequest(dAtA, i, uint64(len(m.RouteRecipient)))
		i--
		dAtA[i] = 0x42
	}
	{
		size := m.RouteAmount.Size()
		i -= size
		if _, err := m.RouteAmount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintCrosschainRequest(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	if len(m.SrcTxOrigin) > 0 {
		i -= len(m.SrcTxOrigin)
		copy(dAtA[i:], m.SrcTxOrigin)
		i = encodeVarintCrosschainRequest(dAtA, i, uint64(len(m.SrcTxOrigin)))
		i--
		dAtA[i] = 0x32
	}
	if m.SrcTimestamp != 0 {
		i = encodeVarintCrosschainRequest(dAtA, i, uint64(m.SrcTimestamp))
		i--
		dAtA[i] = 0x28
	}
	if len(m.SourceTxHash) > 0 {
		i -= len(m.SourceTxHash)
		copy(dAtA[i:], m.SourceTxHash)
		i = encodeVarintCrosschainRequest(dAtA, i, uint64(len(m.SourceTxHash)))
		i--
		dAtA[i] = 0x22
	}
	if m.BlockHeight != 0 {
		i = encodeVarintCrosschainRequest(dAtA, i, uint64(m.BlockHeight))
		i--
		dAtA[i] = 0x18
	}
	if m.EventNonce != 0 {
		i = encodeVarintCrosschainRequest(dAtA, i, uint64(m.EventNonce))
		i--
		dAtA[i] = 0x10
	}
	if len(m.SrcChainId) > 0 {
		i -= len(m.SrcChainId)
		copy(dAtA[i:], m.SrcChainId)
		i = encodeVarintCrosschainRequest(dAtA, i, uint64(len(m.SrcChainId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintCrosschainRequest(dAtA []byte, offset int, v uint64) int {
	offset -= sovCrosschainRequest(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *CrosschainRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.SrcChainId)
	if l > 0 {
		n += 1 + l + sovCrosschainRequest(uint64(l))
	}
	if m.EventNonce != 0 {
		n += 1 + sovCrosschainRequest(uint64(m.EventNonce))
	}
	if m.BlockHeight != 0 {
		n += 1 + sovCrosschainRequest(uint64(m.BlockHeight))
	}
	l = len(m.SourceTxHash)
	if l > 0 {
		n += 1 + l + sovCrosschainRequest(uint64(l))
	}
	if m.SrcTimestamp != 0 {
		n += 1 + sovCrosschainRequest(uint64(m.SrcTimestamp))
	}
	l = len(m.SrcTxOrigin)
	if l > 0 {
		n += 1 + l + sovCrosschainRequest(uint64(l))
	}
	l = m.RouteAmount.Size()
	n += 1 + l + sovCrosschainRequest(uint64(l))
	l = len(m.RouteRecipient)
	if l > 0 {
		n += 1 + l + sovCrosschainRequest(uint64(l))
	}
	l = len(m.DestChainId)
	if l > 0 {
		n += 1 + l + sovCrosschainRequest(uint64(l))
	}
	l = len(m.RequestSender)
	if l > 0 {
		n += 1 + l + sovCrosschainRequest(uint64(l))
	}
	l = len(m.RequestMetadata)
	if l > 0 {
		n += 1 + l + sovCrosschainRequest(uint64(l))
	}
	l = len(m.RequestPacket)
	if l > 0 {
		n += 1 + l + sovCrosschainRequest(uint64(l))
	}
	if m.SrcChainType != 0 {
		n += 1 + sovCrosschainRequest(uint64(m.SrcChainType))
	}
	if m.DestChainType != 0 {
		n += 1 + sovCrosschainRequest(uint64(m.DestChainType))
	}
	l = len(m.Status)
	if l > 0 {
		n += 1 + l + sovCrosschainRequest(uint64(l))
	}
	return n
}

func sovCrosschainRequest(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCrosschainRequest(x uint64) (n int) {
	return sovCrosschainRequest(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CrosschainRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCrosschainRequest
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
			return fmt.Errorf("proto: CrosschainRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CrosschainRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SrcChainId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCrosschainRequest
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
				return ErrInvalidLengthCrosschainRequest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCrosschainRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SrcChainId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EventNonce", wireType)
			}
			m.EventNonce = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCrosschainRequest
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
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockHeight", wireType)
			}
			m.BlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCrosschainRequest
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
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SourceTxHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCrosschainRequest
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
				return ErrInvalidLengthCrosschainRequest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCrosschainRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SourceTxHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SrcTimestamp", wireType)
			}
			m.SrcTimestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCrosschainRequest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SrcTimestamp |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SrcTxOrigin", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCrosschainRequest
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
				return ErrInvalidLengthCrosschainRequest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCrosschainRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SrcTxOrigin = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RouteAmount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCrosschainRequest
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
				return ErrInvalidLengthCrosschainRequest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCrosschainRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.RouteAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RouteRecipient", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCrosschainRequest
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
				return ErrInvalidLengthCrosschainRequest
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthCrosschainRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RouteRecipient = append(m.RouteRecipient[:0], dAtA[iNdEx:postIndex]...)
			if m.RouteRecipient == nil {
				m.RouteRecipient = []byte{}
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DestChainId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCrosschainRequest
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
				return ErrInvalidLengthCrosschainRequest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCrosschainRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DestChainId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestSender", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCrosschainRequest
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
				return ErrInvalidLengthCrosschainRequest
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthCrosschainRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RequestSender = append(m.RequestSender[:0], dAtA[iNdEx:postIndex]...)
			if m.RequestSender == nil {
				m.RequestSender = []byte{}
			}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestMetadata", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCrosschainRequest
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
				return ErrInvalidLengthCrosschainRequest
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthCrosschainRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RequestMetadata = append(m.RequestMetadata[:0], dAtA[iNdEx:postIndex]...)
			if m.RequestMetadata == nil {
				m.RequestMetadata = []byte{}
			}
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestPacket", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCrosschainRequest
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
				return ErrInvalidLengthCrosschainRequest
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthCrosschainRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RequestPacket = append(m.RequestPacket[:0], dAtA[iNdEx:postIndex]...)
			if m.RequestPacket == nil {
				m.RequestPacket = []byte{}
			}
			iNdEx = postIndex
		case 13:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SrcChainType", wireType)
			}
			m.SrcChainType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCrosschainRequest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SrcChainType |= types.ChainType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 14:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DestChainType", wireType)
			}
			m.DestChainType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCrosschainRequest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DestChainType |= types.ChainType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 15:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCrosschainRequest
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
				return ErrInvalidLengthCrosschainRequest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCrosschainRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Status = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCrosschainRequest(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCrosschainRequest
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
func skipCrosschainRequest(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCrosschainRequest
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
					return 0, ErrIntOverflowCrosschainRequest
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
					return 0, ErrIntOverflowCrosschainRequest
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
				return 0, ErrInvalidLengthCrosschainRequest
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCrosschainRequest
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCrosschainRequest
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCrosschainRequest        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCrosschainRequest          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCrosschainRequest = fmt.Errorf("proto: unexpected end of group")
)
