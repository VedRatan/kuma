// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.0
// source: mesh/v1alpha1/kds.proto

package v1alpha1

import (
	v3 "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v3"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type KumaResource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Meta *KumaResource_Meta `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	Spec *anypb.Any         `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec,omitempty"`
}

func (x *KumaResource) Reset() {
	*x = KumaResource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mesh_v1alpha1_kds_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KumaResource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KumaResource) ProtoMessage() {}

func (x *KumaResource) ProtoReflect() protoreflect.Message {
	mi := &file_mesh_v1alpha1_kds_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KumaResource.ProtoReflect.Descriptor instead.
func (*KumaResource) Descriptor() ([]byte, []int) {
	return file_mesh_v1alpha1_kds_proto_rawDescGZIP(), []int{0}
}

func (x *KumaResource) GetMeta() *KumaResource_Meta {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *KumaResource) GetSpec() *anypb.Any {
	if x != nil {
		return x.Spec
	}
	return nil
}

// XDSConfigRequest is a request for XDS Config Dump that is executed on Zone
// CP.
type XDSConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// RequestID is a UUID of a request so we can correlate requests with response
	// on one stream.
	RequestId string `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	// Type of resource (Dataplane, ZoneIngress, ZoneEgress)
	ResourceType string `protobuf:"bytes,2,opt,name=resource_type,json=resourceType,proto3" json:"resource_type,omitempty"`
	// Name of the resource on which we execute config dump.
	ResourceName string `protobuf:"bytes,3,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Mesh of the resource on which we execute config dump. Should be empty for
	// ZoneIngress, ZoneEgress.
	ResourceMesh string `protobuf:"bytes,4,opt,name=resource_mesh,json=resourceMesh,proto3" json:"resource_mesh,omitempty"`
}

func (x *XDSConfigRequest) Reset() {
	*x = XDSConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mesh_v1alpha1_kds_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *XDSConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*XDSConfigRequest) ProtoMessage() {}

func (x *XDSConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mesh_v1alpha1_kds_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use XDSConfigRequest.ProtoReflect.Descriptor instead.
func (*XDSConfigRequest) Descriptor() ([]byte, []int) {
	return file_mesh_v1alpha1_kds_proto_rawDescGZIP(), []int{1}
}

func (x *XDSConfigRequest) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *XDSConfigRequest) GetResourceType() string {
	if x != nil {
		return x.ResourceType
	}
	return ""
}

func (x *XDSConfigRequest) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *XDSConfigRequest) GetResourceMesh() string {
	if x != nil {
		return x.ResourceMesh
	}
	return ""
}

