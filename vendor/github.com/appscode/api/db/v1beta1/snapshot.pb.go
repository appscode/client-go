// Code generated by protoc-gen-go.
// source: snapshot.proto
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

// Next Id: 4
type DatabaseInfo struct {
	Cluster string `protobuf:"bytes,1,opt,name=cluster" json:"cluster,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Type    string `protobuf:"bytes,3,opt,name=type" json:"type,omitempty"`
}

func (m *DatabaseInfo) Reset()                    { *m = DatabaseInfo{} }
func (m *DatabaseInfo) String() string            { return proto.CompactTextString(m) }
func (*DatabaseInfo) ProtoMessage()               {}
func (*DatabaseInfo) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

// Next Id: 15
type Snapshot struct {
	Name              string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Phid              string `protobuf:"bytes,2,opt,name=phid" json:"phid,omitempty"`
	Provider          string `protobuf:"bytes,3,opt,name=provider" json:"provider,omitempty"`
	CloudCredential   string `protobuf:"bytes,4,opt,name=cloud_credential,json=cloudCredential" json:"cloud_credential,omitempty"`
	Bucket            string `protobuf:"bytes,5,opt,name=bucket" json:"bucket,omitempty"`
	Status            string `protobuf:"bytes,8,opt,name=status" json:"status,omitempty"`
	Process           string `protobuf:"bytes,9,opt,name=process" json:"process,omitempty"`
	Type              string `protobuf:"bytes,10,opt,name=type" json:"type,omitempty"`
	IsScheduledBackup int32  `protobuf:"varint,11,opt,name=is_scheduled_backup,json=isScheduledBackup" json:"is_scheduled_backup,omitempty"`
	SourceDatabase    string `protobuf:"bytes,12,opt,name=source_database,json=sourceDatabase" json:"source_database,omitempty"`
	SourceSnapshot    string `protobuf:"bytes,13,opt,name=source_snapshot,json=sourceSnapshot" json:"source_snapshot,omitempty"`
	CreatedAt         int64  `protobuf:"varint,14,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
}

func (m *Snapshot) Reset()                    { *m = Snapshot{} }
func (m *Snapshot) String() string            { return proto.CompactTextString(m) }
func (*Snapshot) ProtoMessage()               {}
func (*Snapshot) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

// Next Id: 3
type SnapshotListRequest struct {
	Cluster string `protobuf:"bytes,1,opt,name=cluster" json:"cluster,omitempty"`
	Uid     string `protobuf:"bytes,2,opt,name=uid" json:"uid,omitempty"`
}

