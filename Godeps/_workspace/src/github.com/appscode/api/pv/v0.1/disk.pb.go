// Code generated by protoc-gen-go.
// source: disk.proto
// DO NOT EDIT!

/*
Package pv is a generated protocol buffer package.

It is generated from these files:
	disk.proto
	pvc.proto
	pv.proto

It has these top-level messages:
	DiskCreateRequest
	DiskDeleteRequest
	DiskListRequest
	DiskListResponse
	Result
	Disk
	PV
	PVC
	DiskDescribeRequest
	DiskDescribeResponse
	PVCRegisterRequest
	PVCUnregisterRequest
	PVCDescribeRequest
	PVCInfo
	PVCDescribeResponse
	PVRegisterRequest
	PVUnregisterRequest
	PVDescribeRequest
	PVInfo
	PVDescribeResponse
*/
package pv

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api"
import dtypes "github.com/appscode/api/dtypes"

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

type DiskCreateRequest struct {
	Cluster  string `protobuf:"bytes,1,opt,name=cluster" json:"cluster,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Zone     string `protobuf:"bytes,3,opt,name=zone" json:"zone,omitempty"`
	DiskType string `protobuf:"bytes,4,opt,name=disk_type,json=diskType" json:"disk_type,omitempty"`
	Size     int64  `protobuf:"varint,5,opt,name=size" json:"size,omitempty"`
}

func (m *DiskCreateRequest) Reset()                    { *m = DiskCreateRequest{} }
func (m *DiskCreateRequest) String() string            { return proto.CompactTextString(m) }
func (*DiskCreateRequest) ProtoMessage()               {}
func (*DiskCreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type DiskDeleteRequest struct {
	Cluster    string `protobuf:"bytes,1,opt,name=cluster" json:"cluster,omitempty"`
	Identifier string `protobuf:"bytes,2,opt,name=identifier" json:"identifier,omitempty"`
}

func (m *DiskDeleteRequest) Reset()                    { *m = DiskDeleteRequest{} }
func (m *DiskDeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*DiskDeleteRequest) ProtoMessage()               {}
func (*DiskDeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type DiskListRequest struct {
	Cluster string `protobuf:"bytes,1,opt,name=cluster" json:"cluster,omitempty"`
}

func (m *DiskListRequest) Reset()                    { *m = DiskListRequest{} }
func (m *DiskListRequest) String() string            { return proto.CompactTextString(m) }
func (*DiskListRequest) ProtoMessage()               {}
func (*DiskListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type DiskListResponse struct {
	Status *dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Kube   string         `protobuf:"bytes,2,opt,name=kube" json:"kube,omitempty"`
	Result []*Result      `protobuf:"bytes,3,rep,name=result" json:"result,omitempty"`
}

func (m *DiskListResponse) Reset()                    { *m = DiskListResponse{} }
func (m *DiskListResponse) String() string            { return proto.CompactTextString(m) }
func (*DiskListResponse) ProtoMessage()               {}
func (*DiskListResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *DiskListResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *DiskListResponse) GetResult() []*Result {
	if m != nil {
		return m.Result
	}
	return nil
}

type Result struct {
	Disk *Disk `protobuf:"bytes,1,opt,name=disk" json:"disk,omitempty"`
	Pv   *PV   `protobuf:"bytes,2,opt,name=pv" json:"pv,omitempty"`
	Pvc  *PVC  `protobuf:"bytes,3,opt,name=pvc" json:"pvc,omitempty"`
}

func (m *Result) Reset()                    { *m = Result{} }
func (m *Result) String() string            { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()               {}
func (*Result) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Result) GetDisk() *Disk {
	if m != nil {
		return m.Disk
	}
	return nil
}

func (m *Result) GetPv() *PV {
	if m != nil {
		return m.Pv
	}
	return nil
}

func (m *Result) GetPvc() *PVC {
	if m != nil {
		return m.Pvc
	}
	return nil
}

type Disk struct {
	Name     string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Id       string   `protobuf:"bytes,2,opt,name=id" json:"id,omitempty"`
	Plugin   string   `protobuf:"bytes,3,opt,name=plugin" json:"plugin,omitempty"`
	Size     int64    `protobuf:"varint,4,opt,name=size" json:"size,omitempty"`
	Type     string   `protobuf:"bytes,5,opt,name=type" json:"type,omitempty"`
	Zone     string   `protobuf:"bytes,6,opt,name=zone" json:"zone,omitempty"`
	Status   string   `protobuf:"bytes,7,opt,name=status" json:"status,omitempty"`
	Users    []string `protobuf:"bytes,8,rep,name=users" json:"users,omitempty"`
	Kind     string   `protobuf:"bytes,9,opt,name=kind" json:"kind,omitempty"`
	Endpoint string   `protobuf:"bytes,10,opt,name=endpoint" json:"endpoint,omitempty"`
	Iops     int64    `protobuf:"varint,11,opt,name=iops" json:"iops,omitempty"`
}

func (m *Disk) Reset()                    { *m = Disk{} }
func (m *Disk) String() string            { return proto.CompactTextString(m) }
func (*Disk) ProtoMessage()               {}
func (*Disk) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type PV struct {
	Name   string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Size   int64  `protobuf:"varint,2,opt,name=size" json:"size,omitempty"`
	Status string `protobuf:"bytes,3,opt,name=status" json:"status,omitempty"`
	Claim  string `protobuf:"bytes,4,opt,name=claim" json:"claim,omitempty"`
	Volume string `protobuf:"bytes,5,opt,name=volume" json:"volume,omitempty"`
	Plugin string `protobuf:"bytes,6,opt,name=plugin" json:"plugin,omitempty"`
	Type   string `protobuf:"bytes,7,opt,name=type" json:"type,omitempty"`
}

func (m *PV) Reset()                    { *m = PV{} }
func (m *PV) String() string            { return proto.CompactTextString(m) }
func (*PV) ProtoMessage()               {}
func (*PV) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

type PVC struct {
	Name      string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Size      int64  `protobuf:"varint,2,opt,name=size" json:"size,omitempty"`
	Namespace string `protobuf:"bytes,3,opt,name=namespace" json:"namespace,omitempty"`
	Status    string `protobuf:"bytes,4,opt,name=status" json:"status,omitempty"`
	Volume    string `protobuf:"bytes,5,opt,name=volume" json:"volume,omitempty"`
}

func (m *PVC) Reset()                    { *m = PVC{} }
func (m *PVC) String() string            { return proto.CompactTextString(m) }
func (*PVC) ProtoMessage()               {}
func (*PVC) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

type DiskDescribeRequest struct {
	Cluster string `protobuf:"bytes,1,opt,name=cluster" json:"cluster,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Plugin  string `protobuf:"bytes,3,opt,name=plugin" json:"plugin,omitempty"`
}