// XDSConfigRequest is a response containing result of XDS Config Dump execution
// on Zone CP.
type XDSConfigResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// RequestID is a UUID that was set by the Global CP.
	RequestId string `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	// Types that are assignable to Result:
	//
	//	*XDSConfigResponse_Error
	//	*XDSConfigResponse_Config
	Result isXDSConfigResponse_Result `protobuf_oneof:"result"`
}

func (x *XDSConfigResponse) Reset() {
	*x = XDSConfigResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mesh_v1alpha1_kds_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *XDSConfigResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*XDSConfigResponse) ProtoMessage() {}

func (x *XDSConfigResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mesh_v1alpha1_kds_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use XDSConfigResponse.ProtoReflect.Descriptor instead.
func (*XDSConfigResponse) Descriptor() ([]byte, []int) {
	return file_mesh_v1alpha1_kds_proto_rawDescGZIP(), []int{2}
}

func (x *XDSConfigResponse) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (m *XDSConfigResponse) GetResult() isXDSConfigResponse_Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (x *XDSConfigResponse) GetError() string {
	if x, ok := x.GetResult().(*XDSConfigResponse_Error); ok {
		return x.Error
	}
	return ""
}

func (x *XDSConfigResponse) GetConfig() []byte {
	if x, ok := x.GetResult().(*XDSConfigResponse_Config); ok {
		return x.Config
	}
	return nil
}

type isXDSConfigResponse_Result interface {
	isXDSConfigResponse_Result()
}

type XDSConfigResponse_Error struct {
	// Error that was captured by the Zone CP when executing XDS Config Dump.
	Error string `protobuf:"bytes,2,opt,name=error,proto3,oneof"`
}

type XDSConfigResponse_Config struct {
	// The XDS Config that is a successful result of XDS Config dump execution.
	Config []byte `protobuf:"bytes,3,opt,name=config,proto3,oneof"`
}

func (*XDSConfigResponse_Error) isXDSConfigResponse_Result() {}

func (*XDSConfigResponse_Config) isXDSConfigResponse_Result() {}

// StatsRequest is a request for kuma-dp stats that is executed on Zone CP.
type StatsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// RequestID is a UUID of a request so we can correlate requests with response
	// on one stream.
	RequestId string `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	// Type of resource (Dataplane, ZoneIngress, ZoneEgress)
	ResourceType string `protobuf:"bytes,2,opt,name=resource_type,json=resourceType,proto3" json:"resource_type,omitempty"`
	// Name of the resource on which we execute kuma-dp stats request.
	ResourceName string `protobuf:"bytes,3,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Mesh of the resource on which we execute kuma-dp stats request.
	// Should be empty for ZoneIngress, ZoneEgress.
	ResourceMesh string `protobuf:"bytes,4,opt,name=resource_mesh,json=resourceMesh,proto3" json:"resource_mesh,omitempty"`
}

func (x *StatsRequest) Reset() {
	*x = StatsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mesh_v1alpha1_kds_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatsRequest) ProtoMessage() {}

func (x *StatsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mesh_v1alpha1_kds_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatsRequest.ProtoReflect.Descriptor instead.
func (*StatsRequest) Descriptor() ([]byte, []int) {
	return file_mesh_v1alpha1_kds_proto_rawDescGZIP(), []int{3}
}

func (x *StatsRequest) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *StatsRequest) GetResourceType() string {
	if x != nil {
		return x.ResourceType
	}
	return ""
}

func (x *StatsRequest) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *StatsRequest) GetResourceMesh() string {
	if x != nil {
		return x.ResourceMesh
	}
	return ""
}

// StatsResponse is a response containing result of kuma-dp stats execution on
// Zone CP.
type StatsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// RequestID is a UUID that was set by the Global CP.
	RequestId string `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	// Types that are assignable to Result:
	//
	//	*StatsResponse_Error
	//	*StatsResponse_Stats
	Result isStatsResponse_Result `protobuf_oneof:"result"`
}

func (x *StatsResponse) Reset() {
	*x = StatsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mesh_v1alpha1_kds_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatsResponse) ProtoMessage() {}

func (x *StatsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mesh_v1alpha1_kds_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatsResponse.ProtoReflect.Descriptor instead.
func (*StatsResponse) Descriptor() ([]byte, []int) {
	return file_mesh_v1alpha1_kds_proto_rawDescGZIP(), []int{4}
}

func (x *StatsResponse) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (m *StatsResponse) GetResult() isStatsResponse_Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (x *StatsResponse) GetError() string {
	if x, ok := x.GetResult().(*StatsResponse_Error); ok {
		return x.Error
	}
	return ""
}

func (x *StatsResponse) GetStats() []byte {
	if x, ok := x.GetResult().(*StatsResponse_Stats); ok {
		return x.Stats
	}
	return nil
}

type isStatsResponse_Result interface {
	isStatsResponse_Result()
}

type StatsResponse_Error struct {
	// Error that was captured by the Zone CP when executing kuma-dp stats
	// request.
	Error string `protobuf:"bytes,2,opt,name=error,proto3,oneof"`
}

type StatsResponse_Stats struct {
	// The stats content that is a successful result of kuma-dp stats execution.
	Stats []byte `protobuf:"bytes,3,opt,name=stats,proto3,oneof"`
}

func (*StatsResponse_Error) isStatsResponse_Result() {}

func (*StatsResponse_Stats) isStatsResponse_Result() {}

// ClustersRequest is a request for kuma-dp clusters that is executed on Zone
// CP.
type ClustersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// RequestID is a UUID of a request so we can correlate requests with response
	// on one stream.
	RequestId string `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	// Type of resource (Dataplane, ZoneIngress, ZoneEgress)
	ResourceType string `protobuf:"bytes,2,opt,name=resource_type,json=resourceType,proto3" json:"resource_type,omitempty"`
	// Name of the resource on which we execute kuma-dp clusters request.
	ResourceName string `protobuf:"bytes,3,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Mesh of the resource on which we execute kuma-dp clusters request.
	// Should be empty for ZoneIngress, ZoneEgress.
	ResourceMesh string `protobuf:"bytes,4,opt,name=resource_mesh,json=resourceMesh,proto3" json:"resource_mesh,omitempty"`
}

