// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/api/v2/eds.proto

package envoy_api_v2

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import wrappers "github.com/golang/protobuf/ptypes/wrappers"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import discovery "google.golang.org/grpc/balancer/xds/internal/proto/envoy/api/v2/discovery"
import endpoint "google.golang.org/grpc/balancer/xds/internal/proto/envoy/api/v2/endpoint/endpoint"
import percent "google.golang.org/grpc/balancer/xds/internal/proto/envoy/type/percent"
import _ "google.golang.org/grpc/balancer/xds/internal/proto/validate"

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

type ClusterLoadAssignment struct {
	ClusterName          string                          `protobuf:"bytes,1,opt,name=cluster_name,json=clusterName,proto3" json:"cluster_name,omitempty"`
	Endpoints            []*endpoint.LocalityLbEndpoints `protobuf:"bytes,2,rep,name=endpoints,proto3" json:"endpoints,omitempty"`
	NamedEndpoints       map[string]*endpoint.Endpoint   `protobuf:"bytes,5,rep,name=named_endpoints,json=namedEndpoints,proto3" json:"named_endpoints,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Policy               *ClusterLoadAssignment_Policy   `protobuf:"bytes,4,opt,name=policy,proto3" json:"policy,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *ClusterLoadAssignment) Reset()         { *m = ClusterLoadAssignment{} }
func (m *ClusterLoadAssignment) String() string { return proto.CompactTextString(m) }
func (*ClusterLoadAssignment) ProtoMessage()    {}
func (*ClusterLoadAssignment) Descriptor() ([]byte, []int) {
	return fileDescriptor_eds_fb0a999149ff4153, []int{0}
}
func (m *ClusterLoadAssignment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterLoadAssignment.Unmarshal(m, b)
}
func (m *ClusterLoadAssignment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterLoadAssignment.Marshal(b, m, deterministic)
}
func (dst *ClusterLoadAssignment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterLoadAssignment.Merge(dst, src)
}
func (m *ClusterLoadAssignment) XXX_Size() int {
	return xxx_messageInfo_ClusterLoadAssignment.Size(m)
}
func (m *ClusterLoadAssignment) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterLoadAssignment.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterLoadAssignment proto.InternalMessageInfo

func (m *ClusterLoadAssignment) GetClusterName() string {
	if m != nil {
		return m.ClusterName
	}
	return ""
}

func (m *ClusterLoadAssignment) GetEndpoints() []*endpoint.LocalityLbEndpoints {
	if m != nil {
		return m.Endpoints
	}
	return nil
}

func (m *ClusterLoadAssignment) GetNamedEndpoints() map[string]*endpoint.Endpoint {
	if m != nil {
		return m.NamedEndpoints
	}
	return nil
}

func (m *ClusterLoadAssignment) GetPolicy() *ClusterLoadAssignment_Policy {
	if m != nil {
		return m.Policy
	}
	return nil
}

type ClusterLoadAssignment_Policy struct {
	DropOverloads          []*ClusterLoadAssignment_Policy_DropOverload `protobuf:"bytes,2,rep,name=drop_overloads,json=dropOverloads,proto3" json:"drop_overloads,omitempty"`
	OverprovisioningFactor *wrappers.UInt32Value                        `protobuf:"bytes,3,opt,name=overprovisioning_factor,json=overprovisioningFactor,proto3" json:"overprovisioning_factor,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}                                     `json:"-"`
	XXX_unrecognized       []byte                                       `json:"-"`
	XXX_sizecache          int32                                        `json:"-"`
}

func (m *ClusterLoadAssignment_Policy) Reset()         { *m = ClusterLoadAssignment_Policy{} }
func (m *ClusterLoadAssignment_Policy) String() string { return proto.CompactTextString(m) }
func (*ClusterLoadAssignment_Policy) ProtoMessage()    {}
func (*ClusterLoadAssignment_Policy) Descriptor() ([]byte, []int) {
	return fileDescriptor_eds_fb0a999149ff4153, []int{0, 1}
}
func (m *ClusterLoadAssignment_Policy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterLoadAssignment_Policy.Unmarshal(m, b)
}
func (m *ClusterLoadAssignment_Policy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterLoadAssignment_Policy.Marshal(b, m, deterministic)
}
func (dst *ClusterLoadAssignment_Policy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterLoadAssignment_Policy.Merge(dst, src)
}
func (m *ClusterLoadAssignment_Policy) XXX_Size() int {
	return xxx_messageInfo_ClusterLoadAssignment_Policy.Size(m)
}
func (m *ClusterLoadAssignment_Policy) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterLoadAssignment_Policy.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterLoadAssignment_Policy proto.InternalMessageInfo

func (m *ClusterLoadAssignment_Policy) GetDropOverloads() []*ClusterLoadAssignment_Policy_DropOverload {
	if m != nil {
		return m.DropOverloads
	}
	return nil
}

func (m *ClusterLoadAssignment_Policy) GetOverprovisioningFactor() *wrappers.UInt32Value {
	if m != nil {
		return m.OverprovisioningFactor
	}
	return nil
}

type ClusterLoadAssignment_Policy_DropOverload struct {
	Category             string                     `protobuf:"bytes,1,opt,name=category,proto3" json:"category,omitempty"`
	DropPercentage       *percent.FractionalPercent `protobuf:"bytes,2,opt,name=drop_percentage,json=dropPercentage,proto3" json:"drop_percentage,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *ClusterLoadAssignment_Policy_DropOverload) Reset() {
	*m = ClusterLoadAssignment_Policy_DropOverload{}
}
func (m *ClusterLoadAssignment_Policy_DropOverload) String() string { return proto.CompactTextString(m) }
func (*ClusterLoadAssignment_Policy_DropOverload) ProtoMessage()    {}
func (*ClusterLoadAssignment_Policy_DropOverload) Descriptor() ([]byte, []int) {
	return fileDescriptor_eds_fb0a999149ff4153, []int{0, 1, 0}
}
func (m *ClusterLoadAssignment_Policy_DropOverload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterLoadAssignment_Policy_DropOverload.Unmarshal(m, b)
}
func (m *ClusterLoadAssignment_Policy_DropOverload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterLoadAssignment_Policy_DropOverload.Marshal(b, m, deterministic)
}
func (dst *ClusterLoadAssignment_Policy_DropOverload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterLoadAssignment_Policy_DropOverload.Merge(dst, src)
}
func (m *ClusterLoadAssignment_Policy_DropOverload) XXX_Size() int {
	return xxx_messageInfo_ClusterLoadAssignment_Policy_DropOverload.Size(m)
}
func (m *ClusterLoadAssignment_Policy_DropOverload) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterLoadAssignment_Policy_DropOverload.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterLoadAssignment_Policy_DropOverload proto.InternalMessageInfo

func (m *ClusterLoadAssignment_Policy_DropOverload) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

func (m *ClusterLoadAssignment_Policy_DropOverload) GetDropPercentage() *percent.FractionalPercent {
	if m != nil {
		return m.DropPercentage
	}
	return nil
}

func init() {
	proto.RegisterType((*ClusterLoadAssignment)(nil), "envoy.api.v2.ClusterLoadAssignment")
	proto.RegisterMapType((map[string]*endpoint.Endpoint)(nil), "envoy.api.v2.ClusterLoadAssignment.NamedEndpointsEntry")
	proto.RegisterType((*ClusterLoadAssignment_Policy)(nil), "envoy.api.v2.ClusterLoadAssignment.Policy")
	proto.RegisterType((*ClusterLoadAssignment_Policy_DropOverload)(nil), "envoy.api.v2.ClusterLoadAssignment.Policy.DropOverload")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// EndpointDiscoveryServiceClient is the client API for EndpointDiscoveryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EndpointDiscoveryServiceClient interface {
	StreamEndpoints(ctx context.Context, opts ...grpc.CallOption) (EndpointDiscoveryService_StreamEndpointsClient, error)
	FetchEndpoints(ctx context.Context, in *discovery.DiscoveryRequest, opts ...grpc.CallOption) (*discovery.DiscoveryResponse, error)
}

type endpointDiscoveryServiceClient struct {
	cc *grpc.ClientConn
}

func NewEndpointDiscoveryServiceClient(cc *grpc.ClientConn) EndpointDiscoveryServiceClient {
	return &endpointDiscoveryServiceClient{cc}
}

func (c *endpointDiscoveryServiceClient) StreamEndpoints(ctx context.Context, opts ...grpc.CallOption) (EndpointDiscoveryService_StreamEndpointsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_EndpointDiscoveryService_serviceDesc.Streams[0], "/envoy.api.v2.EndpointDiscoveryService/StreamEndpoints", opts...)
	if err != nil {
		return nil, err
	}
	x := &endpointDiscoveryServiceStreamEndpointsClient{stream}
	return x, nil
}

