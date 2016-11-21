// Code generated by protoc-gen-go.
// source: incident.proto
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

type Incident struct {
	Phid                 string               `protobuf:"bytes,1,opt,name=phid" json:"phid,omitempty"`
	KubernetesCluster    *appscode_dtypes.Uid `protobuf:"bytes,2,opt,name=kubernetes_cluster,json=kubernetesCluster" json:"kubernetes_cluster,omitempty"`
	KubernetesNamespace  string               `protobuf:"bytes,3,opt,name=kubernetes_namespace,json=kubernetesNamespace" json:"kubernetes_namespace,omitempty"`
	KubernetesObjectType string               `protobuf:"bytes,4,opt,name=kubernetes_object_type,json=kubernetesObjectType" json:"kubernetes_object_type,omitempty"`
	KubernetesObjectName string               `protobuf:"bytes,5,opt,name=kubernetes_object_name,json=kubernetesObjectName" json:"kubernetes_object_name,omitempty"`
	KubernetesAlertName  string               `protobuf:"bytes,6,opt,name=kubernetes_alert_name,json=kubernetesAlertName" json:"kubernetes_alert_name,omitempty"`
	IcingaHost           string               `protobuf:"bytes,7,opt,name=icinga_host,json=icingaHost" json:"icinga_host,omitempty"`
	IcingaService        string               `protobuf:"bytes,8,opt,name=icinga_service,json=icingaService" json:"icinga_service,omitempty"`
	Type                 string               `protobuf:"bytes,9,opt,name=type" json:"type,omitempty"`
	State                string               `protobuf:"bytes,10,opt,name=state" json:"state,omitempty"`
	User                 *appscode_dtypes.Uid `protobuf:"bytes,11,opt,name=user" json:"user,omitempty"`
	// Timestamp of first reported event
	ReportedAt int64 `protobuf:"varint,12,opt,name=reported_at,json=reportedAt" json:"reported_at,omitempty"`
	// Timestamp of first acknowledgement
	AcknowledgedAt int64             `protobuf:"varint,13,opt,name=acknowledged_at,json=acknowledgedAt" json:"acknowledged_at,omitempty"`
	RecoveredAt    int64             `protobuf:"varint,14,opt,name=recovered_at,json=recoveredAt" json:"recovered_at,omitempty"`
	IcingawebUrl   string            `protobuf:"bytes,15,opt,name=icingaweb_url,json=icingawebUrl" json:"icingaweb_url,omitempty"`
	Events         []*Incident_Event `protobuf:"bytes,16,rep,name=events" json:"events,omitempty"`
}

func (m *Incident) Reset()                    { *m = Incident{} }
func (m *Incident) String() string            { return proto.CompactTextString(m) }
func (*Incident) ProtoMessage()               {}
func (*Incident) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *Incident) GetKubernetesCluster() *appscode_dtypes.Uid {
	if m != nil {
		return m.KubernetesCluster
	}
	return nil
}

func (m *Incident) GetUser() *appscode_dtypes.Uid {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *Incident) GetEvents() []*Incident_Event {
	if m != nil {
		return m.Events
	}
	return nil
}

