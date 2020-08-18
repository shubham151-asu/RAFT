// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.3
// source: stub.proto

package gRPC

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type RequestVote struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Term         int64 `protobuf:"varint,1,opt,name=term,proto3" json:"term,omitempty"`
	CandidateID  int64 `protobuf:"varint,2,opt,name=candidateID,proto3" json:"candidateID,omitempty"`
	LastLogIndex int64 `protobuf:"varint,3,opt,name=lastLogIndex,proto3" json:"lastLogIndex,omitempty"`
	LastLogTerm  int64 `protobuf:"varint,4,opt,name=lastLogTerm,proto3" json:"lastLogTerm,omitempty"`
}

func (x *RequestVote) Reset() {
	*x = RequestVote{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stub_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestVote) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestVote) ProtoMessage() {}

func (x *RequestVote) ProtoReflect() protoreflect.Message {
	mi := &file_stub_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestVote.ProtoReflect.Descriptor instead.
func (*RequestVote) Descriptor() ([]byte, []int) {
	return file_stub_proto_rawDescGZIP(), []int{0}
}

func (x *RequestVote) GetTerm() int64 {
	if x != nil {
		return x.Term
	}
	return 0
}

func (x *RequestVote) GetCandidateID() int64 {
	if x != nil {
		return x.CandidateID
	}
	return 0
}

func (x *RequestVote) GetLastLogIndex() int64 {
	if x != nil {
		return x.LastLogIndex
	}
	return 0
}

func (x *RequestVote) GetLastLogTerm() int64 {
	if x != nil {
		return x.LastLogTerm
	}
	return 0
}

type ResponseVote struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Term        int64 `protobuf:"varint,1,opt,name=term,proto3" json:"term,omitempty"`
	VoteGranted bool  `protobuf:"varint,2,opt,name=voteGranted,proto3" json:"voteGranted,omitempty"`
}

func (x *ResponseVote) Reset() {
	*x = ResponseVote{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stub_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseVote) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseVote) ProtoMessage() {}

func (x *ResponseVote) ProtoReflect() protoreflect.Message {
	mi := &file_stub_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseVote.ProtoReflect.Descriptor instead.
func (*ResponseVote) Descriptor() ([]byte, []int) {
	return file_stub_proto_rawDescGZIP(), []int{1}
}

func (x *ResponseVote) GetTerm() int64 {
	if x != nil {
		return x.Term
	}
	return 0
}

func (x *ResponseVote) GetVoteGranted() bool {
	if x != nil {
		return x.VoteGranted
	}
	return false
}

type RequestAppend struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Term         int64                    `protobuf:"varint,1,opt,name=term,proto3" json:"term,omitempty"`
	LeaderId     int64                    `protobuf:"varint,2,opt,name=leaderId,proto3" json:"leaderId,omitempty"`
	PrevLogIndex int64                    `protobuf:"varint,3,opt,name=prevLogIndex,proto3" json:"prevLogIndex,omitempty"`
	PrevLogTerm  int64                    `protobuf:"varint,4,opt,name=prevLogTerm,proto3" json:"prevLogTerm,omitempty"`
	Entries      []*RequestAppendLogEntry `protobuf:"bytes,5,rep,name=entries,proto3" json:"entries,omitempty"`
	LeaderCommit int64                    `protobuf:"varint,6,opt,name=leaderCommit,proto3" json:"leaderCommit,omitempty"`
}

func (x *RequestAppend) Reset() {
	*x = RequestAppend{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stub_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestAppend) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestAppend) ProtoMessage() {}

