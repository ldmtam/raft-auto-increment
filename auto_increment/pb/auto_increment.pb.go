// Code generated by protoc-gen-go. DO NOT EDIT.
// source: auto_increment.proto

package auto_increment

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type GetOneRequest struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetOneRequest) Reset()         { *m = GetOneRequest{} }
func (m *GetOneRequest) String() string { return proto.CompactTextString(m) }
func (*GetOneRequest) ProtoMessage()    {}
func (*GetOneRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_59c5cff9d1a655ea, []int{0}
}

func (m *GetOneRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetOneRequest.Unmarshal(m, b)
}
func (m *GetOneRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetOneRequest.Marshal(b, m, deterministic)
}
func (m *GetOneRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetOneRequest.Merge(m, src)
}
func (m *GetOneRequest) XXX_Size() int {
	return xxx_messageInfo_GetOneRequest.Size(m)
}
func (m *GetOneRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetOneRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetOneRequest proto.InternalMessageInfo

func (m *GetOneRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type GetOneResponse struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                uint64   `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetOneResponse) Reset()         { *m = GetOneResponse{} }
func (m *GetOneResponse) String() string { return proto.CompactTextString(m) }
func (*GetOneResponse) ProtoMessage()    {}
func (*GetOneResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_59c5cff9d1a655ea, []int{1}
}

func (m *GetOneResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetOneResponse.Unmarshal(m, b)
}
func (m *GetOneResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetOneResponse.Marshal(b, m, deterministic)
}
func (m *GetOneResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetOneResponse.Merge(m, src)
}
func (m *GetOneResponse) XXX_Size() int {
	return xxx_messageInfo_GetOneResponse.Size(m)
}
func (m *GetOneResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetOneResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetOneResponse proto.InternalMessageInfo

func (m *GetOneResponse) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *GetOneResponse) GetValue() uint64 {
	if m != nil {
		return m.Value
	}
	return 0
}

type GetManyRequest struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Quantity             uint64   `protobuf:"varint,2,opt,name=quantity,proto3" json:"quantity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetManyRequest) Reset()         { *m = GetManyRequest{} }
func (m *GetManyRequest) String() string { return proto.CompactTextString(m) }
func (*GetManyRequest) ProtoMessage()    {}
func (*GetManyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_59c5cff9d1a655ea, []int{2}
}

func (m *GetManyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetManyRequest.Unmarshal(m, b)
}
func (m *GetManyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetManyRequest.Marshal(b, m, deterministic)
}
func (m *GetManyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetManyRequest.Merge(m, src)
}
func (m *GetManyRequest) XXX_Size() int {
	return xxx_messageInfo_GetManyRequest.Size(m)
}
func (m *GetManyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetManyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetManyRequest proto.InternalMessageInfo

func (m *GetManyRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *GetManyRequest) GetQuantity() uint64 {
	if m != nil {
		return m.Quantity
	}
	return 0
}

type GetManyResponse struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Values               []uint64 `protobuf:"varint,2,rep,packed,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetManyResponse) Reset()         { *m = GetManyResponse{} }
func (m *GetManyResponse) String() string { return proto.CompactTextString(m) }
func (*GetManyResponse) ProtoMessage()    {}
func (*GetManyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_59c5cff9d1a655ea, []int{3}
}

func (m *GetManyResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetManyResponse.Unmarshal(m, b)
}
func (m *GetManyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetManyResponse.Marshal(b, m, deterministic)
}
func (m *GetManyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetManyResponse.Merge(m, src)
}
func (m *GetManyResponse) XXX_Size() int {
	return xxx_messageInfo_GetManyResponse.Size(m)
}
func (m *GetManyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetManyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetManyResponse proto.InternalMessageInfo

func (m *GetManyResponse) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *GetManyResponse) GetValues() []uint64 {
	if m != nil {
		return m.Values
	}
	return nil
}

type GetLastInsertedRequest struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetLastInsertedRequest) Reset()         { *m = GetLastInsertedRequest{} }
func (m *GetLastInsertedRequest) String() string { return proto.CompactTextString(m) }
func (*GetLastInsertedRequest) ProtoMessage()    {}
func (*GetLastInsertedRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_59c5cff9d1a655ea, []int{4}
}

func (m *GetLastInsertedRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetLastInsertedRequest.Unmarshal(m, b)
}
func (m *GetLastInsertedRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetLastInsertedRequest.Marshal(b, m, deterministic)
}
func (m *GetLastInsertedRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetLastInsertedRequest.Merge(m, src)
}
func (m *GetLastInsertedRequest) XXX_Size() int {
	return xxx_messageInfo_GetLastInsertedRequest.Size(m)
}
func (m *GetLastInsertedRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetLastInsertedRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetLastInsertedRequest proto.InternalMessageInfo

func (m *GetLastInsertedRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type GetLastInsertedResponse struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                uint64   `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetLastInsertedResponse) Reset()         { *m = GetLastInsertedResponse{} }
func (m *GetLastInsertedResponse) String() string { return proto.CompactTextString(m) }
func (*GetLastInsertedResponse) ProtoMessage()    {}
func (*GetLastInsertedResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_59c5cff9d1a655ea, []int{5}
}

func (m *GetLastInsertedResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetLastInsertedResponse.Unmarshal(m, b)
}
func (m *GetLastInsertedResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetLastInsertedResponse.Marshal(b, m, deterministic)
}
func (m *GetLastInsertedResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetLastInsertedResponse.Merge(m, src)
}
func (m *GetLastInsertedResponse) XXX_Size() int {
	return xxx_messageInfo_GetLastInsertedResponse.Size(m)
}
func (m *GetLastInsertedResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetLastInsertedResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetLastInsertedResponse proto.InternalMessageInfo

func (m *GetLastInsertedResponse) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *GetLastInsertedResponse) GetValue() uint64 {
	if m != nil {
		return m.Value
	}
	return 0
}

type JoinRequest struct {
	NodeID               string   `protobuf:"bytes,1,opt,name=nodeID,proto3" json:"nodeID,omitempty"`
	NodeAddress          string   `protobuf:"bytes,2,opt,name=nodeAddress,proto3" json:"nodeAddress,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JoinRequest) Reset()         { *m = JoinRequest{} }
func (m *JoinRequest) String() string { return proto.CompactTextString(m) }
func (*JoinRequest) ProtoMessage()    {}
func (*JoinRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_59c5cff9d1a655ea, []int{6}
}

func (m *JoinRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JoinRequest.Unmarshal(m, b)
}
func (m *JoinRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JoinRequest.Marshal(b, m, deterministic)
}
func (m *JoinRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JoinRequest.Merge(m, src)
}
func (m *JoinRequest) XXX_Size() int {
	return xxx_messageInfo_JoinRequest.Size(m)
}
func (m *JoinRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_JoinRequest.DiscardUnknown(m)
}

var xxx_messageInfo_JoinRequest proto.InternalMessageInfo

func (m *JoinRequest) GetNodeID() string {
	if m != nil {
		return m.NodeID
	}
	return ""
}

func (m *JoinRequest) GetNodeAddress() string {
	if m != nil {
		return m.NodeAddress
	}
	return ""
}

type JoinResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JoinResponse) Reset()         { *m = JoinResponse{} }
func (m *JoinResponse) String() string { return proto.CompactTextString(m) }
func (*JoinResponse) ProtoMessage()    {}
func (*JoinResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_59c5cff9d1a655ea, []int{7}
}

func (m *JoinResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JoinResponse.Unmarshal(m, b)
}
func (m *JoinResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JoinResponse.Marshal(b, m, deterministic)
}
func (m *JoinResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JoinResponse.Merge(m, src)
}
func (m *JoinResponse) XXX_Size() int {
	return xxx_messageInfo_JoinResponse.Size(m)
}
func (m *JoinResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_JoinResponse.DiscardUnknown(m)
}

var xxx_messageInfo_JoinResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*GetOneRequest)(nil), "GetOneRequest")
	proto.RegisterType((*GetOneResponse)(nil), "GetOneResponse")
	proto.RegisterType((*GetManyRequest)(nil), "GetManyRequest")
	proto.RegisterType((*GetManyResponse)(nil), "GetManyResponse")
	proto.RegisterType((*GetLastInsertedRequest)(nil), "GetLastInsertedRequest")
	proto.RegisterType((*GetLastInsertedResponse)(nil), "GetLastInsertedResponse")
	proto.RegisterType((*JoinRequest)(nil), "JoinRequest")
	proto.RegisterType((*JoinResponse)(nil), "JoinResponse")
}

func init() { proto.RegisterFile("auto_increment.proto", fileDescriptor_59c5cff9d1a655ea) }

var fileDescriptor_59c5cff9d1a655ea = []byte{
	// 422 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0xd1, 0x8a, 0xd3, 0x40,
	0x14, 0x86, 0x49, 0xb6, 0x46, 0xf7, 0xec, 0xb6, 0xbb, 0x0c, 0xa5, 0x9b, 0x8d, 0x82, 0xd9, 0x91,
	0xc5, 0xb2, 0xd2, 0x09, 0xe8, 0x8d, 0x28, 0x08, 0x85, 0x85, 0x52, 0xa9, 0x08, 0xb9, 0xf5, 0x42,
	0x46, 0x73, 0x28, 0xb1, 0xed, 0x4c, 0x9b, 0x99, 0x08, 0xa1, 0xf4, 0xc6, 0x57, 0xf0, 0xd1, 0x7c,
	0x05, 0xdf, 0xc0, 0x17, 0x90, 0x4c, 0x92, 0x92, 0xa6, 0xe6, 0xc2, 0xbb, 0x39, 0xc9, 0x39, 0xff,
	0x9f, 0xf3, 0x7f, 0x13, 0xe8, 0xf3, 0x54, 0xcb, 0xcf, 0xb1, 0xf8, 0x9a, 0xe0, 0x0a, 0x85, 0x66,
	0xeb, 0x44, 0x6a, 0xe9, 0x3d, 0x99, 0x4b, 0x39, 0x5f, 0x62, 0xc0, 0xd7, 0x71, 0xc0, 0x85, 0x90,
	0x9a, 0xeb, 0x58, 0x0a, 0x55, 0xbc, 0xa5, 0x37, 0xd0, 0x9d, 0xa0, 0xfe, 0x28, 0x30, 0xc4, 0x4d,
	0x8a, 0x4a, 0x93, 0x4b, 0x38, 0x59, 0x60, 0xe6, 0x5a, 0xbe, 0x35, 0x3c, 0x0d, 0xf3, 0x23, 0x7d,
	0x0d, 0xbd, 0xaa, 0x45, 0xad, 0xa5, 0x50, 0x78, 0xdc, 0x43, 0xfa, 0xf0, 0xe0, 0x3b, 0x5f, 0xa6,
	0xe8, 0xda, 0xbe, 0x35, 0xec, 0x84, 0x45, 0x41, 0xdf, 0x99, 0xc9, 0x0f, 0x5c, 0x64, 0xad, 0xea,
	0xc4, 0x83, 0x47, 0x9b, 0x94, 0x0b, 0x1d, 0xeb, 0xac, 0x1c, 0xde, 0xd7, 0xf4, 0x2d, 0x5c, 0xec,
	0xe7, 0x5b, 0xad, 0x07, 0xe0, 0x18, 0x37, 0xe5, 0xda, 0xfe, 0xc9, 0xb0, 0x13, 0x96, 0x15, 0xbd,
	0x83, 0xc1, 0x04, 0xf5, 0x8c, 0x2b, 0x3d, 0x15, 0x0a, 0x13, 0x8d, 0x51, 0xfb, 0x8a, 0x63, 0xb8,
	0x3a, 0xea, 0xfd, 0xcf, 0x5d, 0x27, 0x70, 0xf6, 0x5e, 0xc6, 0xa2, 0xf2, 0x18, 0x80, 0x23, 0x64,
	0x84, 0xd3, 0xfb, 0x72, 0xb2, 0xac, 0x88, 0x0f, 0x67, 0xf9, 0x69, 0x1c, 0x45, 0x09, 0x2a, 0x65,
	0x24, 0x4e, 0xc3, 0xfa, 0x23, 0xda, 0x83, 0xf3, 0x42, 0xa8, 0xf8, 0x80, 0x97, 0x7f, 0x6c, 0xe8,
	0x8e, 0x53, 0x2d, 0xa7, 0x15, 0x57, 0x32, 0x03, 0xa7, 0x00, 0x42, 0x7a, 0xec, 0x00, 0x9e, 0x77,
	0xc1, 0x0e, 0x49, 0xd1, 0x9b, 0x1f, 0xbf, 0x7e, 0xff, 0xb4, 0x1f, 0x93, 0xeb, 0x20, 0xbf, 0x1b,
	0xa3, 0xfd, 0xdd, 0x08, 0xa4, 0xc0, 0x60, 0xbb, 0xc0, 0x6c, 0x47, 0x3e, 0xc1, 0xc3, 0x32, 0x64,
	0x62, 0xc6, 0x6b, 0xb8, 0xbc, 0x4b, 0xd6, 0xc8, 0x9f, 0x8e, 0x8c, 0xe0, 0x73, 0x72, 0xdb, 0x14,
	0x5c, 0x71, 0x91, 0x15, 0x8a, 0xc1, 0xb6, 0x02, 0xb8, 0x23, 0x1b, 0x43, 0xb0, 0x1e, 0x2c, 0xb9,
	0x62, 0xff, 0xc6, 0xe2, 0xb9, 0xac, 0x85, 0x01, 0x7d, 0x61, 0x4c, 0x6f, 0xc9, 0xb3, 0xa6, 0xe9,
	0x92, 0x2b, 0x3d, 0x8a, 0xcb, 0xf6, 0x72, 0x9f, 0x7b, 0xe8, 0xe4, 0xf9, 0x91, 0x73, 0x56, 0xe3,
	0xe1, 0x75, 0x59, 0x3d, 0x54, 0xfa, 0xd4, 0x28, 0x5e, 0xd3, 0x7e, 0x53, 0xf1, 0x9b, 0x8c, 0xc5,
	0x1b, 0xeb, 0xee, 0x8b, 0x63, 0x7e, 0x8f, 0x57, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xc2, 0x40,
	0x18, 0x0c, 0x54, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AutoIncrementClient is the client API for AutoIncrement service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AutoIncrementClient interface {
	GetOne(ctx context.Context, in *GetOneRequest, opts ...grpc.CallOption) (*GetOneResponse, error)
	GetMany(ctx context.Context, in *GetManyRequest, opts ...grpc.CallOption) (*GetManyResponse, error)
	GetLastInserted(ctx context.Context, in *GetLastInsertedRequest, opts ...grpc.CallOption) (*GetLastInsertedResponse, error)
	Join(ctx context.Context, in *JoinRequest, opts ...grpc.CallOption) (*JoinResponse, error)
}

type autoIncrementClient struct {
	cc *grpc.ClientConn
}

func NewAutoIncrementClient(cc *grpc.ClientConn) AutoIncrementClient {
	return &autoIncrementClient{cc}
}

func (c *autoIncrementClient) GetOne(ctx context.Context, in *GetOneRequest, opts ...grpc.CallOption) (*GetOneResponse, error) {
	out := new(GetOneResponse)
	err := c.cc.Invoke(ctx, "/AutoIncrement/GetOne", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *autoIncrementClient) GetMany(ctx context.Context, in *GetManyRequest, opts ...grpc.CallOption) (*GetManyResponse, error) {
	out := new(GetManyResponse)
	err := c.cc.Invoke(ctx, "/AutoIncrement/GetMany", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *autoIncrementClient) GetLastInserted(ctx context.Context, in *GetLastInsertedRequest, opts ...grpc.CallOption) (*GetLastInsertedResponse, error) {
	out := new(GetLastInsertedResponse)
	err := c.cc.Invoke(ctx, "/AutoIncrement/GetLastInserted", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *autoIncrementClient) Join(ctx context.Context, in *JoinRequest, opts ...grpc.CallOption) (*JoinResponse, error) {
	out := new(JoinResponse)
	err := c.cc.Invoke(ctx, "/AutoIncrement/Join", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AutoIncrementServer is the server API for AutoIncrement service.
type AutoIncrementServer interface {
	GetOne(context.Context, *GetOneRequest) (*GetOneResponse, error)
	GetMany(context.Context, *GetManyRequest) (*GetManyResponse, error)
	GetLastInserted(context.Context, *GetLastInsertedRequest) (*GetLastInsertedResponse, error)
	Join(context.Context, *JoinRequest) (*JoinResponse, error)
}

// UnimplementedAutoIncrementServer can be embedded to have forward compatible implementations.
type UnimplementedAutoIncrementServer struct {
}

func (*UnimplementedAutoIncrementServer) GetOne(ctx context.Context, req *GetOneRequest) (*GetOneResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOne not implemented")
}
func (*UnimplementedAutoIncrementServer) GetMany(ctx context.Context, req *GetManyRequest) (*GetManyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMany not implemented")
}
func (*UnimplementedAutoIncrementServer) GetLastInserted(ctx context.Context, req *GetLastInsertedRequest) (*GetLastInsertedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLastInserted not implemented")
}
func (*UnimplementedAutoIncrementServer) Join(ctx context.Context, req *JoinRequest) (*JoinResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Join not implemented")
}

func RegisterAutoIncrementServer(s *grpc.Server, srv AutoIncrementServer) {
	s.RegisterService(&_AutoIncrement_serviceDesc, srv)
}

func _AutoIncrement_GetOne_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AutoIncrementServer).GetOne(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AutoIncrement/GetOne",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AutoIncrementServer).GetOne(ctx, req.(*GetOneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AutoIncrement_GetMany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetManyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AutoIncrementServer).GetMany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AutoIncrement/GetMany",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AutoIncrementServer).GetMany(ctx, req.(*GetManyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AutoIncrement_GetLastInserted_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLastInsertedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AutoIncrementServer).GetLastInserted(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AutoIncrement/GetLastInserted",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AutoIncrementServer).GetLastInserted(ctx, req.(*GetLastInsertedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AutoIncrement_Join_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AutoIncrementServer).Join(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AutoIncrement/Join",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AutoIncrementServer).Join(ctx, req.(*JoinRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AutoIncrement_serviceDesc = grpc.ServiceDesc{
	ServiceName: "AutoIncrement",
	HandlerType: (*AutoIncrementServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetOne",
			Handler:    _AutoIncrement_GetOne_Handler,
		},
		{
			MethodName: "GetMany",
			Handler:    _AutoIncrement_GetMany_Handler,
		},
		{
			MethodName: "GetLastInserted",
			Handler:    _AutoIncrement_GetLastInserted_Handler,
		},
		{
			MethodName: "Join",
			Handler:    _AutoIncrement_Join_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auto_increment.proto",
}
