// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: syncer.proto

package syncerpb

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

// SyncerServiceClient is the client API for SyncerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SyncerServiceClient interface {
	CreateResource(ctx context.Context, in *CreateResourceReq, opts ...grpc.CallOption) (*SyncResp, error)
	DeleteResource(ctx context.Context, in *DeleteResourceReq, opts ...grpc.CallOption) (*SyncResp, error)
	CreateAggregationRel(ctx context.Context, in *CreateAggregationRelReq, opts ...grpc.CallOption) (*SyncResp, error)
	DeleteAggregationRel(ctx context.Context, in *DeleteAggregationRelReq, opts ...grpc.CallOption) (*SyncResp, error)
	CreateCompositionRel(ctx context.Context, in *CreateCompositionRelReq, opts ...grpc.CallOption) (*SyncResp, error)
	DeleteCompositionRel(ctx context.Context, in *DeleteCompositionRelReq, opts ...grpc.CallOption) (*SyncResp, error)
	CreateAttribute(ctx context.Context, in *CreateAttributeReq, opts ...grpc.CallOption) (*SyncResp, error)
	UpdateAttribute(ctx context.Context, in *UpdateAttributeReq, opts ...grpc.CallOption) (*SyncResp, error)
	DeleteAttribute(ctx context.Context, in *DeleteAttributeReq, opts ...grpc.CallOption) (*SyncResp, error)
	CreatePermission(ctx context.Context, in *CreatePermissionReq, opts ...grpc.CallOption) (*SyncResp, error)
	DeletePermission(ctx context.Context, in *DeletePermissionReq, opts ...grpc.CallOption) (*SyncResp, error)
}

type syncerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSyncerServiceClient(cc grpc.ClientConnInterface) SyncerServiceClient {
	return &syncerServiceClient{cc}
}

