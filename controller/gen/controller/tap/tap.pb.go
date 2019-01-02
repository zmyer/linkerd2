// Code generated by protoc-gen-go. DO NOT EDIT.
// source: controller/tap.proto

package tap

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	public "github.com/linkerd/linkerd2/controller/gen/public"
	math "math"
)

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TapClient is the client API for Tap service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TapClient interface {
	Tap(ctx context.Context, in *public.TapRequest, opts ...grpc.CallOption) (Tap_TapClient, error)
	TapByResource(ctx context.Context, in *public.TapByResourceRequest, opts ...grpc.CallOption) (Tap_TapByResourceClient, error)
}

type tapClient struct {
	cc *grpc.ClientConn
}

func NewTapClient(cc *grpc.ClientConn) TapClient {
	return &tapClient{cc}
}

// Deprecated: Do not use.
func (c *tapClient) Tap(ctx context.Context, in *public.TapRequest, opts ...grpc.CallOption) (Tap_TapClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Tap_serviceDesc.Streams[0], "/linkerd2.controller.tap.Tap/Tap", opts...)
	if err != nil {
		return nil, err
	}
	x := &tapTapClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Tap_TapClient interface {
	Recv() (*public.TapEvent, error)
	grpc.ClientStream
}

type tapTapClient struct {
	grpc.ClientStream
}

func (x *tapTapClient) Recv() (*public.TapEvent, error) {
	m := new(public.TapEvent)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *tapClient) TapByResource(ctx context.Context, in *public.TapByResourceRequest, opts ...grpc.CallOption) (Tap_TapByResourceClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Tap_serviceDesc.Streams[1], "/linkerd2.controller.tap.Tap/TapByResource", opts...)
	if err != nil {
		return nil, err
	}
	x := &tapTapByResourceClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Tap_TapByResourceClient interface {
	Recv() (*public.TapEvent, error)
	grpc.ClientStream
}

type tapTapByResourceClient struct {
	grpc.ClientStream
}

func (x *tapTapByResourceClient) Recv() (*public.TapEvent, error) {
	m := new(public.TapEvent)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TapServer is the server API for Tap service.
type TapServer interface {
	Tap(*public.TapRequest, Tap_TapServer) error
	TapByResource(*public.TapByResourceRequest, Tap_TapByResourceServer) error
}

func RegisterTapServer(s *grpc.Server, srv TapServer) {
	s.RegisterService(&_Tap_serviceDesc, srv)
}

func _Tap_Tap_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(public.TapRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TapServer).Tap(m, &tapTapServer{stream})
}

type Tap_TapServer interface {
	Send(*public.TapEvent) error
	grpc.ServerStream
}

type tapTapServer struct {
	grpc.ServerStream
}

func (x *tapTapServer) Send(m *public.TapEvent) error {
	return x.ServerStream.SendMsg(m)
}

func _Tap_TapByResource_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(public.TapByResourceRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TapServer).TapByResource(m, &tapTapByResourceServer{stream})
}

type Tap_TapByResourceServer interface {
	Send(*public.TapEvent) error
	grpc.ServerStream
}

type tapTapByResourceServer struct {
	grpc.ServerStream
}

func (x *tapTapByResourceServer) Send(m *public.TapEvent) error {
	return x.ServerStream.SendMsg(m)
}

var _Tap_serviceDesc = grpc.ServiceDesc{
	ServiceName: "linkerd2.controller.tap.Tap",
	HandlerType: (*TapServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Tap",
			Handler:       _Tap_Tap_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "TapByResource",
			Handler:       _Tap_TapByResource_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "controller/tap.proto",
}

func init() { proto.RegisterFile("controller/tap.proto", fileDescriptor_d1a304d4ff6ff27b) }

var fileDescriptor_d1a304d4ff6ff27b = []byte{
	// 180 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x49, 0xce, 0xcf, 0x2b,
	0x29, 0xca, 0xcf, 0xc9, 0x49, 0x2d, 0xd2, 0x2f, 0x49, 0x2c, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0x12, 0xcf, 0xc9, 0xcc, 0xcb, 0x4e, 0x2d, 0x4a, 0x31, 0xd2, 0x43, 0x48, 0xeb, 0x95, 0x24,
	0x16, 0x48, 0xf1, 0x14, 0x94, 0x26, 0xe5, 0x64, 0x26, 0x43, 0x94, 0x19, 0x2d, 0x62, 0xe4, 0x62,
	0x0e, 0x49, 0x2c, 0x10, 0x72, 0x81, 0x50, 0xd2, 0x7a, 0x70, 0x6d, 0x50, 0x65, 0x21, 0x89, 0x05,
	0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x52, 0x92, 0xd8, 0x24, 0x5d, 0xcb, 0x52, 0xf3, 0x4a,
	0x94, 0x98, 0x3b, 0x98, 0x18, 0x0d, 0x18, 0x85, 0x42, 0xb9, 0x78, 0x43, 0x12, 0x0b, 0x9c, 0x2a,
	0x83, 0x52, 0x8b, 0xf3, 0x4b, 0x8b, 0x92, 0x53, 0x85, 0x54, 0xb1, 0x69, 0x41, 0xc8, 0x13, 0x61,
	0x32, 0x83, 0x01, 0xa3, 0x93, 0x75, 0x94, 0x65, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92, 0x5e, 0x72,
	0x7e, 0xae, 0x3e, 0x54, 0x29, 0x8c, 0x36, 0xd2, 0x47, 0xf2, 0x7f, 0x7a, 0x6a, 0x9e, 0x3e, 0x6a,
	0x70, 0x24, 0xb1, 0x81, 0x3d, 0x6a, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xb2, 0x4e, 0xdc, 0xd1,
	0x27, 0x01, 0x00, 0x00,
}
