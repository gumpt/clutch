// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.4
// source: sourcecontrol/github/v1/github.proto

package githubv1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type RepositoryParameters_Visibility int32

const (
	RepositoryParameters_UNSPECIFIED RepositoryParameters_Visibility = 0
	RepositoryParameters_PUBLIC      RepositoryParameters_Visibility = 1
	RepositoryParameters_PRIVATE     RepositoryParameters_Visibility = 2
)

// Enum value maps for RepositoryParameters_Visibility.
var (
	RepositoryParameters_Visibility_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "PUBLIC",
		2: "PRIVATE",
	}
	RepositoryParameters_Visibility_value = map[string]int32{
		"UNSPECIFIED": 0,
		"PUBLIC":      1,
		"PRIVATE":     2,
	}
)

func (x RepositoryParameters_Visibility) Enum() *RepositoryParameters_Visibility {
	p := new(RepositoryParameters_Visibility)
	*p = x
	return p
}

func (x RepositoryParameters_Visibility) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RepositoryParameters_Visibility) Descriptor() protoreflect.EnumDescriptor {
	return file_sourcecontrol_github_v1_github_proto_enumTypes[0].Descriptor()
}

func (RepositoryParameters_Visibility) Type() protoreflect.EnumType {
	return &file_sourcecontrol_github_v1_github_proto_enumTypes[0]
}

func (x RepositoryParameters_Visibility) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RepositoryParameters_Visibility.Descriptor instead.
func (RepositoryParameters_Visibility) EnumDescriptor() ([]byte, []int) {
	return file_sourcecontrol_github_v1_github_proto_rawDescGZIP(), []int{0, 0}
}

// Common parameters for GitHub repostories.
type RepositoryParameters struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Visibility       RepositoryParameters_Visibility `protobuf:"varint,1,opt,name=visibility,proto3,enum=clutch.sourcecontrol.github.v1.RepositoryParameters_Visibility" json:"visibility,omitempty"`
	AllowMergeCommit *wrappers.BoolValue             `protobuf:"bytes,2,opt,name=allow_merge_commit,json=allowMergeCommit,proto3" json:"allow_merge_commit,omitempty"`
	AllowRebaseMerge *wrappers.BoolValue             `protobuf:"bytes,3,opt,name=allow_rebase_merge,json=allowRebaseMerge,proto3" json:"allow_rebase_merge,omitempty"`
	AllowSquashMerge *wrappers.BoolValue             `protobuf:"bytes,4,opt,name=allow_squash_merge,json=allowSquashMerge,proto3" json:"allow_squash_merge,omitempty"`
}

func (x *RepositoryParameters) Reset() {
	*x = RepositoryParameters{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sourcecontrol_github_v1_github_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RepositoryParameters) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RepositoryParameters) ProtoMessage() {}

func (x *RepositoryParameters) ProtoReflect() protoreflect.Message {
	mi := &file_sourcecontrol_github_v1_github_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RepositoryParameters.ProtoReflect.Descriptor instead.
func (*RepositoryParameters) Descriptor() ([]byte, []int) {
	return file_sourcecontrol_github_v1_github_proto_rawDescGZIP(), []int{0}
}

func (x *RepositoryParameters) GetVisibility() RepositoryParameters_Visibility {
	if x != nil {
		return x.Visibility
	}
	return RepositoryParameters_UNSPECIFIED
}

func (x *RepositoryParameters) GetAllowMergeCommit() *wrappers.BoolValue {
	if x != nil {
		return x.AllowMergeCommit
	}
	return nil
}

func (x *RepositoryParameters) GetAllowRebaseMerge() *wrappers.BoolValue {
	if x != nil {
		return x.AllowRebaseMerge
	}
	return nil
}

func (x *RepositoryParameters) GetAllowSquashMerge() *wrappers.BoolValue {
	if x != nil {
		return x.AllowSquashMerge
	}
	return nil
}

// Options unique to creating a repository.
type CreateRepositoryOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Parameters *RepositoryParameters `protobuf:"bytes,1,opt,name=parameters,proto3" json:"parameters,omitempty"`
	// Pass true to create an initial commit with empty README.
	AutoInit bool `protobuf:"varint,2,opt,name=auto_init,json=autoInit,proto3" json:"auto_init,omitempty"`
}

