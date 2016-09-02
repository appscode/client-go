// Code generated by protoc-gen-go.
// source: ssh.proto
// DO NOT EDIT!

/*
Package ssh is a generated protocol buffer package.

It is generated from these files:
	ssh.proto

It has these top-level messages:
	SSHGetRequest
	SSHGetResponse
	SSHKey
*/
package ssh

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

// Use specific requests for protos
type SSHGetRequest struct {
	Namespace    string `protobuf:"bytes,1,opt,name=namespace" json:"namespace,omitempty"`
	ClusterName  string `protobuf:"bytes,2,opt,name=cluster_name,json=clusterName" json:"cluster_name,omitempty"`
	InstanceName string `protobuf:"bytes,3,opt,name=instance_name,json=instanceName" json:"instance_name,omitempty"`
}

func (m *SSHGetRequest) Reset()                    { *m = SSHGetRequest{} }
func (m *SSHGetRequest) String() string            { return proto.CompactTextString(m) }
func (*SSHGetRequest) ProtoMessage()               {}
func (*SSHGetRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// return phid ?
type SSHGetResponse struct {
	Status       *dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	SshKey       *SSHKey        `protobuf:"bytes,2,opt,name=ssh_key,json=sshKey" json:"ssh_key,omitempty"`
	InstanceAddr string         `protobuf:"bytes,3,opt,name=instance_addr,json=instanceAddr" json:"instance_addr,omitempty"`
	InstancePort int32          `protobuf:"varint,4,opt,name=instance_port,json=instancePort" json:"instance_port,omitempty"`
	User         string         `protobuf:"bytes,5,opt,name=user" json:"user,omitempty"`
}

func (m *SSHGetResponse) Reset()                    { *m = SSHGetResponse{} }
func (m *SSHGetResponse) String() string            { return proto.CompactTextString(m) }
func (*SSHGetResponse) ProtoMessage()               {}
func (*SSHGetResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SSHGetResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *SSHGetResponse) GetSshKey() *SSHKey {
	if m != nil {
		return m.SshKey
	}
	return nil
}

type SSHKey struct {
	PublicKey          []byte `protobuf:"bytes,1,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	PrivateKey         []byte `protobuf:"bytes,2,opt,name=private_key,json=privateKey,proto3" json:"private_key,omitempty"`
	AwsFingerprint     string `protobuf:"bytes,3,opt,name=aws_fingerprint,json=awsFingerprint" json:"aws_fingerprint,omitempty"`
	OpensshFingerprint string `protobuf:"bytes,4,opt,name=openssh_fingerprint,json=opensshFingerprint" json:"openssh_fingerprint,omitempty"`
}

func (m *SSHKey) Reset()                    { *m = SSHKey{} }
func (m *SSHKey) String() string            { return proto.CompactTextString(m) }
func (*SSHKey) ProtoMessage()               {}
func (*SSHKey) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func init() {
	proto.RegisterType((*SSHGetRequest)(nil), "ssh.SSHGetRequest")
	proto.RegisterType((*SSHGetResponse)(nil), "ssh.SSHGetResponse")
	proto.RegisterType((*SSHKey)(nil), "ssh.SSHKey")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for SSH service

type SSHClient interface {
	Get(ctx context.Context, in *SSHGetRequest, opts ...grpc.CallOption) (*SSHGetResponse, error)
}

type sSHClient struct {
	cc *grpc.ClientConn
}

func NewSSHClient(cc *grpc.ClientConn) SSHClient {
	return &sSHClient{cc}
}

func (c *sSHClient) Get(ctx context.Context, in *SSHGetRequest, opts ...grpc.CallOption) (*SSHGetResponse, error) {
	out := new(SSHGetResponse)
	err := grpc.Invoke(ctx, "/ssh.SSH/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SSH service

type SSHServer interface {
	Get(context.Context, *SSHGetRequest) (*SSHGetResponse, error)
}

func RegisterSSHServer(s *grpc.Server, srv SSHServer) {
	s.RegisterService(&_SSH_serviceDesc, srv)
}

func _SSH_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SSHGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SSHServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ssh.SSH/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SSHServer).Get(ctx, req.(*SSHGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SSH_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ssh.SSH",
	HandlerType: (*SSHServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _SSH_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("ssh.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 396 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x5c, 0x92, 0x4f, 0xae, 0xd3, 0x30,
	0x10, 0x87, 0x95, 0x97, 0xbe, 0xa2, 0x4c, 0xf2, 0x8a, 0xe4, 0x27, 0x55, 0x51, 0x55, 0xfe, 0x05,
	0x04, 0xac, 0x12, 0x28, 0x27, 0x60, 0x03, 0x48, 0x48, 0x15, 0x4a, 0xc5, 0xba, 0x72, 0x93, 0xa1,
	0x8d, 0x28, 0xb6, 0x89, 0x9d, 0xa2, 0x6e, 0xb9, 0x02, 0x37, 0xe0, 0x1c, 0xdc, 0x82, 0x2b, 0x70,
	0x10, 0xc6, 0x76, 0x4a, 0x5a, 0x36, 0xad, 0xf5, 0xf9, 0x8b, 0xe7, 0xe7, 0xf1, 0x40, 0xa4, 0xf5,
	0x2e, 0x57, 0xad, 0x34, 0x92, 0x85, 0xb4, 0x9c, 0xcd, 0xb7, 0x52, 0x6e, 0xf7, 0x58, 0x70, 0xd5,
	0x14, 0x5c, 0x08, 0x69, 0xb8, 0x69, 0xa4, 0xd0, 0x5e, 0x99, 0x4d, 0x2d, 0xae, 0xcd, 0x51, 0xa1,
	0x2e, 0xdc, 0xaf, 0xe7, 0x59, 0x07, 0x37, 0xab, 0xd5, 0xbb, 0xb7, 0x68, 0x4a, 0xfc, 0xda, 0xa1,
	0x36, 0x6c, 0x0e, 0x91, 0xe0, 0x5f, 0x50, 0x2b, 0x5e, 0x61, 0x1a, 0x3c, 0x0c, 0x9e, 0x47, 0xe5,
	0x00, 0xd8, 0x23, 0x48, 0xaa, 0x7d, 0xa7, 0x0d, 0xb6, 0x6b, 0x0b, 0xd3, 0x2b, 0x27, 0xc4, 0x3d,
	0x5b, 0x12, 0x62, 0x8f, 0xe1, 0xa6, 0x11, 0xda, 0x70, 0x51, 0xa1, 0x77, 0x42, 0xe7, 0x24, 0x27,
	0x68, 0xa5, 0xec, 0x57, 0x00, 0x93, 0x53, 0x5d, 0xad, 0x28, 0x26, 0xb2, 0xa7, 0x30, 0x26, 0xc1,
	0x74, 0xda, 0x55, 0x8d, 0x17, 0x93, 0xdc, 0xc7, 0xcd, 0x57, 0x8e, 0x96, 0xfd, 0x2e, 0x7b, 0x02,
	0x77, 0xe8, 0xba, 0xeb, 0xcf, 0x78, 0x74, 0xd5, 0xe3, 0x45, 0x9c, 0xdb, 0x4e, 0xd0, 0x69, 0xef,
	0xf1, 0x48, 0x96, 0xde, 0xd1, 0xff, 0x45, 0x0a, 0x5e, 0xd7, 0xed, 0xff, 0x29, 0x5e, 0x13, 0xbb,
	0x90, 0x94, 0x6c, 0x4d, 0x3a, 0x22, 0xe9, 0x7a, 0x90, 0x3e, 0x10, 0x63, 0x0c, 0x46, 0x9d, 0xc6,
	0x36, 0xbd, 0x76, 0x07, 0xb8, 0x75, 0xf6, 0x33, 0x80, 0xb1, 0x2f, 0xc8, 0xee, 0x01, 0xa8, 0x6e,
	0xb3, 0x6f, 0x2a, 0x97, 0xc8, 0x46, 0x4f, 0xca, 0xc8, 0x13, 0xbb, 0xfd, 0x00, 0x62, 0xd5, 0x36,
	0x07, 0x6e, 0xf0, 0x5f, 0xe2, 0xa4, 0x84, 0x1e, 0x59, 0xe1, 0x19, 0xdc, 0xe5, 0xdf, 0xf4, 0xfa,
	0x53, 0x23, 0xb6, 0xd8, 0x12, 0x17, 0xa6, 0x8f, 0x3a, 0x21, 0xfc, 0x66, 0xa0, 0xac, 0x80, 0x5b,
	0xa9, 0x50, 0xd8, 0xbb, 0x9f, 0xcb, 0x23, 0x27, 0xb3, 0x7e, 0xeb, 0xec, 0x83, 0xc5, 0x47, 0x08,
	0x29, 0x23, 0x5b, 0x42, 0x48, 0x6d, 0x66, 0xec, 0xd4, 0xa5, 0xe1, 0xad, 0x67, 0xb7, 0x17, 0xcc,
	0xbf, 0x43, 0x76, 0xff, 0xfb, 0xef, 0x3f, 0x3f, 0xae, 0x52, 0x36, 0xa5, 0x49, 0x52, 0xba, 0x92,
	0xb5, 0x1f, 0x29, 0x32, 0x8b, 0xc3, 0x8b, 0xfc, 0xe5, 0x66, 0xec, 0x06, 0xe7, 0xd5, 0xdf, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x5f, 0x01, 0x35, 0xc3, 0x80, 0x02, 0x00, 0x00,
}
