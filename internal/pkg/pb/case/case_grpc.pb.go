// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: internal/pkg/pb/case/case.proto

package _case

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

// CaseClient is the client API for Case service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CaseClient interface {
	ListCases(ctx context.Context, in *ListCasesRequest, opts ...grpc.CallOption) (Case_ListCasesClient, error)
	GetCase(ctx context.Context, in *GetCaseRequest, opts ...grpc.CallOption) (*GetCaseReply, error)
	CreateCase(ctx context.Context, in *CreateCaseRequest, opts ...grpc.CallOption) (*CreateCaseReply, error)
	UpdateCase(ctx context.Context, in *UpdateCaseRequest, opts ...grpc.CallOption) (*UpdateCaseReply, error)
	DeleteCase(ctx context.Context, in *DeleteCaseRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type caseClient struct {
	cc grpc.ClientConnInterface
}

func NewCaseClient(cc grpc.ClientConnInterface) CaseClient {
	return &caseClient{cc}
}

func (c *caseClient) ListCases(ctx context.Context, in *ListCasesRequest, opts ...grpc.CallOption) (Case_ListCasesClient, error) {
	stream, err := c.cc.NewStream(ctx, &Case_ServiceDesc.Streams[0], "/case.Case/ListCases", opts...)
	if err != nil {
		return nil, err
	}
	x := &caseListCasesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Case_ListCasesClient interface {
	Recv() (*ListCasesReply, error)
	grpc.ClientStream
}

type caseListCasesClient struct {
	grpc.ClientStream
}

func (x *caseListCasesClient) Recv() (*ListCasesReply, error) {
	m := new(ListCasesReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *caseClient) GetCase(ctx context.Context, in *GetCaseRequest, opts ...grpc.CallOption) (*GetCaseReply, error) {
	out := new(GetCaseReply)
	err := c.cc.Invoke(ctx, "/case.Case/GetCase", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *caseClient) CreateCase(ctx context.Context, in *CreateCaseRequest, opts ...grpc.CallOption) (*CreateCaseReply, error) {
	out := new(CreateCaseReply)
	err := c.cc.Invoke(ctx, "/case.Case/CreateCase", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *caseClient) UpdateCase(ctx context.Context, in *UpdateCaseRequest, opts ...grpc.CallOption) (*UpdateCaseReply, error) {
	out := new(UpdateCaseReply)
	err := c.cc.Invoke(ctx, "/case.Case/UpdateCase", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *caseClient) DeleteCase(ctx context.Context, in *DeleteCaseRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/case.Case/DeleteCase", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CaseServer is the server API for Case service.
// All implementations must embed UnimplementedCaseServer
// for forward compatibility
type CaseServer interface {
	ListCases(*ListCasesRequest, Case_ListCasesServer) error
	GetCase(context.Context, *GetCaseRequest) (*GetCaseReply, error)
	CreateCase(context.Context, *CreateCaseRequest) (*CreateCaseReply, error)
	UpdateCase(context.Context, *UpdateCaseRequest) (*UpdateCaseReply, error)
	DeleteCase(context.Context, *DeleteCaseRequest) (*empty.Empty, error)
	mustEmbedUnimplementedCaseServer()
}

// UnimplementedCaseServer must be embedded to have forward compatible implementations.
type UnimplementedCaseServer struct {
}

func (UnimplementedCaseServer) ListCases(*ListCasesRequest, Case_ListCasesServer) error {
	return status.Errorf(codes.Unimplemented, "method ListCases not implemented")
}
func (UnimplementedCaseServer) GetCase(context.Context, *GetCaseRequest) (*GetCaseReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCase not implemented")
}
func (UnimplementedCaseServer) CreateCase(context.Context, *CreateCaseRequest) (*CreateCaseReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCase not implemented")
}
func (UnimplementedCaseServer) UpdateCase(context.Context, *UpdateCaseRequest) (*UpdateCaseReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCase not implemented")
}
func (UnimplementedCaseServer) DeleteCase(context.Context, *DeleteCaseRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCase not implemented")
}
func (UnimplementedCaseServer) mustEmbedUnimplementedCaseServer() {}

// UnsafeCaseServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CaseServer will
// result in compilation errors.
type UnsafeCaseServer interface {
	mustEmbedUnimplementedCaseServer()
}

func RegisterCaseServer(s grpc.ServiceRegistrar, srv CaseServer) {
	s.RegisterService(&Case_ServiceDesc, srv)
}

func _Case_ListCases_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListCasesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CaseServer).ListCases(m, &caseListCasesServer{stream})
}

type Case_ListCasesServer interface {
	Send(*ListCasesReply) error
	grpc.ServerStream
}

type caseListCasesServer struct {
	grpc.ServerStream
}

func (x *caseListCasesServer) Send(m *ListCasesReply) error {
	return x.ServerStream.SendMsg(m)
}

func _Case_GetCase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CaseServer).GetCase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/case.Case/GetCase",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CaseServer).GetCase(ctx, req.(*GetCaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Case_CreateCase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CaseServer).CreateCase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/case.Case/CreateCase",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CaseServer).CreateCase(ctx, req.(*CreateCaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Case_UpdateCase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CaseServer).UpdateCase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/case.Case/UpdateCase",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CaseServer).UpdateCase(ctx, req.(*UpdateCaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Case_DeleteCase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CaseServer).DeleteCase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/case.Case/DeleteCase",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CaseServer).DeleteCase(ctx, req.(*DeleteCaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Case_ServiceDesc is the grpc.ServiceDesc for Case service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Case_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "case.Case",
	HandlerType: (*CaseServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCase",
			Handler:    _Case_GetCase_Handler,
		},
		{
			MethodName: "CreateCase",
			Handler:    _Case_CreateCase_Handler,
		},
		{
			MethodName: "UpdateCase",
			Handler:    _Case_UpdateCase_Handler,
		},
		{
			MethodName: "DeleteCase",
			Handler:    _Case_DeleteCase_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListCases",
			Handler:       _Case_ListCases_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "internal/pkg/pb/case/case.proto",
}
