// Code generated by protoc-gen-go.
// source: disk.proto
// DO NOT EDIT!

/*
Package v1beta1 is a generated protocol buffer package.

It is generated from these files:
	disk.proto
	persistentvolumeclaim.proto
	persistentvolume.proto

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
package v1beta1

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
	SizeGb   int64  `protobuf:"varint,5,opt,name=size_gb,json=sizeGb" json:"size_gb,omitempty"`
}

func (m *DiskCreateRequest) Reset()                    { *m = DiskCreateRequest{} }
func (m *DiskCreateRequest) String() string            { return proto.CompactTextString(m) }
func (*DiskCreateRequest) ProtoMessage()               {}
func (*DiskCreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type DiskDeleteRequest struct {
	Cluster string `protobuf:"bytes,1,opt,name=cluster" json:"cluster,omitempty"`
	Uid     string `protobuf:"bytes,2,opt,name=uid" json:"uid,omitempty"`
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
	SizeGb   int64    `protobuf:"varint,4,opt,name=size_gb,json=sizeGb" json:"size_gb,omitempty"`
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
	SizeGb int64  `protobuf:"varint,2,opt,name=size_gb,json=sizeGb" json:"size_gb,omitempty"`
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
	SizeGb    int64  `protobuf:"varint,2,opt,name=size_gb,json=sizeGb" json:"size_gb,omitempty"`
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
	proto.RegisterType((*DiskCreateRequest)(nil), "pv.v1beta1.DiskCreateRequest")
	proto.RegisterType((*DiskDeleteRequest)(nil), "pv.v1beta1.DiskDeleteRequest")
	proto.RegisterType((*DiskListRequest)(nil), "pv.v1beta1.DiskListRequest")
	proto.RegisterType((*DiskListResponse)(nil), "pv.v1beta1.DiskListResponse")
	proto.RegisterType((*Result)(nil), "pv.v1beta1.Result")
	proto.RegisterType((*Disk)(nil), "pv.v1beta1.Disk")
	proto.RegisterType((*PV)(nil), "pv.v1beta1.PV")
	proto.RegisterType((*PVC)(nil), "pv.v1beta1.PVC")
	proto.RegisterType((*DiskDescribeRequest)(nil), "pv.v1beta1.DiskDescribeRequest")
	proto.RegisterType((*DiskDescribeResponse)(nil), "pv.v1beta1.DiskDescribeResponse")
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
	err := grpc.Invoke(ctx, "/pv.v1beta1.Disks/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *disksClient) Describe(ctx context.Context, in *DiskDescribeRequest, opts ...grpc.CallOption) (*DiskDescribeResponse, error) {
	out := new(DiskDescribeResponse)
	err := grpc.Invoke(ctx, "/pv.v1beta1.Disks/Describe", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *disksClient) Create(ctx context.Context, in *DiskCreateRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/pv.v1beta1.Disks/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *disksClient) Delete(ctx context.Context, in *DiskDeleteRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/pv.v1beta1.Disks/Delete", in, out, c.cc, opts...)
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
		FullMethod: "/pv.v1beta1.Disks/List",
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
		FullMethod: "/pv.v1beta1.Disks/Describe",
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
		FullMethod: "/pv.v1beta1.Disks/Create",
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
		FullMethod: "/pv.v1beta1.Disks/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DisksServer).Delete(ctx, req.(*DiskDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Disks_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pv.v1beta1.Disks",
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
	// 730 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x55, 0x4d, 0x6f, 0xd3, 0x4c,
	0x10, 0x96, 0x3f, 0xe2, 0x24, 0x13, 0xa9, 0xed, 0xbb, 0x6f, 0xd5, 0xd7, 0x4a, 0xfb, 0x42, 0xb0,
	0x10, 0x8a, 0x8a, 0x14, 0xab, 0x01, 0x21, 0x51, 0x8a, 0x2a, 0x35, 0x95, 0xb8, 0x70, 0xa8, 0x0c,
	0xca, 0x01, 0x0e, 0x95, 0x13, 0xaf, 0xa2, 0x55, 0x5d, 0xef, 0x36, 0x6b, 0x47, 0x6a, 0xab, 0xf2,
	0x75, 0xe5, 0x06, 0xbf, 0x00, 0xf1, 0x93, 0xf8, 0x0b, 0xfc, 0x05, 0xee, 0x68, 0x3f, 0x9c, 0xd8,
	0xa1, 0x01, 0x9a, 0x4b, 0x34, 0x33, 0xfb, 0x78, 0xf6, 0x99, 0x67, 0x76, 0x26, 0x00, 0x11, 0xe1,
	0x27, 0x1d, 0x36, 0xa6, 0x29, 0x45, 0xc0, 0x26, 0x9d, 0xc9, 0xce, 0x00, 0xa7, 0xe1, 0x4e, 0x73,
	0x6b, 0x44, 0xe9, 0x28, 0xc6, 0x7e, 0xc8, 0x88, 0x1f, 0x26, 0x09, 0x4d, 0xc3, 0x94, 0xd0, 0x84,
	0x2b, 0x64, 0x73, 0x43, 0x84, 0xa3, 0xf4, 0x9c, 0x61, 0xee, 0xcb, 0x5f, 0x15, 0xf7, 0x3e, 0x1a,
	0xf0, 0xcf, 0x21, 0xe1, 0x27, 0xbd, 0x31, 0x0e, 0x53, 0x1c, 0xe0, 0xb3, 0x0c, 0xf3, 0x14, 0xb9,
	0x50, 0x1d, 0xc6, 0x19, 0x4f, 0xf1, 0xd8, 0x35, 0x5a, 0x46, 0xbb, 0x1e, 0xe4, 0x2e, 0x42, 0x60,
	0x27, 0xe1, 0x29, 0x76, 0x4d, 0x19, 0x96, 0xb6, 0x88, 0x5d, 0xd0, 0x04, 0xbb, 0x96, 0x8a, 0x09,
	0x1b, 0x6d, 0x42, 0x5d, 0xf0, 0x3c, 0x16, 0x77, 0xb9, 0xb6, 0x3c, 0xa8, 0x89, 0xc0, 0xcb, 0x73,
	0x86, 0xd1, 0x7f, 0x50, 0xe5, 0xe4, 0x02, 0x1f, 0x8f, 0x06, 0x6e, 0xa5, 0x65, 0xb4, 0xad, 0xc0,
	0x11, 0xee, 0xb3, 0x81, 0xb7, 0xaf, 0xc8, 0x1c, 0xe2, 0x18, 0xff, 0x0d, 0x99, 0x35, 0xb0, 0x32,
	0x12, 0x69, 0x2e, 0xc2, 0xf4, 0xee, 0xc3, 0xaa, 0x48, 0xf0, 0x9c, 0xf0, 0xf4, 0x8f, 0x9f, 0x7b,
	0x17, 0xb0, 0x36, 0x03, 0x73, 0x46, 0x13, 0x8e, 0xd1, 0x3d, 0x70, 0x78, 0x1a, 0xa6, 0x19, 0x97,
	0xe0, 0x46, 0x77, 0xa5, 0xa3, 0x44, 0xeb, 0xbc, 0x90, 0xd1, 0x40, 0x9f, 0x8a, 0x9a, 0x4f, 0xb2,
	0xc1, 0x54, 0x07, 0x61, 0xa3, 0x6d, 0x70, 0xc6, 0x98, 0x67, 0x71, 0xea, 0x5a, 0x2d, 0xab, 0xdd,
	0xe8, 0xa2, 0xce, 0xac, 0x3d, 0x9d, 0x40, 0x9e, 0x04, 0x1a, 0xe1, 0x9d, 0x81, 0xa3, 0x22, 0xe8,
	0x2e, 0xd8, 0x42, 0x18, 0x7d, 0xdf, 0x5a, 0xf1, 0x1b, 0xc1, 0x2e, 0x90, 0xa7, 0xe8, 0x16, 0x98,
	0x6c, 0x22, 0x6f, 0x13, 0x9c, 0x0a, 0x98, 0xa3, 0x7e, 0x60, 0xb2, 0x09, 0xba, 0x03, 0x16, 0x9b,
	0x0c, 0x65, 0x0b, 0x1a, 0xdd, 0xd5, 0x32, 0xa0, 0x17, 0x88, 0x33, 0xef, 0x87, 0x01, 0xb6, 0xc8,
	0x38, 0xed, 0xa1, 0x51, 0xe8, 0xe1, 0x0a, 0x98, 0x53, 0x25, 0x4d, 0x12, 0xa1, 0x0d, 0x70, 0x58,
	0x9c, 0x8d, 0x48, 0xa2, 0xbb, 0xaa, 0xbd, 0x62, 0xeb, 0xec, 0x62, 0xeb, 0x44, 0x52, 0xd9, 0xeb,
	0x8a, 0x4a, 0x2a, 0xec, 0xe9, 0xc3, 0x70, 0x0a, 0x0f, 0x63, 0x63, 0x2a, 0x70, 0x55, 0x25, 0xd6,
	0x82, 0xae, 0x43, 0x25, 0xe3, 0x78, 0xcc, 0xdd, 0x5a, 0xcb, 0x6a, 0xd7, 0x03, 0xe5, 0x48, 0x99,
	0x49, 0x12, 0xb9, 0x75, 0x2d, 0x33, 0x49, 0x22, 0xd4, 0x84, 0x1a, 0x4e, 0x22, 0x46, 0x49, 0x92,
	0xba, 0xa0, 0x5e, 0x56, 0xee, 0x0b, 0x3c, 0xa1, 0x8c, 0xbb, 0x0d, 0xc9, 0x4d, 0xda, 0xde, 0x57,
	0x03, 0xcc, 0xa3, 0xfe, 0xb5, 0x55, 0x17, 0xaa, 0x31, 0x4b, 0xd5, 0xcc, 0x58, 0x5a, 0xf3, 0x2c,
	0x87, 0x71, 0x48, 0x4e, 0xf5, 0x93, 0x56, 0x8e, 0x40, 0x4f, 0x68, 0x9c, 0x9d, 0xe6, 0xd5, 0x6b,
	0xaf, 0x20, 0xa2, 0x53, 0x12, 0x31, 0xd7, 0xaa, 0x3a, 0xd3, 0xca, 0x7b, 0x67, 0x80, 0x75, 0xd4,
	0xef, 0xdd, 0x8c, 0xe6, 0x16, 0xd4, 0x05, 0x80, 0xb3, 0x70, 0x98, 0x8f, 0xdf, 0x2c, 0x50, 0x28,
	0xc2, 0x2e, 0x15, 0xb1, 0x80, 0xae, 0xf7, 0x1a, 0xfe, 0x55, 0xd3, 0xc7, 0x87, 0x63, 0x32, 0x58,
	0x72, 0x19, 0x2c, 0x78, 0x38, 0x5e, 0x04, 0xeb, 0xe5, 0xe4, 0x37, 0x1c, 0xb8, 0x7c, 0x4c, 0xcc,
	0xdf, 0x8d, 0x49, 0xf7, 0x8b, 0x0d, 0x15, 0xe1, 0x72, 0xf4, 0x16, 0x6c, 0x31, 0xd8, 0x68, 0x73,
	0x1e, 0x59, 0xd8, 0x0d, 0xcd, 0xad, 0xeb, 0x0f, 0x15, 0x35, 0x6f, 0xef, 0xc3, 0xb7, 0xef, 0x9f,
	0xcd, 0x47, 0xe8, 0xa1, 0x1f, 0x32, 0xc6, 0x87, 0x34, 0x52, 0xcb, 0x75, 0x18, 0xd3, 0x2c, 0xf2,
	0xf5, 0x57, 0xbe, 0xd6, 0x82, 0xfb, 0x97, 0xda, 0xba, 0xf2, 0x23, 0x49, 0xe0, 0x93, 0x01, 0xb5,
	0xbc, 0x5a, 0x74, 0x7b, 0xfe, 0xa2, 0x39, 0x91, 0x9b, 0xad, 0xc5, 0x00, 0xcd, 0xa6, 0x27, 0xd9,
	0x3c, 0x45, 0x4f, 0x96, 0x61, 0xe3, 0x5f, 0x8a, 0xe6, 0x5c, 0xa1, 0x37, 0xe0, 0xa8, 0x4d, 0x8f,
	0xfe, 0x9f, 0xbf, 0xb0, 0xf4, 0x0f, 0xd0, 0x5c, 0xcf, 0xdb, 0xd0, 0xa7, 0x24, 0x9a, 0x72, 0xd8,
	0x97, 0x1c, 0x1e, 0x7b, 0x4b, 0x29, 0xb2, 0x6b, 0x6c, 0xa3, 0xf7, 0x06, 0x38, 0x6a, 0xbb, 0xff,
	0x4a, 0xa0, 0xb4, 0xf5, 0x17, 0x10, 0x38, 0x90, 0x04, 0xf6, 0xb6, 0x77, 0x97, 0x13, 0x21, 0x23,
	0xd1, 0xd5, 0x41, 0xfd, 0x55, 0x55, 0x23, 0x07, 0x8e, 0xfc, 0x13, 0x7c, 0xf0, 0x33, 0x00, 0x00,
	0xff, 0xff, 0x46, 0x9d, 0x91, 0x21, 0x54, 0x07, 0x00, 0x00,
}