type EndpointDiscoveryService_StreamEndpointsClient interface {
	Send(*discovery.DiscoveryRequest) error
	Recv() (*discovery.DiscoveryResponse, error)
	grpc.ClientStream
}

type endpointDiscoveryServiceStreamEndpointsClient struct {
	grpc.ClientStream
}

func (x *endpointDiscoveryServiceStreamEndpointsClient) Send(m *discovery.DiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *endpointDiscoveryServiceStreamEndpointsClient) Recv() (*discovery.DiscoveryResponse, error) {
	m := new(discovery.DiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *endpointDiscoveryServiceClient) FetchEndpoints(ctx context.Context, in *discovery.DiscoveryRequest, opts ...grpc.CallOption) (*discovery.DiscoveryResponse, error) {
	out := new(discovery.DiscoveryResponse)
	err := c.cc.Invoke(ctx, "/envoy.api.v2.EndpointDiscoveryService/FetchEndpoints", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EndpointDiscoveryServiceServer is the server API for EndpointDiscoveryService service.
type EndpointDiscoveryServiceServer interface {
	StreamEndpoints(EndpointDiscoveryService_StreamEndpointsServer) error
	FetchEndpoints(context.Context, *discovery.DiscoveryRequest) (*discovery.DiscoveryResponse, error)
}

func RegisterEndpointDiscoveryServiceServer(s *grpc.Server, srv EndpointDiscoveryServiceServer) {
	s.RegisterService(&_EndpointDiscoveryService_serviceDesc, srv)
}

func _EndpointDiscoveryService_StreamEndpoints_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(EndpointDiscoveryServiceServer).StreamEndpoints(&endpointDiscoveryServiceStreamEndpointsServer{stream})
}

type EndpointDiscoveryService_StreamEndpointsServer interface {
	Send(*discovery.DiscoveryResponse) error
	Recv() (*discovery.DiscoveryRequest, error)
	grpc.ServerStream
}

type endpointDiscoveryServiceStreamEndpointsServer struct {
	grpc.ServerStream
}

func (x *endpointDiscoveryServiceStreamEndpointsServer) Send(m *discovery.DiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *endpointDiscoveryServiceStreamEndpointsServer) Recv() (*discovery.DiscoveryRequest, error) {
	m := new(discovery.DiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _EndpointDiscoveryService_FetchEndpoints_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(discovery.DiscoveryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EndpointDiscoveryServiceServer).FetchEndpoints(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/envoy.api.v2.EndpointDiscoveryService/FetchEndpoints",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EndpointDiscoveryServiceServer).FetchEndpoints(ctx, req.(*discovery.DiscoveryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _EndpointDiscoveryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.api.v2.EndpointDiscoveryService",
	HandlerType: (*EndpointDiscoveryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchEndpoints",
			Handler:    _EndpointDiscoveryService_FetchEndpoints_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamEndpoints",
			Handler:       _EndpointDiscoveryService_StreamEndpoints_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "envoy/api/v2/eds.proto",
}

func init() { proto.RegisterFile("envoy/api/v2/eds.proto", fileDescriptor_eds_fb0a999149ff4153) }

var fileDescriptor_eds_fb0a999149ff4153 = []byte{
	// 612 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xee, 0x3a, 0x6d, 0xd5, 0x6e, 0x43, 0x52, 0x2d, 0xa2, 0xb1, 0xac, 0xd0, 0x46, 0x11, 0x48,
	0x51, 0x40, 0x36, 0x4a, 0x85, 0x8a, 0x72, 0x23, 0x34, 0x11, 0xa0, 0x08, 0x22, 0x57, 0x20, 0x4e,
	0x84, 0x8d, 0xbd, 0x35, 0x2b, 0x9c, 0xdd, 0x65, 0xbd, 0x31, 0xf8, 0xc0, 0x85, 0x13, 0x77, 0xde,
	0x86, 0x13, 0x6f, 0xc0, 0x81, 0x57, 0xe0, 0x82, 0xfa, 0x12, 0xc8, 0xbf, 0x89, 0x69, 0x2a, 0x71,
	0xe0, 0xb6, 0xde, 0x99, 0xef, 0x9b, 0x6f, 0xbe, 0x9d, 0x31, 0x3c, 0x20, 0x2c, 0xe4, 0x91, 0x85,
	0x05, 0xb5, 0xc2, 0x9e, 0x45, 0xdc, 0xc0, 0x14, 0x92, 0x2b, 0x8e, 0xaa, 0xc9, 0xbd, 0x89, 0x05,
	0x35, 0xc3, 0x9e, 0xd1, 0x2c, 0x65, 0xb9, 0x34, 0x70, 0x78, 0x48, 0x64, 0x94, 0xe6, 0x1a, 0xb7,
	0xca, 0x1c, 0xcc, 0x15, 0x9c, 0x32, 0x55, 0x1c, 0xb2, 0x2c, 0x3d, 0xcd, 0x52, 0x91, 0x20, 0x96,
	0x20, 0xd2, 0x21, 0x45, 0xa4, 0xe9, 0x71, 0xee, 0xf9, 0x24, 0x21, 0xc0, 0x8c, 0x71, 0x85, 0x15,
	0xe5, 0x2c, 0x53, 0x62, 0x34, 0x42, 0xec, 0x53, 0x17, 0x2b, 0x62, 0xe5, 0x87, 0x2c, 0x70, 0x98,
	0xc1, 0x92, 0xaf, 0xd9, 0xe2, 0xdc, 0xfa, 0x20, 0xb1, 0x10, 0x44, 0x66, 0xc0, 0xf6, 0xc5, 0x16,
	0xbc, 0xf1, 0xc8, 0x5f, 0x04, 0x8a, 0xc8, 0x31, 0xc7, 0xee, 0xc3, 0x20, 0xa0, 0x1e, 0x9b, 0x13,
	0xa6, 0xd0, 0x5d, 0x58, 0x75, 0xd2, 0xc0, 0x94, 0xe1, 0x39, 0xd1, 0x41, 0x0b, 0x74, 0x76, 0x07,
	0xbb, 0xdf, 0x7e, 0x7f, 0xaf, 0x6c, 0x4a, 0xad, 0x05, 0xec, 0xbd, 0x2c, 0xfc, 0x0c, 0xcf, 0x09,
	0x7a, 0x0c, 0x77, 0xf3, 0x56, 0x02, 0x5d, 0x6b, 0x55, 0x3a, 0x7b, 0xbd, 0xae, 0xb9, 0x6a, 0x8f,
	0x59, 0x74, 0x3a, 0xe6, 0x0e, 0xf6, 0xa9, 0x8a, 0xc6, 0xb3, 0x61, 0x8e, 0xb0, 0x97, 0x60, 0xf4,
	0x06, 0xd6, 0xe3, 0x7a, 0xee, 0x74, 0xc9, 0xb7, 0x95, 0xf0, 0x9d, 0x94, 0xf9, 0xd6, 0xaa, 0x36,
	0x63, 0x31, 0x6e, 0xc1, 0x3b, 0x64, 0x4a, 0x46, 0x76, 0x8d, 0x95, 0x2e, 0xd1, 0x00, 0x6e, 0x0b,
	0xee, 0x53, 0x27, 0xd2, 0x37, 0x5b, 0xe0, 0xb2, 0xd0, 0xf5, 0xc4, 0x93, 0x04, 0x61, 0x67, 0x48,
	0x63, 0x06, 0xaf, 0xaf, 0x29, 0x85, 0xf6, 0x61, 0xe5, 0x1d, 0x89, 0x52, 0xaf, 0xec, 0xf8, 0x88,
	0xee, 0xc3, 0xad, 0x10, 0xfb, 0x0b, 0xa2, 0x6b, 0x49, 0xad, 0xa3, 0x2b, 0x4c, 0xc9, 0x79, 0xec,
	0x34, 0xbb, 0xaf, 0x3d, 0x00, 0xc6, 0x0f, 0x0d, 0x6e, 0xa7, 0x65, 0xd1, 0x6b, 0x58, 0x73, 0x25,
	0x17, 0xd3, 0x78, 0xa2, 0x7c, 0x8e, 0xdd, 0xdc, 0xe3, 0x93, 0x7f, 0x97, 0x6e, 0x9e, 0x4a, 0x2e,
	0x9e, 0x67, 0x78, 0xfb, 0x9a, 0xbb, 0xf2, 0x15, 0x9b, 0xde, 0x88, 0xa9, 0x85, 0xe4, 0x21, 0x0d,
	0x28, 0x67, 0x94, 0x79, 0xd3, 0x73, 0xec, 0x28, 0x2e, 0xf5, 0x4a, 0xa2, 0xbb, 0x69, 0xa6, 0x83,
	0x64, 0xe6, 0x83, 0x64, 0xbe, 0x78, 0xc2, 0xd4, 0x71, 0xef, 0x65, 0xac, 0x36, 0x9b, 0x8a, 0xae,
	0xd6, 0xda, 0xb0, 0x0f, 0xfe, 0xe6, 0x19, 0x25, 0x34, 0xc6, 0x27, 0x58, 0x5d, 0x15, 0x80, 0x6e,
	0xc3, 0x1d, 0x07, 0x2b, 0xe2, 0x71, 0x19, 0x5d, 0x1e, 0xad, 0x22, 0x84, 0x46, 0xb0, 0x9e, 0x34,
	0x9e, 0x2d, 0x03, 0xf6, 0x72, 0x23, 0x6f, 0x66, 0x9d, 0xc7, 0xab, 0x62, 0x8e, 0x24, 0x76, 0xe2,
	0x75, 0xc0, 0xfe, 0x24, 0xcd, 0xb3, 0x13, 0xbb, 0x26, 0x05, 0xe8, 0xe9, 0xe6, 0x0e, 0xd8, 0xd7,
	0x7a, 0x17, 0x00, 0xea, 0xb9, 0xd3, 0xa7, 0xf9, 0x82, 0x9e, 0x11, 0x19, 0x52, 0x87, 0xa0, 0x57,
	0xb0, 0x7e, 0xa6, 0x24, 0xc1, 0xf3, 0xe5, 0xa4, 0x1c, 0x96, 0xed, 0x2d, 0x20, 0x36, 0x79, 0xbf,
	0x20, 0x81, 0x32, 0x8e, 0xae, 0x8c, 0x07, 0x82, 0xb3, 0x80, 0xb4, 0x37, 0x3a, 0xe0, 0x1e, 0x40,
	0x0b, 0x58, 0x1b, 0x11, 0xe5, 0xbc, 0xfd, 0x8f, 0xc4, 0xed, 0xcf, 0x3f, 0x7f, 0x7d, 0xd5, 0x9a,
	0xed, 0x46, 0xe9, 0x5f, 0xd3, 0x2f, 0x76, 0xa6, 0x0f, 0xba, 0x83, 0x3b, 0xd0, 0xa0, 0x3c, 0x25,
	0x12, 0x92, 0x7f, 0x8c, 0x4a, 0x9c, 0x83, 0x9d, 0xa1, 0x1b, 0x4c, 0xe2, 0xc7, 0x9c, 0x80, 0x2f,
	0x00, 0xcc, 0xb6, 0x93, 0x87, 0x3d, 0xfe, 0x13, 0x00, 0x00, 0xff, 0xff, 0xe7, 0x56, 0x21, 0x69,
	0xec, 0x04, 0x00, 0x00,
}
