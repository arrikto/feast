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

// EntityServiceClient is the client API for EntityService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EntityServiceClient interface {
	CreateEntity(ctx context.Context, in *CreateEntityRequest, opts ...grpc.CallOption) (*Entity, error)
	GetEntity(ctx context.Context, in *GetEntityRequest, opts ...grpc.CallOption) (*Entity, error)
	UpdateEntity(ctx context.Context, in *UpdateEntityRequest, opts ...grpc.CallOption) (*Entity, error)
	DeleteEntity(ctx context.Context, in *DeleteEntityRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	ListEntities(ctx context.Context, in *ListEntitiesRequest, opts ...grpc.CallOption) (*ListEntitiesResponse, error)
}

type entityServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEntityServiceClient(cc grpc.ClientConnInterface) EntityServiceClient {
	return &entityServiceClient{cc}
}

func (c *entityServiceClient) CreateEntity(ctx context.Context, in *CreateEntityRequest, opts ...grpc.CallOption) (*Entity, error) {
	out := new(Entity)
	err := c.cc.Invoke(ctx, "/api.EntityService/CreateEntity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *entityServiceClient) GetEntity(ctx context.Context, in *GetEntityRequest, opts ...grpc.CallOption) (*Entity, error) {
	out := new(Entity)
	err := c.cc.Invoke(ctx, "/api.EntityService/GetEntity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *entityServiceClient) UpdateEntity(ctx context.Context, in *UpdateEntityRequest, opts ...grpc.CallOption) (*Entity, error) {
	out := new(Entity)
	err := c.cc.Invoke(ctx, "/api.EntityService/UpdateEntity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *entityServiceClient) DeleteEntity(ctx context.Context, in *DeleteEntityRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/api.EntityService/DeleteEntity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *entityServiceClient) ListEntities(ctx context.Context, in *ListEntitiesRequest, opts ...grpc.CallOption) (*ListEntitiesResponse, error) {
	out := new(ListEntitiesResponse)
	err := c.cc.Invoke(ctx, "/api.EntityService/ListEntities", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EntityServiceServer is the server API for EntityService service.
// All implementations should embed UnimplementedEntityServiceServer
// for forward compatibility
type EntityServiceServer interface {
	CreateEntity(context.Context, *CreateEntityRequest) (*Entity, error)
	GetEntity(context.Context, *GetEntityRequest) (*Entity, error)
	UpdateEntity(context.Context, *UpdateEntityRequest) (*Entity, error)
	DeleteEntity(context.Context, *DeleteEntityRequest) (*empty.Empty, error)
	ListEntities(context.Context, *ListEntitiesRequest) (*ListEntitiesResponse, error)
}

// UnimplementedEntityServiceServer should be embedded to have forward compatible implementations.
type UnimplementedEntityServiceServer struct {
}

func (UnimplementedEntityServiceServer) CreateEntity(context.Context, *CreateEntityRequest) (*Entity, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateEntity not implemented")
}
func (UnimplementedEntityServiceServer) GetEntity(context.Context, *GetEntityRequest) (*Entity, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEntity not implemented")
}
func (UnimplementedEntityServiceServer) UpdateEntity(context.Context, *UpdateEntityRequest) (*Entity, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateEntity not implemented")
}
func (UnimplementedEntityServiceServer) DeleteEntity(context.Context, *DeleteEntityRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteEntity not implemented")
}
func (UnimplementedEntityServiceServer) ListEntities(context.Context, *ListEntitiesRequest) (*ListEntitiesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListEntities not implemented")
}

// UnsafeEntityServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EntityServiceServer will
// result in compilation errors.
type UnsafeEntityServiceServer interface {
	mustEmbedUnimplementedEntityServiceServer()
}

func RegisterEntityServiceServer(s grpc.ServiceRegistrar, srv EntityServiceServer) {
	s.RegisterService(&EntityService_ServiceDesc, srv)
}

func _EntityService_CreateEntity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateEntityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EntityServiceServer).CreateEntity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.EntityService/CreateEntity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EntityServiceServer).CreateEntity(ctx, req.(*CreateEntityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EntityService_GetEntity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEntityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EntityServiceServer).GetEntity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.EntityService/GetEntity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EntityServiceServer).GetEntity(ctx, req.(*GetEntityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EntityService_UpdateEntity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateEntityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EntityServiceServer).UpdateEntity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.EntityService/UpdateEntity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EntityServiceServer).UpdateEntity(ctx, req.(*UpdateEntityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EntityService_DeleteEntity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteEntityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EntityServiceServer).DeleteEntity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.EntityService/DeleteEntity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EntityServiceServer).DeleteEntity(ctx, req.(*DeleteEntityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EntityService_ListEntities_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListEntitiesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EntityServiceServer).ListEntities(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.EntityService/ListEntities",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EntityServiceServer).ListEntities(ctx, req.(*ListEntitiesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// EntityService_ServiceDesc is the grpc.ServiceDesc for EntityService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EntityService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.EntityService",
	HandlerType: (*EntityServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateEntity",
			Handler:    _EntityService_CreateEntity_Handler,
		},
		{
			MethodName: "GetEntity",
			Handler:    _EntityService_GetEntity_Handler,
		},
		{
			MethodName: "UpdateEntity",
			Handler:    _EntityService_UpdateEntity_Handler,
		},
		{
			MethodName: "DeleteEntity",
			Handler:    _EntityService_DeleteEntity_Handler,
		},
		{
			MethodName: "ListEntities",
			Handler:    _EntityService_ListEntities_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "EntityService.proto",
}