func (x *CreateRepositoryOptions) Reset() {
	*x = CreateRepositoryOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sourcecontrol_github_v1_github_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRepositoryOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRepositoryOptions) ProtoMessage() {}

func (x *CreateRepositoryOptions) ProtoReflect() protoreflect.Message {
	mi := &file_sourcecontrol_github_v1_github_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRepositoryOptions.ProtoReflect.Descriptor instead.
func (*CreateRepositoryOptions) Descriptor() ([]byte, []int) {
	return file_sourcecontrol_github_v1_github_proto_rawDescGZIP(), []int{1}
}

func (x *CreateRepositoryOptions) GetParameters() *RepositoryParameters {
	if x != nil {
		return x.Parameters
	}
	return nil
}

func (x *CreateRepositoryOptions) GetAutoInit() bool {
	if x != nil {
		return x.AutoInit
	}
	return false
}

// Options unique to updating a repository.
type UpdateRepositoryOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Parameters *RepositoryParameters `protobuf:"bytes,1,opt,name=parameters,proto3" json:"parameters,omitempty"`
	// Pass true to archive this repository. Note: You cannot unarchive repositories through the API.
	Archived bool `protobuf:"varint,2,opt,name=archived,proto3" json:"archived,omitempty"`
}

func (x *UpdateRepositoryOptions) Reset() {
	*x = UpdateRepositoryOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sourcecontrol_github_v1_github_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRepositoryOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRepositoryOptions) ProtoMessage() {}

