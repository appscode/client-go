// Code generated by protoc-gen-go.
// source: metadata.proto
// DO NOT EDIT!

package v1beta1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/appscodeapis/appscode/api"
import appscode_dtypes "github.com/appscode/api/dtypes"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ListRegionsRequest struct {
	Provider        string `protobuf:"bytes,1,opt,name=provider" json:"provider,omitempty"`
	CloudCredential string `protobuf:"bytes,2,opt,name=cloud_credential,json=cloudCredential" json:"cloud_credential,omitempty"`
}

func (m *ListRegionsRequest) Reset()                    { *m = ListRegionsRequest{} }
func (m *ListRegionsRequest) String() string            { return proto.CompactTextString(m) }
func (*ListRegionsRequest) ProtoMessage()               {}
func (*ListRegionsRequest) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

type ListRegionsResponse struct {
	Status  *appscode_dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Regions []string                `protobuf:"bytes,2,rep,name=regions" json:"regions,omitempty"`
}

func (m *ListRegionsResponse) Reset()                    { *m = ListRegionsResponse{} }
func (m *ListRegionsResponse) String() string            { return proto.CompactTextString(m) }
func (*ListRegionsResponse) ProtoMessage()               {}
func (*ListRegionsResponse) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

func (m *ListRegionsResponse) GetStatus() *appscode_dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

type ListZonesRequest struct {
	Provider        string `protobuf:"bytes,1,opt,name=provider" json:"provider,omitempty"`
	Region          string `protobuf:"bytes,2,opt,name=region" json:"region,omitempty"`
	CloudCredential string `protobuf:"bytes,3,opt,name=cloud_credential,json=cloudCredential" json:"cloud_credential,omitempty"`
}

func (m *ListZonesRequest) Reset()                    { *m = ListZonesRequest{} }
func (m *ListZonesRequest) String() string            { return proto.CompactTextString(m) }
func (*ListZonesRequest) ProtoMessage()               {}
func (*ListZonesRequest) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{2} }

type ListZonesResponse struct {
	Status *appscode_dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Zones  []string                `protobuf:"bytes,2,rep,name=zones" json:"zones,omitempty"`
}

func (m *ListZonesResponse) Reset()                    { *m = ListZonesResponse{} }
func (m *ListZonesResponse) String() string            { return proto.CompactTextString(m) }
func (*ListZonesResponse) ProtoMessage()               {}
func (*ListZonesResponse) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{3} }

