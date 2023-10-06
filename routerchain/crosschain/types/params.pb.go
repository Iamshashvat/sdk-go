// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: routerchain/crosschain/params.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

// Params defines the parameters for the module.
type Params struct {
	InboundGasPrice       uint64 `protobuf:"varint,1,opt,name=inboundGasPrice,proto3" json:"inboundGasPrice,omitempty"`
	MinimumRelayerFees    uint64 `protobuf:"varint,2,opt,name=minimumRelayerFees,proto3" json:"minimumRelayerFees,omitempty"`
	CleanupInterval       int64  `protobuf:"varint,3,opt,name=cleanup_interval,json=cleanupInterval,proto3" json:"cleanup_interval,omitempty"`
	BlockedRetryInterval  int64  `protobuf:"varint,4,opt,name=blocked_retry_interval,json=blockedRetryInterval,proto3" json:"blocked_retry_interval,omitempty"`
	RouterAdmin           string `protobuf:"bytes,5,opt,name=router_admin,json=routerAdmin,proto3" json:"router_admin,omitempty"`
	BlockedExpiryInterval int64  `protobuf:"varint,6,opt,name=blocked_expiry_interval,json=blockedExpiryInterval,proto3" json:"blocked_expiry_interval,omitempty"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_fad6f7037c606078, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetInboundGasPrice() uint64 {
	if m != nil {
		return m.InboundGasPrice
	}
	return 0
}

func (m *Params) GetMinimumRelayerFees() uint64 {
	if m != nil {
		return m.MinimumRelayerFees
	}
	return 0
}

func (m *Params) GetCleanupInterval() int64 {
	if m != nil {
		return m.CleanupInterval
	}
	return 0
}

func (m *Params) GetBlockedRetryInterval() int64 {
	if m != nil {
		return m.BlockedRetryInterval
	}
	return 0
}

func (m *Params) GetRouterAdmin() string {
	if m != nil {
		return m.RouterAdmin
	}
	return ""
}

func (m *Params) GetBlockedExpiryInterval() int64 {
	if m != nil {
		return m.BlockedExpiryInterval
	}
	return 0
}

func init() {
	proto.RegisterType((*Params)(nil), "routerprotocol.routerchain.crosschain.Params")
}

func init() {
	proto.RegisterFile("routerchain/crosschain/params.proto", fileDescriptor_fad6f7037c606078)
}

var fileDescriptor_fad6f7037c606078 = []byte{
	// 325 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xb1, 0x4b, 0x03, 0x31,
	0x14, 0xc6, 0x2f, 0x6d, 0x2d, 0x18, 0x85, 0x4a, 0xa8, 0x7a, 0x38, 0x9c, 0x55, 0x11, 0xce, 0xc1,
	0xeb, 0xa0, 0x38, 0x74, 0x53, 0x50, 0x71, 0x2b, 0x87, 0x93, 0x4b, 0xc9, 0xa5, 0x8f, 0x36, 0x78,
	0x97, 0x1c, 0x49, 0x4e, 0xda, 0xff, 0xc2, 0xd1, 0xd1, 0xc1, 0x3f, 0xc6, 0xb1, 0xa3, 0xa3, 0xb4,
	0xff, 0x88, 0x34, 0xb9, 0xb3, 0x45, 0xdc, 0xde, 0xfb, 0xbe, 0xef, 0xf7, 0x85, 0xf0, 0xf0, 0x89,
	0x92, 0x85, 0x01, 0xc5, 0xc6, 0x94, 0x8b, 0x2e, 0x53, 0x52, 0x6b, 0x37, 0xe6, 0x54, 0xd1, 0x4c,
	0x47, 0xb9, 0x92, 0x46, 0x92, 0x53, 0x17, 0xb2, 0x0b, 0x93, 0x69, 0xb4, 0xc6, 0x44, 0x2b, 0xe6,
	0xa0, 0x3d, 0x92, 0x23, 0x69, 0x43, 0xdd, 0xe5, 0xe4, 0xe0, 0xe3, 0x8f, 0x1a, 0x6e, 0xf6, 0x6d,
	0x1b, 0x09, 0x71, 0x8b, 0x8b, 0x44, 0x16, 0x62, 0x78, 0x4f, 0x75, 0x5f, 0x71, 0x06, 0x3e, 0xea,
	0xa0, 0xb0, 0x11, 0xff, 0x95, 0x49, 0x84, 0x49, 0xc6, 0x05, 0xcf, 0x8a, 0x2c, 0x86, 0x94, 0x4e,
	0x41, 0xdd, 0x01, 0x68, 0xbf, 0x66, 0xc3, 0xff, 0x38, 0xe4, 0x0c, 0xef, 0xb0, 0x14, 0xa8, 0x28,
	0xf2, 0x01, 0x17, 0x06, 0xd4, 0x0b, 0x4d, 0xfd, 0x7a, 0x07, 0x85, 0xf5, 0xb8, 0x55, 0xea, 0x0f,
	0xa5, 0x4c, 0x2e, 0xf1, 0x5e, 0x92, 0x4a, 0xf6, 0x0c, 0xc3, 0x81, 0x02, 0xa3, 0xa6, 0x2b, 0xa0,
	0x61, 0x81, 0x76, 0xe9, 0xc6, 0x4b, 0xf3, 0x97, 0x3a, 0xc2, 0xdb, 0xee, 0xd7, 0x03, 0x3a, 0xcc,
	0xb8, 0xf0, 0x37, 0x3a, 0x28, 0xdc, 0x8c, 0xb7, 0x9c, 0x76, 0xbd, 0x94, 0xc8, 0x15, 0xde, 0xaf,
	0x8a, 0x61, 0x92, 0xf3, 0xf5, 0xe6, 0xa6, 0x6d, 0xde, 0x2d, 0xed, 0x5b, 0xeb, 0x56, 0xd5, 0xbd,
	0xc6, 0xdb, 0xfb, 0xa1, 0x77, 0xf3, 0xf8, 0x39, 0x0f, 0xd0, 0x6c, 0x1e, 0xa0, 0xef, 0x79, 0x80,
	0x5e, 0x17, 0x81, 0x37, 0x5b, 0x04, 0xde, 0xd7, 0x22, 0xf0, 0x9e, 0x7a, 0x23, 0x6e, 0xc6, 0x45,
	0x12, 0x31, 0x99, 0x75, 0xdd, 0x7b, 0xe7, 0xd5, 0x25, 0xaa, 0xdd, 0xdd, 0x6c, 0xb2, 0x7e, 0x40,
	0x33, 0xcd, 0x41, 0x27, 0x4d, 0x9b, 0xbc, 0xf8, 0x09, 0x00, 0x00, 0xff, 0xff, 0xab, 0x57, 0x14,
	0x9a, 0xe7, 0x01, 0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.BlockedExpiryInterval != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.BlockedExpiryInterval))
		i--
		dAtA[i] = 0x30
	}
	if len(m.RouterAdmin) > 0 {
		i -= len(m.RouterAdmin)
		copy(dAtA[i:], m.RouterAdmin)
		i = encodeVarintParams(dAtA, i, uint64(len(m.RouterAdmin)))
		i--
		dAtA[i] = 0x2a
	}
	if m.BlockedRetryInterval != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.BlockedRetryInterval))
		i--
		dAtA[i] = 0x20
	}
	if m.CleanupInterval != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.CleanupInterval))
		i--
		dAtA[i] = 0x18
	}
	if m.MinimumRelayerFees != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.MinimumRelayerFees))
		i--
		dAtA[i] = 0x10
	}
	if m.InboundGasPrice != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.InboundGasPrice))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.InboundGasPrice != 0 {
		n += 1 + sovParams(uint64(m.InboundGasPrice))
	}
	if m.MinimumRelayerFees != 0 {
		n += 1 + sovParams(uint64(m.MinimumRelayerFees))
	}
	if m.CleanupInterval != 0 {
		n += 1 + sovParams(uint64(m.CleanupInterval))
	}
	if m.BlockedRetryInterval != 0 {
		n += 1 + sovParams(uint64(m.BlockedRetryInterval))
	}
	l = len(m.RouterAdmin)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	if m.BlockedExpiryInterval != 0 {
		n += 1 + sovParams(uint64(m.BlockedExpiryInterval))
	}
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field InboundGasPrice", wireType)
			}
			m.InboundGasPrice = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.InboundGasPrice |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinimumRelayerFees", wireType)
			}
			m.MinimumRelayerFees = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MinimumRelayerFees |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CleanupInterval", wireType)
			}
			m.CleanupInterval = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CleanupInterval |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockedRetryInterval", wireType)
			}
			m.BlockedRetryInterval = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlockedRetryInterval |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RouterAdmin", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RouterAdmin = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockedExpiryInterval", wireType)
			}
			m.BlockedExpiryInterval = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlockedExpiryInterval |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