func (c *syncerServiceClient) CreateResource(ctx context.Context, in *CreateResourceReq, opts ...grpc.CallOption) (*SyncResp, error) {
	out := new(SyncResp)
	err := c.cc.Invoke(ctx, "/syncerpb.SyncerService/CreateResource", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syncerServiceClient) DeleteResource(ctx context.Context, in *DeleteResourceReq, opts ...grpc.CallOption) (*SyncResp, error) {
	out := new(SyncResp)
	err := c.cc.Invoke(ctx, "/syncerpb.SyncerService/DeleteResource", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syncerServiceClient) CreateAggregationRel(ctx context.Context, in *CreateAggregationRelReq, opts ...grpc.CallOption) (*SyncResp, error) {
	out := new(SyncResp)
	err := c.cc.Invoke(ctx, "/syncerpb.SyncerService/CreateAggregationRel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syncerServiceClient) DeleteAggregationRel(ctx context.Context, in *DeleteAggregationRelReq, opts ...grpc.CallOption) (*SyncResp, error) {
	out := new(SyncResp)
	err := c.cc.Invoke(ctx, "/syncerpb.SyncerService/DeleteAggregationRel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syncerServiceClient) CreateCompositionRel(ctx context.Context, in *CreateCompositionRelReq, opts ...grpc.CallOption) (*SyncResp, error) {
	out := new(SyncResp)
	err := c.cc.Invoke(ctx, "/syncerpb.SyncerService/CreateCompositionRel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syncerServiceClient) DeleteCompositionRel(ctx context.Context, in *DeleteCompositionRelReq, opts ...grpc.CallOption) (*SyncResp, error) {
	out := new(SyncResp)
	err := c.cc.Invoke(ctx, "/syncerpb.SyncerService/DeleteCompositionRel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syncerServiceClient) CreateAttribute(ctx context.Context, in *CreateAttributeReq, opts ...grpc.CallOption) (*SyncResp, error) {
	out := new(SyncResp)
	err := c.cc.Invoke(ctx, "/syncerpb.SyncerService/CreateAttribute", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syncerServiceClient) UpdateAttribute(ctx context.Context, in *UpdateAttributeReq, opts ...grpc.CallOption) (*SyncResp, error) {
	out := new(SyncResp)
	err := c.cc.Invoke(ctx, "/syncerpb.SyncerService/UpdateAttribute", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syncerServiceClient) DeleteAttribute(ctx context.Context, in *DeleteAttributeReq, opts ...grpc.CallOption) (*SyncResp, error) {
	out := new(SyncResp)
	err := c.cc.Invoke(ctx, "/syncerpb.SyncerService/DeleteAttribute", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syncerServiceClient) CreatePermission(ctx context.Context, in *CreatePermissionReq, opts ...grpc.CallOption) (*SyncResp, error) {
	out := new(SyncResp)
	err := c.cc.Invoke(ctx, "/syncerpb.SyncerService/CreatePermission", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syncerServiceClient) DeletePermission(ctx context.Context, in *DeletePermissionReq, opts ...grpc.CallOption) (*SyncResp, error) {
	out := new(SyncResp)
	err := c.cc.Invoke(ctx, "/syncerpb.SyncerService/DeletePermission", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SyncerServiceServer is the server API for SyncerService service.
// All implementations must embed UnimplementedSyncerServiceServer
// for forward compatibility
type SyncerServiceServer interface {
	CreateResource(context.Context, *CreateResourceReq) (*SyncResp, error)
	DeleteResource(context.Context, *DeleteResourceReq) (*SyncResp, error)
	CreateAggregationRel(context.Context, *CreateAggregationRelReq) (*SyncResp, error)
	DeleteAggregationRel(context.Context, *DeleteAggregationRelReq) (*SyncResp, error)
	CreateCompositionRel(context.Context, *CreateCompositionRelReq) (*SyncResp, error)
	DeleteCompositionRel(context.Context, *DeleteCompositionRelReq) (*SyncResp, error)
	CreateAttribute(context.Context, *CreateAttributeReq) (*SyncResp, error)
	UpdateAttribute(context.Context, *UpdateAttributeReq) (*SyncResp, error)
	DeleteAttribute(context.Context, *DeleteAttributeReq) (*SyncResp, error)
	CreatePermission(context.Context, *CreatePermissionReq) (*SyncResp, error)
	DeletePermission(context.Context, *DeletePermissionReq) (*SyncResp, error)
	mustEmbedUnimplementedSyncerServiceServer()
}

// UnimplementedSyncerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSyncerServiceServer struct {
}

func (UnimplementedSyncerServiceServer) CreateResource(context.Context, *CreateResourceReq) (*SyncResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateResource not implemented")
}
func (UnimplementedSyncerServiceServer) DeleteResource(context.Context, *DeleteResourceReq) (*SyncResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteResource not implemented")
}
func (UnimplementedSyncerServiceServer) CreateAggregationRel(context.Context, *CreateAggregationRelReq) (*SyncResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAggregationRel not implemented")
}
func (UnimplementedSyncerServiceServer) DeleteAggregationRel(context.Context, *DeleteAggregationRelReq) (*SyncResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAggregationRel not implemented")
}
func (UnimplementedSyncerServiceServer) CreateCompositionRel(context.Context, *CreateCompositionRelReq) (*SyncResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCompositionRel not implemented")
}
func (UnimplementedSyncerServiceServer) DeleteCompositionRel(context.Context, *DeleteCompositionRelReq) (*SyncResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCompositionRel not implemented")
}
func (UnimplementedSyncerServiceServer) CreateAttribute(context.Context, *CreateAttributeReq) (*SyncResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAttribute not implemented")
}
func (UnimplementedSyncerServiceServer) UpdateAttribute(context.Context, *UpdateAttributeReq) (*SyncResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAttribute not implemented")
}
func (UnimplementedSyncerServiceServer) DeleteAttribute(context.Context, *DeleteAttributeReq) (*SyncResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAttribute not implemented")
}
func (UnimplementedSyncerServiceServer) CreatePermission(context.Context, *CreatePermissionReq) (*SyncResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePermission not implemented")
}
func (UnimplementedSyncerServiceServer) DeletePermission(context.Context, *DeletePermissionReq) (*SyncResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePermission not implemented")
}
func (UnimplementedSyncerServiceServer) mustEmbedUnimplementedSyncerServiceServer() {}

// UnsafeSyncerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SyncerServiceServer will
// result in compilation errors.
type UnsafeSyncerServiceServer interface {
	mustEmbedUnimplementedSyncerServiceServer()
}

func RegisterSyncerServiceServer(s grpc.ServiceRegistrar, srv SyncerServiceServer) {
	s.RegisterService(&SyncerService_ServiceDesc, srv)
}

func _SyncerService_CreateResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateResourceReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncerServiceServer).CreateResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/syncerpb.SyncerService/CreateResource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncerServiceServer).CreateResource(ctx, req.(*CreateResourceReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyncerService_DeleteResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteResourceReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncerServiceServer).DeleteResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/syncerpb.SyncerService/DeleteResource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncerServiceServer).DeleteResource(ctx, req.(*DeleteResourceReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyncerService_CreateAggregationRel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAggregationRelReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncerServiceServer).CreateAggregationRel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/syncerpb.SyncerService/CreateAggregationRel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncerServiceServer).CreateAggregationRel(ctx, req.(*CreateAggregationRelReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyncerService_DeleteAggregationRel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAggregationRelReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncerServiceServer).DeleteAggregationRel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/syncerpb.SyncerService/DeleteAggregationRel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncerServiceServer).DeleteAggregationRel(ctx, req.(*DeleteAggregationRelReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyncerService_CreateCompositionRel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCompositionRelReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncerServiceServer).CreateCompositionRel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/syncerpb.SyncerService/CreateCompositionRel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncerServiceServer).CreateCompositionRel(ctx, req.(*CreateCompositionRelReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyncerService_DeleteCompositionRel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCompositionRelReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncerServiceServer).DeleteCompositionRel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/syncerpb.SyncerService/DeleteCompositionRel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncerServiceServer).DeleteCompositionRel(ctx, req.(*DeleteCompositionRelReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyncerService_CreateAttribute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAttributeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncerServiceServer).CreateAttribute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/syncerpb.SyncerService/CreateAttribute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncerServiceServer).CreateAttribute(ctx, req.(*CreateAttributeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyncerService_UpdateAttribute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAttributeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncerServiceServer).UpdateAttribute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/syncerpb.SyncerService/UpdateAttribute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncerServiceServer).UpdateAttribute(ctx, req.(*UpdateAttributeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyncerService_DeleteAttribute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAttributeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncerServiceServer).DeleteAttribute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/syncerpb.SyncerService/DeleteAttribute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncerServiceServer).DeleteAttribute(ctx, req.(*DeleteAttributeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyncerService_CreatePermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePermissionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncerServiceServer).CreatePermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/syncerpb.SyncerService/CreatePermission",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncerServiceServer).CreatePermission(ctx, req.(*CreatePermissionReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyncerService_DeletePermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePermissionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncerServiceServer).DeletePermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/syncerpb.SyncerService/DeletePermission",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncerServiceServer).DeletePermission(ctx, req.(*DeletePermissionReq))
	}
	return interceptor(ctx, in, info, handler)
}

// SyncerService_ServiceDesc is the grpc.ServiceDesc for SyncerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SyncerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "syncerpb.SyncerService",
	HandlerType: (*SyncerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateResource",
			Handler:    _SyncerService_CreateResource_Handler,
		},
		{
			MethodName: "DeleteResource",
			Handler:    _SyncerService_DeleteResource_Handler,
		},
		{
			MethodName: "CreateAggregationRel",
			Handler:    _SyncerService_CreateAggregationRel_Handler,
		},
		{
			MethodName: "DeleteAggregationRel",
			Handler:    _SyncerService_DeleteAggregationRel_Handler,
		},
		{
			MethodName: "CreateCompositionRel",
			Handler:    _SyncerService_CreateCompositionRel_Handler,
		},
		{
			MethodName: "DeleteCompositionRel",
			Handler:    _SyncerService_DeleteCompositionRel_Handler,
		},
		{
			MethodName: "CreateAttribute",
			Handler:    _SyncerService_CreateAttribute_Handler,
		},
		{
			MethodName: "UpdateAttribute",
			Handler:    _SyncerService_UpdateAttribute_Handler,
		},
		{
			MethodName: "DeleteAttribute",
			Handler:    _SyncerService_DeleteAttribute_Handler,
		},
		{
			MethodName: "CreatePermission",
			Handler:    _SyncerService_CreatePermission_Handler,
		},
		{
			MethodName: "DeletePermission",
			Handler:    _SyncerService_DeletePermission_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "syncer.proto",
}
