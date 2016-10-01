// Code generated by protoc-gen-go.
// source: job.proto
// DO NOT EDIT!

package ci

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

type JobListRequest struct {
	Parents string `protobuf:"bytes,1,opt,name=parents" json:"parents,omitempty"`
}

func (m *JobListRequest) Reset()                    { *m = JobListRequest{} }
func (m *JobListRequest) String() string            { return proto.CompactTextString(m) }
func (*JobListRequest) ProtoMessage()               {}
func (*JobListRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

type JobListResponse struct {
	Status *dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Jobs   []*Job         `protobuf:"bytes,2,rep,name=jobs" json:"jobs,omitempty"`
}

func (m *JobListResponse) Reset()                    { *m = JobListResponse{} }
func (m *JobListResponse) String() string            { return proto.CompactTextString(m) }
func (*JobListResponse) ProtoMessage()               {}
func (*JobListResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *JobListResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *JobListResponse) GetJobs() []*Job {
	if m != nil {
		return m.Jobs
	}
	return nil
}

type JobBuildRequest struct {
	Name    string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Parents string `protobuf:"bytes,2,opt,name=parents" json:"parents,omitempty"`
	Param   string `protobuf:"bytes,3,opt,name=param" json:"param,omitempty"`
}

func (m *JobBuildRequest) Reset()                    { *m = JobBuildRequest{} }
func (m *JobBuildRequest) String() string            { return proto.CompactTextString(m) }
func (*JobBuildRequest) ProtoMessage()               {}
func (*JobBuildRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

type JobDescribeRequest struct {
	Name    string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Parents string `protobuf:"bytes,2,opt,name=parents" json:"parents,omitempty"`
}

func (m *JobDescribeRequest) Reset()                    { *m = JobDescribeRequest{} }
func (m *JobDescribeRequest) String() string            { return proto.CompactTextString(m) }
func (*JobDescribeRequest) ProtoMessage()               {}
func (*JobDescribeRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{3} }

type JobDescribeResponse struct {
	Status                *dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Name                  string         `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Parents               string         `protobuf:"bytes,3,opt,name=parents" json:"parents,omitempty"`
	Description           string         `protobuf:"bytes,4,opt,name=description" json:"description,omitempty"`
	DisplayName           string         `protobuf:"bytes,5,opt,name=display_name,json=displayName" json:"display_name,omitempty"`
	JobType               string         `protobuf:"bytes,6,opt,name=job_type,json=jobType" json:"job_type,omitempty"`
	JobColor              string         `protobuf:"bytes,7,opt,name=job_color,json=jobColor" json:"job_color,omitempty"`
	Buildable             bool           `protobuf:"varint,8,opt,name=buildable" json:"buildable,omitempty"`
	BuildCount            int64          `protobuf:"varint,9,opt,name=build_count,json=buildCount" json:"build_count,omitempty"`
	BuildIds              []int64        `protobuf:"varint,10,rep,packed,name=build_ids,json=buildIds" json:"build_ids,omitempty"`
	FirstBuildId          int64          `protobuf:"varint,11,opt,name=first_build_id,json=firstBuildId" json:"first_build_id,omitempty"`
	LastBuildId           int64          `protobuf:"varint,12,opt,name=last_build_id,json=lastBuildId" json:"last_build_id,omitempty"`
	LastCompletedBuildId  int64          `protobuf:"varint,13,opt,name=last_completed_build_id,json=lastCompletedBuildId" json:"last_completed_build_id,omitempty"`
	LastFailedBuildId     int64          `protobuf:"varint,14,opt,name=last_failed_build_id,json=lastFailedBuildId" json:"last_failed_build_id,omitempty"`
	LastSuccessfulBuildId int64          `protobuf:"varint,15,opt,name=last_successful_build_id,json=lastSuccessfulBuildId" json:"last_successful_build_id,omitempty"`
	NextBuildNumber       int64          `protobuf:"varint,16,opt,name=next_build_number,json=nextBuildNumber" json:"next_build_number,omitempty"`
	HealthReport          []*JobHealth   `protobuf:"bytes,17,rep,name=health_report,json=healthReport" json:"health_report,omitempty"`
}

func (m *JobDescribeResponse) Reset()                    { *m = JobDescribeResponse{} }
func (m *JobDescribeResponse) String() string            { return proto.CompactTextString(m) }
func (*JobDescribeResponse) ProtoMessage()               {}
func (*JobDescribeResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{4} }

func (m *JobDescribeResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *JobDescribeResponse) GetHealthReport() []*JobHealth {
	if m != nil {
		return m.HealthReport
	}
	return nil
}

type JobHealth struct {
	Description string `protobuf:"bytes,1,opt,name=description" json:"description,omitempty"`
	Score       int64  `protobuf:"varint,4,opt,name=score" json:"score,omitempty"`
}

func (m *JobHealth) Reset()                    { *m = JobHealth{} }
func (m *JobHealth) String() string            { return proto.CompactTextString(m) }
func (*JobHealth) ProtoMessage()               {}
func (*JobHealth) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{5} }

type JobDeleteRequest struct {
	Name    string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Parents string `protobuf:"bytes,2,opt,name=parents" json:"parents,omitempty"`
}

func (m *JobDeleteRequest) Reset()                    { *m = JobDeleteRequest{} }
func (m *JobDeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*JobDeleteRequest) ProtoMessage()               {}
func (*JobDeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{6} }

type JobCreateRequest struct {
	Name    string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Parents string `protobuf:"bytes,2,opt,name=parents" json:"parents,omitempty"`
	ShFile  string `protobuf:"bytes,3,opt,name=sh_file,json=shFile" json:"sh_file,omitempty"`
}

func (m *JobCreateRequest) Reset()                    { *m = JobCreateRequest{} }
func (m *JobCreateRequest) String() string            { return proto.CompactTextString(m) }
func (*JobCreateRequest) ProtoMessage()               {}
func (*JobCreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{7} }

type JobCopyRequest struct {
	Source      string `protobuf:"bytes,1,opt,name=source" json:"source,omitempty"`
	Destination string `protobuf:"bytes,2,opt,name=destination" json:"destination,omitempty"`
}

func (m *JobCopyRequest) Reset()                    { *m = JobCopyRequest{} }
func (m *JobCopyRequest) String() string            { return proto.CompactTextString(m) }
func (*JobCopyRequest) ProtoMessage()               {}
func (*JobCopyRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{8} }

type Job struct {
	Name    string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Parents string `protobuf:"bytes,2,opt,name=parents" json:"parents,omitempty"`
	JobType string `protobuf:"bytes,3,opt,name=job_type,json=jobType" json:"job_type,omitempty"`
	Color   string `protobuf:"bytes,4,opt,name=color" json:"color,omitempty"`
	Url     string `protobuf:"bytes,5,opt,name=url" json:"url,omitempty"`
}

func (m *Job) Reset()                    { *m = Job{} }
func (m *Job) String() string            { return proto.CompactTextString(m) }
func (*Job) ProtoMessage()               {}
func (*Job) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{9} }

func init() {
	proto.RegisterType((*JobListRequest)(nil), "ci.JobListRequest")
	proto.RegisterType((*JobListResponse)(nil), "ci.JobListResponse")
	proto.RegisterType((*JobBuildRequest)(nil), "ci.JobBuildRequest")
	proto.RegisterType((*JobDescribeRequest)(nil), "ci.JobDescribeRequest")
	proto.RegisterType((*JobDescribeResponse)(nil), "ci.JobDescribeResponse")
	proto.RegisterType((*JobHealth)(nil), "ci.JobHealth")
	proto.RegisterType((*JobDeleteRequest)(nil), "ci.JobDeleteRequest")
	proto.RegisterType((*JobCreateRequest)(nil), "ci.JobCreateRequest")
	proto.RegisterType((*JobCopyRequest)(nil), "ci.JobCopyRequest")
	proto.RegisterType((*Job)(nil), "ci.Job")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Jobs service

type JobsClient interface {
	List(ctx context.Context, in *JobListRequest, opts ...grpc.CallOption) (*JobListResponse, error)
	Describe(ctx context.Context, in *JobDescribeRequest, opts ...grpc.CallOption) (*JobDescribeResponse, error)
	Create(ctx context.Context, in *JobCreateRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
	Copy(ctx context.Context, in *JobCopyRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
	Build(ctx context.Context, in *JobBuildRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
	Delete(ctx context.Context, in *JobDeleteRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
}

type jobsClient struct {
	cc *grpc.ClientConn
}

func NewJobsClient(cc *grpc.ClientConn) JobsClient {
	return &jobsClient{cc}
}

func (c *jobsClient) List(ctx context.Context, in *JobListRequest, opts ...grpc.CallOption) (*JobListResponse, error) {
	out := new(JobListResponse)
	err := grpc.Invoke(ctx, "/ci.Jobs/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobsClient) Describe(ctx context.Context, in *JobDescribeRequest, opts ...grpc.CallOption) (*JobDescribeResponse, error) {
	out := new(JobDescribeResponse)
	err := grpc.Invoke(ctx, "/ci.Jobs/Describe", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobsClient) Create(ctx context.Context, in *JobCreateRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/ci.Jobs/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobsClient) Copy(ctx context.Context, in *JobCopyRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/ci.Jobs/Copy", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobsClient) Build(ctx context.Context, in *JobBuildRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/ci.Jobs/Build", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobsClient) Delete(ctx context.Context, in *JobDeleteRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/ci.Jobs/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Jobs service

type JobsServer interface {
	List(context.Context, *JobListRequest) (*JobListResponse, error)
	Describe(context.Context, *JobDescribeRequest) (*JobDescribeResponse, error)
	Create(context.Context, *JobCreateRequest) (*dtypes.VoidResponse, error)
	Copy(context.Context, *JobCopyRequest) (*dtypes.VoidResponse, error)
	Build(context.Context, *JobBuildRequest) (*dtypes.VoidResponse, error)
	Delete(context.Context, *JobDeleteRequest) (*dtypes.VoidResponse, error)
}

func RegisterJobsServer(s *grpc.Server, srv JobsServer) {
	s.RegisterService(&_Jobs_serviceDesc, srv)
}

func _Jobs_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JobListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobsServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ci.Jobs/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobsServer).List(ctx, req.(*JobListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Jobs_Describe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JobDescribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobsServer).Describe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ci.Jobs/Describe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobsServer).Describe(ctx, req.(*JobDescribeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Jobs_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JobCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobsServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ci.Jobs/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobsServer).Create(ctx, req.(*JobCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Jobs_Copy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JobCopyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobsServer).Copy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ci.Jobs/Copy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobsServer).Copy(ctx, req.(*JobCopyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Jobs_Build_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JobBuildRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobsServer).Build(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ci.Jobs/Build",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobsServer).Build(ctx, req.(*JobBuildRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Jobs_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JobDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobsServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ci.Jobs/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobsServer).Delete(ctx, req.(*JobDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Jobs_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ci.Jobs",
	HandlerType: (*JobsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _Jobs_List_Handler,
		},
		{
			MethodName: "Describe",
			Handler:    _Jobs_Describe_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _Jobs_Create_Handler,
		},
		{
			MethodName: "Copy",
			Handler:    _Jobs_Copy_Handler,
		},
		{
			MethodName: "Build",
			Handler:    _Jobs_Build_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Jobs_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor2,
}

func init() { proto.RegisterFile("job.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 839 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x55, 0xdd, 0x6e, 0xe3, 0x44,
	0x14, 0x56, 0xe2, 0xfc, 0x9e, 0xfc, 0xb4, 0x9d, 0x0d, 0xad, 0xf1, 0x56, 0x22, 0x6b, 0xd0, 0x12,
	0xa2, 0x55, 0xcc, 0x06, 0x01, 0x12, 0xe2, 0x02, 0x35, 0x68, 0x05, 0x11, 0xea, 0x85, 0x0b, 0x95,
	0xb8, 0xb2, 0xfc, 0x33, 0x69, 0x26, 0x72, 0x3c, 0xc6, 0x63, 0x57, 0x44, 0x55, 0x6f, 0x78, 0x01,
	0x2e, 0x78, 0x34, 0x5e, 0x80, 0x0b, 0x1e, 0x04, 0xcd, 0x19, 0x3b, 0x75, 0x5a, 0xb2, 0x55, 0x7b,
	0x13, 0x65, 0xbe, 0xf3, 0x9d, 0xef, 0x9b, 0x99, 0x33, 0xe7, 0x18, 0xda, 0x2b, 0xee, 0x4d, 0xe2,
	0x84, 0xa7, 0x9c, 0x54, 0x7d, 0x66, 0x9c, 0x5e, 0x71, 0x7e, 0x15, 0x52, 0xcb, 0x8d, 0x99, 0xe5,
	0x46, 0x11, 0x4f, 0xdd, 0x94, 0xf1, 0x48, 0x28, 0x86, 0x71, 0x2c, 0xe1, 0x20, 0xdd, 0xc4, 0x54,
	0x58, 0xf8, 0xab, 0x70, 0x73, 0x0c, 0xfd, 0x39, 0xf7, 0x7e, 0x62, 0x22, 0xb5, 0xe9, 0x6f, 0x19,
	0x15, 0x29, 0xd1, 0xa1, 0x19, 0xbb, 0x09, 0x8d, 0x52, 0xa1, 0x57, 0x86, 0x95, 0x51, 0xdb, 0x2e,
	0x96, 0xe6, 0x25, 0x1c, 0x6c, 0xb9, 0x22, 0xe6, 0x91, 0xa0, 0xe4, 0x35, 0x34, 0x44, 0xea, 0xa6,
	0x99, 0xe2, 0x76, 0xa6, 0xfd, 0x89, 0xf2, 0x98, 0x5c, 0x20, 0x6a, 0xe7, 0x51, 0xf2, 0x12, 0x6a,
	0x2b, 0xee, 0x09, 0xbd, 0x3a, 0xd4, 0x46, 0x9d, 0x69, 0x73, 0xe2, 0xb3, 0xc9, 0x9c, 0x7b, 0x36,
	0x82, 0xe6, 0x2f, 0xa8, 0x7b, 0x96, 0xb1, 0x30, 0x28, 0x36, 0x41, 0xa0, 0x16, 0xb9, 0x6b, 0x9a,
	0xef, 0x00, 0xff, 0x97, 0x37, 0x56, 0xdd, 0xd9, 0x18, 0x19, 0x40, 0x3d, 0x76, 0x13, 0x77, 0xad,
	0x6b, 0x88, 0xab, 0x85, 0x79, 0x06, 0x64, 0xce, 0xbd, 0xef, 0xa9, 0xf0, 0x13, 0xe6, 0xd1, 0x67,
	0x29, 0x9b, 0x7f, 0xd6, 0xe1, 0xc5, 0x8e, 0xc8, 0x13, 0xcf, 0x5d, 0xb8, 0x55, 0xff, 0xdf, 0x4d,
	0xdb, 0x3d, 0xc7, 0x10, 0x3a, 0x01, 0x3a, 0xc5, 0xb2, 0x74, 0x7a, 0x0d, 0xa3, 0x65, 0x88, 0xbc,
	0x82, 0x6e, 0xc0, 0x44, 0x1c, 0xba, 0x1b, 0x07, 0x75, 0xeb, 0x39, 0x45, 0x61, 0xe7, 0x52, 0xfe,
	0x43, 0x68, 0xad, 0xb8, 0xe7, 0xc8, 0xed, 0xe8, 0x0d, 0xa5, 0xbf, 0xe2, 0xde, 0xcf, 0x9b, 0x98,
	0x92, 0x97, 0xf8, 0x66, 0x1c, 0x9f, 0x87, 0x3c, 0xd1, 0x9b, 0x18, 0x93, 0xdc, 0x99, 0x5c, 0x93,
	0x53, 0x68, 0x7b, 0xb2, 0x04, 0xae, 0x17, 0x52, 0xbd, 0x35, 0xac, 0x8c, 0x5a, 0xf6, 0x1d, 0x40,
	0x3e, 0x82, 0x0e, 0x2e, 0x1c, 0x9f, 0x67, 0x51, 0xaa, 0xb7, 0x87, 0x95, 0x91, 0x66, 0x03, 0x42,
	0x33, 0x89, 0x48, 0x6d, 0x45, 0x60, 0x81, 0xd0, 0x61, 0xa8, 0x8d, 0x34, 0xbb, 0x85, 0xc0, 0x8f,
	0x81, 0x20, 0x9f, 0x40, 0x7f, 0xc1, 0x12, 0x91, 0x3a, 0x05, 0x45, 0xef, 0xa0, 0x40, 0x17, 0xd1,
	0x33, 0x45, 0x23, 0x26, 0xf4, 0x42, 0xb7, 0x4c, 0xea, 0x22, 0xa9, 0x23, 0xc1, 0x82, 0xf3, 0x25,
	0x9c, 0x20, 0xc7, 0xe7, 0xeb, 0x38, 0xa4, 0x29, 0x0d, 0xee, 0xd8, 0x3d, 0x64, 0x0f, 0x64, 0x78,
	0x56, 0x44, 0x8b, 0x34, 0x0b, 0x10, 0x77, 0x16, 0x2e, 0x0b, 0xcb, 0x39, 0x7d, 0xcc, 0x39, 0x92,
	0xb1, 0x77, 0x18, 0x2a, 0x12, 0xbe, 0x06, 0x1d, 0x13, 0x44, 0xe6, 0xfb, 0x54, 0x88, 0x45, 0x16,
	0xde, 0x25, 0x1d, 0x60, 0xd2, 0x07, 0x32, 0x7e, 0xb1, 0x0d, 0x17, 0x89, 0x63, 0x38, 0x8a, 0xe8,
	0xef, 0xc5, 0x21, 0xa2, 0x6c, 0xed, 0xd1, 0x44, 0x3f, 0xc4, 0x8c, 0x03, 0x19, 0x40, 0xde, 0x39,
	0xc2, 0x64, 0x0a, 0xbd, 0x25, 0x75, 0xc3, 0x74, 0xe9, 0x24, 0x34, 0xe6, 0x49, 0xaa, 0x1f, 0x61,
	0x7b, 0xf4, 0xf2, 0xf6, 0xf8, 0x01, 0x63, 0x76, 0x57, 0x71, 0x6c, 0xa4, 0x98, 0x33, 0x68, 0x6f,
	0x43, 0xf7, 0x1f, 0x4c, 0xe5, 0xe1, 0x83, 0x19, 0x40, 0x5d, 0xf8, 0x3c, 0xa1, 0xf8, 0x98, 0x34,
	0x5b, 0x2d, 0xcc, 0xef, 0xe0, 0x10, 0x5f, 0xb5, 0xbc, 0xa3, 0xe7, 0x35, 0xc6, 0xaf, 0xa8, 0x30,
	0x4b, 0xa8, 0xfb, 0x4c, 0x05, 0x72, 0x02, 0x4d, 0xb1, 0x74, 0x16, 0x2c, 0xa4, 0x79, 0x1b, 0x34,
	0xc4, 0xf2, 0x1d, 0x0b, 0xa9, 0x39, 0xc7, 0x91, 0x34, 0xe3, 0xf1, 0xa6, 0x10, 0x3e, 0x86, 0x86,
	0xe0, 0x59, 0xe2, 0x17, 0xd2, 0xf9, 0x2a, 0x3f, 0x7e, 0xca, 0x22, 0x1c, 0x75, 0xb9, 0x41, 0x19,
	0x32, 0xaf, 0x41, 0x9b, 0x73, 0xef, 0x89, 0x3b, 0x2b, 0x77, 0x90, 0xb6, 0xdb, 0x41, 0x03, 0xa8,
	0xab, 0xee, 0x51, 0xbd, 0xa9, 0x16, 0xe4, 0x10, 0xb4, 0x2c, 0x09, 0xf3, 0x66, 0x94, 0x7f, 0xa7,
	0xff, 0xd4, 0xa0, 0x36, 0xe7, 0x9e, 0x20, 0x17, 0x50, 0x93, 0x03, 0x93, 0x90, 0xbc, 0xa6, 0xa5,
	0x49, 0x6b, 0xbc, 0xd8, 0xc1, 0xd4, 0x64, 0x31, 0xcd, 0x3f, 0xfe, 0xfe, 0xf7, 0xaf, 0xea, 0x29,
	0x31, 0x2c, 0x37, 0x8e, 0x85, 0xcf, 0x03, 0x35, 0xd1, 0x7d, 0x66, 0x5d, 0x7f, 0x3e, 0x79, 0x6b,
	0xc9, 0x81, 0x49, 0x16, 0xd0, 0x2a, 0x26, 0x12, 0x39, 0xce, 0x45, 0xee, 0xcd, 0x39, 0xe3, 0xe4,
	0x01, 0x9e, 0x1b, 0x7c, 0x86, 0x06, 0x1f, 0x93, 0x57, 0xfb, 0x0d, 0xac, 0x1b, 0x79, 0x43, 0xb7,
	0xc4, 0x83, 0x86, 0xaa, 0x30, 0x19, 0xe4, 0x6a, 0x3b, 0x05, 0x37, 0x06, 0xc5, 0xd4, 0xbb, 0xe4,
	0x2c, 0xd8, 0x1a, 0xbc, 0x41, 0x83, 0xd7, 0xc6, 0xe3, 0x06, 0xdf, 0x54, 0xc6, 0x24, 0x82, 0x9a,
	0x2c, 0xf5, 0xf6, 0x82, 0x4a, 0x75, 0xdf, 0xa3, 0xff, 0x2d, 0xea, 0x7f, 0x65, 0xbe, 0x7d, 0x9f,
	0xbe, 0x7a, 0x21, 0xb7, 0xd6, 0x4d, 0xe9, 0x35, 0xa0, 0xdf, 0x02, 0xea, 0xd8, 0x82, 0xa4, 0xb8,
	0xfd, 0xf2, 0x77, 0x67, 0x8f, 0xe3, 0x14, 0x1d, 0xdf, 0x18, 0x9f, 0x3e, 0x7a, 0x22, 0x0b, 0xdb,
	0x5e, 0xfa, 0x38, 0xd0, 0x50, 0xfd, 0xb5, 0xbd, 0xbb, 0x9d, 0x76, 0xdb, 0xe3, 0x94, 0x17, 0x67,
	0xfc, 0xf8, 0xdd, 0x79, 0x0d, 0xfc, 0x80, 0x7f, 0xf1, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa6,
	0x2b, 0xd7, 0x31, 0x07, 0x08, 0x00, 0x00,
}
