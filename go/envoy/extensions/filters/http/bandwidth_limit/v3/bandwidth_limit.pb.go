// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.19.1
// source: envoy/extensions/filters/http/bandwidth_limit/v3/bandwidth_limit.proto

package bandwidth_limitv3

import (
	v3 "github.com/cilium/proxy/go/envoy/config/core/v3"
	_ "github.com/cncf/xds/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
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

// Defines the mode for the bandwidth limit filter.
// Values represent bitmask.
type BandwidthLimit_EnableMode int32

const (
	// Filter is disabled.
	BandwidthLimit_DISABLED BandwidthLimit_EnableMode = 0
	// Filter enabled only for incoming traffic.
	BandwidthLimit_REQUEST BandwidthLimit_EnableMode = 1
	// Filter enabled only for outgoing traffic.
	BandwidthLimit_RESPONSE BandwidthLimit_EnableMode = 2
	// Filter enabled for both incoming and outgoing traffic.
	BandwidthLimit_REQUEST_AND_RESPONSE BandwidthLimit_EnableMode = 3
)

// Enum value maps for BandwidthLimit_EnableMode.
var (
	BandwidthLimit_EnableMode_name = map[int32]string{
		0: "DISABLED",
		1: "REQUEST",
		2: "RESPONSE",
		3: "REQUEST_AND_RESPONSE",
	}
	BandwidthLimit_EnableMode_value = map[string]int32{
		"DISABLED":             0,
		"REQUEST":              1,
		"RESPONSE":             2,
		"REQUEST_AND_RESPONSE": 3,
	}
)

func (x BandwidthLimit_EnableMode) Enum() *BandwidthLimit_EnableMode {
	p := new(BandwidthLimit_EnableMode)
	*p = x
	return p
}

func (x BandwidthLimit_EnableMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BandwidthLimit_EnableMode) Descriptor() protoreflect.EnumDescriptor {
	return file_envoy_extensions_filters_http_bandwidth_limit_v3_bandwidth_limit_proto_enumTypes[0].Descriptor()
}

func (BandwidthLimit_EnableMode) Type() protoreflect.EnumType {
	return &file_envoy_extensions_filters_http_bandwidth_limit_v3_bandwidth_limit_proto_enumTypes[0]
}

func (x BandwidthLimit_EnableMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BandwidthLimit_EnableMode.Descriptor instead.
func (BandwidthLimit_EnableMode) EnumDescriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_http_bandwidth_limit_v3_bandwidth_limit_proto_rawDescGZIP(), []int{0, 0}
}

// [#next-free-field: 8]
type BandwidthLimit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The human readable prefix to use when emitting stats.
	StatPrefix string `protobuf:"bytes,1,opt,name=stat_prefix,json=statPrefix,proto3" json:"stat_prefix,omitempty"`
	// The enable mode for the bandwidth limit filter.
	// Default is Disabled.
	EnableMode BandwidthLimit_EnableMode `protobuf:"varint,2,opt,name=enable_mode,json=enableMode,proto3,enum=envoy.extensions.filters.http.bandwidth_limit.v3.BandwidthLimit_EnableMode" json:"enable_mode,omitempty"`
	// The limit supplied in KiB/s.
	//
	// .. note::
	//   It's fine for the limit to be unset for the global configuration since the bandwidth limit
	//   can be applied at a the virtual host or route level. Thus, the limit must be set for the
	//   per route configuration otherwise the config will be rejected.
	//
	// .. note::
	//   When using per route configuration, the limit becomes unique to that route.
	//
	LimitKbps *wrapperspb.UInt64Value `protobuf:"bytes,3,opt,name=limit_kbps,json=limitKbps,proto3" json:"limit_kbps,omitempty"`
	// Optional Fill interval in milliseconds for the token refills. Defaults to 50ms.
	// It must be at least 20ms to avoid too aggressive refills.
	FillInterval *durationpb.Duration `protobuf:"bytes,4,opt,name=fill_interval,json=fillInterval,proto3" json:"fill_interval,omitempty"`
	// Runtime flag that controls whether the filter is enabled or not. If not specified, defaults
	// to enabled.
	RuntimeEnabled *v3.RuntimeFeatureFlag `protobuf:"bytes,5,opt,name=runtime_enabled,json=runtimeEnabled,proto3" json:"runtime_enabled,omitempty"`
	// Enable response trailers.
	//
	// .. note::
	//
	//   * If set true, the response trailers *bandwidth-request-delay-ms* and *bandwidth-response-delay-ms* will be added, prefixed by *response_trailer_prefix*.
	//   * bandwidth-request-delay-ms: delay time in milliseconds it took for the request stream transfer.
	//   * bandwidth-response-delay-ms: delay time in milliseconds it took for the response stream transfer.
	//   * If :ref:`enable_mode <envoy_v3_api_field_extensions.filters.http.bandwidth_limit.v3.BandwidthLimit.enable_mode>` is DISABLED or REQUEST, the trailers will not be set.
	//   * If both the request and response delay time is 0, the trailers will not be set.
	//
	EnableResponseTrailers bool `protobuf:"varint,6,opt,name=enable_response_trailers,json=enableResponseTrailers,proto3" json:"enable_response_trailers,omitempty"`
	// Optional The prefix for the response trailers.
	ResponseTrailerPrefix string `protobuf:"bytes,7,opt,name=response_trailer_prefix,json=responseTrailerPrefix,proto3" json:"response_trailer_prefix,omitempty"`
}

