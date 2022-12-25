// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: crosstalk/crosstalk_request_status.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

type CrossTalkRequestStatus int32

const (
	CROSSTALK_REQUEST_CREATED  CrossTalkRequestStatus = 0
	CROSSTALK_REQUEST_OBSERVED CrossTalkRequestStatus = 1
)

var CrossTalkRequestStatus_name = map[int32]string{
	0: "CROSSTALK_REQUEST_CREATED",
	1: "CROSSTALK_REQUEST_OBSERVED",
}

var CrossTalkRequestStatus_value = map[string]int32{
	"CROSSTALK_REQUEST_CREATED":  0,
	"CROSSTALK_REQUEST_OBSERVED": 1,
}

func (x CrossTalkRequestStatus) String() string {
	return proto.EnumName(CrossTalkRequestStatus_name, int32(x))
}

func (CrossTalkRequestStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3b7442485ea34150, []int{0}
}

func init() {
	proto.RegisterEnum("routerprotocol.routerchain.crosstalk.CrossTalkRequestStatus", CrossTalkRequestStatus_name, CrossTalkRequestStatus_value)
}

func init() {
	proto.RegisterFile("crosstalk/crosstalk_request_status.proto", fileDescriptor_3b7442485ea34150)
}

var fileDescriptor_3b7442485ea34150 = []byte{
	// 267 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x48, 0x2e, 0xca, 0x2f,
	0x2e, 0x2e, 0x49, 0xcc, 0xc9, 0xd6, 0x87, 0xb3, 0xe2, 0x8b, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b,
	0xe2, 0x8b, 0x4b, 0x12, 0x4b, 0x4a, 0x8b, 0xf5, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x54, 0x8a,
	0xf2, 0x4b, 0x4b, 0x52, 0x8b, 0xc0, 0x9c, 0xe4, 0xfc, 0x1c, 0x3d, 0x08, 0x37, 0x39, 0x23, 0x31,
	0x33, 0x4f, 0x0f, 0xae, 0x55, 0x4a, 0x24, 0x3d, 0x3f, 0x3d, 0x1f, 0xac, 0x46, 0x1f, 0xc4, 0x82,
	0xe8, 0x95, 0x92, 0x4c, 0xcf, 0xcf, 0x4f, 0xcf, 0x49, 0xd5, 0x07, 0xf3, 0x92, 0x4a, 0xd3, 0xf4,
	0x13, 0xf3, 0x2a, 0x21, 0x52, 0x5a, 0x73, 0x18, 0xb9, 0xc4, 0x9c, 0x41, 0xda, 0x43, 0x12, 0x73,
	0xb2, 0x83, 0x20, 0x16, 0x07, 0x83, 0xed, 0x15, 0xb2, 0xe1, 0x92, 0x74, 0x0e, 0xf2, 0x0f, 0x0e,
	0x0e, 0x71, 0xf4, 0xf1, 0x8e, 0x0f, 0x72, 0x0d, 0x0c, 0x75, 0x0d, 0x0e, 0x89, 0x77, 0x0e, 0x72,
	0x75, 0x0c, 0x71, 0x75, 0x11, 0x60, 0x90, 0x92, 0xed, 0x9a, 0xab, 0x80, 0x5b, 0x81, 0x90, 0x1d,
	0x97, 0x14, 0xa6, 0xa4, 0xbf, 0x53, 0xb0, 0x6b, 0x50, 0x98, 0xab, 0x8b, 0x00, 0xa3, 0x94, 0x5c,
	0xd7, 0x5c, 0x05, 0x3c, 0x2a, 0xa4, 0x58, 0x3a, 0x16, 0xcb, 0x31, 0x38, 0x05, 0x9f, 0x78, 0x24,
	0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x13, 0x1e, 0xcb, 0x31, 0x5c, 0x78,
	0x2c, 0xc7, 0x70, 0xe3, 0xb1, 0x1c, 0x43, 0x94, 0x65, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92, 0x5e,
	0x72, 0x7e, 0xae, 0x3e, 0x24, 0x2c, 0x74, 0x61, 0x61, 0x03, 0xe3, 0x83, 0x03, 0x47, 0xbf, 0x02,
	0x11, 0xb2, 0xfa, 0x25, 0x95, 0x05, 0xa9, 0xc5, 0x49, 0x6c, 0x60, 0x85, 0xc6, 0x80, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x88, 0x25, 0xc5, 0xfa, 0x7d, 0x01, 0x00, 0x00,
}
