// Code generated by protoc-gen-go. DO NOT EDIT.
// source: builds/v1/builds.proto

package buildsv1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type BuildStatus int32

const (
	BuildStatus_NULL    BuildStatus = 0
	BuildStatus_SUCCESS BuildStatus = 1
	BuildStatus_FAILURE BuildStatus = 2
	BuildStatus_PENDING BuildStatus = 3
	BuildStatus_RUNNING BuildStatus = 4
)

var BuildStatus_name = map[int32]string{
	0: "NULL",
	1: "SUCCESS",
	2: "FAILURE",
	3: "PENDING",
	4: "RUNNING",
}

var BuildStatus_value = map[string]int32{
	"NULL":    0,
	"SUCCESS": 1,
	"FAILURE": 2,
	"PENDING": 3,
	"RUNNING": 4,
}

func (x BuildStatus) String() string {
	return proto.EnumName(BuildStatus_name, int32(x))
}

func (BuildStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_05a627abb7f9adb4, []int{0}
}

type ListBuildsRequest struct {
	EntityUri            string   `protobuf:"bytes,1,opt,name=entity_uri,json=entityUri,proto3" json:"entity_uri,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListBuildsRequest) Reset()         { *m = ListBuildsRequest{} }
func (m *ListBuildsRequest) String() string { return proto.CompactTextString(m) }
func (*ListBuildsRequest) ProtoMessage()    {}
func (*ListBuildsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_05a627abb7f9adb4, []int{0}
}

func (m *ListBuildsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListBuildsRequest.Unmarshal(m, b)
}
func (m *ListBuildsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListBuildsRequest.Marshal(b, m, deterministic)
}
func (m *ListBuildsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListBuildsRequest.Merge(m, src)
}
func (m *ListBuildsRequest) XXX_Size() int {
	return xxx_messageInfo_ListBuildsRequest.Size(m)
}
func (m *ListBuildsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListBuildsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListBuildsRequest proto.InternalMessageInfo

func (m *ListBuildsRequest) GetEntityUri() string {
	if m != nil {
		return m.EntityUri
	}
	return ""
}

type ListBuildsReply struct {
	EntityUri            string   `protobuf:"bytes,1,opt,name=entity_uri,json=entityUri,proto3" json:"entity_uri,omitempty"`
	Builds               []*Build `protobuf:"bytes,2,rep,name=builds,proto3" json:"builds,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListBuildsReply) Reset()         { *m = ListBuildsReply{} }
func (m *ListBuildsReply) String() string { return proto.CompactTextString(m) }
func (*ListBuildsReply) ProtoMessage()    {}
func (*ListBuildsReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_05a627abb7f9adb4, []int{1}
}

func (m *ListBuildsReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListBuildsReply.Unmarshal(m, b)
}
func (m *ListBuildsReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListBuildsReply.Marshal(b, m, deterministic)
}
func (m *ListBuildsReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListBuildsReply.Merge(m, src)
}
func (m *ListBuildsReply) XXX_Size() int {
	return xxx_messageInfo_ListBuildsReply.Size(m)
}
func (m *ListBuildsReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ListBuildsReply.DiscardUnknown(m)
}

var xxx_messageInfo_ListBuildsReply proto.InternalMessageInfo

func (m *ListBuildsReply) GetEntityUri() string {
	if m != nil {
		return m.EntityUri
	}
	return ""
}

func (m *ListBuildsReply) GetBuilds() []*Build {
	if m != nil {
		return m.Builds
	}
	return nil
}

type GetBuildRequest struct {
	BuildUri             string   `protobuf:"bytes,1,opt,name=build_uri,json=buildUri,proto3" json:"build_uri,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetBuildRequest) Reset()         { *m = GetBuildRequest{} }
func (m *GetBuildRequest) String() string { return proto.CompactTextString(m) }
func (*GetBuildRequest) ProtoMessage()    {}
func (*GetBuildRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_05a627abb7f9adb4, []int{2}
}

func (m *GetBuildRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBuildRequest.Unmarshal(m, b)
}
func (m *GetBuildRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBuildRequest.Marshal(b, m, deterministic)
}
func (m *GetBuildRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBuildRequest.Merge(m, src)
}
func (m *GetBuildRequest) XXX_Size() int {
	return xxx_messageInfo_GetBuildRequest.Size(m)
}
func (m *GetBuildRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBuildRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetBuildRequest proto.InternalMessageInfo

func (m *GetBuildRequest) GetBuildUri() string {
	if m != nil {
		return m.BuildUri
	}
	return ""
}

type GetBuildReply struct {
	Build                *Build        `protobuf:"bytes,1,opt,name=build,proto3" json:"build,omitempty"`
	Details              *BuildDetails `protobuf:"bytes,2,opt,name=details,proto3" json:"details,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *GetBuildReply) Reset()         { *m = GetBuildReply{} }
func (m *GetBuildReply) String() string { return proto.CompactTextString(m) }
func (*GetBuildReply) ProtoMessage()    {}
func (*GetBuildReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_05a627abb7f9adb4, []int{3}
}

func (m *GetBuildReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBuildReply.Unmarshal(m, b)
}
func (m *GetBuildReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBuildReply.Marshal(b, m, deterministic)
}
func (m *GetBuildReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBuildReply.Merge(m, src)
}
func (m *GetBuildReply) XXX_Size() int {
	return xxx_messageInfo_GetBuildReply.Size(m)
}
func (m *GetBuildReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBuildReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetBuildReply proto.InternalMessageInfo

func (m *GetBuildReply) GetBuild() *Build {
	if m != nil {
		return m.Build
	}
	return nil
}

func (m *GetBuildReply) GetDetails() *BuildDetails {
	if m != nil {
		return m.Details
	}
	return nil
}

type Build struct {
	Uri                  string      `protobuf:"bytes,1,opt,name=uri,proto3" json:"uri,omitempty"`
	CommitId             string      `protobuf:"bytes,2,opt,name=commit_id,json=commitId,proto3" json:"commit_id,omitempty"`
	Message              string      `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	Status               BuildStatus `protobuf:"varint,4,opt,name=status,proto3,enum=spotify.backstage.builds.v1.BuildStatus" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Build) Reset()         { *m = Build{} }
func (m *Build) String() string { return proto.CompactTextString(m) }
func (*Build) ProtoMessage()    {}
func (*Build) Descriptor() ([]byte, []int) {
	return fileDescriptor_05a627abb7f9adb4, []int{4}
}

func (m *Build) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Build.Unmarshal(m, b)
}
func (m *Build) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Build.Marshal(b, m, deterministic)
}
func (m *Build) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Build.Merge(m, src)
}
func (m *Build) XXX_Size() int {
	return xxx_messageInfo_Build.Size(m)
}
func (m *Build) XXX_DiscardUnknown() {
	xxx_messageInfo_Build.DiscardUnknown(m)
}