func (x *RequestAppend) ProtoReflect() protoreflect.Message {
	mi := &file_stub_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestAppend.ProtoReflect.Descriptor instead.
func (*RequestAppend) Descriptor() ([]byte, []int) {
	return file_stub_proto_rawDescGZIP(), []int{2}
}

func (x *RequestAppend) GetTerm() int64 {
	if x != nil {
		return x.Term
	}
	return 0
}

func (x *RequestAppend) GetLeaderId() int64 {
	if x != nil {
		return x.LeaderId
	}
	return 0
}

func (x *RequestAppend) GetPrevLogIndex() int64 {
	if x != nil {
		return x.PrevLogIndex
	}
	return 0
}

func (x *RequestAppend) GetPrevLogTerm() int64 {
	if x != nil {
		return x.PrevLogTerm
	}
	return 0
}

func (x *RequestAppend) GetEntries() []*RequestAppendLogEntry {
	if x != nil {
		return x.Entries
	}
	return nil
}

func (x *RequestAppend) GetLeaderCommit() int64 {
	if x != nil {
		return x.LeaderCommit
	}
	return 0
}

type ResponseAppend struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Term    int64 `protobuf:"varint,1,opt,name=term,proto3" json:"term,omitempty"`
	Success bool  `protobuf:"varint,2,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *ResponseAppend) Reset() {
	*x = ResponseAppend{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stub_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseAppend) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseAppend) ProtoMessage() {}

func (x *ResponseAppend) ProtoReflect() protoreflect.Message {
	mi := &file_stub_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseAppend.ProtoReflect.Descriptor instead.
func (*ResponseAppend) Descriptor() ([]byte, []int) {
	return file_stub_proto_rawDescGZIP(), []int{3}
}

func (x *ResponseAppend) GetTerm() int64 {
	if x != nil {
		return x.Term
	}
	return 0
}

func (x *ResponseAppend) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type ClientResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success  bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Result   string `protobuf:"bytes,2,opt,name=result,proto3" json:"result,omitempty"`
	LeaderId int64  `protobuf:"varint,3,opt,name=leaderId,proto3" json:"leaderId,omitempty"`
}

func (x *ClientResponse) Reset() {
	*x = ClientResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stub_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientResponse) ProtoMessage() {}

func (x *ClientResponse) ProtoReflect() protoreflect.Message {
	mi := &file_stub_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientResponse.ProtoReflect.Descriptor instead.
func (*ClientResponse) Descriptor() ([]byte, []int) {
	return file_stub_proto_rawDescGZIP(), []int{4}
}

func (x *ClientResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *ClientResponse) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

func (x *ClientResponse) GetLeaderId() int64 {
	if x != nil {
		return x.LeaderId
	}
	return 0
}

type ClientRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Command string `protobuf:"bytes,1,opt,name=command,proto3" json:"command,omitempty"`
	Health  string `protobuf:"bytes,2,opt,name=health,proto3" json:"health,omitempty"`
	Key     string `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
	Value   string `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *ClientRequest) Reset() {
	*x = ClientRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stub_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientRequest) ProtoMessage() {}

func (x *ClientRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stub_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientRequest.ProtoReflect.Descriptor instead.
func (*ClientRequest) Descriptor() ([]byte, []int) {
	return file_stub_proto_rawDescGZIP(), []int{5}
}

func (x *ClientRequest) GetCommand() string {
	if x != nil {
		return x.Command
	}
	return ""
}

func (x *ClientRequest) GetHealth() string {
	if x != nil {
		return x.Health
	}
	return ""
}

func (x *ClientRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *ClientRequest) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type LogsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReportLog string `protobuf:"bytes,1,opt,name=reportLog,proto3" json:"reportLog,omitempty"`
}

func (x *LogsRequest) Reset() {
	*x = LogsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stub_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogsRequest) ProtoMessage() {}

func (x *LogsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stub_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogsRequest.ProtoReflect.Descriptor instead.
func (*LogsRequest) Descriptor() ([]byte, []int) {
	return file_stub_proto_rawDescGZIP(), []int{6}
}

func (x *LogsRequest) GetReportLog() string {
	if x != nil {
		return x.ReportLog
	}
	return ""
}

type LogsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entries []*LogsResponseLogEntry `protobuf:"bytes,1,rep,name=entries,proto3" json:"entries,omitempty"`
}

func (x *LogsResponse) Reset() {
	*x = LogsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stub_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogsResponse) ProtoMessage() {}

func (x *LogsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_stub_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogsResponse.ProtoReflect.Descriptor instead.
func (*LogsResponse) Descriptor() ([]byte, []int) {
	return file_stub_proto_rawDescGZIP(), []int{7}
}

func (x *LogsResponse) GetEntries() []*LogsResponseLogEntry {
	if x != nil {
		return x.Entries
	}
	return nil
}

type RequestAppendLogEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Command  string `protobuf:"bytes,1,opt,name=command,proto3" json:"command,omitempty"`
	Key      string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Value    string `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	Term     int64  `protobuf:"varint,4,opt,name=term,proto3" json:"term,omitempty"`
	LogIndex int64  `protobuf:"varint,5,opt,name=logIndex,proto3" json:"logIndex,omitempty"`
}

