// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api.proto

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	api.proto

It has these top-level messages:
	PainUpdate
	GetEventsRequest
	Event
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/empty"
import google_protobuf1 "github.com/golang/protobuf/ptypes/timestamp"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type PainLevel int32

const (
	PainLevel_NO_PAIN     PainLevel = 0
	PainLevel_MINOR       PainLevel = 1
	PainLevel_SIGNIFICANT PainLevel = 2
	PainLevel_SEVERE      PainLevel = 3
)

var PainLevel_name = map[int32]string{
	0: "NO_PAIN",
	1: "MINOR",
	2: "SIGNIFICANT",
	3: "SEVERE",
}
var PainLevel_value = map[string]int32{
	"NO_PAIN":     0,
	"MINOR":       1,
	"SIGNIFICANT": 2,
	"SEVERE":      3,
}

func (x PainLevel) String() string {
	return proto1.EnumName(PainLevel_name, int32(x))
}
func (PainLevel) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type PainUpdate struct {
	PainLevel PainLevel `protobuf:"varint,1,opt,name=pain_level,json=painLevel,enum=PainLevel" json:"pain_level,omitempty"`
}

func (m *PainUpdate) Reset()                    { *m = PainUpdate{} }
func (m *PainUpdate) String() string            { return proto1.CompactTextString(m) }
func (*PainUpdate) ProtoMessage()               {}
func (*PainUpdate) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *PainUpdate) GetPainLevel() PainLevel {
	if m != nil {
		return m.PainLevel
	}
	return PainLevel_NO_PAIN
}

type GetEventsRequest struct {
	Start *google_protobuf1.Timestamp `protobuf:"bytes,1,opt,name=start" json:"start,omitempty"`
	End   *google_protobuf1.Timestamp `protobuf:"bytes,2,opt,name=end" json:"end,omitempty"`
}

func (m *GetEventsRequest) Reset()                    { *m = GetEventsRequest{} }
func (m *GetEventsRequest) String() string            { return proto1.CompactTextString(m) }
func (*GetEventsRequest) ProtoMessage()               {}
func (*GetEventsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *GetEventsRequest) GetStart() *google_protobuf1.Timestamp {
	if m != nil {
		return m.Start
	}
	return nil
}

func (m *GetEventsRequest) GetEnd() *google_protobuf1.Timestamp {
	if m != nil {
		return m.End
	}
	return nil
}

type Event struct {
	Timestamp *google_protobuf1.Timestamp `protobuf:"bytes,1,opt,name=timestamp" json:"timestamp,omitempty"`
	PainLevel PainLevel                   `protobuf:"varint,2,opt,name=pain_level,json=painLevel,enum=PainLevel" json:"pain_level,omitempty"`
	Stressed  bool                        `protobuf:"varint,3,opt,name=stressed" json:"stressed,omitempty"`
	Flare     bool                        `protobuf:"varint,4,opt,name=flare" json:"flare,omitempty"`
}

func (m *Event) Reset()                    { *m = Event{} }
func (m *Event) String() string            { return proto1.CompactTextString(m) }
func (*Event) ProtoMessage()               {}
func (*Event) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Event) GetTimestamp() *google_protobuf1.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *Event) GetPainLevel() PainLevel {
	if m != nil {
		return m.PainLevel
	}
	return PainLevel_NO_PAIN
}

func (m *Event) GetStressed() bool {
	if m != nil {
		return m.Stressed
	}
	return false
}

func (m *Event) GetFlare() bool {
	if m != nil {
		return m.Flare
	}
	return false
}

