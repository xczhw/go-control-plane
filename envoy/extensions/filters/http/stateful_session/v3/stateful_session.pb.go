// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v5.28.2
// source: envoy/extensions/filters/http/stateful_session/v3/stateful_session.proto

package stateful_sessionv3

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	v3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
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

type StatefulSession struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Specific implementation of session state. This session state will be used to store and
	// get address of the upstream host to which the session is assigned.
	//
	// [#extension-category: envoy.http.stateful_session]
	SessionState *v3.TypedExtensionConfig `protobuf:"bytes,1,opt,name=session_state,json=sessionState,proto3" json:"session_state,omitempty"`
	// If set to True, the HTTP request must be routed to the requested destination.
	// If the requested destination is not available, Envoy returns 503. Defaults to False.
	Strict bool `protobuf:"varint,2,opt,name=strict,proto3" json:"strict,omitempty"`
}

func (x *StatefulSession) Reset() {
	*x = StatefulSession{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_http_stateful_session_v3_stateful_session_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatefulSession) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatefulSession) ProtoMessage() {}

func (x *StatefulSession) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_http_stateful_session_v3_stateful_session_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatefulSession.ProtoReflect.Descriptor instead.
func (*StatefulSession) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_http_stateful_session_v3_stateful_session_proto_rawDescGZIP(), []int{0}
}

func (x *StatefulSession) GetSessionState() *v3.TypedExtensionConfig {
	if x != nil {
		return x.SessionState
	}
	return nil
}

func (x *StatefulSession) GetStrict() bool {
	if x != nil {
		return x.Strict
	}
	return false
}

type StatefulSessionPerRoute struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Override:
	//
	//	*StatefulSessionPerRoute_Disabled
	//	*StatefulSessionPerRoute_StatefulSession
	Override isStatefulSessionPerRoute_Override `protobuf_oneof:"override"`
}

func (x *StatefulSessionPerRoute) Reset() {
	*x = StatefulSessionPerRoute{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_http_stateful_session_v3_stateful_session_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatefulSessionPerRoute) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatefulSessionPerRoute) ProtoMessage() {}

func (x *StatefulSessionPerRoute) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_http_stateful_session_v3_stateful_session_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatefulSessionPerRoute.ProtoReflect.Descriptor instead.
func (*StatefulSessionPerRoute) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_http_stateful_session_v3_stateful_session_proto_rawDescGZIP(), []int{1}
}

func (m *StatefulSessionPerRoute) GetOverride() isStatefulSessionPerRoute_Override {
	if m != nil {
		return m.Override
	}
	return nil
}

func (x *StatefulSessionPerRoute) GetDisabled() bool {
	if x, ok := x.GetOverride().(*StatefulSessionPerRoute_Disabled); ok {
		return x.Disabled
	}
	return false
}

func (x *StatefulSessionPerRoute) GetStatefulSession() *StatefulSession {
	if x, ok := x.GetOverride().(*StatefulSessionPerRoute_StatefulSession); ok {
		return x.StatefulSession
	}
	return nil
}

type isStatefulSessionPerRoute_Override interface {
	isStatefulSessionPerRoute_Override()
}

type StatefulSessionPerRoute_Disabled struct {
	// Disable the stateful session filter for this particular vhost or route. If disabled is
	// specified in multiple per-filter-configs, the most specific one will be used.
	Disabled bool `protobuf:"varint,1,opt,name=disabled,proto3,oneof"`
}

type StatefulSessionPerRoute_StatefulSession struct {
	// Per-route stateful session configuration that can be served by RDS or static route table.
	StatefulSession *StatefulSession `protobuf:"bytes,2,opt,name=stateful_session,json=statefulSession,proto3,oneof"`
}

func (*StatefulSessionPerRoute_Disabled) isStatefulSessionPerRoute_Override() {}

func (*StatefulSessionPerRoute_StatefulSession) isStatefulSessionPerRoute_Override() {}

var File_envoy_extensions_filters_http_stateful_session_v3_stateful_session_proto protoreflect.FileDescriptor