func (x *ClustersRequest) Reset() {
	*x = ClustersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mesh_v1alpha1_kds_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClustersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClustersRequest) ProtoMessage() {}

func (x *ClustersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mesh_v1alpha1_kds_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClustersRequest.ProtoReflect.Descriptor instead.
func (*ClustersRequest) Descriptor() ([]byte, []int) {
	return file_mesh_v1alpha1_kds_proto_rawDescGZIP(), []int{5}
}

func (x *ClustersRequest) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *ClustersRequest) GetResourceType() string {
	if x != nil {
		return x.ResourceType
	}
	return ""
}

func (x *ClustersRequest) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *ClustersRequest) GetResourceMesh() string {
	if x != nil {
		return x.ResourceMesh
	}
	return ""
}

// ClustersResponse is a response containing result of kuma-dp clusters
// execution on Zone CP.
type ClustersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// RequestID is a UUID that was set by the Global CP.
	RequestId string `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	// Types that are assignable to Result:
	//
	//	*ClustersResponse_Error
	//	*ClustersResponse_Clusters
	Result isClustersResponse_Result `protobuf_oneof:"result"`
}

func (x *ClustersResponse) Reset() {
	*x = ClustersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mesh_v1alpha1_kds_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClustersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClustersResponse) ProtoMessage() {}

func (x *ClustersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mesh_v1alpha1_kds_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClustersResponse.ProtoReflect.Descriptor instead.
func (*ClustersResponse) Descriptor() ([]byte, []int) {
	return file_mesh_v1alpha1_kds_proto_rawDescGZIP(), []int{6}
}

func (x *ClustersResponse) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (m *ClustersResponse) GetResult() isClustersResponse_Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (x *ClustersResponse) GetError() string {
	if x, ok := x.GetResult().(*ClustersResponse_Error); ok {
		return x.Error
	}
	return ""
}

func (x *ClustersResponse) GetClusters() []byte {
	if x, ok := x.GetResult().(*ClustersResponse_Clusters); ok {
		return x.Clusters
	}
	return nil
}

type isClustersResponse_Result interface {
	isClustersResponse_Result()
}

type ClustersResponse_Error struct {
	// Error that was captured by the Zone CP when executing kuma-dp clusters
	// request.
	Error string `protobuf:"bytes,2,opt,name=error,proto3,oneof"`
}

type ClustersResponse_Clusters struct {
	// The clusters content that is a successful result of kuma-dp clusters
	// execution.
	Clusters []byte `protobuf:"bytes,3,opt,name=clusters,proto3,oneof"`
}

func (*ClustersResponse_Error) isClustersResponse_Result() {}

func (*ClustersResponse_Clusters) isClustersResponse_Result() {}

type KumaResource_Meta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Mesh    string `protobuf:"bytes,2,opt,name=mesh,proto3" json:"mesh,omitempty"`
	Version string `protobuf:"bytes,5,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *KumaResource_Meta) Reset() {
	*x = KumaResource_Meta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mesh_v1alpha1_kds_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KumaResource_Meta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KumaResource_Meta) ProtoMessage() {}

func (x *KumaResource_Meta) ProtoReflect() protoreflect.Message {
	mi := &file_mesh_v1alpha1_kds_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KumaResource_Meta.ProtoReflect.Descriptor instead.
func (*KumaResource_Meta) Descriptor() ([]byte, []int) {
	return file_mesh_v1alpha1_kds_proto_rawDescGZIP(), []int{0, 0}
}

func (x *KumaResource_Meta) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *KumaResource_Meta) GetMesh() string {
	if x != nil {
		return x.Mesh
	}
	return ""
}

func (x *KumaResource_Meta) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

var File_mesh_v1alpha1_kds_proto protoreflect.FileDescriptor

