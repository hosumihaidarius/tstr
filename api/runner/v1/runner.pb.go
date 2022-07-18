// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: runner/v1/runner.proto

package runnerv1

import (
	v1 "github.com/nanzhong/tstr/api/common/v1"
	_ "github.com/nanzhong/tstr/api/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RegisterRunnerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name                     string            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	AcceptTestLabelSelectors map[string]string `protobuf:"bytes,2,rep,name=accept_test_label_selectors,proto3" json:"accept_test_label_selectors,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	RejectTestLabelSelectors map[string]string `protobuf:"bytes,3,rep,name=reject_test_label_selectors,proto3" json:"reject_test_label_selectors,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *RegisterRunnerRequest) Reset() {
	*x = RegisterRunnerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_runner_v1_runner_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterRunnerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterRunnerRequest) ProtoMessage() {}

func (x *RegisterRunnerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_runner_v1_runner_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterRunnerRequest.ProtoReflect.Descriptor instead.
func (*RegisterRunnerRequest) Descriptor() ([]byte, []int) {
	return file_runner_v1_runner_proto_rawDescGZIP(), []int{0}
}

func (x *RegisterRunnerRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RegisterRunnerRequest) GetAcceptTestLabelSelectors() map[string]string {
	if x != nil {
		return x.AcceptTestLabelSelectors
	}
	return nil
}

func (x *RegisterRunnerRequest) GetRejectTestLabelSelectors() map[string]string {
	if x != nil {
		return x.RejectTestLabelSelectors
	}
	return nil
}

type RegisterRunnerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Runner *v1.Runner `protobuf:"bytes,1,opt,name=runner,proto3" json:"runner,omitempty"` // required
}

func (x *RegisterRunnerResponse) Reset() {
	*x = RegisterRunnerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_runner_v1_runner_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterRunnerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterRunnerResponse) ProtoMessage() {}

func (x *RegisterRunnerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_runner_v1_runner_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterRunnerResponse.ProtoReflect.Descriptor instead.
func (*RegisterRunnerResponse) Descriptor() ([]byte, []int) {
	return file_runner_v1_runner_proto_rawDescGZIP(), []int{1}
}

func (x *RegisterRunnerResponse) GetRunner() *v1.Runner {
	if x != nil {
		return x.Runner
	}
	return nil
}

type NextRunRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *NextRunRequest) Reset() {
	*x = NextRunRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_runner_v1_runner_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NextRunRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NextRunRequest) ProtoMessage() {}

func (x *NextRunRequest) ProtoReflect() protoreflect.Message {
	mi := &file_runner_v1_runner_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NextRunRequest.ProtoReflect.Descriptor instead.
func (*NextRunRequest) Descriptor() ([]byte, []int) {
	return file_runner_v1_runner_proto_rawDescGZIP(), []int{2}
}

func (x *NextRunRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type NextRunResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Run *v1.Run `protobuf:"bytes,1,opt,name=run,proto3" json:"run,omitempty"` // required
}

func (x *NextRunResponse) Reset() {
	*x = NextRunResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_runner_v1_runner_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NextRunResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NextRunResponse) ProtoMessage() {}

func (x *NextRunResponse) ProtoReflect() protoreflect.Message {
	mi := &file_runner_v1_runner_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NextRunResponse.ProtoReflect.Descriptor instead.
func (*NextRunResponse) Descriptor() ([]byte, []int) {
	return file_runner_v1_runner_proto_rawDescGZIP(), []int{3}
}

func (x *NextRunResponse) GetRun() *v1.Run {
	if x != nil {
		return x.Run
	}
	return nil
}

type SubmitRunRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	RunId      string                 `protobuf:"bytes,2,opt,name=run_id,json=runId,proto3" json:"run_id,omitempty"`
	Result     v1.Run_Result          `protobuf:"varint,3,opt,name=result,proto3,enum=tstr.common.v1.Run_Result" json:"result,omitempty"`
	Logs       []*v1.Run_Log          `protobuf:"bytes,4,rep,name=logs,proto3" json:"logs,omitempty"`
	StartedAt  *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=started_at,proto3" json:"started_at,omitempty"`
	FinishedAt *timestamppb.Timestamp `protobuf:"bytes,11,opt,name=finished_at,proto3" json:"finished_at,omitempty"`
}

