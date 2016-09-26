// Code generated by protoc-gen-go.
// source: cluster.proto
// DO NOT EDIT!

/*
Package glusterfs is a generated protocol buffer package.

It is generated from these files:
	cluster.proto
	volume.proto

It has these top-level messages:
	Cluster
	ClusterListRequest
	ClusterListResponse
	ClusterDescribeRequest
	ClusterDescribeResponse
	ClusterCreateRequest
	ClusterDeleteRequest
	VolumeCreateRequest
	VolumeDeleteRequest
	VolumeListRequest
	VolumeListResponse
	Volume
*/
package glusterfs

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

type Cluster struct {
	Phid          string `protobuf:"bytes,1,opt,name=phid" json:"phid,omitempty"`
	Name          string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	KubeCluster   string `protobuf:"bytes,3,opt,name=kube_cluster,json=kubeCluster" json:"kube_cluster,omitempty"`
	Mood          string `protobuf:"bytes,4,opt,name=mood" json:"mood,omitempty"`
	KubeNamespace string `protobuf:"bytes,5,opt,name=kube_namespace,json=kubeNamespace" json:"kube_namespace,omitempty"`
	Replica       int64  `protobuf:"varint,6,opt,name=replica" json:"replica,omitempty"`
	BaculaEnabled int32  `protobuf:"varint,7,opt,name=BaculaEnabled,json=baculaEnabled" json:"BaculaEnabled,omitempty"`
	Created       string `protobuf:"bytes,8,opt,name=created" json:"created,omitempty"`
}

func (m *Cluster) Reset()                    { *m = Cluster{} }
func (m *Cluster) String() string            { return proto.CompactTextString(m) }
func (*Cluster) ProtoMessage()               {}
func (*Cluster) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type ClusterListRequest struct {
	KubeCluster string `protobuf:"bytes,1,opt,name=kube_cluster,json=kubeCluster" json:"kube_cluster,omitempty"`
}

