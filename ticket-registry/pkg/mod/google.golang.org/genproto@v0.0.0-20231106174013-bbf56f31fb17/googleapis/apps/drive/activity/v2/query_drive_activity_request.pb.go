// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.12
// source: google/apps/drive/activity/v2/query_drive_activity_request.proto

package activity

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// The request message for querying Drive activity.
type QueryDriveActivityRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The primary criteria in the query. The default is
	// ancestorName = `items/root`, if no key is specified.
	//
	// Types that are assignable to Key:
	//
	//	*QueryDriveActivityRequest_ItemName
	//	*QueryDriveActivityRequest_AncestorName
	Key isQueryDriveActivityRequest_Key `protobuf_oneof:"key"`
	// Details on how to consolidate related actions that make up the activity. If
	// not set, then related actions aren't consolidated.
	ConsolidationStrategy *ConsolidationStrategy `protobuf:"bytes,5,opt,name=consolidation_strategy,json=consolidationStrategy,proto3" json:"consolidation_strategy,omitempty"`
	// The minimum number of activities desired in the response; the server
	// attempts to return at least this quantity. The server may also return fewer
	// activities if it has a partial response ready before the request times out.
	// If not set, a default value is used.
	PageSize int32 `protobuf:"varint,6,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// The token identifies which page of results to return. Set this to the
	// next_page_token value returned from a previous query to obtain the
	// following page of results. If not set, the first page of results is
	// returned.
	PageToken string `protobuf:"bytes,7,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	// The filtering for items returned from this query request. The format of the
	// filter string is a sequence of expressions, joined by an optional "AND",
	// where each expression is of the form "field operator value".
	//
	// Supported fields:
	//
	//   - `time`: Uses numerical operators on date values either in
	//     terms of milliseconds since Jan 1, 1970 or in <a
	//     href="https://www.rfc-editor.org/rfc/rfc3339" target="_blank">RFC
	//     3339</a> format. Examples:
	//
	//   - `time > 1452409200000 AND time <= 1492812924310`
	//
	//   - `time >= "2016-01-10T01:02:03-05:00"`
	//
	//   - `detail.action_detail_case`: Uses the "has" operator (:) and
	//     either a singular value or a list of allowed action types enclosed in
	//     parentheses, separated by a space. To exclude a result from the
	//     response, prepend a hyphen (`-`) to the beginning of the filter string.
	//     Examples:
	//
	//   - `detail.action_detail_case:RENAME`
	//
	//   - `detail.action_detail_case:(CREATE RESTORE)`
	//
	//   - `-detail.action_detail_case:MOVE`
	Filter string `protobuf:"bytes,8,opt,name=filter,proto3" json:"filter,omitempty"`
}

func (x *QueryDriveActivityRequest) Reset() {
	*x = QueryDriveActivityRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_apps_drive_activity_v2_query_drive_activity_request_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryDriveActivityRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryDriveActivityRequest) ProtoMessage() {}

