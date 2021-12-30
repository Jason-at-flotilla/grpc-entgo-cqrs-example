// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package contactpb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ReadContactServiceClient is the client API for ReadContactService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReadContactServiceClient interface {
	Get(ctx context.Context, in *GetContactReq, opts ...grpc.CallOption) (*GetContactResp, error)
	List(ctx context.Context, in *ListContactReq, opts ...grpc.CallOption) (*ListContactResp, error)
}

type readContactServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewReadContactServiceClient(cc grpc.ClientConnInterface) ReadContactServiceClient {
	return &readContactServiceClient{cc}
}

func (c *readContactServiceClient) Get(ctx context.Context, in *GetContactReq, opts ...grpc.CallOption) (*GetContactResp, error) {
	out := new(GetContactResp)
	err := c.cc.Invoke(ctx, "/Contactpb.ReadContactService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *readContactServiceClient) List(ctx context.Context, in *ListContactReq, opts ...grpc.CallOption) (*ListContactResp, error) {
	out := new(ListContactResp)
	err := c.cc.Invoke(ctx, "/Contactpb.ReadContactService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReadContactServiceServer is the server API for ReadContactService service.
// All implementations must embed UnimplementedReadContactServiceServer
// for forward compatibility
type ReadContactServiceServer interface {
	Get(context.Context, *GetContactReq) (*GetContactResp, error)
	List(context.Context, *ListContactReq) (*ListContactResp, error)
	mustEmbedUnimplementedReadContactServiceServer()
}

// UnimplementedReadContactServiceServer must be embedded to have forward compatible implementations.
type UnimplementedReadContactServiceServer struct {
}

func (UnimplementedReadContactServiceServer) Get(context.Context, *GetContactReq) (*GetContactResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedReadContactServiceServer) List(context.Context, *ListContactReq) (*ListContactResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedReadContactServiceServer) mustEmbedUnimplementedReadContactServiceServer() {}

// UnsafeReadContactServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReadContactServiceServer will
// result in compilation errors.
type UnsafeReadContactServiceServer interface {
	mustEmbedUnimplementedReadContactServiceServer()
}

func RegisterReadContactServiceServer(s grpc.ServiceRegistrar, srv ReadContactServiceServer) {
	s.RegisterService(&ReadContactService_ServiceDesc, srv)
}

func _ReadContactService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetContactReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReadContactServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Contactpb.ReadContactService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReadContactServiceServer).Get(ctx, req.(*GetContactReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReadContactService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListContactReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReadContactServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Contactpb.ReadContactService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReadContactServiceServer).List(ctx, req.(*ListContactReq))
	}
	return interceptor(ctx, in, info, handler)
}

// ReadContactService_ServiceDesc is the grpc.ServiceDesc for ReadContactService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ReadContactService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Contactpb.ReadContactService",
	HandlerType: (*ReadContactServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _ReadContactService_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _ReadContactService_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "contact.proto",
}

// WriteContactServiceClient is the client API for WriteContactService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WriteContactServiceClient interface {
	Create(ctx context.Context, in *CreateContactReq, opts ...grpc.CallOption) (*CreateContactResp, error)
	Update(ctx context.Context, in *UpdateContactReq, opts ...grpc.CallOption) (*UpdateContactResp, error)
	Delete(ctx context.Context, in *DeleteContactReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type writeContactServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWriteContactServiceClient(cc grpc.ClientConnInterface) WriteContactServiceClient {
	return &writeContactServiceClient{cc}
}

func (c *writeContactServiceClient) Create(ctx context.Context, in *CreateContactReq, opts ...grpc.CallOption) (*CreateContactResp, error) {
	out := new(CreateContactResp)
	err := c.cc.Invoke(ctx, "/Contactpb.WriteContactService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *writeContactServiceClient) Update(ctx context.Context, in *UpdateContactReq, opts ...grpc.CallOption) (*UpdateContactResp, error) {
	out := new(UpdateContactResp)
	err := c.cc.Invoke(ctx, "/Contactpb.WriteContactService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *writeContactServiceClient) Delete(ctx context.Context, in *DeleteContactReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/Contactpb.WriteContactService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WriteContactServiceServer is the server API for WriteContactService service.
// All implementations must embed UnimplementedWriteContactServiceServer
// for forward compatibility
type WriteContactServiceServer interface {
	Create(context.Context, *CreateContactReq) (*CreateContactResp, error)
	Update(context.Context, *UpdateContactReq) (*UpdateContactResp, error)
	Delete(context.Context, *DeleteContactReq) (*emptypb.Empty, error)
	mustEmbedUnimplementedWriteContactServiceServer()
}

// UnimplementedWriteContactServiceServer must be embedded to have forward compatible implementations.
type UnimplementedWriteContactServiceServer struct {
}

func (UnimplementedWriteContactServiceServer) Create(context.Context, *CreateContactReq) (*CreateContactResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedWriteContactServiceServer) Update(context.Context, *UpdateContactReq) (*UpdateContactResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedWriteContactServiceServer) Delete(context.Context, *DeleteContactReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedWriteContactServiceServer) mustEmbedUnimplementedWriteContactServiceServer() {}

// UnsafeWriteContactServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WriteContactServiceServer will
// result in compilation errors.
type UnsafeWriteContactServiceServer interface {
	mustEmbedUnimplementedWriteContactServiceServer()
}

func RegisterWriteContactServiceServer(s grpc.ServiceRegistrar, srv WriteContactServiceServer) {
	s.RegisterService(&WriteContactService_ServiceDesc, srv)
}

func _WriteContactService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateContactReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WriteContactServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Contactpb.WriteContactService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WriteContactServiceServer).Create(ctx, req.(*CreateContactReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _WriteContactService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateContactReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WriteContactServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Contactpb.WriteContactService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WriteContactServiceServer).Update(ctx, req.(*UpdateContactReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _WriteContactService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteContactReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WriteContactServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Contactpb.WriteContactService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WriteContactServiceServer).Delete(ctx, req.(*DeleteContactReq))
	}
	return interceptor(ctx, in, info, handler)
}

// WriteContactService_ServiceDesc is the grpc.ServiceDesc for WriteContactService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WriteContactService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Contactpb.WriteContactService",
	HandlerType: (*WriteContactServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _WriteContactService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _WriteContactService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _WriteContactService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "contact.proto",
}
