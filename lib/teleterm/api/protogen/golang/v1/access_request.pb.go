// Copyright 2021 Gravitational, Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        (unknown)
// source: v1/access_request.proto

package v1

import (
	proto "github.com/golang/protobuf/proto"
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

// Database describes a database
type AccessRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                 string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	State              string                 `protobuf:"bytes,2,opt,name=state,proto3" json:"state,omitempty"`
	ResolveReason      string                 `protobuf:"bytes,3,opt,name=resolve_reason,json=resolveReason,proto3" json:"resolve_reason,omitempty"`
	RequestReason      string                 `protobuf:"bytes,4,opt,name=request_reason,json=requestReason,proto3" json:"request_reason,omitempty"`
	User               string                 `protobuf:"bytes,5,opt,name=user,proto3" json:"user,omitempty"`
	Roles              []string               `protobuf:"bytes,6,rep,name=roles,proto3" json:"roles,omitempty"`
	Created            string                 `protobuf:"bytes,7,opt,name=created,proto3" json:"created,omitempty"`
	Expires            string                 `protobuf:"bytes,8,opt,name=expires,proto3" json:"expires,omitempty"`
	Reviews            []*AccessRequestReview `protobuf:"bytes,9,rep,name=reviews,proto3" json:"reviews,omitempty"`
	SuggestedReviewers []string               `protobuf:"bytes,10,rep,name=suggested_reviewers,json=suggestedReviewers,proto3" json:"suggested_reviewers,omitempty"`
	ThresholdNames     []string               `protobuf:"bytes,11,rep,name=threshold_names,json=thresholdNames,proto3" json:"threshold_names,omitempty"`
	ResourceIds        []*ResourceID          `protobuf:"bytes,12,rep,name=resource_ids,json=resourceIds,proto3" json:"resource_ids,omitempty"`
}

func (x *AccessRequest) Reset() {
	*x = AccessRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_access_request_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccessRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccessRequest) ProtoMessage() {}

func (x *AccessRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_access_request_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccessRequest.ProtoReflect.Descriptor instead.
func (*AccessRequest) Descriptor() ([]byte, []int) {
	return file_v1_access_request_proto_rawDescGZIP(), []int{0}
}

func (x *AccessRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AccessRequest) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *AccessRequest) GetResolveReason() string {
	if x != nil {
		return x.ResolveReason
	}
	return ""
}

func (x *AccessRequest) GetRequestReason() string {
	if x != nil {
		return x.RequestReason
	}
	return ""
}

func (x *AccessRequest) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *AccessRequest) GetRoles() []string {
	if x != nil {
		return x.Roles
	}
	return nil
}

func (x *AccessRequest) GetCreated() string {
	if x != nil {
		return x.Created
	}
	return ""
}

func (x *AccessRequest) GetExpires() string {
	if x != nil {
		return x.Expires
	}
	return ""
}

func (x *AccessRequest) GetReviews() []*AccessRequestReview {
	if x != nil {
		return x.Reviews
	}
	return nil
}

func (x *AccessRequest) GetSuggestedReviewers() []string {
	if x != nil {
		return x.SuggestedReviewers
	}
	return nil
}

func (x *AccessRequest) GetThresholdNames() []string {
	if x != nil {
		return x.ThresholdNames
	}
	return nil
}

func (x *AccessRequest) GetResourceIds() []*ResourceID {
	if x != nil {
		return x.ResourceIds
	}
	return nil
}

type AccessRequestReview struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Author  string   `protobuf:"bytes,1,opt,name=author,proto3" json:"author,omitempty"`
	Roles   []string `protobuf:"bytes,2,rep,name=roles,proto3" json:"roles,omitempty"`
	State   string   `protobuf:"bytes,3,opt,name=state,proto3" json:"state,omitempty"`
	Reason  string   `protobuf:"bytes,4,opt,name=reason,proto3" json:"reason,omitempty"`
	Created string   `protobuf:"bytes,5,opt,name=created,proto3" json:"created,omitempty"`
}

func (x *AccessRequestReview) Reset() {
	*x = AccessRequestReview{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_access_request_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccessRequestReview) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccessRequestReview) ProtoMessage() {}

func (x *AccessRequestReview) ProtoReflect() protoreflect.Message {
	mi := &file_v1_access_request_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccessRequestReview.ProtoReflect.Descriptor instead.
func (*AccessRequestReview) Descriptor() ([]byte, []int) {
	return file_v1_access_request_proto_rawDescGZIP(), []int{1}
}

func (x *AccessRequestReview) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *AccessRequestReview) GetRoles() []string {
	if x != nil {
		return x.Roles
	}
	return nil
}

