// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.7.0
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

type Frame_DeviceType int32

const (
	Frame_KEYBOARD Frame_DeviceType = 0
)

// Enum value maps for Frame_DeviceType.
var (
	Frame_DeviceType_name = map[int32]string{
		0: "KEYBOARD",
	}
	Frame_DeviceType_value = map[string]int32{
		"KEYBOARD": 0,
	}
)

func (x Frame_DeviceType) Enum() *Frame_DeviceType {
	p := new(Frame_DeviceType)
	*p = x
	return p
}

func (x Frame_DeviceType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Frame_DeviceType) Descriptor() protoreflect.EnumDescriptor {
	return file_rgbx_rgb_proto_enumTypes[0].Descriptor()
}

func (Frame_DeviceType) Type() protoreflect.EnumType {
	return &file_rgbx_rgb_proto_enumTypes[0]
}

func (x Frame_DeviceType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Frame_DeviceType.Descriptor instead.
func (Frame_DeviceType) EnumDescriptor() ([]byte, []int) {
	return file_rgbx_rgb_proto_rawDescGZIP(), []int{0, 0}
}

type Frame struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Priority   uint32           `protobuf:"varint,2,opt,name=priority,proto3" json:"priority,omitempty"`
	DurationMs int32            `protobuf:"varint,3,opt,name=duration_ms,json=durationMs,proto3" json:"duration_ms,omitempty"`
	DeviceType Frame_DeviceType `protobuf:"varint,4,opt,name=device_type,json=deviceType,proto3,enum=rgbx.Frame_DeviceType" json:"device_type,omitempty"`
	// Types that are assignable to Effect:
	//	*Frame_Static
	//	*Frame_Matrix
	Effect isFrame_Effect `protobuf_oneof:"effect"`
}

func (x *Frame) Reset() {
	*x = Frame{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rgbx_rgb_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Frame) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Frame) ProtoMessage() {}

func (x *Frame) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Frame.ProtoReflect.Descriptor instead.
func (*Frame) Descriptor() ([]byte, []int) {
	return file_rgbx_rgb_proto_rawDescGZIP(), []int{0}
}

func (x *Frame) GetPriority() uint32 {
	if x != nil {
		return x.Priority
	}
	return 0
}

func (x *Frame) GetDurationMs() int32 {
	if x != nil {
		return x.DurationMs
	}
	return 0
}

func (x *Frame) GetDeviceType() Frame_DeviceType {
	if x != nil {
		return x.DeviceType
	}
	return Frame_KEYBOARD
}

func (m *Frame) GetEffect() isFrame_Effect {
	if m != nil {
		return m.Effect
	}
	return nil
}

func (x *Frame) GetStatic() *EffectStatic {
	if x, ok := x.GetEffect().(*Frame_Static); ok {
		return x.Static
	}
	return nil
}

func (x *Frame) GetMatrix() *EffectMatrix {
	if x, ok := x.GetEffect().(*Frame_Matrix); ok {
		return x.Matrix
	}
	return nil
}

type isFrame_Effect interface {
	isFrame_Effect()
}

type Frame_Static struct {
	Static *EffectStatic `protobuf:"bytes,100,opt,name=static,proto3,oneof"`
}

type Frame_Matrix struct {
	Matrix *EffectMatrix `protobuf:"bytes,101,opt,name=matrix,proto3,oneof"`
}

func (*Frame_Static) isFrame_Effect() {}

func (*Frame_Matrix) isFrame_Effect() {}

type SetFrameResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *SetFrameResponse) Reset() {
	*x = SetFrameResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rgbx_rgb_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetFrameResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetFrameResponse) ProtoMessage() {}

func (x *SetFrameResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use SetFrameResponse.ProtoReflect.Descriptor instead.
func (*SetFrameResponse) Descriptor() ([]byte, []int) {
	return file_rgbx_rgb_proto_rawDescGZIP(), []int{1}
}

func (x *SetFrameResponse) GetSuccess() bool {
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
		mi := &file_rgbx_rgb_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EffectStatic) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EffectStatic) ProtoMessage() {}

