// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: slinky/module/v1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type GetAllCurrencyPairsRequest struct {
}

func (m *GetAllCurrencyPairsRequest) Reset()         { *m = GetAllCurrencyPairsRequest{} }
func (m *GetAllCurrencyPairsRequest) String() string { return proto.CompactTextString(m) }
func (*GetAllCurrencyPairsRequest) ProtoMessage()    {}
func (*GetAllCurrencyPairsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c7eac55a0ef62584, []int{0}
}
func (m *GetAllCurrencyPairsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetAllCurrencyPairsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetAllCurrencyPairsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetAllCurrencyPairsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllCurrencyPairsRequest.Merge(m, src)
}
func (m *GetAllCurrencyPairsRequest) XXX_Size() int {
	return m.Size()
}
func (m *GetAllCurrencyPairsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllCurrencyPairsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllCurrencyPairsRequest proto.InternalMessageInfo

// GetAllCurrencyPairsResponse returns all CurrencyPairs that the module is
// currently tracking
type GetAllCurrencyPairsResponse struct {
	CurrencyPairs []CurrencyPair `protobuf:"bytes,1,rep,name=currency_pairs,json=currencyPairs,proto3" json:"currency_pairs"`
}

func (m *GetAllCurrencyPairsResponse) Reset()         { *m = GetAllCurrencyPairsResponse{} }
func (m *GetAllCurrencyPairsResponse) String() string { return proto.CompactTextString(m) }
func (*GetAllCurrencyPairsResponse) ProtoMessage()    {}
func (*GetAllCurrencyPairsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c7eac55a0ef62584, []int{1}
}
func (m *GetAllCurrencyPairsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetAllCurrencyPairsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetAllCurrencyPairsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetAllCurrencyPairsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllCurrencyPairsResponse.Merge(m, src)
}
func (m *GetAllCurrencyPairsResponse) XXX_Size() int {
	return m.Size()
}
func (m *GetAllCurrencyPairsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllCurrencyPairsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllCurrencyPairsResponse proto.InternalMessageInfo

func (m *GetAllCurrencyPairsResponse) GetCurrencyPairs() []CurrencyPair {
	if m != nil {
		return m.CurrencyPairs
	}
	return nil
}

// GetPriceRequest either takes a CurrencyPair, or an identifier for the
// CurrencyPair in the format base/quote
type GetPriceRequest struct {
	// CurrencyPairSelector represents the types that the user may provide to the
	// request to identify a CurrencyPair
	//
	// Types that are valid to be assigned to CurrencyPairSelector:
	//	*GetPriceRequest_CurrencyPair
	//	*GetPriceRequest_CurrencyPairId
	CurrencyPairSelector isGetPriceRequest_CurrencyPairSelector `protobuf_oneof:"currency_pair_selector"`
}

func (m *GetPriceRequest) Reset()         { *m = GetPriceRequest{} }
func (m *GetPriceRequest) String() string { return proto.CompactTextString(m) }
func (*GetPriceRequest) ProtoMessage()    {}
func (*GetPriceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c7eac55a0ef62584, []int{2}
}
func (m *GetPriceRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetPriceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetPriceRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetPriceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPriceRequest.Merge(m, src)
}
func (m *GetPriceRequest) XXX_Size() int {
	return m.Size()
}
func (m *GetPriceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPriceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetPriceRequest proto.InternalMessageInfo

type isGetPriceRequest_CurrencyPairSelector interface {
	isGetPriceRequest_CurrencyPairSelector()
	MarshalTo([]byte) (int, error)
	Size() int
}

type GetPriceRequest_CurrencyPair struct {
	CurrencyPair *CurrencyPair `protobuf:"bytes,1,opt,name=currency_pair,json=currencyPair,proto3,oneof" json:"currency_pair,omitempty"`
}
type GetPriceRequest_CurrencyPairId struct {
	CurrencyPairId string `protobuf:"bytes,2,opt,name=currency_pair_id,json=currencyPairId,proto3,oneof" json:"currency_pair_id,omitempty"`
}

func (*GetPriceRequest_CurrencyPair) isGetPriceRequest_CurrencyPairSelector()   {}
func (*GetPriceRequest_CurrencyPairId) isGetPriceRequest_CurrencyPairSelector() {}

func (m *GetPriceRequest) GetCurrencyPairSelector() isGetPriceRequest_CurrencyPairSelector {
	if m != nil {
		return m.CurrencyPairSelector
	}
	return nil
}

func (m *GetPriceRequest) GetCurrencyPair() *CurrencyPair {
	if x, ok := m.GetCurrencyPairSelector().(*GetPriceRequest_CurrencyPair); ok {
		return x.CurrencyPair
	}
	return nil
}

func (m *GetPriceRequest) GetCurrencyPairId() string {
	if x, ok := m.GetCurrencyPairSelector().(*GetPriceRequest_CurrencyPairId); ok {
		return x.CurrencyPairId
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*GetPriceRequest) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*GetPriceRequest_CurrencyPair)(nil),
		(*GetPriceRequest_CurrencyPairId)(nil),
	}
}

// GetPriceResponse is the response from the GetPrice grpc method exposed from
// the x/oracle query service
type GetPriceResponse struct {
	// QuotePrice represents the quote-price for the CurrencyPair given in
	// GetPriceRequest (possibly nil if no update has been made)
	Price *QuotePrice `protobuf:"bytes,1,opt,name=price,proto3" json:"price,omitempty"`
	// nonce represents the nonce for the CurrencyPair if it exists in state
	Nonce uint64 `protobuf:"varint,2,opt,name=nonce,proto3" json:"nonce,omitempty"`
}

func (m *GetPriceResponse) Reset()         { *m = GetPriceResponse{} }
func (m *GetPriceResponse) String() string { return proto.CompactTextString(m) }
func (*GetPriceResponse) ProtoMessage()    {}
func (*GetPriceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c7eac55a0ef62584, []int{3}
}
func (m *GetPriceResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetPriceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetPriceResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetPriceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPriceResponse.Merge(m, src)
}
func (m *GetPriceResponse) XXX_Size() int {
	return m.Size()
}
func (m *GetPriceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPriceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetPriceResponse proto.InternalMessageInfo

func (m *GetPriceResponse) GetPrice() *QuotePrice {
	if m != nil {
		return m.Price
	}
	return nil
}

func (m *GetPriceResponse) GetNonce() uint64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func init() {
	proto.RegisterType((*GetAllCurrencyPairsRequest)(nil), "slinky.module.v1.GetAllCurrencyPairsRequest")
	proto.RegisterType((*GetAllCurrencyPairsResponse)(nil), "slinky.module.v1.GetAllCurrencyPairsResponse")
	proto.RegisterType((*GetPriceRequest)(nil), "slinky.module.v1.GetPriceRequest")
	proto.RegisterType((*GetPriceResponse)(nil), "slinky.module.v1.GetPriceResponse")
}

func init() { proto.RegisterFile("slinky/module/v1/query.proto", fileDescriptor_c7eac55a0ef62584) }

var fileDescriptor_c7eac55a0ef62584 = []byte{
	// 451 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x3f, 0x6f, 0xd3, 0x40,
	0x18, 0xc6, 0x7d, 0xa1, 0x41, 0x70, 0x25, 0x25, 0x3a, 0x2a, 0x14, 0xb9, 0xc1, 0xa4, 0xee, 0x40,
	0x8b, 0xa8, 0x4f, 0x2d, 0x4c, 0x6c, 0xa4, 0x43, 0xcb, 0xd6, 0x78, 0x83, 0x25, 0x72, 0x2f, 0xaf,
	0xcc, 0x29, 0x97, 0x3b, 0xf7, 0xee, 0x1c, 0x91, 0x95, 0x1d, 0x09, 0x89, 0x19, 0xf1, 0x75, 0x3a,
	0x56, 0x62, 0x61, 0x42, 0x28, 0xe1, 0x83, 0xa0, 0xd8, 0x17, 0xe1, 0x90, 0x80, 0xba, 0xd9, 0xfe,
	0x3d, 0x7e, 0x9e, 0xf7, 0x1f, 0x6e, 0x1b, 0xc1, 0xe5, 0x70, 0x42, 0x47, 0x6a, 0x90, 0x0b, 0xa0,
	0xe3, 0x23, 0x7a, 0x99, 0x83, 0x9e, 0x44, 0x99, 0x56, 0x56, 0x91, 0x66, 0x49, 0xa3, 0x92, 0x46,
	0xe3, 0x23, 0x7f, 0x3b, 0x55, 0xa9, 0x2a, 0x20, 0x9d, 0x3f, 0x95, 0x3a, 0xbf, 0x9d, 0x2a, 0x95,
	0x0a, 0xa0, 0x49, 0xc6, 0x69, 0x22, 0xa5, 0xb2, 0x89, 0xe5, 0x4a, 0x1a, 0x47, 0x83, 0x95, 0x8c,
	0x14, 0x24, 0x18, 0xee, 0x78, 0xd8, 0xc6, 0xfe, 0x29, 0xd8, 0x57, 0x42, 0x9c, 0xe4, 0x5a, 0x83,
	0x64, 0x93, 0xf3, 0x84, 0x6b, 0x13, 0xc3, 0x65, 0x0e, 0xc6, 0x86, 0x6f, 0xf0, 0xce, 0x5a, 0x6a,
	0x32, 0x25, 0x0d, 0x90, 0x97, 0x78, 0x8b, 0x39, 0xd0, 0xcf, 0xe6, 0xa4, 0x85, 0x3a, 0xb7, 0xf6,
	0x37, 0x8f, 0x1b, 0x51, 0x55, 0xdf, 0xdd, 0xb8, 0xfa, 0xf1, 0xd8, 0x8b, 0x1b, 0xac, 0xea, 0x11,
	0x7e, 0x44, 0xf8, 0xfe, 0x29, 0xd8, 0x73, 0xcd, 0x19, 0xb8, 0x38, 0xf2, 0x02, 0x37, 0x96, 0xfc,
	0x5a, 0xa8, 0x83, 0x56, 0xec, 0xce, 0xbc, 0xf8, 0x5e, 0xd5, 0x8a, 0x3c, 0xc5, 0xcd, 0xa5, 0xbf,
	0xfa, 0x7c, 0xd0, 0xaa, 0x75, 0xd0, 0xfe, 0xdd, 0x33, 0x2f, 0xde, 0xaa, 0x2a, 0x5f, 0x0f, 0xba,
	0x2d, 0xfc, 0x70, 0x59, 0x6b, 0x40, 0x00, 0xb3, 0x4a, 0x87, 0x3d, 0xdc, 0xfc, 0x53, 0x8e, 0xeb,
	0xef, 0x09, 0xae, 0x67, 0xf3, 0x0f, 0xae, 0x8e, 0xcd, 0xa8, 0x97, 0x2b, 0x0b, 0x85, 0xa6, 0x68,
	0x0a, 0xc5, 0x25, 0x27, 0xdb, 0xb8, 0x2e, 0x95, 0x64, 0x50, 0xe4, 0x6e, 0xc4, 0xe5, 0xcb, 0xf1,
	0xd7, 0x1a, 0xae, 0xf7, 0xe6, 0x1b, 0x25, 0x5f, 0x10, 0x7e, 0xb0, 0x66, 0x90, 0xe4, 0x59, 0xf4,
	0xf7, 0x92, 0xa3, 0x7f, 0x6f, 0xc3, 0x3f, 0xbc, 0xa1, 0xba, 0xac, 0x3e, 0x3c, 0xf8, 0xf0, 0xed,
	0xd7, 0xe7, 0xda, 0x1e, 0xd9, 0xa5, 0xee, 0x06, 0x94, 0x4e, 0xd8, 0xe2, 0x06, 0x6c, 0x3f, 0x11,
	0xa2, 0x6f, 0x39, 0x1b, 0x82, 0x36, 0x64, 0x8c, 0xef, 0x2c, 0x9a, 0x27, 0xbb, 0x6b, 0x53, 0xaa,
	0x7b, 0xf2, 0xc3, 0xff, 0x49, 0x5c, 0xfa, 0x5e, 0x91, 0xfe, 0x88, 0xec, 0xac, 0x4f, 0x2f, 0xe6,
	0xd6, 0x3d, 0xb9, 0x9a, 0x06, 0xe8, 0x7a, 0x1a, 0xa0, 0x9f, 0xd3, 0x00, 0x7d, 0x9a, 0x05, 0xde,
	0xf5, 0x2c, 0xf0, 0xbe, 0xcf, 0x02, 0xef, 0xed, 0x41, 0xca, 0xed, 0xbb, 0xfc, 0x22, 0x62, 0x6a,
	0x44, 0xcd, 0x90, 0x67, 0x87, 0x23, 0x18, 0x2f, 0x9c, 0xde, 0x2f, 0xbc, 0xec, 0x24, 0x03, 0x73,
	0x71, 0xbb, 0xb8, 0xe4, 0xe7, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x73, 0xdc, 0x46, 0x02, 0x4f,
	0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Get all the currency pairs the x/oracle module is tracking price-data for
	GetAllCurrencyPairs(ctx context.Context, in *GetAllCurrencyPairsRequest, opts ...grpc.CallOption) (*GetAllCurrencyPairsResponse, error)
	// Given a CurrencyPair (or its identifier) return the latest QuotePrice for
	// that CurrencyPair
	GetPrice(ctx context.Context, in *GetPriceRequest, opts ...grpc.CallOption) (*GetPriceResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) GetAllCurrencyPairs(ctx context.Context, in *GetAllCurrencyPairsRequest, opts ...grpc.CallOption) (*GetAllCurrencyPairsResponse, error) {
	out := new(GetAllCurrencyPairsResponse)
	err := c.cc.Invoke(ctx, "/slinky.module.v1.Query/GetAllCurrencyPairs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) GetPrice(ctx context.Context, in *GetPriceRequest, opts ...grpc.CallOption) (*GetPriceResponse, error) {
	out := new(GetPriceResponse)
	err := c.cc.Invoke(ctx, "/slinky.module.v1.Query/GetPrice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Get all the currency pairs the x/oracle module is tracking price-data for
	GetAllCurrencyPairs(context.Context, *GetAllCurrencyPairsRequest) (*GetAllCurrencyPairsResponse, error)
	// Given a CurrencyPair (or its identifier) return the latest QuotePrice for
	// that CurrencyPair
	GetPrice(context.Context, *GetPriceRequest) (*GetPriceResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) GetAllCurrencyPairs(ctx context.Context, req *GetAllCurrencyPairsRequest) (*GetAllCurrencyPairsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllCurrencyPairs not implemented")
}
func (*UnimplementedQueryServer) GetPrice(ctx context.Context, req *GetPriceRequest) (*GetPriceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPrice not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_GetAllCurrencyPairs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllCurrencyPairsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GetAllCurrencyPairs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/slinky.module.v1.Query/GetAllCurrencyPairs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GetAllCurrencyPairs(ctx, req.(*GetAllCurrencyPairsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_GetPrice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPriceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GetPrice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/slinky.module.v1.Query/GetPrice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GetPrice(ctx, req.(*GetPriceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "slinky.module.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllCurrencyPairs",
			Handler:    _Query_GetAllCurrencyPairs_Handler,
		},
		{
			MethodName: "GetPrice",
			Handler:    _Query_GetPrice_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "slinky/module/v1/query.proto",
}

func (m *GetAllCurrencyPairsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetAllCurrencyPairsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetAllCurrencyPairsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *GetAllCurrencyPairsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetAllCurrencyPairsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetAllCurrencyPairsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.CurrencyPairs) > 0 {
		for iNdEx := len(m.CurrencyPairs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.CurrencyPairs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *GetPriceRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetPriceRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetPriceRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.CurrencyPairSelector != nil {
		{
			size := m.CurrencyPairSelector.Size()
			i -= size
			if _, err := m.CurrencyPairSelector.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	return len(dAtA) - i, nil
}

func (m *GetPriceRequest_CurrencyPair) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetPriceRequest_CurrencyPair) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.CurrencyPair != nil {
		{
			size, err := m.CurrencyPair.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}
func (m *GetPriceRequest_CurrencyPairId) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetPriceRequest_CurrencyPairId) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	i -= len(m.CurrencyPairId)
	copy(dAtA[i:], m.CurrencyPairId)
	i = encodeVarintQuery(dAtA, i, uint64(len(m.CurrencyPairId)))
	i--
	dAtA[i] = 0x12
	return len(dAtA) - i, nil
}
func (m *GetPriceResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetPriceResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetPriceResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Nonce != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.Nonce))
		i--
		dAtA[i] = 0x10
	}
	if m.Price != nil {
		{
			size, err := m.Price.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GetAllCurrencyPairsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *GetAllCurrencyPairsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.CurrencyPairs) > 0 {
		for _, e := range m.CurrencyPairs {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func (m *GetPriceRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CurrencyPairSelector != nil {
		n += m.CurrencyPairSelector.Size()
	}
	return n
}

func (m *GetPriceRequest_CurrencyPair) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CurrencyPair != nil {
		l = m.CurrencyPair.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}
func (m *GetPriceRequest_CurrencyPairId) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.CurrencyPairId)
	n += 1 + l + sovQuery(uint64(l))
	return n
}
func (m *GetPriceResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Price != nil {
		l = m.Price.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	if m.Nonce != 0 {
		n += 1 + sovQuery(uint64(m.Nonce))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GetAllCurrencyPairsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: GetAllCurrencyPairsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetAllCurrencyPairsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *GetAllCurrencyPairsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: GetAllCurrencyPairsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetAllCurrencyPairsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrencyPairs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CurrencyPairs = append(m.CurrencyPairs, CurrencyPair{})
			if err := m.CurrencyPairs[len(m.CurrencyPairs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *GetPriceRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: GetPriceRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetPriceRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrencyPair", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &CurrencyPair{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.CurrencyPairSelector = &GetPriceRequest_CurrencyPair{v}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrencyPairId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CurrencyPairSelector = &GetPriceRequest_CurrencyPairId{string(dAtA[iNdEx:postIndex])}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *GetPriceResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: GetPriceResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetPriceResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Price", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Price == nil {
				m.Price = &QuotePrice{}
			}
			if err := m.Price.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nonce", wireType)
			}
			m.Nonce = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Nonce |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
