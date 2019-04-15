// Code generated by protoc-gen-go. DO NOT EDIT.
// source: yandex/cloud/containerregistry/v1/image_service.proto

package containerregistry // import "github.com/yandex-cloud/go-genproto/yandex/cloud/containerregistry/v1"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/yandex-cloud/go-genproto/yandex/api"
import operation "github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
import _ "github.com/yandex-cloud/go-genproto/yandex/cloud/validation"
import _ "google.golang.org/genproto/googleapis/api/annotations"

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

type ListImagesRequest struct {
	// ID of the registry to list Docker images in.
	//
	// [registry_id] is ignored if a [ListImagesRequest.repository_name] is specified in the request.
	//
	// To get the registry ID use a [RegistryService.List] request.
	RegistryId string `protobuf:"bytes,1,opt,name=registry_id,json=registryId,proto3" json:"registry_id,omitempty"`
	// Name of the repository to list Docker images in.
	//
	// To get the repository name use a [RepositoryService.List] request.
	RepositoryName string `protobuf:"bytes,2,opt,name=repository_name,json=repositoryName,proto3" json:"repository_name,omitempty"`
	// ID of the folder to list Docker images in.
	//
	// [folder_id] is ignored if a [ListImagesRequest.repository_name] or a [ListImagesRequest.registry_id] are specified in the request.
	//
	// To get the folder ID use a [yandex.cloud.resourcemanager.v1.FolderService.List] request.
	FolderId string `protobuf:"bytes,7,opt,name=folder_id,json=folderId,proto3" json:"folder_id,omitempty"`
	// The maximum number of results per page to return. If the number of available
	// results is larger than [page_size],
	// the service returns a [ListImagesResponse.next_page_token]
	// that can be used to get the next page of results in subsequent list requests.
	// Default value: 100.
	PageSize int64 `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Page token. To get the next page of results, set [page_token] to the
	// [ListImagesResponse.next_page_token] returned by a previous list request.
	PageToken string `protobuf:"bytes,4,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	// A filter expression that filters resources listed in the response.
	// The expression must specify:
	// 1. The field name. Currently you can use filtering only on [Image.name] field.
	// 2. An operator. Can be either `=` or `!=` for single values, `IN` or `NOT IN` for lists of values.
	// 3. The value. Must be a maximum of 256 characters and match the regular expression `[a-z0-9]+(?:[._-][a-z0-9]+)*(/([a-z0-9]+(?:[._-][a-z0-9]+)*))`.
	Filter               string   `protobuf:"bytes,5,opt,name=filter,proto3" json:"filter,omitempty"`
	OrderBy              string   `protobuf:"bytes,6,opt,name=order_by,json=orderBy,proto3" json:"order_by,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListImagesRequest) Reset()         { *m = ListImagesRequest{} }
func (m *ListImagesRequest) String() string { return proto.CompactTextString(m) }
func (*ListImagesRequest) ProtoMessage()    {}
func (*ListImagesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_image_service_181142b342f3b95f, []int{0}
}
func (m *ListImagesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListImagesRequest.Unmarshal(m, b)
}
func (m *ListImagesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListImagesRequest.Marshal(b, m, deterministic)
}
func (dst *ListImagesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListImagesRequest.Merge(dst, src)
}
func (m *ListImagesRequest) XXX_Size() int {
	return xxx_messageInfo_ListImagesRequest.Size(m)
}
func (m *ListImagesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListImagesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListImagesRequest proto.InternalMessageInfo

func (m *ListImagesRequest) GetRegistryId() string {
	if m != nil {
		return m.RegistryId
	}
	return ""
}

func (m *ListImagesRequest) GetRepositoryName() string {
	if m != nil {
		return m.RepositoryName
	}
	return ""
}

func (m *ListImagesRequest) GetFolderId() string {
	if m != nil {
		return m.FolderId
	}
	return ""
}

func (m *ListImagesRequest) GetPageSize() int64 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListImagesRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

func (m *ListImagesRequest) GetFilter() string {
	if m != nil {
		return m.Filter
	}
	return ""
}

func (m *ListImagesRequest) GetOrderBy() string {
	if m != nil {
		return m.OrderBy
	}
	return ""
}

type ListImagesResponse struct {
	// List of Image resources.
	Images []*Image `protobuf:"bytes,1,rep,name=images,proto3" json:"images,omitempty"`
	// This token allows you to get the next page of results for list requests. If the number of results
	// is larger than [ListImagesRequest.page_size], use
	// the [next_page_token] as the value
	// for the [ListImagesRequest.page_token] query parameter
	// in the next list request. Each subsequent list request will have its own
	// [next_page_token] to continue paging through the results.
	NextPageToken        string   `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListImagesResponse) Reset()         { *m = ListImagesResponse{} }
