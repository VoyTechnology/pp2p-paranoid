// Code generated by protoc-gen-go.
// source: discoverynetwork/discovery.proto
// DO NOT EDIT!

/*
Package discoverynetwork is a generated protocol buffer package.

It is generated from these files:
	discoverynetwork/discovery.proto

It has these top-level messages:
	EmptyMessage
	DisconnectRequest
	JoinRequest
	JoinResponse
	Node
*/
package discoverynetwork

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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
const _ = proto.ProtoPackageIsVersion1

type EmptyMessage struct {
}

func (m *EmptyMessage) Reset()                    { *m = EmptyMessage{} }
func (m *EmptyMessage) String() string            { return proto.CompactTextString(m) }
func (*EmptyMessage) ProtoMessage()               {}
func (*EmptyMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type DisconnectRequest struct {
	Pool     string `protobuf:"bytes,1,opt,name=pool" json:"pool,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
	Node     *Node  `protobuf:"bytes,3,opt,name=node" json:"node,omitempty"`
}

func (m *DisconnectRequest) Reset()                    { *m = DisconnectRequest{} }
func (m *DisconnectRequest) String() string            { return proto.CompactTextString(m) }
func (*DisconnectRequest) ProtoMessage()               {}
func (*DisconnectRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *DisconnectRequest) GetNode() *Node {
	if m != nil {
		return m.Node
	}
	return nil
}

type JoinRequest struct {
	Pool     string `protobuf:"bytes,1,opt,name=pool" json:"pool,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
	Node     *Node  `protobuf:"bytes,3,opt,name=node" json:"node,omitempty"`
}

func (m *JoinRequest) Reset()                    { *m = JoinRequest{} }
func (m *JoinRequest) String() string            { return proto.CompactTextString(m) }
func (*JoinRequest) ProtoMessage()               {}
func (*JoinRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *JoinRequest) GetNode() *Node {
	if m != nil {
		return m.Node
	}
	return nil
}

type JoinResponse struct {
	// The time after which the server will remove the node from the list
	// if the node doesn't make the Renew RPC call
	ResetInterval int64   `protobuf:"varint,1,opt,name=reset_interval" json:"reset_interval,omitempty"`
	Nodes         []*Node `protobuf:"bytes,2,rep,name=nodes" json:"nodes,omitempty"`
}

func (m *JoinResponse) Reset()                    { *m = JoinResponse{} }
func (m *JoinResponse) String() string            { return proto.CompactTextString(m) }
func (*JoinResponse) ProtoMessage()               {}
func (*JoinResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *JoinResponse) GetNodes() []*Node {
	if m != nil {
		return m.Nodes
	}
	return nil
}

type Node struct {
	Ip         string `protobuf:"bytes,1,opt,name=ip" json:"ip,omitempty"`
	Port       string `protobuf:"bytes,2,opt,name=port" json:"port,omitempty"`
	CommonName string `protobuf:"bytes,3,opt,name=common_name" json:"common_name,omitempty"`
	Uuid       string `protobuf:"bytes,4,opt,name=uuid" json:"uuid,omitempty"`
}

func (m *Node) Reset()                    { *m = Node{} }
func (m *Node) String() string            { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()               {}
func (*Node) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func init() {
	proto.RegisterType((*EmptyMessage)(nil), "discoverynetwork.EmptyMessage")
	proto.RegisterType((*DisconnectRequest)(nil), "discoverynetwork.DisconnectRequest")
	proto.RegisterType((*JoinRequest)(nil), "discoverynetwork.JoinRequest")
	proto.RegisterType((*JoinResponse)(nil), "discoverynetwork.JoinResponse")
	proto.RegisterType((*Node)(nil), "discoverynetwork.Node")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion2

// Client API for DiscoveryNetwork service

type DiscoveryNetworkClient interface {
	// Discovery Calls
	Join(ctx context.Context, in *JoinRequest, opts ...grpc.CallOption) (*JoinResponse, error)
	Disconnect(ctx context.Context, in *DisconnectRequest, opts ...grpc.CallOption) (*EmptyMessage, error)
}

type discoveryNetworkClient struct {
	cc *grpc.ClientConn
}

func NewDiscoveryNetworkClient(cc *grpc.ClientConn) DiscoveryNetworkClient {
	return &discoveryNetworkClient{cc}
}

func (c *discoveryNetworkClient) Join(ctx context.Context, in *JoinRequest, opts ...grpc.CallOption) (*JoinResponse, error) {
	out := new(JoinResponse)
	err := grpc.Invoke(ctx, "/discoverynetwork.DiscoveryNetwork/Join", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *discoveryNetworkClient) Disconnect(ctx context.Context, in *DisconnectRequest, opts ...grpc.CallOption) (*EmptyMessage, error) {
	out := new(EmptyMessage)
	err := grpc.Invoke(ctx, "/discoverynetwork.DiscoveryNetwork/Disconnect", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DiscoveryNetwork service

type DiscoveryNetworkServer interface {
	// Discovery Calls
	Join(context.Context, *JoinRequest) (*JoinResponse, error)
	Disconnect(context.Context, *DisconnectRequest) (*EmptyMessage, error)
}

func RegisterDiscoveryNetworkServer(s *grpc.Server, srv DiscoveryNetworkServer) {
	s.RegisterService(&_DiscoveryNetwork_serviceDesc, srv)
}

func _DiscoveryNetwork_Join_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscoveryNetworkServer).Join(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/discoverynetwork.DiscoveryNetwork/Join",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscoveryNetworkServer).Join(ctx, req.(*JoinRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DiscoveryNetwork_Disconnect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DisconnectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscoveryNetworkServer).Disconnect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/discoverynetwork.DiscoveryNetwork/Disconnect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscoveryNetworkServer).Disconnect(ctx, req.(*DisconnectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DiscoveryNetwork_serviceDesc = grpc.ServiceDesc{
	ServiceName: "discoverynetwork.DiscoveryNetwork",
	HandlerType: (*DiscoveryNetworkServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Join",
			Handler:    _DiscoveryNetwork_Join_Handler,
		},
		{
			MethodName: "Disconnect",
			Handler:    _DiscoveryNetwork_Disconnect_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor0 = []byte{
	// 288 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x92, 0xcf, 0x4a, 0xc3, 0x40,
	0x10, 0xc6, 0x9b, 0x26, 0x8a, 0x4e, 0x42, 0x89, 0x2b, 0x94, 0x20, 0x28, 0x25, 0x2a, 0x78, 0x8a,
	0x50, 0x1f, 0xc1, 0x8a, 0x20, 0xb4, 0x07, 0x7b, 0xf2, 0x54, 0x62, 0x32, 0x48, 0xd0, 0xec, 0xac,
	0x3b, 0x9b, 0x4a, 0xdf, 0xca, 0x47, 0x34, 0xc9, 0xd6, 0x7f, 0x8d, 0x7a, 0xf2, 0x14, 0x66, 0xe6,
	0xcb, 0x6f, 0xbe, 0xdd, 0x6f, 0x61, 0x94, 0x17, 0x9c, 0xd1, 0x12, 0xf5, 0x4a, 0xa2, 0x79, 0x21,
	0xfd, 0x78, 0xfe, 0xd1, 0x48, 0x94, 0x26, 0x43, 0x22, 0xdc, 0x54, 0xc4, 0x03, 0x08, 0xae, 0x4a,
	0x65, 0x56, 0x53, 0x64, 0x4e, 0x1f, 0x30, 0xbe, 0x83, 0xbd, 0x49, 0xa3, 0x91, 0x12, 0x33, 0x73,
	0x8b, 0xcf, 0x15, 0xb2, 0x11, 0x01, 0x78, 0x8a, 0xe8, 0x29, 0x72, 0x46, 0xce, 0xd9, 0xae, 0x08,
	0x61, 0x47, 0xa5, 0xcc, 0xf5, 0xef, 0x79, 0xd4, 0x6f, 0x3b, 0x27, 0xe0, 0x49, 0xca, 0x31, 0x72,
	0xeb, 0xca, 0x1f, 0x0f, 0x93, 0xcd, 0x2d, 0xc9, 0xac, 0x9e, 0xc6, 0x73, 0xf0, 0x6f, 0xa8, 0x90,
	0xff, 0x0b, 0x9d, 0x42, 0x60, 0xa1, 0xac, 0x48, 0x32, 0x8a, 0x21, 0x0c, 0x34, 0x32, 0x9a, 0x45,
	0x21, 0x0d, 0xea, 0x65, 0x6a, 0xf9, 0xae, 0x38, 0x85, 0xad, 0x86, 0xc6, 0x35, 0xdc, 0xfd, 0x03,
	0x77, 0x09, 0x5e, 0xf3, 0x15, 0x00, 0xfd, 0x42, 0xad, 0xad, 0xb5, 0x46, 0xb5, 0x59, 0xdb, 0xda,
	0x07, 0x3f, 0xa3, 0xb2, 0x24, 0xb9, 0x90, 0x69, 0x69, 0xdd, 0xb5, 0x92, 0xaa, 0x2a, 0xf2, 0xc8,
	0x6b, 0xaa, 0xf1, 0xab, 0x03, 0xe1, 0xe4, 0x1d, 0x3f, 0xb3, 0x78, 0x71, 0x0d, 0x5e, 0x63, 0x54,
	0x1c, 0x76, 0x37, 0x7f, 0xb9, 0x95, 0x83, 0xa3, 0xdf, 0xc6, 0xf6, 0x7c, 0x71, 0x4f, 0xcc, 0x01,
	0x3e, 0x13, 0x12, 0xc7, 0x5d, 0x7d, 0x27, 0xbf, 0x9f, 0xa0, 0xdf, 0x42, 0xef, 0xdd, 0x6f, 0xb7,
	0xef, 0xe3, 0xe2, 0x2d, 0x00, 0x00, 0xff, 0xff, 0x56, 0xb1, 0xe3, 0x5b, 0x43, 0x02, 0x00, 0x00,
}
