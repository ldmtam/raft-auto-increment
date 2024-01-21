// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: pb/auto_increment.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	AutoIncrementService_GetOne_FullMethodName          = "/raftautoincrement.service.v1.AutoIncrementService/GetOne"
	AutoIncrementService_GetMany_FullMethodName         = "/raftautoincrement.service.v1.AutoIncrementService/GetMany"
	AutoIncrementService_GetLastInserted_FullMethodName = "/raftautoincrement.service.v1.AutoIncrementService/GetLastInserted"
	AutoIncrementService_Join_FullMethodName            = "/raftautoincrement.service.v1.AutoIncrementService/Join"
)

// AutoIncrementServiceClient is the client API for AutoIncrementService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AutoIncrementServiceClient interface {
	GetOne(ctx context.Context, in *GetOneRequest, opts ...grpc.CallOption) (*GetOneResponse, error)
	GetMany(ctx context.Context, in *GetManyRequest, opts ...grpc.CallOption) (*GetManyResponse, error)
	GetLastInserted(ctx context.Context, in *GetLastInsertedRequest, opts ...grpc.CallOption) (*GetLastInsertedResponse, error)
	Join(ctx context.Context, in *JoinRequest, opts ...grpc.CallOption) (*JoinResponse, error)
}

type autoIncrementServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAutoIncrementServiceClient(cc grpc.ClientConnInterface) AutoIncrementServiceClient {
	return &autoIncrementServiceClient{cc}
}

func (c *autoIncrementServiceClient) GetOne(ctx context.Context, in *GetOneRequest, opts ...grpc.CallOption) (*GetOneResponse, error) {
	out := new(GetOneResponse)
	err := c.cc.Invoke(ctx, AutoIncrementService_GetOne_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *autoIncrementServiceClient) GetMany(ctx context.Context, in *GetManyRequest, opts ...grpc.CallOption) (*GetManyResponse, error) {
	out := new(GetManyResponse)
	err := c.cc.Invoke(ctx, AutoIncrementService_GetMany_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *autoIncrementServiceClient) GetLastInserted(ctx context.Context, in *GetLastInsertedRequest, opts ...grpc.CallOption) (*GetLastInsertedResponse, error) {
	out := new(GetLastInsertedResponse)
	err := c.cc.Invoke(ctx, AutoIncrementService_GetLastInserted_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *autoIncrementServiceClient) Join(ctx context.Context, in *JoinRequest, opts ...grpc.CallOption) (*JoinResponse, error) {
	out := new(JoinResponse)
	err := c.cc.Invoke(ctx, AutoIncrementService_Join_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AutoIncrementServiceServer is the server API for AutoIncrementService service.
// All implementations should embed UnimplementedAutoIncrementServiceServer
// for forward compatibility
type AutoIncrementServiceServer interface {
	GetOne(context.Context, *GetOneRequest) (*GetOneResponse, error)
	GetMany(context.Context, *GetManyRequest) (*GetManyResponse, error)
	GetLastInserted(context.Context, *GetLastInsertedRequest) (*GetLastInsertedResponse, error)
	Join(context.Context, *JoinRequest) (*JoinResponse, error)
}

// UnimplementedAutoIncrementServiceServer should be embedded to have forward compatible implementations.
type UnimplementedAutoIncrementServiceServer struct {
}

func (UnimplementedAutoIncrementServiceServer) GetOne(context.Context, *GetOneRequest) (*GetOneResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOne not implemented")
}
func (UnimplementedAutoIncrementServiceServer) GetMany(context.Context, *GetManyRequest) (*GetManyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMany not implemented")
}
func (UnimplementedAutoIncrementServiceServer) GetLastInserted(context.Context, *GetLastInsertedRequest) (*GetLastInsertedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLastInserted not implemented")
}
func (UnimplementedAutoIncrementServiceServer) Join(context.Context, *JoinRequest) (*JoinResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Join not implemented")
}

// UnsafeAutoIncrementServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AutoIncrementServiceServer will
// result in compilation errors.
type UnsafeAutoIncrementServiceServer interface {
	mustEmbedUnimplementedAutoIncrementServiceServer()
}

func RegisterAutoIncrementServiceServer(s grpc.ServiceRegistrar, srv AutoIncrementServiceServer) {
	s.RegisterService(&AutoIncrementService_ServiceDesc, srv)
}

func _AutoIncrementService_GetOne_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AutoIncrementServiceServer).GetOne(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AutoIncrementService_GetOne_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AutoIncrementServiceServer).GetOne(ctx, req.(*GetOneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AutoIncrementService_GetMany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetManyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AutoIncrementServiceServer).GetMany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AutoIncrementService_GetMany_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AutoIncrementServiceServer).GetMany(ctx, req.(*GetManyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AutoIncrementService_GetLastInserted_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLastInsertedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AutoIncrementServiceServer).GetLastInserted(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AutoIncrementService_GetLastInserted_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AutoIncrementServiceServer).GetLastInserted(ctx, req.(*GetLastInsertedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AutoIncrementService_Join_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AutoIncrementServiceServer).Join(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AutoIncrementService_Join_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AutoIncrementServiceServer).Join(ctx, req.(*JoinRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AutoIncrementService_ServiceDesc is the grpc.ServiceDesc for AutoIncrementService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AutoIncrementService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "raftautoincrement.service.v1.AutoIncrementService",
	HandlerType: (*AutoIncrementServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetOne",
			Handler:    _AutoIncrementService_GetOne_Handler,
		},
		{
			MethodName: "GetMany",
			Handler:    _AutoIncrementService_GetMany_Handler,
		},
		{
			MethodName: "GetLastInserted",
			Handler:    _AutoIncrementService_GetLastInserted_Handler,
		},
		{
			MethodName: "Join",
			Handler:    _AutoIncrementService_Join_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/auto_increment.proto",
}