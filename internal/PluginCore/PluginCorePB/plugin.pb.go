// cd .\internal\PluginCore\protobuf\
// protoc -I . --go_out=plugins=grpc:. .\plugin.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.2
// source: plugin.proto

package PluginCorePB

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// 插件路由注册请求
type PluginRouteRegisteredRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RouterName        string                      `protobuf:"bytes,1,opt,name=RouterName,proto3" json:"RouterName,omitempty"`
	RouterMethod      string                      `protobuf:"bytes,2,opt,name=RouterMethod,proto3" json:"RouterMethod,omitempty"`
	RouterAuthIF      bool                        `protobuf:"varint,3,opt,name=RouterAuthIF,proto3" json:"RouterAuthIF,omitempty"`
	PluginRouterGroup []*PluginRouterGroupRequest `protobuf:"bytes,4,rep,name=PluginRouterGroup,proto3" json:"PluginRouterGroup,omitempty"`
	PluginInfo        *PluginInfoRequest          `protobuf:"bytes,5,opt,name=PluginInfo,proto3" json:"PluginInfo,omitempty"`
}

func (x *PluginRouteRegisteredRequest) Reset() {
	*x = PluginRouteRegisteredRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PluginRouteRegisteredRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PluginRouteRegisteredRequest) ProtoMessage() {}

