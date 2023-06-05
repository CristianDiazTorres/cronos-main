// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cronos/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// MsgConvertVouchers represents a message to convert ibc voucher coins to cronos evm coins.
type MsgConvertVouchers struct {
	Address string                                   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Coins   github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,2,rep,name=coins,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"coins"`
}

func (m *MsgConvertVouchers) Reset()         { *m = MsgConvertVouchers{} }
func (m *MsgConvertVouchers) String() string { return proto.CompactTextString(m) }
func (*MsgConvertVouchers) ProtoMessage()    {}
func (*MsgConvertVouchers) Descriptor() ([]byte, []int) {
	return fileDescriptor_28e09e4eabb18884, []int{0}
}
func (m *MsgConvertVouchers) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgConvertVouchers) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgConvertVouchers.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgConvertVouchers) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgConvertVouchers.Merge(m, src)
}
func (m *MsgConvertVouchers) XXX_Size() int {
	return m.Size()
}
func (m *MsgConvertVouchers) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgConvertVouchers.DiscardUnknown(m)
}

var xxx_messageInfo_MsgConvertVouchers proto.InternalMessageInfo

func (m *MsgConvertVouchers) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *MsgConvertVouchers) GetCoins() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Coins
	}
	return nil
}

// MsgTransferTokens represents a message to transfer cronos evm coins through ibc.
type MsgTransferTokens struct {
	From  string                                   `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To    string                                   `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	Coins github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,3,rep,name=coins,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"coins"`
}

func (m *MsgTransferTokens) Reset()         { *m = MsgTransferTokens{} }
func (m *MsgTransferTokens) String() string { return proto.CompactTextString(m) }
func (*MsgTransferTokens) ProtoMessage()    {}
func (*MsgTransferTokens) Descriptor() ([]byte, []int) {
	return fileDescriptor_28e09e4eabb18884, []int{1}
}
func (m *MsgTransferTokens) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgTransferTokens) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgTransferTokens.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgTransferTokens) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgTransferTokens.Merge(m, src)
}
func (m *MsgTransferTokens) XXX_Size() int {
	return m.Size()
}
func (m *MsgTransferTokens) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgTransferTokens.DiscardUnknown(m)
}

var xxx_messageInfo_MsgTransferTokens proto.InternalMessageInfo

func (m *MsgTransferTokens) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *MsgTransferTokens) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *MsgTransferTokens) GetCoins() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Coins
	}
	return nil
}

// MsgConvertVouchersResponse defines the ConvertVouchers response type.
type MsgConvertVouchersResponse struct {
}