func (x *RequestAppendLogEntry) Reset() {
	*x = RequestAppendLogEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stub_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestAppendLogEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestAppendLogEntry) ProtoMessage() {}

func (x *RequestAppendLogEntry) ProtoReflect() protoreflect.Message {
	mi := &file_stub_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestAppendLogEntry.ProtoReflect.Descriptor instead.
func (*RequestAppendLogEntry) Descriptor() ([]byte, []int) {
	return file_stub_proto_rawDescGZIP(), []int{2, 0}
}

func (x *RequestAppendLogEntry) GetCommand() string {
	if x != nil {
		return x.Command
	}
	return ""
}

func (x *RequestAppendLogEntry) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *RequestAppendLogEntry) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *RequestAppendLogEntry) GetTerm() int64 {
	if x != nil {
		return x.Term
	}
	return 0
}

func (x *RequestAppendLogEntry) GetLogIndex() int64 {
	if x != nil {
		return x.LogIndex
	}
	return 0
}

type LogsResponseLogEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Command  string `protobuf:"bytes,1,opt,name=command,proto3" json:"command,omitempty"`
	Key      string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Value    string `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	Term     int64  `protobuf:"varint,4,opt,name=term,proto3" json:"term,omitempty"`
	LogIndex int64  `protobuf:"varint,5,opt,name=logIndex,proto3" json:"logIndex,omitempty"`
}

func (x *LogsResponseLogEntry) Reset() {
	*x = LogsResponseLogEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stub_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogsResponseLogEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogsResponseLogEntry) ProtoMessage() {}

func (x *LogsResponseLogEntry) ProtoReflect() protoreflect.Message {
	mi := &file_stub_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogsResponseLogEntry.ProtoReflect.Descriptor instead.
func (*LogsResponseLogEntry) Descriptor() ([]byte, []int) {
	return file_stub_proto_rawDescGZIP(), []int{7, 0}
}

func (x *LogsResponseLogEntry) GetCommand() string {
	if x != nil {
		return x.Command
	}
	return ""
}

func (x *LogsResponseLogEntry) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *LogsResponseLogEntry) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *LogsResponseLogEntry) GetTerm() int64 {
	if x != nil {
		return x.Term
	}
	return 0
}

func (x *LogsResponseLogEntry) GetLogIndex() int64 {
	if x != nil {
		return x.LogIndex
	}
	return 0
}

var File_stub_proto protoreflect.FileDescriptor

var file_stub_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x73, 0x74, 0x75, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x67, 0x52,
	0x50, 0x43, 0x22, 0x89, 0x01, 0x0a, 0x0b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x6f,
	0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x61, 0x6e, 0x64, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x63, 0x61, 0x6e,
	0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x49, 0x44, 0x12, 0x22, 0x0a, 0x0c, 0x6c, 0x61, 0x73, 0x74,
	0x4c, 0x6f, 0x67, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c,
	0x6c, 0x61, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x20, 0x0a, 0x0b,
	0x6c, 0x61, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x54, 0x65, 0x72, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x54, 0x65, 0x72, 0x6d, 0x22, 0x44,
	0x0a, 0x0c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x56, 0x6f, 0x74, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x74, 0x65,
	0x72, 0x6d, 0x12, 0x20, 0x0a, 0x0b, 0x76, 0x6f, 0x74, 0x65, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x65,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x76, 0x6f, 0x74, 0x65, 0x47, 0x72, 0x61,
	0x6e, 0x74, 0x65, 0x64, 0x22, 0xdf, 0x02, 0x0a, 0x0d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x41, 0x70, 0x70, 0x65, 0x6e, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6c, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x70, 0x72, 0x65, 0x76, 0x4c, 0x6f,
	0x67, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x70, 0x72,
	0x65, 0x76, 0x4c, 0x6f, 0x67, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x72,
	0x65, 0x76, 0x4c, 0x6f, 0x67, 0x54, 0x65, 0x72, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0b, 0x70, 0x72, 0x65, 0x76, 0x4c, 0x6f, 0x67, 0x54, 0x65, 0x72, 0x6d, 0x12, 0x36, 0x0a, 0x07,
	0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x67, 0x52, 0x50, 0x43, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x41, 0x70, 0x70, 0x65,
	0x6e, 0x64, 0x2e, 0x6c, 0x6f, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x65, 0x6e, 0x74,
	0x72, 0x69, 0x65, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x43, 0x6f,
	0x6d, 0x6d, 0x69, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x6c, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x1a, 0x7c, 0x0a, 0x08, 0x6c, 0x6f, 0x67, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f,
	0x67, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6c, 0x6f,
	0x67, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x22, 0x3e, 0x0a, 0x0e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x41, 0x70, 0x70, 0x65, 0x6e, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x72, 0x6d,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x5e, 0x0a, 0x0e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6c, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x49, 0x64, 0x22, 0x69, 0x0a, 0x0d, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x22, 0x2b, 0x0a, 0x0b, 0x6c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x4c, 0x6f, 0x67, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x4c, 0x6f, 0x67, 0x22, 0xc3,
	0x01, 0x0a, 0x0c, 0x6c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x35, 0x0a, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1b, 0x2e, 0x67, 0x52, 0x50, 0x43, 0x2e, 0x6c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x6c, 0x6f, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x65,
	0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x1a, 0x7c, 0x0a, 0x08, 0x6c, 0x6f, 0x67, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x67, 0x49,
	0x6e, 0x64, 0x65, 0x78, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6c, 0x6f, 0x67, 0x49,
	0x6e, 0x64, 0x65, 0x78, 0x32, 0x83, 0x02, 0x0a, 0x0a, 0x52, 0x50, 0x43, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x39, 0x0a, 0x0e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x6f,
	0x74, 0x65, 0x52, 0x50, 0x43, 0x12, 0x11, 0x2e, 0x67, 0x52, 0x50, 0x43, 0x2e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x56, 0x6f, 0x74, 0x65, 0x1a, 0x12, 0x2e, 0x67, 0x52, 0x50, 0x43, 0x2e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x56, 0x6f, 0x74, 0x65, 0x22, 0x00, 0x12, 0x3f,
	0x0a, 0x10, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x41, 0x70, 0x70, 0x65, 0x6e, 0x64, 0x52,
	0x50, 0x43, 0x12, 0x13, 0x2e, 0x67, 0x52, 0x50, 0x43, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x41, 0x70, 0x70, 0x65, 0x6e, 0x64, 0x1a, 0x14, 0x2e, 0x67, 0x52, 0x50, 0x43, 0x2e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x41, 0x70, 0x70, 0x65, 0x6e, 0x64, 0x22, 0x00, 0x12,
	0x3f, 0x0a, 0x10, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x52, 0x50, 0x43, 0x12, 0x13, 0x2e, 0x67, 0x52, 0x50, 0x43, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x67, 0x52, 0x50, 0x43, 0x2e,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x38, 0x0a, 0x0d, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x50,
	0x43, 0x12, 0x11, 0x2e, 0x67, 0x52, 0x50, 0x43, 0x2e, 0x6c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x67, 0x52, 0x50, 0x43, 0x2e, 0x6c, 0x6f, 0x67, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_stub_proto_rawDescOnce sync.Once
	file_stub_proto_rawDescData = file_stub_proto_rawDesc
)

func file_stub_proto_rawDescGZIP() []byte {
	file_stub_proto_rawDescOnce.Do(func() {
		file_stub_proto_rawDescData = protoimpl.X.CompressGZIP(file_stub_proto_rawDescData)
	})
	return file_stub_proto_rawDescData
}

var file_stub_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_stub_proto_goTypes = []interface{}{
	(*RequestVote)(nil),           // 0: gRPC.RequestVote
	(*ResponseVote)(nil),          // 1: gRPC.ResponseVote
	(*RequestAppend)(nil),         // 2: gRPC.RequestAppend
	(*ResponseAppend)(nil),        // 3: gRPC.ResponseAppend
	(*ClientResponse)(nil),        // 4: gRPC.ClientResponse
	(*ClientRequest)(nil),         // 5: gRPC.ClientRequest
	(*LogsRequest)(nil),           // 6: gRPC.logsRequest
	(*LogsResponse)(nil),          // 7: gRPC.logsResponse
	(*RequestAppendLogEntry)(nil), // 8: gRPC.RequestAppend.logEntry
	(*LogsResponseLogEntry)(nil),  // 9: gRPC.logsResponse.logEntry
}
var file_stub_proto_depIdxs = []int32{
	8, // 0: gRPC.RequestAppend.entries:type_name -> gRPC.RequestAppend.logEntry
	9, // 1: gRPC.logsResponse.entries:type_name -> gRPC.logsResponse.logEntry
	0, // 2: gRPC.RPCService.RequestVoteRPC:input_type -> gRPC.RequestVote
	2, // 3: gRPC.RPCService.RequestAppendRPC:input_type -> gRPC.RequestAppend
	5, // 4: gRPC.RPCService.ClientRequestRPC:input_type -> gRPC.ClientRequest
	6, // 5: gRPC.RPCService.LogRequestRPC:input_type -> gRPC.logsRequest
	1, // 6: gRPC.RPCService.RequestVoteRPC:output_type -> gRPC.ResponseVote
	3, // 7: gRPC.RPCService.RequestAppendRPC:output_type -> gRPC.ResponseAppend
	4, // 8: gRPC.RPCService.ClientRequestRPC:output_type -> gRPC.ClientResponse
	7, // 9: gRPC.RPCService.LogRequestRPC:output_type -> gRPC.logsResponse
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_stub_proto_init() }
func file_stub_proto_init() {
	if File_stub_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_stub_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestVote); i {
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
		file_stub_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseVote); i {
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
		file_stub_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestAppend); i {
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
		file_stub_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseAppend); i {
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
		file_stub_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientResponse); i {
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
		file_stub_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientRequest); i {
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
		file_stub_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogsRequest); i {
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
		file_stub_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogsResponse); i {
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
		file_stub_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestAppendLogEntry); i {
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
		file_stub_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogsResponseLogEntry); i {
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
			RawDescriptor: file_stub_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_stub_proto_goTypes,
		DependencyIndexes: file_stub_proto_depIdxs,
		MessageInfos:      file_stub_proto_msgTypes,
	}.Build()
	File_stub_proto = out.File
	file_stub_proto_rawDesc = nil
	file_stub_proto_goTypes = nil
	file_stub_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// RPCServiceClient is the client API for RPCService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RPCServiceClient interface {
	RequestVoteRPC(ctx context.Context, in *RequestVote, opts ...grpc.CallOption) (*ResponseVote, error)
	RequestAppendRPC(ctx context.Context, in *RequestAppend, opts ...grpc.CallOption) (*ResponseAppend, error)
	ClientRequestRPC(ctx context.Context, in *ClientRequest, opts ...grpc.CallOption) (*ClientResponse, error)
	LogRequestRPC(ctx context.Context, in *LogsRequest, opts ...grpc.CallOption) (*LogsResponse, error)
}

type rPCServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRPCServiceClient(cc grpc.ClientConnInterface) RPCServiceClient {
	return &rPCServiceClient{cc}
}

func (c *rPCServiceClient) RequestVoteRPC(ctx context.Context, in *RequestVote, opts ...grpc.CallOption) (*ResponseVote, error) {
	out := new(ResponseVote)
	err := c.cc.Invoke(ctx, "/gRPC.RPCService/RequestVoteRPC", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCServiceClient) RequestAppendRPC(ctx context.Context, in *RequestAppend, opts ...grpc.CallOption) (*ResponseAppend, error) {
	out := new(ResponseAppend)
	err := c.cc.Invoke(ctx, "/gRPC.RPCService/RequestAppendRPC", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCServiceClient) ClientRequestRPC(ctx context.Context, in *ClientRequest, opts ...grpc.CallOption) (*ClientResponse, error) {
	out := new(ClientResponse)
	err := c.cc.Invoke(ctx, "/gRPC.RPCService/ClientRequestRPC", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCServiceClient) LogRequestRPC(ctx context.Context, in *LogsRequest, opts ...grpc.CallOption) (*LogsResponse, error) {
	out := new(LogsResponse)
	err := c.cc.Invoke(ctx, "/gRPC.RPCService/LogRequestRPC", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RPCServiceServer is the server API for RPCService service.
type RPCServiceServer interface {
	RequestVoteRPC(context.Context, *RequestVote) (*ResponseVote, error)
	RequestAppendRPC(context.Context, *RequestAppend) (*ResponseAppend, error)
	ClientRequestRPC(context.Context, *ClientRequest) (*ClientResponse, error)
	LogRequestRPC(context.Context, *LogsRequest) (*LogsResponse, error)
}

// UnimplementedRPCServiceServer can be embedded to have forward compatible implementations.
type UnimplementedRPCServiceServer struct {
}

func (*UnimplementedRPCServiceServer) RequestVoteRPC(context.Context, *RequestVote) (*ResponseVote, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestVoteRPC not implemented")
}
func (*UnimplementedRPCServiceServer) RequestAppendRPC(context.Context, *RequestAppend) (*ResponseAppend, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestAppendRPC not implemented")
}
func (*UnimplementedRPCServiceServer) ClientRequestRPC(context.Context, *ClientRequest) (*ClientResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClientRequestRPC not implemented")
}
func (*UnimplementedRPCServiceServer) LogRequestRPC(context.Context, *LogsRequest) (*LogsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LogRequestRPC not implemented")
}

func RegisterRPCServiceServer(s *grpc.Server, srv RPCServiceServer) {
	s.RegisterService(&_RPCService_serviceDesc, srv)
}

func _RPCService_RequestVoteRPC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestVote)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServiceServer).RequestVoteRPC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gRPC.RPCService/RequestVoteRPC",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServiceServer).RequestVoteRPC(ctx, req.(*RequestVote))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCService_RequestAppendRPC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestAppend)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServiceServer).RequestAppendRPC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gRPC.RPCService/RequestAppendRPC",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServiceServer).RequestAppendRPC(ctx, req.(*RequestAppend))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCService_ClientRequestRPC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServiceServer).ClientRequestRPC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gRPC.RPCService/ClientRequestRPC",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServiceServer).ClientRequestRPC(ctx, req.(*ClientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCService_LogRequestRPC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServiceServer).LogRequestRPC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gRPC.RPCService/LogRequestRPC",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServiceServer).LogRequestRPC(ctx, req.(*LogsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RPCService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gRPC.RPCService",
	HandlerType: (*RPCServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RequestVoteRPC",
			Handler:    _RPCService_RequestVoteRPC_Handler,
		},
		{
			MethodName: "RequestAppendRPC",
			Handler:    _RPCService_RequestAppendRPC_Handler,
		},
		{
			MethodName: "ClientRequestRPC",
			Handler:    _RPCService_ClientRequestRPC_Handler,
		},
		{
			MethodName: "LogRequestRPC",
			Handler:    _RPCService_LogRequestRPC_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "stub.proto",
}