var file_envoy_extensions_filters_http_stateful_session_v3_stateful_session_proto_rawDesc = []byte{
	0x0a, 0x48, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x66, 0x75, 0x6c, 0x5f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x2f, 0x76, 0x33, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x66, 0x75, 0x6c, 0x5f, 0x73, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x31, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x66,
	0x75, 0x6c, 0x5f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x33, 0x1a, 0x24, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6f, 0x72, 0x65,
	0x2f, 0x76, 0x33, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7a, 0x0a, 0x0f, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x66, 0x75, 0x6c, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x4f,
	0x0a, 0x0d, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x54, 0x79, 0x70,
	0x65, 0x64, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x52, 0x0c, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x06, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x22, 0xc2, 0x01, 0x0a, 0x17, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x66, 0x75, 0x6c, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x50, 0x65, 0x72, 0x52, 0x6f,
	0x75, 0x74, 0x65, 0x12, 0x25, 0x0a, 0x08, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x6a, 0x02, 0x08, 0x01, 0x48, 0x00,
	0x52, 0x08, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x6f, 0x0a, 0x10, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x66, 0x75, 0x6c, 0x5f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x42, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e,
	0x68, 0x74, 0x74, 0x70, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x66, 0x75, 0x6c, 0x5f, 0x73, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x66, 0x75,
	0x6c, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x0f, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x66, 0x75, 0x6c, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x0f, 0x0a, 0x08, 0x6f,
	0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x12, 0x03, 0xf8, 0x42, 0x01, 0x42, 0xce, 0x01, 0xba,
	0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x02, 0x0a, 0x3f, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e,
	0x68, 0x74, 0x74, 0x70, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x66, 0x75, 0x6c, 0x5f, 0x73, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x33, 0x42, 0x14, 0x53, 0x74, 0x61, 0x74, 0x65, 0x66,
	0x75, 0x6c, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x6b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65,
	0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x73, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x66, 0x75, 0x6c, 0x5f,
	0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x33, 0x3b, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x66, 0x75, 0x6c, 0x5f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x76, 0x33, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_extensions_filters_http_stateful_session_v3_stateful_session_proto_rawDescOnce sync.Once
	file_envoy_extensions_filters_http_stateful_session_v3_stateful_session_proto_rawDescData = file_envoy_extensions_filters_http_stateful_session_v3_stateful_session_proto_rawDesc
)

func file_envoy_extensions_filters_http_stateful_session_v3_stateful_session_proto_rawDescGZIP() []byte {
	file_envoy_extensions_filters_http_stateful_session_v3_stateful_session_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_filters_http_stateful_session_v3_stateful_session_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_filters_http_stateful_session_v3_stateful_session_proto_rawDescData)
	})
	return file_envoy_extensions_filters_http_stateful_session_v3_stateful_session_proto_rawDescData
}

var file_envoy_extensions_filters_http_stateful_session_v3_stateful_session_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_envoy_extensions_filters_http_stateful_session_v3_stateful_session_proto_goTypes = []interface{}{
	(*StatefulSession)(nil),         // 0: envoy.extensions.filters.http.stateful_session.v3.StatefulSession
	(*StatefulSessionPerRoute)(nil), // 1: envoy.extensions.filters.http.stateful_session.v3.StatefulSessionPerRoute
	(*v3.TypedExtensionConfig)(nil), // 2: envoy.config.core.v3.TypedExtensionConfig
}
var file_envoy_extensions_filters_http_stateful_session_v3_stateful_session_proto_depIdxs = []int32{
	2, // 0: envoy.extensions.filters.http.stateful_session.v3.StatefulSession.session_state:type_name -> envoy.config.core.v3.TypedExtensionConfig
	0, // 1: envoy.extensions.filters.http.stateful_session.v3.StatefulSessionPerRoute.stateful_session:type_name -> envoy.extensions.filters.http.stateful_session.v3.StatefulSession
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_envoy_extensions_filters_http_stateful_session_v3_stateful_session_proto_init() }
func file_envoy_extensions_filters_http_stateful_session_v3_stateful_session_proto_init() {
	if File_envoy_extensions_filters_http_stateful_session_v3_stateful_session_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_filters_http_stateful_session_v3_stateful_session_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatefulSession); i {
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
		file_envoy_extensions_filters_http_stateful_session_v3_stateful_session_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatefulSessionPerRoute); i {
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
	file_envoy_extensions_filters_http_stateful_session_v3_stateful_session_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*StatefulSessionPerRoute_Disabled)(nil),
		(*StatefulSessionPerRoute_StatefulSession)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_envoy_extensions_filters_http_stateful_session_v3_stateful_session_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_filters_http_stateful_session_v3_stateful_session_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_filters_http_stateful_session_v3_stateful_session_proto_depIdxs,
		MessageInfos:      file_envoy_extensions_filters_http_stateful_session_v3_stateful_session_proto_msgTypes,
	}.Build()
	File_envoy_extensions_filters_http_stateful_session_v3_stateful_session_proto = out.File
	file_envoy_extensions_filters_http_stateful_session_v3_stateful_session_proto_rawDesc = nil
	file_envoy_extensions_filters_http_stateful_session_v3_stateful_session_proto_goTypes = nil
	file_envoy_extensions_filters_http_stateful_session_v3_stateful_session_proto_depIdxs = nil
}