func (m *ListZonesResponse) GetStatus() *appscode_dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func init() {
	proto.RegisterType((*ListRegionsRequest)(nil), "appscode.kubernetes.v1beta1.ListRegionsRequest")
	proto.RegisterType((*ListRegionsResponse)(nil), "appscode.kubernetes.v1beta1.ListRegionsResponse")
	proto.RegisterType((*ListZonesRequest)(nil), "appscode.kubernetes.v1beta1.ListZonesRequest")
	proto.RegisterType((*ListZonesResponse)(nil), "appscode.kubernetes.v1beta1.ListZonesResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Metadata service

type MetadataClient interface {
	ListRegions(ctx context.Context, in *ListRegionsRequest, opts ...grpc.CallOption) (*ListRegionsResponse, error)
	ListZones(ctx context.Context, in *ListZonesRequest, opts ...grpc.CallOption) (*ListZonesResponse, error)
}

type metadataClient struct {
	cc *grpc.ClientConn
}

func NewMetadataClient(cc *grpc.ClientConn) MetadataClient {
	return &metadataClient{cc}
}

func (c *metadataClient) ListRegions(ctx context.Context, in *ListRegionsRequest, opts ...grpc.CallOption) (*ListRegionsResponse, error) {
	out := new(ListRegionsResponse)
	err := grpc.Invoke(ctx, "/appscode.kubernetes.v1beta1.Metadata/ListRegions", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metadataClient) ListZones(ctx context.Context, in *ListZonesRequest, opts ...grpc.CallOption) (*ListZonesResponse, error) {
	out := new(ListZonesResponse)
	err := grpc.Invoke(ctx, "/appscode.kubernetes.v1beta1.Metadata/ListZones", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Metadata service

type MetadataServer interface {
	ListRegions(context.Context, *ListRegionsRequest) (*ListRegionsResponse, error)
	ListZones(context.Context, *ListZonesRequest) (*ListZonesResponse, error)
}

func RegisterMetadataServer(s *grpc.Server, srv MetadataServer) {
	s.RegisterService(&_Metadata_serviceDesc, srv)
}

func _Metadata_ListRegions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRegionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetadataServer).ListRegions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.kubernetes.v1beta1.Metadata/ListRegions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetadataServer).ListRegions(ctx, req.(*ListRegionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Metadata_ListZones_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListZonesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetadataServer).ListZones(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.kubernetes.v1beta1.Metadata/ListZones",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetadataServer).ListZones(ctx, req.(*ListZonesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Metadata_serviceDesc = grpc.ServiceDesc{
	ServiceName: "appscode.kubernetes.v1beta1.Metadata",
	HandlerType: (*MetadataServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListRegions",
			Handler:    _Metadata_ListRegions_Handler,
		},
		{
			MethodName: "ListZones",
			Handler:    _Metadata_ListZones_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor3,
}

func init() { proto.RegisterFile("metadata.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 425 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xa4, 0x53, 0x4d, 0xae, 0xd3, 0x30,
	0x10, 0x56, 0x52, 0xd1, 0x1f, 0x57, 0x40, 0x31, 0x08, 0xa2, 0x80, 0xa0, 0x0a, 0x9b, 0xb2, 0xc0,
	0xa6, 0xe5, 0x00, 0x88, 0x76, 0x0b, 0xa2, 0x0a, 0xbb, 0x22, 0x01, 0x6e, 0x32, 0xaa, 0x22, 0xda,
	0xd8, 0x8d, 0x9d, 0x4a, 0x50, 0x75, 0xd3, 0x2b, 0xb0, 0xe4, 0x1c, 0x1c, 0x80, 0x2d, 0x5b, 0xae,
	0xc0, 0x41, 0x50, 0x6c, 0x27, 0xa4, 0x7a, 0x7d, 0xef, 0x55, 0xef, 0x6d, 0xa2, 0x8c, 0x67, 0xbe,
	0xf1, 0x37, 0xdf, 0x37, 0x46, 0xb7, 0x56, 0xa0, 0x58, 0xcc, 0x14, 0x23, 0x22, 0xe3, 0x8a, 0xe3,
	0x87, 0x4c, 0x08, 0x19, 0xf1, 0x18, 0xc8, 0x97, 0x7c, 0x0e, 0x59, 0x0a, 0x0a, 0x24, 0xd9, 0x0c,
	0xe7, 0xa0, 0xd8, 0xd0, 0x7f, 0xb4, 0xe0, 0x7c, 0xb1, 0x04, 0xca, 0x44, 0x42, 0x59, 0x9a, 0x72,
	0xc5, 0x54, 0xc2, 0x53, 0x69, 0xa0, 0xfe, 0xe3, 0x12, 0x7a, 0x4e, 0xfe, 0xc9, 0x41, 0x3e, 0x56,
	0x5f, 0x05, 0x48, 0xaa, 0xbf, 0xa6, 0x20, 0xf8, 0x80, 0xf0, 0x9b, 0x44, 0xaa, 0x10, 0x16, 0x05,
	0x2a, 0x84, 0x75, 0x0e, 0x52, 0x61, 0x1f, 0xb5, 0x45, 0xc6, 0x37, 0x49, 0x0c, 0x99, 0xe7, 0xf4,
	0x9d, 0x41, 0x27, 0xac, 0x62, 0xfc, 0x0c, 0xf5, 0xa2, 0x25, 0xcf, 0xe3, 0x4f, 0x51, 0x06, 0x31,
	0xa4, 0x2a, 0x61, 0x4b, 0xcf, 0xd5, 0x35, 0xb7, 0xf5, 0xf9, 0xa4, 0x3a, 0x0e, 0x3e, 0xa3, 0xbb,
	0x07, 0xcd, 0xa5, 0xe0, 0xa9, 0x04, 0x4c, 0x51, 0x53, 0x2a, 0xa6, 0x72, 0xa9, 0x7b, 0x77, 0x47,
	0x0f, 0x48, 0x25, 0x80, 0x61, 0x48, 0xde, 0xeb, 0x74, 0x68, 0xcb, 0xb0, 0x87, 0x5a, 0x99, 0xe9,
	0xe1, 0xb9, 0xfd, 0xc6, 0xa0, 0x13, 0x96, 0x61, 0xb0, 0x46, 0xbd, 0xe2, 0x86, 0x19, 0x4f, 0xe1,
	0x24, 0xf2, 0xf7, 0x51, 0xd3, 0x40, 0x2d, 0x65, 0x1b, 0x1d, 0x1d, 0xaa, 0x71, 0x7c, 0xa8, 0x19,
	0xba, 0x53, 0xbb, 0xf2, 0xaa, 0x23, 0xdd, 0x43, 0x37, 0xbe, 0x15, 0x1d, 0xec, 0x40, 0x26, 0x18,
	0xfd, 0x68, 0xa0, 0xf6, 0x5b, 0xbb, 0x1c, 0xf8, 0x97, 0x83, 0xba, 0x35, 0xf9, 0x30, 0x25, 0x17,
	0xec, 0x09, 0x39, 0xeb, 0xa2, 0xff, 0xe2, 0x74, 0x80, 0x19, 0x23, 0x78, 0xb7, 0xff, 0xe9, 0xb9,
	0x6d, 0x67, 0xff, 0xe7, 0xef, 0x77, 0x77, 0x82, 0x5f, 0xd3, 0x83, 0xed, 0xd1, 0x42, 0x50, 0x0b,
	0xa7, 0xe5, 0x0e, 0xd3, 0x52, 0x5b, 0x49, 0xb7, 0xe5, 0xef, 0x8e, 0x5a, 0x7f, 0xf0, 0x6f, 0x07,
	0x75, 0x2a, 0xb5, 0xf0, 0xf3, 0x4b, 0x09, 0xd5, 0x8d, 0xf4, 0xc9, 0xa9, 0xe5, 0x96, 0xfd, 0xc7,
	0x1a, 0xfb, 0x10, 0x4f, 0xaf, 0xcd, 0x9e, 0x6e, 0xcd, 0xcf, 0x8e, 0x6a, 0x77, 0xc6, 0xaf, 0xd0,
	0xd3, 0x88, 0xaf, 0xfe, 0x93, 0x62, 0x22, 0x39, 0x42, 0x6c, 0x7c, 0xb3, 0x74, 0x70, 0x5a, 0xbc,
	0xb0, 0xa9, 0x33, 0x6b, 0xd9, 0xcc, 0xbc, 0xa9, 0xdf, 0xdc, 0xcb, 0x7f, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x9f, 0x0e, 0x9a, 0xb6, 0x01, 0x04, 0x00, 0x00,
}
