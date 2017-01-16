// Code generated by protoc-gen-go.
// source: agent.proto
// DO NOT EDIT!

/*
Package v1beta1 is a generated protocol buffer package.

It is generated from these files:
	agent.proto
	metadata.proto

It has these top-level messages:
	AgentListResponse
	Agent
	AgentCreateRequest
	AgentCreateResponse
	AgentDeleteRequest
	ServerInfoResponse
*/
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

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type AgentListResponse struct {
	Status *appscode_dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Agents []*Agent                `protobuf:"bytes,2,rep,name=agents" json:"agents,omitempty"`
}

func (m *AgentListResponse) Reset()                    { *m = AgentListResponse{} }
func (m *AgentListResponse) String() string            { return proto.CompactTextString(m) }
func (*AgentListResponse) ProtoMessage()               {}
func (*AgentListResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *AgentListResponse) GetStatus() *appscode_dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *AgentListResponse) GetAgents() []*Agent {
	if m != nil {
		return m.Agents
	}
	return nil
}

type Agent struct {
	Name        string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Status      string `protobuf:"bytes,2,opt,name=status" json:"status,omitempty"`
	AgentStatus string `protobuf:"bytes,3,opt,name=agent_status,json=agentStatus" json:"agent_status,omitempty"`
	IsRefreshed int32  `protobuf:"varint,4,opt,name=is_refreshed,json=isRefreshed" json:"is_refreshed,omitempty"`
	CreatedAt   int64  `protobuf:"varint,5,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
	UpdatedAt   int64  `protobuf:"varint,6,opt,name=updated_at,json=updatedAt" json:"updated_at,omitempty"`
}

func (m *Agent) Reset()                    { *m = Agent{} }
func (m *Agent) String() string            { return proto.CompactTextString(m) }
func (*Agent) ProtoMessage()               {}
func (*Agent) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Agent) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Agent) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Agent) GetAgentStatus() string {
	if m != nil {
		return m.AgentStatus
	}
	return ""
}

func (m *Agent) GetIsRefreshed() int32 {
	if m != nil {
		return m.IsRefreshed
	}
	return 0
}

func (m *Agent) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *Agent) GetUpdatedAt() int64 {
	if m != nil {
		return m.UpdatedAt
	}
	return 0
}

type AgentCreateRequest struct {
	Name            string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Role            string `protobuf:"bytes,2,opt,name=role" json:"role,omitempty"`
	ExternalIp      string `protobuf:"bytes,3,opt,name=external_ip,json=externalIp" json:"external_ip,omitempty"`
	InternalIp      string `protobuf:"bytes,4,opt,name=internal_ip,json=internalIp" json:"internal_ip,omitempty"`
	SshUser         string `protobuf:"bytes,5,opt,name=ssh_user,json=sshUser" json:"ssh_user,omitempty"`
	SshPort         string `protobuf:"bytes,6,opt,name=ssh_port,json=sshPort" json:"ssh_port,omitempty"`
	JenkinsJnlpPort string `protobuf:"bytes,7,opt,name=jenkins_jnlp_port,json=jenkinsJnlpPort" json:"jenkins_jnlp_port,omitempty"`
	GitSshPublicKey string `protobuf:"bytes,8,opt,name=git_ssh_public_key,json=gitSshPublicKey" json:"git_ssh_public_key,omitempty"`
}

func (m *AgentCreateRequest) Reset()                    { *m = AgentCreateRequest{} }
func (m *AgentCreateRequest) String() string            { return proto.CompactTextString(m) }
func (*AgentCreateRequest) ProtoMessage()               {}
func (*AgentCreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *AgentCreateRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AgentCreateRequest) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

func (m *AgentCreateRequest) GetExternalIp() string {
	if m != nil {
		return m.ExternalIp
	}
	return ""
}

func (m *AgentCreateRequest) GetInternalIp() string {
	if m != nil {
		return m.InternalIp
	}
	return ""
}

func (m *AgentCreateRequest) GetSshUser() string {
	if m != nil {
		return m.SshUser
	}
	return ""
}

func (m *AgentCreateRequest) GetSshPort() string {
	if m != nil {
		return m.SshPort
	}
	return ""
}

func (m *AgentCreateRequest) GetJenkinsJnlpPort() string {
	if m != nil {
		return m.JenkinsJnlpPort
	}
	return ""
}

func (m *AgentCreateRequest) GetGitSshPublicKey() string {
	if m != nil {
		return m.GitSshPublicKey
	}
	return ""
}

type AgentCreateResponse struct {
	Status                 *appscode_dtypes.Status          `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Namespace              string                           `protobuf:"bytes,2,opt,name=namespace" json:"namespace,omitempty"`
	SshAuthorizedPublicKey string                           `protobuf:"bytes,3,opt,name=ssh_authorized_public_key,json=sshAuthorizedPublicKey" json:"ssh_authorized_public_key,omitempty"`
	GitHostname            string                           `protobuf:"bytes,4,opt,name=git_hostname,json=gitHostname" json:"git_hostname,omitempty"`
	GitHostPublicKey       string                           `protobuf:"bytes,5,opt,name=git_host_public_key,json=gitHostPublicKey" json:"git_host_public_key,omitempty"`
	GitUser                *AgentCreateResponse_ConduitUser `protobuf:"bytes,6,opt,name=git_user,json=gitUser" json:"git_user,omitempty"`
}

