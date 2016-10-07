// Code generated by protoc-gen-go.
// source: persistentvolume.proto
// DO NOT EDIT!

package v1beta1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/appscodeapis/appscode/api"
import dtypes "github.com/appscode/api/dtypes"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type PVRegisterRequest struct {
	Cluster    string `protobuf:"bytes,1,opt,name=cluster" json:"cluster,omitempty"`
	Name       string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Identifier string `protobuf:"bytes,3,opt,name=identifier" json:"identifier,omitempty"`
	Plugin     string `protobuf:"bytes,4,opt,name=plugin" json:"plugin,omitempty"`
	SizeGb     int64  `protobuf:"varint,5,opt,name=size_gb,json=sizeGb" json:"size_gb,omitempty"`
	Endpoint   string `protobuf:"bytes,6,opt,name=endpoint" json:"endpoint,omitempty"`
}

func (m *PVRegisterRequest) Reset()                    { *m = PVRegisterRequest{} }
func (m *PVRegisterRequest) String() string            { return proto.CompactTextString(m) }
func (*PVRegisterRequest) ProtoMessage()               {}
func (*PVRegisterRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

type PVUnregisterRequest struct {
	Cluster string `protobuf:"bytes,1,opt,name=cluster" json:"cluster,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *PVUnregisterRequest) Reset()                    { *m = PVUnregisterRequest{} }
func (m *PVUnregisterRequest) String() string            { return proto.CompactTextString(m) }
func (*PVUnregisterRequest) ProtoMessage()               {}
func (*PVUnregisterRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

type PVDescribeRequest struct {
	Cluster string `protobuf:"bytes,1,opt,name=cluster" json:"cluster,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *PVDescribeRequest) Reset()                    { *m = PVDescribeRequest{} }
func (m *PVDescribeRequest) String() string            { return proto.CompactTextString(m) }
func (*PVDescribeRequest) ProtoMessage()               {}
func (*PVDescribeRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

type PVInfo struct {
	Name        string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	SizeGb      int64    `protobuf:"varint,2,opt,name=size_gb,json=sizeGb" json:"size_gb,omitempty"`
	Status      string   `protobuf:"bytes,3,opt,name=status" json:"status,omitempty"`
	Volume      string   `protobuf:"bytes,4,opt,name=volume" json:"volume,omitempty"`
	Claim       string   `protobuf:"bytes,5,opt,name=claim" json:"claim,omitempty"`
	Plugin      string   `protobuf:"bytes,6,opt,name=plugin" json:"plugin,omitempty"`
	AccessModes []string `protobuf:"bytes,7,rep,name=accessModes" json:"accessModes,omitempty"`
}

func (m *PVInfo) Reset()                    { *m = PVInfo{} }
func (m *PVInfo) String() string            { return proto.CompactTextString(m) }
func (*PVInfo) ProtoMessage()               {}
func (*PVInfo) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{3} }

type PVDescribeResponse struct {
	Status *dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Pv     *PVInfo        `protobuf:"bytes,2,opt,name=pv" json:"pv,omitempty"`
}

func (m *PVDescribeResponse) Reset()                    { *m = PVDescribeResponse{} }
func (m *PVDescribeResponse) String() string            { return proto.CompactTextString(m) }
func (*PVDescribeResponse) ProtoMessage()               {}
func (*PVDescribeResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{4} }

func (m *PVDescribeResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *PVDescribeResponse) GetPv() *PVInfo {
	if m != nil {
		return m.Pv
	}
	return nil
}

func init() {
	proto.RegisterType((*PVRegisterRequest)(nil), "pv.v1beta1.PVRegisterRequest")
	proto.RegisterType((*PVUnregisterRequest)(nil), "pv.v1beta1.PVUnregisterRequest")
	proto.RegisterType((*PVDescribeRequest)(nil), "pv.v1beta1.PVDescribeRequest")
	proto.RegisterType((*PVInfo)(nil), "pv.v1beta1.PVInfo")
	proto.RegisterType((*PVDescribeResponse)(nil), "pv.v1beta1.PVDescribeResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for PersistentVolumes service

type PersistentVolumesClient interface {
	Describe(ctx context.Context, in *PVDescribeRequest, opts ...grpc.CallOption) (*PVDescribeResponse, error)
	Register(ctx context.Context, in *PVRegisterRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
	Unregister(ctx context.Context, in *PVUnregisterRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
}

type persistentVolumesClient struct {
	cc *grpc.ClientConn
}

func NewPersistentVolumesClient(cc *grpc.ClientConn) PersistentVolumesClient {
	return &persistentVolumesClient{cc}
}

func (c *persistentVolumesClient) Describe(ctx context.Context, in *PVDescribeRequest, opts ...grpc.CallOption) (*PVDescribeResponse, error) {
	out := new(PVDescribeResponse)
	err := grpc.Invoke(ctx, "/pv.v1beta1.PersistentVolumes/Describe", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *persistentVolumesClient) Register(ctx context.Context, in *PVRegisterRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/pv.v1beta1.PersistentVolumes/Register", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *persistentVolumesClient) Unregister(ctx context.Context, in *PVUnregisterRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/pv.v1beta1.PersistentVolumes/Unregister", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PersistentVolumes service

type PersistentVolumesServer interface {
	Describe(context.Context, *PVDescribeRequest) (*PVDescribeResponse, error)
	Register(context.Context, *PVRegisterRequest) (*dtypes.VoidResponse, error)
	Unregister(context.Context, *PVUnregisterRequest) (*dtypes.VoidResponse, error)
}

func RegisterPersistentVolumesServer(s *grpc.Server, srv PersistentVolumesServer) {
	s.RegisterService(&_PersistentVolumes_serviceDesc, srv)
}

func _PersistentVolumes_Describe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PVDescribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PersistentVolumesServer).Describe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pv.v1beta1.PersistentVolumes/Describe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PersistentVolumesServer).Describe(ctx, req.(*PVDescribeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PersistentVolumes_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PVRegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PersistentVolumesServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pv.v1beta1.PersistentVolumes/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PersistentVolumesServer).Register(ctx, req.(*PVRegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PersistentVolumes_Unregister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PVUnregisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PersistentVolumesServer).Unregister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pv.v1beta1.PersistentVolumes/Unregister",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PersistentVolumesServer).Unregister(ctx, req.(*PVUnregisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PersistentVolumes_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pv.v1beta1.PersistentVolumes",
	HandlerType: (*PersistentVolumesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Describe",
			Handler:    _PersistentVolumes_Describe_Handler,
		},
		{
			MethodName: "Register",
			Handler:    _PersistentVolumes_Register_Handler,
		},
		{
			MethodName: "Unregister",
			Handler:    _PersistentVolumes_Unregister_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor2,
}

func init() { proto.RegisterFile("persistentvolume.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 528 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x54, 0xcf, 0x8b, 0xd3, 0x40,
	0x14, 0x66, 0xd2, 0xdd, 0xb4, 0x7d, 0x05, 0x61, 0xc7, 0xa5, 0x86, 0xa0, 0xb5, 0xe4, 0x20, 0x65,
	0x0f, 0x09, 0x5b, 0x6f, 0xde, 0xfc, 0x01, 0x22, 0xa8, 0xd4, 0x88, 0x3d, 0x2c, 0x88, 0xe6, 0xc7,
	0xdb, 0x32, 0xd8, 0xce, 0x8c, 0x99, 0x49, 0x41, 0x97, 0xbd, 0xec, 0xbf, 0xe0, 0x5d, 0x4f, 0x5e,
	0xbd, 0xed, 0x5f, 0xe2, 0xbf, 0xe0, 0xdf, 0x21, 0x92, 0x49, 0x9a, 0x4d, 0x2a, 0xbd, 0xe8, 0x5e,
	0x4a, 0xbe, 0xef, 0x75, 0x32, 0xdf, 0xf7, 0xbd, 0xf7, 0x02, 0x43, 0x89, 0x99, 0x62, 0x4a, 0x23,
	0xd7, 0x6b, 0xb1, 0xcc, 0x57, 0xe8, 0xcb, 0x4c, 0x68, 0x41, 0x41, 0xae, 0xfd, 0xf5, 0x71, 0x8c,
	0x3a, 0x3a, 0x76, 0x6f, 0x2f, 0x84, 0x58, 0x2c, 0x31, 0x88, 0x24, 0x0b, 0x22, 0xce, 0x85, 0x8e,
	0x34, 0x13, 0x5c, 0x95, 0xff, 0x74, 0x47, 0x91, 0x94, 0x2a, 0x11, 0xe9, 0xae, 0xfa, 0xb0, 0xa0,
	0x53, 0xfd, 0x49, 0xa2, 0x0a, 0xcc, 0x6f, 0xc9, 0x7b, 0x3f, 0x08, 0x1c, 0xcc, 0xe6, 0x21, 0x2e,
	0x8a, 0xdb, 0xb3, 0x10, 0x3f, 0xe6, 0xa8, 0x34, 0x75, 0xa0, 0x9b, 0x2c, 0xf3, 0x82, 0x71, 0xc8,
	0x98, 0x4c, 0xfa, 0xe1, 0x06, 0x52, 0x0a, 0x7b, 0x3c, 0x5a, 0xa1, 0x63, 0x19, 0xda, 0x3c, 0xd3,
	0x11, 0x00, 0x4b, 0x91, 0x6b, 0x76, 0xca, 0x30, 0x73, 0x3a, 0xa6, 0xd2, 0x60, 0xe8, 0x10, 0x6c,
	0xb9, 0xcc, 0x17, 0x8c, 0x3b, 0x7b, 0xa6, 0x56, 0x21, 0x7a, 0x0b, 0xba, 0x8a, 0x7d, 0xc6, 0x77,
	0x8b, 0xd8, 0xd9, 0x1f, 0x93, 0x49, 0x27, 0xb4, 0x0b, 0xf8, 0x34, 0xa6, 0x2e, 0xf4, 0x90, 0xa7,
	0x52, 0x30, 0xae, 0x1d, 0xdb, 0x1c, 0xa9, 0xb1, 0xf7, 0x18, 0x6e, 0xce, 0xe6, 0x6f, 0x78, 0xf6,
	0x3f, 0x8a, 0xbd, 0x87, 0x85, 0xe9, 0x27, 0xa8, 0x92, 0x8c, 0xc5, 0xf8, 0x6f, 0xaf, 0xb8, 0x24,
	0x60, 0xcf, 0xe6, 0xcf, 0xf8, 0xa9, 0xa8, 0xcb, 0xa4, 0x91, 0x49, 0xc3, 0x9b, 0xd5, 0xf2, 0x36,
	0x04, 0x5b, 0xe9, 0x48, 0xe7, 0xaa, 0x0a, 0xaa, 0x42, 0x05, 0x5f, 0xb6, 0x7e, 0x13, 0x52, 0x89,
	0xe8, 0x21, 0xec, 0x27, 0xcb, 0x88, 0xad, 0x4c, 0x44, 0xfd, 0xb0, 0x04, 0x8d, 0x48, 0xed, 0x56,
	0xa4, 0x63, 0x18, 0x44, 0x49, 0x82, 0x4a, 0xbd, 0x10, 0x29, 0x2a, 0xa7, 0x3b, 0xee, 0x4c, 0xfa,
	0x61, 0x93, 0xf2, 0xde, 0x03, 0x6d, 0x5a, 0x57, 0x52, 0x70, 0x85, 0xf4, 0x5e, 0xad, 0xaa, 0x30,
	0x31, 0x98, 0xde, 0xf0, 0xcb, 0x59, 0xf1, 0x5f, 0x1b, 0xb6, 0x56, 0xe9, 0x81, 0x25, 0xd7, 0xc6,
	0xd1, 0x60, 0x4a, 0xfd, 0xab, 0xe9, 0xf4, 0xcb, 0x28, 0x42, 0x4b, 0xae, 0xa7, 0xbf, 0x3b, 0x70,
	0x30, 0xab, 0xe7, 0x79, 0x6e, 0x6c, 0x28, 0xfa, 0x9d, 0x40, 0x6f, 0x73, 0x2d, 0xbd, 0xd3, 0x3e,
	0xba, 0xd5, 0x09, 0x77, 0xb4, 0xab, 0x5c, 0xaa, 0xf5, 0x4e, 0x2e, 0x2e, 0x1d, 0xab, 0x47, 0x2e,
	0x7e, 0xfe, 0xfa, 0x62, 0xbd, 0xa4, 0xcf, 0x83, 0xd6, 0xec, 0x7f, 0xc8, 0x63, 0xcc, 0x38, 0x6a,
	0x54, 0x41, 0xf5, 0x8a, 0xa0, 0xea, 0xa3, 0x0a, 0xce, 0xaa, 0xa7, 0xf3, 0x60, 0x7b, 0xe1, 0x54,
	0x70, 0x56, 0xf4, 0xed, 0x9c, 0x7e, 0x25, 0xd0, 0xdb, 0xac, 0xc3, 0xb6, 0xce, 0xad, 0x35, 0x71,
	0x0f, 0x37, 0x29, 0xcd, 0x05, 0x4b, 0x6b, 0x75, 0x6f, 0x1b, 0xea, 0x5e, 0xb9, 0xd7, 0xaa, 0xee,
	0x01, 0x39, 0xa2, 0xdf, 0x08, 0xc0, 0xd5, 0xfc, 0xd3, 0xbb, 0x6d, 0x89, 0x7f, 0x6d, 0xc6, 0x0e,
	0x91, 0xad, 0x08, 0x8f, 0xae, 0x55, 0xe4, 0xa3, 0xfe, 0x49, 0xb7, 0x3a, 0x16, 0xdb, 0xe6, 0x2b,
	0x73, 0xff, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x1d, 0x22, 0x19, 0x2f, 0xe1, 0x04, 0x00, 0x00,
}