func init() {
	proto1.RegisterType((*PainUpdate)(nil), "PainUpdate")
	proto1.RegisterType((*GetEventsRequest)(nil), "GetEventsRequest")
	proto1.RegisterType((*Event)(nil), "Event")
	proto1.RegisterEnum("PainLevel", PainLevel_name, PainLevel_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for PainTracker service

type PainTrackerClient interface {
	SetPainLevel(ctx context.Context, in *PainUpdate, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
	ToggleFlare(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
	ToggleStress(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
	ToggleNoPain(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
}

type painTrackerClient struct {
	cc *grpc.ClientConn
}

func NewPainTrackerClient(cc *grpc.ClientConn) PainTrackerClient {
	return &painTrackerClient{cc}
}

func (c *painTrackerClient) SetPainLevel(ctx context.Context, in *PainUpdate, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/PainTracker/SetPainLevel", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *painTrackerClient) ToggleFlare(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/PainTracker/ToggleFlare", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *painTrackerClient) ToggleStress(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/PainTracker/ToggleStress", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *painTrackerClient) ToggleNoPain(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/PainTracker/ToggleNoPain", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PainTracker service

type PainTrackerServer interface {
	SetPainLevel(context.Context, *PainUpdate) (*google_protobuf.Empty, error)
	ToggleFlare(context.Context, *google_protobuf.Empty) (*google_protobuf.Empty, error)
	ToggleStress(context.Context, *google_protobuf.Empty) (*google_protobuf.Empty, error)
	ToggleNoPain(context.Context, *google_protobuf.Empty) (*google_protobuf.Empty, error)
}

func RegisterPainTrackerServer(s *grpc.Server, srv PainTrackerServer) {
	s.RegisterService(&_PainTracker_serviceDesc, srv)
}

func _PainTracker_SetPainLevel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PainUpdate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PainTrackerServer).SetPainLevel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PainTracker/SetPainLevel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PainTrackerServer).SetPainLevel(ctx, req.(*PainUpdate))
	}
	return interceptor(ctx, in, info, handler)
}

func _PainTracker_ToggleFlare_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PainTrackerServer).ToggleFlare(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PainTracker/ToggleFlare",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PainTrackerServer).ToggleFlare(ctx, req.(*google_protobuf.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _PainTracker_ToggleStress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PainTrackerServer).ToggleStress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PainTracker/ToggleStress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PainTrackerServer).ToggleStress(ctx, req.(*google_protobuf.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _PainTracker_ToggleNoPain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PainTrackerServer).ToggleNoPain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PainTracker/ToggleNoPain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PainTrackerServer).ToggleNoPain(ctx, req.(*google_protobuf.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _PainTracker_serviceDesc = grpc.ServiceDesc{
	ServiceName: "PainTracker",
	HandlerType: (*PainTrackerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetPainLevel",
			Handler:    _PainTracker_SetPainLevel_Handler,
		},
		{
			MethodName: "ToggleFlare",
			Handler:    _PainTracker_ToggleFlare_Handler,
		},
		{
			MethodName: "ToggleStress",
			Handler:    _PainTracker_ToggleStress_Handler,
		},
		{
			MethodName: "ToggleNoPain",
			Handler:    _PainTracker_ToggleNoPain_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}

// Client API for Monitor service

type MonitorClient interface {
	GetEvents(ctx context.Context, in *GetEventsRequest, opts ...grpc.CallOption) (Monitor_GetEventsClient, error)
}

type monitorClient struct {
	cc *grpc.ClientConn
}

func NewMonitorClient(cc *grpc.ClientConn) MonitorClient {
	return &monitorClient{cc}
}

func (c *monitorClient) GetEvents(ctx context.Context, in *GetEventsRequest, opts ...grpc.CallOption) (Monitor_GetEventsClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Monitor_serviceDesc.Streams[0], c.cc, "/Monitor/GetEvents", opts...)
	if err != nil {
		return nil, err
	}
	x := &monitorGetEventsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Monitor_GetEventsClient interface {
	Recv() (*Event, error)
	grpc.ClientStream
}

type monitorGetEventsClient struct {
	grpc.ClientStream
}

func (x *monitorGetEventsClient) Recv() (*Event, error) {
	m := new(Event)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Monitor service

type MonitorServer interface {
	GetEvents(*GetEventsRequest, Monitor_GetEventsServer) error
}

func RegisterMonitorServer(s *grpc.Server, srv MonitorServer) {
	s.RegisterService(&_Monitor_serviceDesc, srv)
}

func _Monitor_GetEvents_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetEventsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MonitorServer).GetEvents(m, &monitorGetEventsServer{stream})
}

type Monitor_GetEventsServer interface {
	Send(*Event) error
	grpc.ServerStream
}

type monitorGetEventsServer struct {
	grpc.ServerStream
}

func (x *monitorGetEventsServer) Send(m *Event) error {
	return x.ServerStream.SendMsg(m)
}

var _Monitor_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Monitor",
	HandlerType: (*MonitorServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetEvents",
			Handler:       _Monitor_GetEvents_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "api.proto",
}

func init() { proto1.RegisterFile("api.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 505 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x93, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0x71, 0x42, 0xd2, 0x7a, 0x1c, 0x5a, 0x77, 0x81, 0x12, 0x0c, 0x12, 0x95, 0x4f, 0x25,
	0xa2, 0x76, 0x09, 0x42, 0x45, 0x3d, 0x51, 0x2a, 0xb7, 0xb2, 0xd4, 0xba, 0x95, 0x63, 0x7a, 0xe0,
	0x52, 0x6d, 0x92, 0xad, 0x63, 0x70, 0x76, 0xcd, 0xee, 0x26, 0x12, 0x57, 0x5e, 0x81, 0x3b, 0x2f,
	0xc5, 0x13, 0x20, 0xf1, 0x20, 0x68, 0xd7, 0x71, 0x82, 0x82, 0xca, 0x1f, 0xf5, 0x64, 0x8f, 0xe7,
	0xdb, 0xdf, 0xec, 0x37, 0x33, 0x06, 0x13, 0x17, 0x99, 0x57, 0x70, 0x26, 0x99, 0xf3, 0x28, 0x65,
	0x2c, 0xcd, 0x89, 0xaf, 0xa3, 0xfe, 0xe4, 0xca, 0x27, 0xe3, 0x42, 0x7e, 0x9a, 0x25, 0x9f, 0x2c,
	0x27, 0x65, 0x36, 0x26, 0x42, 0xe2, 0x71, 0x31, 0x13, 0x3c, 0x9e, 0x09, 0x70, 0x91, 0xf9, 0x98,
	0x52, 0x26, 0xb1, 0xcc, 0x18, 0x15, 0x65, 0xd6, 0xdd, 0x03, 0x38, 0xc7, 0x19, 0x7d, 0x5b, 0x0c,
	0xb1, 0x24, 0xe8, 0x29, 0x40, 0x81, 0x33, 0x7a, 0x99, 0x93, 0x29, 0xc9, 0xdb, 0xc6, 0x96, 0xb1,
	0xbd, 0xd6, 0x05, 0x4f, 0x09, 0x4e, 0xd4, 0x97, 0xd8, 0x2c, 0xaa, 0x57, 0x97, 0x83, 0x7d, 0x4c,
	0x64, 0x30, 0x25, 0x54, 0x8a, 0x98, 0x7c, 0x9c, 0x10, 0x21, 0xd1, 0x2e, 0x34, 0x84, 0xc4, 0x5c,
	0xea, 0x93, 0x56, 0xd7, 0xf1, 0xca, 0xd2, 0x5e, 0x75, 0x37, 0x2f, 0xa9, 0xee, 0x16, 0x97, 0x42,
	0xf4, 0x0c, 0xea, 0x84, 0x0e, 0xdb, 0xb5, 0xbf, 0xea, 0x95, 0xcc, 0xfd, 0x6a, 0x40, 0x43, 0x57,
	0x44, 0xaf, 0xc0, 0x9c, 0xfb, 0xfc, 0x87, 0x6a, 0x0b, 0xf1, 0x92, 0xc5, 0xda, 0x1f, 0x2c, 0x22,
	0x07, 0x56, 0x85, 0xe4, 0x44, 0x08, 0x32, 0x6c, 0xd7, 0xb7, 0x8c, 0xed, 0xd5, 0x78, 0x1e, 0xa3,
	0x7b, 0xd0, 0xb8, 0xca, 0x31, 0x27, 0xed, 0xdb, 0x3a, 0x51, 0x06, 0x9d, 0xd7, 0x60, 0xce, 0x49,
	0xc8, 0x82, 0x95, 0xe8, 0xec, 0xf2, 0xfc, 0x20, 0x8c, 0xec, 0x5b, 0xc8, 0x84, 0xc6, 0x69, 0x18,
	0x9d, 0xc5, 0xb6, 0x81, 0xd6, 0xc1, 0xea, 0x85, 0xc7, 0x51, 0x78, 0x14, 0x1e, 0x1e, 0x44, 0x89,
	0x5d, 0x43, 0x00, 0xcd, 0x5e, 0x70, 0x11, 0xc4, 0x81, 0x5d, 0xef, 0x7e, 0xaf, 0x81, 0xa5, 0x10,
	0x09, 0xc7, 0x83, 0x0f, 0x84, 0xa3, 0x13, 0x68, 0xf5, 0x88, 0xfc, 0x05, 0xea, 0x2d, 0xc6, 0xe5,
	0x6c, 0xfe, 0x66, 0x39, 0x50, 0x9b, 0xe1, 0x3e, 0xf8, 0xfc, 0xed, 0xc7, 0x97, 0xda, 0x86, 0xdb,
	0xd2, 0x43, 0x9f, 0x3e, 0xf7, 0x95, 0xa7, 0x7d, 0xa3, 0x83, 0x12, 0xb0, 0x12, 0x96, 0xa6, 0x39,
	0x39, 0x52, 0xd7, 0x45, 0xd7, 0x9c, 0xbf, 0x96, 0x7b, 0x5f, 0x73, 0xd7, 0xdd, 0x3b, 0x15, 0x57,
	0xbb, 0x46, 0x17, 0xd0, 0x2a, 0xa9, 0x3d, 0xdd, 0x9d, 0xff, 0xc6, 0x6e, 0x6a, 0xac, 0xed, 0xae,
	0x55, 0xd8, 0xb2, 0xcb, 0x0b, 0x6e, 0xc4, 0x94, 0xe9, 0x9b, 0x73, 0x29, 0x53, 0x8d, 0xe8, 0x46,
	0xb0, 0x72, 0xca, 0x68, 0x26, 0x19, 0x47, 0x87, 0x60, 0xce, 0xb7, 0x18, 0x6d, 0x78, 0xcb, 0x1b,
	0xed, 0x34, 0x3d, 0x1d, 0xbb, 0x0f, 0x35, 0xea, 0xee, 0x02, 0x45, 0xb4, 0x6c, 0xdf, 0xe8, 0xec,
	0x1a, 0x6f, 0xf6, 0xde, 0xbd, 0x4c, 0x33, 0x39, 0x9a, 0xf4, 0xbd, 0x01, 0x1b, 0xfb, 0xef, 0xd9,
	0x08, 0xd3, 0x3e, 0xc7, 0x74, 0x38, 0x62, 0x5c, 0x48, 0x7f, 0x30, 0xe2, 0x8c, 0x66, 0x83, 0x1d,
	0x55, 0x7a, 0x47, 0x96, 0x63, 0x9d, 0xfd, 0xad, 0x4d, 0xfd, 0x78, 0xf1, 0x33, 0x00, 0x00, 0xff,
	0xff, 0xa8, 0x72, 0xb1, 0x32, 0xec, 0x03, 0x00, 0x00,
}
