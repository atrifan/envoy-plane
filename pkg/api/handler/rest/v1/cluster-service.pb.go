// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cluster-service.proto

package v1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

//Hosts message
//last id = 2
type Hosts struct {
	Ip                   string   `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
	Port                 uint32   `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Hosts) Reset()         { *m = Hosts{} }
func (m *Hosts) String() string { return proto.CompactTextString(m) }
func (*Hosts) ProtoMessage()    {}
func (*Hosts) Descriptor() ([]byte, []int) {
	return fileDescriptor_e8681304beda874e, []int{0}
}

func (m *Hosts) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Hosts.Unmarshal(m, b)
}
func (m *Hosts) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Hosts.Marshal(b, m, deterministic)
}
func (m *Hosts) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Hosts.Merge(m, src)
}
func (m *Hosts) XXX_Size() int {
	return xxx_messageInfo_Hosts.Size(m)
}
func (m *Hosts) XXX_DiscardUnknown() {
	xxx_messageInfo_Hosts.DiscardUnknown(m)
}

var xxx_messageInfo_Hosts proto.InternalMessageInfo

func (m *Hosts) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *Hosts) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

// Cluster message
// last id 4
type Cluster struct {
	// Unique clusterName
	ClusterName string `protobuf:"bytes,1,opt,name=clusterName,proto3" json:"clusterName,omitempty"`
	// Type of lb policy
	LbPolicy string `protobuf:"bytes,2,opt,name=LbPolicy,proto3" json:"LbPolicy,omitempty"`
	// list of hosts
	Hosts                []*Hosts `protobuf:"bytes,3,rep,name=hosts,proto3" json:"hosts,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Cluster) Reset()         { *m = Cluster{} }
func (m *Cluster) String() string { return proto.CompactTextString(m) }
func (*Cluster) ProtoMessage()    {}
func (*Cluster) Descriptor() ([]byte, []int) {
	return fileDescriptor_e8681304beda874e, []int{1}
}

func (m *Cluster) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Cluster.Unmarshal(m, b)
}
func (m *Cluster) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Cluster.Marshal(b, m, deterministic)
}
func (m *Cluster) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Cluster.Merge(m, src)
}
func (m *Cluster) XXX_Size() int {
	return xxx_messageInfo_Cluster.Size(m)
}
func (m *Cluster) XXX_DiscardUnknown() {
	xxx_messageInfo_Cluster.DiscardUnknown(m)
}

var xxx_messageInfo_Cluster proto.InternalMessageInfo

func (m *Cluster) GetClusterName() string {
	if m != nil {
		return m.ClusterName
	}
	return ""
}

func (m *Cluster) GetLbPolicy() string {
	if m != nil {
		return m.LbPolicy
	}
	return ""
}

func (m *Cluster) GetHosts() []*Hosts {
	if m != nil {
		return m.Hosts
	}
	return nil
}

// Request data to create new Cluster
type CreateRequest struct {
	// for which node to create the cluster
	NodeId string `protobuf:"bytes,1,opt,name=nodeId,proto3" json:"nodeId,omitempty"`
	// List of clusters
	Cluster []*Cluster `protobuf:"bytes,2,rep,name=cluster,proto3" json:"cluster,omitempty"`
	// the version
	Version              string   `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e8681304beda874e, []int{2}
}

func (m *CreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRequest.Unmarshal(m, b)
}
func (m *CreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRequest.Marshal(b, m, deterministic)
}
func (m *CreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequest.Merge(m, src)
}
func (m *CreateRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRequest.Size(m)
}
func (m *CreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequest proto.InternalMessageInfo

func (m *CreateRequest) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *CreateRequest) GetCluster() []*Cluster {
	if m != nil {
		return m.Cluster
	}
	return nil
}

func (m *CreateRequest) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

// Contains data of created Cluster
// last id 2
type CreateResponse struct {
	// The node where the cluster has been created
	NodeId string `protobuf:"bytes,1,opt,name=nodeId,proto3" json:"nodeId,omitempty"`
	// versions of the created clusters
	Version              string   `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateResponse) Reset()         { *m = CreateResponse{} }
func (m *CreateResponse) String() string { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()    {}
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e8681304beda874e, []int{3}
}

func (m *CreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateResponse.Unmarshal(m, b)
}
func (m *CreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateResponse.Marshal(b, m, deterministic)
}
func (m *CreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateResponse.Merge(m, src)
}
func (m *CreateResponse) XXX_Size() int {
	return xxx_messageInfo_CreateResponse.Size(m)
}
func (m *CreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateResponse proto.InternalMessageInfo

func (m *CreateResponse) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *CreateResponse) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

// Request data to update a cluster
type UpdateRequest struct {
	// node for which we create a cluster
	NodeId string `protobuf:"bytes,1,opt,name=nodeId,proto3" json:"nodeId,omitempty"`
	// Cluster object to update
	Cluster              *Cluster `protobuf:"bytes,2,opt,name=cluster,proto3" json:"cluster,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateRequest) Reset()         { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()    {}
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e8681304beda874e, []int{4}
}

