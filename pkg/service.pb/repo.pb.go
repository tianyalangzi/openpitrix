// Code generated by protoc-gen-go. DO NOT EDIT.
// source: repo.proto

package openpitrix

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import _ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
import google_protobuf2 "github.com/golang/protobuf/ptypes/empty"
import google_protobuf3 "github.com/golang/protobuf/ptypes/timestamp"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Repo struct {
	Id               *string                     `protobuf:"bytes,1,req,name=id" json:"id,omitempty"`
	Name             *string                     `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Description      *string                     `protobuf:"bytes,3,opt,name=description" json:"description,omitempty"`
	Url              *string                     `protobuf:"bytes,4,opt,name=url" json:"url,omitempty"`
	Created          *google_protobuf3.Timestamp `protobuf:"bytes,5,opt,name=created" json:"created,omitempty"`
	LastModified     *google_protobuf3.Timestamp `protobuf:"bytes,6,opt,name=last_modified,json=lastModified" json:"last_modified,omitempty"`
	XXX_unrecognized []byte                      `json:"-"`
}

func (m *Repo) Reset()                    { *m = Repo{} }
func (m *Repo) String() string            { return proto.CompactTextString(m) }
func (*Repo) ProtoMessage()               {}
func (*Repo) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *Repo) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

func (m *Repo) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Repo) GetDescription() string {
	if m != nil && m.Description != nil {
		return *m.Description
	}
	return ""
}

func (m *Repo) GetUrl() string {
	if m != nil && m.Url != nil {
		return *m.Url
	}
	return ""
}

func (m *Repo) GetCreated() *google_protobuf3.Timestamp {
	if m != nil {
		return m.Created
	}
	return nil
}

func (m *Repo) GetLastModified() *google_protobuf3.Timestamp {
	if m != nil {
		return m.LastModified
	}
	return nil
}

type RepoLabel struct {
	RepoId           *string `protobuf:"bytes,1,req,name=repo_id,json=repoId" json:"repo_id,omitempty"`
	LabelKey         *string `protobuf:"bytes,2,req,name=label_key,json=labelKey" json:"label_key,omitempty"`
	LabelValue       *string `protobuf:"bytes,3,req,name=label_value,json=labelValue" json:"label_value,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *RepoLabel) Reset()                    { *m = RepoLabel{} }
func (m *RepoLabel) String() string            { return proto.CompactTextString(m) }
func (*RepoLabel) ProtoMessage()               {}
func (*RepoLabel) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

func (m *RepoLabel) GetRepoId() string {
	if m != nil && m.RepoId != nil {
		return *m.RepoId
	}
	return ""
}

func (m *RepoLabel) GetLabelKey() string {
	if m != nil && m.LabelKey != nil {
		return *m.LabelKey
	}
	return ""
}

func (m *RepoLabel) GetLabelValue() string {
	if m != nil && m.LabelValue != nil {
		return *m.LabelValue
	}
	return ""
}

type RepoSelector struct {
	RepoId           *string `protobuf:"bytes,1,req,name=repo_id,json=repoId" json:"repo_id,omitempty"`
	SelectorKey      *string `protobuf:"bytes,2,req,name=selector_key,json=selectorKey" json:"selector_key,omitempty"`
	SelectorValue    *string `protobuf:"bytes,3,req,name=selector_value,json=selectorValue" json:"selector_value,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *RepoSelector) Reset()                    { *m = RepoSelector{} }
func (m *RepoSelector) String() string            { return proto.CompactTextString(m) }
func (*RepoSelector) ProtoMessage()               {}
func (*RepoSelector) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{2} }

func (m *RepoSelector) GetRepoId() string {
	if m != nil && m.RepoId != nil {
		return *m.RepoId
	}
	return ""
}

func (m *RepoSelector) GetSelectorKey() string {
	if m != nil && m.SelectorKey != nil {
		return *m.SelectorKey
	}
	return ""
}

func (m *RepoSelector) GetSelectorValue() string {
	if m != nil && m.SelectorValue != nil {
		return *m.SelectorValue
	}
	return ""
}

type RepoId struct {
	Id               *string `protobuf:"bytes,1,req,name=id" json:"id,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *RepoId) Reset()                    { *m = RepoId{} }
func (m *RepoId) String() string            { return proto.CompactTextString(m) }
func (*RepoId) ProtoMessage()               {}
func (*RepoId) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{3} }

func (m *RepoId) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

type RepoListRequest struct {
	PageSize         *int32 `protobuf:"varint,1,opt,name=page_size,json=pageSize,def=10" json:"page_size,omitempty"`
	PageNumber       *int32 `protobuf:"varint,2,opt,name=page_number,json=pageNumber,def=1" json:"page_number,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *RepoListRequest) Reset()                    { *m = RepoListRequest{} }
func (m *RepoListRequest) String() string            { return proto.CompactTextString(m) }
func (*RepoListRequest) ProtoMessage()               {}
func (*RepoListRequest) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{4} }

const Default_RepoListRequest_PageSize int32 = 10
const Default_RepoListRequest_PageNumber int32 = 1

func (m *RepoListRequest) GetPageSize() int32 {
	if m != nil && m.PageSize != nil {
		return *m.PageSize
	}
	return Default_RepoListRequest_PageSize
}

func (m *RepoListRequest) GetPageNumber() int32 {
	if m != nil && m.PageNumber != nil {
		return *m.PageNumber
	}
	return Default_RepoListRequest_PageNumber
}

type RepoListResponse struct {
	TotalItems       *int32  `protobuf:"varint,1,opt,name=total_items,json=totalItems" json:"total_items,omitempty"`
	TotalPages       *int32  `protobuf:"varint,2,opt,name=total_pages,json=totalPages" json:"total_pages,omitempty"`
	PageSize         *int32  `protobuf:"varint,3,opt,name=page_size,json=pageSize" json:"page_size,omitempty"`
	CurrentPage      *int32  `protobuf:"varint,4,opt,name=current_page,json=currentPage" json:"current_page,omitempty"`
	Items            []*Repo `protobuf:"bytes,5,rep,name=items" json:"items,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *RepoListResponse) Reset()                    { *m = RepoListResponse{} }
func (m *RepoListResponse) String() string            { return proto.CompactTextString(m) }
func (*RepoListResponse) ProtoMessage()               {}
func (*RepoListResponse) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{5} }

func (m *RepoListResponse) GetTotalItems() int32 {
	if m != nil && m.TotalItems != nil {
		return *m.TotalItems
	}
	return 0
}

func (m *RepoListResponse) GetTotalPages() int32 {
	if m != nil && m.TotalPages != nil {
		return *m.TotalPages
	}
	return 0
}

func (m *RepoListResponse) GetPageSize() int32 {
	if m != nil && m.PageSize != nil {
		return *m.PageSize
	}
	return 0
}

func (m *RepoListResponse) GetCurrentPage() int32 {
	if m != nil && m.CurrentPage != nil {
		return *m.CurrentPage
	}
	return 0
}

func (m *RepoListResponse) GetItems() []*Repo {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
	proto.RegisterType((*Repo)(nil), "openpitrix.Repo")
	proto.RegisterType((*RepoLabel)(nil), "openpitrix.RepoLabel")
	proto.RegisterType((*RepoSelector)(nil), "openpitrix.RepoSelector")
	proto.RegisterType((*RepoId)(nil), "openpitrix.RepoId")
	proto.RegisterType((*RepoListRequest)(nil), "openpitrix.RepoListRequest")
	proto.RegisterType((*RepoListResponse)(nil), "openpitrix.RepoListResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for RepoService service

type RepoServiceClient interface {
	GetRepo(ctx context.Context, in *RepoId, opts ...grpc.CallOption) (*Repo, error)
	GetRepoList(ctx context.Context, in *RepoListRequest, opts ...grpc.CallOption) (*RepoListResponse, error)
	CreateRepo(ctx context.Context, in *Repo, opts ...grpc.CallOption) (*google_protobuf2.Empty, error)
	UpdateRepo(ctx context.Context, in *Repo, opts ...grpc.CallOption) (*google_protobuf2.Empty, error)
	DeleteRepo(ctx context.Context, in *RepoId, opts ...grpc.CallOption) (*google_protobuf2.Empty, error)
}

type repoServiceClient struct {
	cc *grpc.ClientConn
}

func NewRepoServiceClient(cc *grpc.ClientConn) RepoServiceClient {
	return &repoServiceClient{cc}
}

func (c *repoServiceClient) GetRepo(ctx context.Context, in *RepoId, opts ...grpc.CallOption) (*Repo, error) {
	out := new(Repo)
	err := grpc.Invoke(ctx, "/openpitrix.RepoService/GetRepo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *repoServiceClient) GetRepoList(ctx context.Context, in *RepoListRequest, opts ...grpc.CallOption) (*RepoListResponse, error) {
	out := new(RepoListResponse)
	err := grpc.Invoke(ctx, "/openpitrix.RepoService/GetRepoList", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *repoServiceClient) CreateRepo(ctx context.Context, in *Repo, opts ...grpc.CallOption) (*google_protobuf2.Empty, error) {
	out := new(google_protobuf2.Empty)
	err := grpc.Invoke(ctx, "/openpitrix.RepoService/CreateRepo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *repoServiceClient) UpdateRepo(ctx context.Context, in *Repo, opts ...grpc.CallOption) (*google_protobuf2.Empty, error) {
	out := new(google_protobuf2.Empty)
	err := grpc.Invoke(ctx, "/openpitrix.RepoService/UpdateRepo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *repoServiceClient) DeleteRepo(ctx context.Context, in *RepoId, opts ...grpc.CallOption) (*google_protobuf2.Empty, error) {
	out := new(google_protobuf2.Empty)
	err := grpc.Invoke(ctx, "/openpitrix.RepoService/DeleteRepo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RepoService service

type RepoServiceServer interface {
	GetRepo(context.Context, *RepoId) (*Repo, error)
	GetRepoList(context.Context, *RepoListRequest) (*RepoListResponse, error)
	CreateRepo(context.Context, *Repo) (*google_protobuf2.Empty, error)
	UpdateRepo(context.Context, *Repo) (*google_protobuf2.Empty, error)
	DeleteRepo(context.Context, *RepoId) (*google_protobuf2.Empty, error)
}

func RegisterRepoServiceServer(s *grpc.Server, srv RepoServiceServer) {
	s.RegisterService(&_RepoService_serviceDesc, srv)
}

func _RepoService_GetRepo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RepoId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepoServiceServer).GetRepo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.RepoService/GetRepo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepoServiceServer).GetRepo(ctx, req.(*RepoId))
	}
	return interceptor(ctx, in, info, handler)
}

func _RepoService_GetRepoList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RepoListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepoServiceServer).GetRepoList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.RepoService/GetRepoList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepoServiceServer).GetRepoList(ctx, req.(*RepoListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RepoService_CreateRepo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Repo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepoServiceServer).CreateRepo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.RepoService/CreateRepo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepoServiceServer).CreateRepo(ctx, req.(*Repo))
	}
	return interceptor(ctx, in, info, handler)
}

func _RepoService_UpdateRepo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Repo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepoServiceServer).UpdateRepo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.RepoService/UpdateRepo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepoServiceServer).UpdateRepo(ctx, req.(*Repo))
	}
	return interceptor(ctx, in, info, handler)
}