var file_mesh_v1alpha1_kds_proto_rawDesc = []byte{
	0x0a, 0x17, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f,
	0x6b, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x6b, 0x75, 0x6d, 0x61, 0x2e,
	0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x2a, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x64, 0x69, 0x73,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2f, 0x76, 0x33, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76,
	0x65, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbd, 0x01, 0x0a, 0x0c, 0x4b, 0x75, 0x6d, 0x61, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x39, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x6b, 0x75, 0x6d, 0x61, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4b, 0x75, 0x6d, 0x61, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61,
	0x12, 0x28, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x41, 0x6e, 0x79, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x1a, 0x48, 0x0a, 0x04, 0x4d, 0x65,
	0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x65, 0x73, 0x68, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x65, 0x73, 0x68, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x22, 0xa0, 0x01, 0x0a, 0x10, 0x58, 0x44, 0x53, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x23, 0x0a,
	0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6d,
	0x65, 0x73, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x4d, 0x65, 0x73, 0x68, 0x22, 0x6e, 0x0a, 0x11, 0x58, 0x44, 0x53, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x12, 0x18, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x08, 0x0a,
	0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x9c, 0x01, 0x0a, 0x0c, 0x53, 0x74, 0x61, 0x74,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x23, 0x0a, 0x0d,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6d, 0x65,
	0x73, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x4d, 0x65, 0x73, 0x68, 0x22, 0x68, 0x0a, 0x0d, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x16,
	0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52,
	0x05, 0x73, 0x74, 0x61, 0x74, 0x73, 0x42, 0x08, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x22, 0x9f, 0x01, 0x0a, 0x0f, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a,
	0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6d, 0x65, 0x73, 0x68, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4d, 0x65,
	0x73, 0x68, 0x22, 0x71, 0x0a, 0x10, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x1c, 0x0a,
	0x08, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x48,
	0x00, 0x52, 0x08, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x42, 0x08, 0x0a, 0x06, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0x8e, 0x01, 0x0a, 0x14, 0x4b, 0x75, 0x6d, 0x61, 0x44, 0x69,
	0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x76,
	0x0a, 0x13, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4b, 0x75, 0x6d, 0x61, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x12, 0x2c, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e,
	0x76, 0x33, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x33,
	0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x28, 0x01, 0x30, 0x01, 0x32, 0xb0, 0x02, 0x0a, 0x10, 0x47, 0x6c, 0x6f, 0x62, 0x61,
	0x6c, 0x4b, 0x44, 0x53, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x63, 0x0a, 0x10, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x58, 0x44, 0x53, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x12,
	0x25, 0x2e, 0x6b, 0x75, 0x6d, 0x61, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x58, 0x44, 0x53, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x1a, 0x24, 0x2e, 0x6b, 0x75, 0x6d, 0x61, 0x2e, 0x6d, 0x65,
	0x73, 0x68, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x58, 0x44, 0x53, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x28, 0x01, 0x30, 0x01,
	0x12, 0x56, 0x0a, 0x0b, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12,
	0x21, 0x2e, 0x6b, 0x75, 0x6d, 0x61, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x1a, 0x20, 0x2e, 0x6b, 0x75, 0x6d, 0x61, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x28, 0x01, 0x30, 0x01, 0x12, 0x5f, 0x0a, 0x0e, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x12, 0x24, 0x2e, 0x6b, 0x75, 0x6d,
	0x61, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x1a, 0x23, 0x2e, 0x6b, 0x75, 0x6d, 0x61, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x28, 0x01, 0x30, 0x01, 0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x75, 0x6d, 0x61, 0x68, 0x71, 0x2f, 0x6b,
	0x75, 0x6d, 0x61, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mesh_v1alpha1_kds_proto_rawDescOnce sync.Once
	file_mesh_v1alpha1_kds_proto_rawDescData = file_mesh_v1alpha1_kds_proto_rawDesc
)

func file_mesh_v1alpha1_kds_proto_rawDescGZIP() []byte {
	file_mesh_v1alpha1_kds_proto_rawDescOnce.Do(func() {
		file_mesh_v1alpha1_kds_proto_rawDescData = protoimpl.X.CompressGZIP(file_mesh_v1alpha1_kds_proto_rawDescData)
	})
	return file_mesh_v1alpha1_kds_proto_rawDescData
}

