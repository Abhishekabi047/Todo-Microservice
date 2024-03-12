// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: pkg/pb/todo.proto

package pb

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

type CreateTaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Task        string `protobuf:"bytes,1,opt,name=task,proto3" json:"task,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Done        bool   `protobuf:"varint,3,opt,name=done,proto3" json:"done,omitempty"`
}

func (x *CreateTaskRequest) Reset() {
	*x = CreateTaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_todo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTaskRequest) ProtoMessage() {}

func (x *CreateTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_todo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTaskRequest.ProtoReflect.Descriptor instead.
func (*CreateTaskRequest) Descriptor() ([]byte, []int) {
	return file_pkg_pb_todo_proto_rawDescGZIP(), []int{0}
}

func (x *CreateTaskRequest) GetTask() string {
	if x != nil {
		return x.Task
	}
	return ""
}

func (x *CreateTaskRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateTaskRequest) GetDone() bool {
	if x != nil {
		return x.Done
	}
	return false
}

type FindOne struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Task        string `protobuf:"bytes,2,opt,name=task,proto3" json:"task,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Done        bool   `protobuf:"varint,4,opt,name=done,proto3" json:"done,omitempty"`
}

func (x *FindOne) Reset() {
	*x = FindOne{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_todo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindOne) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindOne) ProtoMessage() {}

func (x *FindOne) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_todo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindOne.ProtoReflect.Descriptor instead.
func (*FindOne) Descriptor() ([]byte, []int) {
	return file_pkg_pb_todo_proto_rawDescGZIP(), []int{1}
}

func (x *FindOne) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *FindOne) GetTask() string {
	if x != nil {
		return x.Task
	}
	return ""
}

func (x *FindOne) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *FindOne) GetDone() bool {
	if x != nil {
		return x.Done
	}
	return false
}

type CreateTaskResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int64  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Error  string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *CreateTaskResponse) Reset() {
	*x = CreateTaskResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_todo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTaskResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTaskResponse) ProtoMessage() {}

func (x *CreateTaskResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_todo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTaskResponse.ProtoReflect.Descriptor instead.
func (*CreateTaskResponse) Descriptor() ([]byte, []int) {
	return file_pkg_pb_todo_proto_rawDescGZIP(), []int{2}
}

func (x *CreateTaskResponse) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *CreateTaskResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type DeleteTaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteTaskRequest) Reset() {
	*x = DeleteTaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_todo_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTaskRequest) ProtoMessage() {}

func (x *DeleteTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_todo_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTaskRequest.ProtoReflect.Descriptor instead.
func (*DeleteTaskRequest) Descriptor() ([]byte, []int) {
	return file_pkg_pb_todo_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteTaskRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteTaskResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int64  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Error  string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *DeleteTaskResponse) Reset() {
	*x = DeleteTaskResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_todo_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTaskResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTaskResponse) ProtoMessage() {}

func (x *DeleteTaskResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_todo_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTaskResponse.ProtoReflect.Descriptor instead.
func (*DeleteTaskResponse) Descriptor() ([]byte, []int) {
	return file_pkg_pb_todo_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteTaskResponse) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *DeleteTaskResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type CompleteListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CompleteListRequest) Reset() {
	*x = CompleteListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_todo_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompleteListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompleteListRequest) ProtoMessage() {}

func (x *CompleteListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_todo_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompleteListRequest.ProtoReflect.Descriptor instead.
func (*CompleteListRequest) Descriptor() ([]byte, []int) {
	return file_pkg_pb_todo_proto_rawDescGZIP(), []int{5}
}

type CompleteListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int64      `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Error  string     `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Data   []*FindOne `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *CompleteListResponse) Reset() {
	*x = CompleteListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_todo_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompleteListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompleteListResponse) ProtoMessage() {}

func (x *CompleteListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_todo_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompleteListResponse.ProtoReflect.Descriptor instead.
func (*CompleteListResponse) Descriptor() ([]byte, []int) {
	return file_pkg_pb_todo_proto_rawDescGZIP(), []int{6}
}

func (x *CompleteListResponse) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *CompleteListResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *CompleteListResponse) GetData() []*FindOne {
	if x != nil {
		return x.Data
	}
	return nil
}

type ListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListRequest) Reset() {
	*x = ListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_todo_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRequest) ProtoMessage() {}

func (x *ListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_todo_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRequest.ProtoReflect.Descriptor instead.
func (*ListRequest) Descriptor() ([]byte, []int) {
	return file_pkg_pb_todo_proto_rawDescGZIP(), []int{7}
}

type ListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int64      `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Error  string     `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Data   []*FindOne `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *ListResponse) Reset() {
	*x = ListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_todo_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResponse) ProtoMessage() {}

func (x *ListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_todo_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListResponse.ProtoReflect.Descriptor instead.
func (*ListResponse) Descriptor() ([]byte, []int) {
	return file_pkg_pb_todo_proto_rawDescGZIP(), []int{8}
}

func (x *ListResponse) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *ListResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *ListResponse) GetData() []*FindOne {
	if x != nil {
		return x.Data
	}
	return nil
}

type CompleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CompleteRequest) Reset() {
	*x = CompleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_todo_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompleteRequest) ProtoMessage() {}