func (m *MsgConvertVouchersResponse) Reset()         { *m = MsgConvertVouchersResponse{} }
func (m *MsgConvertVouchersResponse) String() string { return proto.CompactTextString(m) }
func (*MsgConvertVouchersResponse) ProtoMessage()    {}
func (*MsgConvertVouchersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_28e09e4eabb18884, []int{2}
}
func (m *MsgConvertVouchersResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgConvertVouchersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgConvertVouchersResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgConvertVouchersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgConvertVouchersResponse.Merge(m, src)
}
func (m *MsgConvertVouchersResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgConvertVouchersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgConvertVouchersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgConvertVouchersResponse proto.InternalMessageInfo

// MsgTransferTokensResponse defines the TransferTokens response type.
type MsgTransferTokensResponse struct {
}

func (m *MsgTransferTokensResponse) Reset()         { *m = MsgTransferTokensResponse{} }
func (m *MsgTransferTokensResponse) String() string { return proto.CompactTextString(m) }
func (*MsgTransferTokensResponse) ProtoMessage()    {}
func (*MsgTransferTokensResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_28e09e4eabb18884, []int{3}
}
func (m *MsgTransferTokensResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgTransferTokensResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgTransferTokensResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgTransferTokensResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgTransferTokensResponse.Merge(m, src)
}
func (m *MsgTransferTokensResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgTransferTokensResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgTransferTokensResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgTransferTokensResponse proto.InternalMessageInfo

// MsgUpdateTokenMapping defines the request type
type MsgUpdateTokenMapping struct {
	Sender   string `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	Denom    string `protobuf:"bytes,2,opt,name=denom,proto3" json:"denom,omitempty"`
	Contract string `protobuf:"bytes,3,opt,name=contract,proto3" json:"contract,omitempty"`
}

func (m *MsgUpdateTokenMapping) Reset()         { *m = MsgUpdateTokenMapping{} }
func (m *MsgUpdateTokenMapping) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateTokenMapping) ProtoMessage()    {}
func (*MsgUpdateTokenMapping) Descriptor() ([]byte, []int) {
	return fileDescriptor_28e09e4eabb18884, []int{4}
}
func (m *MsgUpdateTokenMapping) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateTokenMapping) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateTokenMapping.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateTokenMapping) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateTokenMapping.Merge(m, src)
}
func (m *MsgUpdateTokenMapping) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateTokenMapping) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateTokenMapping.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateTokenMapping proto.InternalMessageInfo

func (m *MsgUpdateTokenMapping) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *MsgUpdateTokenMapping) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

func (m *MsgUpdateTokenMapping) GetContract() string {
	if m != nil {
		return m.Contract
	}
	return ""
}

// MsgUpdateTokenMappingResponse defines the response type
type MsgUpdateTokenMappingResponse struct {
}

func (m *MsgUpdateTokenMappingResponse) Reset()         { *m = MsgUpdateTokenMappingResponse{} }
func (m *MsgUpdateTokenMappingResponse) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateTokenMappingResponse) ProtoMessage()    {}
func (*MsgUpdateTokenMappingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_28e09e4eabb18884, []int{5}
}
func (m *MsgUpdateTokenMappingResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateTokenMappingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateTokenMappingResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateTokenMappingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateTokenMappingResponse.Merge(m, src)
}
func (m *MsgUpdateTokenMappingResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateTokenMappingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateTokenMappingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateTokenMappingResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgConvertVouchers)(nil), "cronos.MsgConvertVouchers")
	proto.RegisterType((*MsgTransferTokens)(nil), "cronos.MsgTransferTokens")
	proto.RegisterType((*MsgConvertVouchersResponse)(nil), "cronos.MsgConvertVouchersResponse")
	proto.RegisterType((*MsgTransferTokensResponse)(nil), "cronos.MsgTransferTokensResponse")
	proto.RegisterType((*MsgUpdateTokenMapping)(nil), "cronos.MsgUpdateTokenMapping")
	proto.RegisterType((*MsgUpdateTokenMappingResponse)(nil), "cronos.MsgUpdateTokenMappingResponse")
}

func init() { proto.RegisterFile("cronos/tx.proto", fileDescriptor_28e09e4eabb18884) }

var fileDescriptor_28e09e4eabb18884 = []byte{
	// 449 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x93, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x86, 0x63, 0x87, 0x06, 0x18, 0xa4, 0x56, 0xac, 0x0a, 0x72, 0x0c, 0x75, 0x8a, 0x25, 0xa4,
	0x5c, 0x62, 0x93, 0xf2, 0x06, 0xed, 0x11, 0x19, 0x89, 0xa8, 0x70, 0xe8, 0x6d, 0x63, 0x6f, 0x37,
	0x56, 0x95, 0x1d, 0x6b, 0x67, 0x5b, 0xb5, 0x77, 0x1e, 0x00, 0xae, 0x3c, 0x02, 0x4f, 0xd2, 0x63,
	0x8f, 0x9c, 0x00, 0x25, 0x2f, 0x82, 0xb2, 0x6b, 0x87, 0xd0, 0x34, 0xdc, 0x7a, 0xda, 0x99, 0xf9,
	0x34, 0x33, 0xff, 0xec, 0xce, 0xc2, 0x4e, 0xae, 0x51, 0x21, 0xa5, 0xe6, 0x32, 0xa9, 0x34, 0x1a,
	0x64, 0x1d, 0x17, 0x08, 0x77, 0x25, 0x4a, 0xb4, 0xa1, 0x74, 0x61, 0x39, 0x1a, 0x46, 0x39, 0xd2,
	0x14, 0x29, 0x1d, 0x73, 0x12, 0xe9, 0xc5, 0x70, 0x2c, 0x0c, 0x1f, 0xa6, 0x39, 0x96, 0xca, 0xf1,
	0xf8, 0xab, 0x07, 0x2c, 0x23, 0x79, 0x84, 0xea, 0x42, 0x68, 0xf3, 0x09, 0xcf, 0xf3, 0x89, 0xd0,
	0xc4, 0x02, 0x78, 0xc8, 0x8b, 0x42, 0x0b, 0xa2, 0xc0, 0xdb, 0xf7, 0xfa, 0x8f, 0x47, 0x8d, 0xcb,
	0x38, 0x6c, 0x2d, 0xd2, 0x29, 0xf0, 0xf7, 0xdb, 0xfd, 0x27, 0x07, 0xdd, 0xc4, 0x35, 0x48, 0x16,
	0x0d, 0x92, 0xba, 0x41, 0x72, 0x84, 0xa5, 0x3a, 0x7c, 0x73, 0xfd, 0xb3, 0xd7, 0xfa, 0xfe, 0xab,
	0xd7, 0x97, 0xa5, 0x99, 0x9c, 0x8f, 0x93, 0x1c, 0xa7, 0x69, 0xad, 0xc6, 0x1d, 0x03, 0x2a, 0xce,
	0x52, 0x73, 0x55, 0x09, 0xb2, 0x09, 0x34, 0x72, 0x95, 0xe3, 0x6f, 0x1e, 0x3c, 0xcd, 0x48, 0x1e,
	0x6b, 0xae, 0xe8, 0x54, 0xe8, 0x63, 0x3c, 0x13, 0x8a, 0x18, 0x83, 0x07, 0xa7, 0x1a, 0xa7, 0xb5,
	0x1e, 0x6b, 0xb3, 0x6d, 0xf0, 0x0d, 0x06, 0xbe, 0x8d, 0xf8, 0x06, 0xff, 0x8a, 0x6b, 0xdf, 0x9b,
	0xb8, 0x97, 0x10, 0xae, 0xdf, 0xd7, 0x48, 0x50, 0x85, 0x8a, 0x44, 0xfc, 0x02, 0xba, 0x6b, 0xca,
	0x97, 0x90, 0xc3, 0xb3, 0x8c, 0xe4, 0xc7, 0xaa, 0xe0, 0x46, 0x58, 0x94, 0xf1, 0xaa, 0x2a, 0x95,
	0x64, 0xcf, 0xa1, 0x43, 0x42, 0x15, 0x42, 0xd7, 0xc3, 0xd5, 0x1e, 0xdb, 0x85, 0xad, 0x42, 0x28,
	0x9c, 0xd6, 0x13, 0x3a, 0x87, 0x85, 0xf0, 0x28, 0x47, 0x65, 0x34, 0xcf, 0x4d, 0xd0, 0xb6, 0x60,
	0xe9, 0xc7, 0x3d, 0xd8, 0xbb, 0xb3, 0x45, 0xa3, 0xe1, 0xe0, 0xb3, 0x0f, 0xed, 0x8c, 0x24, 0xfb,
	0x00, 0x3b, 0xb7, 0xdf, 0x3c, 0x4c, 0xdc, 0x26, 0x25, 0xeb, 0xf3, 0x85, 0xf1, 0x66, 0xd6, 0x94,
	0x66, 0xef, 0x61, 0xfb, 0xd6, 0x93, 0x75, 0x57, 0xb2, 0xfe, 0x45, 0xe1, 0xab, 0x8d, 0x68, 0x59,
	0xef, 0x04, 0xd8, 0x1d, 0x77, 0xb5, 0xb7, 0x92, 0xb8, 0x8e, 0xc3, 0xd7, 0xff, 0xc5, 0x4d, 0xed,
	0xc3, 0x77, 0xd7, 0xb3, 0xc8, 0xbb, 0x99, 0x45, 0xde, 0xef, 0x59, 0xe4, 0x7d, 0x99, 0x47, 0xad,
	0x9b, 0x79, 0xd4, 0xfa, 0x31, 0x8f, 0x5a, 0x27, 0xc3, 0xd5, 0x85, 0xd0, 0x57, 0x95, 0xc1, 0x01,
	0x6a, 0x39, 0xc8, 0x27, 0xbc, 0x54, 0x69, 0xfd, 0xf7, 0x2e, 0x1b, 0xc3, 0xee, 0xc7, 0xb8, 0x63,
	0xbf, 0xd2, 0xdb, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x45, 0x34, 0x86, 0x15, 0x9b, 0x03, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	// ConvertVouchers defines a method for converting ibc voucher to cronos evm coins.
	ConvertVouchers(ctx context.Context, in *MsgConvertVouchers, opts ...grpc.CallOption) (*MsgConvertVouchersResponse, error)
	// TransferTokens defines a method to transfer cronos evm coins to another chain through IBC
	TransferTokens(ctx context.Context, in *MsgTransferTokens, opts ...grpc.CallOption) (*MsgTransferTokensResponse, error)
	// UpdateTokenMapping defines a method to update token mapping
	UpdateTokenMapping(ctx context.Context, in *MsgUpdateTokenMapping, opts ...grpc.CallOption) (*MsgUpdateTokenMappingResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) ConvertVouchers(ctx context.Context, in *MsgConvertVouchers, opts ...grpc.CallOption) (*MsgConvertVouchersResponse, error) {
	out := new(MsgConvertVouchersResponse)
	err := c.cc.Invoke(ctx, "/cronos.Msg/ConvertVouchers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) TransferTokens(ctx context.Context, in *MsgTransferTokens, opts ...grpc.CallOption) (*MsgTransferTokensResponse, error) {
	out := new(MsgTransferTokensResponse)
	err := c.cc.Invoke(ctx, "/cronos.Msg/TransferTokens", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateTokenMapping(ctx context.Context, in *MsgUpdateTokenMapping, opts ...grpc.CallOption) (*MsgUpdateTokenMappingResponse, error) {
	out := new(MsgUpdateTokenMappingResponse)
	err := c.cc.Invoke(ctx, "/cronos.Msg/UpdateTokenMapping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// ConvertVouchers defines a method for converting ibc voucher to cronos evm coins.
	ConvertVouchers(context.Context, *MsgConvertVouchers) (*MsgConvertVouchersResponse, error)
	// TransferTokens defines a method to transfer cronos evm coins to another chain through IBC
	TransferTokens(context.Context, *MsgTransferTokens) (*MsgTransferTokensResponse, error)
	// UpdateTokenMapping defines a method to update token mapping
	UpdateTokenMapping(context.Context, *MsgUpdateTokenMapping) (*MsgUpdateTokenMappingResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) ConvertVouchers(ctx context.Context, req *MsgConvertVouchers) (*MsgConvertVouchersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConvertVouchers not implemented")
}
func (*UnimplementedMsgServer) TransferTokens(ctx context.Context, req *MsgTransferTokens) (*MsgTransferTokensResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TransferTokens not implemented")
}
func (*UnimplementedMsgServer) UpdateTokenMapping(ctx context.Context, req *MsgUpdateTokenMapping) (*MsgUpdateTokenMappingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTokenMapping not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_ConvertVouchers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgConvertVouchers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).ConvertVouchers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cronos.Msg/ConvertVouchers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).ConvertVouchers(ctx, req.(*MsgConvertVouchers))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_TransferTokens_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgTransferTokens)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).TransferTokens(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cronos.Msg/TransferTokens",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).TransferTokens(ctx, req.(*MsgTransferTokens))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateTokenMapping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateTokenMapping)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateTokenMapping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cronos.Msg/UpdateTokenMapping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateTokenMapping(ctx, req.(*MsgUpdateTokenMapping))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cronos.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ConvertVouchers",
			Handler:    _Msg_ConvertVouchers_Handler,
		},
		{
			MethodName: "TransferTokens",
			Handler:    _Msg_TransferTokens_Handler,
		},
		{
			MethodName: "UpdateTokenMapping",
			Handler:    _Msg_UpdateTokenMapping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cronos/tx.proto",
}

func (m *MsgConvertVouchers) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgConvertVouchers) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgConvertVouchers) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Coins) > 0 {
		for iNdEx := len(m.Coins) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Coins[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgTransferTokens) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgTransferTokens) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgTransferTokens) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Coins) > 0 {
		for iNdEx := len(m.Coins) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Coins[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.To) > 0 {
		i -= len(m.To)
		copy(dAtA[i:], m.To)
		i = encodeVarintTx(dAtA, i, uint64(len(m.To)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.From) > 0 {
		i -= len(m.From)
		copy(dAtA[i:], m.From)
		i = encodeVarintTx(dAtA, i, uint64(len(m.From)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgConvertVouchersResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgConvertVouchersResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgConvertVouchersResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgTransferTokensResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgTransferTokensResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgTransferTokensResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgUpdateTokenMapping) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateTokenMapping) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateTokenMapping) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Contract) > 0 {
		i -= len(m.Contract)
		copy(dAtA[i:], m.Contract)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Contract)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgUpdateTokenMappingResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateTokenMappingResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateTokenMappingResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgConvertVouchers) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if len(m.Coins) > 0 {
		for _, e := range m.Coins {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	return n
}

func (m *MsgTransferTokens) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.From)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.To)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if len(m.Coins) > 0 {
		for _, e := range m.Coins {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	return n
}

func (m *MsgConvertVouchersResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgTransferTokensResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgUpdateTokenMapping) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Contract)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgUpdateTokenMappingResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgConvertVouchers) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgConvertVouchers: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgConvertVouchers: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Coins", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Coins = append(m.Coins, types.Coin{})
			if err := m.Coins[len(m.Coins)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgTransferTokens) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgTransferTokens: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgTransferTokens: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.From = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field To", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.To = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Coins", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Coins = append(m.Coins, types.Coin{})
			if err := m.Coins[len(m.Coins)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgConvertVouchersResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgConvertVouchersResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgConvertVouchersResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgTransferTokensResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgTransferTokensResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgTransferTokensResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgUpdateTokenMapping) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgUpdateTokenMapping: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateTokenMapping: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Contract", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Contract = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgUpdateTokenMappingResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgUpdateTokenMappingResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateTokenMappingResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)