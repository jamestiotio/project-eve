// Copyright(c) 2017-2018 Zededa, Inc.
// All rights reserved.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.9
// source: config/fw.proto

package config

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

type ACEDirection int32

const (
	ACEDirection_BOTH    ACEDirection = 0
	ACEDirection_INGRESS ACEDirection = 1
	ACEDirection_EGRESS  ACEDirection = 2
)

// Enum value maps for ACEDirection.
var (
	ACEDirection_name = map[int32]string{
		0: "BOTH",
		1: "INGRESS",
		2: "EGRESS",
	}
	ACEDirection_value = map[string]int32{
		"BOTH":    0,
		"INGRESS": 1,
		"EGRESS":  2,
	}
)

func (x ACEDirection) Enum() *ACEDirection {
	p := new(ACEDirection)
	*p = x
	return p
}

func (x ACEDirection) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ACEDirection) Descriptor() protoreflect.EnumDescriptor {
	return file_config_fw_proto_enumTypes[0].Descriptor()
}

func (ACEDirection) Type() protoreflect.EnumType {
	return &file_config_fw_proto_enumTypes[0]
}

func (x ACEDirection) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ACEDirection.Descriptor instead.
func (ACEDirection) EnumDescriptor() ([]byte, []int) {
	return file_config_fw_proto_rawDescGZIP(), []int{0}
}

type ACEMatch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// FIXME: We should convert this to enum
	Type  string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *ACEMatch) Reset() {
	*x = ACEMatch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_fw_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ACEMatch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ACEMatch) ProtoMessage() {}