func (x *CompleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_todo_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompleteRequest.ProtoReflect.Descriptor instead.
func (*CompleteRequest) Descriptor() ([]byte, []int) {
	return file_pkg_pb_todo_proto_rawDescGZIP(), []int{9}
}

func (x *CompleteRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type CompleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int64    `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Error  string   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Data   *FindOne `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *CompleteResponse) Reset() {
	*x = CompleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_todo_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompleteResponse) ProtoMessage() {}

func (x *CompleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_todo_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompleteResponse.ProtoReflect.Descriptor instead.
func (*CompleteResponse) Descriptor() ([]byte, []int) {
	return file_pkg_pb_todo_proto_rawDescGZIP(), []int{10}
}

func (x *CompleteResponse) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *CompleteResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *CompleteResponse) GetData() *FindOne {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_pkg_pb_todo_proto protoreflect.FileDescriptor

var file_pkg_pb_todo_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x04, 0x74, 0x6f, 0x64, 0x6f, 0x22, 0x5d, 0x0a, 0x11, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61,
	0x73, 0x6b, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x6f, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x04, 0x64, 0x6f, 0x6e, 0x65, 0x22, 0x63, 0x0a, 0x07, 0x46, 0x69, 0x6e, 0x64,
	0x4f, 0x6e, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x6f, 0x6e,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x64, 0x6f, 0x6e, 0x65, 0x22, 0x42, 0x0a,
	0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x22, 0x23, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x42, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x15, 0x0a, 0x13, 0x43, 0x6f,
	0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x67, 0x0a, 0x14, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x21, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x46, 0x69, 0x6e,
	0x64, 0x4f, 0x6e, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x0d, 0x0a, 0x0b, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x5f, 0x0a, 0x0c, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x21, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x46, 0x69, 0x6e,
	0x64, 0x4f, 0x6e, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x21, 0x0a, 0x0f, 0x43, 0x6f,
	0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x63, 0x0a,
	0x10, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12,
	0x21, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e,
	0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x4f, 0x6e, 0x65, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x32, 0xce, 0x02, 0x0a, 0x0b, 0x54, 0x6f, 0x64, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x41, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b,
	0x12, 0x17, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61,
	0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x74, 0x6f, 0x64, 0x6f,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54,
	0x61, 0x73, 0x6b, 0x12, 0x17, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x74,
	0x6f, 0x64, 0x6f, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x47, 0x0a, 0x0c, 0x43, 0x6f, 0x6d, 0x70,
	0x6c, 0x65, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x19, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e,
	0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c,
	0x65, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x33, 0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x11, 0x2e,
	0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x12, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x08, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65,
	0x74, 0x65, 0x12, 0x15, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x74, 0x6f, 0x64, 0x6f,
	0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_pb_todo_proto_rawDescOnce sync.Once
	file_pkg_pb_todo_proto_rawDescData = file_pkg_pb_todo_proto_rawDesc
)

func file_pkg_pb_todo_proto_rawDescGZIP() []byte {
	file_pkg_pb_todo_proto_rawDescOnce.Do(func() {
		file_pkg_pb_todo_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_pb_todo_proto_rawDescData)
	})
	return file_pkg_pb_todo_proto_rawDescData
}

