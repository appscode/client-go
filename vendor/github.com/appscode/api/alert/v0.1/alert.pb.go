// Code generated by protoc-gen-go.
// source: alert.proto
// DO NOT EDIT!

/*
Package alert is a generated protocol buffer package.

It is generated from these files:
	alert.proto
	incident.proto

It has these top-level messages:
	IcingaParam
	IcingaState
	NotifierParam
	Alert
	AlertListRequest
	AlertListResponse
	AlertCreateRequest
	AlertDescribeRequest
	AlertDescribeResponse
	AlertUpdateRequest
	AlertSyncRequest
	AlertDeleteRequest
	Incident
	IncidentListRequest
	IncidentListResponse
	IncidentDescribeRequest
	IncidentDescribeResponse
	IncidentNotifyRequest
	IncidentEventCreateRequest
*/
package alert

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

type IcingaParam struct {
	CheckIntervalSec int64 `protobuf:"varint,1,opt,name=check_interval_sec,json=checkIntervalSec" json:"check_interval_sec,omitempty"`
	AlertIntervalSec int64 `protobuf:"varint,2,opt,name=alert_interval_sec,json=alertIntervalSec" json:"alert_interval_sec,omitempty"`
}

func (m *IcingaParam) Reset()                    { *m = IcingaParam{} }
func (m *IcingaParam) String() string            { return proto.CompactTextString(m) }
func (*IcingaParam) ProtoMessage()               {}
func (*IcingaParam) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// This one is used in kubernetes/v0.1/client.proto
// To send information if alert is set for app/pod
type IcingaState struct {
	OK       int32 `protobuf:"varint,1,opt,name=OK,json=oK" json:"OK,omitempty"`
	Warning  int32 `protobuf:"varint,2,opt,name=Warning,json=warning" json:"Warning,omitempty"`
	Critical int32 `protobuf:"varint,3,opt,name=Critical,json=critical" json:"Critical,omitempty"`
	Unknown  int32 `protobuf:"varint,4,opt,name=Unknown,json=unknown" json:"Unknown,omitempty"`
}

func (m *IcingaState) Reset()                    { *m = IcingaState{} }
func (m *IcingaState) String() string            { return proto.CompactTextString(m) }
func (*IcingaState) ProtoMessage()               {}
func (*IcingaState) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type NotifierParam struct {
	State   string `protobuf:"bytes,1,opt,name=state" json:"state,omitempty"`
	UserUid string `protobuf:"bytes,2,opt,name=user_uid,json=userUid" json:"user_uid,omitempty"`
	Method  string `protobuf:"bytes,3,opt,name=method" json:"method,omitempty"`
}