func (m *ClusterListRequest) Reset()                    { *m = ClusterListRequest{} }
func (m *ClusterListRequest) String() string            { return proto.CompactTextString(m) }
func (*ClusterListRequest) ProtoMessage()               {}
func (*ClusterListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type ClusterListResponse struct {
	Status    *dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Glusterfs []*Cluster     `protobuf:"bytes,2,rep,name=glusterfs" json:"glusterfs,omitempty"`
}

func (m *ClusterListResponse) Reset()                    { *m = ClusterListResponse{} }
func (m *ClusterListResponse) String() string            { return proto.CompactTextString(m) }
func (*ClusterListResponse) ProtoMessage()               {}
func (*ClusterListResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ClusterListResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ClusterListResponse) GetGlusterfs() []*Cluster {
	if m != nil {
		return m.Glusterfs
	}
	return nil
}

type ClusterDescribeRequest struct {
	KubeCluster string `protobuf:"bytes,1,opt,name=kube_cluster,json=kubeCluster" json:"kube_cluster,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *ClusterDescribeRequest) Reset()                    { *m = ClusterDescribeRequest{} }
func (m *ClusterDescribeRequest) String() string            { return proto.CompactTextString(m) }
func (*ClusterDescribeRequest) ProtoMessage()               {}
func (*ClusterDescribeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type ClusterDescribeResponse struct {
	Status *dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
}

func (m *ClusterDescribeResponse) Reset()                    { *m = ClusterDescribeResponse{} }
func (m *ClusterDescribeResponse) String() string            { return proto.CompactTextString(m) }
func (*ClusterDescribeResponse) ProtoMessage()               {}
func (*ClusterDescribeResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ClusterDescribeResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

type ClusterCreateRequest struct {
	Name        string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Node        int64    `protobuf:"varint,2,opt,name=node" json:"node,omitempty"`
	Mood        string   `protobuf:"bytes,3,opt,name=mood" json:"mood,omitempty"`
	Disks       []string `protobuf:"bytes,4,rep,name=disks" json:"disks,omitempty"`
	KubeCluster string   `protobuf:"bytes,5,opt,name=kube_cluster,json=kubeCluster" json:"kube_cluster,omitempty"`
}

func (m *ClusterCreateRequest) Reset()                    { *m = ClusterCreateRequest{} }
func (m *ClusterCreateRequest) String() string            { return proto.CompactTextString(m) }
func (*ClusterCreateRequest) ProtoMessage()               {}
func (*ClusterCreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type ClusterDeleteRequest struct {
	KubeCluster string `protobuf:"bytes,1,opt,name=kube_cluster,json=kubeCluster" json:"kube_cluster,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *ClusterDeleteRequest) Reset()                    { *m = ClusterDeleteRequest{} }
func (m *ClusterDeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*ClusterDeleteRequest) ProtoMessage()               {}
func (*ClusterDeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func init() {
	proto.RegisterType((*Cluster)(nil), "glusterfs.Cluster")
	proto.RegisterType((*ClusterListRequest)(nil), "glusterfs.ClusterListRequest")
	proto.RegisterType((*ClusterListResponse)(nil), "glusterfs.ClusterListResponse")
	proto.RegisterType((*ClusterDescribeRequest)(nil), "glusterfs.ClusterDescribeRequest")
	proto.RegisterType((*ClusterDescribeResponse)(nil), "glusterfs.ClusterDescribeResponse")
	proto.RegisterType((*ClusterCreateRequest)(nil), "glusterfs.ClusterCreateRequest")
	proto.RegisterType((*ClusterDeleteRequest)(nil), "glusterfs.ClusterDeleteRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Clusters service

type ClustersClient interface {
	// Glusterfs cluster list. Needs to work with two modes.
	// First is to list all the glusterfs cluster through the
	// space with out considering the kubernetes cluster. if the
	// cluster_name is provided then list all the glusterfs cluster
	// with respect to the provided kube cluster space.
	List(ctx context.Context, in *ClusterListRequest, opts ...grpc.CallOption) (*ClusterListResponse, error)
	Describe(ctx context.Context, in *ClusterDescribeRequest, opts ...grpc.CallOption) (*ClusterDescribeResponse, error)
	Create(ctx context.Context, in *ClusterCreateRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
	Delete(ctx context.Context, in *ClusterDeleteRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
}

type clustersClient struct {
	cc *grpc.ClientConn
}

func NewClustersClient(cc *grpc.ClientConn) ClustersClient {
	return &clustersClient{cc}
}

func (c *clustersClient) List(ctx context.Context, in *ClusterListRequest, opts ...grpc.CallOption) (*ClusterListResponse, error) {
	out := new(ClusterListResponse)
	err := grpc.Invoke(ctx, "/glusterfs.Clusters/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clustersClient) Describe(ctx context.Context, in *ClusterDescribeRequest, opts ...grpc.CallOption) (*ClusterDescribeResponse, error) {
	out := new(ClusterDescribeResponse)
	err := grpc.Invoke(ctx, "/glusterfs.Clusters/Describe", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clustersClient) Create(ctx context.Context, in *ClusterCreateRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/glusterfs.Clusters/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clustersClient) Delete(ctx context.Context, in *ClusterDeleteRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/glusterfs.Clusters/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Clusters service

type ClustersServer interface {
	// Glusterfs cluster list. Needs to work with two modes.
	// First is to list all the glusterfs cluster through the
	// space with out considering the kubernetes cluster. if the
	// cluster_name is provided then list all the glusterfs cluster
	// with respect to the provided kube cluster space.
	List(context.Context, *ClusterListRequest) (*ClusterListResponse, error)
	Describe(context.Context, *ClusterDescribeRequest) (*ClusterDescribeResponse, error)
	Create(context.Context, *ClusterCreateRequest) (*dtypes.VoidResponse, error)
	Delete(context.Context, *ClusterDeleteRequest) (*dtypes.VoidResponse, error)
}

func RegisterClustersServer(s *grpc.Server, srv ClustersServer) {
	s.RegisterService(&_Clusters_serviceDesc, srv)
}

func _Clusters_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClusterListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClustersServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/glusterfs.Clusters/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClustersServer).List(ctx, req.(*ClusterListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Clusters_Describe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClusterDescribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClustersServer).Describe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/glusterfs.Clusters/Describe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClustersServer).Describe(ctx, req.(*ClusterDescribeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Clusters_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClusterCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClustersServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/glusterfs.Clusters/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClustersServer).Create(ctx, req.(*ClusterCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Clusters_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClusterDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClustersServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/glusterfs.Clusters/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClustersServer).Delete(ctx, req.(*ClusterDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Clusters_serviceDesc = grpc.ServiceDesc{
	ServiceName: "glusterfs.Clusters",
	HandlerType: (*ClustersServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _Clusters_List_Handler,
		},
		{
			MethodName: "Describe",
			Handler:    _Clusters_Describe_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _Clusters_Create_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Clusters_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("cluster.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 548 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xa4, 0x54, 0x5b, 0x6b, 0xd4, 0x40,
	0x14, 0x66, 0x9a, 0xbd, 0xf5, 0xd4, 0xed, 0xc3, 0x71, 0xa9, 0x21, 0x78, 0x6b, 0x50, 0x91, 0x3e,
	0x24, 0xb5, 0x8a, 0xe2, 0x0d, 0xd1, 0xad, 0x6f, 0x5e, 0x20, 0x82, 0xaf, 0x32, 0x9b, 0x8c, 0x6b,
	0x68, 0xba, 0x13, 0x77, 0x12, 0x41, 0x4a, 0x41, 0xd4, 0x37, 0x1f, 0x05, 0xff, 0x88, 0x3f, 0xc5,
	0xbf, 0xe0, 0x8b, 0xff, 0xc2, 0xb9, 0x25, 0xcd, 0x9a, 0x56, 0xb4, 0x7d, 0x59, 0x66, 0xbe, 0x6f,
	0xce, 0x99, 0x6f, 0xbe, 0xf3, 0x65, 0x61, 0x18, 0x67, 0xa5, 0x28, 0xd8, 0x3c, 0xc8, 0xe7, 0xbc,
	0xe0, 0xb8, 0x3c, 0x35, 0xdb, 0xd7, 0xc2, 0x3b, 0x3b, 0xe5, 0x7c, 0x9a, 0xb1, 0x90, 0xe6, 0x69,
	0x48, 0x67, 0x33, 0x5e, 0xd0, 0x22, 0xe5, 0x33, 0x61, 0x0e, 0x7a, 0x6b, 0x0a, 0x4e, 0x8a, 0xf7,
	0x39, 0x13, 0xa1, 0xfe, 0x35, 0xb8, 0xff, 0x8b, 0x40, 0x7f, 0x6c, 0x7a, 0x20, 0x42, 0x27, 0x7f,
	0x93, 0x26, 0x2e, 0xb9, 0x48, 0xae, 0x2e, 0x47, 0x7a, 0xad, 0xb0, 0x19, 0xdd, 0x65, 0xee, 0x92,
	0xc1, 0xd4, 0x1a, 0xd7, 0xe1, 0xd4, 0x4e, 0x39, 0x61, 0xaf, 0xac, 0x14, 0xd7, 0xd1, 0xdc, 0x8a,
	0xc2, 0x1a, 0xad, 0x76, 0x39, 0x4f, 0xdc, 0x8e, 0x29, 0x53, 0x6b, 0xbc, 0x0c, 0xab, 0xba, 0x4c,
	0xf5, 0x10, 0x39, 0x8d, 0x99, 0xdb, 0xd5, 0xec, 0x50, 0xa1, 0xcf, 0x2a, 0x10, 0x5d, 0xe8, 0xcf,
	0x59, 0x9e, 0xa5, 0x31, 0x75, 0x7b, 0x92, 0x77, 0xa2, 0x6a, 0x8b, 0x97, 0x60, 0xf8, 0x88, 0xc6,
	0x65, 0x46, 0x1f, 0xcf, 0xe8, 0x24, 0x63, 0x89, 0xdb, 0x97, 0x7c, 0x37, 0x1a, 0x4e, 0x9a, 0xa0,
	0xaa, 0x8f, 0xe7, 0x8c, 0x16, 0x92, 0x1f, 0xe8, 0xfe, 0xd5, 0xd6, 0xbf, 0x05, 0x68, 0xf5, 0x3d,
	0x49, 0x45, 0x11, 0xb1, 0xb7, 0x25, 0x13, 0x45, 0xeb, 0x35, 0xa4, 0xf5, 0x1a, 0x9f, 0xc3, 0xe9,
	0x85, 0x42, 0x91, 0x4b, 0x63, 0x19, 0x5e, 0x81, 0x9e, 0x90, 0x2e, 0x97, 0x42, 0xd7, 0xac, 0x6c,
	0xad, 0x06, 0xc6, 0xe0, 0xe0, 0x85, 0x46, 0x23, 0xcb, 0xe2, 0x26, 0x1c, 0x8c, 0x49, 0x1a, 0xe9,
	0xc8, 0xa3, 0x18, 0xd4, 0x48, 0x60, 0x5b, 0x47, 0x07, 0x87, 0xfc, 0xe7, 0xb0, 0x66, 0xd1, 0x6d,
	0x26, 0xe2, 0x79, 0x3a, 0x61, 0xff, 0xae, 0xf6, 0xb0, 0x91, 0xf9, 0x0f, 0xe1, 0x4c, 0xab, 0xe1,
	0xff, 0xbd, 0xc2, 0xff, 0x42, 0x60, 0x64, 0x7b, 0x8c, 0xb5, 0xa1, 0x95, 0xa4, 0xea, 0x3e, 0xd2,
	0x88, 0x88, 0xc2, 0x78, 0x62, 0x34, 0x38, 0x91, 0x5e, 0xd7, 0x99, 0x70, 0x1a, 0x99, 0x18, 0x41,
	0x37, 0x49, 0xc5, 0x8e, 0x90, 0x41, 0x71, 0x24, 0x68, 0x36, 0xad, 0x47, 0x76, 0xdb, 0x23, 0x79,
	0x5a, 0x8b, 0xd9, 0x66, 0x19, 0x2b, 0x4e, 0xe8, 0xcf, 0xd6, 0xf7, 0x0e, 0x0c, 0x2c, 0x2f, 0xf0,
	0x13, 0x81, 0x8e, 0x1a, 0x34, 0x9e, 0x6b, 0x4f, 0xa9, 0x91, 0x1c, 0xef, 0xfc, 0x51, 0xb4, 0x71,
	0xd6, 0xbf, 0xf7, 0xf1, 0xc7, 0xcf, 0xaf, 0x4b, 0x37, 0xf1, 0x86, 0xfc, 0x26, 0x73, 0x11, 0x4b,
	0x0f, 0xf4, 0xc7, 0x59, 0x17, 0x85, 0xef, 0x36, 0x83, 0x6b, 0xa1, 0x55, 0x2a, 0xc2, 0xbd, 0xa6,
	0xf0, 0x7d, 0xfc, 0x46, 0x60, 0x50, 0x0d, 0x0b, 0xd7, 0xdb, 0x57, 0xfd, 0x91, 0x0c, 0xcf, 0xff,
	0xdb, 0x11, 0xab, 0x68, 0xac, 0x15, 0xdd, 0xc7, 0xbb, 0xc7, 0x51, 0x14, 0xee, 0x29, 0xab, 0xf6,
	0xf1, 0x03, 0x81, 0x9e, 0x49, 0x00, 0x5e, 0x68, 0xdf, 0xb9, 0x90, 0x0d, 0x6f, 0x54, 0x85, 0xe9,
	0x25, 0x4f, 0x93, 0x5a, 0xc6, 0x03, 0x2d, 0xe3, 0xb6, 0x77, 0x2c, 0x63, 0xee, 0x90, 0x0d, 0xfc,
	0x2c, 0x25, 0x98, 0xb9, 0x1f, 0x26, 0x61, 0x21, 0x11, 0x47, 0x48, 0xb0, 0x4e, 0x6c, 0x9c, 0xc4,
	0x89, 0x49, 0x4f, 0xff, 0x87, 0x5e, 0xff, 0x1d, 0x00, 0x00, 0xff, 0xff, 0xf0, 0xeb, 0x87, 0xb1,
	0x95, 0x05, 0x00, 0x00,
}