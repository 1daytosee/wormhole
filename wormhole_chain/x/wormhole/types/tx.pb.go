// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: wormhole/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type MsgExecuteGovernanceVAA struct {
	Vaa    []byte `protobuf:"bytes,1,opt,name=vaa,proto3" json:"vaa,omitempty"`
	Signer string `protobuf:"bytes,2,opt,name=signer,proto3" json:"signer,omitempty"`
}

func (m *MsgExecuteGovernanceVAA) Reset()         { *m = MsgExecuteGovernanceVAA{} }
func (m *MsgExecuteGovernanceVAA) String() string { return proto.CompactTextString(m) }
func (*MsgExecuteGovernanceVAA) ProtoMessage()    {}
func (*MsgExecuteGovernanceVAA) Descriptor() ([]byte, []int) {
	return fileDescriptor_55f7aa067b0c517b, []int{0}
}
func (m *MsgExecuteGovernanceVAA) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgExecuteGovernanceVAA) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgExecuteGovernanceVAA.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgExecuteGovernanceVAA) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgExecuteGovernanceVAA.Merge(m, src)
}
func (m *MsgExecuteGovernanceVAA) XXX_Size() int {
	return m.Size()
}
func (m *MsgExecuteGovernanceVAA) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgExecuteGovernanceVAA.DiscardUnknown(m)
}

var xxx_messageInfo_MsgExecuteGovernanceVAA proto.InternalMessageInfo

func (m *MsgExecuteGovernanceVAA) GetVaa() []byte {
	if m != nil {
		return m.Vaa
	}
	return nil
}

func (m *MsgExecuteGovernanceVAA) GetSigner() string {
	if m != nil {
		return m.Signer
	}
	return ""
}

type MsgExecuteGovernanceVAAResponse struct {
}