func (x *BandwidthLimit) Reset() {
	*x = BandwidthLimit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_http_bandwidth_limit_v3_bandwidth_limit_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BandwidthLimit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BandwidthLimit) ProtoMessage() {}

func (x *BandwidthLimit) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_http_bandwidth_limit_v3_bandwidth_limit_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BandwidthLimit.ProtoReflect.Descriptor instead.
func (*BandwidthLimit) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_http_bandwidth_limit_v3_bandwidth_limit_proto_rawDescGZIP(), []int{0}
}

func (x *BandwidthLimit) GetStatPrefix() string {
	if x != nil {
		return x.StatPrefix
	}
	return ""
}

func (x *BandwidthLimit) GetEnableMode() BandwidthLimit_EnableMode {
	if x != nil {
		return x.EnableMode
	}
	return BandwidthLimit_DISABLED
}

func (x *BandwidthLimit) GetLimitKbps() *wrapperspb.UInt64Value {
	if x != nil {
		return x.LimitKbps
	}
	return nil
}

func (x *BandwidthLimit) GetFillInterval() *durationpb.Duration {
	if x != nil {
		return x.FillInterval
	}
	return nil
}

func (x *BandwidthLimit) GetRuntimeEnabled() *v3.RuntimeFeatureFlag {
	if x != nil {
		return x.RuntimeEnabled
	}
	return nil
}

func (x *BandwidthLimit) GetEnableResponseTrailers() bool {
	if x != nil {
		return x.EnableResponseTrailers
	}
	return false
}

func (x *BandwidthLimit) GetResponseTrailerPrefix() string {
	if x != nil {
		return x.ResponseTrailerPrefix
	}
	return ""
}

var File_envoy_extensions_filters_http_bandwidth_limit_v3_bandwidth_limit_proto protoreflect.FileDescriptor

