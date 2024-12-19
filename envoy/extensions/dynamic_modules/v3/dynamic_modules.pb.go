// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v5.29.1
// source: envoy/extensions/dynamic_modules/v3/dynamic_modules.proto

package dynamic_modulesv3

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	_ "github.com/cncf/xds/go/xds/annotations/v3"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

// Configuration of a dynamic module. A dynamic module is a shared object file that can be loaded via dlopen
// by various Envoy extension points. Currently, only HTTP filter (envoy.filters.http.dynamic_modules) is supported.
//
// How a module is loaded is determined by the extension point that uses it. For example, the HTTP filter
// loads the module with dlopen when Envoy receives a configuration that references the module at load time.
// If loading the module fails, the configuration will be rejected.
//
// Whether or not the shared object is the same is determined by the file path as well as the file's inode depending
// on the platform. Notably, if the file path and the content of the file are the same, the shared object will be reused.
//
// A module must be compatible with the ABI specified in :repo:`abi.h <source/extensions/dynamic_modules/abi.h>`.
// Currently, compatibility is only guaranteed by an exact version match between the Envoy
// codebase and the dynamic module SDKs. In the future, after the ABI is stabilized, we will revisit
// this restriction and hopefully provide a wider compatibility guarantee. Until then, Envoy
// checks the hash of the ABI header files to ensure that the dynamic modules are built against the
// same version of the ABI.
//
// Currently, the implementation is work in progress and not usable.
type DynamicModuleConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the dynamic module. The client is expected to have some configuration indicating where to search for the module.
	// In Envoy, the search path can only be configured via the environment variable “ENVOY_DYNAMIC_MODULES_SEARCH_PATH“.
	// The actual search path is “${ENVOY_DYNAMIC_MODULES_SEARCH_PATH}/lib${name}.so“. TODO: make the search path configurable via
	// command line options.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Set true to prevent the module from being unloaded with dlclose.
	// This is useful for modules that have global state that should not be unloaded.
	// A module is closed when no more references to it exist in the process. For example,
	// no HTTP filters are using the module (e.g. after configuration update).
	DoNotClose bool `protobuf:"varint,3,opt,name=do_not_close,json=doNotClose,proto3" json:"do_not_close,omitempty"`
}

func (x *DynamicModuleConfig) Reset() {
	*x = DynamicModuleConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_dynamic_modules_v3_dynamic_modules_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DynamicModuleConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DynamicModuleConfig) ProtoMessage() {}

func (x *DynamicModuleConfig) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_dynamic_modules_v3_dynamic_modules_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DynamicModuleConfig.ProtoReflect.Descriptor instead.
func (*DynamicModuleConfig) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_dynamic_modules_v3_dynamic_modules_proto_rawDescGZIP(), []int{0}
}

func (x *DynamicModuleConfig) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DynamicModuleConfig) GetDoNotClose() bool {
	if x != nil {
		return x.DoNotClose
	}
	return false
}

var File_envoy_extensions_dynamic_modules_v3_dynamic_modules_proto protoreflect.FileDescriptor

var file_envoy_extensions_dynamic_modules_v3_dynamic_modules_proto_rawDesc = []byte{
	0x0a, 0x39, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x5f, 0x6d, 0x6f, 0x64, 0x75, 0x6c,
	0x65, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x5f, 0x6d, 0x6f,
	0x64, 0x75, 0x6c, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x23, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x64, 0x79,
	0x6e, 0x61, 0x6d, 0x69, 0x63, 0x5f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x2e, 0x76, 0x33,
	0x1a, 0x1f, 0x78, 0x64, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x76, 0x33, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x54, 0x0a, 0x13, 0x44, 0x79, 0x6e,
	0x61, 0x6d, 0x69, 0x63, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x12, 0x1b, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07,
	0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a,
	0x0c, 0x64, 0x6f, 0x5f, 0x6e, 0x6f, 0x74, 0x5f, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0a, 0x64, 0x6f, 0x4e, 0x6f, 0x74, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x42,
	0xb8, 0x01, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x02, 0xd2, 0xc6, 0xa4, 0xe1, 0x06, 0x02,
	0x08, 0x01, 0x0a, 0x31, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78,
	0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x5f, 0x6d, 0x6f, 0x64, 0x75, 0x6c,
	0x65, 0x73, 0x2e, 0x76, 0x33, 0x42, 0x13, 0x44, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x4d, 0x6f,
	0x64, 0x75, 0x6c, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x5c, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72,
	0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x70,
	0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x5f, 0x6d, 0x6f,
	0x64, 0x75, 0x6c, 0x65, 0x73, 0x2f, 0x76, 0x33, 0x3b, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63,
	0x5f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x76, 0x33, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_envoy_extensions_dynamic_modules_v3_dynamic_modules_proto_rawDescOnce sync.Once
	file_envoy_extensions_dynamic_modules_v3_dynamic_modules_proto_rawDescData = file_envoy_extensions_dynamic_modules_v3_dynamic_modules_proto_rawDesc
)

func file_envoy_extensions_dynamic_modules_v3_dynamic_modules_proto_rawDescGZIP() []byte {
	file_envoy_extensions_dynamic_modules_v3_dynamic_modules_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_dynamic_modules_v3_dynamic_modules_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_dynamic_modules_v3_dynamic_modules_proto_rawDescData)
	})
	return file_envoy_extensions_dynamic_modules_v3_dynamic_modules_proto_rawDescData
}

var file_envoy_extensions_dynamic_modules_v3_dynamic_modules_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_envoy_extensions_dynamic_modules_v3_dynamic_modules_proto_goTypes = []interface{}{
	(*DynamicModuleConfig)(nil), // 0: envoy.extensions.dynamic_modules.v3.DynamicModuleConfig
}
var file_envoy_extensions_dynamic_modules_v3_dynamic_modules_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_envoy_extensions_dynamic_modules_v3_dynamic_modules_proto_init() }
func file_envoy_extensions_dynamic_modules_v3_dynamic_modules_proto_init() {
	if File_envoy_extensions_dynamic_modules_v3_dynamic_modules_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_dynamic_modules_v3_dynamic_modules_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DynamicModuleConfig); i {
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
			RawDescriptor: file_envoy_extensions_dynamic_modules_v3_dynamic_modules_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_dynamic_modules_v3_dynamic_modules_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_dynamic_modules_v3_dynamic_modules_proto_depIdxs,
		MessageInfos:      file_envoy_extensions_dynamic_modules_v3_dynamic_modules_proto_msgTypes,
	}.Build()
	File_envoy_extensions_dynamic_modules_v3_dynamic_modules_proto = out.File
	file_envoy_extensions_dynamic_modules_v3_dynamic_modules_proto_rawDesc = nil
	file_envoy_extensions_dynamic_modules_v3_dynamic_modules_proto_goTypes = nil
	file_envoy_extensions_dynamic_modules_v3_dynamic_modules_proto_depIdxs = nil
}