type Incident_Event struct {
	Type       string               `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	State      string               `protobuf:"bytes,2,opt,name=state" json:"state,omitempty"`
	ReportedAt int64                `protobuf:"varint,3,opt,name=reported_at,json=reportedAt" json:"reported_at,omitempty"`
	User       *appscode_dtypes.Uid `protobuf:"bytes,4,opt,name=user" json:"user,omitempty"`
	Comment    string               `protobuf:"bytes,5,opt,name=comment" json:"comment,omitempty"`
}

func (m *Incident_Event) Reset()                    { *m = Incident_Event{} }
func (m *Incident_Event) String() string            { return proto.CompactTextString(m) }
func (*Incident_Event) ProtoMessage()               {}
func (*Incident_Event) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0, 0} }

func (m *Incident_Event) GetUser() *appscode_dtypes.Uid {
	if m != nil {
		return m.User
	}
	return nil
}

// Next Id: 6
type IncidentListRequest struct {
	KubernetesCluster    string   `protobuf:"bytes,1,opt,name=kubernetes_cluster,json=kubernetesCluster" json:"kubernetes_cluster,omitempty"`
	KubernetesNamespace  string   `protobuf:"bytes,2,opt,name=kubernetes_namespace,json=kubernetesNamespace" json:"kubernetes_namespace,omitempty"`
	KubernetesObjectType string   `protobuf:"bytes,3,opt,name=kubernetes_object_type,json=kubernetesObjectType" json:"kubernetes_object_type,omitempty"`
	KubernetesObjectName string   `protobuf:"bytes,4,opt,name=kubernetes_object_name,json=kubernetesObjectName" json:"kubernetes_object_name,omitempty"`
	States               []string `protobuf:"bytes,5,rep,name=states" json:"states,omitempty"`
}

func (m *IncidentListRequest) Reset()                    { *m = IncidentListRequest{} }
func (m *IncidentListRequest) String() string            { return proto.CompactTextString(m) }
func (*IncidentListRequest) ProtoMessage()               {}
func (*IncidentListRequest) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

type IncidentListResponse struct {
	Status    *appscode_dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Incidents []*Incident             `protobuf:"bytes,2,rep,name=incidents" json:"incidents,omitempty"`
}

func (m *IncidentListResponse) Reset()                    { *m = IncidentListResponse{} }
func (m *IncidentListResponse) String() string            { return proto.CompactTextString(m) }
func (*IncidentListResponse) ProtoMessage()               {}
func (*IncidentListResponse) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{2} }

func (m *IncidentListResponse) GetStatus() *appscode_dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *IncidentListResponse) GetIncidents() []*Incident {
	if m != nil {
		return m.Incidents
	}
	return nil
}

// Next Id: 2
type IncidentDescribeRequest struct {
	Phid string `protobuf:"bytes,1,opt,name=phid" json:"phid,omitempty"`
}

func (m *IncidentDescribeRequest) Reset()                    { *m = IncidentDescribeRequest{} }
func (m *IncidentDescribeRequest) String() string            { return proto.CompactTextString(m) }
func (*IncidentDescribeRequest) ProtoMessage()               {}
func (*IncidentDescribeRequest) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{3} }

type IncidentDescribeResponse struct {
	Status   *appscode_dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Incident *Incident               `protobuf:"bytes,2,opt,name=incident" json:"incident,omitempty"`
}

func (m *IncidentDescribeResponse) Reset()                    { *m = IncidentDescribeResponse{} }
func (m *IncidentDescribeResponse) String() string            { return proto.CompactTextString(m) }
func (*IncidentDescribeResponse) ProtoMessage()               {}
func (*IncidentDescribeResponse) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{4} }

func (m *IncidentDescribeResponse) GetStatus() *appscode_dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *IncidentDescribeResponse) GetIncident() *Incident {
	if m != nil {
		return m.Incident
	}
	return nil
}

// Next Id: 12
type IncidentNotifyRequest struct {
	AlertPhid string `protobuf:"bytes,1,opt,name=alert_phid,json=alertPhid" json:"alert_phid,omitempty"`
	HostName  string `protobuf:"bytes,2,opt,name=host_name,json=hostName" json:"host_name,omitempty"`
	Type      string `protobuf:"bytes,3,opt,name=type" json:"type,omitempty"`
	State     string `protobuf:"bytes,4,opt,name=state" json:"state,omitempty"`
	Output    string `protobuf:"bytes,5,opt,name=output" json:"output,omitempty"`
	// The time object is used in icinga to send request. This
	// indicates detection time from icinga.
	Time                int64  `protobuf:"varint,6,opt,name=time" json:"time,omitempty"`
	Author              string `protobuf:"bytes,7,opt,name=author" json:"author,omitempty"`
	Comment             string `protobuf:"bytes,8,opt,name=comment" json:"comment,omitempty"`
	KubernetesAlertName string `protobuf:"bytes,9,opt,name=kubernetes_alert_name,json=kubernetesAlertName" json:"kubernetes_alert_name,omitempty"`
	KubernetesCluster   string `protobuf:"bytes,10,opt,name=kubernetes_cluster,json=kubernetesCluster" json:"kubernetes_cluster,omitempty"`
	KubernetesNamespace string `protobuf:"bytes,11,opt,name=kubernetes_namespace,json=kubernetesNamespace" json:"kubernetes_namespace,omitempty"`
}

func (m *IncidentNotifyRequest) Reset()                    { *m = IncidentNotifyRequest{} }
func (m *IncidentNotifyRequest) String() string            { return proto.CompactTextString(m) }
func (*IncidentNotifyRequest) ProtoMessage()               {}
func (*IncidentNotifyRequest) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{5} }

// Next Id: 4
type IncidentEventCreateRequest struct {
	// Incident PHID
	Phid        string `protobuf:"bytes,1,opt,name=phid" json:"phid,omitempty"`
	Comment     string `protobuf:"bytes,2,opt,name=comment" json:"comment,omitempty"`
	Acknowledge bool   `protobuf:"varint,3,opt,name=acknowledge" json:"acknowledge,omitempty"`
}

func (m *IncidentEventCreateRequest) Reset()                    { *m = IncidentEventCreateRequest{} }
func (m *IncidentEventCreateRequest) String() string            { return proto.CompactTextString(m) }
func (*IncidentEventCreateRequest) ProtoMessage()               {}
func (*IncidentEventCreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{6} }

func init() {
	proto.RegisterType((*Incident)(nil), "appscode.kubernetes.v1beta1.Incident")
	proto.RegisterType((*Incident_Event)(nil), "appscode.kubernetes.v1beta1.Incident.Event")
	proto.RegisterType((*IncidentListRequest)(nil), "appscode.kubernetes.v1beta1.IncidentListRequest")
	proto.RegisterType((*IncidentListResponse)(nil), "appscode.kubernetes.v1beta1.IncidentListResponse")
	proto.RegisterType((*IncidentDescribeRequest)(nil), "appscode.kubernetes.v1beta1.IncidentDescribeRequest")
	proto.RegisterType((*IncidentDescribeResponse)(nil), "appscode.kubernetes.v1beta1.IncidentDescribeResponse")
	proto.RegisterType((*IncidentNotifyRequest)(nil), "appscode.kubernetes.v1beta1.IncidentNotifyRequest")
	proto.RegisterType((*IncidentEventCreateRequest)(nil), "appscode.kubernetes.v1beta1.IncidentEventCreateRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Incidents service

type IncidentsClient interface {
	List(ctx context.Context, in *IncidentListRequest, opts ...grpc.CallOption) (*IncidentListResponse, error)
	Describe(ctx context.Context, in *IncidentDescribeRequest, opts ...grpc.CallOption) (*IncidentDescribeResponse, error)
	Notify(ctx context.Context, in *IncidentNotifyRequest, opts ...grpc.CallOption) (*appscode_dtypes.VoidResponse, error)
	CreateEvent(ctx context.Context, in *IncidentEventCreateRequest, opts ...grpc.CallOption) (*appscode_dtypes.VoidResponse, error)
}

type incidentsClient struct {
	cc *grpc.ClientConn
}

func NewIncidentsClient(cc *grpc.ClientConn) IncidentsClient {
	return &incidentsClient{cc}
}

func (c *incidentsClient) List(ctx context.Context, in *IncidentListRequest, opts ...grpc.CallOption) (*IncidentListResponse, error) {
	out := new(IncidentListResponse)
	err := grpc.Invoke(ctx, "/appscode.kubernetes.v1beta1.Incidents/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *incidentsClient) Describe(ctx context.Context, in *IncidentDescribeRequest, opts ...grpc.CallOption) (*IncidentDescribeResponse, error) {
	out := new(IncidentDescribeResponse)
	err := grpc.Invoke(ctx, "/appscode.kubernetes.v1beta1.Incidents/Describe", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *incidentsClient) Notify(ctx context.Context, in *IncidentNotifyRequest, opts ...grpc.CallOption) (*appscode_dtypes.VoidResponse, error) {
	out := new(appscode_dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/appscode.kubernetes.v1beta1.Incidents/Notify", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *incidentsClient) CreateEvent(ctx context.Context, in *IncidentEventCreateRequest, opts ...grpc.CallOption) (*appscode_dtypes.VoidResponse, error) {
	out := new(appscode_dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/appscode.kubernetes.v1beta1.Incidents/CreateEvent", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Incidents service

type IncidentsServer interface {
	List(context.Context, *IncidentListRequest) (*IncidentListResponse, error)
	Describe(context.Context, *IncidentDescribeRequest) (*IncidentDescribeResponse, error)
	Notify(context.Context, *IncidentNotifyRequest) (*appscode_dtypes.VoidResponse, error)
	CreateEvent(context.Context, *IncidentEventCreateRequest) (*appscode_dtypes.VoidResponse, error)
}

func RegisterIncidentsServer(s *grpc.Server, srv IncidentsServer) {
	s.RegisterService(&_Incidents_serviceDesc, srv)
}

func _Incidents_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IncidentListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IncidentsServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.kubernetes.v1beta1.Incidents/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IncidentsServer).List(ctx, req.(*IncidentListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Incidents_Describe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IncidentDescribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IncidentsServer).Describe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.kubernetes.v1beta1.Incidents/Describe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IncidentsServer).Describe(ctx, req.(*IncidentDescribeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Incidents_Notify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IncidentNotifyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IncidentsServer).Notify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.kubernetes.v1beta1.Incidents/Notify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IncidentsServer).Notify(ctx, req.(*IncidentNotifyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Incidents_CreateEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IncidentEventCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IncidentsServer).CreateEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.kubernetes.v1beta1.Incidents/CreateEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IncidentsServer).CreateEvent(ctx, req.(*IncidentEventCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Incidents_serviceDesc = grpc.ServiceDesc{
	ServiceName: "appscode.kubernetes.v1beta1.Incidents",
	HandlerType: (*IncidentsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _Incidents_List_Handler,
		},
		{
			MethodName: "Describe",
			Handler:    _Incidents_Describe_Handler,
		},
		{
			MethodName: "Notify",
			Handler:    _Incidents_Notify_Handler,
		},
		{
			MethodName: "CreateEvent",
			Handler:    _Incidents_CreateEvent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor3,
}

func init() { proto.RegisterFile("incident.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 968 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x56, 0x4f, 0x6f, 0x1b, 0x45,
	0x14, 0xd7, 0xac, 0xff, 0xc4, 0x7e, 0x9b, 0xa4, 0x30, 0x4d, 0xd3, 0x91, 0x0b, 0xd4, 0xdd, 0xaa,
	0xc2, 0x0a, 0xc4, 0x8b, 0xdd, 0x22, 0xaa, 0x4a, 0xa8, 0x72, 0x0c, 0x12, 0x95, 0x50, 0x89, 0xb6,
	0xb4, 0x07, 0x38, 0x58, 0xeb, 0xdd, 0xc1, 0x59, 0xea, 0xec, 0x2c, 0x3b, 0xb3, 0xa9, 0xa2, 0xaa,
	0x97, 0x9e, 0x11, 0x17, 0x24, 0xbe, 0x01, 0x77, 0x2e, 0x48, 0x9c, 0x11, 0xdf, 0x80, 0xaf, 0xc0,
	0x81, 0x2f, 0xc0, 0x89, 0x4b, 0xb5, 0xf3, 0xc7, 0x5e, 0xc7, 0x76, 0x62, 0xe7, 0x62, 0xed, 0xbc,
	0xf7, 0x7e, 0x33, 0xbf, 0xf7, 0xe6, 0xfd, 0x9e, 0x07, 0xb6, 0xa3, 0x38, 0x88, 0x42, 0x1a, 0x8b,
	0x76, 0x92, 0x32, 0xc1, 0xf0, 0x0d, 0x3f, 0x49, 0x78, 0xc0, 0x42, 0xda, 0x7e, 0x9e, 0x0d, 0x69,
	0x1a, 0x53, 0x41, 0x79, 0xfb, 0xa4, 0x33, 0xa4, 0xc2, 0xef, 0x34, 0xde, 0x19, 0x31, 0x36, 0x1a,
	0x53, 0xd7, 0x4f, 0x22, 0xd7, 0x8f, 0x63, 0x26, 0x7c, 0x11, 0xb1, 0x98, 0x2b, 0x68, 0xe3, 0x3d,
	0x03, 0x5d, 0xe2, 0xbf, 0x39, 0xe3, 0x0f, 0xc5, 0x69, 0x42, 0xb9, 0x2b, 0x7f, 0x55, 0x80, 0xf3,
	0x67, 0x15, 0x6a, 0x8f, 0x34, 0x1d, 0x8c, 0xa1, 0x9c, 0x1c, 0x45, 0x21, 0x41, 0x4d, 0xd4, 0xaa,
	0x7b, 0xf2, 0x1b, 0xf7, 0x01, 0x4f, 0x59, 0x0d, 0x82, 0x71, 0xc6, 0x05, 0x4d, 0x89, 0xd5, 0x44,
	0x2d, 0xbb, 0xbb, 0xd3, 0x9e, 0x30, 0x57, 0x5b, 0xb7, 0x9f, 0x46, 0xa1, 0xf7, 0xf6, 0x34, 0xbe,
	0xaf, 0xc2, 0x71, 0x07, 0x76, 0x0a, 0x9b, 0xc4, 0xfe, 0x31, 0xe5, 0x89, 0x1f, 0x50, 0x52, 0x92,
	0x07, 0x5d, 0x9d, 0xfa, 0x1e, 0x1b, 0x17, 0xbe, 0x07, 0xbb, 0x05, 0x08, 0x1b, 0x7e, 0x4f, 0x03,
	0x31, 0xc8, 0x0f, 0x21, 0x65, 0x09, 0x2a, 0x6c, 0xf8, 0x95, 0x74, 0x7e, 0x7d, 0x9a, 0x2c, 0x41,
	0xe5, 0xe7, 0x91, 0xca, 0x62, 0x54, 0x7e, 0x20, 0xee, 0xc2, 0xb5, 0x02, 0xca, 0x1f, 0xd3, 0x54,
	0x83, 0xaa, 0x67, 0xf9, 0xf5, 0x72, 0x9f, 0xc4, 0xdc, 0x04, 0x3b, 0x0a, 0xa2, 0x78, 0xe4, 0x0f,
	0x8e, 0x18, 0x17, 0x64, 0x43, 0x46, 0x82, 0x32, 0x7d, 0xc1, 0xb8, 0xc0, 0x77, 0x60, 0x5b, 0x07,
	0x70, 0x9a, 0x9e, 0x44, 0x01, 0x25, 0x35, 0x19, 0xb3, 0xa5, 0xac, 0x4f, 0x94, 0x31, 0xaf, 0xb9,
	0xcc, 0xaa, 0xae, 0x6a, 0x9e, 0x7f, 0xe3, 0x1d, 0xa8, 0x70, 0xe1, 0x0b, 0x4a, 0x40, 0x1a, 0xd5,
	0x02, 0xb7, 0xa0, 0x9c, 0x71, 0x9a, 0x12, 0xfb, 0x9c, 0xda, 0xcb, 0x88, 0x9c, 0x5b, 0x4a, 0x13,
	0x96, 0x0a, 0x1a, 0x0e, 0x7c, 0x41, 0x36, 0x9b, 0xa8, 0x55, 0xf2, 0xc0, 0x98, 0x7a, 0x02, 0xbf,
	0x0f, 0x57, 0xfc, 0xe0, 0x79, 0xcc, 0x5e, 0x8c, 0x69, 0x38, 0x52, 0x41, 0x5b, 0x32, 0x68, 0xbb,
	0x68, 0xee, 0x09, 0x7c, 0x0b, 0x36, 0x53, 0x1a, 0xb0, 0x13, 0x9a, 0xaa, 0xa8, 0x6d, 0x19, 0x65,
	0x4f, 0x6c, 0x3d, 0x81, 0x6f, 0x83, 0xce, 0xe8, 0x05, 0x1d, 0x0e, 0xb2, 0x74, 0x4c, 0xae, 0x48,
	0xd2, 0x9b, 0x13, 0xe3, 0xd3, 0x74, 0x8c, 0xfb, 0x50, 0xa5, 0x27, 0x34, 0x16, 0x9c, 0xbc, 0xd5,
	0x2c, 0xb5, 0xec, 0xee, 0x07, 0xed, 0x73, 0x7a, 0xbe, 0x6d, 0x1a, 0xb2, 0xfd, 0x79, 0x8e, 0xf1,
	0x34, 0xb4, 0xf1, 0x0b, 0x82, 0x8a, 0xb4, 0x4c, 0x8a, 0x86, 0x16, 0x15, 0xcd, 0x2a, 0x16, 0xed,
	0x4c, 0x29, 0x4a, 0x73, 0xa5, 0x30, 0x55, 0x2d, 0x5f, 0x58, 0x55, 0x02, 0x1b, 0x01, 0x3b, 0x3e,
	0xa6, 0xb1, 0xd0, 0xcd, 0x64, 0x96, 0xce, 0xff, 0x08, 0xae, 0x1a, 0xce, 0x5f, 0x46, 0x5c, 0x78,
	0xf4, 0x87, 0x8c, 0x72, 0x81, 0xf7, 0x17, 0x6a, 0x47, 0x91, 0x5e, 0x43, 0x25, 0xd6, 0x65, 0x54,
	0x52, 0xba, 0x94, 0x4a, 0xca, 0xe7, 0xa8, 0x64, 0x17, 0xaa, 0xb2, 0xa6, 0x9c, 0x54, 0x9a, 0xa5,
	0x56, 0xdd, 0xd3, 0x2b, 0xe7, 0x47, 0x04, 0x3b, 0xb3, 0xd9, 0xf3, 0x84, 0xc5, 0x9c, 0x62, 0x57,
	0x01, 0x32, 0x2e, 0x53, 0xb6, 0xbb, 0xd7, 0xe7, 0x8a, 0xfb, 0x44, 0xba, 0x3d, 0x1d, 0x86, 0xfb,
	0x50, 0x37, 0xa3, 0x91, 0x13, 0x4b, 0x36, 0xca, 0x9d, 0x95, 0x1a, 0xc5, 0x9b, 0xe2, 0x9c, 0x7d,
	0xb8, 0x6e, 0xcc, 0x9f, 0x51, 0x1e, 0xa4, 0xd1, 0x90, 0x9a, 0xfb, 0x58, 0x30, 0xdf, 0x9c, 0x9f,
	0x10, 0x90, 0xf9, 0xf8, 0xcb, 0x66, 0xd0, 0x83, 0x9a, 0x61, 0xa2, 0x67, 0xe4, 0x8a, 0x09, 0x4c,
	0x60, 0xce, 0x7f, 0x16, 0x5c, 0x33, 0xe6, 0xc7, 0x4c, 0x44, 0xdf, 0x9d, 0x1a, 0xfa, 0xb7, 0x00,
	0xd4, 0x6c, 0x9a, 0x26, 0x71, 0x60, 0x11, 0xe4, 0xd5, 0xa5, 0xf5, 0x30, 0x9f, 0xd6, 0x37, 0xa0,
	0x9e, 0x8f, 0x23, 0x75, 0x99, 0xaa, 0x6f, 0x6a, 0xb9, 0x41, 0x5e, 0xa0, 0x51, 0x4d, 0x69, 0x91,
	0x6a, 0xca, 0x45, 0xd5, 0xec, 0x42, 0x95, 0x65, 0x22, 0xc9, 0x4c, 0xa7, 0xeb, 0x95, 0xdc, 0x21,
	0xd2, 0x73, 0xb1, 0xe4, 0xc9, 0xef, 0x3c, 0xd6, 0xcf, 0xc4, 0x11, 0x4b, 0xf5, 0x0c, 0xd4, 0xab,
	0xa2, 0x5c, 0x6a, 0x33, 0x72, 0x59, 0x3e, 0x6e, 0xeb, 0xcb, 0xc7, 0xed, 0x62, 0x29, 0xc1, 0xba,
	0x52, 0xb2, 0x97, 0x4a, 0xc9, 0x19, 0x43, 0xc3, 0x94, 0x5d, 0x0e, 0x99, 0x7e, 0x4a, 0x7d, 0x71,
	0x5e, 0xeb, 0x14, 0x33, 0xb4, 0x66, 0x33, 0x6c, 0x82, 0x5d, 0x18, 0xa4, 0xb2, 0xe0, 0x35, 0xaf,
	0x68, 0xea, 0xfe, 0x5b, 0x81, 0xba, 0x39, 0x8e, 0xe3, 0x5f, 0x11, 0x94, 0x73, 0xe9, 0xe0, 0x8f,
	0x56, 0xea, 0x96, 0xc2, 0x8c, 0x69, 0x74, 0xd6, 0x40, 0xa8, 0xae, 0x76, 0xee, 0xbf, 0xfe, 0x9d,
	0x58, 0x35, 0xf4, 0xfa, 0xef, 0x7f, 0x7e, 0xb6, 0x3e, 0xc4, 0x7b, 0xee, 0xcc, 0x1b, 0x61, 0xba,
	0x89, 0xab, 0x37, 0x71, 0x27, 0xda, 0xc2, 0x7f, 0x20, 0xa8, 0x19, 0x91, 0xe0, 0x7b, 0x2b, 0x9d,
	0x7c, 0x46, 0x83, 0x8d, 0x8f, 0xd7, 0x44, 0x69, 0xce, 0x0f, 0x0b, 0x9c, 0xef, 0xe2, 0xce, 0xea,
	0x9c, 0xdd, 0x97, 0xf9, 0x5d, 0xbd, 0xc2, 0x7f, 0x21, 0xa8, 0x2a, 0x39, 0xe1, 0xee, 0x4a, 0x14,
	0x66, 0xb4, 0xd7, 0x78, 0x77, 0x4e, 0xf9, 0xcf, 0x58, 0x14, 0x4e, 0xe8, 0x8d, 0x0a, 0xf4, 0xbe,
	0x75, 0x9e, 0x5d, 0x48, 0x4f, 0x37, 0x30, 0x77, 0x5f, 0xce, 0x77, 0xf5, 0x2b, 0xd7, 0x0f, 0xe4,
	0x1b, 0xce, 0x8d, 0x25, 0x85, 0x7d, 0x93, 0xcb, 0x03, 0xb4, 0x87, 0x7f, 0x43, 0x60, 0xab, 0xce,
	0x54, 0xff, 0x84, 0x9f, 0xac, 0x94, 0xcb, 0x7c, 0x43, 0x5f, 0x94, 0xd0, 0xa3, 0x42, 0x42, 0x9f,
	0x3a, 0xf7, 0xd7, 0xae, 0xb7, 0xab, 0xfe, 0xb3, 0x1f, 0xa0, 0xbd, 0x83, 0x87, 0x70, 0x3b, 0x60,
	0xc7, 0xd3, 0xe3, 0xfc, 0x24, 0x5a, 0xc0, 0xf5, 0x60, 0xcb, 0x90, 0x3d, 0xcc, 0x1f, 0xa6, 0x87,
	0xe8, 0x9b, 0x0d, 0xed, 0x19, 0x56, 0xe5, 0x53, 0xf5, 0xee, 0x9b, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x1e, 0x60, 0xef, 0xc2, 0x38, 0x0b, 0x00, 0x00,
}