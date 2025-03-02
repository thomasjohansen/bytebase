// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: v1/sheet_service.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
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

// Type of the SheetPayload.
type SheetPayload_Type int32

const (
	SheetPayload_TYPE_UNSPECIFIED SheetPayload_Type = 0
	SheetPayload_SCHEMA_DESIGN    SheetPayload_Type = 1
)

// Enum value maps for SheetPayload_Type.
var (
	SheetPayload_Type_name = map[int32]string{
		0: "TYPE_UNSPECIFIED",
		1: "SCHEMA_DESIGN",
	}
	SheetPayload_Type_value = map[string]int32{
		"TYPE_UNSPECIFIED": 0,
		"SCHEMA_DESIGN":    1,
	}
)

func (x SheetPayload_Type) Enum() *SheetPayload_Type {
	p := new(SheetPayload_Type)
	*p = x
	return p
}

func (x SheetPayload_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SheetPayload_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_v1_sheet_service_proto_enumTypes[0].Descriptor()
}

func (SheetPayload_Type) Type() protoreflect.EnumType {
	return &file_v1_sheet_service_proto_enumTypes[0]
}

func (x SheetPayload_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SheetPayload_Type.Descriptor instead.
func (SheetPayload_Type) EnumDescriptor() ([]byte, []int) {
	return file_v1_sheet_service_proto_rawDescGZIP(), []int{5, 0}
}

type CreateSheetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The parent resource where this sheet will be created.
	// Format: projects/{project}
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// The sheet to create.
	Sheet *Sheet `protobuf:"bytes,2,opt,name=sheet,proto3" json:"sheet,omitempty"`
}

func (x *CreateSheetRequest) Reset() {
	*x = CreateSheetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_sheet_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSheetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSheetRequest) ProtoMessage() {}

func (x *CreateSheetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_sheet_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSheetRequest.ProtoReflect.Descriptor instead.
func (*CreateSheetRequest) Descriptor() ([]byte, []int) {
	return file_v1_sheet_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateSheetRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *CreateSheetRequest) GetSheet() *Sheet {
	if x != nil {
		return x.Sheet
	}
	return nil
}

type GetSheetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the sheet to retrieve.
	// Format: projects/{project}/sheets/{sheet}
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// By default, the content of the sheet is cut off, set the `raw` to true to retrieve the full content.
	Raw bool `protobuf:"varint,2,opt,name=raw,proto3" json:"raw,omitempty"`
}

func (x *GetSheetRequest) Reset() {
	*x = GetSheetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_sheet_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSheetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSheetRequest) ProtoMessage() {}

