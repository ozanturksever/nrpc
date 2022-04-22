// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.12.4
// source: nrpc.proto

package nrpc

import (
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
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

type SubjectRule int32

const (
	SubjectRule_COPY    SubjectRule = 0
	SubjectRule_TOLOWER SubjectRule = 1
)

// Enum value maps for SubjectRule.
var (
	SubjectRule_name = map[int32]string{
		0: "COPY",
		1: "TOLOWER",
	}
	SubjectRule_value = map[string]int32{
		"COPY":    0,
		"TOLOWER": 1,
	}
)

func (x SubjectRule) Enum() *SubjectRule {
	p := new(SubjectRule)
	*p = x
	return p
}

func (x SubjectRule) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SubjectRule) Descriptor() protoreflect.EnumDescriptor {
	return file_nrpc_proto_enumTypes[0].Descriptor()
}

func (SubjectRule) Type() protoreflect.EnumType {
	return &file_nrpc_proto_enumTypes[0]
}

func (x SubjectRule) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SubjectRule.Descriptor instead.
func (SubjectRule) EnumDescriptor() ([]byte, []int) {
	return file_nrpc_proto_rawDescGZIP(), []int{0}
}

type Error_Type int32

const (
	Error_CLIENT        Error_Type = 0
	Error_SERVER        Error_Type = 1
	Error_EOS           Error_Type = 3
	Error_SERVERTOOBUSY Error_Type = 4
)

// Enum value maps for Error_Type.
var (
	Error_Type_name = map[int32]string{
		0: "CLIENT",
		1: "SERVER",
		3: "EOS",
		4: "SERVERTOOBUSY",
	}
	Error_Type_value = map[string]int32{
		"CLIENT":        0,
		"SERVER":        1,
		"EOS":           3,
		"SERVERTOOBUSY": 4,
	}
)

func (x Error_Type) Enum() *Error_Type {
	p := new(Error_Type)
	*p = x
	return p
}

func (x Error_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Error_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_nrpc_proto_enumTypes[1].Descriptor()
}

func (Error_Type) Type() protoreflect.EnumType {
	return &file_nrpc_proto_enumTypes[1]
}