func (m *UpdateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateRequest.Unmarshal(m, b)
}
func (m *UpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateRequest.Marshal(b, m, deterministic)
}
func (m *UpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateRequest.Merge(m, src)
}
func (m *UpdateRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateRequest.Size(m)
}
func (m *UpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateRequest proto.InternalMessageInfo

func (m *UpdateRequest) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *UpdateRequest) GetCluster() *Cluster {
	if m != nil {
		return m.Cluster
	}
	return nil
}

// Contains status of update operation
type UpdateResponse struct {
	// API versioning: it is my best practice to specify version explicitly
	NodeId string `protobuf:"bytes,1,opt,name=nodeId,proto3" json:"nodeId,omitempty"`
	// Contains number of entities have beed updated
	// Equals 1 in case of succesfull update
	Version              string   `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateResponse) Reset()         { *m = UpdateResponse{} }
func (m *UpdateResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateResponse) ProtoMessage()    {}
func (*UpdateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e8681304beda874e, []int{5}
}

func (m *UpdateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateResponse.Unmarshal(m, b)
}
func (m *UpdateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateResponse.Marshal(b, m, deterministic)
}
func (m *UpdateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateResponse.Merge(m, src)
}
func (m *UpdateResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateResponse.Size(m)
}
func (m *UpdateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateResponse proto.InternalMessageInfo

func (m *UpdateResponse) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *UpdateResponse) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

// Request data to delete cluster from a node
type DeleteRequest struct {
	// node id
	NodeId string `protobuf:"bytes,1,opt,name=nodeId,proto3" json:"nodeId,omitempty"`
	// version of cluster to delete
	VersionId            string   `protobuf:"bytes,2,opt,name=versionId,proto3" json:"versionId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRequest) Reset()         { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()    {}
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e8681304beda874e, []int{6}
}

func (m *DeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRequest.Unmarshal(m, b)
}
func (m *DeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRequest.Marshal(b, m, deterministic)
}
func (m *DeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRequest.Merge(m, src)
}
func (m *DeleteRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteRequest.Size(m)
}
func (m *DeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRequest proto.InternalMessageInfo

func (m *DeleteRequest) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *DeleteRequest) GetVersionId() string {
	if m != nil {
		return m.VersionId
	}
	return ""
}

// Contains status of delete operation
type DeleteResponse struct {
	// nodeId response
	NodeId string `protobuf:"bytes,1,opt,name=nodeId,proto3" json:"nodeId,omitempty"`
	// contains the versions of deleted clusters
	Deleted              []string `protobuf:"bytes,2,rep,name=deleted,proto3" json:"deleted,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteResponse) Reset()         { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()    {}
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e8681304beda874e, []int{7}
}

func (m *DeleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteResponse.Unmarshal(m, b)
}
func (m *DeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteResponse.Marshal(b, m, deterministic)
}
func (m *DeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteResponse.Merge(m, src)
}
func (m *DeleteResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteResponse.Size(m)
}
func (m *DeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteResponse proto.InternalMessageInfo

func (m *DeleteResponse) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *DeleteResponse) GetDeleted() []string {
	if m != nil {
		return m.Deleted
	}
	return nil
}

// Request data to read all clusters for a node
type ReadAllRequestForNode struct {
	// API versioning: it is my best practice to specify version explicitly
	NodeId               string   `protobuf:"bytes,1,opt,name=nodeId,proto3" json:"nodeId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadAllRequestForNode) Reset()         { *m = ReadAllRequestForNode{} }
func (m *ReadAllRequestForNode) String() string { return proto.CompactTextString(m) }
func (*ReadAllRequestForNode) ProtoMessage()    {}
func (*ReadAllRequestForNode) Descriptor() ([]byte, []int) {
	return fileDescriptor_e8681304beda874e, []int{8}
}

func (m *ReadAllRequestForNode) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadAllRequestForNode.Unmarshal(m, b)
}
func (m *ReadAllRequestForNode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadAllRequestForNode.Marshal(b, m, deterministic)
}
func (m *ReadAllRequestForNode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadAllRequestForNode.Merge(m, src)
}
func (m *ReadAllRequestForNode) XXX_Size() int {
	return xxx_messageInfo_ReadAllRequestForNode.Size(m)
}
func (m *ReadAllRequestForNode) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadAllRequestForNode.DiscardUnknown(m)
}

