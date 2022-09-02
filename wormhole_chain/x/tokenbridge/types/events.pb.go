// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tokenbridge/events.proto

package types

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	proto "github.com/gogo/protobuf/proto"
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

type EventChainRegistered struct {
	ChainID        uint32 `protobuf:"varint,1,opt,name=chainID,proto3" json:"chainID,omitempty"`
	EmitterAddress []byte `protobuf:"bytes,2,opt,name=emitterAddress,proto3" json:"emitterAddress,omitempty"`
}

func (m *EventChainRegistered) Reset()         { *m = EventChainRegistered{} }
func (m *EventChainRegistered) String() string { return proto.CompactTextString(m) }
func (*EventChainRegistered) ProtoMessage()    {}
func (*EventChainRegistered) Descriptor() ([]byte, []int) {
	return fileDescriptor_2327fbee5f2854e8, []int{0}
}
func (m *EventChainRegistered) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventChainRegistered) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventChainRegistered.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventChainRegistered) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventChainRegistered.Merge(m, src)
}
func (m *EventChainRegistered) XXX_Size() int {
	return m.Size()
}
func (m *EventChainRegistered) XXX_DiscardUnknown() {
	xxx_messageInfo_EventChainRegistered.DiscardUnknown(m)
}

var xxx_messageInfo_EventChainRegistered proto.InternalMessageInfo

func (m *EventChainRegistered) GetChainID() uint32 {
	if m != nil {
		return m.ChainID
	}
	return 0
}

func (m *EventChainRegistered) GetEmitterAddress() []byte {
	if m != nil {
		return m.EmitterAddress
	}
	return nil
}

