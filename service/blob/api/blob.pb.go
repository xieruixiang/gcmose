// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.15.7
// source: blob.proto

package blobpb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Operation_Status int32

const (
	Operation_Status_DEFAULT  Operation_Status = 0
	Operation_Status_DOWNLOAD Operation_Status = 1
	Operation_Status_UPLOAD   Operation_Status = 2
)

// Enum value maps for Operation_Status.
var (
	Operation_Status_name = map[int32]string{
		0: "DEFAULT",
		1: "DOWNLOAD",
		2: "UPLOAD",
	}
	Operation_Status_value = map[string]int32{
		"DEFAULT":  0,
		"DOWNLOAD": 1,
		"UPLOAD":   2,
	}
)

func (x Operation_Status) Enum() *Operation_Status {
	p := new(Operation_Status)
	*p = x
	return p
}

func (x Operation_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Operation_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_blob_proto_enumTypes[0].Descriptor()
}

func (Operation_Status) Type() protoreflect.EnumType {
	return &file_blob_proto_enumTypes[0]
}

func (x Operation_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Operation_Status.Descriptor instead.
func (Operation_Status) EnumDescriptor() ([]byte, []int) {
	return file_blob_proto_rawDescGZIP(), []int{0}
}

type CreateUrlRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path      string           `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	Operation Operation_Status `protobuf:"varint,2,opt,name=operation,proto3,enum=blob.v1.Operation_Status" json:"operation,omitempty"`
}

func (x *CreateUrlRequest) Reset() {
	*x = CreateUrlRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blob_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUrlRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUrlRequest) ProtoMessage() {}

func (x *CreateUrlRequest) ProtoReflect() protoreflect.Message {
	mi := &file_blob_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUrlRequest.ProtoReflect.Descriptor instead.
func (*CreateUrlRequest) Descriptor() ([]byte, []int) {
	return file_blob_proto_rawDescGZIP(), []int{0}
}

func (x *CreateUrlRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *CreateUrlRequest) GetOperation() Operation_Status {
	if x != nil {
		return x.Operation
	}
	return Operation_Status_DEFAULT
}

type CreateUrlResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *CreateUrlResponse) Reset() {
	*x = CreateUrlResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blob_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUrlResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUrlResponse) ProtoMessage() {}

func (x *CreateUrlResponse) ProtoReflect() protoreflect.Message {
	mi := &file_blob_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUrlResponse.ProtoReflect.Descriptor instead.
func (*CreateUrlResponse) Descriptor() ([]byte, []int) {
	return file_blob_proto_rawDescGZIP(), []int{1}
}

func (x *CreateUrlResponse) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

var File_blob_proto protoreflect.FileDescriptor

var file_blob_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x62, 0x6c, 0x6f, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x62, 0x6c,
	0x6f, 0x62, 0x2e, 0x76, 0x31, 0x22, 0x5f, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55,
	0x72, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74,
	0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x37, 0x0a,
	0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x19, 0x2e, 0x62, 0x6c, 0x6f, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x09, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x25, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x55, 0x72, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75,
	0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x2a, 0x39, 0x0a,
	0x10, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x0b, 0x0a, 0x07, 0x44, 0x45, 0x46, 0x41, 0x55, 0x4c, 0x54, 0x10, 0x00, 0x12, 0x0c,
	0x0a, 0x08, 0x44, 0x4f, 0x57, 0x4e, 0x4c, 0x4f, 0x41, 0x44, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06,
	0x55, 0x50, 0x4c, 0x4f, 0x41, 0x44, 0x10, 0x02, 0x32, 0x53, 0x0a, 0x0b, 0x42, 0x6c, 0x6f, 0x62,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x44, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x55, 0x72, 0x6c, 0x12, 0x19, 0x2e, 0x62, 0x6c, 0x6f, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x72, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1a, 0x2e, 0x62, 0x6c, 0x6f, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x55, 0x72, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x21, 0x5a,
	0x1f, 0x2f, 0x63, 0x63, 0x6d, 0x6f, 0x73, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2f, 0x62, 0x6c, 0x6f, 0x62, 0x2f, 0x61, 0x70, 0x69, 0x3b, 0x62, 0x6c, 0x6f, 0x62, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_blob_proto_rawDescOnce sync.Once
	file_blob_proto_rawDescData = file_blob_proto_rawDesc
)

func file_blob_proto_rawDescGZIP() []byte {
	file_blob_proto_rawDescOnce.Do(func() {
		file_blob_proto_rawDescData = protoimpl.X.CompressGZIP(file_blob_proto_rawDescData)
	})
	return file_blob_proto_rawDescData
}

var file_blob_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_blob_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_blob_proto_goTypes = []interface{}{
	(Operation_Status)(0),     // 0: blob.v1.Operation_Status
	(*CreateUrlRequest)(nil),  // 1: blob.v1.CreateUrlRequest
	(*CreateUrlResponse)(nil), // 2: blob.v1.CreateUrlResponse
}
var file_blob_proto_depIdxs = []int32{
	0, // 0: blob.v1.CreateUrlRequest.operation:type_name -> blob.v1.Operation_Status
	1, // 1: blob.v1.BlobService.CreateUrl:input_type -> blob.v1.CreateUrlRequest
	2, // 2: blob.v1.BlobService.CreateUrl:output_type -> blob.v1.CreateUrlResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_blob_proto_init() }
func file_blob_proto_init() {
	if File_blob_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_blob_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUrlRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_blob_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUrlResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_blob_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_blob_proto_goTypes,
		DependencyIndexes: file_blob_proto_depIdxs,
		EnumInfos:         file_blob_proto_enumTypes,
		MessageInfos:      file_blob_proto_msgTypes,
	}.Build()
	File_blob_proto = out.File
	file_blob_proto_rawDesc = nil
	file_blob_proto_goTypes = nil
	file_blob_proto_depIdxs = nil
}
