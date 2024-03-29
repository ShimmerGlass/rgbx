// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.15.6
// source: rgbx/rgb.proto

package rgbx

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

type DeviceType int32

const (
	DeviceType_KEYBOARD DeviceType = 0
)

// Enum value maps for DeviceType.
var (
	DeviceType_name = map[int32]string{
		0: "KEYBOARD",
	}
	DeviceType_value = map[string]int32{
		"KEYBOARD": 0,
	}
)

func (x DeviceType) Enum() *DeviceType {
	p := new(DeviceType)
	*p = x
	return p
}

func (x DeviceType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DeviceType) Descriptor() protoreflect.EnumDescriptor {
	return file_rgbx_rgb_proto_enumTypes[0].Descriptor()
}

func (DeviceType) Type() protoreflect.EnumType {
	return &file_rgbx_rgb_proto_enumTypes[0]
}

func (x DeviceType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DeviceType.Descriptor instead.
func (DeviceType) EnumDescriptor() ([]byte, []int) {
	return file_rgbx_rgb_proto_rawDescGZIP(), []int{0}
}

type SetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Priority   uint32     `protobuf:"varint,1,opt,name=priority,proto3" json:"priority,omitempty"`
	DurationMs int32      `protobuf:"varint,2,opt,name=duration_ms,json=durationMs,proto3" json:"duration_ms,omitempty"`
	DeviceType DeviceType `protobuf:"varint,3,opt,name=device_type,json=deviceType,proto3,enum=rgbx.DeviceType" json:"device_type,omitempty"`
	// Types that are assignable to Effect:
	//	*SetRequest_Static
	//	*SetRequest_Matrix
	//	*SetRequest_Wave
	//	*SetRequest_Nightsky
	//	*SetRequest_K2000
	//	*SetRequest_Progress
	//	*SetRequest_VUMeter
	Effect isSetRequest_Effect `protobuf_oneof:"effect"`
}

func (x *SetRequest) Reset() {
	*x = SetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rgbx_rgb_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetRequest) ProtoMessage() {}

func (x *SetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rgbx_rgb_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetRequest.ProtoReflect.Descriptor instead.
func (*SetRequest) Descriptor() ([]byte, []int) {
	return file_rgbx_rgb_proto_rawDescGZIP(), []int{0}
}

func (x *SetRequest) GetPriority() uint32 {
	if x != nil {
		return x.Priority
	}
	return 0
}

func (x *SetRequest) GetDurationMs() int32 {
	if x != nil {
		return x.DurationMs
	}
	return 0
}

func (x *SetRequest) GetDeviceType() DeviceType {
	if x != nil {
		return x.DeviceType
	}
	return DeviceType_KEYBOARD
}

func (m *SetRequest) GetEffect() isSetRequest_Effect {
	if m != nil {
		return m.Effect
	}
	return nil
}

func (x *SetRequest) GetStatic() *EffectStatic {
	if x, ok := x.GetEffect().(*SetRequest_Static); ok {
		return x.Static
	}
	return nil
}

func (x *SetRequest) GetMatrix() *EffectMatrix {
	if x, ok := x.GetEffect().(*SetRequest_Matrix); ok {
		return x.Matrix
	}
	return nil
}

func (x *SetRequest) GetWave() *EffectWave {
	if x, ok := x.GetEffect().(*SetRequest_Wave); ok {
		return x.Wave
	}
	return nil
}

func (x *SetRequest) GetNightsky() *EffectNightSky {
	if x, ok := x.GetEffect().(*SetRequest_Nightsky); ok {
		return x.Nightsky
	}
	return nil
}

func (x *SetRequest) GetK2000() *EffectK2000 {
	if x, ok := x.GetEffect().(*SetRequest_K2000); ok {
		return x.K2000
	}
	return nil
}

func (x *SetRequest) GetProgress() *EffectProgress {
	if x, ok := x.GetEffect().(*SetRequest_Progress); ok {
		return x.Progress
	}
	return nil
}

func (x *SetRequest) GetVUMeter() *EffectVUMeter {
	if x, ok := x.GetEffect().(*SetRequest_VUMeter); ok {
		return x.VUMeter
	}
	return nil
}

type isSetRequest_Effect interface {
	isSetRequest_Effect()
}

type SetRequest_Static struct {
	Static *EffectStatic `protobuf:"bytes,100,opt,name=static,proto3,oneof"`
}

type SetRequest_Matrix struct {
	Matrix *EffectMatrix `protobuf:"bytes,101,opt,name=matrix,proto3,oneof"`
}

type SetRequest_Wave struct {
	Wave *EffectWave `protobuf:"bytes,102,opt,name=wave,proto3,oneof"`
}

type SetRequest_Nightsky struct {
	Nightsky *EffectNightSky `protobuf:"bytes,103,opt,name=nightsky,proto3,oneof"`
}

type SetRequest_K2000 struct {
	K2000 *EffectK2000 `protobuf:"bytes,104,opt,name=k2000,proto3,oneof"`
}

type SetRequest_Progress struct {
	Progress *EffectProgress `protobuf:"bytes,105,opt,name=progress,proto3,oneof"`
}

type SetRequest_VUMeter struct {
	VUMeter *EffectVUMeter `protobuf:"bytes,106,opt,name=VUMeter,proto3,oneof"`
}

func (*SetRequest_Static) isSetRequest_Effect() {}

func (*SetRequest_Matrix) isSetRequest_Effect() {}

func (*SetRequest_Wave) isSetRequest_Effect() {}

func (*SetRequest_Nightsky) isSetRequest_Effect() {}

func (*SetRequest_K2000) isSetRequest_Effect() {}

func (*SetRequest_Progress) isSetRequest_Effect() {}

func (*SetRequest_VUMeter) isSetRequest_Effect() {}

type RemoveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Priority   uint32     `protobuf:"varint,1,opt,name=priority,proto3" json:"priority,omitempty"`
	DeviceType DeviceType `protobuf:"varint,2,opt,name=device_type,json=deviceType,proto3,enum=rgbx.DeviceType" json:"device_type,omitempty"`
}

