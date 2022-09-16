// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: outbound/outbound_ack.proto

package types

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/router-protocol/router-chain/x/multichain/types"
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

type OutboundAck struct {
	ChainType             types.ChainType        `protobuf:"varint,1,opt,name=chainType,proto3,enum=routerprotocol.routerchain.multichain.ChainType" json:"chainType,omitempty"`
	ChainId               string                 `protobuf:"bytes,2,opt,name=chainId,proto3" json:"chainId,omitempty"`
	EventNonce            uint64                 `protobuf:"varint,3,opt,name=eventNonce,proto3" json:"eventNonce,omitempty"`
	BlockHeight           uint64                 `protobuf:"varint,4,opt,name=blockHeight,proto3" json:"blockHeight,omitempty"`
	OutboundTxNonce       uint64                 `protobuf:"varint,5,opt,name=outboundTxNonce,proto3" json:"outboundTxNonce,omitempty"`
	OutboundTxRequestedBy string                 `protobuf:"bytes,6,opt,name=outboundTxRequestedBy,proto3" json:"outboundTxRequestedBy,omitempty"`
	RelayerRouterAddress  string                 `protobuf:"bytes,7,opt,name=relayerRouterAddress,proto3" json:"relayerRouterAddress,omitempty"`
	DestinationTxHash     string                 `protobuf:"bytes,8,opt,name=destinationTxHash,proto3" json:"destinationTxHash,omitempty"`
	ContractAckResponses  []*ContractAckResponse `protobuf:"bytes,9,rep,name=contractAckResponses,proto3" json:"contractAckResponses,omitempty"`
	FeeConsumed           uint64                 `protobuf:"varint,10,opt,name=feeConsumed,proto3" json:"feeConsumed,omitempty"`
	Status                OutgoingTxStatus       `protobuf:"varint,11,opt,name=status,proto3,enum=routerprotocol.routerchain.outbound.OutgoingTxStatus" json:"status,omitempty"`
}

func (m *OutboundAck) Reset()         { *m = OutboundAck{} }
func (m *OutboundAck) String() string { return proto.CompactTextString(m) }
func (*OutboundAck) ProtoMessage()    {}
func (*OutboundAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a18876d7f6adbe4, []int{0}
}
func (m *OutboundAck) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OutboundAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OutboundAck.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OutboundAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OutboundAck.Merge(m, src)
}
func (m *OutboundAck) XXX_Size() int {
	return m.Size()
}
func (m *OutboundAck) XXX_DiscardUnknown() {
	xxx_messageInfo_OutboundAck.DiscardUnknown(m)
}

var xxx_messageInfo_OutboundAck proto.InternalMessageInfo

func (m *OutboundAck) GetChainType() types.ChainType {
	if m != nil {
		return m.ChainType
	}
	return types.CHAIN_TYPE_EVM
}

func (m *OutboundAck) GetChainId() string {
	if m != nil {
		return m.ChainId
	}
	return ""
}

func (m *OutboundAck) GetEventNonce() uint64 {
	if m != nil {
		return m.EventNonce
	}
	return 0
}

func (m *OutboundAck) GetBlockHeight() uint64 {
	if m != nil {
		return m.BlockHeight
	}
	return 0
}

func (m *OutboundAck) GetOutboundTxNonce() uint64 {
	if m != nil {
		return m.OutboundTxNonce
	}
	return 0
}

func (m *OutboundAck) GetOutboundTxRequestedBy() string {
	if m != nil {
		return m.OutboundTxRequestedBy
	}
	return ""
}

func (m *OutboundAck) GetRelayerRouterAddress() string {
	if m != nil {
		return m.RelayerRouterAddress
	}
	return ""
}

func (m *OutboundAck) GetDestinationTxHash() string {
	if m != nil {
		return m.DestinationTxHash
	}
	return ""
}

func (m *OutboundAck) GetContractAckResponses() []*ContractAckResponse {
	if m != nil {
		return m.ContractAckResponses
	}
	return nil
}

func (m *OutboundAck) GetFeeConsumed() uint64 {
	if m != nil {
		return m.FeeConsumed
	}
	return 0
}