var xxx_messageInfo_Build proto.InternalMessageInfo

func (m *Build) GetUri() string {
	if m != nil {
		return m.Uri
	}
	return ""
}

func (m *Build) GetCommitId() string {
	if m != nil {
		return m.CommitId
	}
	return ""
}

func (m *Build) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Build) GetStatus() BuildStatus {
	if m != nil {
		return m.Status
	}
	return BuildStatus_NULL
}

type BuildDetails struct {
	Author               string   `protobuf:"bytes,1,opt,name=author,proto3" json:"author,omitempty"`
	OverviewUrl          string   `protobuf:"bytes,2,opt,name=overview_url,json=overviewUrl,proto3" json:"overview_url,omitempty"`
	LogUrl               string   `protobuf:"bytes,3,opt,name=log_url,json=logUrl,proto3" json:"log_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BuildDetails) Reset()         { *m = BuildDetails{} }
func (m *BuildDetails) String() string { return proto.CompactTextString(m) }
func (*BuildDetails) ProtoMessage()    {}
func (*BuildDetails) Descriptor() ([]byte, []int) {
	return fileDescriptor_05a627abb7f9adb4, []int{5}
}

func (m *BuildDetails) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BuildDetails.Unmarshal(m, b)
}
func (m *BuildDetails) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BuildDetails.Marshal(b, m, deterministic)
}
func (m *BuildDetails) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BuildDetails.Merge(m, src)
}
func (m *BuildDetails) XXX_Size() int {
	return xxx_messageInfo_BuildDetails.Size(m)
}
func (m *BuildDetails) XXX_DiscardUnknown() {
	xxx_messageInfo_BuildDetails.DiscardUnknown(m)
}

var xxx_messageInfo_BuildDetails proto.InternalMessageInfo

func (m *BuildDetails) GetAuthor() string {
	if m != nil {
		return m.Author
	}
	return ""
}

func (m *BuildDetails) GetOverviewUrl() string {
	if m != nil {
		return m.OverviewUrl
	}
	return ""
}

func (m *BuildDetails) GetLogUrl() string {
	if m != nil {
		return m.LogUrl
	}
	return ""
}

func init() {
	proto.RegisterEnum("spotify.backstage.builds.v1.BuildStatus", BuildStatus_name, BuildStatus_value)
	proto.RegisterType((*ListBuildsRequest)(nil), "spotify.backstage.builds.v1.ListBuildsRequest")
	proto.RegisterType((*ListBuildsReply)(nil), "spotify.backstage.builds.v1.ListBuildsReply")
	proto.RegisterType((*GetBuildRequest)(nil), "spotify.backstage.builds.v1.GetBuildRequest")
	proto.RegisterType((*GetBuildReply)(nil), "spotify.backstage.builds.v1.GetBuildReply")
	proto.RegisterType((*Build)(nil), "spotify.backstage.builds.v1.Build")
	proto.RegisterType((*BuildDetails)(nil), "spotify.backstage.builds.v1.BuildDetails")
}

func init() { proto.RegisterFile("builds/v1/builds.proto", fileDescriptor_05a627abb7f9adb4) }

var fileDescriptor_05a627abb7f9adb4 = []byte{
	// 447 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0x41, 0x6f, 0xd3, 0x30,
	0x18, 0x25, 0x6b, 0x97, 0xb6, 0x5f, 0x07, 0x0b, 0x3e, 0x8c, 0x68, 0x13, 0x52, 0xc9, 0xa9, 0x4c,
	0x28, 0x53, 0xcb, 0x05, 0x71, 0x82, 0x75, 0x65, 0xaa, 0xa8, 0x22, 0xe4, 0x2a, 0x17, 0x2e, 0x55,
	0xda, 0x98, 0x62, 0x70, 0x71, 0xb0, 0x9d, 0xa0, 0xfc, 0x09, 0x0e, 0xfc, 0x3c, 0x7e, 0x0d, 0xb2,
	0x9d, 0xa8, 0x11, 0x93, 0xda, 0xde, 0xfc, 0xbe, 0xf7, 0xbd, 0xbc, 0xf7, 0x49, 0x2f, 0x70, 0xb1,
	0xca, 0x29, 0x4b, 0xe5, 0x4d, 0x31, 0xba, 0xb1, 0xaf, 0x30, 0x13, 0x5c, 0x71, 0x74, 0x25, 0x33,
	0xae, 0xe8, 0x97, 0x32, 0x5c, 0x25, 0xeb, 0xef, 0x52, 0x25, 0x1b, 0x12, 0x56, 0x7c, 0x31, 0x0a,
	0xc6, 0xf0, 0x74, 0x4e, 0xa5, 0xba, 0x35, 0x03, 0x4c, 0x7e, 0xe6, 0x44, 0x2a, 0xf4, 0x1c, 0x80,
	0xfc, 0x50, 0x54, 0x95, 0xcb, 0x5c, 0x50, 0xdf, 0x19, 0x38, 0xc3, 0x1e, 0xee, 0xd9, 0x49, 0x2c,
	0x68, 0xc0, 0xe0, 0xbc, 0xa9, 0xc9, 0x58, 0x79, 0x40, 0x81, 0xde, 0x82, 0x6b, 0x2d, 0xfd, 0x93,
	0x41, 0x6b, 0xd8, 0x1f, 0x07, 0xe1, 0x9e, 0x4c, 0xa1, 0xf9, 0x30, 0xae, 0x14, 0x41, 0x08, 0xe7,
	0xf7, 0xc4, 0x9a, 0xd5, 0xf9, 0xae, 0xa0, 0x67, 0xc8, 0x86, 0x59, 0xd7, 0x0c, 0x74, 0xba, 0xdf,
	0x0e, 0x3c, 0xde, 0x09, 0x74, 0xb8, 0x37, 0x70, 0x6a, 0x58, 0xb3, 0x7a, 0x9c, 0xb9, 0x15, 0xa0,
	0x09, 0x74, 0x52, 0xa2, 0x12, 0xca, 0x74, 0x70, 0xad, 0x7d, 0x79, 0x58, 0x7b, 0x67, 0x05, 0xb8,
	0x56, 0x06, 0x7f, 0x1c, 0x38, 0x35, 0x0c, 0xf2, 0xa0, 0xb5, 0x4b, 0xac, 0x9f, 0xfa, 0x92, 0x35,
	0xdf, 0x6e, 0xa9, 0x5a, 0xd2, 0xd4, 0x58, 0xf4, 0x70, 0xd7, 0x0e, 0x66, 0x29, 0xf2, 0xa1, 0xb3,
	0x25, 0x52, 0x26, 0x1b, 0xe2, 0xb7, 0x0c, 0x55, 0x43, 0xf4, 0x0e, 0x5c, 0xa9, 0x12, 0x95, 0x4b,
	0xbf, 0x3d, 0x70, 0x86, 0x4f, 0xc6, 0xc3, 0xc3, 0xb1, 0x16, 0x66, 0x1f, 0x57, 0xba, 0x60, 0x05,
	0x67, 0xcd, 0xb4, 0xe8, 0x02, 0xdc, 0x24, 0x57, 0x5f, 0xb9, 0xa8, 0xd2, 0x55, 0x08, 0xbd, 0x80,
	0x33, 0x5e, 0x10, 0x51, 0x50, 0xf2, 0x6b, 0x99, 0x0b, 0x56, 0x65, 0xec, 0xd7, 0xb3, 0x58, 0x30,
	0xf4, 0x0c, 0x3a, 0x8c, 0x6f, 0x0c, 0x6b, 0x63, 0xba, 0x8c, 0x6f, 0x62, 0xc1, 0xae, 0x3f, 0x42,
	0xbf, 0x61, 0x8d, 0xba, 0xd0, 0x8e, 0xe2, 0xf9, 0xdc, 0x7b, 0x84, 0xfa, 0xd0, 0x59, 0xc4, 0x93,
	0xc9, 0x74, 0xb1, 0xf0, 0x1c, 0x0d, 0x3e, 0xbc, 0x9f, 0xcd, 0x63, 0x3c, 0xf5, 0x4e, 0x34, 0xf8,
	0x34, 0x8d, 0xee, 0x66, 0xd1, 0xbd, 0xd7, 0xd2, 0x00, 0xc7, 0x51, 0xa4, 0x41, 0x7b, 0xfc, 0xd7,
	0x01, 0xd7, 0x36, 0x0e, 0x7d, 0x03, 0xd8, 0xf5, 0x0f, 0x85, 0x7b, 0x6f, 0x7f, 0x50, 0xee, 0xcb,
	0x57, 0x47, 0xef, 0xeb, 0xee, 0xa4, 0xd0, 0xad, 0xcb, 0x84, 0xf6, 0x2b, 0xff, 0x2b, 0xe9, 0xe5,
	0xf5, 0x91, 0xdb, 0x19, 0x2b, 0x6f, 0xe1, 0xb3, 0xed, 0xaf, 0x2c, 0x46, 0x2b, 0xd7, 0xfc, 0xb5,
	0xaf, 0xff, 0x05, 0x00, 0x00, 0xff, 0xff, 0x40, 0xf3, 0x96, 0x0b, 0xcf, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// BuildsClient is the client API for Builds service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BuildsClient interface {
	ListBuilds(ctx context.Context, in *ListBuildsRequest, opts ...grpc.CallOption) (*ListBuildsReply, error)
	GetBuild(ctx context.Context, in *GetBuildRequest, opts ...grpc.CallOption) (*GetBuildReply, error)
}

type buildsClient struct {
	cc grpc.ClientConnInterface
}

func NewBuildsClient(cc grpc.ClientConnInterface) BuildsClient {
	return &buildsClient{cc}
}

func (c *buildsClient) ListBuilds(ctx context.Context, in *ListBuildsRequest, opts ...grpc.CallOption) (*ListBuildsReply, error) {
	out := new(ListBuildsReply)
	err := c.cc.Invoke(ctx, "/spotify.backstage.builds.v1.Builds/ListBuilds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *buildsClient) GetBuild(ctx context.Context, in *GetBuildRequest, opts ...grpc.CallOption) (*GetBuildReply, error) {
	out := new(GetBuildReply)
	err := c.cc.Invoke(ctx, "/spotify.backstage.builds.v1.Builds/GetBuild", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BuildsServer is the server API for Builds service.
type BuildsServer interface {
	ListBuilds(context.Context, *ListBuildsRequest) (*ListBuildsReply, error)
	GetBuild(context.Context, *GetBuildRequest) (*GetBuildReply, error)
}

// UnimplementedBuildsServer can be embedded to have forward compatible implementations.
type UnimplementedBuildsServer struct {
}

func (*UnimplementedBuildsServer) ListBuilds(ctx context.Context, req *ListBuildsRequest) (*ListBuildsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListBuilds not implemented")
}
func (*UnimplementedBuildsServer) GetBuild(ctx context.Context, req *GetBuildRequest) (*GetBuildReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBuild not implemented")
}

func RegisterBuildsServer(s *grpc.Server, srv BuildsServer) {
	s.RegisterService(&_Builds_serviceDesc, srv)
}

func _Builds_ListBuilds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListBuildsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BuildsServer).ListBuilds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spotify.backstage.builds.v1.Builds/ListBuilds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BuildsServer).ListBuilds(ctx, req.(*ListBuildsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Builds_GetBuild_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBuildRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BuildsServer).GetBuild(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spotify.backstage.builds.v1.Builds/GetBuild",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BuildsServer).GetBuild(ctx, req.(*GetBuildRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Builds_serviceDesc = grpc.ServiceDesc{
	ServiceName: "spotify.backstage.builds.v1.Builds",
	HandlerType: (*BuildsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListBuilds",
			Handler:    _Builds_ListBuilds_Handler,
		},
		{
			MethodName: "GetBuild",
			Handler:    _Builds_GetBuild_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "builds/v1/builds.proto",
}