func (x *UpdateRepositoryOptions) ProtoReflect() protoreflect.Message {
	mi := &file_sourcecontrol_github_v1_github_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRepositoryOptions.ProtoReflect.Descriptor instead.
func (*UpdateRepositoryOptions) Descriptor() ([]byte, []int) {
	return file_sourcecontrol_github_v1_github_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateRepositoryOptions) GetParameters() *RepositoryParameters {
	if x != nil {
		return x.Parameters
	}
	return nil
}

func (x *UpdateRepositoryOptions) GetArchived() bool {
	if x != nil {
		return x.Archived
	}
	return false
}

var File_sourcecontrol_github_v1_github_proto protoreflect.FileDescriptor

var file_sourcecontrol_github_v1_github_proto_rawDesc = []byte{
	0x0a, 0x24, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2f,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2e, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x76, 0x31, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x99, 0x03, 0x0a, 0x14, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x12, 0x6b, 0x0a, 0x0a, 0x76, 0x69, 0x73, 0x69,
	0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x3f, 0x2e, 0x63,
	0x6c, 0x75, 0x74, 0x63, 0x68, 0x2e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65,
	0x72, 0x73, 0x2e, 0x56, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x42, 0x0a, 0xfa,
	0x42, 0x07, 0x82, 0x01, 0x04, 0x10, 0x01, 0x20, 0x00, 0x52, 0x0a, 0x76, 0x69, 0x73, 0x69, 0x62,
	0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x48, 0x0a, 0x12, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x6d,
	0x65, 0x72, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x10, 0x61,
	0x6c, 0x6c, 0x6f, 0x77, 0x4d, 0x65, 0x72, 0x67, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x12,
	0x48, 0x0a, 0x12, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x72, 0x65, 0x62, 0x61, 0x73, 0x65, 0x5f,
	0x6d, 0x65, 0x72, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f,
	0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x10, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x52, 0x65,
	0x62, 0x61, 0x73, 0x65, 0x4d, 0x65, 0x72, 0x67, 0x65, 0x12, 0x48, 0x0a, 0x12, 0x61, 0x6c, 0x6c,
	0x6f, 0x77, 0x5f, 0x73, 0x71, 0x75, 0x61, 0x73, 0x68, 0x5f, 0x6d, 0x65, 0x72, 0x67, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x10, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x53, 0x71, 0x75, 0x61, 0x73, 0x68, 0x4d, 0x65,
	0x72, 0x67, 0x65, 0x22, 0x36, 0x0a, 0x0a, 0x56, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74,
	0x79, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x50, 0x55, 0x42, 0x4c, 0x49, 0x43, 0x10, 0x01, 0x12, 0x0b,
	0x0a, 0x07, 0x50, 0x52, 0x49, 0x56, 0x41, 0x54, 0x45, 0x10, 0x02, 0x22, 0x8c, 0x01, 0x0a, 0x17,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x54, 0x0a, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d,
	0x65, 0x74, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x63, 0x6c,
	0x75, 0x74, 0x63, 0x68, 0x2e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72,
	0x73, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x12, 0x1b, 0x0a,
	0x09, 0x61, 0x75, 0x74, 0x6f, 0x5f, 0x69, 0x6e, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x08, 0x61, 0x75, 0x74, 0x6f, 0x49, 0x6e, 0x69, 0x74, 0x22, 0x8b, 0x01, 0x0a, 0x17, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x54, 0x0a, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65,
	0x74, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x63, 0x6c, 0x75,
	0x74, 0x63, 0x68, 0x2e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73,
	0x52, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x12, 0x1a, 0x0a, 0x08,
	0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08,
	0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x64, 0x42, 0x0a, 0x5a, 0x08, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sourcecontrol_github_v1_github_proto_rawDescOnce sync.Once
	file_sourcecontrol_github_v1_github_proto_rawDescData = file_sourcecontrol_github_v1_github_proto_rawDesc
)

func file_sourcecontrol_github_v1_github_proto_rawDescGZIP() []byte {
	file_sourcecontrol_github_v1_github_proto_rawDescOnce.Do(func() {
		file_sourcecontrol_github_v1_github_proto_rawDescData = protoimpl.X.CompressGZIP(file_sourcecontrol_github_v1_github_proto_rawDescData)
	})
	return file_sourcecontrol_github_v1_github_proto_rawDescData
}

var file_sourcecontrol_github_v1_github_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_sourcecontrol_github_v1_github_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_sourcecontrol_github_v1_github_proto_goTypes = []interface{}{
	(RepositoryParameters_Visibility)(0), // 0: clutch.sourcecontrol.github.v1.RepositoryParameters.Visibility
	(*RepositoryParameters)(nil),         // 1: clutch.sourcecontrol.github.v1.RepositoryParameters
	(*CreateRepositoryOptions)(nil),      // 2: clutch.sourcecontrol.github.v1.CreateRepositoryOptions
	(*UpdateRepositoryOptions)(nil),      // 3: clutch.sourcecontrol.github.v1.UpdateRepositoryOptions
	(*wrappers.BoolValue)(nil),           // 4: google.protobuf.BoolValue
}
var file_sourcecontrol_github_v1_github_proto_depIdxs = []int32{
	0, // 0: clutch.sourcecontrol.github.v1.RepositoryParameters.visibility:type_name -> clutch.sourcecontrol.github.v1.RepositoryParameters.Visibility
	4, // 1: clutch.sourcecontrol.github.v1.RepositoryParameters.allow_merge_commit:type_name -> google.protobuf.BoolValue
	4, // 2: clutch.sourcecontrol.github.v1.RepositoryParameters.allow_rebase_merge:type_name -> google.protobuf.BoolValue
	4, // 3: clutch.sourcecontrol.github.v1.RepositoryParameters.allow_squash_merge:type_name -> google.protobuf.BoolValue
	1, // 4: clutch.sourcecontrol.github.v1.CreateRepositoryOptions.parameters:type_name -> clutch.sourcecontrol.github.v1.RepositoryParameters
	1, // 5: clutch.sourcecontrol.github.v1.UpdateRepositoryOptions.parameters:type_name -> clutch.sourcecontrol.github.v1.RepositoryParameters
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_sourcecontrol_github_v1_github_proto_init() }
func file_sourcecontrol_github_v1_github_proto_init() {
	if File_sourcecontrol_github_v1_github_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sourcecontrol_github_v1_github_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RepositoryParameters); i {
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
		file_sourcecontrol_github_v1_github_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRepositoryOptions); i {
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
		file_sourcecontrol_github_v1_github_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRepositoryOptions); i {
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
			RawDescriptor: file_sourcecontrol_github_v1_github_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sourcecontrol_github_v1_github_proto_goTypes,
		DependencyIndexes: file_sourcecontrol_github_v1_github_proto_depIdxs,
		EnumInfos:         file_sourcecontrol_github_v1_github_proto_enumTypes,
		MessageInfos:      file_sourcecontrol_github_v1_github_proto_msgTypes,
	}.Build()
	File_sourcecontrol_github_v1_github_proto = out.File
	file_sourcecontrol_github_v1_github_proto_rawDesc = nil
	file_sourcecontrol_github_v1_github_proto_goTypes = nil
	file_sourcecontrol_github_v1_github_proto_depIdxs = nil
}
