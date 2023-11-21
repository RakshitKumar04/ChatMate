// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// GrpcServerServiceClient is the client API for GrpcServerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GrpcServerServiceClient interface {
	SignUp(ctx context.Context, in *SignupRequestMessage, opts ...grpc.CallOption) (*SignupResponseMessage, error)
	Login(ctx context.Context, in *LoginRequestMessage, opts ...grpc.CallOption) (*LoginResponseMessage, error)
	GetUser(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*GetUserResponse, error)
}

type grpcServerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGrpcServerServiceClient(cc grpc.ClientConnInterface) GrpcServerServiceClient {
	return &grpcServerServiceClient{cc}
}

func (c *grpcServerServiceClient) SignUp(ctx context.Context, in *SignupRequestMessage, opts ...grpc.CallOption) (*SignupResponseMessage, error) {
	out := new(SignupResponseMessage)
	err := c.cc.Invoke(ctx, "/pb.GrpcServerService/SignUp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *grpcServerServiceClient) Login(ctx context.Context, in *LoginRequestMessage, opts ...grpc.CallOption) (*LoginResponseMessage, error) {
	out := new(LoginResponseMessage)
	err := c.cc.Invoke(ctx, "/pb.GrpcServerService/login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *grpcServerServiceClient) GetUser(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*GetUserResponse, error) {
	out := new(GetUserResponse)
	err := c.cc.Invoke(ctx, "/pb.GrpcServerService/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GrpcServerServiceServer is the server API for GrpcServerService service.
// All implementations must embed UnimplementedGrpcServerServiceServer
// for forward compatibility
type GrpcServerServiceServer interface {
	SignUp(context.Context, *SignupRequestMessage) (*SignupResponseMessage, error)
	Login(context.Context, *LoginRequestMessage) (*LoginResponseMessage, error)
	GetUser(context.Context, *EmptyRequest) (*GetUserResponse, error)
	mustEmbedUnimplementedGrpcServerServiceServer()
}

// UnimplementedGrpcServerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGrpcServerServiceServer struct {
}

func (UnimplementedGrpcServerServiceServer) SignUp(context.Context, *SignupRequestMessage) (*SignupResponseMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignUp not implemented")
}
func (UnimplementedGrpcServerServiceServer) Login(context.Context, *LoginRequestMessage) (*LoginResponseMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedGrpcServerServiceServer) GetUser(context.Context, *EmptyRequest) (*GetUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedGrpcServerServiceServer) mustEmbedUnimplementedGrpcServerServiceServer() {}

// UnsafeGrpcServerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GrpcServerServiceServer will
// result in compilation errors.
type UnsafeGrpcServerServiceServer interface {
	mustEmbedUnimplementedGrpcServerServiceServer()
}

func RegisterGrpcServerServiceServer(s *grpc.Server, srv GrpcServerServiceServer) {
	s.RegisterService(&_GrpcServerService_serviceDesc, srv)
}

func _GrpcServerService_SignUp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignupRequestMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GrpcServerServiceServer).SignUp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.GrpcServerService/SignUp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GrpcServerServiceServer).SignUp(ctx, req.(*SignupRequestMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _GrpcServerService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequestMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GrpcServerServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.GrpcServerService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GrpcServerServiceServer).Login(ctx, req.(*LoginRequestMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _GrpcServerService_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GrpcServerServiceServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.GrpcServerService/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GrpcServerServiceServer).GetUser(ctx, req.(*EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GrpcServerService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.GrpcServerService",
	HandlerType: (*GrpcServerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SignUp",
			Handler:    _GrpcServerService_SignUp_Handler,
		},
		{
			MethodName: "login",
			Handler:    _GrpcServerService_Login_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _GrpcServerService_GetUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc_services.proto",
}