func (x *EffectStatic) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use EffectStatic.ProtoReflect.Descriptor instead.
func (*EffectStatic) Descriptor() ([]byte, []int) {
	return file_rgbx_rgb_proto_rawDescGZIP(), []int{2}
}

func (x *EffectStatic) GetColor() *Color {
	if x != nil {
		return x.Color
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
		mi := &file_rgbx_rgb_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EffectMatrix) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EffectMatrix) ProtoMessage() {}

func (x *EffectMatrix) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use EffectMatrix.ProtoReflect.Descriptor instead.
func (*EffectMatrix) Descriptor() ([]byte, []int) {
	return file_rgbx_rgb_proto_rawDescGZIP(), []int{3}
}

func (x *EffectMatrix) GetRows() []*Row {
	if x != nil {
		return x.Rows
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
		mi := &file_rgbx_rgb_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Row) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Row) ProtoMessage() {}

func (x *Row) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Row.ProtoReflect.Descriptor instead.
func (*Row) Descriptor() ([]byte, []int) {
	return file_rgbx_rgb_proto_rawDescGZIP(), []int{4}
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
		mi := &file_rgbx_rgb_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Color) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Color) ProtoMessage() {}

func (x *Color) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Color.ProtoReflect.Descriptor instead.
func (*Color) Descriptor() ([]byte, []int) {
	return file_rgbx_rgb_proto_rawDescGZIP(), []int{5}
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
	0x0a, 0x11, 0x72, 0x67, 0x62, 0x69, 0x7a, 0x65, 0x72, 0x2f, 0x72, 0x67, 0x62, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x07, 0x72, 0x67, 0x62, 0x69, 0x7a, 0x65, 0x72, 0x22, 0x88, 0x02, 0x0a,
	0x05, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69,
	0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69,
	0x74, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x4d, 0x73, 0x12, 0x3a, 0x0a, 0x0b, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x72, 0x67, 0x62, 0x69, 0x7a,
	0x65, 0x72, 0x2e, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x0a, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x2f, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x15, 0x2e, 0x72, 0x67, 0x62, 0x69, 0x7a, 0x65, 0x72, 0x2e, 0x45, 0x66, 0x66, 0x65, 0x63, 0x74,
	0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x48, 0x00, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63,
	0x12, 0x2f, 0x0a, 0x06, 0x6d, 0x61, 0x74, 0x72, 0x69, 0x78, 0x18, 0x65, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x15, 0x2e, 0x72, 0x67, 0x62, 0x69, 0x7a, 0x65, 0x72, 0x2e, 0x45, 0x66, 0x66, 0x65, 0x63,
	0x74, 0x4d, 0x61, 0x74, 0x72, 0x69, 0x78, 0x48, 0x00, 0x52, 0x06, 0x6d, 0x61, 0x74, 0x72, 0x69,
	0x78, 0x22, 0x1a, 0x0a, 0x0a, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x0c, 0x0a, 0x08, 0x4b, 0x45, 0x59, 0x42, 0x4f, 0x41, 0x52, 0x44, 0x10, 0x00, 0x42, 0x08, 0x0a,
	0x06, 0x65, 0x66, 0x66, 0x65, 0x63, 0x74, 0x22, 0x2c, 0x0a, 0x10, 0x53, 0x65, 0x74, 0x46, 0x72,
	0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x34, 0x0a, 0x0c, 0x45, 0x66, 0x66, 0x65, 0x63, 0x74, 0x53,
	0x74, 0x61, 0x74, 0x69, 0x63, 0x12, 0x24, 0x0a, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x72, 0x67, 0x62, 0x69, 0x7a, 0x65, 0x72, 0x2e, 0x43,
	0x6f, 0x6c, 0x6f, 0x72, 0x52, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x22, 0x30, 0x0a, 0x0c, 0x45,
	0x66, 0x66, 0x65, 0x63, 0x74, 0x4d, 0x61, 0x74, 0x72, 0x69, 0x78, 0x12, 0x20, 0x0a, 0x04, 0x72,
	0x6f, 0x77, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x72, 0x67, 0x62, 0x69,
	0x7a, 0x65, 0x72, 0x2e, 0x52, 0x6f, 0x77, 0x52, 0x04, 0x72, 0x6f, 0x77, 0x73, 0x22, 0x29, 0x0a,
	0x03, 0x52, 0x6f, 0x77, 0x12, 0x22, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x72, 0x67, 0x62, 0x69, 0x7a, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6c,
	0x6f, 0x72, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x31, 0x0a, 0x05, 0x43, 0x6f, 0x6c, 0x6f,
	0x72, 0x12, 0x0c, 0x0a, 0x01, 0x52, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x52, 0x12,
	0x0c, 0x0a, 0x01, 0x47, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x47, 0x12, 0x0c, 0x0a,
	0x01, 0x42, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x42, 0x32, 0x42, 0x0a, 0x07, 0x52,
	0x47, 0x42, 0x69, 0x7a, 0x65, 0x72, 0x12, 0x37, 0x0a, 0x08, 0x53, 0x65, 0x74, 0x46, 0x72, 0x61,
	0x6d, 0x65, 0x12, 0x0e, 0x2e, 0x72, 0x67, 0x62, 0x69, 0x7a, 0x65, 0x72, 0x2e, 0x46, 0x72, 0x61,
	0x6d, 0x65, 0x1a, 0x19, 0x2e, 0x72, 0x67, 0x62, 0x69, 0x7a, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x74,
	0x46, 0x72, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x68,
	0x69, 0x6d, 0x6d, 0x65, 0x72, 0x67, 0x6c, 0x61, 0x73, 0x73, 0x2f, 0x72, 0x67, 0x62, 0x69, 0x7a,
	0x65, 0x72, 0x2f, 0x72, 0x67, 0x62, 0x69, 0x7a, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
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
var file_rgbx_rgb_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_rgbx_rgb_proto_goTypes = []interface{}{
	(Frame_DeviceType)(0),    // 0: rgbx.Frame.DeviceType
	(*Frame)(nil),            // 1: rgbx.Frame
	(*SetFrameResponse)(nil), // 2: rgbx.SetFrameResponse
	(*EffectStatic)(nil),     // 3: rgbx.EffectStatic
	(*EffectMatrix)(nil),     // 4: rgbx.EffectMatrix
	(*Row)(nil),              // 5: rgbx.Row
	(*Color)(nil),            // 6: rgbx.Color
}
var file_rgbx_rgb_proto_depIdxs = []int32{
	0, // 0: rgbx.Frame.device_type:type_name -> rgbx.Frame.DeviceType
	3, // 1: rgbx.Frame.static:type_name -> rgbx.EffectStatic
	4, // 2: rgbx.Frame.matrix:type_name -> rgbx.EffectMatrix
	6, // 3: rgbx.EffectStatic.color:type_name -> rgbx.Color
	5, // 4: rgbx.EffectMatrix.rows:type_name -> rgbx.Row
	6, // 5: rgbx.Row.data:type_name -> rgbx.Color
	1, // 6: rgbx.RGBizer.SetFrame:input_type -> rgbx.Frame
	2, // 7: rgbx.RGBizer.SetFrame:output_type -> rgbx.SetFrameResponse
	7, // [7:8] is the sub-list for method output_type
	6, // [6:7] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_rgbx_rgb_proto_init() }
func file_rgbx_rgb_proto_init() {
	if File_rgbx_rgb_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rgbx_rgb_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Frame); i {
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
			switch v := v.(*SetFrameResponse); i {
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
		file_rgbx_rgb_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_rgbx_rgb_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
		file_rgbx_rgb_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
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
		(*Frame_Static)(nil),
		(*Frame_Matrix)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_rgbx_rgb_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
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