func (m *AgentCreateResponse) Reset()                    { *m = AgentCreateResponse{} }
func (m *AgentCreateResponse) String() string            { return proto.CompactTextString(m) }
func (*AgentCreateResponse) ProtoMessage()               {}
func (*AgentCreateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *AgentCreateResponse) GetStatus() *appscode_dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *AgentCreateResponse) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *AgentCreateResponse) GetSshAuthorizedPublicKey() string {
	if m != nil {
		return m.SshAuthorizedPublicKey
	}
	return ""
}

func (m *AgentCreateResponse) GetGitHostname() string {
	if m != nil {
		return m.GitHostname
	}
	return ""
}

func (m *AgentCreateResponse) GetGitHostPublicKey() string {
	if m != nil {
		return m.GitHostPublicKey
	}
	return ""
}

func (m *AgentCreateResponse) GetGitUser() *AgentCreateResponse_ConduitUser {
	if m != nil {
		return m.GitUser
	}
	return nil
}

type AgentCreateResponse_ConduitUser struct {
	Phid     string `protobuf:"bytes,1,opt,name=phid" json:"phid,omitempty"`
	UserName string `protobuf:"bytes,2,opt,name=user_name,json=userName" json:"user_name,omitempty"`
	Token    string `protobuf:"bytes,3,opt,name=token" json:"token,omitempty"`
	Email    string `protobuf:"bytes,4,opt,name=email" json:"email,omitempty"`
}

func (m *AgentCreateResponse_ConduitUser) Reset()         { *m = AgentCreateResponse_ConduitUser{} }
func (m *AgentCreateResponse_ConduitUser) String() string { return proto.CompactTextString(m) }
func (*AgentCreateResponse_ConduitUser) ProtoMessage()    {}
func (*AgentCreateResponse_ConduitUser) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{3, 0}
}

func (m *AgentCreateResponse_ConduitUser) GetPhid() string {
	if m != nil {
		return m.Phid
	}
	return ""
}

func (m *AgentCreateResponse_ConduitUser) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *AgentCreateResponse_ConduitUser) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *AgentCreateResponse_ConduitUser) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

type AgentDeleteRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *AgentDeleteRequest) Reset()                    { *m = AgentDeleteRequest{} }
func (m *AgentDeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*AgentDeleteRequest) ProtoMessage()               {}
func (*AgentDeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *AgentDeleteRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*AgentListResponse)(nil), "appscode.ci.v1beta1.AgentListResponse")
	proto.RegisterType((*Agent)(nil), "appscode.ci.v1beta1.Agent")
	proto.RegisterType((*AgentCreateRequest)(nil), "appscode.ci.v1beta1.AgentCreateRequest")
	proto.RegisterType((*AgentCreateResponse)(nil), "appscode.ci.v1beta1.AgentCreateResponse")
	proto.RegisterType((*AgentCreateResponse_ConduitUser)(nil), "appscode.ci.v1beta1.AgentCreateResponse.ConduitUser")
	proto.RegisterType((*AgentDeleteRequest)(nil), "appscode.ci.v1beta1.AgentDeleteRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Agents service

type AgentsClient interface {
	List(ctx context.Context, in *appscode_dtypes.VoidRequest, opts ...grpc.CallOption) (*AgentListResponse, error)
	Create(ctx context.Context, in *AgentCreateRequest, opts ...grpc.CallOption) (*AgentCreateResponse, error)
	Delete(ctx context.Context, in *AgentDeleteRequest, opts ...grpc.CallOption) (*appscode_dtypes.VoidResponse, error)
}

type agentsClient struct {
	cc *grpc.ClientConn
}

func NewAgentsClient(cc *grpc.ClientConn) AgentsClient {
	return &agentsClient{cc}
}

func (c *agentsClient) List(ctx context.Context, in *appscode_dtypes.VoidRequest, opts ...grpc.CallOption) (*AgentListResponse, error) {
	out := new(AgentListResponse)
	err := grpc.Invoke(ctx, "/appscode.ci.v1beta1.Agents/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentsClient) Create(ctx context.Context, in *AgentCreateRequest, opts ...grpc.CallOption) (*AgentCreateResponse, error) {
	out := new(AgentCreateResponse)
	err := grpc.Invoke(ctx, "/appscode.ci.v1beta1.Agents/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentsClient) Delete(ctx context.Context, in *AgentDeleteRequest, opts ...grpc.CallOption) (*appscode_dtypes.VoidResponse, error) {
	out := new(appscode_dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/appscode.ci.v1beta1.Agents/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Agents service

type AgentsServer interface {
	List(context.Context, *appscode_dtypes.VoidRequest) (*AgentListResponse, error)
	Create(context.Context, *AgentCreateRequest) (*AgentCreateResponse, error)
	Delete(context.Context, *AgentDeleteRequest) (*appscode_dtypes.VoidResponse, error)
}

func RegisterAgentsServer(s *grpc.Server, srv AgentsServer) {
	s.RegisterService(&_Agents_serviceDesc, srv)
}

func _Agents_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(appscode_dtypes.VoidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentsServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.ci.v1beta1.Agents/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentsServer).List(ctx, req.(*appscode_dtypes.VoidRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agents_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AgentCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentsServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.ci.v1beta1.Agents/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentsServer).Create(ctx, req.(*AgentCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agents_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AgentDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentsServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.ci.v1beta1.Agents/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentsServer).Delete(ctx, req.(*AgentDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Agents_serviceDesc = grpc.ServiceDesc{
	ServiceName: "appscode.ci.v1beta1.Agents",
	HandlerType: (*AgentsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _Agents_List_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _Agents_Create_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Agents_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "agent.proto",
}

func init() { proto.RegisterFile("agent.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 739 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x54, 0xcb, 0x6e, 0x13, 0x49,
	0x14, 0x55, 0xfb, 0xd1, 0xb6, 0xaf, 0x47, 0x9a, 0x49, 0x65, 0x94, 0x71, 0x9c, 0x64, 0xe2, 0x69,
	0x69, 0x26, 0x9e, 0x8c, 0xa6, 0x5b, 0xf1, 0x64, 0x16, 0xb0, 0x73, 0xc2, 0x82, 0x00, 0x02, 0xab,
	0x23, 0x58, 0xb0, 0x69, 0x55, 0xdc, 0x85, 0xbb, 0x92, 0x4e, 0x57, 0xd1, 0x55, 0x46, 0x09, 0x8f,
	0x4d, 0x36, 0xac, 0x60, 0xc3, 0x17, 0xf0, 0x03, 0xec, 0x10, 0x1f, 0xc2, 0x2f, 0xf0, 0x21, 0xa8,
	0x1e, 0xed, 0x87, 0x88, 0x13, 0xc4, 0xa6, 0x55, 0x7d, 0xce, 0xa9, 0xaa, 0x73, 0x4f, 0x5d, 0x5d,
	0x68, 0xe2, 0x11, 0xc9, 0xa4, 0xcf, 0x73, 0x26, 0x19, 0x5a, 0xc6, 0x9c, 0x8b, 0x21, 0x8b, 0x89,
	0x3f, 0xa4, 0xfe, 0xb3, 0x9d, 0x23, 0x22, 0xf1, 0x4e, 0x7b, 0x7d, 0xc4, 0xd8, 0x28, 0x25, 0x01,
	0xe6, 0x34, 0xc0, 0x59, 0xc6, 0x24, 0x96, 0x94, 0x65, 0xc2, 0x6c, 0x69, 0xff, 0x5e, 0x6c, 0x59,
	0xc0, 0x6f, 0xce, 0xf1, 0xb1, 0x3c, 0xe7, 0x44, 0x04, 0xfa, 0x6b, 0x04, 0xde, 0x19, 0x2c, 0xf5,
	0x95, 0x85, 0x7b, 0x54, 0xc8, 0x90, 0x08, 0xce, 0x32, 0x41, 0x50, 0x00, 0xae, 0x90, 0x58, 0x8e,
	0x45, 0xcb, 0xe9, 0x38, 0xdd, 0x66, 0xef, 0x37, 0x7f, 0xe2, 0xcc, 0x1c, 0xe1, 0x1f, 0x6a, 0x3a,
	0xb4, 0x32, 0xd4, 0x03, 0x57, 0x17, 0x22, 0x5a, 0xa5, 0x4e, 0xb9, 0xdb, 0xec, 0xb5, 0xfd, 0x4b,
	0x4a, 0xf1, 0xf5, 0x45, 0xa1, 0x55, 0x7a, 0x9f, 0x1c, 0xa8, 0x6a, 0x04, 0x21, 0xa8, 0x64, 0xf8,
	0x94, 0xe8, 0xcb, 0x1a, 0xa1, 0x5e, 0xa3, 0x95, 0x89, 0x85, 0x92, 0x46, 0x8b, 0x9b, 0xfe, 0x80,
	0x9f, 0xf4, 0xfe, 0xc8, 0xb2, 0x65, 0xcd, 0x9a, 0x18, 0x0f, 0x27, 0x12, 0x2a, 0xa2, 0x9c, 0x3c,
	0xc9, 0x89, 0x48, 0x48, 0xdc, 0xaa, 0x74, 0x9c, 0x6e, 0x35, 0x6c, 0x52, 0x11, 0x16, 0x10, 0xda,
	0x00, 0x18, 0xe6, 0x04, 0x4b, 0x12, 0x47, 0x58, 0xb6, 0xaa, 0x1d, 0xa7, 0x5b, 0x0e, 0x1b, 0x16,
	0xe9, 0x4b, 0x45, 0x8f, 0x79, 0x5c, 0xd0, 0xae, 0xa1, 0x2d, 0xd2, 0x97, 0xde, 0x9b, 0x12, 0x20,
	0xed, 0x7c, 0x5f, 0xef, 0x08, 0xc9, 0xd3, 0x31, 0x11, 0x97, 0x97, 0x81, 0xa0, 0x92, 0xb3, 0x94,
	0xd8, 0x22, 0xf4, 0x1a, 0x6d, 0x42, 0x93, 0x9c, 0x49, 0x92, 0x67, 0x38, 0x8d, 0x28, 0xb7, 0x15,
	0x40, 0x01, 0x1d, 0x70, 0x25, 0xa0, 0xd9, 0x54, 0x50, 0x31, 0x82, 0x02, 0x3a, 0xe0, 0x68, 0x15,
	0xea, 0x42, 0x24, 0xd1, 0x58, 0x90, 0x5c, 0x9b, 0x6f, 0x84, 0x35, 0x21, 0x92, 0x87, 0x82, 0xe4,
	0x05, 0xc5, 0x59, 0x6e, 0x8c, 0x1b, 0x6a, 0xc0, 0x72, 0x89, 0xb6, 0x61, 0xe9, 0x98, 0x64, 0x27,
	0x34, 0x13, 0xd1, 0x71, 0x96, 0x72, 0xa3, 0xa9, 0x69, 0xcd, 0xcf, 0x96, 0xb8, 0x93, 0xa5, 0x5c,
	0x6b, 0xff, 0x01, 0x34, 0xa2, 0x32, 0xd2, 0x47, 0x8d, 0x8f, 0x52, 0x3a, 0x8c, 0x4e, 0xc8, 0x79,
	0xab, 0x6e, 0xc4, 0x23, 0x2a, 0x0f, 0x45, 0x32, 0xd0, 0xf8, 0x5d, 0x72, 0xee, 0x7d, 0x28, 0xc3,
	0xf2, 0x5c, 0x1e, 0x3f, 0xda, 0x46, 0xeb, 0xd0, 0x50, 0xa9, 0x09, 0x8e, 0x87, 0x45, 0x64, 0x53,
	0x00, 0xdd, 0x80, 0x55, 0xe5, 0x07, 0x8f, 0x65, 0xc2, 0x72, 0xfa, 0x9c, 0xc4, 0xb3, 0xd6, 0x4c,
	0x8a, 0x2b, 0x42, 0x24, 0xfd, 0x09, 0x3f, 0x71, 0xa8, 0x5a, 0x42, 0x95, 0x93, 0x30, 0x21, 0xf5,
	0x13, 0x99, 0x48, 0x9b, 0x23, 0x2a, 0x6f, 0x5b, 0x08, 0xfd, 0x0b, 0xcb, 0x85, 0x64, 0xf6, 0x5c,
	0x13, 0xef, 0x2f, 0x56, 0x39, 0x3d, 0xf1, 0x01, 0xd4, 0x95, 0x5c, 0x3f, 0x81, 0xab, 0xab, 0xdb,
	0x5d, 0xdc, 0xf3, 0xf3, 0xb9, 0xf8, 0xfb, 0x2c, 0x8b, 0xc7, 0x54, 0xaa, 0xf7, 0x0a, 0x6b, 0x23,
	0xb3, 0x68, 0x1f, 0x43, 0x73, 0x06, 0x57, 0x8d, 0xc3, 0x13, 0x1a, 0x17, 0xcd, 0xa4, 0xd6, 0x68,
	0x0d, 0x1a, 0xea, 0xbe, 0x48, 0x97, 0x60, 0xe2, 0xa9, 0x2b, 0xe0, 0xbe, 0xf2, 0xff, 0x2b, 0x54,
	0x25, 0x3b, 0x21, 0x99, 0x4d, 0xc2, 0xfc, 0x28, 0x94, 0x9c, 0x62, 0x9a, 0xda, 0x8a, 0xcd, 0x8f,
	0xd7, 0xb5, 0xfd, 0x7b, 0x8b, 0xa4, 0xe4, 0xca, 0xfe, 0xed, 0xbd, 0x2f, 0x83, 0xab, 0xa5, 0x02,
	0xbd, 0x84, 0x8a, 0x1a, 0x12, 0x68, 0xfd, 0x9b, 0x57, 0x7c, 0xc4, 0x68, 0x6c, 0x0f, 0x69, 0xff,
	0xb5, 0x38, 0x85, 0xd9, 0x11, 0xe3, 0xf9, 0x17, 0x1f, 0x5b, 0xa5, 0xba, 0x73, 0xf1, 0xf9, 0xcb,
	0xbb, 0x92, 0x87, 0x3a, 0x41, 0x34, 0x37, 0xa8, 0x86, 0x34, 0xb0, 0x5b, 0x03, 0x33, 0x2d, 0xd0,
	0x5b, 0x07, 0x5c, 0x13, 0x23, 0xda, 0xba, 0x3e, 0x68, 0xe3, 0xa5, 0xfb, 0xbd, 0x2f, 0xe2, 0xed,
	0xcc, 0xb8, 0xf9, 0xd3, 0xbb, 0xd6, 0xcd, 0x4d, 0x67, 0x1b, 0xbd, 0x76, 0xc0, 0x35, 0xf9, 0x5d,
	0x65, 0x68, 0x2e, 0xe1, 0xf6, 0xc6, 0x82, 0xe8, 0xac, 0x8b, 0xff, 0x67, 0x5c, 0xfc, 0xbd, 0xbd,
	0x75, 0x9d, 0x8b, 0xe0, 0x85, 0x7a, 0xa2, 0x57, 0x7b, 0xbb, 0xb0, 0x36, 0x64, 0xa7, 0xd3, 0xa3,
	0x31, 0xa7, 0x33, 0x3e, 0xf6, 0x40, 0x1b, 0x19, 0xa8, 0x69, 0x3f, 0x70, 0x1e, 0xd7, 0x2c, 0x7c,
	0xe4, 0xea, 0xf9, 0xff, 0xdf, 0xd7, 0x00, 0x00, 0x00, 0xff, 0xff, 0x42, 0xc3, 0xe7, 0x1d, 0x82,
	0x06, 0x00, 0x00,
}
