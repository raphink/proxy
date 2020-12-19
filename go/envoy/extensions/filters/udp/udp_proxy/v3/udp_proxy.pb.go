// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.10.1
// source: envoy/extensions/filters/udp/udp_proxy/v3/udp_proxy.proto

package envoy_extensions_filters_udp_udp_proxy_v3

import (
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
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

// Configuration for the UDP proxy filter.
type UdpProxyConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The stat prefix used when emitting UDP proxy filter stats.
	StatPrefix string `protobuf:"bytes,1,opt,name=stat_prefix,json=statPrefix,proto3" json:"stat_prefix,omitempty"`
	// Types that are assignable to RouteSpecifier:
	//	*UdpProxyConfig_Cluster
	RouteSpecifier isUdpProxyConfig_RouteSpecifier `protobuf_oneof:"route_specifier"`
	// The idle timeout for sessions. Idle is defined as no datagrams between received or sent by
	// the session. The default if not specified is 1 minute.
	IdleTimeout *duration.Duration `protobuf:"bytes,3,opt,name=idle_timeout,json=idleTimeout,proto3" json:"idle_timeout,omitempty"`
}

func (x *UdpProxyConfig) Reset() {
	*x = UdpProxyConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_udp_udp_proxy_v3_udp_proxy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UdpProxyConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UdpProxyConfig) ProtoMessage() {}

func (x *UdpProxyConfig) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_udp_udp_proxy_v3_udp_proxy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UdpProxyConfig.ProtoReflect.Descriptor instead.
func (*UdpProxyConfig) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_udp_udp_proxy_v3_udp_proxy_proto_rawDescGZIP(), []int{0}
}

func (x *UdpProxyConfig) GetStatPrefix() string {
	if x != nil {
		return x.StatPrefix
	}
	return ""
}

func (m *UdpProxyConfig) GetRouteSpecifier() isUdpProxyConfig_RouteSpecifier {
	if m != nil {
		return m.RouteSpecifier
	}
	return nil
}

func (x *UdpProxyConfig) GetCluster() string {
	if x, ok := x.GetRouteSpecifier().(*UdpProxyConfig_Cluster); ok {
		return x.Cluster
	}
	return ""
}

func (x *UdpProxyConfig) GetIdleTimeout() *duration.Duration {
	if x != nil {
		return x.IdleTimeout
	}
	return nil
}

type isUdpProxyConfig_RouteSpecifier interface {
	isUdpProxyConfig_RouteSpecifier()
}

type UdpProxyConfig_Cluster struct {
	// The upstream cluster to connect to.
	Cluster string `protobuf:"bytes,2,opt,name=cluster,proto3,oneof"`
}

func (*UdpProxyConfig_Cluster) isUdpProxyConfig_RouteSpecifier() {}

var File_envoy_extensions_filters_udp_udp_proxy_v3_udp_proxy_proto protoreflect.FileDescriptor

var file_envoy_extensions_filters_udp_udp_proxy_v3_udp_proxy_proto_rawDesc = []byte{
	0x0a, 0x39, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x75, 0x64, 0x70, 0x2f, 0x75,
	0x64, 0x70, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x76, 0x33, 0x2f, 0x75, 0x64, 0x70, 0x5f,
	0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x29, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x75, 0x64, 0x70, 0x2e, 0x75, 0x64, 0x70, 0x5f, 0x70, 0x72,
	0x6f, 0x78, 0x79, 0x2e, 0x76, 0x33, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x69,
	0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xf6, 0x01, 0x0a, 0x0e, 0x55, 0x64, 0x70, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x28, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x5f, 0x70, 0x72, 0x65,
	0x66, 0x69, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02,
	0x20, 0x01, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x12, 0x23,
	0x0a, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x20, 0x01, 0x48, 0x00, 0x52, 0x07, 0x63, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x12, 0x3c, 0x0a, 0x0c, 0x69, 0x64, 0x6c, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x6f, 0x75, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x69, 0x64, 0x6c, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75,
	0x74, 0x3a, 0x3f, 0x9a, 0xc5, 0x88, 0x1e, 0x3a, 0x0a, 0x38, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x75, 0x64,
	0x70, 0x2e, 0x75, 0x64, 0x70, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76, 0x32, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x2e, 0x55, 0x64, 0x70, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x42, 0x16, 0x0a, 0x0f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f, 0x73, 0x70, 0x65, 0x63,
	0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x03, 0xf8, 0x42, 0x01, 0x42, 0x52, 0x0a, 0x37, 0x69, 0x6f,
	0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x73, 0x2e, 0x75, 0x64, 0x70, 0x2e, 0x75, 0x64, 0x70, 0x5f, 0x70, 0x72, 0x6f,
	0x78, 0x79, 0x2e, 0x76, 0x33, 0x42, 0x0d, 0x55, 0x64, 0x70, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x02, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_extensions_filters_udp_udp_proxy_v3_udp_proxy_proto_rawDescOnce sync.Once
	file_envoy_extensions_filters_udp_udp_proxy_v3_udp_proxy_proto_rawDescData = file_envoy_extensions_filters_udp_udp_proxy_v3_udp_proxy_proto_rawDesc
)

func file_envoy_extensions_filters_udp_udp_proxy_v3_udp_proxy_proto_rawDescGZIP() []byte {
	file_envoy_extensions_filters_udp_udp_proxy_v3_udp_proxy_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_filters_udp_udp_proxy_v3_udp_proxy_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_filters_udp_udp_proxy_v3_udp_proxy_proto_rawDescData)
	})
	return file_envoy_extensions_filters_udp_udp_proxy_v3_udp_proxy_proto_rawDescData
}

var file_envoy_extensions_filters_udp_udp_proxy_v3_udp_proxy_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_envoy_extensions_filters_udp_udp_proxy_v3_udp_proxy_proto_goTypes = []interface{}{
	(*UdpProxyConfig)(nil),    // 0: envoy.extensions.filters.udp.udp_proxy.v3.UdpProxyConfig
	(*duration.Duration)(nil), // 1: google.protobuf.Duration
}
var file_envoy_extensions_filters_udp_udp_proxy_v3_udp_proxy_proto_depIdxs = []int32{
	1, // 0: envoy.extensions.filters.udp.udp_proxy.v3.UdpProxyConfig.idle_timeout:type_name -> google.protobuf.Duration
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_envoy_extensions_filters_udp_udp_proxy_v3_udp_proxy_proto_init() }
func file_envoy_extensions_filters_udp_udp_proxy_v3_udp_proxy_proto_init() {
	if File_envoy_extensions_filters_udp_udp_proxy_v3_udp_proxy_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_filters_udp_udp_proxy_v3_udp_proxy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UdpProxyConfig); i {
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
	file_envoy_extensions_filters_udp_udp_proxy_v3_udp_proxy_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*UdpProxyConfig_Cluster)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_envoy_extensions_filters_udp_udp_proxy_v3_udp_proxy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_filters_udp_udp_proxy_v3_udp_proxy_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_filters_udp_udp_proxy_v3_udp_proxy_proto_depIdxs,
		MessageInfos:      file_envoy_extensions_filters_udp_udp_proxy_v3_udp_proxy_proto_msgTypes,
	}.Build()
	File_envoy_extensions_filters_udp_udp_proxy_v3_udp_proxy_proto = out.File
	file_envoy_extensions_filters_udp_udp_proxy_v3_udp_proxy_proto_rawDesc = nil
	file_envoy_extensions_filters_udp_udp_proxy_v3_udp_proxy_proto_goTypes = nil
	file_envoy_extensions_filters_udp_udp_proxy_v3_udp_proxy_proto_depIdxs = nil
}