func (m *SnapshotListRequest) Reset()                    { *m = SnapshotListRequest{} }
func (m *SnapshotListRequest) String() string            { return proto.CompactTextString(m) }
func (*SnapshotListRequest) ProtoMessage()               {}
func (*SnapshotListRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

// Next Id: 4
type SnapshotListResponse struct {
	Status       *dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	DatabaseInfo *DatabaseInfo  `protobuf:"bytes,2,opt,name=database_info,json=databaseInfo" json:"database_info,omitempty"`
	List         []*Snapshot    `protobuf:"bytes,3,rep,name=list" json:"list,omitempty"`
}

func (m *SnapshotListResponse) Reset()                    { *m = SnapshotListResponse{} }
func (m *SnapshotListResponse) String() string            { return proto.CompactTextString(m) }
func (*SnapshotListResponse) ProtoMessage()               {}
func (*SnapshotListResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *SnapshotListResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *SnapshotListResponse) GetDatabaseInfo() *DatabaseInfo {
	if m != nil {
		return m.DatabaseInfo
	}
	return nil
}

func (m *SnapshotListResponse) GetList() []*Snapshot {
	if m != nil {
		return m.List
	}
	return nil
}

// Next Id: 13
type BackupScheduleRequest struct {
	Cluster          string `protobuf:"bytes,1,opt,name=cluster" json:"cluster,omitempty"`
	Uid              string `protobuf:"bytes,2,opt,name=uid" json:"uid,omitempty"`
	CredentialUid    string `protobuf:"bytes,4,opt,name=credential_uid,json=credentialUid" json:"credential_uid,omitempty"`
	BucketName       string `protobuf:"bytes,5,opt,name=bucket_name,json=bucketName" json:"bucket_name,omitempty"`
	SnapshotName     string `protobuf:"bytes,6,opt,name=snapshot_name,json=snapshotName" json:"snapshot_name,omitempty"`
	AuthSecretName   string `protobuf:"bytes,8,opt,name=auth_secret_name,json=authSecretName" json:"auth_secret_name,omitempty"`
	ScheduleCronExpr string `protobuf:"bytes,9,opt,name=schedule_cron_expr,json=scheduleCronExpr" json:"schedule_cron_expr,omitempty"`
}

func (m *BackupScheduleRequest) Reset()                    { *m = BackupScheduleRequest{} }
func (m *BackupScheduleRequest) String() string            { return proto.CompactTextString(m) }
func (*BackupScheduleRequest) ProtoMessage()               {}
func (*BackupScheduleRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

// Next Id: 3
type BackupUnscheduleRequest struct {
	Cluster string `protobuf:"bytes,1,opt,name=cluster" json:"cluster,omitempty"`
	Uid     string `protobuf:"bytes,2,opt,name=uid" json:"uid,omitempty"`
}

func (m *BackupUnscheduleRequest) Reset()                    { *m = BackupUnscheduleRequest{} }
func (m *BackupUnscheduleRequest) String() string            { return proto.CompactTextString(m) }
func (*BackupUnscheduleRequest) ProtoMessage()               {}
func (*BackupUnscheduleRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

// Next Id: 19
type SnapshotRestoreRequest struct {
	Cluster        string `protobuf:"bytes,1,opt,name=cluster" json:"cluster,omitempty"`
	Uid            string `protobuf:"bytes,2,opt,name=uid" json:"uid,omitempty"`
	SnapshotPhid   string `protobuf:"bytes,8,opt,name=snapshot_phid,json=snapshotPhid" json:"snapshot_phid,omitempty"`
	AuthSecretName string `protobuf:"bytes,13,opt,name=auth_secret_name,json=authSecretName" json:"auth_secret_name,omitempty"`
}

func (m *SnapshotRestoreRequest) Reset()                    { *m = SnapshotRestoreRequest{} }
func (m *SnapshotRestoreRequest) String() string            { return proto.CompactTextString(m) }
func (*SnapshotRestoreRequest) ProtoMessage()               {}
func (*SnapshotRestoreRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{6} }

// Next Id: 4
type SnapshotCheckRequest struct {
	Cluster string `protobuf:"bytes,1,opt,name=cluster" json:"cluster,omitempty"`
	Phid    string `protobuf:"bytes,2,opt,name=phid" json:"phid,omitempty"`
}

func (m *SnapshotCheckRequest) Reset()                    { *m = SnapshotCheckRequest{} }
func (m *SnapshotCheckRequest) String() string            { return proto.CompactTextString(m) }
func (*SnapshotCheckRequest) ProtoMessage()               {}
func (*SnapshotCheckRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{7} }

func init() {
	proto.RegisterType((*DatabaseInfo)(nil), "db.v1beta1.DatabaseInfo")
	proto.RegisterType((*Snapshot)(nil), "db.v1beta1.Snapshot")
	proto.RegisterType((*SnapshotListRequest)(nil), "db.v1beta1.SnapshotListRequest")
	proto.RegisterType((*SnapshotListResponse)(nil), "db.v1beta1.SnapshotListResponse")
	proto.RegisterType((*BackupScheduleRequest)(nil), "db.v1beta1.BackupScheduleRequest")
	proto.RegisterType((*BackupUnscheduleRequest)(nil), "db.v1beta1.BackupUnscheduleRequest")
	proto.RegisterType((*SnapshotRestoreRequest)(nil), "db.v1beta1.SnapshotRestoreRequest")
	proto.RegisterType((*SnapshotCheckRequest)(nil), "db.v1beta1.SnapshotCheckRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Snapshots service

type SnapshotsClient interface {
	List(ctx context.Context, in *SnapshotListRequest, opts ...grpc.CallOption) (*SnapshotListResponse, error)
	BackupSchedule(ctx context.Context, in *BackupScheduleRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
	BackupUnschedule(ctx context.Context, in *BackupUnscheduleRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
	Restore(ctx context.Context, in *SnapshotRestoreRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
}

type snapshotsClient struct {
	cc *grpc.ClientConn
}

func NewSnapshotsClient(cc *grpc.ClientConn) SnapshotsClient {
	return &snapshotsClient{cc}
}

func (c *snapshotsClient) List(ctx context.Context, in *SnapshotListRequest, opts ...grpc.CallOption) (*SnapshotListResponse, error) {
	out := new(SnapshotListResponse)
	err := grpc.Invoke(ctx, "/db.v1beta1.Snapshots/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *snapshotsClient) BackupSchedule(ctx context.Context, in *BackupScheduleRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/db.v1beta1.Snapshots/BackupSchedule", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *snapshotsClient) BackupUnschedule(ctx context.Context, in *BackupUnscheduleRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/db.v1beta1.Snapshots/BackupUnschedule", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *snapshotsClient) Restore(ctx context.Context, in *SnapshotRestoreRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/db.v1beta1.Snapshots/Restore", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Snapshots service

type SnapshotsServer interface {
	List(context.Context, *SnapshotListRequest) (*SnapshotListResponse, error)
	BackupSchedule(context.Context, *BackupScheduleRequest) (*dtypes.VoidResponse, error)
	BackupUnschedule(context.Context, *BackupUnscheduleRequest) (*dtypes.VoidResponse, error)
	Restore(context.Context, *SnapshotRestoreRequest) (*dtypes.VoidResponse, error)
}

func RegisterSnapshotsServer(s *grpc.Server, srv SnapshotsServer) {
	s.RegisterService(&_Snapshots_serviceDesc, srv)
}

func _Snapshots_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SnapshotListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SnapshotsServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/db.v1beta1.Snapshots/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SnapshotsServer).List(ctx, req.(*SnapshotListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Snapshots_BackupSchedule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BackupScheduleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SnapshotsServer).BackupSchedule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/db.v1beta1.Snapshots/BackupSchedule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SnapshotsServer).BackupSchedule(ctx, req.(*BackupScheduleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Snapshots_BackupUnschedule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BackupUnscheduleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SnapshotsServer).BackupUnschedule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/db.v1beta1.Snapshots/BackupUnschedule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SnapshotsServer).BackupUnschedule(ctx, req.(*BackupUnscheduleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Snapshots_Restore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SnapshotRestoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SnapshotsServer).Restore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/db.v1beta1.Snapshots/Restore",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SnapshotsServer).Restore(ctx, req.(*SnapshotRestoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Snapshots_serviceDesc = grpc.ServiceDesc{
	ServiceName: "db.v1beta1.Snapshots",
	HandlerType: (*SnapshotsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _Snapshots_List_Handler,
		},
		{
			MethodName: "BackupSchedule",
			Handler:    _Snapshots_BackupSchedule_Handler,
		},
		{
			MethodName: "BackupUnschedule",
			Handler:    _Snapshots_BackupUnschedule_Handler,
		},
		{
			MethodName: "Restore",
			Handler:    _Snapshots_Restore_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor1,
}

func init() { proto.RegisterFile("snapshot.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 808 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x55, 0xcb, 0x6e, 0xe3, 0x36,
	0x14, 0x85, 0x6c, 0x27, 0xb1, 0xaf, 0x1f, 0x71, 0x99, 0x34, 0x15, 0x8c, 0xb6, 0x71, 0x15, 0xb4,
	0x75, 0x83, 0xd6, 0x42, 0xdc, 0x5d, 0x80, 0x2e, 0xf2, 0x5a, 0x14, 0xe8, 0x23, 0xb0, 0x91, 0xa2,
	0x68, 0x50, 0xa8, 0xb4, 0xc8, 0xc4, 0x82, 0x5d, 0x51, 0x25, 0xa9, 0x20, 0x45, 0x90, 0x4d, 0x56,
	0xb3, 0xcf, 0x3f, 0xcc, 0x60, 0x3e, 0x60, 0x30, 0xc0, 0xfc, 0xc6, 0x2c, 0xe6, 0x07, 0xe6, 0x23,
	0x66, 0x39, 0x10, 0x29, 0xca, 0x72, 0x62, 0xcf, 0x00, 0x41, 0x36, 0x01, 0x79, 0x78, 0x74, 0x73,
	0xee, 0xe1, 0xb9, 0x34, 0x34, 0x44, 0x88, 0x23, 0x31, 0x62, 0xb2, 0x1b, 0x71, 0x26, 0x19, 0x02,
	0x32, 0xec, 0x5e, 0xec, 0x0c, 0xa9, 0xc4, 0x3b, 0xad, 0xcf, 0xcf, 0x19, 0x3b, 0x9f, 0x50, 0x17,
	0x47, 0x81, 0x8b, 0xc3, 0x90, 0x49, 0x2c, 0x03, 0x16, 0x0a, 0xcd, 0x6c, 0x7d, 0x89, 0xa3, 0x48,
	0xf8, 0x8c, 0x2c, 0x3a, 0xdf, 0x48, 0x60, 0x22, 0xff, 0x8f, 0xa8, 0x70, 0xd5, 0x5f, 0x8d, 0x3b,
	0xc7, 0x50, 0x3b, 0xc4, 0x12, 0x0f, 0xb1, 0xa0, 0x3f, 0x87, 0x67, 0x0c, 0xd9, 0xb0, 0xe2, 0x4f,
	0x62, 0x21, 0x29, 0xb7, 0xad, 0xb6, 0xd5, 0xa9, 0xf4, 0xcd, 0x16, 0x21, 0x28, 0x85, 0xf8, 0x5f,
	0x6a, 0x17, 0x14, 0xac, 0xd6, 0x09, 0x96, 0x14, 0xb3, 0x8b, 0x1a, 0x4b, 0xd6, 0xce, 0xbb, 0x02,
	0x94, 0x07, 0x69, 0x1b, 0xd9, 0x47, 0xd6, 0xec, 0x47, 0xd1, 0x28, 0x20, 0xa6, 0x50, 0xb2, 0x46,
	0x2d, 0x28, 0x47, 0x9c, 0x5d, 0x04, 0x84, 0xf2, 0xb4, 0x58, 0xb6, 0x47, 0xdf, 0x41, 0xd3, 0x9f,
	0xb0, 0x98, 0x78, 0x3e, 0xa7, 0x84, 0x86, 0x32, 0xc0, 0x13, 0xbb, 0xa4, 0x38, 0xab, 0x0a, 0x3f,
	0xc8, 0x60, 0xb4, 0x01, 0xcb, 0xc3, 0xd8, 0x1f, 0x53, 0x69, 0x2f, 0x29, 0x42, 0xba, 0x4b, 0x70,
	0x21, 0xb1, 0x8c, 0x85, 0x5d, 0xd6, 0xb8, 0xde, 0x25, 0xdd, 0x46, 0x9c, 0xf9, 0x54, 0x08, 0xbb,
	0xa2, 0xbb, 0x4d, 0xb7, 0x59, 0x67, 0x30, 0xed, 0x0c, 0x75, 0x61, 0x2d, 0x10, 0x9e, 0xf0, 0x47,
	0x94, 0xc4, 0x13, 0x4a, 0xbc, 0x21, 0xf6, 0xc7, 0x71, 0x64, 0x57, 0xdb, 0x56, 0x67, 0xa9, 0xff,
	0x49, 0x20, 0x06, 0xe6, 0x64, 0x5f, 0x1d, 0xa0, 0x6f, 0x61, 0x55, 0xb0, 0x98, 0xfb, 0xd4, 0x23,
	0xa9, 0xc5, 0x76, 0x4d, 0x95, 0x6b, 0x68, 0xd8, 0x18, 0x9f, 0x23, 0x9a, 0xfb, 0xb7, 0xeb, 0x79,
	0x62, 0x66, 0xe7, 0x17, 0x00, 0x3e, 0xa7, 0x58, 0x52, 0xe2, 0x61, 0x69, 0x37, 0xda, 0x56, 0xa7,
	0xd8, 0xaf, 0xa4, 0xc8, 0x9e, 0x74, 0xf6, 0x60, 0xcd, 0x50, 0x7f, 0x09, 0x84, 0xec, 0xd3, 0xff,
	0x62, 0x2a, 0xe4, 0x07, 0xee, 0xb4, 0x09, 0xc5, 0x38, 0xbb, 0x89, 0x64, 0xe9, 0x3c, 0xb3, 0x60,
	0x7d, 0xb6, 0x86, 0x88, 0x58, 0x28, 0x28, 0xfa, 0x26, 0xb3, 0x30, 0xa9, 0x51, 0xed, 0x35, 0xba,
	0x3a, 0x4d, 0xdd, 0x81, 0x42, 0x33, 0x4b, 0x7f, 0x82, 0xba, 0xe9, 0xd6, 0x0b, 0xc2, 0x33, 0xa6,
	0x8a, 0x57, 0x7b, 0x76, 0x77, 0x1a, 0xe5, 0x6e, 0x3e, 0x71, 0xfd, 0x1a, 0xc9, 0xe7, 0xaf, 0x03,
	0xa5, 0x49, 0x20, 0xa4, 0x5d, 0x6c, 0x17, 0x3b, 0xd5, 0xde, 0x7a, 0xfe, 0x2b, 0x23, 0xab, 0xaf,
	0x18, 0xce, 0x93, 0x02, 0x7c, 0xaa, 0x8d, 0x36, 0xbe, 0x3f, 0xa0, 0x5f, 0xf4, 0x35, 0x34, 0xa6,
	0xb1, 0xf2, 0x92, 0x43, 0x1d, 0xad, 0xfa, 0x14, 0x3d, 0x09, 0x08, 0xda, 0x84, 0xaa, 0x8e, 0x92,
	0xa7, 0xe2, 0xac, 0xd3, 0x05, 0x1a, 0xfa, 0x2d, 0x09, 0xf5, 0x16, 0xd4, 0xcd, 0xdd, 0x69, 0xca,
	0xb2, 0xa2, 0xd4, 0x0c, 0xa8, 0x48, 0x1d, 0x68, 0xe2, 0x58, 0x8e, 0x3c, 0x41, 0x7d, 0x6e, 0x4a,
	0xe9, 0x40, 0x36, 0x12, 0x7c, 0xa0, 0x60, 0xc5, 0xfc, 0x1e, 0x90, 0xc9, 0x99, 0xe7, 0x73, 0x16,
	0x7a, 0xf4, 0x32, 0xe2, 0x69, 0x46, 0x9b, 0xe6, 0xe4, 0x80, 0xb3, 0xf0, 0xe8, 0x32, 0xe2, 0xce,
	0x11, 0x7c, 0xa6, 0x9d, 0x38, 0x09, 0xc5, 0xc3, 0xbd, 0x70, 0x6e, 0x2d, 0xd8, 0xc8, 0x4c, 0xa6,
	0x42, 0x32, 0xfe, 0x20, 0x4b, 0xf3, 0x56, 0xa8, 0x41, 0x2f, 0xcf, 0x5a, 0x71, 0x9c, 0x0c, 0xfc,
	0x3c, 0x2b, 0xea, 0xf3, 0xac, 0x70, 0x0e, 0xa7, 0x81, 0x3c, 0x18, 0x51, 0x7f, 0xfc, 0x71, 0x49,
	0x73, 0x1e, 0x98, 0xde, 0x9b, 0x25, 0xa8, 0x98, 0x32, 0x02, 0x3d, 0xb5, 0xa0, 0x94, 0xa4, 0x1b,
	0x6d, 0xce, 0x0b, 0x58, 0x6e, 0x76, 0x5a, 0xed, 0xc5, 0x04, 0x3d, 0x18, 0xce, 0xe9, 0xcd, 0x0b,
	0xbb, 0x50, 0xb6, 0x6e, 0x5e, 0xbf, 0xbd, 0x2d, 0xfc, 0x8e, 0x7e, 0x75, 0x67, 0x1e, 0xe2, 0x71,
	0x3c, 0xa4, 0x3c, 0xa4, 0x92, 0x0a, 0x37, 0x2d, 0xe2, 0xa6, 0x2a, 0x85, 0x7b, 0x95, 0xae, 0xae,
	0x5d, 0x33, 0x01, 0xc2, 0xbd, 0x8a, 0x03, 0x72, 0xed, 0x8a, 0x4c, 0xe8, 0x4b, 0x0b, 0x1a, 0xb3,
	0x21, 0x47, 0x5f, 0xe5, 0x15, 0xcd, 0x1d, 0x80, 0xd6, 0xba, 0x99, 0xcd, 0x3f, 0x58, 0x40, 0x32,
	0xa1, 0x93, 0x9c, 0xd0, 0x7f, 0x9c, 0xd3, 0xc7, 0x10, 0x8a, 0x7d, 0xf5, 0x1b, 0xe3, 0x9a, 0xd4,
	0xfd, 0xa0, 0x9f, 0xc4, 0x5d, 0x6b, 0x1b, 0xbd, 0xb2, 0xa0, 0x79, 0x37, 0x94, 0x68, 0xeb, 0xbe,
	0xf6, 0x7b, 0x91, 0x5d, 0xa0, 0x7e, 0x9c, 0x53, 0xef, 0xb5, 0xfe, 0x7e, 0x4c, 0xf5, 0x71, 0x78,
	0x47, 0x3f, 0x7a, 0x6e, 0xc1, 0x4a, 0x3a, 0x01, 0xc8, 0x99, 0xfb, 0x06, 0xcd, 0x8c, 0xc7, 0x02,
	0xc9, 0x24, 0x27, 0xf9, 0x4f, 0x67, 0xf0, 0x98, 0x92, 0xb9, 0xfe, 0xff, 0xbb, 0xd6, 0xf6, 0x7e,
	0xe5, 0xaf, 0x95, 0xb4, 0xc0, 0x70, 0x59, 0xfd, 0xa6, 0xff, 0xf8, 0x3e, 0x00, 0x00, 0xff, 0xff,
	0x04, 0xe2, 0x2d, 0x2a, 0x47, 0x08, 0x00, 0x00,
}