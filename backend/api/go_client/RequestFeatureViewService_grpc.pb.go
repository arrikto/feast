// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package go_client

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// RequestFeatureViewServiceClient is the client API for RequestFeatureViewService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RequestFeatureViewServiceClient interface {
	CreateRequestFeatureView(ctx context.Context, in *CreateRequestFeatureViewRequest, opts ...grpc.CallOption) (*RequestFeatureView, error)
	GetRequestFeatureView(ctx context.Context, in *GetRequestFeatureViewRequest, opts ...grpc.CallOption) (*RequestFeatureView, error)
	UpdateRequestFeatureView(ctx context.Context, in *UpdateRequestFeatureViewRequest, opts ...grpc.CallOption) (*RequestFeatureView, error)
	DeleteRequestFeatureView(ctx context.Context, in *DeleteRequestFeatureViewRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	ListRequestFeatureViews(ctx context.Context, in *ListRequestFeatureViewsRequest, opts ...grpc.CallOption) (*ListRequestFeatureViewsResponse, error)
}

type requestFeatureViewServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRequestFeatureViewServiceClient(cc grpc.ClientConnInterface) RequestFeatureViewServiceClient {
	return &requestFeatureViewServiceClient{cc}
}

func (c *requestFeatureViewServiceClient) CreateRequestFeatureView(ctx context.Context, in *CreateRequestFeatureViewRequest, opts ...grpc.CallOption) (*RequestFeatureView, error) {
	out := new(RequestFeatureView)
	err := c.cc.Invoke(ctx, "/api.RequestFeatureViewService/CreateRequestFeatureView", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *requestFeatureViewServiceClient) GetRequestFeatureView(ctx context.Context, in *GetRequestFeatureViewRequest, opts ...grpc.CallOption) (*RequestFeatureView, error) {
	out := new(RequestFeatureView)
	err := c.cc.Invoke(ctx, "/api.RequestFeatureViewService/GetRequestFeatureView", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *requestFeatureViewServiceClient) UpdateRequestFeatureView(ctx context.Context, in *UpdateRequestFeatureViewRequest, opts ...grpc.CallOption) (*RequestFeatureView, error) {
	out := new(RequestFeatureView)
	err := c.cc.Invoke(ctx, "/api.RequestFeatureViewService/UpdateRequestFeatureView", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *requestFeatureViewServiceClient) DeleteRequestFeatureView(ctx context.Context, in *DeleteRequestFeatureViewRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/api.RequestFeatureViewService/DeleteRequestFeatureView", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *requestFeatureViewServiceClient) ListRequestFeatureViews(ctx context.Context, in *ListRequestFeatureViewsRequest, opts ...grpc.CallOption) (*ListRequestFeatureViewsResponse, error) {
	out := new(ListRequestFeatureViewsResponse)
	err := c.cc.Invoke(ctx, "/api.RequestFeatureViewService/ListRequestFeatureViews", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RequestFeatureViewServiceServer is the server API for RequestFeatureViewService service.
// All implementations should embed UnimplementedRequestFeatureViewServiceServer
// for forward compatibility
type RequestFeatureViewServiceServer interface {
	CreateRequestFeatureView(context.Context, *CreateRequestFeatureViewRequest) (*RequestFeatureView, error)
	GetRequestFeatureView(context.Context, *GetRequestFeatureViewRequest) (*RequestFeatureView, error)
	UpdateRequestFeatureView(context.Context, *UpdateRequestFeatureViewRequest) (*RequestFeatureView, error)
	DeleteRequestFeatureView(context.Context, *DeleteRequestFeatureViewRequest) (*empty.Empty, error)
	ListRequestFeatureViews(context.Context, *ListRequestFeatureViewsRequest) (*ListRequestFeatureViewsResponse, error)
}

// UnimplementedRequestFeatureViewServiceServer should be embedded to have forward compatible implementations.
type UnimplementedRequestFeatureViewServiceServer struct {
}

func (UnimplementedRequestFeatureViewServiceServer) CreateRequestFeatureView(context.Context, *CreateRequestFeatureViewRequest) (*RequestFeatureView, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRequestFeatureView not implemented")
}
func (UnimplementedRequestFeatureViewServiceServer) GetRequestFeatureView(context.Context, *GetRequestFeatureViewRequest) (*RequestFeatureView, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRequestFeatureView not implemented")
}
func (UnimplementedRequestFeatureViewServiceServer) UpdateRequestFeatureView(context.Context, *UpdateRequestFeatureViewRequest) (*RequestFeatureView, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRequestFeatureView not implemented")
}
func (UnimplementedRequestFeatureViewServiceServer) DeleteRequestFeatureView(context.Context, *DeleteRequestFeatureViewRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRequestFeatureView not implemented")
}
func (UnimplementedRequestFeatureViewServiceServer) ListRequestFeatureViews(context.Context, *ListRequestFeatureViewsRequest) (*ListRequestFeatureViewsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRequestFeatureViews not implemented")
}

// UnsafeRequestFeatureViewServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RequestFeatureViewServiceServer will
// result in compilation errors.
type UnsafeRequestFeatureViewServiceServer interface {
	mustEmbedUnimplementedRequestFeatureViewServiceServer()
}

func RegisterRequestFeatureViewServiceServer(s grpc.ServiceRegistrar, srv RequestFeatureViewServiceServer) {
	s.RegisterService(&RequestFeatureViewService_ServiceDesc, srv)
}

func _RequestFeatureViewService_CreateRequestFeatureView_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequestFeatureViewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RequestFeatureViewServiceServer).CreateRequestFeatureView(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.RequestFeatureViewService/CreateRequestFeatureView",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RequestFeatureViewServiceServer).CreateRequestFeatureView(ctx, req.(*CreateRequestFeatureViewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RequestFeatureViewService_GetRequestFeatureView_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequestFeatureViewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RequestFeatureViewServiceServer).GetRequestFeatureView(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.RequestFeatureViewService/GetRequestFeatureView",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RequestFeatureViewServiceServer).GetRequestFeatureView(ctx, req.(*GetRequestFeatureViewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RequestFeatureViewService_UpdateRequestFeatureView_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequestFeatureViewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RequestFeatureViewServiceServer).UpdateRequestFeatureView(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.RequestFeatureViewService/UpdateRequestFeatureView",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RequestFeatureViewServiceServer).UpdateRequestFeatureView(ctx, req.(*UpdateRequestFeatureViewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RequestFeatureViewService_DeleteRequestFeatureView_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequestFeatureViewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RequestFeatureViewServiceServer).DeleteRequestFeatureView(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.RequestFeatureViewService/DeleteRequestFeatureView",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RequestFeatureViewServiceServer).DeleteRequestFeatureView(ctx, req.(*DeleteRequestFeatureViewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RequestFeatureViewService_ListRequestFeatureViews_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequestFeatureViewsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RequestFeatureViewServiceServer).ListRequestFeatureViews(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.RequestFeatureViewService/ListRequestFeatureViews",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RequestFeatureViewServiceServer).ListRequestFeatureViews(ctx, req.(*ListRequestFeatureViewsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RequestFeatureViewService_ServiceDesc is the grpc.ServiceDesc for RequestFeatureViewService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RequestFeatureViewService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.RequestFeatureViewService",
	HandlerType: (*RequestFeatureViewServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateRequestFeatureView",
			Handler:    _RequestFeatureViewService_CreateRequestFeatureView_Handler,
		},
		{
			MethodName: "GetRequestFeatureView",
			Handler:    _RequestFeatureViewService_GetRequestFeatureView_Handler,
		},
		{
			MethodName: "UpdateRequestFeatureView",
			Handler:    _RequestFeatureViewService_UpdateRequestFeatureView_Handler,
		},
		{
			MethodName: "DeleteRequestFeatureView",
			Handler:    _RequestFeatureViewService_DeleteRequestFeatureView_Handler,
		},
		{
			MethodName: "ListRequestFeatureViews",
			Handler:    _RequestFeatureViewService_ListRequestFeatureViews_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "RequestFeatureViewService.proto",
}