func (x *ACEMatch) ProtoReflect() protoreflect.Message {
	mi := &file_config_fw_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ACEMatch.ProtoReflect.Descriptor instead.
func (*ACEMatch) Descriptor() ([]byte, []int) {
	return file_config_fw_proto_rawDescGZIP(), []int{0}
}

func (x *ACEMatch) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *ACEMatch) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type ACEAction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Drop bool `protobuf:"varint,1,opt,name=drop,proto3" json:"drop,omitempty"`
	// limit action, and its associated parameter
	Limit      bool   `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Limitrate  uint32 `protobuf:"varint,3,opt,name=limitrate,proto3" json:"limitrate,omitempty"`
	Limitunit  string `protobuf:"bytes,4,opt,name=limitunit,proto3" json:"limitunit,omitempty"`
	Limitburst uint32 `protobuf:"varint,5,opt,name=limitburst,proto3" json:"limitburst,omitempty"`
	// port map action, and its associated parameter
	Portmap bool   `protobuf:"varint,6,opt,name=portmap,proto3" json:"portmap,omitempty"`
	AppPort uint32 `protobuf:"varint,7,opt,name=appPort,proto3" json:"appPort,omitempty"`
}

func (x *ACEAction) Reset() {
	*x = ACEAction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_fw_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ACEAction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ACEAction) ProtoMessage() {}

func (x *ACEAction) ProtoReflect() protoreflect.Message {
	mi := &file_config_fw_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ACEAction.ProtoReflect.Descriptor instead.
func (*ACEAction) Descriptor() ([]byte, []int) {
	return file_config_fw_proto_rawDescGZIP(), []int{1}
}

func (x *ACEAction) GetDrop() bool {
	if x != nil {
		return x.Drop
	}
	return false
}

func (x *ACEAction) GetLimit() bool {
	if x != nil {
		return x.Limit
	}
	return false
}

func (x *ACEAction) GetLimitrate() uint32 {
	if x != nil {
		return x.Limitrate
	}
	return 0
}

func (x *ACEAction) GetLimitunit() string {
	if x != nil {
		return x.Limitunit
	}
	return ""
}

func (x *ACEAction) GetLimitburst() uint32 {
	if x != nil {
		return x.Limitburst
	}
	return 0
}

func (x *ACEAction) GetPortmap() bool {
	if x != nil {
		return x.Portmap
	}
	return false
}

func (x *ACEAction) GetAppPort() uint32 {
	if x != nil {
		return x.AppPort
	}
	return 0
}

type ACE struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// multiple matches here is for various fields of 6 tuples
	//
	//	for example
	//	   1) host=www.example.com & port=http
	//	   2) ip=8.8.8.8 & port=53 & proto=UDP
	Matches []*ACEMatch `protobuf:"bytes,1,rep,name=matches,proto3" json:"matches,omitempty"`
	// Expect only single action...repeated here is
	// for future work.
	Actions []*ACEAction `protobuf:"bytes,2,rep,name=actions,proto3" json:"actions,omitempty"`
	Name    string       `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`                                        // User visible name of the ACL
	Id      int32        `protobuf:"varint,4,opt,name=id,proto3" json:"id,omitempty"`                                           // identifier
	Dir     ACEDirection `protobuf:"varint,5,opt,name=dir,proto3,enum=org.lfedge.eve.config.ACEDirection" json:"dir,omitempty"` // direction
}

func (x *ACE) Reset() {
	*x = ACE{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_fw_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ACE) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ACE) ProtoMessage() {}

func (x *ACE) ProtoReflect() protoreflect.Message {
	mi := &file_config_fw_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ACE.ProtoReflect.Descriptor instead.
func (*ACE) Descriptor() ([]byte, []int) {
	return file_config_fw_proto_rawDescGZIP(), []int{2}
}

func (x *ACE) GetMatches() []*ACEMatch {
	if x != nil {
		return x.Matches
	}
	return nil
}

func (x *ACE) GetActions() []*ACEAction {
	if x != nil {
		return x.Actions
	}
	return nil
}

func (x *ACE) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ACE) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ACE) GetDir() ACEDirection {
	if x != nil {
		return x.Dir
	}
	return ACEDirection_BOTH
}

var File_config_fw_proto protoreflect.FileDescriptor

var file_config_fw_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x66, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x15, 0x6f, 0x72, 0x67, 0x2e, 0x6c, 0x66, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x65, 0x76,
	0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x34, 0x0a, 0x08, 0x41, 0x43, 0x45, 0x4d,
	0x61, 0x74, 0x63, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xc5,
	0x01, 0x0a, 0x09, 0x41, 0x43, 0x45, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04,
	0x64, 0x72, 0x6f, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x64, 0x72, 0x6f, 0x70,
	0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x72,
	0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x72, 0x61, 0x74, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x75, 0x6e, 0x69,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x75, 0x6e,
	0x69, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x62, 0x75, 0x72, 0x73, 0x74,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x62, 0x75, 0x72,
	0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x6f, 0x72, 0x74, 0x6d, 0x61, 0x70, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x70, 0x6f, 0x72, 0x74, 0x6d, 0x61, 0x70, 0x12, 0x18, 0x0a, 0x07,
	0x61, 0x70, 0x70, 0x50, 0x6f, 0x72, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x61,
	0x70, 0x70, 0x50, 0x6f, 0x72, 0x74, 0x22, 0xd7, 0x01, 0x0a, 0x03, 0x41, 0x43, 0x45, 0x12, 0x39,
	0x0a, 0x07, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1f, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x6c, 0x66, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x65, 0x76, 0x65,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x41, 0x43, 0x45, 0x4d, 0x61, 0x74, 0x63, 0x68,
	0x52, 0x07, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x12, 0x3a, 0x0a, 0x07, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x6f, 0x72, 0x67,
	0x2e, 0x6c, 0x66, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x41, 0x43, 0x45, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x35, 0x0a, 0x03, 0x64, 0x69, 0x72,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x23, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x6c, 0x66, 0x65,
	0x64, 0x67, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x41,
	0x43, 0x45, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x03, 0x64, 0x69, 0x72,
	0x2a, 0x31, 0x0a, 0x0c, 0x41, 0x43, 0x45, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x08, 0x0a, 0x04, 0x42, 0x4f, 0x54, 0x48, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x49, 0x4e,
	0x47, 0x52, 0x45, 0x53, 0x53, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x45, 0x47, 0x52, 0x45, 0x53,
	0x53, 0x10, 0x02, 0x42, 0x3d, 0x0a, 0x15, 0x6f, 0x72, 0x67, 0x2e, 0x6c, 0x66, 0x65, 0x64, 0x67,
	0x65, 0x2e, 0x65, 0x76, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5a, 0x24, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x66, 0x2d, 0x65, 0x64, 0x67, 0x65,
	0x2f, 0x65, 0x76, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_config_fw_proto_rawDescOnce sync.Once
	file_config_fw_proto_rawDescData = file_config_fw_proto_rawDesc
)

func file_config_fw_proto_rawDescGZIP() []byte {
	file_config_fw_proto_rawDescOnce.Do(func() {
		file_config_fw_proto_rawDescData = protoimpl.X.CompressGZIP(file_config_fw_proto_rawDescData)
	})
	return file_config_fw_proto_rawDescData
}

var file_config_fw_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_config_fw_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_config_fw_proto_goTypes = []interface{}{
	(ACEDirection)(0), // 0: org.lfedge.eve.config.ACEDirection
	(*ACEMatch)(nil),  // 1: org.lfedge.eve.config.ACEMatch
	(*ACEAction)(nil), // 2: org.lfedge.eve.config.ACEAction
	(*ACE)(nil),       // 3: org.lfedge.eve.config.ACE
}
var file_config_fw_proto_depIdxs = []int32{
	1, // 0: org.lfedge.eve.config.ACE.matches:type_name -> org.lfedge.eve.config.ACEMatch
	2, // 1: org.lfedge.eve.config.ACE.actions:type_name -> org.lfedge.eve.config.ACEAction
	0, // 2: org.lfedge.eve.config.ACE.dir:type_name -> org.lfedge.eve.config.ACEDirection
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_config_fw_proto_init() }
func file_config_fw_proto_init() {
	if File_config_fw_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_config_fw_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ACEMatch); i {
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
		file_config_fw_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ACEAction); i {
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
		file_config_fw_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ACE); i {
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
			RawDescriptor: file_config_fw_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_config_fw_proto_goTypes,
		DependencyIndexes: file_config_fw_proto_depIdxs,
		EnumInfos:         file_config_fw_proto_enumTypes,
		MessageInfos:      file_config_fw_proto_msgTypes,
	}.Build()
	File_config_fw_proto = out.File
	file_config_fw_proto_rawDesc = nil
	file_config_fw_proto_goTypes = nil
	file_config_fw_proto_depIdxs = nil
}