func (x *SubmitRunRequest) Reset() {
	*x = SubmitRunRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_runner_v1_runner_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubmitRunRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmitRunRequest) ProtoMessage() {}

func (x *SubmitRunRequest) ProtoReflect() protoreflect.Message {
	mi := &file_runner_v1_runner_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmitRunRequest.ProtoReflect.Descriptor instead.
func (*SubmitRunRequest) Descriptor() ([]byte, []int) {
	return file_runner_v1_runner_proto_rawDescGZIP(), []int{4}
}

func (x *SubmitRunRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SubmitRunRequest) GetRunId() string {
	if x != nil {
		return x.RunId
	}
	return ""
}

func (x *SubmitRunRequest) GetResult() v1.Run_Result {
	if x != nil {
		return x.Result
	}
	return v1.Run_Result(0)
}

func (x *SubmitRunRequest) GetLogs() []*v1.Run_Log {
	if x != nil {
		return x.Logs
	}
	return nil
}

func (x *SubmitRunRequest) GetStartedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.StartedAt
	}
	return nil
}

func (x *SubmitRunRequest) GetFinishedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.FinishedAt
	}
	return nil
}

type SubmitRunResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SubmitRunResponse) Reset() {
	*x = SubmitRunResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_runner_v1_runner_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubmitRunResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmitRunResponse) ProtoMessage() {}