func (m *DiskDescribeRequest) Reset()                    { *m = DiskDescribeRequest{} }
func (m *DiskDescribeRequest) String() string            { return proto.CompactTextString(m) }
func (*DiskDescribeRequest) ProtoMessage()               {}
func (*DiskDescribeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

type DiskDescribeResponse struct {
	Status *dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Disk   *Disk          `protobuf:"bytes,2,opt,name=disk" json:"disk,omitempty"`
}

func (m *DiskDescribeResponse) Reset()                    { *m = DiskDescribeResponse{} }
func (m *DiskDescribeResponse) String() string            { return proto.CompactTextString(m) }
func (*DiskDescribeResponse) ProtoMessage()               {}
func (*DiskDescribeResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *DiskDescribeResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *DiskDescribeResponse) GetDisk() *Disk {
	if m != nil {
		return m.Disk
	}
	return nil
}

func init() {
	proto.RegisterType((*DiskCreateRequest)(nil), "pv.DiskCreateRequest")
	proto.RegisterType((*DiskDeleteRequest)(nil), "pv.DiskDeleteRequest")
	proto.RegisterType((*DiskListRequest)(nil), "pv.DiskListRequest")
	proto.RegisterType((*DiskListResponse)(nil), "pv.DiskListResponse")
	proto.RegisterType((*Result)(nil), "pv.Result")
	proto.RegisterType((*Disk)(nil), "pv.Disk")
	proto.RegisterType((*PV)(nil), "pv.PV")
	proto.RegisterType((*PVC)(nil), "pv.PVC")
	proto.RegisterType((*DiskDescribeRequest)(nil), "pv.DiskDescribeRequest")
	proto.RegisterType((*DiskDescribeResponse)(nil), "pv.DiskDescribeResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Disks service

type DisksClient interface {
	List(ctx context.Context, in *DiskListRequest, opts ...grpc.CallOption) (*DiskListResponse, error)
	Describe(ctx context.Context, in *DiskDescribeRequest, opts ...grpc.CallOption) (*DiskDescribeResponse, error)
	Create(ctx context.Context, in *DiskCreateRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
	Delete(ctx context.Context, in *DiskDeleteRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
}

type disksClient struct {
	cc *grpc.ClientConn
}

func NewDisksClient(cc *grpc.ClientConn) DisksClient {
	return &disksClient{cc}
}

func (c *disksClient) List(ctx context.Context, in *DiskListRequest, opts ...grpc.CallOption) (*DiskListResponse, error) {
	out := new(DiskListResponse)
	err := grpc.Invoke(ctx, "/pv.Disks/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *disksClient) Describe(ctx context.Context, in *DiskDescribeRequest, opts ...grpc.CallOption) (*DiskDescribeResponse, error) {
	out := new(DiskDescribeResponse)
	err := grpc.Invoke(ctx, "/pv.Disks/Describe", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *disksClient) Create(ctx context.Context, in *DiskCreateRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/pv.Disks/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *disksClient) Delete(ctx context.Context, in *DiskDeleteRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/pv.Disks/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Disks service

type DisksServer interface {
	List(context.Context, *DiskListRequest) (*DiskListResponse, error)
	Describe(context.Context, *DiskDescribeRequest) (*DiskDescribeResponse, error)
	Create(context.Context, *DiskCreateRequest) (*dtypes.VoidResponse, error)
	Delete(context.Context, *DiskDeleteRequest) (*dtypes.VoidResponse, error)
}

func RegisterDisksServer(s *grpc.Server, srv DisksServer) {
	s.RegisterService(&_Disks_serviceDesc, srv)
}

func _Disks_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DiskListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DisksServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pv.Disks/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DisksServer).List(ctx, req.(*DiskListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Disks_Describe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DiskDescribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DisksServer).Describe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pv.Disks/Describe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DisksServer).Describe(ctx, req.(*DiskDescribeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Disks_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DiskCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DisksServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pv.Disks/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DisksServer).Create(ctx, req.(*DiskCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Disks_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DiskDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DisksServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pv.Disks/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DisksServer).Delete(ctx, req.(*DiskDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Disks_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pv.Disks",
	HandlerType: (*DisksServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _Disks_List_Handler,
		},
		{
			MethodName: "Describe",
			Handler:    _Disks_Describe_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _Disks_Create_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Disks_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("disk.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 705 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x55, 0xcd, 0x6e, 0xd3, 0x4a,
	0x14, 0x96, 0x7f, 0xe2, 0x24, 0x27, 0x52, 0xef, 0xbd, 0xd3, 0xde, 0x5e, 0xdf, 0x50, 0x21, 0x64,
	0x89, 0x1f, 0x15, 0x88, 0x69, 0x00, 0x55, 0xaa, 0xc4, 0xaa, 0x5d, 0x82, 0x54, 0x19, 0x54, 0x09,
	0x81, 0x04, 0x6e, 0x3c, 0x54, 0xa3, 0xba, 0xf6, 0xd4, 0x63, 0x47, 0xa2, 0xa5, 0x1b, 0x16, 0xbc,
	0x00, 0x7b, 0x1e, 0x88, 0x2d, 0xaf, 0xc0, 0x03, 0xf0, 0x08, 0xcc, 0x99, 0x19, 0xc7, 0x4e, 0x69,
	0x45, 0xca, 0xa6, 0x9d, 0xf3, 0x9d, 0xf1, 0xf9, 0xbe, 0xf3, 0x9d, 0x99, 0x09, 0x40, 0xc2, 0xc4,
	0xe1, 0x88, 0x17, 0x79, 0x99, 0x13, 0x9b, 0x4f, 0x87, 0x6b, 0x07, 0x79, 0x7e, 0x90, 0xd2, 0x30,
	0xe6, 0x2c, 0x8c, 0xb3, 0x2c, 0x2f, 0xe3, 0x92, 0xe5, 0x99, 0xd0, 0x3b, 0x86, 0xab, 0x08, 0x27,
	0xe5, 0x7b, 0x4e, 0x45, 0xa8, 0xfe, 0x6a, 0x3c, 0xf8, 0x64, 0xc1, 0x3f, 0x3b, 0xb2, 0xd0, 0x76,
	0x41, 0xe3, 0x92, 0x46, 0xf4, 0xb8, 0xa2, 0xa2, 0x24, 0x3e, 0x74, 0x27, 0x69, 0x25, 0x4a, 0x5a,
	0xf8, 0xd6, 0x0d, 0xeb, 0x4e, 0x3f, 0xaa, 0x43, 0x42, 0xc0, 0xcd, 0xe2, 0x23, 0xea, 0xdb, 0x0a,
	0x56, 0x6b, 0xc4, 0x4e, 0xf2, 0x8c, 0xfa, 0x8e, 0xc6, 0x70, 0x4d, 0xae, 0x41, 0x1f, 0xf5, 0xbd,
	0x41, 0x2e, 0xdf, 0x55, 0x89, 0x1e, 0x02, 0x2f, 0x64, 0x8c, 0x1f, 0x08, 0x76, 0x42, 0xfd, 0x8e,
	0xc4, 0x9d, 0x48, 0xad, 0x83, 0x67, 0x5a, 0xc7, 0x0e, 0x4d, 0xe9, 0x22, 0x3a, 0xae, 0x03, 0xb0,
	0x84, 0x66, 0x25, 0x7b, 0xc7, 0x64, 0x52, 0xab, 0x69, 0x21, 0xc1, 0x5d, 0xf8, 0x0b, 0xcb, 0x3d,
	0x65, 0xa2, 0xfc, 0x6d, 0xb1, 0xa0, 0x80, 0xbf, 0x9b, 0xcd, 0x82, 0x4b, 0xd7, 0x28, 0xb9, 0x05,
	0x9e, 0x90, 0x16, 0x56, 0x42, 0x6d, 0x1e, 0x8c, 0x97, 0x46, 0xda, 0xbd, 0xd1, 0x73, 0x85, 0x46,
	0x26, 0x8b, 0xbd, 0x1c, 0x56, 0xfb, 0x33, 0x43, 0x70, 0x4d, 0x02, 0xf0, 0x0a, 0x2a, 0xaa, 0xb4,
	0x94, 0x96, 0x38, 0xf2, 0x5b, 0x18, 0xf1, 0xe9, 0x28, 0x52, 0x48, 0x64, 0x32, 0xc1, 0x4b, 0xf0,
	0x34, 0x42, 0xd6, 0xc0, 0x45, 0x67, 0x0c, 0x4f, 0x0f, 0xf7, 0xa2, 0x9a, 0x48, 0xa1, 0x64, 0x15,
	0xe4, 0x70, 0x55, 0xf5, 0xc1, 0xd8, 0xc3, 0xdc, 0xee, 0x5e, 0x24, 0x11, 0xf2, 0x3f, 0x38, 0x7c,
	0x3a, 0x51, 0x9e, 0x0f, 0xc6, 0x5d, 0x9d, 0xd8, 0x8e, 0x10, 0x0b, 0x7e, 0x58, 0xe0, 0x62, 0x85,
	0xd9, 0xb0, 0xac, 0xd6, 0xb0, 0x96, 0xc0, 0x66, 0x89, 0x51, 0x2b, 0x57, 0xb2, 0xbe, 0xc7, 0xd3,
	0xea, 0x80, 0x65, 0x66, 0x7c, 0x26, 0x9a, 0xcd, 0xc8, 0x6d, 0x66, 0x84, 0x98, 0x9a, 0x67, 0x47,
	0xd7, 0x2b, 0xcd, 0x2c, 0xd5, 0xf0, 0xbd, 0xd6, 0xf0, 0x57, 0x67, 0xde, 0x75, 0x75, 0x4d, 0xe3,
	0xd5, 0x0a, 0x74, 0x2a, 0x41, 0x0b, 0xe1, 0xf7, 0xa4, 0x2d, 0xfd, 0x48, 0x07, 0xca, 0x41, 0x96,
	0x25, 0x7e, 0xdf, 0x38, 0x28, 0xd7, 0x64, 0x08, 0x3d, 0x9a, 0x25, 0x3c, 0x67, 0x59, 0xe9, 0x83,
	0x3e, 0x3d, 0x75, 0x8c, 0xfb, 0x59, 0xce, 0x85, 0x3f, 0xd0, 0xca, 0x70, 0x1d, 0x7c, 0xb1, 0xc0,
	0xde, 0xdd, 0xbb, 0xb0, 0xe1, 0xba, 0x11, 0xbb, 0xd5, 0x48, 0x23, 0xd0, 0x39, 0x2f, 0x70, 0x92,
	0xc6, 0xec, 0xc8, 0x9c, 0x58, 0x1d, 0xe0, 0xee, 0x69, 0x9e, 0x56, 0x47, 0x75, 0xe3, 0x26, 0x6a,
	0x59, 0xe7, 0x9d, 0xb7, 0x4e, 0xd9, 0xd4, 0x6d, 0x6c, 0x0a, 0x4e, 0xc1, 0x91, 0xf3, 0x59, 0x58,
	0xe0, 0x1a, 0xf4, 0x31, 0x27, 0x78, 0x3c, 0xa9, 0xef, 0x55, 0x03, 0xb4, 0xe4, 0xbb, 0x73, 0xf2,
	0x2f, 0x11, 0x1a, 0xbc, 0x82, 0x65, 0x7d, 0xb7, 0xc4, 0xa4, 0x60, 0xfb, 0x7f, 0x78, 0xcb, 0x2f,
	0x39, 0x28, 0xc1, 0x6b, 0x58, 0x99, 0x2f, 0x7e, 0xc5, 0x0b, 0x54, 0x1f, 0x7f, 0xfb, 0xa2, 0xe3,
	0x3f, 0xfe, 0xea, 0x40, 0x07, 0x43, 0x41, 0xde, 0x82, 0x8b, 0x17, 0x94, 0x2c, 0xd7, 0x3b, 0x5a,
	0x77, 0x7b, 0xb8, 0x32, 0x0f, 0x6a, 0x09, 0xc1, 0xfd, 0x8f, 0xdf, 0xbe, 0x7f, 0xb6, 0x6f, 0x93,
	0x9b, 0xf2, 0x51, 0xe4, 0x62, 0x92, 0x27, 0xfa, 0x75, 0xe4, 0xd3, 0x70, 0xfa, 0x60, 0xb4, 0x11,
	0x22, 0x81, 0x08, 0x4f, 0x4d, 0xd3, 0x67, 0xe4, 0x03, 0xf4, 0xea, 0x2e, 0xc8, 0x7f, 0x75, 0xc1,
	0x73, 0xa6, 0x0d, 0xfd, 0x5f, 0x13, 0x86, 0xed, 0x89, 0x62, 0xdb, 0x24, 0x8f, 0x17, 0x62, 0x0b,
	0x4f, 0xd1, 0x55, 0xf9, 0x4f, 0xdb, 0x78, 0x46, 0x32, 0xf0, 0xf4, 0x23, 0x4c, 0xfe, 0xad, 0x29,
	0xe6, 0x1e, 0x65, 0xd9, 0xa3, 0x31, 0x70, 0x2f, 0x67, 0xc9, 0x8c, 0x75, 0x53, 0xb1, 0x6e, 0x0c,
	0xef, 0x5d, 0x85, 0x75, 0xcb, 0x5a, 0x27, 0xc7, 0xe0, 0xe9, 0xc7, 0xb6, 0xe1, 0x9b, 0x7b, 0x7c,
	0x2f, 0xe1, 0xdb, 0x52, 0x7c, 0x8f, 0xd6, 0xc7, 0x0b, 0xf2, 0x35, 0x6f, 0xf2, 0xd9, 0xbe, 0xa7,
	0x7e, 0x73, 0x1e, 0xfe, 0x0c, 0x00, 0x00, 0xff, 0xff, 0x9f, 0x96, 0x55, 0x47, 0xbb, 0x06, 0x00,
	0x00,
}