func (x *AccessRequestReview) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *AccessRequestReview) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

func (x *AccessRequestReview) GetCreated() string {
	if x != nil {
		return x.Created
	}
	return ""
}

type ResourceID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Kind        string `protobuf:"bytes,1,opt,name=kind,proto3" json:"kind,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	ClusterName string `protobuf:"bytes,3,opt,name=cluster_name,json=clusterName,proto3" json:"cluster_name,omitempty"`
}

func (x *ResourceID) Reset() {
	*x = ResourceID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_access_request_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResourceID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceID) ProtoMessage() {}

func (x *ResourceID) ProtoReflect() protoreflect.Message {
	mi := &file_v1_access_request_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceID.ProtoReflect.Descriptor instead.
func (*ResourceID) Descriptor() ([]byte, []int) {
	return file_v1_access_request_proto_rawDescGZIP(), []int{2}
}

func (x *ResourceID) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *ResourceID) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ResourceID) GetClusterName() string {
	if x != nil {
		return x.ClusterName
	}
	return ""
}

var File_v1_access_request_proto protoreflect.FileDescriptor

var file_v1_access_request_proto_rawDesc = []byte{
	0x0a, 0x17, 0x76, 0x31, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x74, 0x65, 0x6c, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x2e, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x22,
	0xc5, 0x03, 0x0a, 0x0d, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x72, 0x65, 0x73, 0x6f, 0x6c,
	0x76, 0x65, 0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x25,
	0x0a, 0x0e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52,
	0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x6f, 0x6c,
	0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x78, 0x70,
	0x69, 0x72, 0x65, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x78, 0x70, 0x69,
	0x72, 0x65, 0x73, 0x12, 0x43, 0x0a, 0x07, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x18, 0x09,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e,
	0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52,
	0x07, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x12, 0x2f, 0x0a, 0x13, 0x73, 0x75, 0x67, 0x67,
	0x65, 0x73, 0x74, 0x65, 0x64, 0x5f, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x65, 0x72, 0x73, 0x18,
	0x0a, 0x20, 0x03, 0x28, 0x09, 0x52, 0x12, 0x73, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x65, 0x64,
	0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x65, 0x72, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x74, 0x68, 0x72,
	0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x0b, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x0e, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x4e, 0x61, 0x6d,
	0x65, 0x73, 0x12, 0x43, 0x0a, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69,
	0x64, 0x73, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x2e, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x2e,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x44, 0x52, 0x0b, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x73, 0x22, 0x8b, 0x01, 0x0a, 0x13, 0x41, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x12,
	0x16, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x12, 0x14, 0x0a,
	0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x22, 0x57, 0x0a, 0x0a, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x42, 0x33,
	0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x72, 0x61,
	0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x6c, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x2f, 0x6c, 0x69, 0x62, 0x2f, 0x74, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x72, 0x6d,
	0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_access_request_proto_rawDescOnce sync.Once
	file_v1_access_request_proto_rawDescData = file_v1_access_request_proto_rawDesc
)

func file_v1_access_request_proto_rawDescGZIP() []byte {
	file_v1_access_request_proto_rawDescOnce.Do(func() {
		file_v1_access_request_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_access_request_proto_rawDescData)
	})
	return file_v1_access_request_proto_rawDescData
}

var file_v1_access_request_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_v1_access_request_proto_goTypes = []interface{}{
	(*AccessRequest)(nil),       // 0: teleport.terminal.v1.AccessRequest
	(*AccessRequestReview)(nil), // 1: teleport.terminal.v1.AccessRequestReview
	(*ResourceID)(nil),          // 2: teleport.terminal.v1.ResourceID
}
var file_v1_access_request_proto_depIdxs = []int32{
	1, // 0: teleport.terminal.v1.AccessRequest.reviews:type_name -> teleport.terminal.v1.AccessRequestReview
	2, // 1: teleport.terminal.v1.AccessRequest.resource_ids:type_name -> teleport.terminal.v1.ResourceID
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_v1_access_request_proto_init() }
func file_v1_access_request_proto_init() {
	if File_v1_access_request_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_access_request_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccessRequest); i {
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
		file_v1_access_request_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccessRequestReview); i {
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
		file_v1_access_request_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResourceID); i {
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
			RawDescriptor: file_v1_access_request_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v1_access_request_proto_goTypes,
		DependencyIndexes: file_v1_access_request_proto_depIdxs,
		MessageInfos:      file_v1_access_request_proto_msgTypes,
	}.Build()
	File_v1_access_request_proto = out.File
	file_v1_access_request_proto_rawDesc = nil
	file_v1_access_request_proto_goTypes = nil
	file_v1_access_request_proto_depIdxs = nil
}