func (m *OutboundAck) GetStatus() OutgoingTxStatus {
	if m != nil {
		return m.Status
	}
	return OUTGOING_TX_CREATED
}

func init() {
	proto.RegisterType((*OutboundAck)(nil), "routerprotocol.routerchain.outbound.OutboundAck")
}

func init() { proto.RegisterFile("outbound/outbound_ack.proto", fileDescriptor_2a18876d7f6adbe4) }

var fileDescriptor_2a18876d7f6adbe4 = []byte{
	// 448 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x8d, 0x69, 0x48, 0xc9, 0x46, 0x02, 0xb1, 0x2a, 0xd2, 0xaa, 0x95, 0x2c, 0xf3, 0x71, 0xf0,
	0x01, 0x1c, 0x14, 0x40, 0xea, 0x35, 0xcd, 0xa5, 0x1c, 0x68, 0xa5, 0x25, 0x27, 0x2e, 0x96, 0xb3,
	0x1e, 0x12, 0x2b, 0xce, 0x6e, 0xd8, 0x9d, 0x45, 0xce, 0xbf, 0xe0, 0x67, 0x71, 0xec, 0x11, 0x71,
	0x42, 0xc9, 0x1f, 0x41, 0x59, 0x7b, 0x9b, 0x88, 0x46, 0xa8, 0x17, 0xcb, 0x33, 0xf3, 0xde, 0xf3,
	0xdb, 0x9d, 0x67, 0x72, 0xa6, 0x2c, 0x4e, 0x94, 0x95, 0x79, 0xdf, 0xbf, 0xa4, 0x99, 0x98, 0x27,
	0x4b, 0xad, 0x50, 0xd1, 0x97, 0x5a, 0x59, 0x04, 0xed, 0x0a, 0xa1, 0xca, 0xa4, 0x2e, 0xc5, 0x2c,
	0x2b, 0x64, 0xe2, 0xe1, 0xa7, 0xcf, 0xf7, 0x15, 0xa6, 0xaa, 0x90, 0xd3, 0x14, 0xab, 0xd4, 0x60,
	0x86, 0xd6, 0xd4, 0x3a, 0xa7, 0x67, 0x0b, 0x5b, 0x62, 0xe1, 0x78, 0x7d, 0xf7, 0x4c, 0x71, 0xb5,
	0x84, 0x66, 0xf8, 0xea, 0x96, 0x2f, 0x94, 0x44, 0x9d, 0x09, 0xdc, 0x3a, 0x48, 0x35, 0x98, 0xa5,
	0x92, 0xa6, 0x41, 0xbd, 0xf8, 0xdd, 0x26, 0xbd, 0xeb, 0x06, 0x38, 0x14, 0x73, 0x7a, 0x45, 0xba,
	0x4e, 0x69, 0xbc, 0x5a, 0x02, 0x0b, 0xa2, 0x20, 0x7e, 0x3c, 0x78, 0x9b, 0xfc, 0xc7, 0xee, 0xce,
	0x41, 0x32, 0xf2, 0x3c, 0xbe, 0x93, 0xa0, 0x8c, 0x1c, 0xbb, 0xe2, 0x63, 0xce, 0x1e, 0x44, 0x41,
	0xdc, 0xe5, 0xbe, 0xa4, 0x21, 0x21, 0xf0, 0x1d, 0x24, 0x5e, 0x29, 0x29, 0x80, 0x1d, 0x45, 0x41,
	0xdc, 0xe6, 0x7b, 0x1d, 0x1a, 0x91, 0xde, 0xa4, 0x54, 0x62, 0x7e, 0x09, 0xc5, 0x74, 0x86, 0xac,
	0xed, 0x00, 0xfb, 0x2d, 0x1a, 0x93, 0x27, 0xfe, 0x8c, 0xe3, 0xaa, 0x96, 0x79, 0xe8, 0x50, 0xff,
	0xb6, 0xe9, 0x7b, 0xf2, 0x6c, 0xd7, 0xe2, 0xf0, 0xcd, 0x82, 0x41, 0xc8, 0x2f, 0x56, 0xac, 0xe3,
	0x3c, 0x1d, 0x1e, 0xd2, 0x01, 0x39, 0xd1, 0x50, 0x66, 0x2b, 0xd0, 0xdc, 0x9d, 0x78, 0x98, 0xe7,
	0x1a, 0x8c, 0x61, 0xc7, 0x8e, 0x74, 0x70, 0x46, 0x5f, 0x93, 0xa7, 0x39, 0x18, 0x2c, 0x64, 0x86,
	0x85, 0x92, 0xe3, 0xea, 0x32, 0x33, 0x33, 0xf6, 0xc8, 0x11, 0xee, 0x0e, 0x68, 0x49, 0x4e, 0xfc,
	0x72, 0x86, 0x62, 0xce, 0x9b, 0xd5, 0x18, 0xd6, 0x8d, 0x8e, 0xe2, 0xde, 0xe0, 0x3c, 0xb9, 0x47,
	0x4e, 0x92, 0xd1, 0x5d, 0x01, 0x7e, 0x50, 0x75, 0x7b, 0xa3, 0x5f, 0x01, 0x46, 0x4a, 0x1a, 0xbb,
	0x80, 0x9c, 0x91, 0xfa, 0x46, 0xf7, 0x5a, 0xf4, 0x13, 0xe9, 0xd4, 0x01, 0x63, 0x3d, 0xb7, 0xfa,
	0x0f, 0xf7, 0x72, 0x70, 0xdd, 0xe4, 0x73, 0x5c, 0x7d, 0x76, 0x64, 0xde, 0x88, 0x5c, 0xf0, 0x9f,
	0xeb, 0x30, 0xb8, 0x59, 0x87, 0xc1, 0x9f, 0x75, 0x18, 0xfc, 0xd8, 0x84, 0xad, 0x9b, 0x4d, 0xd8,
	0xfa, 0xb5, 0x09, 0x5b, 0x5f, 0xce, 0xa7, 0x05, 0xce, 0xec, 0x24, 0x11, 0x6a, 0xd1, 0xaf, 0x35,
	0xdf, 0xf8, 0x6f, 0xf8, 0xba, 0x8e, 0x75, 0x75, 0xfb, 0xff, 0xf4, 0xb7, 0xd9, 0x36, 0x93, 0x8e,
	0xc3, 0xbd, 0xfb, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x2f, 0xa0, 0x22, 0xa2, 0x61, 0x03, 0x00, 0x00,
}