func (m *MsgExecuteGovernanceVAAResponse) Reset()         { *m = MsgExecuteGovernanceVAAResponse{} }
func (m *MsgExecuteGovernanceVAAResponse) String() string { return proto.CompactTextString(m) }
func (*MsgExecuteGovernanceVAAResponse) ProtoMessage()    {}
func (*MsgExecuteGovernanceVAAResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_55f7aa067b0c517b, []int{1}
}
func (m *MsgExecuteGovernanceVAAResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgExecuteGovernanceVAAResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgExecuteGovernanceVAAResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgExecuteGovernanceVAAResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgExecuteGovernanceVAAResponse.Merge(m, src)
}
func (m *MsgExecuteGovernanceVAAResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgExecuteGovernanceVAAResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgExecuteGovernanceVAAResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgExecuteGovernanceVAAResponse proto.InternalMessageInfo

type MsgRegisterAccountAsGuardian struct {
	Signer         string       `protobuf:"bytes,1,opt,name=signer,proto3" json:"signer,omitempty"`
	GuardianPubkey *GuardianKey `protobuf:"bytes,2,opt,name=guardianPubkey,proto3" json:"guardianPubkey,omitempty"`
	Signature      []byte       `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *MsgRegisterAccountAsGuardian) Reset()         { *m = MsgRegisterAccountAsGuardian{} }
func (m *MsgRegisterAccountAsGuardian) String() string { return proto.CompactTextString(m) }
func (*MsgRegisterAccountAsGuardian) ProtoMessage()    {}
func (*MsgRegisterAccountAsGuardian) Descriptor() ([]byte, []int) {
	return fileDescriptor_55f7aa067b0c517b, []int{2}
}
func (m *MsgRegisterAccountAsGuardian) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgRegisterAccountAsGuardian) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRegisterAccountAsGuardian.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgRegisterAccountAsGuardian) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRegisterAccountAsGuardian.Merge(m, src)
}
func (m *MsgRegisterAccountAsGuardian) XXX_Size() int {
	return m.Size()
}
func (m *MsgRegisterAccountAsGuardian) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRegisterAccountAsGuardian.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRegisterAccountAsGuardian proto.InternalMessageInfo

func (m *MsgRegisterAccountAsGuardian) GetSigner() string {
	if m != nil {
		return m.Signer
	}
	return ""
}

func (m *MsgRegisterAccountAsGuardian) GetGuardianPubkey() *GuardianKey {
	if m != nil {
		return m.GuardianPubkey
	}
	return nil
}

func (m *MsgRegisterAccountAsGuardian) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

type MsgRegisterAccountAsGuardianResponse struct {
}

func (m *MsgRegisterAccountAsGuardianResponse) Reset()         { *m = MsgRegisterAccountAsGuardianResponse{} }
func (m *MsgRegisterAccountAsGuardianResponse) String() string { return proto.CompactTextString(m) }
func (*MsgRegisterAccountAsGuardianResponse) ProtoMessage()    {}
func (*MsgRegisterAccountAsGuardianResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_55f7aa067b0c517b, []int{3}
}
func (m *MsgRegisterAccountAsGuardianResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgRegisterAccountAsGuardianResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRegisterAccountAsGuardianResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgRegisterAccountAsGuardianResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRegisterAccountAsGuardianResponse.Merge(m, src)
}
func (m *MsgRegisterAccountAsGuardianResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgRegisterAccountAsGuardianResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRegisterAccountAsGuardianResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRegisterAccountAsGuardianResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgExecuteGovernanceVAA)(nil), "certusone.wormholechain.wormhole.MsgExecuteGovernanceVAA")
	proto.RegisterType((*MsgExecuteGovernanceVAAResponse)(nil), "certusone.wormholechain.wormhole.MsgExecuteGovernanceVAAResponse")
	proto.RegisterType((*MsgRegisterAccountAsGuardian)(nil), "certusone.wormholechain.wormhole.MsgRegisterAccountAsGuardian")
	proto.RegisterType((*MsgRegisterAccountAsGuardianResponse)(nil), "certusone.wormholechain.wormhole.MsgRegisterAccountAsGuardianResponse")
}

func init() { proto.RegisterFile("wormhole/tx.proto", fileDescriptor_55f7aa067b0c517b) }

var fileDescriptor_55f7aa067b0c517b = []byte{
	// 360 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2c, 0xcf, 0x2f, 0xca,
	0xcd, 0xc8, 0xcf, 0x49, 0xd5, 0x2f, 0xa9, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x52, 0x48,
	0x4e, 0x2d, 0x2a, 0x29, 0x2d, 0xce, 0xcf, 0x4b, 0xd5, 0x83, 0x49, 0x26, 0x67, 0x24, 0x66, 0xe6,
	0xc1, 0x79, 0x52, 0xd2, 0x70, 0x4d, 0xe9, 0xa5, 0x89, 0x45, 0x29, 0x99, 0x89, 0x79, 0xf1, 0xd9,
	0xa9, 0x95, 0x10, 0xed, 0x4a, 0xce, 0x5c, 0xe2, 0xbe, 0xc5, 0xe9, 0xae, 0x15, 0xa9, 0xc9, 0xa5,
	0x25, 0xa9, 0xee, 0xf9, 0x65, 0xa9, 0x45, 0x79, 0x89, 0x79, 0xc9, 0xa9, 0x61, 0x8e, 0x8e, 0x42,
	0x02, 0x5c, 0xcc, 0x65, 0x89, 0x89, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x3c, 0x41, 0x20, 0xa6, 0x90,
	0x18, 0x17, 0x5b, 0x71, 0x66, 0x7a, 0x5e, 0x6a, 0x91, 0x04, 0x93, 0x02, 0xa3, 0x06, 0x67, 0x10,
	0x94, 0xa7, 0xa4, 0xc8, 0x25, 0x8f, 0xc3, 0x90, 0xa0, 0xd4, 0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54,
	0xa5, 0xd5, 0x8c, 0x5c, 0x32, 0xbe, 0xc5, 0xe9, 0x41, 0xa9, 0xe9, 0x99, 0xc5, 0x25, 0xa9, 0x45,
	0x8e, 0xc9, 0xc9, 0xf9, 0xa5, 0x79, 0x25, 0x8e, 0xc5, 0xee, 0x50, 0x27, 0x21, 0x99, 0xcd, 0x88,
	0x6c, 0xb6, 0x50, 0x28, 0x17, 0x1f, 0xcc, 0xd9, 0x01, 0xa5, 0x49, 0xd9, 0xa9, 0x95, 0x60, 0xbb,
	0xb9, 0x8d, 0x74, 0xf5, 0x08, 0x79, 0x5c, 0x0f, 0x66, 0xb6, 0x77, 0x6a, 0x65, 0x10, 0x9a, 0x21,
	0x42, 0x32, 0x5c, 0x9c, 0x20, 0x0b, 0x12, 0x4b, 0x4a, 0x8b, 0x52, 0x25, 0x98, 0xc1, 0x5e, 0x44,
	0x08, 0x28, 0xa9, 0x71, 0xa9, 0xe0, 0x73, 0x2c, 0xcc, 0x57, 0x46, 0x87, 0x98, 0xb8, 0x98, 0x7d,
	0x8b, 0xd3, 0x85, 0xa6, 0x30, 0x72, 0x89, 0x60, 0x0d, 0x43, 0x4b, 0xc2, 0xae, 0xc4, 0x11, 0x72,
	0x52, 0x8e, 0x64, 0x6b, 0x85, 0x39, 0x4f, 0x68, 0x31, 0x23, 0x97, 0x24, 0xee, 0x10, 0xb7, 0x23,
	0xca, 0x02, 0x9c, 0xfa, 0xa5, 0xdc, 0x28, 0xd3, 0x0f, 0x73, 0xa5, 0x93, 0xdf, 0x89, 0x47, 0x72,
	0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0xc3, 0x85, 0xc7,
	0x72, 0x0c, 0x37, 0x1e, 0xcb, 0x31, 0x44, 0x99, 0xa4, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25,
	0xe7, 0xe7, 0xea, 0xc3, 0xed, 0xd2, 0x87, 0x99, 0xae, 0x0b, 0xb6, 0x4c, 0xbf, 0x42, 0x1f, 0x91,
	0x29, 0x2a, 0x0b, 0x52, 0x8b, 0x93, 0xd8, 0xc0, 0x29, 0xdb, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff,
	0xd2, 0x9e, 0x5e, 0x5d, 0x2d, 0x03, 0x00, 0x00,
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
	ExecuteGovernanceVAA(ctx context.Context, in *MsgExecuteGovernanceVAA, opts ...grpc.CallOption) (*MsgExecuteGovernanceVAAResponse, error)
	RegisterAccountAsGuardian(ctx context.Context, in *MsgRegisterAccountAsGuardian, opts ...grpc.CallOption) (*MsgRegisterAccountAsGuardianResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) ExecuteGovernanceVAA(ctx context.Context, in *MsgExecuteGovernanceVAA, opts ...grpc.CallOption) (*MsgExecuteGovernanceVAAResponse, error) {
	out := new(MsgExecuteGovernanceVAAResponse)
	err := c.cc.Invoke(ctx, "/certusone.wormholechain.wormhole.Msg/ExecuteGovernanceVAA", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) RegisterAccountAsGuardian(ctx context.Context, in *MsgRegisterAccountAsGuardian, opts ...grpc.CallOption) (*MsgRegisterAccountAsGuardianResponse, error) {
	out := new(MsgRegisterAccountAsGuardianResponse)
	err := c.cc.Invoke(ctx, "/certusone.wormholechain.wormhole.Msg/RegisterAccountAsGuardian", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	ExecuteGovernanceVAA(context.Context, *MsgExecuteGovernanceVAA) (*MsgExecuteGovernanceVAAResponse, error)
	RegisterAccountAsGuardian(context.Context, *MsgRegisterAccountAsGuardian) (*MsgRegisterAccountAsGuardianResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) ExecuteGovernanceVAA(ctx context.Context, req *MsgExecuteGovernanceVAA) (*MsgExecuteGovernanceVAAResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExecuteGovernanceVAA not implemented")
}
func (*UnimplementedMsgServer) RegisterAccountAsGuardian(ctx context.Context, req *MsgRegisterAccountAsGuardian) (*MsgRegisterAccountAsGuardianResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterAccountAsGuardian not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_ExecuteGovernanceVAA_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgExecuteGovernanceVAA)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).ExecuteGovernanceVAA(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/certusone.wormholechain.wormhole.Msg/ExecuteGovernanceVAA",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).ExecuteGovernanceVAA(ctx, req.(*MsgExecuteGovernanceVAA))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_RegisterAccountAsGuardian_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRegisterAccountAsGuardian)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).RegisterAccountAsGuardian(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/certusone.wormholechain.wormhole.Msg/RegisterAccountAsGuardian",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).RegisterAccountAsGuardian(ctx, req.(*MsgRegisterAccountAsGuardian))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "certusone.wormholechain.wormhole.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ExecuteGovernanceVAA",
			Handler:    _Msg_ExecuteGovernanceVAA_Handler,
		},
		{
			MethodName: "RegisterAccountAsGuardian",
			Handler:    _Msg_RegisterAccountAsGuardian_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "wormhole/tx.proto",
}

func (m *MsgExecuteGovernanceVAA) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgExecuteGovernanceVAA) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgExecuteGovernanceVAA) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Signer) > 0 {
		i -= len(m.Signer)
		copy(dAtA[i:], m.Signer)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Signer)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Vaa) > 0 {
		i -= len(m.Vaa)
		copy(dAtA[i:], m.Vaa)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Vaa)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgExecuteGovernanceVAAResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgExecuteGovernanceVAAResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgExecuteGovernanceVAAResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgRegisterAccountAsGuardian) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRegisterAccountAsGuardian) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRegisterAccountAsGuardian) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Signature) > 0 {
		i -= len(m.Signature)
		copy(dAtA[i:], m.Signature)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Signature)))
		i--
		dAtA[i] = 0x1a
	}
	if m.GuardianPubkey != nil {
		{
			size, err := m.GuardianPubkey.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Signer) > 0 {
		i -= len(m.Signer)
		copy(dAtA[i:], m.Signer)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Signer)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgRegisterAccountAsGuardianResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRegisterAccountAsGuardianResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRegisterAccountAsGuardianResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
func (m *MsgExecuteGovernanceVAA) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Vaa)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Signer)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgExecuteGovernanceVAAResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgRegisterAccountAsGuardian) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Signer)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.GuardianPubkey != nil {
		l = m.GuardianPubkey.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Signature)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgRegisterAccountAsGuardianResponse) Size() (n int) {
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
func (m *MsgExecuteGovernanceVAA) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgExecuteGovernanceVAA: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgExecuteGovernanceVAA: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Vaa", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Vaa = append(m.Vaa[:0], dAtA[iNdEx:postIndex]...)
			if m.Vaa == nil {
				m.Vaa = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signer", wireType)
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
			m.Signer = string(dAtA[iNdEx:postIndex])
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
func (m *MsgExecuteGovernanceVAAResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgExecuteGovernanceVAAResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgExecuteGovernanceVAAResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *MsgRegisterAccountAsGuardian) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgRegisterAccountAsGuardian: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRegisterAccountAsGuardian: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signer", wireType)
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
			m.Signer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GuardianPubkey", wireType)
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
			if m.GuardianPubkey == nil {
				m.GuardianPubkey = &GuardianKey{}
			}
			if err := m.GuardianPubkey.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signature", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signature = append(m.Signature[:0], dAtA[iNdEx:postIndex]...)
			if m.Signature == nil {
				m.Signature = []byte{}
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
func (m *MsgRegisterAccountAsGuardianResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgRegisterAccountAsGuardianResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRegisterAccountAsGuardianResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