func (x *RemoveRequest) Reset() {
	*x = RemoveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rgbx_rgb_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveRequest) ProtoMessage() {}

func (x *RemoveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rgbx_rgb_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveRequest.ProtoReflect.Descriptor instead.
func (*RemoveRequest) Descriptor() ([]byte, []int) {
	return file_rgbx_rgb_proto_rawDescGZIP(), []int{1}
}

func (x *RemoveRequest) GetPriority() uint32 {
	if x != nil {
		return x.Priority
	}
	return 0
}

func (x *RemoveRequest) GetDeviceType() DeviceType {
	if x != nil {
		return x.DeviceType
	}
	return DeviceType_KEYBOARD
}

type SuccessResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *SuccessResponse) Reset() {
	*x = SuccessResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rgbx_rgb_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SuccessResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SuccessResponse) ProtoMessage() {}

func (x *SuccessResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rgbx_rgb_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SuccessResponse.ProtoReflect.Descriptor instead.
func (*SuccessResponse) Descriptor() ([]byte, []int) {
	return file_rgbx_rgb_proto_rawDescGZIP(), []int{2}
}

func (x *SuccessResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type EffectStatic struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Color *Color `protobuf:"bytes,1,opt,name=color,proto3" json:"color,omitempty"`
}

func (x *EffectStatic) Reset() {
	*x = EffectStatic{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rgbx_rgb_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EffectStatic) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EffectStatic) ProtoMessage() {}

func (x *EffectStatic) ProtoReflect() protoreflect.Message {
	mi := &file_rgbx_rgb_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EffectStatic.ProtoReflect.Descriptor instead.
func (*EffectStatic) Descriptor() ([]byte, []int) {
	return file_rgbx_rgb_proto_rawDescGZIP(), []int{3}
}

func (x *EffectStatic) GetColor() *Color {
	if x != nil {
		return x.Color
	}
	return nil
}

type EffectNightSky struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Color1 *Color `protobuf:"bytes,1,opt,name=color1,proto3" json:"color1,omitempty"`
	Color2 *Color `protobuf:"bytes,2,opt,name=color2,proto3" json:"color2,omitempty"`
}

func (x *EffectNightSky) Reset() {
	*x = EffectNightSky{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rgbx_rgb_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EffectNightSky) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EffectNightSky) ProtoMessage() {}

func (x *EffectNightSky) ProtoReflect() protoreflect.Message {
	mi := &file_rgbx_rgb_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EffectNightSky.ProtoReflect.Descriptor instead.
func (*EffectNightSky) Descriptor() ([]byte, []int) {
	return file_rgbx_rgb_proto_rawDescGZIP(), []int{4}
}

func (x *EffectNightSky) GetColor1() *Color {
	if x != nil {
		return x.Color1
	}
	return nil
}

func (x *EffectNightSky) GetColor2() *Color {
	if x != nil {
		return x.Color2
	}
	return nil
}

type EffectMatrix struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rows []*Row `protobuf:"bytes,1,rep,name=rows,proto3" json:"rows,omitempty"`
}

func (x *EffectMatrix) Reset() {
	*x = EffectMatrix{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rgbx_rgb_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EffectMatrix) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EffectMatrix) ProtoMessage() {}

func (x *EffectMatrix) ProtoReflect() protoreflect.Message {
	mi := &file_rgbx_rgb_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EffectMatrix.ProtoReflect.Descriptor instead.
func (*EffectMatrix) Descriptor() ([]byte, []int) {
	return file_rgbx_rgb_proto_rawDescGZIP(), []int{5}
}

func (x *EffectMatrix) GetRows() []*Row {
	if x != nil {
		return x.Rows
	}
	return nil
}

type EffectWave struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EffectWave) Reset() {
	*x = EffectWave{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rgbx_rgb_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EffectWave) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EffectWave) ProtoMessage() {}

func (x *EffectWave) ProtoReflect() protoreflect.Message {
	mi := &file_rgbx_rgb_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EffectWave.ProtoReflect.Descriptor instead.
func (*EffectWave) Descriptor() ([]byte, []int) {
	return file_rgbx_rgb_proto_rawDescGZIP(), []int{6}
}

type EffectK2000 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Color *Color `protobuf:"bytes,1,opt,name=color,proto3" json:"color,omitempty"`
	Row   int32  `protobuf:"varint,2,opt,name=row,proto3" json:"row,omitempty"`
}

func (x *EffectK2000) Reset() {
	*x = EffectK2000{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rgbx_rgb_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EffectK2000) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EffectK2000) ProtoMessage() {}

func (x *EffectK2000) ProtoReflect() protoreflect.Message {
	mi := &file_rgbx_rgb_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EffectK2000.ProtoReflect.Descriptor instead.
func (*EffectK2000) Descriptor() ([]byte, []int) {
	return file_rgbx_rgb_proto_rawDescGZIP(), []int{7}
}

func (x *EffectK2000) GetColor() *Color {
	if x != nil {
		return x.Color
	}
	return nil
}

func (x *EffectK2000) GetRow() int32 {
	if x != nil {
		return x.Row
	}
	return 0
}

type EffectProgress struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Color *Color  `protobuf:"bytes,1,opt,name=color,proto3" json:"color,omitempty"`
	Value float64 `protobuf:"fixed64,2,opt,name=value,proto3" json:"value,omitempty"`
	Rows  []int32 `protobuf:"varint,3,rep,packed,name=rows,proto3" json:"rows,omitempty"`
}

func (x *EffectProgress) Reset() {
	*x = EffectProgress{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rgbx_rgb_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EffectProgress) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EffectProgress) ProtoMessage() {}

func (x *EffectProgress) ProtoReflect() protoreflect.Message {
	mi := &file_rgbx_rgb_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EffectProgress.ProtoReflect.Descriptor instead.
func (*EffectProgress) Descriptor() ([]byte, []int) {
	return file_rgbx_rgb_proto_rawDescGZIP(), []int{8}
}

func (x *EffectProgress) GetColor() *Color {
	if x != nil {
		return x.Color
	}
	return nil
}

func (x *EffectProgress) GetValue() float64 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *EffectProgress) GetRows() []int32 {
	if x != nil {
		return x.Rows
	}
	return nil
}

type EffectVUMeter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Color1 *Color `protobuf:"bytes,1,opt,name=color1,proto3" json:"color1,omitempty"`
	Color2 *Color `protobuf:"bytes,2,opt,name=color2,proto3" json:"color2,omitempty"`
}

