// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: osmosis/protorev/v1beta1/protorev.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
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

// TokenPairArbRoutes tracks all of the hot routes for a given pair of tokens
type TokenPairArbRoutes struct {
	// Stores all of the possible hot paths for a given pair of tokens
	ArbRoutes []*Route `protobuf:"bytes,1,rep,name=arb_routes,json=arbRoutes,proto3" json:"arb_routes,omitempty"`
	// Token denomination of the first asset
	TokenIn string `protobuf:"bytes,2,opt,name=token_in,json=tokenIn,proto3" json:"token_in,omitempty"`
	// Token denomination of the second asset
	TokenOut string `protobuf:"bytes,3,opt,name=token_out,json=tokenOut,proto3" json:"token_out,omitempty"`
}

func (m *TokenPairArbRoutes) Reset()         { *m = TokenPairArbRoutes{} }
func (m *TokenPairArbRoutes) String() string { return proto.CompactTextString(m) }
func (*TokenPairArbRoutes) ProtoMessage()    {}
func (*TokenPairArbRoutes) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e9f2391fd9fec01, []int{0}
}
func (m *TokenPairArbRoutes) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TokenPairArbRoutes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TokenPairArbRoutes.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TokenPairArbRoutes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TokenPairArbRoutes.Merge(m, src)
}
func (m *TokenPairArbRoutes) XXX_Size() int {
	return m.Size()
}
func (m *TokenPairArbRoutes) XXX_DiscardUnknown() {
	xxx_messageInfo_TokenPairArbRoutes.DiscardUnknown(m)
}

var xxx_messageInfo_TokenPairArbRoutes proto.InternalMessageInfo

func (m *TokenPairArbRoutes) GetArbRoutes() []*Route {
	if m != nil {
		return m.ArbRoutes
	}
	return nil
}

func (m *TokenPairArbRoutes) GetTokenIn() string {
	if m != nil {
		return m.TokenIn
	}
	return ""
}

func (m *TokenPairArbRoutes) GetTokenOut() string {
	if m != nil {
		return m.TokenOut
	}
	return ""
}

// Route is a hot route for a given pair of tokens
type Route struct {
	// The pool IDs that are travered in the directed cyclic graph (traversed left
	// -> right)
	Trades []*Trade `protobuf:"bytes,1,rep,name=trades,proto3" json:"trades,omitempty"`
}

func (m *Route) Reset()         { *m = Route{} }
func (m *Route) String() string { return proto.CompactTextString(m) }
func (*Route) ProtoMessage()    {}
func (*Route) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e9f2391fd9fec01, []int{1}
}
func (m *Route) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Route) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Route.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Route) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Route.Merge(m, src)
}
func (m *Route) XXX_Size() int {
	return m.Size()
}
func (m *Route) XXX_DiscardUnknown() {
	xxx_messageInfo_Route.DiscardUnknown(m)
}

var xxx_messageInfo_Route proto.InternalMessageInfo

func (m *Route) GetTrades() []*Trade {
	if m != nil {
		return m.Trades
	}
	return nil
}

// Trade is a single trade in a route
type Trade struct {
	// The pool IDs that are travered in the directed cyclic graph (traversed left
	// -> right)
	Pool uint64 `protobuf:"varint,1,opt,name=pool,proto3" json:"pool,omitempty"`
	// The denom of token A that is traded
	TokenIn string `protobuf:"bytes,2,opt,name=token_in,json=tokenIn,proto3" json:"token_in,omitempty"`
	// The denom of token B that is traded
	TokenOut string `protobuf:"bytes,3,opt,name=token_out,json=tokenOut,proto3" json:"token_out,omitempty"`
}

func (m *Trade) Reset()         { *m = Trade{} }
func (m *Trade) String() string { return proto.CompactTextString(m) }
func (*Trade) ProtoMessage()    {}
func (*Trade) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e9f2391fd9fec01, []int{2}
}
func (m *Trade) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Trade) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Trade.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Trade) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Trade.Merge(m, src)
}
func (m *Trade) XXX_Size() int {
	return m.Size()
}
func (m *Trade) XXX_DiscardUnknown() {
	xxx_messageInfo_Trade.DiscardUnknown(m)
}