func (m *OutboundAck) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OutboundAck) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OutboundAck) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Status != 0 {
		i = encodeVarintOutboundAck(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x58
	}
	if m.FeeConsumed != 0 {
		i = encodeVarintOutboundAck(dAtA, i, uint64(m.FeeConsumed))
		i--
		dAtA[i] = 0x50
	}
	if len(m.ContractAckResponses) > 0 {
		for iNdEx := len(m.ContractAckResponses) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ContractAckResponses[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintOutboundAck(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x4a
		}
	}
	if len(m.DestinationTxHash) > 0 {
		i -= len(m.DestinationTxHash)
		copy(dAtA[i:], m.DestinationTxHash)
		i = encodeVarintOutboundAck(dAtA, i, uint64(len(m.DestinationTxHash)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.RelayerRouterAddress) > 0 {
		i -= len(m.RelayerRouterAddress)
		copy(dAtA[i:], m.RelayerRouterAddress)
		i = encodeVarintOutboundAck(dAtA, i, uint64(len(m.RelayerRouterAddress)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.OutboundTxRequestedBy) > 0 {
		i -= len(m.OutboundTxRequestedBy)
		copy(dAtA[i:], m.OutboundTxRequestedBy)
		i = encodeVarintOutboundAck(dAtA, i, uint64(len(m.OutboundTxRequestedBy)))
		i--
		dAtA[i] = 0x32
	}
	if m.OutboundTxNonce != 0 {
		i = encodeVarintOutboundAck(dAtA, i, uint64(m.OutboundTxNonce))
		i--
		dAtA[i] = 0x28
	}
	if m.BlockHeight != 0 {
		i = encodeVarintOutboundAck(dAtA, i, uint64(m.BlockHeight))
		i--
		dAtA[i] = 0x20
	}
	if m.EventNonce != 0 {
		i = encodeVarintOutboundAck(dAtA, i, uint64(m.EventNonce))
		i--
		dAtA[i] = 0x18
	}
	if len(m.ChainId) > 0 {
		i -= len(m.ChainId)
		copy(dAtA[i:], m.ChainId)
		i = encodeVarintOutboundAck(dAtA, i, uint64(len(m.ChainId)))
		i--
		dAtA[i] = 0x12
	}
	if m.ChainType != 0 {
		i = encodeVarintOutboundAck(dAtA, i, uint64(m.ChainType))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintOutboundAck(dAtA []byte, offset int, v uint64) int {
	offset -= sovOutboundAck(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *OutboundAck) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ChainType != 0 {
		n += 1 + sovOutboundAck(uint64(m.ChainType))
	}
	l = len(m.ChainId)
	if l > 0 {
		n += 1 + l + sovOutboundAck(uint64(l))
	}
	if m.EventNonce != 0 {
		n += 1 + sovOutboundAck(uint64(m.EventNonce))
	}
	if m.BlockHeight != 0 {
		n += 1 + sovOutboundAck(uint64(m.BlockHeight))
	}
	if m.OutboundTxNonce != 0 {
		n += 1 + sovOutboundAck(uint64(m.OutboundTxNonce))
	}
	l = len(m.OutboundTxRequestedBy)
	if l > 0 {
		n += 1 + l + sovOutboundAck(uint64(l))
	}
	l = len(m.RelayerRouterAddress)
	if l > 0 {
		n += 1 + l + sovOutboundAck(uint64(l))
	}
	l = len(m.DestinationTxHash)
	if l > 0 {
		n += 1 + l + sovOutboundAck(uint64(l))
	}
	if len(m.ContractAckResponses) > 0 {
		for _, e := range m.ContractAckResponses {
			l = e.Size()
			n += 1 + l + sovOutboundAck(uint64(l))
		}
	}
	if m.FeeConsumed != 0 {
		n += 1 + sovOutboundAck(uint64(m.FeeConsumed))
	}
	if m.Status != 0 {
		n += 1 + sovOutboundAck(uint64(m.Status))
	}
	return n
}

func sovOutboundAck(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozOutboundAck(x uint64) (n int) {
	return sovOutboundAck(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *OutboundAck) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOutboundAck
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
			return fmt.Errorf("proto: OutboundAck: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OutboundAck: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainType", wireType)
			}
			m.ChainType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOutboundAck
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
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOutboundAck
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
				return ErrInvalidLengthOutboundAck
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOutboundAck
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EventNonce", wireType)
			}
			m.EventNonce = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOutboundAck
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
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockHeight", wireType)
			}
			m.BlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOutboundAck
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
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OutboundTxNonce", wireType)
			}
			m.OutboundTxNonce = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOutboundAck
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.OutboundTxNonce |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OutboundTxRequestedBy", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOutboundAck
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
				return ErrInvalidLengthOutboundAck
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOutboundAck
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OutboundTxRequestedBy = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RelayerRouterAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOutboundAck
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
				return ErrInvalidLengthOutboundAck
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOutboundAck
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RelayerRouterAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DestinationTxHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOutboundAck
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
				return ErrInvalidLengthOutboundAck
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOutboundAck
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DestinationTxHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractAckResponses", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOutboundAck
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthOutboundAck
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOutboundAck
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContractAckResponses = append(m.ContractAckResponses, &ContractAckResponse{})
			if err := m.ContractAckResponses[len(m.ContractAckResponses)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeeConsumed", wireType)
			}
			m.FeeConsumed = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOutboundAck
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FeeConsumed |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOutboundAck
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= OutgoingTxStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipOutboundAck(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOutboundAck
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
func skipOutboundAck(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowOutboundAck
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
					return 0, ErrIntOverflowOutboundAck
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
					return 0, ErrIntOverflowOutboundAck
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
				return 0, ErrInvalidLengthOutboundAck
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupOutboundAck
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthOutboundAck
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthOutboundAck        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowOutboundAck          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupOutboundAck = fmt.Errorf("proto: unexpected end of group")
)