func _RepoService_DeleteRepo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RepoId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepoServiceServer).DeleteRepo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.RepoService/DeleteRepo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepoServiceServer).DeleteRepo(ctx, req.(*RepoId))
	}
	return interceptor(ctx, in, info, handler)
}

var _RepoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "openpitrix.RepoService",
	HandlerType: (*RepoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRepo",
			Handler:    _RepoService_GetRepo_Handler,
		},
		{
			MethodName: "GetRepoList",
			Handler:    _RepoService_GetRepoList_Handler,
		},
		{
			MethodName: "CreateRepo",
			Handler:    _RepoService_CreateRepo_Handler,
		},
		{
			MethodName: "UpdateRepo",
			Handler:    _RepoService_UpdateRepo_Handler,
		},
		{
			MethodName: "DeleteRepo",
			Handler:    _RepoService_DeleteRepo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "repo.proto",
}

func init() { proto.RegisterFile("repo.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 802 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xd1, 0x6e, 0x1b, 0x45,
	0x14, 0xd5, 0xda, 0x75, 0x52, 0xdf, 0x4d, 0x83, 0x19, 0xa0, 0x5d, 0x9c, 0x4a, 0x99, 0xae, 0x04,
	0xb2, 0x42, 0xba, 0x9b, 0x58, 0x3c, 0xf9, 0x05, 0x85, 0x82, 0xaa, 0xa8, 0x05, 0xd2, 0x0d, 0x94,
	0x07, 0x1e, 0xac, 0xf1, 0xee, 0x8d, 0x3b, 0xed, 0xee, 0xcc, 0x74, 0x66, 0x36, 0x34, 0x45, 0xbc,
	0xf0, 0x09, 0xf0, 0x0f, 0x7c, 0x05, 0x6f, 0x7c, 0x02, 0xe2, 0x0f, 0x78, 0xe5, 0x81, 0x3f, 0x40,
	0x33, 0xbb, 0x8e, 0x2d, 0x27, 0x15, 0xe2, 0xc9, 0xbb, 0xe7, 0x9e, 0x39, 0xe7, 0xdc, 0x7b, 0x77,
	0x0c, 0xa0, 0x51, 0xc9, 0x44, 0x69, 0x69, 0x25, 0x01, 0xa9, 0x50, 0x28, 0x6e, 0x35, 0x7f, 0x35,
	0xbc, 0x3b, 0x97, 0x72, 0x5e, 0x62, 0xca, 0x14, 0x4f, 0x99, 0x10, 0xd2, 0x32, 0xcb, 0xa5, 0x30,
	0x0d, 0x73, 0xb8, 0xef, 0x7f, 0xf2, 0xfb, 0x73, 0x14, 0xf7, 0xcd, 0xf7, 0x6c, 0x3e, 0x47, 0x9d,
	0x4a, 0xe5, 0x19, 0xd7, 0xb0, 0x77, 0x5a, 0x2d, 0xff, 0x36, 0xab, 0xcf, 0x52, 0xac, 0x94, 0xbd,
	0x68, 0x8b, 0xbb, 0xeb, 0x45, 0xcb, 0x2b, 0x34, 0x96, 0x55, 0xaa, 0x21, 0xc4, 0x7f, 0x06, 0x70,
	0x23, 0x43, 0x25, 0xc9, 0x36, 0x74, 0x78, 0x11, 0x05, 0xb4, 0x33, 0xea, 0x67, 0x1d, 0x5e, 0x10,
	0x02, 0x37, 0x04, 0xab, 0x30, 0xea, 0xd0, 0x60, 0xd4, 0xcf, 0xfc, 0x33, 0xa1, 0x10, 0x16, 0x68,
	0x72, 0xcd, 0x7d, 0x98, 0xa8, 0xeb, 0x4b, 0xab, 0x10, 0x19, 0x40, 0xb7, 0xd6, 0x65, 0x74, 0xc3,
	0x57, 0xdc, 0x23, 0xf9, 0x18, 0x36, 0x73, 0x8d, 0xcc, 0x62, 0x11, 0xf5, 0x68, 0x30, 0x0a, 0xc7,
	0xc3, 0xa4, 0xc9, 0x94, 0x2c, 0x32, 0x25, 0x5f, 0x2f, 0x32, 0x65, 0x0b, 0x2a, 0xf9, 0x04, 0x6e,
	0x95, 0xcc, 0xd8, 0x69, 0x25, 0x0b, 0x7e, 0xc6, 0xb1, 0x88, 0x36, 0xfe, 0xf3, 0xec, 0x96, 0x3b,
	0xf0, 0x45, 0xcb, 0x8f, 0x67, 0xd0, 0x77, 0x6d, 0x3d, 0x66, 0x33, 0x2c, 0xc9, 0x1d, 0xd8, 0x74,
	0x8b, 0x98, 0x5e, 0x36, 0xb8, 0xe1, 0x5e, 0x8f, 0x0b, 0xb2, 0x03, 0xfd, 0xd2, 0x31, 0xa6, 0x2f,
	0xf0, 0x22, 0xea, 0xf8, 0xd2, 0x4d, 0x0f, 0x3c, 0xc2, 0x0b, 0xb2, 0x0b, 0x61, 0x53, 0x3c, 0x67,
	0x65, 0x8d, 0x51, 0xd7, 0x97, 0xc1, 0x43, 0x4f, 0x1d, 0x12, 0xbf, 0x84, 0x2d, 0xe7, 0x71, 0x8a,
	0x25, 0xe6, 0x56, 0xea, 0x37, 0xdb, 0xdc, 0x83, 0x2d, 0xd3, 0x92, 0x56, 0x9c, 0xc2, 0x05, 0xe6,
	0xcc, 0x3e, 0x80, 0xed, 0x4b, 0xca, 0xaa, 0xdf, 0xad, 0x05, 0xda, 0x58, 0x46, 0xb0, 0x91, 0x35,
	0x9a, 0x6b, 0xfb, 0x8a, 0x9f, 0xc2, 0x5b, 0xbe, 0x61, 0x6e, 0x6c, 0x86, 0x2f, 0x6b, 0x34, 0x96,
	0xec, 0x42, 0x5f, 0xb1, 0x39, 0x4e, 0x0d, 0x7f, 0x8d, 0x51, 0x40, 0x83, 0x51, 0x6f, 0xd2, 0x39,
	0x3c, 0xc8, 0x6e, 0x3a, 0xf0, 0x94, 0xbf, 0x46, 0x12, 0x43, 0xe8, 0x09, 0xa2, 0xae, 0x66, 0xa8,
	0xfd, 0xaa, 0x7b, 0x93, 0xe0, 0x30, 0x03, 0x87, 0x7e, 0xe9, 0xc1, 0xf8, 0xb7, 0x00, 0x06, 0x4b,
	0x61, 0xa3, 0xa4, 0x30, 0xe8, 0x46, 0x63, 0xa5, 0x65, 0xe5, 0x94, 0x5b, 0xac, 0x4c, 0xa3, 0x9d,
	0x81, 0x87, 0x8e, 0x1d, 0xb2, 0x24, 0x38, 0x25, 0xd3, 0x28, 0xb7, 0x84, 0x13, 0x87, 0xb8, 0xc9,
	0x2f, 0xb3, 0x75, 0x7d, 0x79, 0x99, 0xeb, 0x1e, 0x6c, 0xe5, 0xb5, 0xd6, 0x28, 0xac, 0x3f, 0xef,
	0x3f, 0xa7, 0x5e, 0x16, 0xb6, 0x98, 0x13, 0x20, 0x1f, 0x42, 0xaf, 0xf1, 0xee, 0xd1, 0xee, 0x28,
	0x1c, 0x0f, 0x92, 0xe5, 0xed, 0x4a, 0x5c, 0xdc, 0xac, 0x29, 0x8f, 0x7f, 0xed, 0x42, 0xd8, 0x2c,
	0x49, 0x9f, 0xf3, 0x1c, 0xc9, 0x43, 0xd8, 0x7c, 0x88, 0xd6, 0x7f, 0xf1, 0x64, 0xfd, 0xcc, 0x71,
	0x31, 0xbc, 0xa2, 0x13, 0xdf, 0xfe, 0xe9, 0x8f, 0xbf, 0x7e, 0xe9, 0x0c, 0xc8, 0x76, 0x7a, 0x7e,
	0x98, 0xba, 0x7d, 0x9a, 0xf4, 0x07, 0x5e, 0xfc, 0x48, 0xbe, 0x83, 0xb0, 0x15, 0x72, 0x93, 0x21,
	0x3b, 0xeb, 0x07, 0x57, 0x16, 0x31, 0xbc, 0x7b, 0x7d, 0xb1, 0x19, 0x66, 0xfc, 0xb6, 0x77, 0x08,
	0x49, 0xff, 0xd2, 0x81, 0x3c, 0x06, 0x78, 0xe0, 0x6f, 0x82, 0x0f, 0x7a, 0x25, 0xd4, 0xf0, 0xf6,
	0x95, 0x7b, 0xf0, 0xb9, 0xbb, 0xf4, 0xf1, 0xbb, 0x5e, 0x6a, 0x3b, 0x5e, 0x4a, 0x4d, 0x82, 0x3d,
	0xf2, 0x04, 0xe0, 0x1b, 0x55, 0xfc, 0x7f, 0xb5, 0xf7, 0xbd, 0xda, 0x3b, 0xf1, 0x5a, 0xeb, 0x4e,
	0xf2, 0x04, 0xe0, 0x33, 0x2c, 0xb1, 0x95, 0xbc, 0x6e, 0x92, 0x6f, 0x12, 0x6d, 0xe7, 0xb9, 0xb7,
	0x26, 0xfa, 0xe9, 0x3f, 0xc1, 0xcf, 0x47, 0x7f, 0x07, 0xe4, 0xf7, 0x00, 0xee, 0x7c, 0xa5, 0x50,
	0x9c, 0x78, 0x31, 0xea, 0xc4, 0xe8, 0x62, 0x77, 0xf6, 0x48, 0x50, 0x67, 0x44, 0x55, 0xc9, 0xec,
	0x99, 0xd4, 0x15, 0xb5, 0x92, 0x2a, 0x96, 0xbf, 0x60, 0x73, 0xa4, 0x4c, 0x14, 0xb4, 0x40, 0x55,
	0xca, 0x0b, 0xca, 0x94, 0x2a, 0x79, 0xde, 0xfc, 0x49, 0x52, 0x2e, 0xac, 0xa4, 0x55, 0x5d, 0x5a,
	0xae, 0x4a, 0x44, 0x71, 0xce, 0xb5, 0x14, 0x15, 0x0a, 0x4b, 0x4d, 0x9d, 0x3f, 0xa3, 0xcc, 0xd0,
	0x27, 0x5c, 0xcc, 0x1f, 0x94, 0xb2, 0x2e, 0xf6, 0xe9, 0xd1, 0xb7, 0xa7, 0xfb, 0xf4, 0x51, 0x3d,
	0x43, 0x2d, 0xd0, 0xa2, 0xa1, 0x68, 0xf3, 0x24, 0xfe, 0x08, 0x48, 0x13, 0x48, 0x23, 0x7f, 0x45,
	0x4f, 0xb4, 0x7c, 0x8e, 0xb9, 0x25, 0xef, 0x3d, 0xb3, 0x56, 0x99, 0x49, 0x9a, 0xae, 0x74, 0xce,
	0xe5, 0xb8, 0x77, 0x90, 0x1c, 0x24, 0x87, 0x7b, 0x41, 0x30, 0x1e, 0xac, 0xa4, 0x48, 0x9f, 0x1b,
	0x29, 0x26, 0x57, 0x90, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x3d, 0xd7, 0x22, 0xe7, 0x1f, 0x06,
	0x00, 0x00,
}