var xxx_messageInfo_Trade proto.InternalMessageInfo

func (m *Trade) GetPool() uint64 {
	if m != nil {
		return m.Pool
	}
	return 0
}

func (m *Trade) GetTokenIn() string {
	if m != nil {
		return m.TokenIn
	}
	return ""
}

func (m *Trade) GetTokenOut() string {
	if m != nil {
		return m.TokenOut
	}
	return ""
}

// RouteStatistics contains the number of trades the module has executed after a
// swap on a given route and the profits from the trades
type RouteStatistics struct {
	// profits is the total profit from all trades on this route
	Profits []*types.Coin `protobuf:"bytes,1,rep,name=profits,proto3" json:"profits,omitempty"`
	// number_of_trades is the number of trades the module has executed using this
	// route
	NumberOfTrades github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,2,opt,name=number_of_trades,json=numberOfTrades,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"number_of_trades"`
	// route is the route that was used
	Route []uint64 `protobuf:"varint,3,rep,packed,name=route,proto3" json:"route,omitempty"`
}

func (m *RouteStatistics) Reset()         { *m = RouteStatistics{} }
func (m *RouteStatistics) String() string { return proto.CompactTextString(m) }
func (*RouteStatistics) ProtoMessage()    {}
func (*RouteStatistics) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e9f2391fd9fec01, []int{3}
}
func (m *RouteStatistics) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RouteStatistics) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RouteStatistics.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RouteStatistics) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteStatistics.Merge(m, src)
}
func (m *RouteStatistics) XXX_Size() int {
	return m.Size()
}
func (m *RouteStatistics) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteStatistics.DiscardUnknown(m)
}

var xxx_messageInfo_RouteStatistics proto.InternalMessageInfo

func (m *RouteStatistics) GetProfits() []*types.Coin {
	if m != nil {
		return m.Profits
	}
	return nil
}

func (m *RouteStatistics) GetRoute() []uint64 {
	if m != nil {
		return m.Route
	}
	return nil
}

// PoolWeights contains the weights of all of the different pool types. This
// distinction is made and necessary because the execution time ranges
// significantly between the different pool types. Each weight roughly
// corresponds to the amount of time (in ms) it takes to execute a swap on that
// pool type.
type PoolWeights struct {
	// The weight of a stableswap pool
	StableWeight uint64 `protobuf:"varint,1,opt,name=stable_weight,json=stableWeight,proto3" json:"stable_weight,omitempty"`
	// The weight of a balancer pool
	BalancerWeight uint64 `protobuf:"varint,2,opt,name=balancer_weight,json=balancerWeight,proto3" json:"balancer_weight,omitempty"`
	// The weight of a concentrated pool
	ConcentratedWeight uint64 `protobuf:"varint,3,opt,name=concentrated_weight,json=concentratedWeight,proto3" json:"concentrated_weight,omitempty"`
}

func (m *PoolWeights) Reset()         { *m = PoolWeights{} }
func (m *PoolWeights) String() string { return proto.CompactTextString(m) }
func (*PoolWeights) ProtoMessage()    {}
func (*PoolWeights) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e9f2391fd9fec01, []int{4}
}
func (m *PoolWeights) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PoolWeights) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PoolWeights.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PoolWeights) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PoolWeights.Merge(m, src)
}
func (m *PoolWeights) XXX_Size() int {
	return m.Size()
}
func (m *PoolWeights) XXX_DiscardUnknown() {
	xxx_messageInfo_PoolWeights.DiscardUnknown(m)
}

var xxx_messageInfo_PoolWeights proto.InternalMessageInfo

func (m *PoolWeights) GetStableWeight() uint64 {
	if m != nil {
		return m.StableWeight
	}
	return 0
}

func (m *PoolWeights) GetBalancerWeight() uint64 {
	if m != nil {
		return m.BalancerWeight
	}
	return 0
}

func (m *PoolWeights) GetConcentratedWeight() uint64 {
	if m != nil {
		return m.ConcentratedWeight
	}
	return 0
}

func init() {
	proto.RegisterType((*TokenPairArbRoutes)(nil), "osmosis.protorev.v1beta1.TokenPairArbRoutes")
	proto.RegisterType((*Route)(nil), "osmosis.protorev.v1beta1.Route")
	proto.RegisterType((*Trade)(nil), "osmosis.protorev.v1beta1.Trade")
	proto.RegisterType((*RouteStatistics)(nil), "osmosis.protorev.v1beta1.RouteStatistics")
	proto.RegisterType((*PoolWeights)(nil), "osmosis.protorev.v1beta1.PoolWeights")
}

func init() {
	proto.RegisterFile("osmosis/protorev/v1beta1/protorev.proto", fileDescriptor_1e9f2391fd9fec01)
}

var fileDescriptor_1e9f2391fd9fec01 = []byte{
	// 494 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x93, 0x4f, 0x6b, 0x13, 0x41,
	0x18, 0xc6, 0x33, 0xdd, 0xa4, 0x35, 0x53, 0x6d, 0x65, 0xec, 0x61, 0x53, 0x61, 0x13, 0x22, 0xd8,
	0x5c, 0xba, 0x4b, 0x6c, 0x41, 0xf0, 0x20, 0x58, 0x41, 0xe8, 0xc5, 0x96, 0xb5, 0xe0, 0x9f, 0xcb,
	0x32, 0xb3, 0x99, 0xa4, 0x43, 0x37, 0x33, 0x61, 0x66, 0x36, 0xea, 0x77, 0xf0, 0x20, 0x7e, 0x02,
	0x3f, 0x85, 0x9f, 0xa1, 0xc7, 0x1e, 0xc5, 0x43, 0x91, 0xe4, 0xe2, 0xc7, 0x90, 0xbc, 0x33, 0xbb,
	0xf6, 0x22, 0x08, 0x9e, 0x32, 0xef, 0xf3, 0xfc, 0xe6, 0xcd, 0x33, 0xef, 0xcb, 0xe2, 0x3d, 0x65,
	0xa6, 0xca, 0x08, 0x93, 0xcc, 0xb4, 0xb2, 0x4a, 0xf3, 0x79, 0x32, 0x1f, 0x32, 0x6e, 0xe9, 0xb0,
	0x16, 0x62, 0x38, 0x90, 0xd0, 0x83, 0x71, 0xad, 0x7b, 0x70, 0xb7, 0x93, 0x83, 0x95, 0x81, 0x91,
	0xb8, 0xc2, 0x51, 0xbb, 0x3b, 0x13, 0x35, 0x51, 0x4e, 0x5f, 0x9d, 0xbc, 0x1a, 0x39, 0x26, 0x61,
	0xd4, 0xf0, 0xfa, 0xef, 0x72, 0x25, 0xa4, 0xf3, 0xfb, 0x5f, 0x10, 0x26, 0x67, 0xea, 0x82, 0xcb,
	0x53, 0x2a, 0xf4, 0x33, 0xcd, 0x52, 0x55, 0x5a, 0x6e, 0xc8, 0x53, 0x8c, 0xa9, 0x66, 0x99, 0x86,
	0x2a, 0x44, 0xbd, 0x60, 0xb0, 0xf9, 0xa8, 0x1b, 0xff, 0x2d, 0x56, 0x0c, 0xb7, 0xd2, 0x36, 0xad,
	0xef, 0x77, 0xf0, 0x2d, 0xbb, 0xea, 0x9a, 0x09, 0x19, 0xae, 0xf5, 0xd0, 0xa0, 0x9d, 0x6e, 0x40,
	0x7d, 0x2c, 0xc9, 0x7d, 0xdc, 0x76, 0x96, 0x2a, 0x6d, 0x18, 0x80, 0xe7, 0xd8, 0x93, 0xd2, 0x3e,
	0x69, 0xfe, 0xfa, 0xda, 0x45, 0xfd, 0x17, 0xb8, 0x05, 0x7d, 0xc8, 0x63, 0xbc, 0x6e, 0x35, 0x1d,
	0xfd, 0x4b, 0x84, 0xb3, 0x15, 0x97, 0x7a, 0xdc, 0xf7, 0x79, 0x8b, 0x5b, 0x20, 0x13, 0x82, 0x9b,
	0x33, 0xa5, 0x8a, 0x10, 0xf5, 0xd0, 0xa0, 0x99, 0xc2, 0xf9, 0x3f, 0x23, 0x7e, 0x43, 0x78, 0x1b,
	0x32, 0xbe, 0xb2, 0xd4, 0x0a, 0x63, 0x45, 0x6e, 0xc8, 0x01, 0xde, 0x98, 0x69, 0x35, 0x16, 0xb6,
	0x8a, 0xdb, 0x89, 0xfd, 0x86, 0x56, 0xd3, 0xaf, 0x93, 0x3e, 0x57, 0x42, 0xa6, 0x15, 0x49, 0xde,
	0xe0, 0xbb, 0xb2, 0x9c, 0x32, 0xae, 0x33, 0x35, 0xce, 0xfc, 0x63, 0x21, 0xce, 0x51, 0x7c, 0x79,
	0xdd, 0x6d, 0xfc, 0xb8, 0xee, 0x3e, 0x9c, 0x08, 0x7b, 0x5e, 0xb2, 0x38, 0x57, 0x53, 0xbf, 0x71,
	0xff, 0xb3, 0x6f, 0x46, 0x17, 0x89, 0xfd, 0x38, 0xe3, 0x26, 0x3e, 0x96, 0x36, 0xdd, 0x72, 0x7d,
	0x4e, 0xc6, 0xf0, 0x66, 0x43, 0x76, 0x70, 0x0b, 0xf6, 0x17, 0x06, 0xbd, 0x60, 0xd0, 0x4c, 0x5d,
	0xd1, 0xff, 0x84, 0xf0, 0xe6, 0xa9, 0x52, 0xc5, 0x6b, 0x2e, 0x26, 0xe7, 0xd6, 0x90, 0x07, 0xf8,
	0x8e, 0xb1, 0x94, 0x15, 0x3c, 0x7b, 0x0f, 0x8a, 0x9f, 0xd1, 0x6d, 0x27, 0x3a, 0x8a, 0xec, 0xe1,
	0x6d, 0x46, 0x0b, 0x2a, 0x73, 0xae, 0x2b, 0x6c, 0x0d, 0xb0, 0xad, 0x4a, 0xf6, 0x60, 0x82, 0xef,
	0xe5, 0x4a, 0xe6, 0x5c, 0x5a, 0x4d, 0x2d, 0x1f, 0x55, 0x70, 0x00, 0x30, 0xb9, 0x69, 0xb9, 0x0b,
	0x47, 0x2f, 0x2f, 0x17, 0x11, 0xba, 0x5a, 0x44, 0xe8, 0xe7, 0x22, 0x42, 0x9f, 0x97, 0x51, 0xe3,
	0x6a, 0x19, 0x35, 0xbe, 0x2f, 0xa3, 0xc6, 0xbb, 0xc3, 0x1b, 0xcf, 0xf6, 0x5b, 0xdf, 0x2f, 0x28,
	0x33, 0x55, 0x91, 0xcc, 0x87, 0x87, 0xc9, 0x87, 0x3f, 0xdf, 0x12, 0x0c, 0x82, 0xad, 0x43, 0x7d,
	0xf0, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x5e, 0x2d, 0x41, 0x5f, 0x6c, 0x03, 0x00, 0x00,
}

func (this *TokenPairArbRoutes) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*TokenPairArbRoutes)
	if !ok {
		that2, ok := that.(TokenPairArbRoutes)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.ArbRoutes) != len(that1.ArbRoutes) {
		return false
	}
	for i := range this.ArbRoutes {
		if !this.ArbRoutes[i].Equal(that1.ArbRoutes[i]) {
			return false
		}
	}
	if this.TokenIn != that1.TokenIn {
		return false
	}
	if this.TokenOut != that1.TokenOut {
		return false
	}
	return true
}
func (this *Route) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Route)
	if !ok {
		that2, ok := that.(Route)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.Trades) != len(that1.Trades) {
		return false
	}
	for i := range this.Trades {
		if !this.Trades[i].Equal(that1.Trades[i]) {
			return false
		}
	}
	return true
}
func (this *Trade) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Trade)
	if !ok {
		that2, ok := that.(Trade)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Pool != that1.Pool {
		return false
	}
	if this.TokenIn != that1.TokenIn {
		return false
	}
	if this.TokenOut != that1.TokenOut {
		return false
	}
	return true
}
func (m *TokenPairArbRoutes) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TokenPairArbRoutes) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TokenPairArbRoutes) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TokenOut) > 0 {
		i -= len(m.TokenOut)
		copy(dAtA[i:], m.TokenOut)
		i = encodeVarintProtorev(dAtA, i, uint64(len(m.TokenOut)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.TokenIn) > 0 {
		i -= len(m.TokenIn)
		copy(dAtA[i:], m.TokenIn)
		i = encodeVarintProtorev(dAtA, i, uint64(len(m.TokenIn)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ArbRoutes) > 0 {
		for iNdEx := len(m.ArbRoutes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ArbRoutes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintProtorev(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *Route) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Route) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Route) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Trades) > 0 {
		for iNdEx := len(m.Trades) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Trades[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintProtorev(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *Trade) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Trade) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Trade) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TokenOut) > 0 {
		i -= len(m.TokenOut)
		copy(dAtA[i:], m.TokenOut)
		i = encodeVarintProtorev(dAtA, i, uint64(len(m.TokenOut)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.TokenIn) > 0 {
		i -= len(m.TokenIn)
		copy(dAtA[i:], m.TokenIn)
		i = encodeVarintProtorev(dAtA, i, uint64(len(m.TokenIn)))
		i--
		dAtA[i] = 0x12
	}
	if m.Pool != 0 {
		i = encodeVarintProtorev(dAtA, i, uint64(m.Pool))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *RouteStatistics) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RouteStatistics) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RouteStatistics) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Route) > 0 {
		dAtA2 := make([]byte, len(m.Route)*10)
		var j1 int
		for _, num := range m.Route {
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		i -= j1
		copy(dAtA[i:], dAtA2[:j1])
		i = encodeVarintProtorev(dAtA, i, uint64(j1))
		i--
		dAtA[i] = 0x1a
	}
	{
		size := m.NumberOfTrades.Size()
		i -= size
		if _, err := m.NumberOfTrades.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintProtorev(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Profits) > 0 {
		for iNdEx := len(m.Profits) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Profits[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintProtorev(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *PoolWeights) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PoolWeights) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PoolWeights) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ConcentratedWeight != 0 {
		i = encodeVarintProtorev(dAtA, i, uint64(m.ConcentratedWeight))
		i--
		dAtA[i] = 0x18
	}
	if m.BalancerWeight != 0 {
		i = encodeVarintProtorev(dAtA, i, uint64(m.BalancerWeight))
		i--
		dAtA[i] = 0x10
	}
	if m.StableWeight != 0 {
		i = encodeVarintProtorev(dAtA, i, uint64(m.StableWeight))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintProtorev(dAtA []byte, offset int, v uint64) int {
	offset -= sovProtorev(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TokenPairArbRoutes) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.ArbRoutes) > 0 {
		for _, e := range m.ArbRoutes {
			l = e.Size()
			n += 1 + l + sovProtorev(uint64(l))
		}
	}
	l = len(m.TokenIn)
	if l > 0 {
		n += 1 + l + sovProtorev(uint64(l))
	}
	l = len(m.TokenOut)
	if l > 0 {
		n += 1 + l + sovProtorev(uint64(l))
	}
	return n
}

func (m *Route) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Trades) > 0 {
		for _, e := range m.Trades {
			l = e.Size()
			n += 1 + l + sovProtorev(uint64(l))
		}
	}
	return n
}

func (m *Trade) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Pool != 0 {
		n += 1 + sovProtorev(uint64(m.Pool))
	}
	l = len(m.TokenIn)
	if l > 0 {
		n += 1 + l + sovProtorev(uint64(l))
	}
	l = len(m.TokenOut)
	if l > 0 {
		n += 1 + l + sovProtorev(uint64(l))
	}
	return n
}

func (m *RouteStatistics) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Profits) > 0 {
		for _, e := range m.Profits {
			l = e.Size()
			n += 1 + l + sovProtorev(uint64(l))
		}
	}
	l = m.NumberOfTrades.Size()
	n += 1 + l + sovProtorev(uint64(l))
	if len(m.Route) > 0 {
		l = 0
		for _, e := range m.Route {
			l += sovProtorev(uint64(e))
		}
		n += 1 + sovProtorev(uint64(l)) + l
	}
	return n
}

func (m *PoolWeights) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.StableWeight != 0 {
		n += 1 + sovProtorev(uint64(m.StableWeight))
	}
	if m.BalancerWeight != 0 {
		n += 1 + sovProtorev(uint64(m.BalancerWeight))
	}
	if m.ConcentratedWeight != 0 {
		n += 1 + sovProtorev(uint64(m.ConcentratedWeight))
	}
	return n
}

func sovProtorev(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozProtorev(x uint64) (n int) {
	return sovProtorev(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TokenPairArbRoutes) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProtorev
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
			return fmt.Errorf("proto: TokenPairArbRoutes: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TokenPairArbRoutes: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ArbRoutes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtorev
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
				return ErrInvalidLengthProtorev
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProtorev
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ArbRoutes = append(m.ArbRoutes, &Route{})
			if err := m.ArbRoutes[len(m.ArbRoutes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenIn", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtorev
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
				return ErrInvalidLengthProtorev
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProtorev
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TokenIn = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenOut", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtorev
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
				return ErrInvalidLengthProtorev
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProtorev
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TokenOut = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProtorev(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProtorev
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
func (m *Route) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProtorev
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
			return fmt.Errorf("proto: Route: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Route: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Trades", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtorev
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
				return ErrInvalidLengthProtorev
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProtorev
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Trades = append(m.Trades, &Trade{})
			if err := m.Trades[len(m.Trades)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProtorev(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProtorev
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
func (m *Trade) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProtorev
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
			return fmt.Errorf("proto: Trade: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Trade: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pool", wireType)
			}
			m.Pool = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtorev
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Pool |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenIn", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtorev
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
				return ErrInvalidLengthProtorev
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProtorev
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TokenIn = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenOut", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtorev
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
				return ErrInvalidLengthProtorev
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProtorev
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TokenOut = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProtorev(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProtorev
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
func (m *RouteStatistics) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProtorev
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
			return fmt.Errorf("proto: RouteStatistics: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RouteStatistics: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Profits", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtorev
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
				return ErrInvalidLengthProtorev
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProtorev
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Profits = append(m.Profits, &types.Coin{})
			if err := m.Profits[len(m.Profits)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NumberOfTrades", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtorev
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
				return ErrInvalidLengthProtorev
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProtorev
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.NumberOfTrades.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType == 0 {
				var v uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowProtorev
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Route = append(m.Route, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowProtorev
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthProtorev
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthProtorev
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.Route) == 0 {
					m.Route = make([]uint64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowProtorev
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Route = append(m.Route, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Route", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipProtorev(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProtorev
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
func (m *PoolWeights) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProtorev
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
			return fmt.Errorf("proto: PoolWeights: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PoolWeights: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StableWeight", wireType)
			}
			m.StableWeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtorev
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StableWeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BalancerWeight", wireType)
			}
			m.BalancerWeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtorev
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BalancerWeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConcentratedWeight", wireType)
			}
			m.ConcentratedWeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtorev
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ConcentratedWeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipProtorev(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProtorev
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
func skipProtorev(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProtorev
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
					return 0, ErrIntOverflowProtorev
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
					return 0, ErrIntOverflowProtorev
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
				return 0, ErrInvalidLengthProtorev
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupProtorev
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthProtorev
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthProtorev        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProtorev          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupProtorev = fmt.Errorf("proto: unexpected end of group")
)