type EventAssetRegistrationUpdate struct {
	TokenChain   uint32 `protobuf:"varint,1,opt,name=tokenChain,proto3" json:"tokenChain,omitempty"`
	TokenAddress []byte `protobuf:"bytes,2,opt,name=tokenAddress,proto3" json:"tokenAddress,omitempty"`
	Name         string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Symbol       string `protobuf:"bytes,4,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Decimals     uint32 `protobuf:"varint,5,opt,name=decimals,proto3" json:"decimals,omitempty"`
}

func (m *EventAssetRegistrationUpdate) Reset()         { *m = EventAssetRegistrationUpdate{} }
func (m *EventAssetRegistrationUpdate) String() string { return proto.CompactTextString(m) }
func (*EventAssetRegistrationUpdate) ProtoMessage()    {}
func (*EventAssetRegistrationUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_2327fbee5f2854e8, []int{1}
}
func (m *EventAssetRegistrationUpdate) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventAssetRegistrationUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventAssetRegistrationUpdate.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventAssetRegistrationUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventAssetRegistrationUpdate.Merge(m, src)
}
func (m *EventAssetRegistrationUpdate) XXX_Size() int {
	return m.Size()
}
func (m *EventAssetRegistrationUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_EventAssetRegistrationUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_EventAssetRegistrationUpdate proto.InternalMessageInfo

func (m *EventAssetRegistrationUpdate) GetTokenChain() uint32 {
	if m != nil {
		return m.TokenChain
	}
	return 0
}

func (m *EventAssetRegistrationUpdate) GetTokenAddress() []byte {
	if m != nil {
		return m.TokenAddress
	}
	return nil
}

func (m *EventAssetRegistrationUpdate) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *EventAssetRegistrationUpdate) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *EventAssetRegistrationUpdate) GetDecimals() uint32 {
	if m != nil {
		return m.Decimals
	}
	return 0
}

type EventTransferReceived struct {
	TokenChain   uint32 `protobuf:"varint,1,opt,name=tokenChain,proto3" json:"tokenChain,omitempty"`
	TokenAddress []byte `protobuf:"bytes,2,opt,name=tokenAddress,proto3" json:"tokenAddress,omitempty"`
	To           string `protobuf:"bytes,3,opt,name=to,proto3" json:"to,omitempty"`
	FeeRecipient string `protobuf:"bytes,4,opt,name=feeRecipient,proto3" json:"feeRecipient,omitempty"`
	Amount       string `protobuf:"bytes,5,opt,name=amount,proto3" json:"amount,omitempty"`
	Fee          string `protobuf:"bytes,6,opt,name=fee,proto3" json:"fee,omitempty"`
	LocalDenom   string `protobuf:"bytes,7,opt,name=localDenom,proto3" json:"localDenom,omitempty"`
}

func (m *EventTransferReceived) Reset()         { *m = EventTransferReceived{} }
func (m *EventTransferReceived) String() string { return proto.CompactTextString(m) }
func (*EventTransferReceived) ProtoMessage()    {}
func (*EventTransferReceived) Descriptor() ([]byte, []int) {
	return fileDescriptor_2327fbee5f2854e8, []int{2}
}
func (m *EventTransferReceived) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventTransferReceived) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventTransferReceived.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventTransferReceived) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventTransferReceived.Merge(m, src)
}
func (m *EventTransferReceived) XXX_Size() int {
	return m.Size()
}
func (m *EventTransferReceived) XXX_DiscardUnknown() {
	xxx_messageInfo_EventTransferReceived.DiscardUnknown(m)
}

var xxx_messageInfo_EventTransferReceived proto.InternalMessageInfo

func (m *EventTransferReceived) GetTokenChain() uint32 {
	if m != nil {
		return m.TokenChain
	}
	return 0
}

func (m *EventTransferReceived) GetTokenAddress() []byte {
	if m != nil {
		return m.TokenAddress
	}
	return nil
}

func (m *EventTransferReceived) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *EventTransferReceived) GetFeeRecipient() string {
	if m != nil {
		return m.FeeRecipient
	}
	return ""
}

func (m *EventTransferReceived) GetAmount() string {
	if m != nil {
		return m.Amount
	}
	return ""
}

func (m *EventTransferReceived) GetFee() string {
	if m != nil {
		return m.Fee
	}
	return ""
}

func (m *EventTransferReceived) GetLocalDenom() string {
	if m != nil {
		return m.LocalDenom
	}
	return ""
}

func init() {
	proto.RegisterType((*EventChainRegistered)(nil), "certusone.wormholechain.tokenbridge.EventChainRegistered")
	proto.RegisterType((*EventAssetRegistrationUpdate)(nil), "certusone.wormholechain.tokenbridge.EventAssetRegistrationUpdate")
	proto.RegisterType((*EventTransferReceived)(nil), "certusone.wormholechain.tokenbridge.EventTransferReceived")
}

func init() { proto.RegisterFile("tokenbridge/events.proto", fileDescriptor_2327fbee5f2854e8) }

var fileDescriptor_2327fbee5f2854e8 = []byte{
	// 380 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0x41, 0x6f, 0xda, 0x30,
	0x18, 0x86, 0x31, 0x30, 0x18, 0x16, 0x43, 0x93, 0xb5, 0x4d, 0xd6, 0x34, 0x45, 0x28, 0x93, 0x26,
	0x2e, 0x4b, 0x0e, 0x3b, 0xec, 0xcc, 0xc6, 0x0e, 0x3b, 0x2e, 0x5a, 0xa5, 0xaa, 0x37, 0x27, 0xf9,
	0x00, 0xab, 0xb1, 0x1d, 0xd9, 0x1f, 0xb4, 0xfc, 0x8b, 0xfe, 0x8e, 0xfe, 0x92, 0x1e, 0x39, 0xb6,
	0xb7, 0x0a, 0xfe, 0x48, 0x15, 0x37, 0xa0, 0xd0, 0x73, 0x6f, 0x7e, 0x5f, 0xe7, 0xfb, 0xde, 0x27,
	0xf2, 0x4b, 0x39, 0x9a, 0x4b, 0xd0, 0xa9, 0x95, 0xf9, 0x02, 0x62, 0x58, 0x83, 0x46, 0x17, 0x95,
	0xd6, 0xa0, 0x61, 0x5f, 0x33, 0xb0, 0xb8, 0x72, 0x46, 0x43, 0x74, 0x65, 0xac, 0x5a, 0x9a, 0x02,
	0xb2, 0xa5, 0x90, 0x3a, 0x6a, 0x4c, 0x84, 0xe7, 0xf4, 0xc3, 0x9f, 0x6a, 0xe8, 0x77, 0x75, 0x93,
	0xc0, 0x42, 0x3a, 0x04, 0x0b, 0x39, 0xe3, 0xb4, 0xef, 0x3f, 0xfe, 0x3b, 0xe3, 0x64, 0x4c, 0x26,
	0xef, 0x92, 0x83, 0x64, 0xdf, 0xe8, 0x08, 0x94, 0x44, 0x04, 0x3b, 0xcd, 0x73, 0x0b, 0xce, 0xf1,
	0xf6, 0x98, 0x4c, 0x86, 0xc9, 0x0b, 0x37, 0xbc, 0x25, 0xf4, 0x8b, 0x5f, 0x3d, 0x75, 0x0e, 0xf0,
	0x79, 0xb5, 0x15, 0x28, 0x8d, 0x3e, 0x2b, 0x73, 0x81, 0xc0, 0x02, 0x4a, 0x3d, 0x89, 0x8f, 0xae,
	0x53, 0x1a, 0x0e, 0x0b, 0xe9, 0xd0, 0xab, 0xd3, 0x98, 0x13, 0x8f, 0x31, 0xda, 0xd5, 0x42, 0x01,
	0xef, 0x8c, 0xc9, 0x64, 0x90, 0xf8, 0x33, 0xfb, 0x44, 0x7b, 0x6e, 0xa3, 0x52, 0x53, 0xf0, 0xae,
	0x77, 0x6b, 0xc5, 0x3e, 0xd3, 0xb7, 0x39, 0x64, 0x52, 0x89, 0xc2, 0xf1, 0x37, 0x3e, 0xed, 0xa8,
	0xc3, 0x07, 0x42, 0x3f, 0x7a, 0xd8, 0xff, 0x56, 0x68, 0x37, 0x07, 0x9b, 0x40, 0x06, 0x72, 0x0d,
	0xf9, 0xab, 0x50, 0x8e, 0x68, 0x1b, 0x4d, 0xcd, 0xd8, 0x46, 0x53, 0xcd, 0xcc, 0x01, 0x12, 0xc8,
	0x64, 0x29, 0x41, 0x63, 0xcd, 0x79, 0xe2, 0x55, 0x7f, 0x21, 0x94, 0x59, 0x69, 0xf4, 0xac, 0x83,
	0xa4, 0x56, 0xec, 0x3d, 0xed, 0xcc, 0x01, 0x78, 0xcf, 0x9b, 0xd5, 0xb1, 0x22, 0x2c, 0x4c, 0x26,
	0x8a, 0x19, 0x68, 0xa3, 0x78, 0xdf, 0x5f, 0x34, 0x9c, 0x5f, 0xff, 0xee, 0x76, 0x01, 0xd9, 0xee,
	0x02, 0xf2, 0xb8, 0x0b, 0xc8, 0xcd, 0x3e, 0x68, 0x6d, 0xf7, 0x41, 0xeb, 0x7e, 0x1f, 0xb4, 0x2e,
	0x7e, 0x2e, 0x24, 0x2e, 0x57, 0x69, 0x94, 0x19, 0x15, 0x1f, 0xcb, 0x12, 0x1f, 0xca, 0xf2, 0xdd,
	0xbf, 0x78, 0x7c, 0x1d, 0x37, 0x1b, 0x86, 0x9b, 0x12, 0x5c, 0xda, 0xf3, 0x0d, 0xfb, 0xf1, 0x14,
	0x00, 0x00, 0xff, 0xff, 0xfe, 0xe7, 0xc2, 0xae, 0x7d, 0x02, 0x00, 0x00,
}

func (m *EventChainRegistered) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventChainRegistered) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventChainRegistered) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.EmitterAddress) > 0 {
		i -= len(m.EmitterAddress)
		copy(dAtA[i:], m.EmitterAddress)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.EmitterAddress)))
		i--
		dAtA[i] = 0x12
	}
	if m.ChainID != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.ChainID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *EventAssetRegistrationUpdate) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventAssetRegistrationUpdate) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventAssetRegistrationUpdate) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Decimals != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.Decimals))
		i--
		dAtA[i] = 0x28
	}
	if len(m.Symbol) > 0 {
		i -= len(m.Symbol)
		copy(dAtA[i:], m.Symbol)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Symbol)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.TokenAddress) > 0 {
		i -= len(m.TokenAddress)
		copy(dAtA[i:], m.TokenAddress)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.TokenAddress)))
		i--
		dAtA[i] = 0x12
	}
	if m.TokenChain != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.TokenChain))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *EventTransferReceived) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventTransferReceived) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventTransferReceived) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.LocalDenom) > 0 {
		i -= len(m.LocalDenom)
		copy(dAtA[i:], m.LocalDenom)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.LocalDenom)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.Fee) > 0 {
		i -= len(m.Fee)
		copy(dAtA[i:], m.Fee)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Fee)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Amount) > 0 {
		i -= len(m.Amount)
		copy(dAtA[i:], m.Amount)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Amount)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.FeeRecipient) > 0 {
		i -= len(m.FeeRecipient)
		copy(dAtA[i:], m.FeeRecipient)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.FeeRecipient)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.To) > 0 {
		i -= len(m.To)
		copy(dAtA[i:], m.To)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.To)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.TokenAddress) > 0 {
		i -= len(m.TokenAddress)
		copy(dAtA[i:], m.TokenAddress)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.TokenAddress)))
		i--
		dAtA[i] = 0x12
	}
	if m.TokenChain != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.TokenChain))
		i--
		dAtA[i] = 0x8
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
func (m *EventChainRegistered) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ChainID != 0 {
		n += 1 + sovEvents(uint64(m.ChainID))
	}
	l = len(m.EmitterAddress)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *EventAssetRegistrationUpdate) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.TokenChain != 0 {
		n += 1 + sovEvents(uint64(m.TokenChain))
	}
	l = len(m.TokenAddress)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.Symbol)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	if m.Decimals != 0 {
		n += 1 + sovEvents(uint64(m.Decimals))
	}
	return n
}

func (m *EventTransferReceived) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.TokenChain != 0 {
		n += 1 + sovEvents(uint64(m.TokenChain))
	}
	l = len(m.TokenAddress)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.To)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.FeeRecipient)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.Amount)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.Fee)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.LocalDenom)
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
func (m *EventChainRegistered) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: EventChainRegistered: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventChainRegistered: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainID", wireType)
			}
			m.ChainID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ChainID |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EmitterAddress", wireType)
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
			m.EmitterAddress = append(m.EmitterAddress[:0], dAtA[iNdEx:postIndex]...)
			if m.EmitterAddress == nil {
				m.EmitterAddress = []byte{}
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
func (m *EventAssetRegistrationUpdate) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: EventAssetRegistrationUpdate: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventAssetRegistrationUpdate: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenChain", wireType)
			}
			m.TokenChain = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TokenChain |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenAddress", wireType)
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
			m.TokenAddress = append(m.TokenAddress[:0], dAtA[iNdEx:postIndex]...)
			if m.TokenAddress == nil {
				m.TokenAddress = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
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
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Symbol", wireType)
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
			m.Symbol = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Decimals", wireType)
			}
			m.Decimals = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Decimals |= uint32(b&0x7F) << shift
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
func (m *EventTransferReceived) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: EventTransferReceived: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventTransferReceived: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenChain", wireType)
			}
			m.TokenChain = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TokenChain |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenAddress", wireType)
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
			m.TokenAddress = append(m.TokenAddress[:0], dAtA[iNdEx:postIndex]...)
			if m.TokenAddress == nil {
				m.TokenAddress = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field To", wireType)
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
			m.To = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeeRecipient", wireType)
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
			m.FeeRecipient = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
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
			m.Amount = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fee", wireType)
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
			m.Fee = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LocalDenom", wireType)
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
			m.LocalDenom = string(dAtA[iNdEx:postIndex])
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