var xxx_messageInfo_ReadAllRequestForNode proto.InternalMessageInfo

func (m *ReadAllRequestForNode) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

// Contains list of all clusters on a node
type ReadAllResponseForNode struct {
	// API versioning: it is my best practice to specify version explicitly
	NodeId string `protobuf:"bytes,1,opt,name=nodeId,proto3" json:"nodeId,omitempty"`
	// List of all todo tasks
	Clusters             []*Cluster `protobuf:"bytes,2,rep,name=clusters,proto3" json:"clusters,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ReadAllResponseForNode) Reset()         { *m = ReadAllResponseForNode{} }
func (m *ReadAllResponseForNode) String() string { return proto.CompactTextString(m) }
func (*ReadAllResponseForNode) ProtoMessage()    {}
func (*ReadAllResponseForNode) Descriptor() ([]byte, []int) {
	return fileDescriptor_e8681304beda874e, []int{9}
}

func (m *ReadAllResponseForNode) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadAllResponseForNode.Unmarshal(m, b)
}
func (m *ReadAllResponseForNode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadAllResponseForNode.Marshal(b, m, deterministic)
}
func (m *ReadAllResponseForNode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadAllResponseForNode.Merge(m, src)
}
func (m *ReadAllResponseForNode) XXX_Size() int {
	return xxx_messageInfo_ReadAllResponseForNode.Size(m)
}
func (m *ReadAllResponseForNode) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadAllResponseForNode.DiscardUnknown(m)
}

var xxx_messageInfo_ReadAllResponseForNode proto.InternalMessageInfo

func (m *ReadAllResponseForNode) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *ReadAllResponseForNode) GetClusters() []*Cluster {
	if m != nil {
		return m.Clusters
	}
	return nil
}

// Request to get a specific cluster
type ReadRequest struct {
	// The nodeId where you want to get the cluster from
	NodeId string `protobuf:"bytes,1,opt,name=nodeId,proto3" json:"nodeId,omitempty"`
	// The version of the cluster that you want to retrieve
	Version              string   `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadRequest) Reset()         { *m = ReadRequest{} }
func (m *ReadRequest) String() string { return proto.CompactTextString(m) }
func (*ReadRequest) ProtoMessage()    {}
func (*ReadRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e8681304beda874e, []int{10}
}

func (m *ReadRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadRequest.Unmarshal(m, b)
}
func (m *ReadRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadRequest.Marshal(b, m, deterministic)
}
func (m *ReadRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadRequest.Merge(m, src)
}
func (m *ReadRequest) XXX_Size() int {
	return xxx_messageInfo_ReadRequest.Size(m)
}
func (m *ReadRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReadRequest proto.InternalMessageInfo

func (m *ReadRequest) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *ReadRequest) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