func (m *ListImagesResponse) String() string { return proto.CompactTextString(m) }
func (*ListImagesResponse) ProtoMessage()    {}
func (*ListImagesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_image_service_181142b342f3b95f, []int{1}
}
func (m *ListImagesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListImagesResponse.Unmarshal(m, b)
}
func (m *ListImagesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListImagesResponse.Marshal(b, m, deterministic)
}
func (dst *ListImagesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListImagesResponse.Merge(dst, src)
}
func (m *ListImagesResponse) XXX_Size() int {
	return xxx_messageInfo_ListImagesResponse.Size(m)
}
func (m *ListImagesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListImagesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListImagesResponse proto.InternalMessageInfo

func (m *ListImagesResponse) GetImages() []*Image {
	if m != nil {
		return m.Images
	}
	return nil
}

func (m *ListImagesResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

type GetImageRequest struct {
	// ID of the Docker image resource to return.
	//
	// To get the Docker image ID use a [ImageService.List] request.
	ImageId              string   `protobuf:"bytes,1,opt,name=image_id,json=imageId,proto3" json:"image_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetImageRequest) Reset()         { *m = GetImageRequest{} }
func (m *GetImageRequest) String() string { return proto.CompactTextString(m) }
func (*GetImageRequest) ProtoMessage()    {}
func (*GetImageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_image_service_181142b342f3b95f, []int{2}
}
func (m *GetImageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetImageRequest.Unmarshal(m, b)
}
func (m *GetImageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetImageRequest.Marshal(b, m, deterministic)
}
func (dst *GetImageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetImageRequest.Merge(dst, src)
}
func (m *GetImageRequest) XXX_Size() int {
	return xxx_messageInfo_GetImageRequest.Size(m)
}
func (m *GetImageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetImageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetImageRequest proto.InternalMessageInfo

func (m *GetImageRequest) GetImageId() string {
	if m != nil {
		return m.ImageId
	}
	return ""
}

type DeleteImageRequest struct {
	// ID of the Docker image to delete.
	//
	// To get Docker image ID use a [ImageService.List] request.
	ImageId              string   `protobuf:"bytes,1,opt,name=image_id,json=imageId,proto3" json:"image_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteImageRequest) Reset()         { *m = DeleteImageRequest{} }
func (m *DeleteImageRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteImageRequest) ProtoMessage()    {}
func (*DeleteImageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_image_service_181142b342f3b95f, []int{3}
}
func (m *DeleteImageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteImageRequest.Unmarshal(m, b)
}
func (m *DeleteImageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteImageRequest.Marshal(b, m, deterministic)
}
func (dst *DeleteImageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteImageRequest.Merge(dst, src)
}
func (m *DeleteImageRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteImageRequest.Size(m)
}
func (m *DeleteImageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteImageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteImageRequest proto.InternalMessageInfo

func (m *DeleteImageRequest) GetImageId() string {
	if m != nil {
		return m.ImageId
	}
	return ""
}

type DeleteImageMetadata struct {
	// ID of the Docker image that is being deleted.
	ImageId              string   `protobuf:"bytes,1,opt,name=image_id,json=imageId,proto3" json:"image_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteImageMetadata) Reset()         { *m = DeleteImageMetadata{} }
func (m *DeleteImageMetadata) String() string { return proto.CompactTextString(m) }
func (*DeleteImageMetadata) ProtoMessage()    {}
func (*DeleteImageMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_image_service_181142b342f3b95f, []int{4}
}
func (m *DeleteImageMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteImageMetadata.Unmarshal(m, b)
}
func (m *DeleteImageMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteImageMetadata.Marshal(b, m, deterministic)
}
func (dst *DeleteImageMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteImageMetadata.Merge(dst, src)
}
func (m *DeleteImageMetadata) XXX_Size() int {
	return xxx_messageInfo_DeleteImageMetadata.Size(m)
}
func (m *DeleteImageMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteImageMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteImageMetadata proto.InternalMessageInfo

func (m *DeleteImageMetadata) GetImageId() string {
	if m != nil {
		return m.ImageId
	}
	return ""
}

func init() {
	proto.RegisterType((*ListImagesRequest)(nil), "yandex.cloud.containerregistry.v1.ListImagesRequest")
	proto.RegisterType((*ListImagesResponse)(nil), "yandex.cloud.containerregistry.v1.ListImagesResponse")
	proto.RegisterType((*GetImageRequest)(nil), "yandex.cloud.containerregistry.v1.GetImageRequest")
	proto.RegisterType((*DeleteImageRequest)(nil), "yandex.cloud.containerregistry.v1.DeleteImageRequest")
	proto.RegisterType((*DeleteImageMetadata)(nil), "yandex.cloud.containerregistry.v1.DeleteImageMetadata")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ImageServiceClient is the client API for ImageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ImageServiceClient interface {
	// Retrieves the list of Image resources in the specified registry or repository.
	List(ctx context.Context, in *ListImagesRequest, opts ...grpc.CallOption) (*ListImagesResponse, error)
	// Returns the specified Image resource.
	//
	// To get the list of available Image resources, make a [List] request.
	Get(ctx context.Context, in *GetImageRequest, opts ...grpc.CallOption) (*Image, error)
	// Deletes the specified Docker image.
	Delete(ctx context.Context, in *DeleteImageRequest, opts ...grpc.CallOption) (*operation.Operation, error)
}

type imageServiceClient struct {
	cc *grpc.ClientConn
}

func NewImageServiceClient(cc *grpc.ClientConn) ImageServiceClient {
	return &imageServiceClient{cc}
}

func (c *imageServiceClient) List(ctx context.Context, in *ListImagesRequest, opts ...grpc.CallOption) (*ListImagesResponse, error) {
	out := new(ListImagesResponse)
	err := c.cc.Invoke(ctx, "/yandex.cloud.containerregistry.v1.ImageService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imageServiceClient) Get(ctx context.Context, in *GetImageRequest, opts ...grpc.CallOption) (*Image, error) {
	out := new(Image)
	err := c.cc.Invoke(ctx, "/yandex.cloud.containerregistry.v1.ImageService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imageServiceClient) Delete(ctx context.Context, in *DeleteImageRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, "/yandex.cloud.containerregistry.v1.ImageService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ImageServiceServer is the server API for ImageService service.
type ImageServiceServer interface {
	// Retrieves the list of Image resources in the specified registry or repository.
	List(context.Context, *ListImagesRequest) (*ListImagesResponse, error)
	// Returns the specified Image resource.
	//
	// To get the list of available Image resources, make a [List] request.
	Get(context.Context, *GetImageRequest) (*Image, error)
	// Deletes the specified Docker image.
	Delete(context.Context, *DeleteImageRequest) (*operation.Operation, error)
}

func RegisterImageServiceServer(s *grpc.Server, srv ImageServiceServer) {
	s.RegisterService(&_ImageService_serviceDesc, srv)
}

func _ImageService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListImagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.containerregistry.v1.ImageService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageServiceServer).List(ctx, req.(*ListImagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ImageService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.containerregistry.v1.ImageService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageServiceServer).Get(ctx, req.(*GetImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ImageService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.containerregistry.v1.ImageService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageServiceServer).Delete(ctx, req.(*DeleteImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ImageService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "yandex.cloud.containerregistry.v1.ImageService",
	HandlerType: (*ImageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _ImageService_List_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _ImageService_Get_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ImageService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yandex/cloud/containerregistry/v1/image_service.proto",
}

func init() {
	proto.RegisterFile("yandex/cloud/containerregistry/v1/image_service.proto", fileDescriptor_image_service_181142b342f3b95f)
}

var fileDescriptor_image_service_181142b342f3b95f = []byte{
	// 667 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0x4b, 0x6f, 0xd3, 0x4c,
	0x14, 0x95, 0x9b, 0x34, 0x4d, 0xa6, 0xfd, 0xbe, 0x8a, 0x41, 0x48, 0x26, 0xa2, 0xa2, 0xb5, 0x68,
	0xeb, 0x04, 0xfc, 0x2a, 0x74, 0x41, 0x69, 0x55, 0x14, 0x40, 0x55, 0x24, 0x5e, 0x72, 0x91, 0x10,
	0x54, 0x55, 0x70, 0xe2, 0x5b, 0x33, 0xc2, 0xf1, 0x18, 0x7b, 0x12, 0x35, 0xe5, 0xb1, 0x60, 0x99,
	0x2d, 0x62, 0xc3, 0xcf, 0xe0, 0x37, 0x20, 0xb5, 0x6b, 0xf8, 0x0b, 0x2c, 0x58, 0x22, 0x96, 0xac,
	0x90, 0x67, 0x92, 0x34, 0x0f, 0xd4, 0x06, 0x76, 0x96, 0xcf, 0x3d, 0x67, 0xce, 0xbd, 0x67, 0xe6,
	0xa2, 0xd5, 0x96, 0x13, 0xb8, 0xb0, 0x6f, 0xd4, 0x7c, 0xda, 0x70, 0x8d, 0x1a, 0x0d, 0x98, 0x43,
	0x02, 0x88, 0x22, 0xf0, 0x48, 0xcc, 0xa2, 0x96, 0xd1, 0xb4, 0x0c, 0x52, 0x77, 0x3c, 0xa8, 0xc4,
	0x10, 0x35, 0x49, 0x0d, 0xf4, 0x30, 0xa2, 0x8c, 0xe2, 0x05, 0x41, 0xd3, 0x39, 0x4d, 0x1f, 0xa1,
	0xe9, 0x4d, 0x2b, 0x9f, 0xef, 0x28, 0x3b, 0x21, 0x31, 0x68, 0x08, 0x91, 0xc3, 0x08, 0x0d, 0x04,
	0x3d, 0xaf, 0x8d, 0x79, 0x6a, 0xa7, 0x7c, 0x69, 0xa0, 0xbc, 0x27, 0x36, 0x22, 0x3b, 0x37, 0x50,
	0xd7, 0x74, 0x7c, 0xe2, 0xf6, 0xc3, 0x17, 0x3c, 0x4a, 0x3d, 0x1f, 0xb8, 0x23, 0x27, 0x08, 0x28,
	0xe3, 0x60, 0x2c, 0x50, 0xe5, 0xc7, 0x04, 0x3a, 0x73, 0x97, 0xc4, 0xac, 0x9c, 0x1c, 0x1c, 0xdb,
	0xf0, 0xb2, 0x01, 0x31, 0xc3, 0x05, 0x34, 0xdd, 0x75, 0x55, 0x21, 0xae, 0x2c, 0xcd, 0x4b, 0x6a,
	0xae, 0x94, 0x6d, 0x1f, 0x59, 0xe9, 0xf5, 0x8d, 0x55, 0xd3, 0x46, 0x5d, 0xb0, 0xec, 0x62, 0x1f,
	0xcd, 0x46, 0x10, 0xd2, 0x98, 0x30, 0x1a, 0xb5, 0x2a, 0x81, 0x53, 0x07, 0x79, 0x82, 0x97, 0xdf,
	0xfa, 0x79, 0x68, 0x6d, 0xbe, 0xde, 0x71, 0xb4, 0x03, 0x53, 0xbb, 0xbe, 0x7b, 0x59, 0xdd, 0x5c,
	0xdb, 0xd1, 0x2b, 0xda, 0x6e, 0xef, 0x47, 0xa1, 0xa8, 0x1a, 0xea, 0x49, 0x70, 0xa1, 0x50, 0xb4,
	0xff, 0x3f, 0xd6, 0xbe, 0xef, 0xd4, 0x01, 0x2f, 0xa2, 0xdc, 0x1e, 0xf5, 0x5d, 0x88, 0x12, 0x5b,
	0x53, 0x43, 0xb6, 0xb2, 0x02, 0x2a, 0xbb, 0x78, 0x19, 0xe5, 0x42, 0x1e, 0x1f, 0x39, 0x00, 0x39,
	0x35, 0x2f, 0xa9, 0xa9, 0x12, 0xfa, 0x75, 0x68, 0x65, 0xd6, 0x37, 0x2c, 0xd3, 0x34, 0xed, 0x6c,
	0x02, 0x6e, 0x93, 0x03, 0xc0, 0x2a, 0x42, 0xbc, 0x90, 0xd1, 0x17, 0x10, 0xc8, 0x69, 0x2e, 0x98,
	0x6b, 0x1f, 0x59, 0x93, 0xbc, 0xd2, 0xe6, 0x2a, 0x8f, 0x12, 0x0c, 0x2b, 0x28, 0xb3, 0x47, 0x7c,
	0x06, 0x91, 0x3c, 0xc9, 0xab, 0x50, 0xfb, 0xa8, 0xa7, 0xd7, 0x41, 0xf0, 0x25, 0x94, 0xa5, 0x51,
	0x62, 0xae, 0xda, 0x92, 0x33, 0xc3, 0x5a, 0x53, 0x1c, 0x2a, 0xb5, 0x94, 0xb7, 0x08, 0xf7, 0x4f,
	0x3c, 0x0e, 0x69, 0x10, 0x03, 0xbe, 0x89, 0x32, 0x3c, 0xfc, 0x58, 0x96, 0xe6, 0x53, 0xea, 0xf4,
	0x8a, 0xaa, 0x9f, 0x7a, 0xd9, 0x74, 0x2e, 0x61, 0x77, 0x78, 0x78, 0x09, 0xcd, 0x06, 0xb0, 0xcf,
	0x2a, 0x7d, 0x0d, 0xf1, 0x24, 0xec, 0xff, 0x92, 0xdf, 0x0f, 0xbb, 0x9d, 0x28, 0x6b, 0x68, 0x76,
	0x0b, 0xc4, 0xf1, 0xdd, 0xbc, 0x97, 0x51, 0x56, 0xdc, 0xf7, 0x5e, 0xd8, 0x33, 0xdf, 0x0f, 0x2d,
	0xa9, 0x37, 0xd9, 0x29, 0x8e, 0x96, 0x5d, 0x65, 0x03, 0xe1, 0xdb, 0xe0, 0x03, 0x83, 0x7f, 0xa3,
	0x9b, 0xe8, 0x6c, 0x1f, 0xfd, 0x1e, 0x30, 0xc7, 0x75, 0x98, 0x83, 0xcf, 0x0f, 0xf3, 0x7b, 0x8c,
	0x95, 0x76, 0x1a, 0xcd, 0xf0, 0xe2, 0x6d, 0xf1, 0x12, 0xf1, 0x47, 0x09, 0xa5, 0x93, 0xf1, 0xe1,
	0x6b, 0x63, 0x0c, 0x68, 0xe4, 0x66, 0xe7, 0x57, 0xff, 0x92, 0x25, 0xd2, 0x51, 0x16, 0xdf, 0x7d,
	0xfd, 0xf6, 0x7e, 0xe2, 0x22, 0x9e, 0x3b, 0x7e, 0xb6, 0xda, 0xc8, 0xbb, 0x8d, 0xf1, 0x07, 0x09,
	0xa5, 0xb6, 0x80, 0xe1, 0x95, 0x31, 0x4e, 0x19, 0xca, 0x20, 0x3f, 0x76, 0xe0, 0x8a, 0xc9, 0xcd,
	0x14, 0xb1, 0x7a, 0xa2, 0x19, 0xe3, 0x55, 0x77, 0xa6, 0x6f, 0xf0, 0x67, 0x09, 0x65, 0xc4, 0xe0,
	0xf1, 0x38, 0x03, 0x18, 0x8d, 0x38, 0xbf, 0x30, 0x48, 0x3b, 0xde, 0x41, 0x0f, 0xba, 0x5f, 0xca,
	0xb3, 0x4f, 0x5f, 0x8a, 0x57, 0xfe, 0x1c, 0xf0, 0x39, 0xb1, 0x83, 0xc4, 0xce, 0xa9, 0x36, 0xf6,
	0xf4, 0x3b, 0xf5, 0x90, 0xb5, 0x44, 0x1b, 0xc5, 0xb1, 0xdb, 0x28, 0x3d, 0x79, 0xfa, 0xd8, 0x23,
	0xec, 0x79, 0xa3, 0xaa, 0xd7, 0x68, 0xdd, 0x10, 0x86, 0x34, 0xb1, 0xf6, 0x3c, 0xaa, 0x79, 0x10,
	0x70, 0x7d, 0xe3, 0xd4, 0x35, 0x7b, 0x63, 0xe4, 0x67, 0x35, 0xc3, 0xa9, 0x57, 0x7f, 0x07, 0x00,
	0x00, 0xff, 0xff, 0xa8, 0xf4, 0x0e, 0x50, 0x1a, 0x06, 0x00, 0x00,
}
