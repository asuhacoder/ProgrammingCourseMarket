// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package course

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

// CourseClient is the client API for Course service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CourseClient interface {
	ListCourses(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (Course_ListCoursesClient, error)
	GetCourse(ctx context.Context, in *GetCourseRequest, opts ...grpc.CallOption) (*GetCourseReply, error)
	CreateCourse(ctx context.Context, in *CreateCourseRequest, opts ...grpc.CallOption) (*CreateCourseReply, error)
	UpdateCourse(ctx context.Context, in *UpdateCourseRequest, opts ...grpc.CallOption) (*UpdateCourseReply, error)
	DeleteCourse(ctx context.Context, in *DeleteCourseRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type courseClient struct {
	cc grpc.ClientConnInterface
}

func NewCourseClient(cc grpc.ClientConnInterface) CourseClient {
	return &courseClient{cc}
}

func (c *courseClient) ListCourses(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (Course_ListCoursesClient, error) {
	stream, err := c.cc.NewStream(ctx, &Course_ServiceDesc.Streams[0], "/course.Course/ListCourses", opts...)
	if err != nil {
		return nil, err
	}
	x := &courseListCoursesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Course_ListCoursesClient interface {
	Recv() (*ListCoursesReply, error)
	grpc.ClientStream
}

type courseListCoursesClient struct {
	grpc.ClientStream
}

func (x *courseListCoursesClient) Recv() (*ListCoursesReply, error) {
	m := new(ListCoursesReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *courseClient) GetCourse(ctx context.Context, in *GetCourseRequest, opts ...grpc.CallOption) (*GetCourseReply, error) {
	out := new(GetCourseReply)
	err := c.cc.Invoke(ctx, "/course.Course/GetCourse", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courseClient) CreateCourse(ctx context.Context, in *CreateCourseRequest, opts ...grpc.CallOption) (*CreateCourseReply, error) {
	out := new(CreateCourseReply)
	err := c.cc.Invoke(ctx, "/course.Course/CreateCourse", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courseClient) UpdateCourse(ctx context.Context, in *UpdateCourseRequest, opts ...grpc.CallOption) (*UpdateCourseReply, error) {
	out := new(UpdateCourseReply)
	err := c.cc.Invoke(ctx, "/course.Course/UpdateCourse", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courseClient) DeleteCourse(ctx context.Context, in *DeleteCourseRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/course.Course/DeleteCourse", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CourseServer is the server API for Course service.
// All implementations must embed UnimplementedCourseServer
// for forward compatibility
type CourseServer interface {
	ListCourses(*empty.Empty, Course_ListCoursesServer) error
	GetCourse(context.Context, *GetCourseRequest) (*GetCourseReply, error)
	CreateCourse(context.Context, *CreateCourseRequest) (*CreateCourseReply, error)
	UpdateCourse(context.Context, *UpdateCourseRequest) (*UpdateCourseReply, error)
	DeleteCourse(context.Context, *DeleteCourseRequest) (*empty.Empty, error)
	mustEmbedUnimplementedCourseServer()
}

// UnimplementedCourseServer must be embedded to have forward compatible implementations.
type UnimplementedCourseServer struct {
}

func (UnimplementedCourseServer) ListCourses(*empty.Empty, Course_ListCoursesServer) error {
	return status.Errorf(codes.Unimplemented, "method ListCourses not implemented")
}
func (UnimplementedCourseServer) GetCourse(context.Context, *GetCourseRequest) (*GetCourseReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCourse not implemented")
}
func (UnimplementedCourseServer) CreateCourse(context.Context, *CreateCourseRequest) (*CreateCourseReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCourse not implemented")
}
func (UnimplementedCourseServer) UpdateCourse(context.Context, *UpdateCourseRequest) (*UpdateCourseReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCourse not implemented")
}
func (UnimplementedCourseServer) DeleteCourse(context.Context, *DeleteCourseRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCourse not implemented")
}
func (UnimplementedCourseServer) mustEmbedUnimplementedCourseServer() {}

// UnsafeCourseServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CourseServer will
// result in compilation errors.
type UnsafeCourseServer interface {
	mustEmbedUnimplementedCourseServer()
}

func RegisterCourseServer(s grpc.ServiceRegistrar, srv CourseServer) {
	s.RegisterService(&Course_ServiceDesc, srv)
}

func _Course_ListCourses_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(empty.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CourseServer).ListCourses(m, &courseListCoursesServer{stream})
}

type Course_ListCoursesServer interface {
	Send(*ListCoursesReply) error
	grpc.ServerStream
}

type courseListCoursesServer struct {
	grpc.ServerStream
}

func (x *courseListCoursesServer) Send(m *ListCoursesReply) error {
	return x.ServerStream.SendMsg(m)
}

func _Course_GetCourse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCourseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CourseServer).GetCourse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/course.Course/GetCourse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CourseServer).GetCourse(ctx, req.(*GetCourseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Course_CreateCourse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCourseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CourseServer).CreateCourse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/course.Course/CreateCourse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CourseServer).CreateCourse(ctx, req.(*CreateCourseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Course_UpdateCourse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCourseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CourseServer).UpdateCourse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/course.Course/UpdateCourse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CourseServer).UpdateCourse(ctx, req.(*UpdateCourseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Course_DeleteCourse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCourseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CourseServer).DeleteCourse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/course.Course/DeleteCourse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CourseServer).DeleteCourse(ctx, req.(*DeleteCourseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Course_ServiceDesc is the grpc.ServiceDesc for Course service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Course_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "course.Course",
	HandlerType: (*CourseServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCourse",
			Handler:    _Course_GetCourse_Handler,
		},
		{
			MethodName: "CreateCourse",
			Handler:    _Course_CreateCourse_Handler,
		},
		{
			MethodName: "UpdateCourse",
			Handler:    _Course_UpdateCourse_Handler,
		},
		{
			MethodName: "DeleteCourse",
			Handler:    _Course_DeleteCourse_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListCourses",
			Handler:       _Course_ListCourses_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "internal/pkg/pb/course/course.proto",
}