func (m *NotifierParam) Reset()                    { *m = NotifierParam{} }
func (m *NotifierParam) String() string            { return proto.CompactTextString(m) }
func (*NotifierParam) ProtoMessage()               {}
func (*NotifierParam) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type Alert struct {
	Phid                 string            `protobuf:"bytes,1,opt,name=phid" json:"phid,omitempty"`
	KubernetesCluster    string            `protobuf:"bytes,2,opt,name=kubernetes_cluster,json=kubernetesCluster" json:"kubernetes_cluster,omitempty"`
	KubernetesNamespace  string            `protobuf:"bytes,3,opt,name=kubernetes_namespace,json=kubernetesNamespace" json:"kubernetes_namespace,omitempty"`
	KubernetesObjectType string            `protobuf:"bytes,4,opt,name=kubernetes_objectType,json=kubernetesObjectType" json:"kubernetes_objectType,omitempty"`
	KubernetesObjectName string            `protobuf:"bytes,5,opt,name=kubernetes_objectName,json=kubernetesObjectName" json:"kubernetes_objectName,omitempty"`
	IcingaService        string            `protobuf:"bytes,7,opt,name=icinga_service,json=icingaService" json:"icinga_service,omitempty"`
	IcingaParam          *IcingaParam      `protobuf:"bytes,8,opt,name=icinga_param,json=icingaParam" json:"icinga_param,omitempty"`
	CheckCommand         string            `protobuf:"bytes,9,opt,name=check_command,json=checkCommand" json:"check_command,omitempty"`
	NotifierParam        []*NotifierParam  `protobuf:"bytes,11,rep,name=notifier_param,json=notifierParam" json:"notifier_param,omitempty"`
	IcingawebUrl         string            `protobuf:"bytes,12,opt,name=icingaweb_url,json=icingawebUrl" json:"icingaweb_url,omitempty"`
	Vars                 map[string]string `protobuf:"bytes,13,rep,name=vars" json:"vars,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Alert) Reset()                    { *m = Alert{} }
func (m *Alert) String() string            { return proto.CompactTextString(m) }
func (*Alert) ProtoMessage()               {}
func (*Alert) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Alert) GetIcingaParam() *IcingaParam {
	if m != nil {
		return m.IcingaParam
	}
	return nil
}

func (m *Alert) GetNotifierParam() []*NotifierParam {
	if m != nil {
		return m.NotifierParam
	}
	return nil
}

func (m *Alert) GetVars() map[string]string {
	if m != nil {
		return m.Vars
	}
	return nil
}

// Next Id: 5
type AlertListRequest struct {
	KubernetesCluster    string `protobuf:"bytes,1,opt,name=kubernetes_cluster,json=kubernetesCluster" json:"kubernetes_cluster,omitempty"`
	KubernetesNamespace  string `protobuf:"bytes,2,opt,name=kubernetes_namespace,json=kubernetesNamespace" json:"kubernetes_namespace,omitempty"`
	KubernetesObjectType string `protobuf:"bytes,3,opt,name=kubernetes_objectType,json=kubernetesObjectType" json:"kubernetes_objectType,omitempty"`
	KubernetesObjectName string `protobuf:"bytes,4,opt,name=kubernetes_objectName,json=kubernetesObjectName" json:"kubernetes_objectName,omitempty"`
}

func (m *AlertListRequest) Reset()                    { *m = AlertListRequest{} }
func (m *AlertListRequest) String() string            { return proto.CompactTextString(m) }
func (*AlertListRequest) ProtoMessage()               {}
func (*AlertListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type AlertListResponse struct {
	Status *dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Alerts []*Alert       `protobuf:"bytes,2,rep,name=alerts" json:"alerts,omitempty"`
}

func (m *AlertListResponse) Reset()                    { *m = AlertListResponse{} }
func (m *AlertListResponse) String() string            { return proto.CompactTextString(m) }
func (*AlertListResponse) ProtoMessage()               {}
func (*AlertListResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *AlertListResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *AlertListResponse) GetAlerts() []*Alert {
	if m != nil {
		return m.Alerts
	}
	return nil
}

// Next Id: 12
type AlertCreateRequest struct {
	KubernetesCluster    string            `protobuf:"bytes,1,opt,name=kubernetes_cluster,json=kubernetesCluster" json:"kubernetes_cluster,omitempty"`
	KubernetesNamespace  string            `protobuf:"bytes,2,opt,name=kubernetes_namespace,json=kubernetesNamespace" json:"kubernetes_namespace,omitempty"`
	KubernetesObjectType string            `protobuf:"bytes,3,opt,name=kubernetes_objectType,json=kubernetesObjectType" json:"kubernetes_objectType,omitempty"`
	KubernetesObjectName string            `protobuf:"bytes,4,opt,name=kubernetes_objectName,json=kubernetesObjectName" json:"kubernetes_objectName,omitempty"`
	IcingaService        string            `protobuf:"bytes,6,opt,name=icinga_service,json=icingaService" json:"icinga_service,omitempty"`
	IcingaParam          *IcingaParam      `protobuf:"bytes,7,opt,name=icinga_param,json=icingaParam" json:"icinga_param,omitempty"`
	CheckCommand         string            `protobuf:"bytes,8,opt,name=check_command,json=checkCommand" json:"check_command,omitempty"`
	NotifierParam        []*NotifierParam  `protobuf:"bytes,10,rep,name=notifier_param,json=notifierParam" json:"notifier_param,omitempty"`
	Vars                 map[string]string `protobuf:"bytes,11,rep,name=vars" json:"vars,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *AlertCreateRequest) Reset()                    { *m = AlertCreateRequest{} }
func (m *AlertCreateRequest) String() string            { return proto.CompactTextString(m) }
func (*AlertCreateRequest) ProtoMessage()               {}
func (*AlertCreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *AlertCreateRequest) GetIcingaParam() *IcingaParam {
	if m != nil {
		return m.IcingaParam
	}
	return nil
}

func (m *AlertCreateRequest) GetNotifierParam() []*NotifierParam {
	if m != nil {
		return m.NotifierParam
	}
	return nil
}

func (m *AlertCreateRequest) GetVars() map[string]string {
	if m != nil {
		return m.Vars
	}
	return nil
}

// Next Id: 2
type AlertDescribeRequest struct {
	Phid string `protobuf:"bytes,1,opt,name=phid" json:"phid,omitempty"`
}

func (m *AlertDescribeRequest) Reset()                    { *m = AlertDescribeRequest{} }
func (m *AlertDescribeRequest) String() string            { return proto.CompactTextString(m) }
func (*AlertDescribeRequest) ProtoMessage()               {}
func (*AlertDescribeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

type AlertDescribeResponse struct {
	Status *dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Alerts *Alert         `protobuf:"bytes,2,opt,name=alerts" json:"alerts,omitempty"`
}

func (m *AlertDescribeResponse) Reset()                    { *m = AlertDescribeResponse{} }
func (m *AlertDescribeResponse) String() string            { return proto.CompactTextString(m) }
func (*AlertDescribeResponse) ProtoMessage()               {}
func (*AlertDescribeResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *AlertDescribeResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *AlertDescribeResponse) GetAlerts() *Alert {
	if m != nil {
		return m.Alerts
	}
	return nil
}

// Next Id: 6
type AlertUpdateRequest struct {
	Phid          string            `protobuf:"bytes,1,opt,name=phid" json:"phid,omitempty"`
	IcingaParam   *IcingaParam      `protobuf:"bytes,2,opt,name=icinga_param,json=icingaParam" json:"icinga_param,omitempty"`
	NotifierParam []*NotifierParam  `protobuf:"bytes,4,rep,name=notifier_param,json=notifierParam" json:"notifier_param,omitempty"`
	Vars          map[string]string `protobuf:"bytes,5,rep,name=vars" json:"vars,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *AlertUpdateRequest) Reset()                    { *m = AlertUpdateRequest{} }
func (m *AlertUpdateRequest) String() string            { return proto.CompactTextString(m) }
func (*AlertUpdateRequest) ProtoMessage()               {}
func (*AlertUpdateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *AlertUpdateRequest) GetIcingaParam() *IcingaParam {
	if m != nil {
		return m.IcingaParam
	}
	return nil
}

func (m *AlertUpdateRequest) GetNotifierParam() []*NotifierParam {
	if m != nil {
		return m.NotifierParam
	}
	return nil
}

func (m *AlertUpdateRequest) GetVars() map[string]string {
	if m != nil {
		return m.Vars
	}
	return nil
}

// Next Id: 6
type AlertSyncRequest struct {
	KubernetesCluster    string                          `protobuf:"bytes,1,opt,name=kubernetes_cluster,json=kubernetesCluster" json:"kubernetes_cluster,omitempty"`
	KubernetesNamespace  string                          `protobuf:"bytes,2,opt,name=kubernetes_namespace,json=kubernetesNamespace" json:"kubernetes_namespace,omitempty"`
	KubernetesObjectType string                          `protobuf:"bytes,3,opt,name=kubernetes_objectType,json=kubernetesObjectType" json:"kubernetes_objectType,omitempty"`
	KubernetesObjectName string                          `protobuf:"bytes,4,opt,name=kubernetes_objectName,json=kubernetesObjectName" json:"kubernetes_objectName,omitempty"`
	PodAncestor          []*AlertSyncRequest_PodAncestor `protobuf:"bytes,5,rep,name=pod_ancestor,json=podAncestor" json:"pod_ancestor,omitempty"`
}

func (m *AlertSyncRequest) Reset()                    { *m = AlertSyncRequest{} }
func (m *AlertSyncRequest) String() string            { return proto.CompactTextString(m) }
func (*AlertSyncRequest) ProtoMessage()               {}
func (*AlertSyncRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *AlertSyncRequest) GetPodAncestor() []*AlertSyncRequest_PodAncestor {
	if m != nil {
		return m.PodAncestor
	}
	return nil
}

// Next Id: 3
type AlertSyncRequest_PodAncestor struct {
	Type string `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *AlertSyncRequest_PodAncestor) Reset()         { *m = AlertSyncRequest_PodAncestor{} }
func (m *AlertSyncRequest_PodAncestor) String() string { return proto.CompactTextString(m) }
func (*AlertSyncRequest_PodAncestor) ProtoMessage()    {}
func (*AlertSyncRequest_PodAncestor) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{10, 0}
}

// Next Id: 2
type AlertDeleteRequest struct {
	Phid string `protobuf:"bytes,1,opt,name=phid" json:"phid,omitempty"`
}

func (m *AlertDeleteRequest) Reset()                    { *m = AlertDeleteRequest{} }
func (m *AlertDeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*AlertDeleteRequest) ProtoMessage()               {}
func (*AlertDeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func init() {
	proto.RegisterType((*IcingaParam)(nil), "alert.IcingaParam")
	proto.RegisterType((*IcingaState)(nil), "alert.IcingaState")
	proto.RegisterType((*NotifierParam)(nil), "alert.NotifierParam")
	proto.RegisterType((*Alert)(nil), "alert.Alert")
	proto.RegisterType((*AlertListRequest)(nil), "alert.AlertListRequest")
	proto.RegisterType((*AlertListResponse)(nil), "alert.AlertListResponse")
	proto.RegisterType((*AlertCreateRequest)(nil), "alert.AlertCreateRequest")
	proto.RegisterType((*AlertDescribeRequest)(nil), "alert.AlertDescribeRequest")
	proto.RegisterType((*AlertDescribeResponse)(nil), "alert.AlertDescribeResponse")
	proto.RegisterType((*AlertUpdateRequest)(nil), "alert.AlertUpdateRequest")
	proto.RegisterType((*AlertSyncRequest)(nil), "alert.AlertSyncRequest")
	proto.RegisterType((*AlertSyncRequest_PodAncestor)(nil), "alert.AlertSyncRequest.PodAncestor")
	proto.RegisterType((*AlertDeleteRequest)(nil), "alert.AlertDeleteRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Alerts service

type AlertsClient interface {
	List(ctx context.Context, in *AlertListRequest, opts ...grpc.CallOption) (*AlertListResponse, error)
	Create(ctx context.Context, in *AlertCreateRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
	Describe(ctx context.Context, in *AlertDescribeRequest, opts ...grpc.CallOption) (*AlertDescribeResponse, error)
	Update(ctx context.Context, in *AlertUpdateRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
	Sync(ctx context.Context, in *AlertSyncRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
	Delete(ctx context.Context, in *AlertDeleteRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
}

type alertsClient struct {
	cc *grpc.ClientConn
}

func NewAlertsClient(cc *grpc.ClientConn) AlertsClient {
	return &alertsClient{cc}
}

func (c *alertsClient) List(ctx context.Context, in *AlertListRequest, opts ...grpc.CallOption) (*AlertListResponse, error) {
	out := new(AlertListResponse)
	err := grpc.Invoke(ctx, "/alert.Alerts/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alertsClient) Create(ctx context.Context, in *AlertCreateRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/alert.Alerts/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alertsClient) Describe(ctx context.Context, in *AlertDescribeRequest, opts ...grpc.CallOption) (*AlertDescribeResponse, error) {
	out := new(AlertDescribeResponse)
	err := grpc.Invoke(ctx, "/alert.Alerts/Describe", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alertsClient) Update(ctx context.Context, in *AlertUpdateRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/alert.Alerts/Update", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alertsClient) Sync(ctx context.Context, in *AlertSyncRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/alert.Alerts/Sync", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alertsClient) Delete(ctx context.Context, in *AlertDeleteRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/alert.Alerts/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Alerts service

type AlertsServer interface {
	List(context.Context, *AlertListRequest) (*AlertListResponse, error)
	Create(context.Context, *AlertCreateRequest) (*dtypes.VoidResponse, error)
	Describe(context.Context, *AlertDescribeRequest) (*AlertDescribeResponse, error)
	Update(context.Context, *AlertUpdateRequest) (*dtypes.VoidResponse, error)
	Sync(context.Context, *AlertSyncRequest) (*dtypes.VoidResponse, error)
	Delete(context.Context, *AlertDeleteRequest) (*dtypes.VoidResponse, error)
}

func RegisterAlertsServer(s *grpc.Server, srv AlertsServer) {
	s.RegisterService(&_Alerts_serviceDesc, srv)
}

func _Alerts_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AlertListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertsServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/alert.Alerts/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertsServer).List(ctx, req.(*AlertListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Alerts_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AlertCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertsServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/alert.Alerts/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertsServer).Create(ctx, req.(*AlertCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Alerts_Describe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AlertDescribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertsServer).Describe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/alert.Alerts/Describe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertsServer).Describe(ctx, req.(*AlertDescribeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Alerts_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AlertUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertsServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/alert.Alerts/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertsServer).Update(ctx, req.(*AlertUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Alerts_Sync_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AlertSyncRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertsServer).Sync(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/alert.Alerts/Sync",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertsServer).Sync(ctx, req.(*AlertSyncRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Alerts_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AlertDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertsServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/alert.Alerts/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertsServer).Delete(ctx, req.(*AlertDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Alerts_serviceDesc = grpc.ServiceDesc{
	ServiceName: "alert.Alerts",
	HandlerType: (*AlertsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _Alerts_List_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _Alerts_Create_Handler,
		},
		{
			MethodName: "Describe",
			Handler:    _Alerts_Describe_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Alerts_Update_Handler,
		},
		{
			MethodName: "Sync",
			Handler:    _Alerts_Sync_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Alerts_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("alert.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 973 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xdc, 0x57, 0x4f, 0x6f, 0x1b, 0x45,
	0x14, 0x97, 0xd7, 0xff, 0x9f, 0xed, 0xa8, 0x1d, 0xdc, 0xb0, 0x35, 0x95, 0x08, 0x9b, 0x50, 0x8c,
	0x01, 0x9b, 0xa4, 0xaa, 0x8a, 0xca, 0xa9, 0x72, 0x41, 0xaa, 0x8a, 0xd2, 0x6a, 0x4d, 0x0a, 0x9c,
	0xac, 0xf1, 0xee, 0x34, 0x19, 0xb2, 0x9e, 0x5d, 0x66, 0x66, 0x1d, 0x59, 0x88, 0x0b, 0x57, 0x24,
	0x0e, 0xf0, 0x49, 0xf8, 0x2c, 0x9c, 0xe1, 0xc4, 0x07, 0x41, 0xf3, 0xc7, 0xce, 0x6e, 0x62, 0x13,
	0x37, 0xf4, 0xd4, 0x4b, 0x35, 0xef, 0xcf, 0xbc, 0xdf, 0xbe, 0xf7, 0xfb, 0xbd, 0x71, 0x03, 0x0d,
	0x1c, 0x11, 0x2e, 0xfb, 0x09, 0x8f, 0x65, 0x8c, 0xca, 0xda, 0xe8, 0xdc, 0x39, 0x8e, 0xe3, 0xe3,
	0x88, 0x0c, 0x70, 0x42, 0x07, 0x98, 0xb1, 0x58, 0x62, 0x49, 0x63, 0x26, 0x4c, 0x52, 0x67, 0x5b,
	0xb9, 0x43, 0x39, 0x4f, 0x88, 0x18, 0xe8, 0x7f, 0x8d, 0xdf, 0xa3, 0xd0, 0x78, 0x12, 0x50, 0x76,
	0x8c, 0x9f, 0x63, 0x8e, 0xa7, 0xe8, 0x63, 0x40, 0xc1, 0x09, 0x09, 0x4e, 0xc7, 0x94, 0x49, 0xc2,
	0x67, 0x38, 0x1a, 0x0b, 0x12, 0xb8, 0x85, 0x9d, 0x42, 0xb7, 0xe8, 0xdf, 0xd0, 0x91, 0x27, 0x36,
	0x30, 0x22, 0x81, 0xca, 0xd6, 0xd8, 0xf9, 0x6c, 0xc7, 0x64, 0xeb, 0x48, 0x26, 0xdb, 0x9b, 0x2e,
	0xa0, 0x46, 0x12, 0x4b, 0x82, 0xb6, 0xc0, 0x79, 0xf6, 0x54, 0x97, 0x2e, 0xfb, 0x4e, 0xfc, 0x14,
	0xb9, 0x50, 0xfd, 0x06, 0x73, 0x46, 0xd9, 0xb1, 0xae, 0x50, 0xf6, 0xab, 0x67, 0xc6, 0x44, 0x1d,
	0xa8, 0x0d, 0x39, 0x95, 0x34, 0xc0, 0x91, 0x5b, 0xd4, 0xa1, 0x5a, 0x60, 0x6d, 0x75, 0xeb, 0x88,
	0x9d, 0xb2, 0xf8, 0x8c, 0xb9, 0x25, 0x73, 0x2b, 0x35, 0xa6, 0xf7, 0x2d, 0xb4, 0x0e, 0x63, 0x49,
	0x5f, 0x52, 0xc2, 0x4d, 0x6f, 0x6d, 0x28, 0x0b, 0x85, 0xac, 0x31, 0xeb, 0xbe, 0x31, 0xd0, 0x6d,
	0xa8, 0xa5, 0x82, 0xf0, 0x71, 0x4a, 0x43, 0x8d, 0x5b, 0xf7, 0xab, 0xca, 0x3e, 0xa2, 0x21, 0xda,
	0x86, 0xca, 0x94, 0xc8, 0x93, 0x38, 0xd4, 0xa8, 0x75, 0xdf, 0x5a, 0xde, 0x1f, 0x25, 0x28, 0x3f,
	0x52, 0xdd, 0x21, 0x04, 0xa5, 0xe4, 0x84, 0x86, 0xb6, 0xa2, 0x3e, 0xa3, 0x4f, 0x00, 0x9d, 0xa6,
	0x13, 0xc2, 0x19, 0x91, 0x44, 0x8c, 0x83, 0x28, 0x15, 0x92, 0x70, 0x5b, 0xfa, 0xe6, 0x79, 0x64,
	0x68, 0x02, 0x68, 0x1f, 0xda, 0x99, 0x74, 0x86, 0xa7, 0x44, 0x24, 0x38, 0x20, 0x16, 0xf2, 0xad,
	0xf3, 0xd8, 0xe1, 0x22, 0x84, 0xee, 0xc1, 0xad, 0xcc, 0x95, 0x78, 0xf2, 0x3d, 0x09, 0xe4, 0xd7,
	0xf3, 0x84, 0xe8, 0x09, 0xd4, 0xfd, 0x4c, 0xbd, 0x67, 0xcb, 0xd8, 0xca, 0x4b, 0xaa, 0xa4, 0x5b,
	0x5e, 0x7d, 0x49, 0xc5, 0xd0, 0xfb, 0xb0, 0x45, 0x35, 0x65, 0x63, 0x41, 0xf8, 0x8c, 0x06, 0xc4,
	0xad, 0xea, 0xec, 0x96, 0xf1, 0x8e, 0x8c, 0x13, 0xdd, 0x87, 0xa6, 0x4d, 0x4b, 0xd4, 0xa4, 0xdd,
	0xda, 0x4e, 0xa1, 0xdb, 0x38, 0x40, 0x7d, 0xa3, 0xd2, 0x8c, 0xbe, 0xfc, 0x06, 0xcd, 0x88, 0x6d,
	0x17, 0x5a, 0x46, 0x6c, 0x41, 0x3c, 0x9d, 0x62, 0x16, 0xba, 0x75, 0x5d, 0xbc, 0xa9, 0x9d, 0x43,
	0xe3, 0x43, 0x9f, 0xc3, 0x16, 0xb3, 0x34, 0xda, 0xea, 0x8d, 0x9d, 0x62, 0xb7, 0x71, 0xd0, 0xb6,
	0xd5, 0x73, 0x1c, 0xfb, 0x2d, 0x96, 0xa3, 0x7c, 0x17, 0xec, 0x97, 0x9e, 0x91, 0xc9, 0x38, 0xe5,
	0x91, 0xdb, 0x34, 0x08, 0x4b, 0xe7, 0x11, 0x8f, 0x50, 0x0f, 0x4a, 0x33, 0xcc, 0x85, 0xdb, 0xd2,
	0x75, 0xb7, 0x6d, 0x5d, 0x4d, 0x70, 0xff, 0x05, 0xe6, 0xe2, 0x0b, 0x26, 0xf9, 0xdc, 0xd7, 0x39,
	0x9d, 0x07, 0x50, 0x5f, 0xba, 0xd0, 0x0d, 0x28, 0x9e, 0x92, 0xb9, 0x25, 0x5f, 0x1d, 0x95, 0xc4,
	0x66, 0x38, 0x4a, 0x89, 0xa5, 0xdb, 0x18, 0x0f, 0x9d, 0xcf, 0x0a, 0xde, 0xdf, 0x05, 0xb8, 0xa1,
	0x4b, 0x7e, 0x45, 0x85, 0xf4, 0xc9, 0x0f, 0x29, 0x11, 0x72, 0x8d, 0x54, 0x0a, 0xaf, 0x2a, 0x15,
	0xe7, 0x1a, 0x52, 0x29, 0x5e, 0x47, 0x2a, 0xa5, 0xf5, 0x52, 0xf1, 0x30, 0xdc, 0xcc, 0xf4, 0x27,
	0x92, 0x98, 0x09, 0x82, 0xee, 0x42, 0x45, 0x6d, 0x59, 0x2a, 0x74, 0x53, 0x8d, 0x83, 0xad, 0xbe,
	0x79, 0x82, 0xfa, 0x23, 0xed, 0xf5, 0x6d, 0x14, 0xed, 0x41, 0x45, 0x4f, 0x5d, 0xb8, 0x8e, 0x26,
	0xa1, 0x99, 0x25, 0xc1, 0xb7, 0x31, 0xef, 0xb7, 0x12, 0x20, 0xed, 0x19, 0x72, 0x82, 0x25, 0x79,
	0x03, 0xa7, 0xb8, 0x62, 0xe1, 0x2a, 0x9b, 0x2c, 0x5c, 0xf5, 0x9a, 0x0b, 0x57, 0xdb, 0x68, 0xe1,
	0x60, 0xf3, 0x85, 0x7b, 0x60, 0x77, 0xc9, 0xec, 0xe8, 0x6e, 0x96, 0xc6, 0x1c, 0x69, 0xaf, 0x6f,
	0xb1, 0x7a, 0xd0, 0xd6, 0xe5, 0x1f, 0x13, 0x11, 0x70, 0x3a, 0x59, 0xaa, 0x62, 0xc5, 0xd3, 0xec,
	0x11, 0xb8, 0x75, 0x21, 0xf7, 0x7f, 0xe8, 0xb4, 0xb0, 0x56, 0xa7, 0xbf, 0x3a, 0x56, 0xa7, 0x47,
	0x49, 0x98, 0xd1, 0xe9, 0xaa, 0x1f, 0x8b, 0x8b, 0x44, 0x3a, 0x9b, 0x11, 0x79, 0x99, 0xa3, 0xd2,
	0xab, 0x73, 0x54, 0xbe, 0xcc, 0x51, 0xee, 0x83, 0x5f, 0x1f, 0x47, 0x7f, 0x39, 0xf6, 0xf1, 0x1b,
	0xcd, 0x59, 0xf0, 0x26, 0xae, 0xed, 0x97, 0xd0, 0x4c, 0xe2, 0x70, 0x8c, 0x59, 0x40, 0x84, 0x8c,
	0xf9, 0xaa, 0xd1, 0x66, 0x5a, 0xef, 0x3f, 0x8f, 0xc3, 0x47, 0x36, 0xd5, 0x6f, 0x24, 0xe7, 0x46,
	0xe7, 0x3e, 0x34, 0x32, 0x31, 0xa5, 0x18, 0x25, 0xc3, 0x85, 0x62, 0xd4, 0x59, 0xf9, 0x54, 0xf3,
	0xb6, 0x6f, 0x7d, 0xf6, 0xba, 0x56, 0x6f, 0x8f, 0x49, 0x44, 0xfe, 0x53, 0x6f, 0x07, 0xbf, 0x94,
	0xa1, 0xa2, 0x53, 0x05, 0x9a, 0x40, 0x49, 0xbd, 0xd5, 0xe8, 0xed, 0xec, 0x57, 0x66, 0x7e, 0x9d,
	0x3a, 0xee, 0xe5, 0x80, 0x59, 0x17, 0xef, 0x83, 0x9f, 0xff, 0xfc, 0xe7, 0x77, 0xe7, 0x3d, 0xf4,
	0xee, 0x00, 0x27, 0x89, 0x08, 0xe2, 0xd0, 0xfe, 0xaf, 0x53, 0x25, 0x0e, 0x66, 0x9f, 0xf6, 0xf7,
	0xcd, 0x51, 0x20, 0x02, 0x15, 0xb3, 0xf6, 0xe8, 0xf6, 0xda, 0xa7, 0xa0, 0xd3, 0x5e, 0x2c, 0xdb,
	0x8b, 0x98, 0x86, 0x4b, 0x8c, 0x9e, 0xc6, 0xd8, 0xeb, 0x5c, 0x85, 0xf1, 0xb0, 0xd0, 0x43, 0x29,
	0xd4, 0x16, 0x2b, 0x8d, 0xde, 0xc9, 0x02, 0x5d, 0x78, 0x14, 0x3a, 0x77, 0x56, 0x07, 0x2d, 0x64,
	0x5f, 0x43, 0x76, 0xd1, 0xdd, 0x2b, 0x20, 0x07, 0x3f, 0xaa, 0x59, 0xfe, 0x84, 0x22, 0xa8, 0x98,
	0x85, 0xc9, 0x77, 0x97, 0x5b, 0xa2, 0x35, 0xdd, 0xed, 0x6b, 0xa8, 0x8f, 0xbc, 0x0d, 0xa1, 0x54,
	0x93, 0x2f, 0xa1, 0xa4, 0x34, 0x94, 0xe7, 0x2b, 0xa3, 0xaa, 0x35, 0x48, 0x03, 0x8d, 0xf4, 0xa1,
	0xb7, 0x77, 0xd5, 0x1c, 0x47, 0xdf, 0x1d, 0x0e, 0x15, 0x0e, 0x85, 0x8a, 0xd1, 0x51, 0xbe, 0xab,
	0x9c, 0xb6, 0xd6, 0x60, 0xd9, 0x01, 0xf6, 0x36, 0xec, 0x6a, 0x52, 0xd1, 0x7f, 0x83, 0xdc, 0xfb,
	0x37, 0x00, 0x00, 0xff, 0xff, 0xcb, 0x8c, 0xc8, 0xf1, 0xcf, 0x0c, 0x00, 0x00,
}