var file_mesh_v1alpha1_kds_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_mesh_v1alpha1_kds_proto_goTypes = []interface{}{
	(*KumaResource)(nil),         // 0: kuma.mesh.v1alpha1.KumaResource
	(*XDSConfigRequest)(nil),     // 1: kuma.mesh.v1alpha1.XDSConfigRequest
	(*XDSConfigResponse)(nil),    // 2: kuma.mesh.v1alpha1.XDSConfigResponse
	(*StatsRequest)(nil),         // 3: kuma.mesh.v1alpha1.StatsRequest
	(*StatsResponse)(nil),        // 4: kuma.mesh.v1alpha1.StatsResponse
	(*ClustersRequest)(nil),      // 5: kuma.mesh.v1alpha1.ClustersRequest
	(*ClustersResponse)(nil),     // 6: kuma.mesh.v1alpha1.ClustersResponse
	(*KumaResource_Meta)(nil),    // 7: kuma.mesh.v1alpha1.KumaResource.Meta
	(*anypb.Any)(nil),            // 8: google.protobuf.Any
	(*v3.DiscoveryRequest)(nil),  // 9: envoy.service.discovery.v3.DiscoveryRequest
	(*v3.DiscoveryResponse)(nil), // 10: envoy.service.discovery.v3.DiscoveryResponse
}
var file_mesh_v1alpha1_kds_proto_depIdxs = []int32{
	7,  // 0: kuma.mesh.v1alpha1.KumaResource.meta:type_name -> kuma.mesh.v1alpha1.KumaResource.Meta
	8,  // 1: kuma.mesh.v1alpha1.KumaResource.spec:type_name -> google.protobuf.Any
	9,  // 2: kuma.mesh.v1alpha1.KumaDiscoveryService.StreamKumaResources:input_type -> envoy.service.discovery.v3.DiscoveryRequest
	2,  // 3: kuma.mesh.v1alpha1.GlobalKDSService.StreamXDSConfigs:input_type -> kuma.mesh.v1alpha1.XDSConfigResponse
	4,  // 4: kuma.mesh.v1alpha1.GlobalKDSService.StreamStats:input_type -> kuma.mesh.v1alpha1.StatsResponse
	6,  // 5: kuma.mesh.v1alpha1.GlobalKDSService.StreamClusters:input_type -> kuma.mesh.v1alpha1.ClustersResponse
	10, // 6: kuma.mesh.v1alpha1.KumaDiscoveryService.StreamKumaResources:output_type -> envoy.service.discovery.v3.DiscoveryResponse
	1,  // 7: kuma.mesh.v1alpha1.GlobalKDSService.StreamXDSConfigs:output_type -> kuma.mesh.v1alpha1.XDSConfigRequest
	3,  // 8: kuma.mesh.v1alpha1.GlobalKDSService.StreamStats:output_type -> kuma.mesh.v1alpha1.StatsRequest
	5,  // 9: kuma.mesh.v1alpha1.GlobalKDSService.StreamClusters:output_type -> kuma.mesh.v1alpha1.ClustersRequest
	6,  // [6:10] is the sub-list for method output_type
	2,  // [2:6] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_mesh_v1alpha1_kds_proto_init() }
func file_mesh_v1alpha1_kds_proto_init() {
	if File_mesh_v1alpha1_kds_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mesh_v1alpha1_kds_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KumaResource); i {
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
		file_mesh_v1alpha1_kds_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*XDSConfigRequest); i {
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
		file_mesh_v1alpha1_kds_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*XDSConfigResponse); i {
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
		file_mesh_v1alpha1_kds_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatsRequest); i {
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
		file_mesh_v1alpha1_kds_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatsResponse); i {
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
		file_mesh_v1alpha1_kds_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClustersRequest); i {
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
		file_mesh_v1alpha1_kds_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClustersResponse); i {
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
		file_mesh_v1alpha1_kds_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KumaResource_Meta); i {
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
	file_mesh_v1alpha1_kds_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*XDSConfigResponse_Error)(nil),
		(*XDSConfigResponse_Config)(nil),
	}
	file_mesh_v1alpha1_kds_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*StatsResponse_Error)(nil),
		(*StatsResponse_Stats)(nil),
	}
	file_mesh_v1alpha1_kds_proto_msgTypes[6].OneofWrappers = []interface{}{
		(*ClustersResponse_Error)(nil),
		(*ClustersResponse_Clusters)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_mesh_v1alpha1_kds_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_mesh_v1alpha1_kds_proto_goTypes,
		DependencyIndexes: file_mesh_v1alpha1_kds_proto_depIdxs,
		MessageInfos:      file_mesh_v1alpha1_kds_proto_msgTypes,
	}.Build()
	File_mesh_v1alpha1_kds_proto = out.File
	file_mesh_v1alpha1_kds_proto_rawDesc = nil
	file_mesh_v1alpha1_kds_proto_goTypes = nil
	file_mesh_v1alpha1_kds_proto_depIdxs = nil
}