func (x *EffectVUMeter) Reset() {
	*x = EffectVUMeter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rgbx_rgb_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EffectVUMeter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EffectVUMeter) ProtoMessage() {}

func (x *EffectVUMeter) ProtoReflect() protoreflect.Message {
	mi := &file_rgbx_rgb_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EffectVUMeter.ProtoReflect.Descriptor instead.
func (*EffectVUMeter) Descriptor() ([]byte, []int) {
	return file_rgbx_rgb_proto_rawDescGZIP(), []int{9}
}

func (x *EffectVUMeter) GetColor1() *Color {
	if x != nil {
		return x.Color1
	}
	return nil
}

func (x *EffectVUMeter) GetColor2() *Color {
	if x != nil {
		return x.Color2
	}
	return nil
}

type Row struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*Color `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *Row) Reset() {
	*x = Row{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rgbx_rgb_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Row) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Row) ProtoMessage() {}

func (x *Row) ProtoReflect() protoreflect.Message {
	mi := &file_rgbx_rgb_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Row.ProtoReflect.Descriptor instead.
func (*Row) Descriptor() ([]byte, []int) {
	return file_rgbx_rgb_proto_rawDescGZIP(), []int{10}
}

func (x *Row) GetData() []*Color {
	if x != nil {
		return x.Data
	}
	return nil
}

type Color struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	R int32 `protobuf:"varint,1,opt,name=R,proto3" json:"R,omitempty"`
	G int32 `protobuf:"varint,2,opt,name=G,proto3" json:"G,omitempty"`
	B int32 `protobuf:"varint,3,opt,name=B,proto3" json:"B,omitempty"`
}

func (x *Color) Reset() {
	*x = Color{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rgbx_rgb_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Color) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Color) ProtoMessage() {}

func (x *Color) ProtoReflect() protoreflect.Message {
	mi := &file_rgbx_rgb_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Color.ProtoReflect.Descriptor instead.
func (*Color) Descriptor() ([]byte, []int) {
	return file_rgbx_rgb_proto_rawDescGZIP(), []int{11}
}

func (x *Color) GetR() int32 {
	if x != nil {
		return x.R
	}
	return 0
}

func (x *Color) GetG() int32 {
	if x != nil {
		return x.G
	}
	return 0
}

func (x *Color) GetB() int32 {
	if x != nil {
		return x.B
	}
	return 0
}

var File_rgbx_rgb_proto protoreflect.FileDescriptor

var file_rgbx_rgb_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x72, 0x67, 0x62, 0x78, 0x2f, 0x72, 0x67, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x04, 0x72, 0x67, 0x62, 0x78, 0x22, 0xce, 0x03, 0x0a, 0x0a, 0x53, 0x65, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74,
	0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x4d, 0x73, 0x12, 0x31, 0x0a, 0x0b, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x72, 0x67, 0x62, 0x78, 0x2e, 0x44,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2c, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x18,
	0x64, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x72, 0x67, 0x62, 0x78, 0x2e, 0x45, 0x66, 0x66,
	0x65, 0x63, 0x74, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x48, 0x00, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x69, 0x63, 0x12, 0x2c, 0x0a, 0x06, 0x6d, 0x61, 0x74, 0x72, 0x69, 0x78, 0x18, 0x65, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x72, 0x67, 0x62, 0x78, 0x2e, 0x45, 0x66, 0x66, 0x65, 0x63,
	0x74, 0x4d, 0x61, 0x74, 0x72, 0x69, 0x78, 0x48, 0x00, 0x52, 0x06, 0x6d, 0x61, 0x74, 0x72, 0x69,
	0x78, 0x12, 0x26, 0x0a, 0x04, 0x77, 0x61, 0x76, 0x65, 0x18, 0x66, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x72, 0x67, 0x62, 0x78, 0x2e, 0x45, 0x66, 0x66, 0x65, 0x63, 0x74, 0x57, 0x61, 0x76,
	0x65, 0x48, 0x00, 0x52, 0x04, 0x77, 0x61, 0x76, 0x65, 0x12, 0x32, 0x0a, 0x08, 0x6e, 0x69, 0x67,
	0x68, 0x74, 0x73, 0x6b, 0x79, 0x18, 0x67, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x72, 0x67,
	0x62, 0x78, 0x2e, 0x45, 0x66, 0x66, 0x65, 0x63, 0x74, 0x4e, 0x69, 0x67, 0x68, 0x74, 0x53, 0x6b,
	0x79, 0x48, 0x00, 0x52, 0x08, 0x6e, 0x69, 0x67, 0x68, 0x74, 0x73, 0x6b, 0x79, 0x12, 0x29, 0x0a,
	0x05, 0x6b, 0x32, 0x30, 0x30, 0x30, 0x18, 0x68, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x72,
	0x67, 0x62, 0x78, 0x2e, 0x45, 0x66, 0x66, 0x65, 0x63, 0x74, 0x4b, 0x32, 0x30, 0x30, 0x30, 0x48,
	0x00, 0x52, 0x05, 0x6b, 0x32, 0x30, 0x30, 0x30, 0x12, 0x32, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x67,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x69, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x72, 0x67, 0x62,
	0x78, 0x2e, 0x45, 0x66, 0x66, 0x65, 0x63, 0x74, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73,
	0x48, 0x00, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x2f, 0x0a, 0x07,
	0x56, 0x55, 0x4d, 0x65, 0x74, 0x65, 0x72, 0x18, 0x6a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e,
	0x72, 0x67, 0x62, 0x78, 0x2e, 0x45, 0x66, 0x66, 0x65, 0x63, 0x74, 0x56, 0x55, 0x4d, 0x65, 0x74,
	0x65, 0x72, 0x48, 0x00, 0x52, 0x07, 0x56, 0x55, 0x4d, 0x65, 0x74, 0x65, 0x72, 0x42, 0x08, 0x0a,
	0x06, 0x65, 0x66, 0x66, 0x65, 0x63, 0x74, 0x22, 0x5e, 0x0a, 0x0d, 0x52, 0x65, 0x6d, 0x6f, 0x76,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x69, 0x6f,
	0x72, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x70, 0x72, 0x69, 0x6f,
	0x72, 0x69, 0x74, 0x79, 0x12, 0x31, 0x0a, 0x0b, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x72, 0x67, 0x62, 0x78,
	0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x22, 0x2b, 0x0a, 0x0f, 0x53, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x22, 0x31, 0x0a, 0x0c, 0x45, 0x66, 0x66, 0x65, 0x63, 0x74, 0x53, 0x74,
	0x61, 0x74, 0x69, 0x63, 0x12, 0x21, 0x0a, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x72, 0x67, 0x62, 0x78, 0x2e, 0x43, 0x6f, 0x6c, 0x6f, 0x72,
	0x52, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x22, 0x5a, 0x0a, 0x0e, 0x45, 0x66, 0x66, 0x65, 0x63,
	0x74, 0x4e, 0x69, 0x67, 0x68, 0x74, 0x53, 0x6b, 0x79, 0x12, 0x23, 0x0a, 0x06, 0x63, 0x6f, 0x6c,
	0x6f, 0x72, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x72, 0x67, 0x62, 0x78,
	0x2e, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x52, 0x06, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x31, 0x12, 0x23,
	0x0a, 0x06, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b,
	0x2e, 0x72, 0x67, 0x62, 0x78, 0x2e, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x52, 0x06, 0x63, 0x6f, 0x6c,
	0x6f, 0x72, 0x32, 0x22, 0x2d, 0x0a, 0x0c, 0x45, 0x66, 0x66, 0x65, 0x63, 0x74, 0x4d, 0x61, 0x74,
	0x72, 0x69, 0x78, 0x12, 0x1d, 0x0a, 0x04, 0x72, 0x6f, 0x77, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x09, 0x2e, 0x72, 0x67, 0x62, 0x78, 0x2e, 0x52, 0x6f, 0x77, 0x52, 0x04, 0x72, 0x6f,
	0x77, 0x73, 0x22, 0x0c, 0x0a, 0x0a, 0x45, 0x66, 0x66, 0x65, 0x63, 0x74, 0x57, 0x61, 0x76, 0x65,
	0x22, 0x42, 0x0a, 0x0b, 0x45, 0x66, 0x66, 0x65, 0x63, 0x74, 0x4b, 0x32, 0x30, 0x30, 0x30, 0x12,
	0x21, 0x0a, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b,
	0x2e, 0x72, 0x67, 0x62, 0x78, 0x2e, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x52, 0x05, 0x63, 0x6f, 0x6c,
	0x6f, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x6f, 0x77, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x03, 0x72, 0x6f, 0x77, 0x22, 0x5d, 0x0a, 0x0e, 0x45, 0x66, 0x66, 0x65, 0x63, 0x74, 0x50, 0x72,
	0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x21, 0x0a, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x72, 0x67, 0x62, 0x78, 0x2e, 0x43, 0x6f, 0x6c,
	0x6f, 0x72, 0x52, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x72, 0x6f, 0x77, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x05, 0x52, 0x04, 0x72,
	0x6f, 0x77, 0x73, 0x22, 0x59, 0x0a, 0x0d, 0x45, 0x66, 0x66, 0x65, 0x63, 0x74, 0x56, 0x55, 0x4d,
	0x65, 0x74, 0x65, 0x72, 0x12, 0x23, 0x0a, 0x06, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x31, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x72, 0x67, 0x62, 0x78, 0x2e, 0x43, 0x6f, 0x6c, 0x6f,
	0x72, 0x52, 0x06, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x31, 0x12, 0x23, 0x0a, 0x06, 0x63, 0x6f, 0x6c,
	0x6f, 0x72, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x72, 0x67, 0x62, 0x78,
	0x2e, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x52, 0x06, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x32, 0x22, 0x26,
	0x0a, 0x03, 0x52, 0x6f, 0x77, 0x12, 0x1f, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x72, 0x67, 0x62, 0x78, 0x2e, 0x43, 0x6f, 0x6c, 0x6f, 0x72,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x31, 0x0a, 0x05, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x12,
	0x0c, 0x0a, 0x01, 0x52, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x52, 0x12, 0x0c, 0x0a,
	0x01, 0x47, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x47, 0x12, 0x0c, 0x0a, 0x01, 0x42,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x42, 0x2a, 0x1a, 0x0a, 0x0a, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0c, 0x0a, 0x08, 0x4b, 0x45, 0x59, 0x42, 0x4f,
	0x41, 0x52, 0x44, 0x10, 0x00, 0x32, 0x73, 0x0a, 0x07, 0x52, 0x47, 0x42, 0x69, 0x7a, 0x65, 0x72,
	0x12, 0x30, 0x0a, 0x03, 0x53, 0x65, 0x74, 0x12, 0x10, 0x2e, 0x72, 0x67, 0x62, 0x78, 0x2e, 0x53,
	0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x72, 0x67, 0x62, 0x78,
	0x2e, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x36, 0x0a, 0x06, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x12, 0x13, 0x2e, 0x72,
	0x67, 0x62, 0x78, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x15, 0x2e, 0x72, 0x67, 0x62, 0x78, 0x2e, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x23, 0x5a, 0x21, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x68, 0x69, 0x6d, 0x6d, 0x65, 0x72,
	0x67, 0x6c, 0x61, 0x73, 0x73, 0x2f, 0x72, 0x67, 0x62, 0x78, 0x2f, 0x72, 0x67, 0x62, 0x78, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rgbx_rgb_proto_rawDescOnce sync.Once
	file_rgbx_rgb_proto_rawDescData = file_rgbx_rgb_proto_rawDesc
)

func file_rgbx_rgb_proto_rawDescGZIP() []byte {
	file_rgbx_rgb_proto_rawDescOnce.Do(func() {
		file_rgbx_rgb_proto_rawDescData = protoimpl.X.CompressGZIP(file_rgbx_rgb_proto_rawDescData)
	})
	return file_rgbx_rgb_proto_rawDescData
}

var file_rgbx_rgb_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_rgbx_rgb_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_rgbx_rgb_proto_goTypes = []interface{}{
	(DeviceType)(0),         // 0: rgbx.DeviceType
	(*SetRequest)(nil),      // 1: rgbx.SetRequest
	(*RemoveRequest)(nil),   // 2: rgbx.RemoveRequest
	(*SuccessResponse)(nil), // 3: rgbx.SuccessResponse
	(*EffectStatic)(nil),    // 4: rgbx.EffectStatic
	(*EffectNightSky)(nil),  // 5: rgbx.EffectNightSky
	(*EffectMatrix)(nil),    // 6: rgbx.EffectMatrix
	(*EffectWave)(nil),      // 7: rgbx.EffectWave
	(*EffectK2000)(nil),     // 8: rgbx.EffectK2000
	(*EffectProgress)(nil),  // 9: rgbx.EffectProgress
	(*EffectVUMeter)(nil),   // 10: rgbx.EffectVUMeter
	(*Row)(nil),             // 11: rgbx.Row
	(*Color)(nil),           // 12: rgbx.Color
}
var file_rgbx_rgb_proto_depIdxs = []int32{
	0,  // 0: rgbx.SetRequest.device_type:type_name -> rgbx.DeviceType
	4,  // 1: rgbx.SetRequest.static:type_name -> rgbx.EffectStatic
	6,  // 2: rgbx.SetRequest.matrix:type_name -> rgbx.EffectMatrix
	7,  // 3: rgbx.SetRequest.wave:type_name -> rgbx.EffectWave
	5,  // 4: rgbx.SetRequest.nightsky:type_name -> rgbx.EffectNightSky
	8,  // 5: rgbx.SetRequest.k2000:type_name -> rgbx.EffectK2000
	9,  // 6: rgbx.SetRequest.progress:type_name -> rgbx.EffectProgress
	10, // 7: rgbx.SetRequest.VUMeter:type_name -> rgbx.EffectVUMeter
	0,  // 8: rgbx.RemoveRequest.device_type:type_name -> rgbx.DeviceType
	12, // 9: rgbx.EffectStatic.color:type_name -> rgbx.Color
	12, // 10: rgbx.EffectNightSky.color1:type_name -> rgbx.Color
	12, // 11: rgbx.EffectNightSky.color2:type_name -> rgbx.Color
	11, // 12: rgbx.EffectMatrix.rows:type_name -> rgbx.Row
	12, // 13: rgbx.EffectK2000.color:type_name -> rgbx.Color
	12, // 14: rgbx.EffectProgress.color:type_name -> rgbx.Color
	12, // 15: rgbx.EffectVUMeter.color1:type_name -> rgbx.Color
	12, // 16: rgbx.EffectVUMeter.color2:type_name -> rgbx.Color
	12, // 17: rgbx.Row.data:type_name -> rgbx.Color
	1,  // 18: rgbx.RGBizer.Set:input_type -> rgbx.SetRequest
	2,  // 19: rgbx.RGBizer.Remove:input_type -> rgbx.RemoveRequest
	3,  // 20: rgbx.RGBizer.Set:output_type -> rgbx.SuccessResponse
	3,  // 21: rgbx.RGBizer.Remove:output_type -> rgbx.SuccessResponse
	20, // [20:22] is the sub-list for method output_type
	18, // [18:20] is the sub-list for method input_type
	18, // [18:18] is the sub-list for extension type_name
	18, // [18:18] is the sub-list for extension extendee
	0,  // [0:18] is the sub-list for field type_name
}

func init() { file_rgbx_rgb_proto_init() }
func file_rgbx_rgb_proto_init() {
	if File_rgbx_rgb_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rgbx_rgb_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetRequest); i {
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
		file_rgbx_rgb_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveRequest); i {
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
		file_rgbx_rgb_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SuccessResponse); i {
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
		file_rgbx_rgb_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EffectStatic); i {
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
		file_rgbx_rgb_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EffectNightSky); i {
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
		file_rgbx_rgb_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EffectMatrix); i {
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
		file_rgbx_rgb_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EffectWave); i {
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
		file_rgbx_rgb_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EffectK2000); i {
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
		file_rgbx_rgb_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EffectProgress); i {
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
		file_rgbx_rgb_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EffectVUMeter); i {
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
		file_rgbx_rgb_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Row); i {
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
		file_rgbx_rgb_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Color); i {
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
	file_rgbx_rgb_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*SetRequest_Static)(nil),
		(*SetRequest_Matrix)(nil),
		(*SetRequest_Wave)(nil),
		(*SetRequest_Nightsky)(nil),
		(*SetRequest_K2000)(nil),
		(*SetRequest_Progress)(nil),
		(*SetRequest_VUMeter)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_rgbx_rgb_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rgbx_rgb_proto_goTypes,
		DependencyIndexes: file_rgbx_rgb_proto_depIdxs,
		EnumInfos:         file_rgbx_rgb_proto_enumTypes,
		MessageInfos:      file_rgbx_rgb_proto_msgTypes,
	}.Build()
	File_rgbx_rgb_proto = out.File
	file_rgbx_rgb_proto_rawDesc = nil
	file_rgbx_rgb_proto_goTypes = nil
	file_rgbx_rgb_proto_depIdxs = nil
}