func (x *SubmitRunResponse) ProtoReflect() protoreflect.Message {
	mi := &file_runner_v1_runner_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmitRunResponse.ProtoReflect.Descriptor instead.
func (*SubmitRunResponse) Descriptor() ([]byte, []int) {
	return file_runner_v1_runner_proto_rawDescGZIP(), []int{5}
}

var File_runner_v1_runner_proto protoreflect.FileDescriptor

var file_runner_v1_runner_proto_rawDesc = []byte{
	0x0a, 0x16, 0x72, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x75, 0x6e, 0x6e,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x74, 0x73, 0x74, 0x72, 0x2e, 0x72,
	0x75, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x16, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe1, 0x03, 0x0a, 0x15, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x0a, 0xfa, 0x42, 0x07, 0x72, 0x05, 0x10, 0x01, 0x18, 0xc8, 0x01, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x85, 0x01, 0x0a, 0x1b, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x5f,
	0x74, 0x65, 0x73, 0x74, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x5f, 0x73, 0x65, 0x6c, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x43, 0x2e, 0x74, 0x73, 0x74,
	0x72, 0x2e, 0x72, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x52, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x54, 0x65, 0x73, 0x74, 0x4c, 0x61, 0x62, 0x65,
	0x6c, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x1b, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x6c, 0x61, 0x62,
	0x65, 0x6c, 0x5f, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x12, 0x85, 0x01, 0x0a,
	0x1b, 0x72, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x6c, 0x61, 0x62,
	0x65, 0x6c, 0x5f, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x43, 0x2e, 0x74, 0x73, 0x74, 0x72, 0x2e, 0x72, 0x75, 0x6e, 0x6e, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x75, 0x6e, 0x6e,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x52, 0x65, 0x6a, 0x65, 0x63, 0x74,
	0x54, 0x65, 0x73, 0x74, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x1b, 0x72, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x5f,
	0x74, 0x65, 0x73, 0x74, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x5f, 0x73, 0x65, 0x6c, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x73, 0x1a, 0x4b, 0x0a, 0x1d, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x54, 0x65,
	0x73, 0x74, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x1a, 0x4b, 0x0a, 0x1d, 0x52, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x54, 0x65, 0x73, 0x74, 0x4c,
	0x61, 0x62, 0x65, 0x6c, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x48,
	0x0a, 0x16, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x75, 0x6e, 0x6e, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x06, 0x72, 0x75, 0x6e, 0x6e,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x74, 0x73, 0x74, 0x72, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x75, 0x6e, 0x6e, 0x65, 0x72,
	0x52, 0x06, 0x72, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x22, 0x2a, 0x0a, 0x0e, 0x4e, 0x65, 0x78, 0x74,
	0x52, 0x75, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0xb0, 0x01, 0x01,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x38, 0x0a, 0x0f, 0x4e, 0x65, 0x78, 0x74, 0x52, 0x75, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x03, 0x72, 0x75, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x74, 0x73, 0x74, 0x72, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x75, 0x6e, 0x52, 0x03, 0x72, 0x75, 0x6e, 0x22, 0xae,
	0x02, 0x0a, 0x10, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x52, 0x75, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x0b, 0xfa, 0x42, 0x08, 0x72, 0x06, 0xb0, 0x01, 0x01, 0xd0, 0x01, 0x01, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x22, 0x0a, 0x06, 0x72, 0x75, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x0b, 0xfa, 0x42, 0x08, 0x72, 0x06, 0xb0, 0x01, 0x01, 0xd0, 0x01, 0x01, 0x52, 0x05, 0x72,
	0x75, 0x6e, 0x49, 0x64, 0x12, 0x32, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x74, 0x73, 0x74, 0x72, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x75, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x2b, 0x0a, 0x04, 0x6c, 0x6f, 0x67, 0x73,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x74, 0x73, 0x74, 0x72, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x75, 0x6e, 0x2e, 0x4c, 0x6f, 0x67, 0x52,
	0x04, 0x6c, 0x6f, 0x67, 0x73, 0x12, 0x3a, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x12, 0x3c, 0x0a, 0x0b, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x0b, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x22,
	0x13, 0x0a, 0x11, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x52, 0x75, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x32, 0x90, 0x02, 0x0a, 0x0d, 0x52, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5f, 0x0a, 0x0e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x52, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x12, 0x25, 0x2e, 0x74, 0x73, 0x74, 0x72, 0x2e,
	0x72, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x52, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x26, 0x2e, 0x74, 0x73, 0x74, 0x72, 0x2e, 0x72, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4a, 0x0a, 0x07, 0x4e, 0x65, 0x78, 0x74, 0x52,
	0x75, 0x6e, 0x12, 0x1e, 0x2e, 0x74, 0x73, 0x74, 0x72, 0x2e, 0x72, 0x75, 0x6e, 0x6e, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x65, 0x78, 0x74, 0x52, 0x75, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x74, 0x73, 0x74, 0x72, 0x2e, 0x72, 0x75, 0x6e, 0x6e, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x65, 0x78, 0x74, 0x52, 0x75, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x52, 0x0a, 0x09, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x52, 0x75, 0x6e,
	0x12, 0x20, 0x2e, 0x74, 0x73, 0x74, 0x72, 0x2e, 0x72, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x52, 0x75, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x21, 0x2e, 0x74, 0x73, 0x74, 0x72, 0x2e, 0x72, 0x75, 0x6e, 0x6e, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x52, 0x75, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x42, 0xac, 0x01, 0x0a, 0x12, 0x63, 0x6f, 0x6d, 0x2e,
	0x74, 0x73, 0x74, 0x72, 0x2e, 0x72, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x0b,
	0x52, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2f, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x61, 0x6e, 0x7a, 0x68, 0x6f,
	0x6e, 0x67, 0x2f, 0x74, 0x73, 0x74, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x75, 0x6e, 0x6e,
	0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x72, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x76, 0x31, 0xa2, 0x02,
	0x03, 0x54, 0x52, 0x58, 0xaa, 0x02, 0x0e, 0x54, 0x73, 0x74, 0x72, 0x2e, 0x52, 0x75, 0x6e, 0x6e,
	0x65, 0x72, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0e, 0x54, 0x73, 0x74, 0x72, 0x5c, 0x52, 0x75, 0x6e,
	0x6e, 0x65, 0x72, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1a, 0x54, 0x73, 0x74, 0x72, 0x5c, 0x52, 0x75,
	0x6e, 0x6e, 0x65, 0x72, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x10, 0x54, 0x73, 0x74, 0x72, 0x3a, 0x3a, 0x52, 0x75, 0x6e, 0x6e,
	0x65, 0x72, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_runner_v1_runner_proto_rawDescOnce sync.Once
	file_runner_v1_runner_proto_rawDescData = file_runner_v1_runner_proto_rawDesc
)

func file_runner_v1_runner_proto_rawDescGZIP() []byte {
	file_runner_v1_runner_proto_rawDescOnce.Do(func() {
		file_runner_v1_runner_proto_rawDescData = protoimpl.X.CompressGZIP(file_runner_v1_runner_proto_rawDescData)
	})
	return file_runner_v1_runner_proto_rawDescData
}

var file_runner_v1_runner_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_runner_v1_runner_proto_goTypes = []interface{}{
	(*RegisterRunnerRequest)(nil),  // 0: tstr.runner.v1.RegisterRunnerRequest
	(*RegisterRunnerResponse)(nil), // 1: tstr.runner.v1.RegisterRunnerResponse
	(*NextRunRequest)(nil),         // 2: tstr.runner.v1.NextRunRequest
	(*NextRunResponse)(nil),        // 3: tstr.runner.v1.NextRunResponse
	(*SubmitRunRequest)(nil),       // 4: tstr.runner.v1.SubmitRunRequest
	(*SubmitRunResponse)(nil),      // 5: tstr.runner.v1.SubmitRunResponse
	nil,                            // 6: tstr.runner.v1.RegisterRunnerRequest.AcceptTestLabelSelectorsEntry
	nil,                            // 7: tstr.runner.v1.RegisterRunnerRequest.RejectTestLabelSelectorsEntry
	(*v1.Runner)(nil),              // 8: tstr.common.v1.Runner
	(*v1.Run)(nil),                 // 9: tstr.common.v1.Run
	(v1.Run_Result)(0),             // 10: tstr.common.v1.Run.Result
	(*v1.Run_Log)(nil),             // 11: tstr.common.v1.Run.Log
	(*timestamppb.Timestamp)(nil),  // 12: google.protobuf.Timestamp
}
var file_runner_v1_runner_proto_depIdxs = []int32{
	6,  // 0: tstr.runner.v1.RegisterRunnerRequest.accept_test_label_selectors:type_name -> tstr.runner.v1.RegisterRunnerRequest.AcceptTestLabelSelectorsEntry
	7,  // 1: tstr.runner.v1.RegisterRunnerRequest.reject_test_label_selectors:type_name -> tstr.runner.v1.RegisterRunnerRequest.RejectTestLabelSelectorsEntry
	8,  // 2: tstr.runner.v1.RegisterRunnerResponse.runner:type_name -> tstr.common.v1.Runner
	9,  // 3: tstr.runner.v1.NextRunResponse.run:type_name -> tstr.common.v1.Run
	10, // 4: tstr.runner.v1.SubmitRunRequest.result:type_name -> tstr.common.v1.Run.Result
	11, // 5: tstr.runner.v1.SubmitRunRequest.logs:type_name -> tstr.common.v1.Run.Log
	12, // 6: tstr.runner.v1.SubmitRunRequest.started_at:type_name -> google.protobuf.Timestamp
	12, // 7: tstr.runner.v1.SubmitRunRequest.finished_at:type_name -> google.protobuf.Timestamp
	0,  // 8: tstr.runner.v1.RunnerService.RegisterRunner:input_type -> tstr.runner.v1.RegisterRunnerRequest
	2,  // 9: tstr.runner.v1.RunnerService.NextRun:input_type -> tstr.runner.v1.NextRunRequest
	4,  // 10: tstr.runner.v1.RunnerService.SubmitRun:input_type -> tstr.runner.v1.SubmitRunRequest
	1,  // 11: tstr.runner.v1.RunnerService.RegisterRunner:output_type -> tstr.runner.v1.RegisterRunnerResponse
	3,  // 12: tstr.runner.v1.RunnerService.NextRun:output_type -> tstr.runner.v1.NextRunResponse
	5,  // 13: tstr.runner.v1.RunnerService.SubmitRun:output_type -> tstr.runner.v1.SubmitRunResponse
	11, // [11:14] is the sub-list for method output_type
	8,  // [8:11] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_runner_v1_runner_proto_init() }
func file_runner_v1_runner_proto_init() {
	if File_runner_v1_runner_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_runner_v1_runner_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterRunnerRequest); i {
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
		file_runner_v1_runner_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterRunnerResponse); i {
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
		file_runner_v1_runner_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NextRunRequest); i {
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
		file_runner_v1_runner_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NextRunResponse); i {
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
		file_runner_v1_runner_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubmitRunRequest); i {
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
		file_runner_v1_runner_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubmitRunResponse); i {
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
			RawDescriptor: file_runner_v1_runner_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_runner_v1_runner_proto_goTypes,
		DependencyIndexes: file_runner_v1_runner_proto_depIdxs,
		MessageInfos:      file_runner_v1_runner_proto_msgTypes,
	}.Build()
	File_runner_v1_runner_proto = out.File
	file_runner_v1_runner_proto_rawDesc = nil
	file_runner_v1_runner_proto_goTypes = nil
	file_runner_v1_runner_proto_depIdxs = nil
}