func (x Error_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Error_Type.Descriptor instead.
func (Error_Type) EnumDescriptor() ([]byte, []int) {
	return file_nrpc_proto_rawDescGZIP(), []int{0, 0}
}

type Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type     Error_Type `protobuf:"varint,1,opt,name=type,proto3,enum=nrpc.Error_Type" json:"type,omitempty"`
	Message  string     `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	MsgCount uint32     `protobuf:"varint,3,opt,name=msgCount,proto3" json:"msgCount,omitempty"`
}

func (x *Error) Reset() {
	*x = Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nrpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_nrpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (*Error) Descriptor() ([]byte, []int) {
	return file_nrpc_proto_rawDescGZIP(), []int{0}
}

func (x *Error) GetType() Error_Type {
	if x != nil {
		return x.Type
	}
	return Error_CLIENT
}

func (x *Error) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Error) GetMsgCount() uint32 {
	if x != nil {
		return x.MsgCount
	}
	return 0
}

type Void struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Void) Reset() {
	*x = Void{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nrpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Void) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Void) ProtoMessage() {}

func (x *Void) ProtoReflect() protoreflect.Message {
	mi := &file_nrpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Void.ProtoReflect.Descriptor instead.
func (*Void) Descriptor() ([]byte, []int) {
	return file_nrpc_proto_rawDescGZIP(), []int{1}
}

type NoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *NoRequest) Reset() {
	*x = NoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nrpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NoRequest) ProtoMessage() {}

func (x *NoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nrpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NoRequest.ProtoReflect.Descriptor instead.
func (*NoRequest) Descriptor() ([]byte, []int) {
	return file_nrpc_proto_rawDescGZIP(), []int{2}
}

type NoReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *NoReply) Reset() {
	*x = NoReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nrpc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NoReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NoReply) ProtoMessage() {}

func (x *NoReply) ProtoReflect() protoreflect.Message {
	mi := &file_nrpc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NoReply.ProtoReflect.Descriptor instead.
func (*NoReply) Descriptor() ([]byte, []int) {
	return file_nrpc_proto_rawDescGZIP(), []int{3}
}

type HeartBeat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Lastbeat bool `protobuf:"varint,1,opt,name=lastbeat,proto3" json:"lastbeat,omitempty"`
}

func (x *HeartBeat) Reset() {
	*x = HeartBeat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nrpc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeartBeat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeartBeat) ProtoMessage() {}

func (x *HeartBeat) ProtoReflect() protoreflect.Message {
	mi := &file_nrpc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeartBeat.ProtoReflect.Descriptor instead.
func (*HeartBeat) Descriptor() ([]byte, []int) {
	return file_nrpc_proto_rawDescGZIP(), []int{4}
}

func (x *HeartBeat) GetLastbeat() bool {
	if x != nil {
		return x.Lastbeat
	}
	return false
}

var file_nrpc_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptor.FileOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50000,
		Name:          "nrpc.packageSubject",
		Tag:           "bytes,50000,opt,name=packageSubject",
		Filename:      "nrpc.proto",
	},
	{
		ExtendedType:  (*descriptor.FileOptions)(nil),
		ExtensionType: ([]string)(nil),
		Field:         50001,
		Name:          "nrpc.packageSubjectParams",
		Tag:           "bytes,50001,rep,name=packageSubjectParams",
		Filename:      "nrpc.proto",
	},
	{
		ExtendedType:  (*descriptor.FileOptions)(nil),
		ExtensionType: (*SubjectRule)(nil),
		Field:         50002,
		Name:          "nrpc.serviceSubjectRule",
		Tag:           "varint,50002,opt,name=serviceSubjectRule,enum=nrpc.SubjectRule",
		Filename:      "nrpc.proto",
	},
	{
		ExtendedType:  (*descriptor.FileOptions)(nil),
		ExtensionType: (*SubjectRule)(nil),
		Field:         50003,
		Name:          "nrpc.methodSubjectRule",
		Tag:           "varint,50003,opt,name=methodSubjectRule,enum=nrpc.SubjectRule",
		Filename:      "nrpc.proto",
	},
	{
		ExtendedType:  (*descriptor.ServiceOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         51000,
		Name:          "nrpc.serviceSubject",
		Tag:           "bytes,51000,opt,name=serviceSubject",
		Filename:      "nrpc.proto",
	},
	{
		ExtendedType:  (*descriptor.ServiceOptions)(nil),
		ExtensionType: ([]string)(nil),
		Field:         51001,
		Name:          "nrpc.serviceSubjectParams",
		Tag:           "bytes,51001,rep,name=serviceSubjectParams",
		Filename:      "nrpc.proto",
	},
	{
		ExtendedType:  (*descriptor.ServiceOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         51002,
		Name:          "nrpc.serviceCustomSubject",
		Tag:           "varint,51002,opt,name=serviceCustomSubject",
		Filename:      "nrpc.proto",
	},
	{
		ExtendedType:  (*descriptor.MethodOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         52000,
		Name:          "nrpc.methodSubject",
		Tag:           "bytes,52000,opt,name=methodSubject",
		Filename:      "nrpc.proto",
	},
	{
		ExtendedType:  (*descriptor.MethodOptions)(nil),
		ExtensionType: ([]string)(nil),
		Field:         52001,
		Name:          "nrpc.methodSubjectParams",
		Tag:           "bytes,52001,rep,name=methodSubjectParams",
		Filename:      "nrpc.proto",
	},
	{
		ExtendedType:  (*descriptor.MethodOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         52002,
		Name:          "nrpc.streamedReply",
		Tag:           "varint,52002,opt,name=streamedReply",
		Filename:      "nrpc.proto",
	},
}

// Extension fields to descriptor.FileOptions.
var (
	// A custom subject prefix to use instead of the package name
	//
	// optional string packageSubject = 50000;
	E_PackageSubject = &file_nrpc_proto_extTypes[0]
	// Parameters included in the subject at the package level
	//
	// repeated string packageSubjectParams = 50001;
	E_PackageSubjectParams = &file_nrpc_proto_extTypes[1]
	// Default rule to build a service subject from the service name
	//
	// optional nrpc.SubjectRule serviceSubjectRule = 50002;
	E_ServiceSubjectRule = &file_nrpc_proto_extTypes[2]
	// Default rule to build a method subject from its name
	//
	// optional nrpc.SubjectRule methodSubjectRule = 50003;
	E_MethodSubjectRule = &file_nrpc_proto_extTypes[3]
)

// Extension fields to descriptor.ServiceOptions.
var (
	// A custom subject token to use instead of (service name + serviceSubjectRule)
	//
	// optional string serviceSubject = 51000;
	E_ServiceSubject = &file_nrpc_proto_extTypes[4]
	// Parameters included in the subject at the service level
	//
	// repeated string serviceSubjectParams = 51001;
	E_ServiceSubjectParams = &file_nrpc_proto_extTypes[5]
	// optional bool serviceCustomSubject = 51002;
	E_ServiceCustomSubject = &file_nrpc_proto_extTypes[6]
)

// Extension fields to descriptor.MethodOptions.
var (
	// A custom subject to use instead of (methor name + methodSubjectRule)
	//
	// optional string methodSubject = 52000;
	E_MethodSubject = &file_nrpc_proto_extTypes[7]
	// Parameters included in the subject at the method level
	//
	// repeated string methodSubjectParams = 52001;
	E_MethodSubjectParams = &file_nrpc_proto_extTypes[8]
	// If true, the method returns a stream of reply messages instead of just one
	//
	// optional bool streamedReply = 52002;
	E_StreamedReply = &file_nrpc_proto_extTypes[9]
)

var File_nrpc_proto protoreflect.FileDescriptor

var file_nrpc_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x6e, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x6e, 0x72,
	0x70, 0x63, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9f, 0x01, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x24,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x6e,
	0x72, 0x70, 0x63, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x6d, 0x73, 0x67, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x08, 0x6d, 0x73, 0x67, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x3a, 0x0a, 0x04, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x10, 0x00, 0x12, 0x0a,
	0x0a, 0x06, 0x53, 0x45, 0x52, 0x56, 0x45, 0x52, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x45, 0x4f,
	0x53, 0x10, 0x03, 0x12, 0x11, 0x0a, 0x0d, 0x53, 0x45, 0x52, 0x56, 0x45, 0x52, 0x54, 0x4f, 0x4f,
	0x42, 0x55, 0x53, 0x59, 0x10, 0x04, 0x22, 0x06, 0x0a, 0x04, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x0b,
	0x0a, 0x09, 0x4e, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x09, 0x0a, 0x07, 0x4e,
	0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x27, 0x0a, 0x09, 0x48, 0x65, 0x61, 0x72, 0x74, 0x42,
	0x65, 0x61, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x62, 0x65, 0x61, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x62, 0x65, 0x61, 0x74, 0x2a,
	0x24, 0x0a, 0x0b, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x75, 0x6c, 0x65, 0x12, 0x08,
	0x0a, 0x04, 0x43, 0x4f, 0x50, 0x59, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x54, 0x4f, 0x4c, 0x4f,
	0x57, 0x45, 0x52, 0x10, 0x01, 0x3a, 0x46, 0x0a, 0x0e, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65,
	0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xd0, 0x86, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x70,
	0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x3a, 0x52, 0x0a,
	0x14, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0xd1, 0x86, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x14, 0x70, 0x61, 0x63,
	0x6b, 0x61, 0x67, 0x65, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x3a, 0x61, 0x0a, 0x12, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x75, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x52, 0x75, 0x6c, 0x65, 0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xd2, 0x86, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e,
	0x6e, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x75, 0x6c, 0x65,
	0x52, 0x12, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x52, 0x75, 0x6c, 0x65, 0x3a, 0x5f, 0x0a, 0x11, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x53, 0x75,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x75, 0x6c, 0x65, 0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x6c, 0x65,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xd3, 0x86, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x11, 0x2e, 0x6e, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x75,
	0x6c, 0x65, 0x52, 0x11, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x52, 0x75, 0x6c, 0x65, 0x3a, 0x49, 0x0a, 0x0e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xb8, 0x8e, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x3a, 0x55, 0x0a, 0x14, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x75, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xb9, 0x8e, 0x03, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x14, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x3a, 0x55, 0x0a, 0x14, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12,
	0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0xba, 0x8e, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x14, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x3a, 0x46,
	0x0a, 0x0d, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12,
	0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0xa0, 0x96, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x53,
	0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x3a, 0x52, 0x0a, 0x13, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x1e, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xa1, 0x96,
	0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x13, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x53, 0x75, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x3a, 0x46, 0x0a, 0x0d, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x65, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x1e, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xa2, 0x96, 0x03, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0d, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x65, 0x64, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x42, 0x1f, 0x5a, 0x1d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6f, 0x7a, 0x61, 0x6e, 0x74, 0x75, 0x72, 0x6b, 0x73, 0x65, 0x76, 0x65, 0x72, 0x2f, 0x6e,
	0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_nrpc_proto_rawDescOnce sync.Once
	file_nrpc_proto_rawDescData = file_nrpc_proto_rawDesc
)

func file_nrpc_proto_rawDescGZIP() []byte {
	file_nrpc_proto_rawDescOnce.Do(func() {
		file_nrpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_nrpc_proto_rawDescData)
	})
	return file_nrpc_proto_rawDescData
}

var file_nrpc_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_nrpc_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_nrpc_proto_goTypes = []interface{}{
	(SubjectRule)(0),                  // 0: nrpc.SubjectRule
	(Error_Type)(0),                   // 1: nrpc.Error.Type
	(*Error)(nil),                     // 2: nrpc.Error
	(*Void)(nil),                      // 3: nrpc.Void
	(*NoRequest)(nil),                 // 4: nrpc.NoRequest
	(*NoReply)(nil),                   // 5: nrpc.NoReply
	(*HeartBeat)(nil),                 // 6: nrpc.HeartBeat
	(*descriptor.FileOptions)(nil),    // 7: google.protobuf.FileOptions
	(*descriptor.ServiceOptions)(nil), // 8: google.protobuf.ServiceOptions
	(*descriptor.MethodOptions)(nil),  // 9: google.protobuf.MethodOptions
}
var file_nrpc_proto_depIdxs = []int32{
	1,  // 0: nrpc.Error.type:type_name -> nrpc.Error.Type
	7,  // 1: nrpc.packageSubject:extendee -> google.protobuf.FileOptions
	7,  // 2: nrpc.packageSubjectParams:extendee -> google.protobuf.FileOptions
	7,  // 3: nrpc.serviceSubjectRule:extendee -> google.protobuf.FileOptions
	7,  // 4: nrpc.methodSubjectRule:extendee -> google.protobuf.FileOptions
	8,  // 5: nrpc.serviceSubject:extendee -> google.protobuf.ServiceOptions
	8,  // 6: nrpc.serviceSubjectParams:extendee -> google.protobuf.ServiceOptions
	8,  // 7: nrpc.serviceCustomSubject:extendee -> google.protobuf.ServiceOptions
	9,  // 8: nrpc.methodSubject:extendee -> google.protobuf.MethodOptions
	9,  // 9: nrpc.methodSubjectParams:extendee -> google.protobuf.MethodOptions
	9,  // 10: nrpc.streamedReply:extendee -> google.protobuf.MethodOptions
	0,  // 11: nrpc.serviceSubjectRule:type_name -> nrpc.SubjectRule
	0,  // 12: nrpc.methodSubjectRule:type_name -> nrpc.SubjectRule
	13, // [13:13] is the sub-list for method output_type
	13, // [13:13] is the sub-list for method input_type
	11, // [11:13] is the sub-list for extension type_name
	1,  // [1:11] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_nrpc_proto_init() }
func file_nrpc_proto_init() {
	if File_nrpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_nrpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Error); i {
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
		file_nrpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Void); i {
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
		file_nrpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NoRequest); i {
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
		file_nrpc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NoReply); i {
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
		file_nrpc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeartBeat); i {
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
			RawDescriptor: file_nrpc_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   5,
			NumExtensions: 10,
			NumServices:   0,
		},
		GoTypes:           file_nrpc_proto_goTypes,
		DependencyIndexes: file_nrpc_proto_depIdxs,
		EnumInfos:         file_nrpc_proto_enumTypes,
		MessageInfos:      file_nrpc_proto_msgTypes,
		ExtensionInfos:    file_nrpc_proto_extTypes,
	}.Build()
	File_nrpc_proto = out.File
	file_nrpc_proto_rawDesc = nil
	file_nrpc_proto_goTypes = nil
	file_nrpc_proto_depIdxs = nil
}
