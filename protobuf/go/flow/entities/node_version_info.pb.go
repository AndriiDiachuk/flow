// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: flow/entities/node_version_info.proto

package entities

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

type NodeVersionInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The currently running node software version.
	Semver string `protobuf:"bytes,1,opt,name=semver,proto3" json:"semver,omitempty"`
	// The git commit hash of the currently running node software.
	Commit string `protobuf:"bytes,2,opt,name=commit,proto3" json:"commit,omitempty"`
	// The unique identifier for the node's network within the current spork.
	SporkId []byte `protobuf:"bytes,3,opt,name=spork_id,json=sporkId,proto3" json:"spork_id,omitempty"`
	// The protocol version of the currently running node software.
	ProtocolVersion uint64 `protobuf:"varint,4,opt,name=protocol_version,json=protocolVersion,proto3" json:"protocol_version,omitempty"`
	// The spork root block height. This is the height of the first sealed block in the spork network.
	SporkRootBlockHeight uint64 `protobuf:"varint,5,opt,name=spork_root_block_height,json=sporkRootBlockHeight,proto3" json:"spork_root_block_height,omitempty"`
	// The node's root block height. This is the first sealed block in the node's protocol database.
	// If the node started at the beginning of the spork, it is the same as the spork root block height.
	// If the node started after the beginning of the spork, it is the height of the first sealed block
	// indexed.
	NodeRootBlockHeight uint64 `protobuf:"varint,6,opt,name=node_root_block_height,json=nodeRootBlockHeight,proto3" json:"node_root_block_height,omitempty"`
}

func (x *NodeVersionInfo) Reset() {
	*x = NodeVersionInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flow_entities_node_version_info_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeVersionInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeVersionInfo) ProtoMessage() {}

func (x *NodeVersionInfo) ProtoReflect() protoreflect.Message {
	mi := &file_flow_entities_node_version_info_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeVersionInfo.ProtoReflect.Descriptor instead.
func (*NodeVersionInfo) Descriptor() ([]byte, []int) {
	return file_flow_entities_node_version_info_proto_rawDescGZIP(), []int{0}
}

func (x *NodeVersionInfo) GetSemver() string {
	if x != nil {
		return x.Semver
	}
	return ""
}

func (x *NodeVersionInfo) GetCommit() string {
	if x != nil {
		return x.Commit
	}
	return ""
}

func (x *NodeVersionInfo) GetSporkId() []byte {
	if x != nil {
		return x.SporkId
	}
	return nil
}

func (x *NodeVersionInfo) GetProtocolVersion() uint64 {
	if x != nil {
		return x.ProtocolVersion
	}
	return 0
}

func (x *NodeVersionInfo) GetSporkRootBlockHeight() uint64 {
	if x != nil {
		return x.SporkRootBlockHeight
	}
	return 0
}

func (x *NodeVersionInfo) GetNodeRootBlockHeight() uint64 {
	if x != nil {
		return x.NodeRootBlockHeight
	}
	return 0
}

var File_flow_entities_node_version_info_proto protoreflect.FileDescriptor

var file_flow_entities_node_version_info_proto_rawDesc = []byte{
	0x0a, 0x25, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f,
	0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x66,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x22, 0xf3, 0x01, 0x0a, 0x0f, 0x4e, 0x6f, 0x64, 0x65, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65,
	0x6d, 0x76, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x6d, 0x76,
	0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x70,
	0x6f, 0x72, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x73, 0x70,
	0x6f, 0x72, 0x6b, 0x49, 0x64, 0x12, 0x29, 0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x0f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x35, 0x0a, 0x17, 0x73, 0x70, 0x6f, 0x72, 0x6b, 0x5f, 0x72, 0x6f, 0x6f, 0x74, 0x5f, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x14, 0x73, 0x70, 0x6f, 0x72, 0x6b, 0x52, 0x6f, 0x6f, 0x74, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x33, 0x0a, 0x16, 0x6e, 0x6f, 0x64, 0x65, 0x5f,
	0x72, 0x6f, 0x6f, 0x74, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x68, 0x65, 0x69, 0x67, 0x68,
	0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x13, 0x6e, 0x6f, 0x64, 0x65, 0x52, 0x6f, 0x6f,
	0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x42, 0x50, 0x0a, 0x1c,
	0x6f, 0x72, 0x67, 0x2e, 0x6f, 0x6e, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x5a, 0x30, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x6e, 0x66, 0x6c, 0x6f, 0x77, 0x2f,
	0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x67, 0x6f,
	0x2f, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_flow_entities_node_version_info_proto_rawDescOnce sync.Once
	file_flow_entities_node_version_info_proto_rawDescData = file_flow_entities_node_version_info_proto_rawDesc
)

func file_flow_entities_node_version_info_proto_rawDescGZIP() []byte {
	file_flow_entities_node_version_info_proto_rawDescOnce.Do(func() {
		file_flow_entities_node_version_info_proto_rawDescData = protoimpl.X.CompressGZIP(file_flow_entities_node_version_info_proto_rawDescData)
	})
	return file_flow_entities_node_version_info_proto_rawDescData
}

var file_flow_entities_node_version_info_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_flow_entities_node_version_info_proto_goTypes = []interface{}{
	(*NodeVersionInfo)(nil), // 0: flow.entities.NodeVersionInfo
}
var file_flow_entities_node_version_info_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_flow_entities_node_version_info_proto_init() }
func file_flow_entities_node_version_info_proto_init() {
	if File_flow_entities_node_version_info_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_flow_entities_node_version_info_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeVersionInfo); i {
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
			RawDescriptor: file_flow_entities_node_version_info_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_flow_entities_node_version_info_proto_goTypes,
		DependencyIndexes: file_flow_entities_node_version_info_proto_depIdxs,
		MessageInfos:      file_flow_entities_node_version_info_proto_msgTypes,
	}.Build()
	File_flow_entities_node_version_info_proto = out.File
	file_flow_entities_node_version_info_proto_rawDesc = nil
	file_flow_entities_node_version_info_proto_goTypes = nil
	file_flow_entities_node_version_info_proto_depIdxs = nil
}