var file_pkg_pb_todo_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_pkg_pb_todo_proto_goTypes = []interface{}{
	(*CreateTaskRequest)(nil),    // 0: todo.CreateTaskRequest
	(*FindOne)(nil),              // 1: todo.FindOne
	(*CreateTaskResponse)(nil),   // 2: todo.CreateTaskResponse
	(*DeleteTaskRequest)(nil),    // 3: todo.DeleteTaskRequest
	(*DeleteTaskResponse)(nil),   // 4: todo.DeleteTaskResponse
	(*CompleteListRequest)(nil),  // 5: todo.CompleteListRequest
	(*CompleteListResponse)(nil), // 6: todo.CompleteListResponse
	(*ListRequest)(nil),          // 7: todo.ListRequest
	(*ListResponse)(nil),         // 8: todo.ListResponse
	(*CompleteRequest)(nil),      // 9: todo.CompleteRequest
	(*CompleteResponse)(nil),     // 10: todo.CompleteResponse
}
var file_pkg_pb_todo_proto_depIdxs = []int32{
	1,  // 0: todo.CompleteListResponse.data:type_name -> todo.FindOne
	1,  // 1: todo.ListResponse.data:type_name -> todo.FindOne
	1,  // 2: todo.CompleteResponse.data:type_name -> todo.FindOne
	0,  // 3: todo.TodoService.CreateTask:input_type -> todo.CreateTaskRequest
	3,  // 4: todo.TodoService.DeleteTask:input_type -> todo.DeleteTaskRequest
	5,  // 5: todo.TodoService.CompleteList:input_type -> todo.CompleteListRequest
	7,  // 6: todo.TodoService.ListTask:input_type -> todo.ListRequest
	9,  // 7: todo.TodoService.Complete:input_type -> todo.CompleteRequest
	2,  // 8: todo.TodoService.CreateTask:output_type -> todo.CreateTaskResponse
	4,  // 9: todo.TodoService.DeleteTask:output_type -> todo.DeleteTaskResponse
	6,  // 10: todo.TodoService.CompleteList:output_type -> todo.CompleteListResponse
	8,  // 11: todo.TodoService.ListTask:output_type -> todo.ListResponse
	10, // 12: todo.TodoService.Complete:output_type -> todo.CompleteResponse
	8,  // [8:13] is the sub-list for method output_type
	3,  // [3:8] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_pkg_pb_todo_proto_init() }
func file_pkg_pb_todo_proto_init() {
	if File_pkg_pb_todo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_pb_todo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTaskRequest); i {
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
		file_pkg_pb_todo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindOne); i {
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
		file_pkg_pb_todo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTaskResponse); i {
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
		file_pkg_pb_todo_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteTaskRequest); i {
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
		file_pkg_pb_todo_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteTaskResponse); i {
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
		file_pkg_pb_todo_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompleteListRequest); i {
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
		file_pkg_pb_todo_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompleteListResponse); i {
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
		file_pkg_pb_todo_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRequest); i {
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
		file_pkg_pb_todo_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListResponse); i {
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
		file_pkg_pb_todo_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompleteRequest); i {
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
		file_pkg_pb_todo_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompleteResponse); i {
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
			RawDescriptor: file_pkg_pb_todo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_pb_todo_proto_goTypes,
		DependencyIndexes: file_pkg_pb_todo_proto_depIdxs,
		MessageInfos:      file_pkg_pb_todo_proto_msgTypes,
	}.Build()
	File_pkg_pb_todo_proto = out.File
	file_pkg_pb_todo_proto_rawDesc = nil
	file_pkg_pb_todo_proto_goTypes = nil
	file_pkg_pb_todo_proto_depIdxs = nil
}