func (x *GetSheetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_sheet_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSheetRequest.ProtoReflect.Descriptor instead.
func (*GetSheetRequest) Descriptor() ([]byte, []int) {
	return file_v1_sheet_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetSheetRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetSheetRequest) GetRaw() bool {
	if x != nil {
		return x.Raw
	}
	return false
}

type UpdateSheetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The sheet to update.
	//
	// The sheet's `name` field is used to identify the sheet to update.
	// Format: projects/{project}/sheets/{sheet}
	Sheet *Sheet `protobuf:"bytes,1,opt,name=sheet,proto3" json:"sheet,omitempty"`
	// The list of fields to be updated.
	// Fields are specified relative to the sheet.
	// (e.g. `title`, `statement`; *not* `sheet.title` or `sheet.statement`)
	// Only support update the following fields for now:
	// - `title`
	// - `statement`
	// - `starred`
	// - `visibility`
	UpdateMask *fieldmaskpb.FieldMask `protobuf:"bytes,2,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
}

func (x *UpdateSheetRequest) Reset() {
	*x = UpdateSheetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_sheet_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateSheetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSheetRequest) ProtoMessage() {}

func (x *UpdateSheetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_sheet_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSheetRequest.ProtoReflect.Descriptor instead.
func (*UpdateSheetRequest) Descriptor() ([]byte, []int) {
	return file_v1_sheet_service_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateSheetRequest) GetSheet() *Sheet {
	if x != nil {
		return x.Sheet
	}
	return nil
}

func (x *UpdateSheetRequest) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

type DeleteSheetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the sheet to delete.
	// Format: projects/{project}/sheets/{sheet}
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *DeleteSheetRequest) Reset() {
	*x = DeleteSheetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_sheet_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteSheetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSheetRequest) ProtoMessage() {}

func (x *DeleteSheetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_sheet_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSheetRequest.ProtoReflect.Descriptor instead.
func (*DeleteSheetRequest) Descriptor() ([]byte, []int) {
	return file_v1_sheet_service_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteSheetRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Sheet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the sheet resource, generated by the server.
	// Canonical parent is project.
	// Format: projects/{project}/sheets/{sheet}
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The database resource name.
	// Format: instances/{instance}/databases/{database}
	// If the database parent doesn't exist, the database field is empty.
	Database string `protobuf:"bytes,2,opt,name=database,proto3" json:"database,omitempty"`
	// The title of the sheet.
	Title string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	// The creator of the Sheet.
	// Format: users/{email}
	Creator string `protobuf:"bytes,4,opt,name=creator,proto3" json:"creator,omitempty"`
	// The create time of the sheet.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// The last update time of the sheet.
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// The content of the sheet.
	// By default, it will be cut off, if it doesn't match the `content_size`, you can
	// set the `raw` to true in GetSheet request to retrieve the full content.
	Content []byte `protobuf:"bytes,7,opt,name=content,proto3" json:"content,omitempty"`
	// content_size is the full size of the content, may not match the size of the `content` field.
	ContentSize int64         `protobuf:"varint,8,opt,name=content_size,json=contentSize,proto3" json:"content_size,omitempty"`
	Payload     *SheetPayload `protobuf:"bytes,13,opt,name=payload,proto3" json:"payload,omitempty"`
	PushEvent   *PushEvent    `protobuf:"bytes,14,opt,name=push_event,json=pushEvent,proto3" json:"push_event,omitempty"`
}

func (x *Sheet) Reset() {
	*x = Sheet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_sheet_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sheet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sheet) ProtoMessage() {}

func (x *Sheet) ProtoReflect() protoreflect.Message {
	mi := &file_v1_sheet_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sheet.ProtoReflect.Descriptor instead.
func (*Sheet) Descriptor() ([]byte, []int) {
	return file_v1_sheet_service_proto_rawDescGZIP(), []int{4}
}

func (x *Sheet) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Sheet) GetDatabase() string {
	if x != nil {
		return x.Database
	}
	return ""
}

func (x *Sheet) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Sheet) GetCreator() string {
	if x != nil {
		return x.Creator
	}
	return ""
}

func (x *Sheet) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *Sheet) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *Sheet) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

func (x *Sheet) GetContentSize() int64 {
	if x != nil {
		return x.ContentSize
	}
	return 0
}

func (x *Sheet) GetPayload() *SheetPayload {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *Sheet) GetPushEvent() *PushEvent {
	if x != nil {
		return x.PushEvent
	}
	return nil
}

type SheetPayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type SheetPayload_Type `protobuf:"varint,1,opt,name=type,proto3,enum=bytebase.v1.SheetPayload_Type" json:"type,omitempty"`
	// The snapshot of the database config when creating the sheet, be used to compare with the baseline_database_config and apply the diff to the database.
	DatabaseConfig *DatabaseConfig `protobuf:"bytes,2,opt,name=database_config,json=databaseConfig,proto3" json:"database_config,omitempty"`
	// The snapshot of the baseline database config when creating the sheet.
	BaselineDatabaseConfig *DatabaseConfig `protobuf:"bytes,3,opt,name=baseline_database_config,json=baselineDatabaseConfig,proto3" json:"baseline_database_config,omitempty"`
}

func (x *SheetPayload) Reset() {
	*x = SheetPayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_sheet_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SheetPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SheetPayload) ProtoMessage() {}

func (x *SheetPayload) ProtoReflect() protoreflect.Message {
	mi := &file_v1_sheet_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SheetPayload.ProtoReflect.Descriptor instead.
func (*SheetPayload) Descriptor() ([]byte, []int) {
	return file_v1_sheet_service_proto_rawDescGZIP(), []int{5}
}

func (x *SheetPayload) GetType() SheetPayload_Type {
	if x != nil {
		return x.Type
	}
	return SheetPayload_TYPE_UNSPECIFIED
}

func (x *SheetPayload) GetDatabaseConfig() *DatabaseConfig {
	if x != nil {
		return x.DatabaseConfig
	}
	return nil
}

func (x *SheetPayload) GetBaselineDatabaseConfig() *DatabaseConfig {
	if x != nil {
		return x.BaselineDatabaseConfig
	}
	return nil
}

var File_v1_sheet_service_proto protoreflect.FileDescriptor

var file_v1_sheet_service_proto_rawDesc = []byte{
	0x0a, 0x16, 0x76, 0x31, 0x2f, 0x73, 0x68, 0x65, 0x65, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61,
	0x73, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62,
	0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65,
	0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x76,
	0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x76, 0x31, 0x2f, 0x76, 0x63, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x60, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x53, 0x68, 0x65, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x06,
	0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41,
	0x02, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x2d, 0x0a, 0x05, 0x73, 0x68, 0x65,
	0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62,
	0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x68, 0x65, 0x65, 0x74, 0x42, 0x03, 0xe0, 0x41,
	0x02, 0x52, 0x05, 0x73, 0x68, 0x65, 0x65, 0x74, 0x22, 0x3c, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x53,
	0x68, 0x65, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x61, 0x77, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x03, 0x72, 0x61, 0x77, 0x22, 0x80, 0x01, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x53, 0x68, 0x65, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a,
	0x05, 0x73, 0x68, 0x65, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x62,
	0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x68, 0x65, 0x65, 0x74,
	0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x05, 0x73, 0x68, 0x65, 0x65, 0x74, 0x12, 0x3b, 0x0a, 0x0b,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x52, 0x0a, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x73, 0x6b, 0x22, 0x2d, 0x0a, 0x12, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x53, 0x68, 0x65, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0,
	0x41, 0x02, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xb0, 0x03, 0x0a, 0x05, 0x53, 0x68, 0x65,
	0x65, 0x74, 0x12, 0x1a, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x06, 0xe0, 0x41, 0x02, 0xe0, 0x41, 0x05, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1d, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x07, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x6f, 0x72, 0x12, 0x40, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x40, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0a, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x26, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x42, 0x03, 0xe0,
	0x41, 0x03, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x53, 0x69, 0x7a, 0x65, 0x12,
	0x33, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x68, 0x65, 0x65, 0x74, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x07, 0x70, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x12, 0x35, 0x0a, 0x0a, 0x70, 0x75, 0x73, 0x68, 0x5f, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62,
	0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x75, 0x73, 0x68, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x52, 0x09, 0x70, 0x75, 0x73, 0x68, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x22, 0x90, 0x02, 0x0a, 0x0c,
	0x53, 0x68, 0x65, 0x65, 0x74, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x32, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x62, 0x79, 0x74,
	0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x68, 0x65, 0x65, 0x74, 0x50, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x44, 0x0a, 0x0f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x62, 0x79, 0x74, 0x65,
	0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x55, 0x0a, 0x18, 0x62, 0x61, 0x73, 0x65, 0x6c, 0x69,
	0x6e, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62,
	0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x16, 0x62, 0x61, 0x73, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x44,
	0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x2f, 0x0a,
	0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x10, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x53,
	0x43, 0x48, 0x45, 0x4d, 0x41, 0x5f, 0x44, 0x45, 0x53, 0x49, 0x47, 0x4e, 0x10, 0x01, 0x32, 0x83,
	0x04, 0x0a, 0x0c, 0x53, 0x68, 0x65, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x80, 0x01, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x68, 0x65, 0x65, 0x74, 0x12,
	0x1f, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x53, 0x68, 0x65, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x12, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x68, 0x65, 0x65, 0x74, 0x22, 0x3c, 0xda, 0x41, 0x0c, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x2c,
	0x73, 0x68, 0x65, 0x65, 0x74, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x27, 0x3a, 0x05, 0x73, 0x68, 0x65,
	0x65, 0x74, 0x22, 0x1e, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x3d,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x7d, 0x2f, 0x73, 0x68, 0x65, 0x65,
	0x74, 0x73, 0x12, 0x6b, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x53, 0x68, 0x65, 0x65, 0x74, 0x12, 0x1c,
	0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x53, 0x68, 0x65, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x62,
	0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x68, 0x65, 0x65, 0x74,
	0x22, 0x2d, 0xda, 0x41, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20, 0x12,
	0x1e, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x73, 0x68, 0x65, 0x65, 0x74, 0x73, 0x2f, 0x2a, 0x7d, 0x12,
	0x8b, 0x01, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x68, 0x65, 0x65, 0x74, 0x12,
	0x1f, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x53, 0x68, 0x65, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x12, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x68, 0x65, 0x65, 0x74, 0x22, 0x47, 0xda, 0x41, 0x11, 0x73, 0x68, 0x65, 0x65, 0x74, 0x2c, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2d,
	0x3a, 0x05, 0x73, 0x68, 0x65, 0x65, 0x74, 0x32, 0x24, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x73, 0x68,
	0x65, 0x65, 0x74, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x73, 0x2f, 0x2a, 0x2f, 0x73, 0x68, 0x65, 0x65, 0x74, 0x73, 0x2f, 0x2a, 0x7d, 0x12, 0x75, 0x0a,
	0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x68, 0x65, 0x65, 0x74, 0x12, 0x1f, 0x2e, 0x62,
	0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x53, 0x68, 0x65, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x2d, 0xda, 0x41, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x20, 0x2a, 0x1e, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x3d,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x73, 0x68, 0x65, 0x65, 0x74,
	0x73, 0x2f, 0x2a, 0x7d, 0x42, 0x11, 0x5a, 0x0f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65,
	0x64, 0x2d, 0x67, 0x6f, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_sheet_service_proto_rawDescOnce sync.Once
	file_v1_sheet_service_proto_rawDescData = file_v1_sheet_service_proto_rawDesc
)

func file_v1_sheet_service_proto_rawDescGZIP() []byte {
	file_v1_sheet_service_proto_rawDescOnce.Do(func() {
		file_v1_sheet_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_sheet_service_proto_rawDescData)
	})
	return file_v1_sheet_service_proto_rawDescData
}

var file_v1_sheet_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_v1_sheet_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_v1_sheet_service_proto_goTypes = []interface{}{
	(SheetPayload_Type)(0),        // 0: bytebase.v1.SheetPayload.Type
	(*CreateSheetRequest)(nil),    // 1: bytebase.v1.CreateSheetRequest
	(*GetSheetRequest)(nil),       // 2: bytebase.v1.GetSheetRequest
	(*UpdateSheetRequest)(nil),    // 3: bytebase.v1.UpdateSheetRequest
	(*DeleteSheetRequest)(nil),    // 4: bytebase.v1.DeleteSheetRequest
	(*Sheet)(nil),                 // 5: bytebase.v1.Sheet
	(*SheetPayload)(nil),          // 6: bytebase.v1.SheetPayload
	(*fieldmaskpb.FieldMask)(nil), // 7: google.protobuf.FieldMask
	(*timestamppb.Timestamp)(nil), // 8: google.protobuf.Timestamp
	(*PushEvent)(nil),             // 9: bytebase.v1.PushEvent
	(*DatabaseConfig)(nil),        // 10: bytebase.v1.DatabaseConfig
	(*emptypb.Empty)(nil),         // 11: google.protobuf.Empty
}
var file_v1_sheet_service_proto_depIdxs = []int32{
	5,  // 0: bytebase.v1.CreateSheetRequest.sheet:type_name -> bytebase.v1.Sheet
	5,  // 1: bytebase.v1.UpdateSheetRequest.sheet:type_name -> bytebase.v1.Sheet
	7,  // 2: bytebase.v1.UpdateSheetRequest.update_mask:type_name -> google.protobuf.FieldMask
	8,  // 3: bytebase.v1.Sheet.create_time:type_name -> google.protobuf.Timestamp
	8,  // 4: bytebase.v1.Sheet.update_time:type_name -> google.protobuf.Timestamp
	6,  // 5: bytebase.v1.Sheet.payload:type_name -> bytebase.v1.SheetPayload
	9,  // 6: bytebase.v1.Sheet.push_event:type_name -> bytebase.v1.PushEvent
	0,  // 7: bytebase.v1.SheetPayload.type:type_name -> bytebase.v1.SheetPayload.Type
	10, // 8: bytebase.v1.SheetPayload.database_config:type_name -> bytebase.v1.DatabaseConfig
	10, // 9: bytebase.v1.SheetPayload.baseline_database_config:type_name -> bytebase.v1.DatabaseConfig
	1,  // 10: bytebase.v1.SheetService.CreateSheet:input_type -> bytebase.v1.CreateSheetRequest
	2,  // 11: bytebase.v1.SheetService.GetSheet:input_type -> bytebase.v1.GetSheetRequest
	3,  // 12: bytebase.v1.SheetService.UpdateSheet:input_type -> bytebase.v1.UpdateSheetRequest
	4,  // 13: bytebase.v1.SheetService.DeleteSheet:input_type -> bytebase.v1.DeleteSheetRequest
	5,  // 14: bytebase.v1.SheetService.CreateSheet:output_type -> bytebase.v1.Sheet
	5,  // 15: bytebase.v1.SheetService.GetSheet:output_type -> bytebase.v1.Sheet
	5,  // 16: bytebase.v1.SheetService.UpdateSheet:output_type -> bytebase.v1.Sheet
	11, // 17: bytebase.v1.SheetService.DeleteSheet:output_type -> google.protobuf.Empty
	14, // [14:18] is the sub-list for method output_type
	10, // [10:14] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_v1_sheet_service_proto_init() }
func file_v1_sheet_service_proto_init() {
	if File_v1_sheet_service_proto != nil {
		return
	}
	file_v1_database_service_proto_init()
	file_v1_vcs_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_v1_sheet_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSheetRequest); i {
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
		file_v1_sheet_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSheetRequest); i {
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
		file_v1_sheet_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateSheetRequest); i {
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
		file_v1_sheet_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteSheetRequest); i {
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
		file_v1_sheet_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Sheet); i {
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
		file_v1_sheet_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SheetPayload); i {
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
			RawDescriptor: file_v1_sheet_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_sheet_service_proto_goTypes,
		DependencyIndexes: file_v1_sheet_service_proto_depIdxs,
		EnumInfos:         file_v1_sheet_service_proto_enumTypes,
		MessageInfos:      file_v1_sheet_service_proto_msgTypes,
	}.Build()
	File_v1_sheet_service_proto = out.File
	file_v1_sheet_service_proto_rawDesc = nil
	file_v1_sheet_service_proto_goTypes = nil
	file_v1_sheet_service_proto_depIdxs = nil
}