func (x *PluginRouteRegisteredRequest) ProtoReflect() protoreflect.Message {
	mi := &file_plugin_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PluginRouteRegisteredRequest.ProtoReflect.Descriptor instead.
func (*PluginRouteRegisteredRequest) Descriptor() ([]byte, []int) {
	return file_plugin_proto_rawDescGZIP(), []int{0}
}

func (x *PluginRouteRegisteredRequest) GetRouterName() string {
	if x != nil {
		return x.RouterName
	}
	return ""
}

func (x *PluginRouteRegisteredRequest) GetRouterMethod() string {
	if x != nil {
		return x.RouterMethod
	}
	return ""
}

func (x *PluginRouteRegisteredRequest) GetRouterAuthIF() bool {
	if x != nil {
		return x.RouterAuthIF
	}
	return false
}

func (x *PluginRouteRegisteredRequest) GetPluginRouterGroup() []*PluginRouterGroupRequest {
	if x != nil {
		return x.PluginRouterGroup
	}
	return nil
}

func (x *PluginRouteRegisteredRequest) GetPluginInfo() *PluginInfoRequest {
	if x != nil {
		return x.PluginInfo
	}
	return nil
}

// 插件路由注册响应
type PluginRouteRegisteredResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code       string            `protobuf:"bytes,1,opt,name=Code,proto3" json:"Code,omitempty"`
	Data       map[string]string `protobuf:"bytes,2,rep,name=Data,proto3" json:"Data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	RouterPath string            `protobuf:"bytes,3,opt,name=RouterPath,proto3" json:"RouterPath,omitempty"`
	Message    string            `protobuf:"bytes,4,opt,name=Message,proto3" json:"Message,omitempty"`
}

func (x *PluginRouteRegisteredResponse) Reset() {
	*x = PluginRouteRegisteredResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugin_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PluginRouteRegisteredResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PluginRouteRegisteredResponse) ProtoMessage() {}

func (x *PluginRouteRegisteredResponse) ProtoReflect() protoreflect.Message {
	mi := &file_plugin_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PluginRouteRegisteredResponse.ProtoReflect.Descriptor instead.
func (*PluginRouteRegisteredResponse) Descriptor() ([]byte, []int) {
	return file_plugin_proto_rawDescGZIP(), []int{1}
}

func (x *PluginRouteRegisteredResponse) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *PluginRouteRegisteredResponse) GetData() map[string]string {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *PluginRouteRegisteredResponse) GetRouterPath() string {
	if x != nil {
		return x.RouterPath
	}
	return ""
}

func (x *PluginRouteRegisteredResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// 插件路由组请求
type PluginRouterGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RouterPath   string `protobuf:"bytes,1,opt,name=routerPath,proto3" json:"routerPath,omitempty"`
	RouterMethod string `protobuf:"bytes,2,opt,name=RouterMethod,proto3" json:"RouterMethod,omitempty"`
	RouterAuthIF bool   `protobuf:"varint,3,opt,name=RouterAuthIF,proto3" json:"RouterAuthIF,omitempty"`
}

func (x *PluginRouterGroupRequest) Reset() {
	*x = PluginRouterGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugin_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PluginRouterGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PluginRouterGroupRequest) ProtoMessage() {}

func (x *PluginRouterGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_plugin_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PluginRouterGroupRequest.ProtoReflect.Descriptor instead.
func (*PluginRouterGroupRequest) Descriptor() ([]byte, []int) {
	return file_plugin_proto_rawDescGZIP(), []int{2}
}

func (x *PluginRouterGroupRequest) GetRouterPath() string {
	if x != nil {
		return x.RouterPath
	}
	return ""
}

func (x *PluginRouterGroupRequest) GetRouterMethod() string {
	if x != nil {
		return x.RouterMethod
	}
	return ""
}

func (x *PluginRouterGroupRequest) GetRouterAuthIF() bool {
	if x != nil {
		return x.RouterAuthIF
	}
	return false
}

// 插件信息请求
type PluginInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PluginName    string `protobuf:"bytes,1,opt,name=PluginName,proto3" json:"PluginName,omitempty"`
	PluginPort    string `protobuf:"bytes,2,opt,name=PluginPort,proto3" json:"PluginPort,omitempty"`
	PluginVersion string `protobuf:"bytes,3,opt,name=PluginVersion,proto3" json:"PluginVersion,omitempty"`
}

func (x *PluginInfoRequest) Reset() {
	*x = PluginInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugin_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PluginInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PluginInfoRequest) ProtoMessage() {}

func (x *PluginInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_plugin_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PluginInfoRequest.ProtoReflect.Descriptor instead.
func (*PluginInfoRequest) Descriptor() ([]byte, []int) {
	return file_plugin_proto_rawDescGZIP(), []int{3}
}

func (x *PluginInfoRequest) GetPluginName() string {
	if x != nil {
		return x.PluginName
	}
	return ""
}

func (x *PluginInfoRequest) GetPluginPort() string {
	if x != nil {
		return x.PluginPort
	}
	return ""
}

func (x *PluginInfoRequest) GetPluginVersion() string {
	if x != nil {
		return x.PluginVersion
	}
	return ""
}

var File_plugin_proto protoreflect.FileDescriptor

var file_plugin_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c,
	0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x43, 0x6f, 0x72, 0x65, 0x50, 0x42, 0x22, 0x9d, 0x02, 0x0a,
	0x1c, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a,
	0x0a, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a,
	0x0c, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x12, 0x22, 0x0a, 0x0c, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x41, 0x75, 0x74, 0x68, 0x49,
	0x46, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x41,
	0x75, 0x74, 0x68, 0x49, 0x46, 0x12, 0x54, 0x0a, 0x11, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x52,
	0x6f, 0x75, 0x74, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x26, 0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x43, 0x6f, 0x72, 0x65, 0x50, 0x42, 0x2e,
	0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x11, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x52, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x3f, 0x0a, 0x0a, 0x50,
	0x6c, 0x75, 0x67, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1f, 0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x43, 0x6f, 0x72, 0x65, 0x50, 0x42, 0x2e, 0x50,
	0x6c, 0x75, 0x67, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x52, 0x0a, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0xf1, 0x01, 0x0a,
	0x1d, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x49, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x35, 0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x43, 0x6f, 0x72, 0x65, 0x50, 0x42, 0x2e,
	0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x65, 0x72, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x44, 0x61,
	0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1e, 0x0a,
	0x0a, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x50, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x50, 0x61, 0x74, 0x68, 0x12, 0x18, 0x0a,
	0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x37, 0x0a, 0x09, 0x44, 0x61, 0x74, 0x61, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x22, 0x82, 0x01, 0x0a, 0x18, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x52, 0x6f, 0x75, 0x74, 0x65,
	0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a,
	0x0a, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x50, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x50, 0x61, 0x74, 0x68, 0x12, 0x22, 0x0a,
	0x0c, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x12, 0x22, 0x0a, 0x0c, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x41, 0x75, 0x74, 0x68, 0x49,
	0x46, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x41,
	0x75, 0x74, 0x68, 0x49, 0x46, 0x22, 0x79, 0x0a, 0x11, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x50, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x50, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x50, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x50, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x32, 0x87, 0x01, 0x0a, 0x0d, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x76, 0x0a, 0x15, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x52, 0x6f, 0x75, 0x74,
	0x65, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x65, 0x64, 0x12, 0x2a, 0x2e, 0x50, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x43, 0x6f, 0x72, 0x65, 0x50, 0x42, 0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x65, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x43, 0x6f, 0x72, 0x65, 0x50, 0x42, 0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x52, 0x6f, 0x75,
	0x74, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x11, 0x5a, 0x0f, 0x2e, 0x2e,
	0x2f, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x43, 0x6f, 0x72, 0x65, 0x50, 0x42, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_plugin_proto_rawDescOnce sync.Once
	file_plugin_proto_rawDescData = file_plugin_proto_rawDesc
)

func file_plugin_proto_rawDescGZIP() []byte {
	file_plugin_proto_rawDescOnce.Do(func() {
		file_plugin_proto_rawDescData = protoimpl.X.CompressGZIP(file_plugin_proto_rawDescData)
	})
	return file_plugin_proto_rawDescData
}

var file_plugin_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_plugin_proto_goTypes = []interface{}{
	(*PluginRouteRegisteredRequest)(nil),  // 0: PluginCorePB.PluginRouteRegisteredRequest
	(*PluginRouteRegisteredResponse)(nil), // 1: PluginCorePB.PluginRouteRegisteredResponse
	(*PluginRouterGroupRequest)(nil),      // 2: PluginCorePB.PluginRouterGroupRequest
	(*PluginInfoRequest)(nil),             // 3: PluginCorePB.PluginInfoRequest
	nil,                                   // 4: PluginCorePB.PluginRouteRegisteredResponse.DataEntry
}
var file_plugin_proto_depIdxs = []int32{
	2, // 0: PluginCorePB.PluginRouteRegisteredRequest.PluginRouterGroup:type_name -> PluginCorePB.PluginRouterGroupRequest
	3, // 1: PluginCorePB.PluginRouteRegisteredRequest.PluginInfo:type_name -> PluginCorePB.PluginInfoRequest
	4, // 2: PluginCorePB.PluginRouteRegisteredResponse.Data:type_name -> PluginCorePB.PluginRouteRegisteredResponse.DataEntry
	0, // 3: PluginCorePB.PluginService.PluginRouteRegistered:input_type -> PluginCorePB.PluginRouteRegisteredRequest
	1, // 4: PluginCorePB.PluginService.PluginRouteRegistered:output_type -> PluginCorePB.PluginRouteRegisteredResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_plugin_proto_init() }
func file_plugin_proto_init() {
	if File_plugin_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_plugin_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PluginRouteRegisteredRequest); i {
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
		file_plugin_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PluginRouteRegisteredResponse); i {
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
		file_plugin_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PluginRouterGroupRequest); i {
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
		file_plugin_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PluginInfoRequest); i {
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
			RawDescriptor: file_plugin_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_plugin_proto_goTypes,
		DependencyIndexes: file_plugin_proto_depIdxs,
		MessageInfos:      file_plugin_proto_msgTypes,
	}.Build()
	File_plugin_proto = out.File
	file_plugin_proto_rawDesc = nil
	file_plugin_proto_goTypes = nil
	file_plugin_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// PluginServiceClient is the client API for PluginService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PluginServiceClient interface {
	PluginRouteRegistered(ctx context.Context, opts ...grpc.CallOption) (PluginService_PluginRouteRegisteredClient, error)
}

type pluginServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPluginServiceClient(cc grpc.ClientConnInterface) PluginServiceClient {
	return &pluginServiceClient{cc}
}

func (c *pluginServiceClient) PluginRouteRegistered(ctx context.Context, opts ...grpc.CallOption) (PluginService_PluginRouteRegisteredClient, error) {
	stream, err := c.cc.NewStream(ctx, &_PluginService_serviceDesc.Streams[0], "/PluginCorePB.PluginService/PluginRouteRegistered", opts...)
	if err != nil {
		return nil, err
	}
	x := &pluginServicePluginRouteRegisteredClient{stream}
	return x, nil
}

type PluginService_PluginRouteRegisteredClient interface {
	Send(*PluginRouteRegisteredRequest) error
	Recv() (*PluginRouteRegisteredResponse, error)
	grpc.ClientStream
}

type pluginServicePluginRouteRegisteredClient struct {
	grpc.ClientStream
}

func (x *pluginServicePluginRouteRegisteredClient) Send(m *PluginRouteRegisteredRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *pluginServicePluginRouteRegisteredClient) Recv() (*PluginRouteRegisteredResponse, error) {
	m := new(PluginRouteRegisteredResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PluginServiceServer is the server API for PluginService service.
type PluginServiceServer interface {
	PluginRouteRegistered(PluginService_PluginRouteRegisteredServer) error
}

// UnimplementedPluginServiceServer can be embedded to have forward compatible implementations.
type UnimplementedPluginServiceServer struct {
}

func (*UnimplementedPluginServiceServer) PluginRouteRegistered(PluginService_PluginRouteRegisteredServer) error {
	return status.Errorf(codes.Unimplemented, "method PluginRouteRegistered not implemented")
}

func RegisterPluginServiceServer(s *grpc.Server, srv PluginServiceServer) {
	s.RegisterService(&_PluginService_serviceDesc, srv)
}

func _PluginService_PluginRouteRegistered_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PluginServiceServer).PluginRouteRegistered(&pluginServicePluginRouteRegisteredServer{stream})
}

type PluginService_PluginRouteRegisteredServer interface {
	Send(*PluginRouteRegisteredResponse) error
	Recv() (*PluginRouteRegisteredRequest, error)
	grpc.ServerStream
}

type pluginServicePluginRouteRegisteredServer struct {
	grpc.ServerStream
}

func (x *pluginServicePluginRouteRegisteredServer) Send(m *PluginRouteRegisteredResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *pluginServicePluginRouteRegisteredServer) Recv() (*PluginRouteRegisteredRequest, error) {
	m := new(PluginRouteRegisteredRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _PluginService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "PluginCorePB.PluginService",
	HandlerType: (*PluginServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PluginRouteRegistered",
			Handler:       _PluginService_PluginRouteRegistered_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "plugin.proto",
}
