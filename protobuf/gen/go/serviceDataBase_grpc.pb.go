// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: serviceDataBase.proto

package serv1

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
	DataBaseService_GetUser_FullMethodName    = "/services.DataBaseService/GetUser"
	DataBaseService_SaveUser_FullMethodName   = "/services.DataBaseService/SaveUser"
	DataBaseService_GetTask_FullMethodName    = "/services.DataBaseService/GetTask"
	DataBaseService_SaveTask_FullMethodName   = "/services.DataBaseService/SaveTask"
	DataBaseService_GetAnswer_FullMethodName  = "/services.DataBaseService/GetAnswer"
	DataBaseService_SaveAnswer_FullMethodName = "/services.DataBaseService/SaveAnswer"
)

// DataBaseServiceClient is the client API for DataBaseService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DataBaseServiceClient interface {
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error)
	SaveUser(ctx context.Context, in *SaveUserRequest, opts ...grpc.CallOption) (*SaveUserResponse, error)
	GetTask(ctx context.Context, in *GetTaskRequest, opts ...grpc.CallOption) (*GetTaskResponse, error)
	SaveTask(ctx context.Context, in *SaveTaskRequest, opts ...grpc.CallOption) (*SaveTaskResponse, error)
	GetAnswer(ctx context.Context, in *GetAnswerRequest, opts ...grpc.CallOption) (*GetAnswerResponse, error)
	SaveAnswer(ctx context.Context, in *SaveAnswerRequest, opts ...grpc.CallOption) (*SaveAnswerResponse, error)
}

type dataBaseServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDataBaseServiceClient(cc grpc.ClientConnInterface) DataBaseServiceClient {
	return &dataBaseServiceClient{cc}
}

func (c *dataBaseServiceClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error) {
	out := new(GetUserResponse)
	err := c.cc.Invoke(ctx, DataBaseService_GetUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataBaseServiceClient) SaveUser(ctx context.Context, in *SaveUserRequest, opts ...grpc.CallOption) (*SaveUserResponse, error) {
	out := new(SaveUserResponse)
	err := c.cc.Invoke(ctx, DataBaseService_SaveUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataBaseServiceClient) GetTask(ctx context.Context, in *GetTaskRequest, opts ...grpc.CallOption) (*GetTaskResponse, error) {
	out := new(GetTaskResponse)
	err := c.cc.Invoke(ctx, DataBaseService_GetTask_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataBaseServiceClient) SaveTask(ctx context.Context, in *SaveTaskRequest, opts ...grpc.CallOption) (*SaveTaskResponse, error) {
	out := new(SaveTaskResponse)
	err := c.cc.Invoke(ctx, DataBaseService_SaveTask_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataBaseServiceClient) GetAnswer(ctx context.Context, in *GetAnswerRequest, opts ...grpc.CallOption) (*GetAnswerResponse, error) {
	out := new(GetAnswerResponse)
	err := c.cc.Invoke(ctx, DataBaseService_GetAnswer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataBaseServiceClient) SaveAnswer(ctx context.Context, in *SaveAnswerRequest, opts ...grpc.CallOption) (*SaveAnswerResponse, error) {
	out := new(SaveAnswerResponse)
	err := c.cc.Invoke(ctx, DataBaseService_SaveAnswer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DataBaseServiceServer is the server API for DataBaseService service.
// All implementations must embed UnimplementedDataBaseServiceServer
// for forward compatibility
type DataBaseServiceServer interface {
	GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error)
	SaveUser(context.Context, *SaveUserRequest) (*SaveUserResponse, error)
	GetTask(context.Context, *GetTaskRequest) (*GetTaskResponse, error)
	SaveTask(context.Context, *SaveTaskRequest) (*SaveTaskResponse, error)
	GetAnswer(context.Context, *GetAnswerRequest) (*GetAnswerResponse, error)
	SaveAnswer(context.Context, *SaveAnswerRequest) (*SaveAnswerResponse, error)
	mustEmbedUnimplementedDataBaseServiceServer()
}

// UnimplementedDataBaseServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDataBaseServiceServer struct {
}

func (UnimplementedDataBaseServiceServer) GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedDataBaseServiceServer) SaveUser(context.Context, *SaveUserRequest) (*SaveUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveUser not implemented")
}
func (UnimplementedDataBaseServiceServer) GetTask(context.Context, *GetTaskRequest) (*GetTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTask not implemented")
}
func (UnimplementedDataBaseServiceServer) SaveTask(context.Context, *SaveTaskRequest) (*SaveTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveTask not implemented")
}
func (UnimplementedDataBaseServiceServer) GetAnswer(context.Context, *GetAnswerRequest) (*GetAnswerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAnswer not implemented")
}
func (UnimplementedDataBaseServiceServer) SaveAnswer(context.Context, *SaveAnswerRequest) (*SaveAnswerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveAnswer not implemented")
}
func (UnimplementedDataBaseServiceServer) mustEmbedUnimplementedDataBaseServiceServer() {}

// UnsafeDataBaseServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DataBaseServiceServer will
// result in compilation errors.
type UnsafeDataBaseServiceServer interface {
	mustEmbedUnimplementedDataBaseServiceServer()
}

func RegisterDataBaseServiceServer(s grpc.ServiceRegistrar, srv DataBaseServiceServer) {
	s.RegisterService(&DataBaseService_ServiceDesc, srv)
}

func _DataBaseService_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataBaseServiceServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DataBaseService_GetUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataBaseServiceServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataBaseService_SaveUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataBaseServiceServer).SaveUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DataBaseService_SaveUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataBaseServiceServer).SaveUser(ctx, req.(*SaveUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataBaseService_GetTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataBaseServiceServer).GetTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DataBaseService_GetTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataBaseServiceServer).GetTask(ctx, req.(*GetTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataBaseService_SaveTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataBaseServiceServer).SaveTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DataBaseService_SaveTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataBaseServiceServer).SaveTask(ctx, req.(*SaveTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataBaseService_GetAnswer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAnswerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataBaseServiceServer).GetAnswer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DataBaseService_GetAnswer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataBaseServiceServer).GetAnswer(ctx, req.(*GetAnswerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataBaseService_SaveAnswer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveAnswerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataBaseServiceServer).SaveAnswer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DataBaseService_SaveAnswer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataBaseServiceServer).SaveAnswer(ctx, req.(*SaveAnswerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DataBaseService_ServiceDesc is the grpc.ServiceDesc for DataBaseService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DataBaseService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "services.DataBaseService",
	HandlerType: (*DataBaseServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUser",
			Handler:    _DataBaseService_GetUser_Handler,
		},
		{
			MethodName: "SaveUser",
			Handler:    _DataBaseService_SaveUser_Handler,
		},
		{
			MethodName: "GetTask",
			Handler:    _DataBaseService_GetTask_Handler,
		},
		{
			MethodName: "SaveTask",
			Handler:    _DataBaseService_SaveTask_Handler,
		},
		{
			MethodName: "GetAnswer",
			Handler:    _DataBaseService_GetAnswer_Handler,
		},
		{
			MethodName: "SaveAnswer",
			Handler:    _DataBaseService_SaveAnswer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "serviceDataBase.proto",
}