var file_envoy_extensions_filters_http_bandwidth_limit_v3_bandwidth_limit_proto_rawDesc = []byte{
	0x0a, 0x46, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f,
	0x62, 0x61, 0x6e, 0x64, 0x77, 0x69, 0x64, 0x74, 0x68, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2f,
	0x76, 0x33, 0x2f, 0x62, 0x61, 0x6e, 0x64, 0x77, 0x69, 0x64, 0x74, 0x68, 0x5f, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x30, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e,
	0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x62, 0x61, 0x6e, 0x64, 0x77, 0x69, 0x64, 0x74,
	0x68, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x76, 0x33, 0x1a, 0x1f, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x33,
	0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61,
	0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70,
	0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xe1, 0x04, 0x0a, 0x0e, 0x42, 0x61, 0x6e, 0x64, 0x77, 0x69, 0x64, 0x74,
	0x68, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x28, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x5f, 0x70,
	0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04,
	0x72, 0x02, 0x10, 0x01, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78,
	0x12, 0x76, 0x0a, 0x0b, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x4b, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78,
	0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73,
	0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x62, 0x61, 0x6e, 0x64, 0x77, 0x69, 0x64, 0x74, 0x68, 0x5f,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x76, 0x33, 0x2e, 0x42, 0x61, 0x6e, 0x64, 0x77, 0x69, 0x64,
	0x74, 0x68, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x4d, 0x6f,
	0x64, 0x65, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x82, 0x01, 0x02, 0x10, 0x01, 0x52, 0x0a, 0x65, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x44, 0x0a, 0x0a, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x5f, 0x6b, 0x62, 0x70, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55,
	0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x32,
	0x02, 0x28, 0x01, 0x52, 0x09, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x4b, 0x62, 0x70, 0x73, 0x12, 0x51,
	0x0a, 0x0d, 0x66, 0x69, 0x6c, 0x6c, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x42, 0x11, 0xfa, 0x42, 0x0e, 0xaa, 0x01, 0x0b, 0x22, 0x02, 0x08, 0x01, 0x32, 0x05, 0x10, 0x80,
	0xda, 0xc4, 0x09, 0x52, 0x0c, 0x66, 0x69, 0x6c, 0x6c, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61,
	0x6c, 0x12, 0x51, 0x0a, 0x0f, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x65, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76,
	0x33, 0x2e, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x46, 0x6c, 0x61, 0x67, 0x52, 0x0e, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x45, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x64, 0x12, 0x38, 0x0a, 0x18, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x74, 0x72, 0x61, 0x69, 0x6c, 0x65, 0x72, 0x73,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x16, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x54, 0x72, 0x61, 0x69, 0x6c, 0x65, 0x72, 0x73, 0x12, 0x36,
	0x0a, 0x17, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x74, 0x72, 0x61, 0x69, 0x6c,
	0x65, 0x72, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x15, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x54, 0x72, 0x61, 0x69, 0x6c, 0x65, 0x72,
	0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x22, 0x4f, 0x0a, 0x0a, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x4d, 0x6f, 0x64, 0x65, 0x12, 0x0c, 0x0a, 0x08, 0x44, 0x49, 0x53, 0x41, 0x42, 0x4c, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x10, 0x01, 0x12,
	0x0c, 0x0a, 0x08, 0x52, 0x45, 0x53, 0x50, 0x4f, 0x4e, 0x53, 0x45, 0x10, 0x02, 0x12, 0x18, 0x0a,
	0x14, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x5f, 0x41, 0x4e, 0x44, 0x5f, 0x52, 0x45, 0x53,
	0x50, 0x4f, 0x4e, 0x53, 0x45, 0x10, 0x03, 0x42, 0xca, 0x01, 0x0a, 0x3e, 0x69, 0x6f, 0x2e, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e,
	0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x62, 0x61, 0x6e, 0x64, 0x77, 0x69, 0x64, 0x74,
	0x68, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x76, 0x33, 0x42, 0x13, 0x42, 0x61, 0x6e, 0x64,
	0x77, 0x69, 0x64, 0x74, 0x68, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x69, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f,
	0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x73, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f, 0x62, 0x61, 0x6e, 0x64, 0x77, 0x69, 0x64, 0x74,
	0x68, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2f, 0x76, 0x33, 0x3b, 0x62, 0x61, 0x6e, 0x64, 0x77,
	0x69, 0x64, 0x74, 0x68, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x76, 0x33, 0xba, 0x80, 0xc8, 0xd1,
	0x06, 0x02, 0x10, 0x02, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_extensions_filters_http_bandwidth_limit_v3_bandwidth_limit_proto_rawDescOnce sync.Once
	file_envoy_extensions_filters_http_bandwidth_limit_v3_bandwidth_limit_proto_rawDescData = file_envoy_extensions_filters_http_bandwidth_limit_v3_bandwidth_limit_proto_rawDesc
)