func (x *QueryDriveActivityRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_apps_drive_activity_v2_query_drive_activity_request_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryDriveActivityRequest.ProtoReflect.Descriptor instead.
func (*QueryDriveActivityRequest) Descriptor() ([]byte, []int) {
	return file_google_apps_drive_activity_v2_query_drive_activity_request_proto_rawDescGZIP(), []int{0}
}

func (m *QueryDriveActivityRequest) GetKey() isQueryDriveActivityRequest_Key {
	if m != nil {
		return m.Key
	}
	return nil
}

func (x *QueryDriveActivityRequest) GetItemName() string {
	if x, ok := x.GetKey().(*QueryDriveActivityRequest_ItemName); ok {
		return x.ItemName
	}
	return ""
}

func (x *QueryDriveActivityRequest) GetAncestorName() string {
	if x, ok := x.GetKey().(*QueryDriveActivityRequest_AncestorName); ok {
		return x.AncestorName
	}
	return ""
}

func (x *QueryDriveActivityRequest) GetConsolidationStrategy() *ConsolidationStrategy {
	if x != nil {
		return x.ConsolidationStrategy
	}
	return nil
}

func (x *QueryDriveActivityRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *QueryDriveActivityRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

func (x *QueryDriveActivityRequest) GetFilter() string {
	if x != nil {
		return x.Filter
	}
	return ""
}

type isQueryDriveActivityRequest_Key interface {
	isQueryDriveActivityRequest_Key()
}

type QueryDriveActivityRequest_ItemName struct {
	// Return activities for this Drive item. The format is
	// `items/ITEM_ID`.
	ItemName string `protobuf:"bytes,1,opt,name=item_name,json=itemName,proto3,oneof"`
}

type QueryDriveActivityRequest_AncestorName struct {
	// Return activities for this Drive folder, plus all children and
	// descendants. The format is `items/ITEM_ID`.
	AncestorName string `protobuf:"bytes,2,opt,name=ancestor_name,json=ancestorName,proto3,oneof"`
}

func (*QueryDriveActivityRequest_ItemName) isQueryDriveActivityRequest_Key() {}

func (*QueryDriveActivityRequest_AncestorName) isQueryDriveActivityRequest_Key() {}

// How the individual activities are consolidated. If a set of activities is
// related they can be consolidated into one combined activity, such as one
// actor performing the same action on multiple targets, or multiple actors
// performing the same action on a single target. The strategy defines the rules
// for which activities are related.
type ConsolidationStrategy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// How the individual activities are consolidated.
	//
	// Types that are assignable to Strategy:
	//
	//	*ConsolidationStrategy_None
	//	*ConsolidationStrategy_Legacy_
	Strategy isConsolidationStrategy_Strategy `protobuf_oneof:"strategy"`
}

func (x *ConsolidationStrategy) Reset() {
	*x = ConsolidationStrategy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_apps_drive_activity_v2_query_drive_activity_request_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConsolidationStrategy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConsolidationStrategy) ProtoMessage() {}

func (x *ConsolidationStrategy) ProtoReflect() protoreflect.Message {
	mi := &file_google_apps_drive_activity_v2_query_drive_activity_request_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConsolidationStrategy.ProtoReflect.Descriptor instead.
func (*ConsolidationStrategy) Descriptor() ([]byte, []int) {
	return file_google_apps_drive_activity_v2_query_drive_activity_request_proto_rawDescGZIP(), []int{1}
}

func (m *ConsolidationStrategy) GetStrategy() isConsolidationStrategy_Strategy {
	if m != nil {
		return m.Strategy
	}
	return nil
}

func (x *ConsolidationStrategy) GetNone() *ConsolidationStrategy_NoConsolidation {
	if x, ok := x.GetStrategy().(*ConsolidationStrategy_None); ok {
		return x.None
	}
	return nil
}

func (x *ConsolidationStrategy) GetLegacy() *ConsolidationStrategy_Legacy {
	if x, ok := x.GetStrategy().(*ConsolidationStrategy_Legacy_); ok {
		return x.Legacy
	}
	return nil
}

type isConsolidationStrategy_Strategy interface {
	isConsolidationStrategy_Strategy()
}

type ConsolidationStrategy_None struct {
	// The individual activities are not consolidated.
	None *ConsolidationStrategy_NoConsolidation `protobuf:"bytes,1,opt,name=none,proto3,oneof"`
}

type ConsolidationStrategy_Legacy_ struct {
	// The individual activities are consolidated using the legacy strategy.
	Legacy *ConsolidationStrategy_Legacy `protobuf:"bytes,2,opt,name=legacy,proto3,oneof"`
}

func (*ConsolidationStrategy_None) isConsolidationStrategy_Strategy() {}

func (*ConsolidationStrategy_Legacy_) isConsolidationStrategy_Strategy() {}

// A strategy that does no consolidation of individual activities.
type ConsolidationStrategy_NoConsolidation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ConsolidationStrategy_NoConsolidation) Reset() {
	*x = ConsolidationStrategy_NoConsolidation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_apps_drive_activity_v2_query_drive_activity_request_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConsolidationStrategy_NoConsolidation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConsolidationStrategy_NoConsolidation) ProtoMessage() {}

func (x *ConsolidationStrategy_NoConsolidation) ProtoReflect() protoreflect.Message {
	mi := &file_google_apps_drive_activity_v2_query_drive_activity_request_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConsolidationStrategy_NoConsolidation.ProtoReflect.Descriptor instead.
func (*ConsolidationStrategy_NoConsolidation) Descriptor() ([]byte, []int) {
	return file_google_apps_drive_activity_v2_query_drive_activity_request_proto_rawDescGZIP(), []int{1, 0}
}

// A strategy that consolidates activities using the grouping rules from the
// legacy V1 Activity API. Similar actions occurring within a window of time
// can be grouped across multiple targets (such as moving a set of files at
// once) or multiple actors (such as several users editing the same item).
// Grouping rules for this strategy are specific to each type of action.
type ConsolidationStrategy_Legacy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ConsolidationStrategy_Legacy) Reset() {
	*x = ConsolidationStrategy_Legacy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_apps_drive_activity_v2_query_drive_activity_request_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConsolidationStrategy_Legacy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConsolidationStrategy_Legacy) ProtoMessage() {}

func (x *ConsolidationStrategy_Legacy) ProtoReflect() protoreflect.Message {
	mi := &file_google_apps_drive_activity_v2_query_drive_activity_request_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConsolidationStrategy_Legacy.ProtoReflect.Descriptor instead.
func (*ConsolidationStrategy_Legacy) Descriptor() ([]byte, []int) {
	return file_google_apps_drive_activity_v2_query_drive_activity_request_proto_rawDescGZIP(), []int{1, 1}
}

var File_google_apps_drive_activity_v2_query_drive_activity_request_proto protoreflect.FileDescriptor

var file_google_apps_drive_activity_v2_query_drive_activity_request_proto_rawDesc = []byte{
	0x0a, 0x40, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x64, 0x72,
	0x69, 0x76, 0x65, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2f, 0x76, 0x32, 0x2f,
	0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x64, 0x72, 0x69, 0x76, 0x65, 0x5f, 0x61, 0x63, 0x74, 0x69,
	0x76, 0x69, 0x74, 0x79, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x1d, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x73, 0x2e,
	0x64, 0x72, 0x69, 0x76, 0x65, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2e, 0x76,
	0x32, 0x22, 0xa9, 0x02, 0x0a, 0x19, 0x51, 0x75, 0x65, 0x72, 0x79, 0x44, 0x72, 0x69, 0x76, 0x65,
	0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1d, 0x0a, 0x09, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x00, 0x52, 0x08, 0x69, 0x74, 0x65, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x25,
	0x0a, 0x0d, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0c, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x74, 0x6f,
	0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x6b, 0x0a, 0x16, 0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61,
	0x70, 0x70, 0x73, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x65, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69,
	0x74, 0x79, 0x2e, 0x76, 0x32, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x52, 0x15, 0x63, 0x6f, 0x6e,
	0x73, 0x6f, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65,
	0x67, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12,
	0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x16,
	0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x42, 0x05, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x22, 0xf3, 0x01,
	0x0a, 0x15, 0x43, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x12, 0x5a, 0x0a, 0x04, 0x6e, 0x6f, 0x6e, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x44, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61,
	0x70, 0x70, 0x73, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x65, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69,
	0x74, 0x79, 0x2e, 0x76, 0x32, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x2e, 0x4e, 0x6f, 0x43, 0x6f,
	0x6e, 0x73, 0x6f, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x04, 0x6e,
	0x6f, 0x6e, 0x65, 0x12, 0x55, 0x0a, 0x06, 0x6c, 0x65, 0x67, 0x61, 0x63, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x3b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x70,
	0x73, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x65, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79,
	0x2e, 0x76, 0x32, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x2e, 0x4c, 0x65, 0x67, 0x61, 0x63, 0x79,
	0x48, 0x00, 0x52, 0x06, 0x6c, 0x65, 0x67, 0x61, 0x63, 0x79, 0x1a, 0x11, 0x0a, 0x0f, 0x4e, 0x6f,
	0x43, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x08, 0x0a,
	0x06, 0x4c, 0x65, 0x67, 0x61, 0x63, 0x79, 0x42, 0x0a, 0x0a, 0x08, 0x73, 0x74, 0x72, 0x61, 0x74,
	0x65, 0x67, 0x79, 0x42, 0xd3, 0x01, 0x0a, 0x21, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x73, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x65, 0x2e, 0x61, 0x63,
	0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x32, 0x42, 0x1e, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x44, 0x72, 0x69, 0x76, 0x65, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x45, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67,
	0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70,
	0x69, 0x73, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x64, 0x72, 0x69, 0x76, 0x65, 0x2f, 0x61, 0x63,
	0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2f, 0x76, 0x32, 0x3b, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69,
	0x74, 0x79, 0xa2, 0x02, 0x04, 0x47, 0x41, 0x44, 0x41, 0xaa, 0x02, 0x1d, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x41, 0x70, 0x70, 0x73, 0x2e, 0x44, 0x72, 0x69, 0x76, 0x65, 0x2e, 0x41, 0x63,
	0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2e, 0x56, 0x32, 0xca, 0x02, 0x1d, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x5c, 0x41, 0x70, 0x70, 0x73, 0x5c, 0x44, 0x72, 0x69, 0x76, 0x65, 0x5c, 0x41, 0x63,
	0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x5c, 0x56, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_google_apps_drive_activity_v2_query_drive_activity_request_proto_rawDescOnce sync.Once
	file_google_apps_drive_activity_v2_query_drive_activity_request_proto_rawDescData = file_google_apps_drive_activity_v2_query_drive_activity_request_proto_rawDesc
)

func file_google_apps_drive_activity_v2_query_drive_activity_request_proto_rawDescGZIP() []byte {
	file_google_apps_drive_activity_v2_query_drive_activity_request_proto_rawDescOnce.Do(func() {
		file_google_apps_drive_activity_v2_query_drive_activity_request_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_apps_drive_activity_v2_query_drive_activity_request_proto_rawDescData)
	})
	return file_google_apps_drive_activity_v2_query_drive_activity_request_proto_rawDescData
}

var file_google_apps_drive_activity_v2_query_drive_activity_request_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_google_apps_drive_activity_v2_query_drive_activity_request_proto_goTypes = []interface{}{
	(*QueryDriveActivityRequest)(nil),             // 0: google.apps.drive.activity.v2.QueryDriveActivityRequest
	(*ConsolidationStrategy)(nil),                 // 1: google.apps.drive.activity.v2.ConsolidationStrategy
	(*ConsolidationStrategy_NoConsolidation)(nil), // 2: google.apps.drive.activity.v2.ConsolidationStrategy.NoConsolidation
	(*ConsolidationStrategy_Legacy)(nil),          // 3: google.apps.drive.activity.v2.ConsolidationStrategy.Legacy
}
var file_google_apps_drive_activity_v2_query_drive_activity_request_proto_depIdxs = []int32{
	1, // 0: google.apps.drive.activity.v2.QueryDriveActivityRequest.consolidation_strategy:type_name -> google.apps.drive.activity.v2.ConsolidationStrategy
	2, // 1: google.apps.drive.activity.v2.ConsolidationStrategy.none:type_name -> google.apps.drive.activity.v2.ConsolidationStrategy.NoConsolidation
	3, // 2: google.apps.drive.activity.v2.ConsolidationStrategy.legacy:type_name -> google.apps.drive.activity.v2.ConsolidationStrategy.Legacy
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_google_apps_drive_activity_v2_query_drive_activity_request_proto_init() }
func file_google_apps_drive_activity_v2_query_drive_activity_request_proto_init() {
	if File_google_apps_drive_activity_v2_query_drive_activity_request_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_apps_drive_activity_v2_query_drive_activity_request_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryDriveActivityRequest); i {
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
		file_google_apps_drive_activity_v2_query_drive_activity_request_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConsolidationStrategy); i {
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
		file_google_apps_drive_activity_v2_query_drive_activity_request_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConsolidationStrategy_NoConsolidation); i {
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
		file_google_apps_drive_activity_v2_query_drive_activity_request_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConsolidationStrategy_Legacy); i {
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
	file_google_apps_drive_activity_v2_query_drive_activity_request_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*QueryDriveActivityRequest_ItemName)(nil),
		(*QueryDriveActivityRequest_AncestorName)(nil),
	}
	file_google_apps_drive_activity_v2_query_drive_activity_request_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*ConsolidationStrategy_None)(nil),
		(*ConsolidationStrategy_Legacy_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_apps_drive_activity_v2_query_drive_activity_request_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_apps_drive_activity_v2_query_drive_activity_request_proto_goTypes,
		DependencyIndexes: file_google_apps_drive_activity_v2_query_drive_activity_request_proto_depIdxs,
		MessageInfos:      file_google_apps_drive_activity_v2_query_drive_activity_request_proto_msgTypes,
	}.Build()
	File_google_apps_drive_activity_v2_query_drive_activity_request_proto = out.File
	file_google_apps_drive_activity_v2_query_drive_activity_request_proto_rawDesc = nil
	file_google_apps_drive_activity_v2_query_drive_activity_request_proto_goTypes = nil
	file_google_apps_drive_activity_v2_query_drive_activity_request_proto_depIdxs = nil
}