type ReadResponse struct {
	NodeId               string   `protobuf:"bytes,1,opt,name=nodeId,proto3" json:"nodeId,omitempty"`
	Cluster              *Cluster `protobuf:"bytes,2,opt,name=cluster,proto3" json:"cluster,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadResponse) Reset()         { *m = ReadResponse{} }
func (m *ReadResponse) String() string { return proto.CompactTextString(m) }
func (*ReadResponse) ProtoMessage()    {}
func (*ReadResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e8681304beda874e, []int{11}
}

func (m *ReadResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadResponse.Unmarshal(m, b)
}
func (m *ReadResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadResponse.Marshal(b, m, deterministic)
}
func (m *ReadResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadResponse.Merge(m, src)
}
func (m *ReadResponse) XXX_Size() int {
	return xxx_messageInfo_ReadResponse.Size(m)
}
func (m *ReadResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReadResponse proto.InternalMessageInfo

func (m *ReadResponse) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *ReadResponse) GetCluster() *Cluster {
	if m != nil {
		return m.Cluster
	}
	return nil
}

func init() {
	proto.RegisterType((*Hosts)(nil), "v1.Hosts")
	proto.RegisterType((*Cluster)(nil), "v1.Cluster")
	proto.RegisterType((*CreateRequest)(nil), "v1.CreateRequest")
	proto.RegisterType((*CreateResponse)(nil), "v1.CreateResponse")
	proto.RegisterType((*UpdateRequest)(nil), "v1.UpdateRequest")
	proto.RegisterType((*UpdateResponse)(nil), "v1.UpdateResponse")
	proto.RegisterType((*DeleteRequest)(nil), "v1.DeleteRequest")
	proto.RegisterType((*DeleteResponse)(nil), "v1.DeleteResponse")
	proto.RegisterType((*ReadAllRequestForNode)(nil), "v1.ReadAllRequestForNode")
	proto.RegisterType((*ReadAllResponseForNode)(nil), "v1.ReadAllResponseForNode")
	proto.RegisterType((*ReadRequest)(nil), "v1.ReadRequest")
	proto.RegisterType((*ReadResponse)(nil), "v1.ReadResponse")
}

func init() { proto.RegisterFile("cluster-service.proto", fileDescriptor_e8681304beda874e) }

var fileDescriptor_e8681304beda874e = []byte{
	// 652 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x94, 0x4f, 0x6f, 0xd3, 0x30,
	0x14, 0xc0, 0x95, 0x74, 0x6b, 0xd7, 0xd7, 0x75, 0x7f, 0x2c, 0x6d, 0x2b, 0xd1, 0x10, 0x51, 0x00,
	0x31, 0x15, 0xd6, 0xac, 0x65, 0xa7, 0xed, 0x00, 0xdb, 0x00, 0x6d, 0x12, 0x54, 0x28, 0x88, 0x03,
	0x70, 0xca, 0x92, 0x47, 0x1b, 0x94, 0xd9, 0x21, 0x76, 0x3b, 0x10, 0xe2, 0xc2, 0x47, 0x80, 0x1b,
	0x37, 0x0e, 0x7c, 0x10, 0x3e, 0x03, 0x5f, 0x81, 0x0f, 0x82, 0x62, 0x3b, 0x5d, 0x43, 0x57, 0x15,
	0x38, 0xb5, 0x7e, 0xb6, 0x7f, 0xef, 0x97, 0x67, 0x3f, 0xc3, 0x5a, 0x10, 0x0f, 0xb8, 0xc0, 0x74,
	0x9b, 0x63, 0x3a, 0x8c, 0x02, 0x6c, 0x25, 0x29, 0x13, 0x8c, 0x98, 0xc3, 0xb6, 0xb5, 0xd9, 0x63,
	0xac, 0x17, 0xa3, 0xeb, 0x27, 0x91, 0xeb, 0x53, 0xca, 0x84, 0x2f, 0x22, 0x46, 0xb9, 0x5a, 0x61,
	0xdd, 0x91, 0x3f, 0xc1, 0x76, 0x0f, 0xe9, 0x36, 0x3f, 0xf7, 0x7b, 0x3d, 0x4c, 0x5d, 0x96, 0xc8,
	0x15, 0x93, 0xab, 0x9d, 0xdb, 0x30, 0x7f, 0xcc, 0xb8, 0xe0, 0x64, 0x09, 0xcc, 0x28, 0x69, 0x18,
	0xb6, 0xb1, 0x55, 0xf5, 0xcc, 0x28, 0x21, 0x04, 0xe6, 0x12, 0x96, 0x8a, 0x86, 0x69, 0x1b, 0x5b,
	0x75, 0x4f, 0xfe, 0x77, 0xfa, 0x50, 0x39, 0x52, 0x56, 0xc4, 0x86, 0x9a, 0x16, 0xec, 0xfa, 0x67,
	0xa8, 0xf7, 0x8d, 0x87, 0x88, 0x05, 0x0b, 0x8f, 0x4f, 0x9f, 0xb2, 0x38, 0x0a, 0xde, 0x4b, 0x48,
	0xd5, 0x1b, 0x8d, 0xc9, 0x35, 0x98, 0xef, 0x67, 0x59, 0x1b, 0x25, 0xbb, 0xb4, 0x55, 0xeb, 0x54,
	0x5b, 0xc3, 0x76, 0x4b, 0x6a, 0x78, 0x2a, 0xee, 0xf4, 0xa1, 0x7e, 0x94, 0xa2, 0x2f, 0xd0, 0xc3,
	0xb7, 0x03, 0xe4, 0x82, 0xac, 0x43, 0x99, 0xb2, 0x10, 0x4f, 0x42, 0x9d, 0x4a, 0x8f, 0xc8, 0x4d,
	0xa8, 0xe8, 0xa4, 0x0d, 0x53, 0xb2, 0x6a, 0x19, 0x4b, 0x5b, 0x7a, 0xf9, 0x1c, 0x69, 0x40, 0x65,
	0x88, 0x29, 0x8f, 0x18, 0x6d, 0x94, 0xe4, 0xfe, 0x7c, 0xe8, 0x1c, 0xc2, 0x52, 0x9e, 0x89, 0x27,
	0x8c, 0x72, 0x9c, 0x9a, 0x6a, 0x8c, 0x61, 0x16, 0x19, 0x5d, 0xa8, 0x3f, 0x4f, 0xc2, 0x7f, 0xb5,
	0x35, 0xa6, 0xd9, 0x66, 0x4e, 0x39, 0xef, 0xbf, 0x9d, 0x1e, 0x42, 0xfd, 0x01, 0xc6, 0x38, 0xdb,
	0x69, 0x13, 0xaa, 0x7a, 0xcf, 0x49, 0xa8, 0x21, 0x17, 0x81, 0x4c, 0x25, 0xc7, 0xcc, 0x56, 0x09,
	0xe5, 0xca, 0x50, 0x9e, 0x44, 0xd5, 0xcb, 0x87, 0x8e, 0x0b, 0x6b, 0x1e, 0xfa, 0xe1, 0x41, 0x1c,
	0x6b, 0x97, 0x47, 0x2c, 0xed, 0xb2, 0x70, 0x2a, 0xca, 0x79, 0x01, 0xeb, 0xa3, 0x0d, 0x2a, 0xeb,
	0x8c, 0x1d, 0xe4, 0x16, 0x2c, 0xe8, 0xe2, 0xf1, 0xcb, 0xee, 0xc1, 0x68, 0xd2, 0xb9, 0x07, 0xb5,
	0x0c, 0x3d, 0xab, 0x28, 0xd3, 0xeb, 0xfa, 0x04, 0x16, 0x15, 0x60, 0x46, 0x39, 0xfe, 0xee, 0xa8,
	0x3b, 0x3f, 0x4a, 0xb0, 0xa4, 0x83, 0xcf, 0x54, 0xa3, 0x13, 0x3a, 0xfa, 0x7a, 0x3d, 0xc1, 0xf3,
	0xaf, 0xbf, 0x92, 0x21, 0x2e, 0x2d, 0xa5, 0x65, 0x15, 0xa6, 0x0a, 0x45, 0x73, 0xae, 0x7e, 0xfa,
	0xf9, 0xeb, 0x8b, 0xb9, 0x41, 0xd6, 0xdc, 0x61, 0xdb, 0xcd, 0x2b, 0xe1, 0x7e, 0x50, 0xa2, 0x1f,
	0xc9, 0x31, 0x94, 0x55, 0x07, 0x90, 0x55, 0xa9, 0x38, 0xde, 0x77, 0x16, 0x19, 0x0f, 0x29, 0xac,
	0xb3, 0x21, 0x79, 0xab, 0xce, 0xe2, 0x38, 0x6f, 0xcf, 0x68, 0x92, 0x63, 0x98, 0xcb, 0x14, 0xc8,
	0x72, 0x2e, 0x93, 0x53, 0x56, 0x2e, 0x02, 0x9a, 0x31, 0xc3, 0xe9, 0x15, 0x94, 0x55, 0x07, 0x28,
	0xa7, 0x42, 0x77, 0x29, 0xa7, 0x62, 0x83, 0x38, 0x4d, 0xc9, 0xbb, 0x61, 0xfd, 0xe9, 0xf4, 0x72,
	0xb5, 0x33, 0xa1, 0xd9, 0x85, 0xb2, 0xba, 0xd3, 0x0a, 0x5e, 0x68, 0x13, 0x05, 0x2f, 0x5e, 0xf9,
	0x5c, 0xb6, 0x79, 0xb9, 0xec, 0xe1, 0x77, 0xe3, 0xf3, 0xc1, 0x37, 0x83, 0xec, 0xc3, 0xb2, 0x3e,
	0x30, 0x5b, 0xbf, 0xd9, 0xce, 0xf5, 0x89, 0x90, 0xb5, 0x22, 0xd2, 0xe8, 0xb5, 0x4f, 0xef, 0xfb,
	0x21, 0x3b, 0xc5, 0x56, 0xc0, 0xce, 0x3a, 0xa5, 0x76, 0x6b, 0xa7, 0x69, 0x18, 0x9d, 0x15, 0x3f,
	0x49, 0xe2, 0x28, 0x90, 0xcf, 0xb3, 0xfb, 0x86, 0x33, 0xba, 0x37, 0x11, 0xf1, 0xf6, 0xa1, 0xb4,
	0xbb, 0xb3, 0x4b, 0x76, 0xa1, 0xe9, 0xa1, 0x18, 0xa4, 0x14, 0x43, 0xfb, 0xbc, 0x8f, 0xd4, 0x16,
	0x7d, 0xb4, 0x53, 0xe4, 0x6c, 0x90, 0x06, 0x68, 0x87, 0x0c, 0xb9, 0x4d, 0x99, 0xb0, 0xf1, 0x5d,
	0xc4, 0x45, 0x8b, 0x94, 0x61, 0xee, 0xab, 0x69, 0x54, 0x4e, 0xcb, 0xf2, 0xc9, 0xbf, 0xfb, 0x3b,
	0x00, 0x00, 0xff, 0xff, 0x4a, 0x0f, 0xdb, 0x7f, 0x5b, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ClusterServiceClient is the client API for ClusterService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ClusterServiceClient interface {
	// Read all clusters for a node
	ReadAllClustersForNode(ctx context.Context, in *ReadAllRequestForNode, opts ...grpc.CallOption) (*ReadAllResponseForNode, error)
	// Create new clusters
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	// Read cluster from node
	Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadResponse, error)
	// Update cluster on node
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
	// Delete cluster from node
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
}

type clusterServiceClient struct {
	cc *grpc.ClientConn
}

func NewClusterServiceClient(cc *grpc.ClientConn) ClusterServiceClient {
	return &clusterServiceClient{cc}
}

func (c *clusterServiceClient) ReadAllClustersForNode(ctx context.Context, in *ReadAllRequestForNode, opts ...grpc.CallOption) (*ReadAllResponseForNode, error) {
	out := new(ReadAllResponseForNode)
	err := c.cc.Invoke(ctx, "/v1.ClusterService/ReadAllClustersForNode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterServiceClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/v1.ClusterService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterServiceClient) Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadResponse, error) {
	out := new(ReadResponse)
	err := c.cc.Invoke(ctx, "/v1.ClusterService/Read", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterServiceClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, "/v1.ClusterService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterServiceClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/v1.ClusterService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClusterServiceServer is the server API for ClusterService service.
type ClusterServiceServer interface {
	// Read all clusters for a node
	ReadAllClustersForNode(context.Context, *ReadAllRequestForNode) (*ReadAllResponseForNode, error)
	// Create new clusters
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	// Read cluster from node
	Read(context.Context, *ReadRequest) (*ReadResponse, error)
	// Update cluster on node
	Update(context.Context, *UpdateRequest) (*UpdateResponse, error)
	// Delete cluster from node
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
}

// UnimplementedClusterServiceServer can be embedded to have forward compatible implementations.
type UnimplementedClusterServiceServer struct {
}

func (*UnimplementedClusterServiceServer) ReadAllClustersForNode(ctx context.Context, req *ReadAllRequestForNode) (*ReadAllResponseForNode, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadAllClustersForNode not implemented")
}
func (*UnimplementedClusterServiceServer) Create(ctx context.Context, req *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedClusterServiceServer) Read(ctx context.Context, req *ReadRequest) (*ReadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Read not implemented")
}
func (*UnimplementedClusterServiceServer) Update(ctx context.Context, req *UpdateRequest) (*UpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedClusterServiceServer) Delete(ctx context.Context, req *DeleteRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

func RegisterClusterServiceServer(s *grpc.Server, srv ClusterServiceServer) {
	s.RegisterService(&_ClusterService_serviceDesc, srv)
}

func _ClusterService_ReadAllClustersForNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadAllRequestForNode)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServiceServer).ReadAllClustersForNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.ClusterService/ReadAllClustersForNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServiceServer).ReadAllClustersForNode(ctx, req.(*ReadAllRequestForNode))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.ClusterService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServiceServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterService_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServiceServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.ClusterService/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServiceServer).Read(ctx, req.(*ReadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.ClusterService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServiceServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.ClusterService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServiceServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ClusterService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1.ClusterService",
	HandlerType: (*ClusterServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReadAllClustersForNode",
			Handler:    _ClusterService_ReadAllClustersForNode_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _ClusterService_Create_Handler,
		},
		{
			MethodName: "Read",
			Handler:    _ClusterService_Read_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ClusterService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ClusterService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cluster-service.proto",
}