func file_envoy_extensions_filters_http_bandwidth_limit_v3_bandwidth_limit_proto_rawDescGZIP() []byte {
	file_envoy_extensions_filters_http_bandwidth_limit_v3_bandwidth_limit_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_filters_http_bandwidth_limit_v3_bandwidth_limit_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_filters_http_bandwidth_limit_v3_bandwidth_limit_proto_rawDescData)
	})
	return file_envoy_extensions_filters_http_bandwidth_limit_v3_bandwidth_limit_proto_rawDescData
}

var file_envoy_extensions_filters_http_bandwidth_limit_v3_bandwidth_limit_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_envoy_extensions_filters_http_bandwidth_limit_v3_bandwidth_limit_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_envoy_extensions_filters_http_bandwidth_limit_v3_bandwidth_limit_proto_goTypes = []interface{}{
	(BandwidthLimit_EnableMode)(0), // 0: envoy.extensions.filters.http.bandwidth_limit.v3.BandwidthLimit.EnableMode
	(*BandwidthLimit)(nil),         // 1: envoy.extensions.filters.http.bandwidth_limit.v3.BandwidthLimit
	(*wrapperspb.UInt64Value)(nil), // 2: google.protobuf.UInt64Value
	(*durationpb.Duration)(nil),    // 3: google.protobuf.Duration
	(*v3.RuntimeFeatureFlag)(nil),  // 4: envoy.config.core.v3.RuntimeFeatureFlag
}
var file_envoy_extensions_filters_http_bandwidth_limit_v3_bandwidth_limit_proto_depIdxs = []int32{
	0, // 0: envoy.extensions.filters.http.bandwidth_limit.v3.BandwidthLimit.enable_mode:type_name -> envoy.extensions.filters.http.bandwidth_limit.v3.BandwidthLimit.EnableMode
	2, // 1: envoy.extensions.filters.http.bandwidth_limit.v3.BandwidthLimit.limit_kbps:type_name -> google.protobuf.UInt64Value
	3, // 2: envoy.extensions.filters.http.bandwidth_limit.v3.BandwidthLimit.fill_interval:type_name -> google.protobuf.Duration
	4, // 3: envoy.extensions.filters.http.bandwidth_limit.v3.BandwidthLimit.runtime_enabled:type_name -> envoy.config.core.v3.RuntimeFeatureFlag
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_envoy_extensions_filters_http_bandwidth_limit_v3_bandwidth_limit_proto_init() }
func file_envoy_extensions_filters_http_bandwidth_limit_v3_bandwidth_limit_proto_init() {
	if File_envoy_extensions_filters_http_bandwidth_limit_v3_bandwidth_limit_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_filters_http_bandwidth_limit_v3_bandwidth_limit_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BandwidthLimit); i {
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
			RawDescriptor: file_envoy_extensions_filters_http_bandwidth_limit_v3_bandwidth_limit_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_filters_http_bandwidth_limit_v3_bandwidth_limit_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_filters_http_bandwidth_limit_v3_bandwidth_limit_proto_depIdxs,
		EnumInfos:         file_envoy_extensions_filters_http_bandwidth_limit_v3_bandwidth_limit_proto_enumTypes,
		MessageInfos:      file_envoy_extensions_filters_http_bandwidth_limit_v3_bandwidth_limit_proto_msgTypes,
	}.Build()
	File_envoy_extensions_filters_http_bandwidth_limit_v3_bandwidth_limit_proto = out.File
	file_envoy_extensions_filters_http_bandwidth_limit_v3_bandwidth_limit_proto_rawDesc = nil
	file_envoy_extensions_filters_http_bandwidth_limit_v3_bandwidth_limit_proto_goTypes = nil
	file_envoy_extensions_filters_http_bandwidth_limit_v3_bandwidth_limit_proto_depIdxs = nil
}