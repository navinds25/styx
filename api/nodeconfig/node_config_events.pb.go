// Code generated by protoc-gen-go. DO NOT EDIT.
// source: node_config_events.proto

package nodeconfig

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type SFTPAuth_SFTPAuthType int32

const (
	SFTPAuth_SFTP_AUTH_TYPE_UNSPECIFIED SFTPAuth_SFTPAuthType = 0
	SFTPAuth_PASSWORD                   SFTPAuth_SFTPAuthType = 1
	SFTPAuth_KEYBASED                   SFTPAuth_SFTPAuthType = 2
	SFTPAuth_KEYPASSPHRASE              SFTPAuth_SFTPAuthType = 3
)

var SFTPAuth_SFTPAuthType_name = map[int32]string{
	0: "SFTP_AUTH_TYPE_UNSPECIFIED",
	1: "PASSWORD",
	2: "KEYBASED",
	3: "KEYPASSPHRASE",
}

var SFTPAuth_SFTPAuthType_value = map[string]int32{
	"SFTP_AUTH_TYPE_UNSPECIFIED": 0,
	"PASSWORD":                   1,
	"KEYBASED":                   2,
	"KEYPASSPHRASE":              3,
}

func (x SFTPAuth_SFTPAuthType) String() string {
	return proto.EnumName(SFTPAuth_SFTPAuthType_name, int32(x))
}

func (SFTPAuth_SFTPAuthType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_98a96787b4b9490b, []int{1, 0}
}

type NodeConfig_NodeType int32

const (
	NodeConfig_NODETYPE_UNSPECIFIED NodeConfig_NodeType = 0
	NodeConfig_INTERNAL             NodeConfig_NodeType = 1
	NodeConfig_EXTERNAL             NodeConfig_NodeType = 2
)

var NodeConfig_NodeType_name = map[int32]string{
	0: "NODETYPE_UNSPECIFIED",
	1: "INTERNAL",
	2: "EXTERNAL",
}

var NodeConfig_NodeType_value = map[string]int32{
	"NODETYPE_UNSPECIFIED": 0,
	"INTERNAL":             1,
	"EXTERNAL":             2,
}

func (x NodeConfig_NodeType) String() string {
	return proto.EnumName(NodeConfig_NodeType_name, int32(x))
}

func (NodeConfig_NodeType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_98a96787b4b9490b, []int{2, 0}
}

type NodeConfig_EnvSec int32

const (
	NodeConfig_ENVSEC_UNSPECIFIED NodeConfig_EnvSec = 0
	NodeConfig_DMZ                NodeConfig_EnvSec = 1
	NodeConfig_APPTIER            NodeConfig_EnvSec = 2
	NodeConfig_DATATIER           NodeConfig_EnvSec = 3
	NodeConfig_CONTROLPLANE       NodeConfig_EnvSec = 4
)

var NodeConfig_EnvSec_name = map[int32]string{
	0: "ENVSEC_UNSPECIFIED",
	1: "DMZ",
	2: "APPTIER",
	3: "DATATIER",
	4: "CONTROLPLANE",
}

var NodeConfig_EnvSec_value = map[string]int32{
	"ENVSEC_UNSPECIFIED": 0,
	"DMZ":                1,
	"APPTIER":            2,
	"DATATIER":           3,
	"CONTROLPLANE":       4,
}

func (x NodeConfig_EnvSec) String() string {
	return proto.EnumName(NodeConfig_EnvSec_name, int32(x))
}

func (NodeConfig_EnvSec) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_98a96787b4b9490b, []int{2, 1}
}

type AddNodeConfigResponse_NodeEntryStatus int32

const (
	AddNodeConfigResponse_STATUS_UNSPECIFIED AddNodeConfigResponse_NodeEntryStatus = 0
	AddNodeConfigResponse_STATUS_SUCCESS     AddNodeConfigResponse_NodeEntryStatus = 1
	AddNodeConfigResponse_STATUS_ERROR       AddNodeConfigResponse_NodeEntryStatus = 2
)

var AddNodeConfigResponse_NodeEntryStatus_name = map[int32]string{
	0: "STATUS_UNSPECIFIED",
	1: "STATUS_SUCCESS",
	2: "STATUS_ERROR",
}

var AddNodeConfigResponse_NodeEntryStatus_value = map[string]int32{
	"STATUS_UNSPECIFIED": 0,
	"STATUS_SUCCESS":     1,
	"STATUS_ERROR":       2,
}

func (x AddNodeConfigResponse_NodeEntryStatus) String() string {
	return proto.EnumName(AddNodeConfigResponse_NodeEntryStatus_name, int32(x))
}

func (AddNodeConfigResponse_NodeEntryStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_98a96787b4b9490b, []int{3, 0}
}

// node config events
type GRPCAuth struct {
	TlsCert              string   `protobuf:"bytes,1,opt,name=tls_cert,json=tlsCert,proto3" json:"tls_cert,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GRPCAuth) Reset()         { *m = GRPCAuth{} }
func (m *GRPCAuth) String() string { return proto.CompactTextString(m) }
func (*GRPCAuth) ProtoMessage()    {}
func (*GRPCAuth) Descriptor() ([]byte, []int) {
	return fileDescriptor_98a96787b4b9490b, []int{0}
}

func (m *GRPCAuth) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GRPCAuth.Unmarshal(m, b)
}
func (m *GRPCAuth) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GRPCAuth.Marshal(b, m, deterministic)
}
func (m *GRPCAuth) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GRPCAuth.Merge(m, src)
}
func (m *GRPCAuth) XXX_Size() int {
	return xxx_messageInfo_GRPCAuth.Size(m)
}
func (m *GRPCAuth) XXX_DiscardUnknown() {
	xxx_messageInfo_GRPCAuth.DiscardUnknown(m)
}

var xxx_messageInfo_GRPCAuth proto.InternalMessageInfo

func (m *GRPCAuth) GetTlsCert() string {
	if m != nil {
		return m.TlsCert
	}
	return ""
}

type SFTPAuth struct {
	SftpAuthType         SFTPAuth_SFTPAuthType `protobuf:"varint,1,opt,name=sftp_auth_type,json=sftpAuthType,proto3,enum=nodeconfig.SFTPAuth_SFTPAuthType" json:"sftp_auth_type,omitempty"`
	Username             string                `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Password             string                `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	Key                  string                `protobuf:"bytes,4,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *SFTPAuth) Reset()         { *m = SFTPAuth{} }
func (m *SFTPAuth) String() string { return proto.CompactTextString(m) }
func (*SFTPAuth) ProtoMessage()    {}
func (*SFTPAuth) Descriptor() ([]byte, []int) {
	return fileDescriptor_98a96787b4b9490b, []int{1}
}

func (m *SFTPAuth) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SFTPAuth.Unmarshal(m, b)
}
func (m *SFTPAuth) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SFTPAuth.Marshal(b, m, deterministic)
}
func (m *SFTPAuth) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SFTPAuth.Merge(m, src)
}
func (m *SFTPAuth) XXX_Size() int {
	return xxx_messageInfo_SFTPAuth.Size(m)
}
func (m *SFTPAuth) XXX_DiscardUnknown() {
	xxx_messageInfo_SFTPAuth.DiscardUnknown(m)
}

var xxx_messageInfo_SFTPAuth proto.InternalMessageInfo

func (m *SFTPAuth) GetSftpAuthType() SFTPAuth_SFTPAuthType {
	if m != nil {
		return m.SftpAuthType
	}
	return SFTPAuth_SFTP_AUTH_TYPE_UNSPECIFIED
}

func (m *SFTPAuth) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *SFTPAuth) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *SFTPAuth) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type NodeConfig struct {
	NodeId               string              `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	Nodetype             NodeConfig_NodeType `protobuf:"varint,2,opt,name=nodetype,proto3,enum=nodeconfig.NodeConfig_NodeType" json:"nodetype,omitempty"`
	IpAddress            string              `protobuf:"bytes,3,opt,name=ip_address,json=ipAddress,proto3" json:"ip_address,omitempty"`
	GrpcPort             string              `protobuf:"bytes,4,opt,name=grpc_port,json=grpcPort,proto3" json:"grpc_port,omitempty"`
	SftpPort             string              `protobuf:"bytes,5,opt,name=sftp_port,json=sftpPort,proto3" json:"sftp_port,omitempty"`
	EnvSec               NodeConfig_EnvSec   `protobuf:"varint,6,opt,name=env_sec,json=envSec,proto3,enum=nodeconfig.NodeConfig_EnvSec" json:"env_sec,omitempty"`
	GrpcAuth             *GRPCAuth           `protobuf:"bytes,7,opt,name=grpc_auth,json=grpcAuth,proto3" json:"grpc_auth,omitempty"`
	SftpAuth             *SFTPAuth           `protobuf:"bytes,8,opt,name=sftp_auth,json=sftpAuth,proto3" json:"sftp_auth,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *NodeConfig) Reset()         { *m = NodeConfig{} }
func (m *NodeConfig) String() string { return proto.CompactTextString(m) }
func (*NodeConfig) ProtoMessage()    {}
func (*NodeConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_98a96787b4b9490b, []int{2}
}

func (m *NodeConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeConfig.Unmarshal(m, b)
}
func (m *NodeConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeConfig.Marshal(b, m, deterministic)
}
func (m *NodeConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeConfig.Merge(m, src)
}
func (m *NodeConfig) XXX_Size() int {
	return xxx_messageInfo_NodeConfig.Size(m)
}
func (m *NodeConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeConfig.DiscardUnknown(m)
}

var xxx_messageInfo_NodeConfig proto.InternalMessageInfo

func (m *NodeConfig) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *NodeConfig) GetNodetype() NodeConfig_NodeType {
	if m != nil {
		return m.Nodetype
	}
	return NodeConfig_NODETYPE_UNSPECIFIED
}

func (m *NodeConfig) GetIpAddress() string {
	if m != nil {
		return m.IpAddress
	}
	return ""
}

func (m *NodeConfig) GetGrpcPort() string {
	if m != nil {
		return m.GrpcPort
	}
	return ""
}

func (m *NodeConfig) GetSftpPort() string {
	if m != nil {
		return m.SftpPort
	}
	return ""
}

func (m *NodeConfig) GetEnvSec() NodeConfig_EnvSec {
	if m != nil {
		return m.EnvSec
	}
	return NodeConfig_ENVSEC_UNSPECIFIED
}

func (m *NodeConfig) GetGrpcAuth() *GRPCAuth {
	if m != nil {
		return m.GrpcAuth
	}
	return nil
}

func (m *NodeConfig) GetSftpAuth() *SFTPAuth {
	if m != nil {
		return m.SftpAuth
	}
	return nil
}

type AddNodeConfigResponse struct {
	NodeId               string                                `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	NodeEntryMessage     string                                `protobuf:"bytes,2,opt,name=node_entry_message,json=nodeEntryMessage,proto3" json:"node_entry_message,omitempty"`
	NodeEntryStatus      AddNodeConfigResponse_NodeEntryStatus `protobuf:"varint,3,opt,name=node_entry_status,json=nodeEntryStatus,proto3,enum=nodeconfig.AddNodeConfigResponse_NodeEntryStatus" json:"node_entry_status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                              `json:"-"`
	XXX_unrecognized     []byte                                `json:"-"`
	XXX_sizecache        int32                                 `json:"-"`
}

func (m *AddNodeConfigResponse) Reset()         { *m = AddNodeConfigResponse{} }
func (m *AddNodeConfigResponse) String() string { return proto.CompactTextString(m) }
func (*AddNodeConfigResponse) ProtoMessage()    {}
func (*AddNodeConfigResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_98a96787b4b9490b, []int{3}
}

func (m *AddNodeConfigResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddNodeConfigResponse.Unmarshal(m, b)
}
func (m *AddNodeConfigResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddNodeConfigResponse.Marshal(b, m, deterministic)
}
func (m *AddNodeConfigResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddNodeConfigResponse.Merge(m, src)
}
func (m *AddNodeConfigResponse) XXX_Size() int {
	return xxx_messageInfo_AddNodeConfigResponse.Size(m)
}
func (m *AddNodeConfigResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AddNodeConfigResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AddNodeConfigResponse proto.InternalMessageInfo

func (m *AddNodeConfigResponse) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *AddNodeConfigResponse) GetNodeEntryMessage() string {
	if m != nil {
		return m.NodeEntryMessage
	}
	return ""
}

func (m *AddNodeConfigResponse) GetNodeEntryStatus() AddNodeConfigResponse_NodeEntryStatus {
	if m != nil {
		return m.NodeEntryStatus
	}
	return AddNodeConfigResponse_STATUS_UNSPECIFIED
}

type NodeID struct {
	NodeId               string   `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeID) Reset()         { *m = NodeID{} }
func (m *NodeID) String() string { return proto.CompactTextString(m) }
func (*NodeID) ProtoMessage()    {}
func (*NodeID) Descriptor() ([]byte, []int) {
	return fileDescriptor_98a96787b4b9490b, []int{4}
}

func (m *NodeID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeID.Unmarshal(m, b)
}
func (m *NodeID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeID.Marshal(b, m, deterministic)
}
func (m *NodeID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeID.Merge(m, src)
}
func (m *NodeID) XXX_Size() int {
	return xxx_messageInfo_NodeID.Size(m)
}
func (m *NodeID) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeID.DiscardUnknown(m)
}

var xxx_messageInfo_NodeID proto.InternalMessageInfo

func (m *NodeID) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

type ListNodeConfigRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListNodeConfigRequest) Reset()         { *m = ListNodeConfigRequest{} }
func (m *ListNodeConfigRequest) String() string { return proto.CompactTextString(m) }
func (*ListNodeConfigRequest) ProtoMessage()    {}
func (*ListNodeConfigRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_98a96787b4b9490b, []int{5}
}

func (m *ListNodeConfigRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListNodeConfigRequest.Unmarshal(m, b)
}
func (m *ListNodeConfigRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListNodeConfigRequest.Marshal(b, m, deterministic)
}
func (m *ListNodeConfigRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListNodeConfigRequest.Merge(m, src)
}
func (m *ListNodeConfigRequest) XXX_Size() int {
	return xxx_messageInfo_ListNodeConfigRequest.Size(m)
}
func (m *ListNodeConfigRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListNodeConfigRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListNodeConfigRequest proto.InternalMessageInfo

type AllNodeConfig struct {
	AllNodeConfig        []*NodeConfig `protobuf:"bytes,1,rep,name=all_node_config,json=allNodeConfig,proto3" json:"all_node_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *AllNodeConfig) Reset()         { *m = AllNodeConfig{} }
func (m *AllNodeConfig) String() string { return proto.CompactTextString(m) }
func (*AllNodeConfig) ProtoMessage()    {}
func (*AllNodeConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_98a96787b4b9490b, []int{6}
}

func (m *AllNodeConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AllNodeConfig.Unmarshal(m, b)
}
func (m *AllNodeConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AllNodeConfig.Marshal(b, m, deterministic)
}
func (m *AllNodeConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AllNodeConfig.Merge(m, src)
}
func (m *AllNodeConfig) XXX_Size() int {
	return xxx_messageInfo_AllNodeConfig.Size(m)
}
func (m *AllNodeConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_AllNodeConfig.DiscardUnknown(m)
}

var xxx_messageInfo_AllNodeConfig proto.InternalMessageInfo

func (m *AllNodeConfig) GetAllNodeConfig() []*NodeConfig {
	if m != nil {
		return m.AllNodeConfig
	}
	return nil
}

func init() {
	proto.RegisterEnum("nodeconfig.SFTPAuth_SFTPAuthType", SFTPAuth_SFTPAuthType_name, SFTPAuth_SFTPAuthType_value)
	proto.RegisterEnum("nodeconfig.NodeConfig_NodeType", NodeConfig_NodeType_name, NodeConfig_NodeType_value)
	proto.RegisterEnum("nodeconfig.NodeConfig_EnvSec", NodeConfig_EnvSec_name, NodeConfig_EnvSec_value)
	proto.RegisterEnum("nodeconfig.AddNodeConfigResponse_NodeEntryStatus", AddNodeConfigResponse_NodeEntryStatus_name, AddNodeConfigResponse_NodeEntryStatus_value)
	proto.RegisterType((*GRPCAuth)(nil), "nodeconfig.GRPCAuth")
	proto.RegisterType((*SFTPAuth)(nil), "nodeconfig.SFTPAuth")
	proto.RegisterType((*NodeConfig)(nil), "nodeconfig.NodeConfig")
	proto.RegisterType((*AddNodeConfigResponse)(nil), "nodeconfig.AddNodeConfigResponse")
	proto.RegisterType((*NodeID)(nil), "nodeconfig.NodeID")
	proto.RegisterType((*ListNodeConfigRequest)(nil), "nodeconfig.ListNodeConfigRequest")
	proto.RegisterType((*AllNodeConfig)(nil), "nodeconfig.AllNodeConfig")
}

func init() { proto.RegisterFile("node_config_events.proto", fileDescriptor_98a96787b4b9490b) }

var fileDescriptor_98a96787b4b9490b = []byte{
	// 665 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x54, 0xdf, 0x4f, 0xdb, 0x3e,
	0x10, 0xff, 0xa6, 0xe5, 0xdb, 0x86, 0xa3, 0x94, 0x60, 0xf1, 0x23, 0x63, 0x62, 0x83, 0x48, 0x93,
	0x78, 0x98, 0x90, 0xe8, 0xa4, 0xbd, 0x4c, 0x9a, 0xe6, 0xa5, 0x06, 0x2a, 0x4a, 0x1a, 0xd9, 0x29,
	0x1b, 0x93, 0x50, 0x94, 0x35, 0x06, 0xaa, 0x95, 0x24, 0x8b, 0x5d, 0xa6, 0xfe, 0x13, 0xfb, 0x67,
	0xf7, 0xba, 0x87, 0xc9, 0x4e, 0x0a, 0x19, 0x83, 0xa7, 0xdc, 0xdd, 0xe7, 0xee, 0xfc, 0xf1, 0xe7,
	0x2e, 0x06, 0x3b, 0x49, 0x63, 0x1e, 0x8e, 0xd2, 0xe4, 0x72, 0x7c, 0x15, 0xf2, 0x5b, 0x9e, 0x48,
	0xb1, 0x9f, 0xe5, 0xa9, 0x4c, 0x11, 0x28, 0xa4, 0x00, 0x9c, 0x57, 0x60, 0x1e, 0x51, 0xdf, 0xc5,
	0x53, 0x79, 0x8d, 0x9e, 0x81, 0x29, 0x27, 0x22, 0x1c, 0xf1, 0x5c, 0xda, 0xc6, 0x8e, 0xb1, 0xb7,
	0x48, 0x9b, 0x72, 0x22, 0x5c, 0x9e, 0x4b, 0xe7, 0xb7, 0x01, 0x26, 0x3b, 0x0c, 0x7c, 0x9d, 0x77,
	0x04, 0x6d, 0x71, 0x29, 0xb3, 0x30, 0x9a, 0xca, 0xeb, 0x50, 0xce, 0x32, 0xae, 0xb3, 0xdb, 0x9d,
	0xdd, 0xfd, 0xfb, 0xc6, 0xfb, 0xf3, 0xec, 0x3b, 0x23, 0x98, 0x65, 0x9c, 0xb6, 0x54, 0xe1, 0xdc,
	0x43, 0x5b, 0x60, 0x4e, 0x05, 0xcf, 0x93, 0xe8, 0x86, 0xdb, 0x35, 0x7d, 0xe0, 0x9d, 0xaf, 0xb0,
	0x2c, 0x12, 0xe2, 0x47, 0x9a, 0xc7, 0x76, 0xbd, 0xc0, 0xe6, 0x3e, 0xb2, 0xa0, 0xfe, 0x8d, 0xcf,
	0xec, 0x05, 0x1d, 0x56, 0xa6, 0x73, 0x01, 0xad, 0xea, 0x39, 0xe8, 0x05, 0x6c, 0x29, 0x3f, 0xc4,
	0xc3, 0xe0, 0x38, 0x0c, 0xce, 0x7d, 0x12, 0x0e, 0x3d, 0xe6, 0x13, 0xb7, 0x77, 0xd8, 0x23, 0x5d,
	0xeb, 0x3f, 0xd4, 0x02, 0xd3, 0xc7, 0x8c, 0x7d, 0x1a, 0xd0, 0xae, 0x65, 0x28, 0xef, 0x84, 0x9c,
	0x7f, 0xc4, 0x8c, 0x74, 0xad, 0x1a, 0x5a, 0x85, 0xe5, 0x13, 0x72, 0xae, 0x60, 0xff, 0x98, 0x62,
	0x46, 0xac, 0xba, 0xf3, 0xab, 0x0e, 0xe0, 0xa5, 0x31, 0x77, 0xf5, 0xdd, 0xd0, 0x26, 0x34, 0xb5,
	0xb8, 0xe3, 0xb8, 0xd4, 0xa9, 0xa1, 0xdc, 0x5e, 0x8c, 0xde, 0x81, 0xa9, 0x2c, 0xad, 0x49, 0x4d,
	0x6b, 0xf2, 0xb2, 0xaa, 0xc9, 0x7d, 0x0b, 0x6d, 0x6a, 0x45, 0xee, 0x0a, 0xd0, 0x36, 0xc0, 0x38,
	0x0b, 0xa3, 0x38, 0xce, 0xb9, 0x10, 0xe5, 0x9d, 0x17, 0xc7, 0x19, 0x2e, 0x02, 0xe8, 0x39, 0x2c,
	0x5e, 0xe5, 0xd9, 0x28, 0xcc, 0xd2, 0x5c, 0x96, 0x57, 0x37, 0x55, 0xc0, 0x4f, 0x73, 0xa9, 0x40,
	0x3d, 0x12, 0x0d, 0xfe, 0x5f, 0x80, 0x2a, 0xa0, 0xc1, 0xb7, 0xd0, 0xe4, 0xc9, 0x6d, 0x28, 0xf8,
	0xc8, 0x6e, 0x68, 0x52, 0xdb, 0x4f, 0x90, 0x22, 0xc9, 0x2d, 0xe3, 0x23, 0xda, 0xe0, 0xfa, 0x8b,
	0x0e, 0xca, 0x13, 0xd5, 0x9c, 0xed, 0xe6, 0x8e, 0xb1, 0xb7, 0xd4, 0x59, 0xab, 0x56, 0xce, 0x17,
	0xa7, 0xe0, 0xa1, 0x57, 0xe3, 0xa0, 0xe4, 0xa1, 0x4b, 0xcc, 0x7f, 0x4b, 0xe6, 0x43, 0x2a, 0xd8,
	0x29, 0xcb, 0xf9, 0x00, 0xe6, 0x5c, 0x0c, 0x64, 0xc3, 0x9a, 0x37, 0xe8, 0x92, 0xc7, 0x07, 0xd6,
	0xf3, 0x02, 0x42, 0x3d, 0xdc, 0x2f, 0x06, 0x46, 0x3e, 0x97, 0x5e, 0xcd, 0x39, 0x83, 0x46, 0xc1,
	0x1c, 0x6d, 0x00, 0x22, 0xde, 0x19, 0x23, 0xee, 0x83, 0xea, 0x26, 0xd4, 0xbb, 0xa7, 0x5f, 0x2c,
	0x03, 0x2d, 0x41, 0x13, 0xfb, 0x7e, 0xd0, 0x23, 0xd4, 0xaa, 0xa9, 0x2e, 0x5d, 0x1c, 0x60, 0xed,
	0xd5, 0x91, 0x05, 0x2d, 0x77, 0xe0, 0x05, 0x74, 0xd0, 0xf7, 0xfb, 0xd8, 0x23, 0xd6, 0x82, 0xf3,
	0xb3, 0x06, 0xeb, 0x38, 0x8e, 0xef, 0x05, 0xa2, 0x5c, 0x64, 0x69, 0x22, 0xf8, 0xd3, 0x0b, 0xf0,
	0x1a, 0x90, 0x06, 0x78, 0x22, 0xf3, 0x59, 0x78, 0xc3, 0x85, 0x88, 0xae, 0xe6, 0xbb, 0x6d, 0x29,
	0x84, 0x28, 0xe0, 0xb4, 0x88, 0xa3, 0x0b, 0x58, 0xad, 0x64, 0x0b, 0x19, 0xc9, 0x69, 0x31, 0xf8,
	0x76, 0xe7, 0xa0, 0xaa, 0xda, 0xa3, 0x24, 0xf4, 0xe0, 0x74, 0x3b, 0xa6, 0x0b, 0xe9, 0x4a, 0xf2,
	0x77, 0xc0, 0x19, 0xc0, 0xca, 0x83, 0x1c, 0x25, 0x10, 0x0b, 0x70, 0x30, 0x64, 0x0f, 0x04, 0x42,
	0xd0, 0x2e, 0xe3, 0x6c, 0xe8, 0xba, 0x84, 0x31, 0xcb, 0x50, 0x82, 0x94, 0x31, 0x42, 0xe9, 0x80,
	0x5a, 0x35, 0x67, 0x17, 0x1a, 0xaa, 0x61, 0xaf, 0xfb, 0xa4, 0x00, 0xce, 0x26, 0xac, 0xf7, 0xc7,
	0x42, 0x56, 0xe9, 0x7e, 0x9f, 0x72, 0x21, 0x9d, 0x01, 0x2c, 0xe3, 0xc9, 0xa4, 0xf2, 0x13, 0xbd,
	0x87, 0x95, 0x68, 0x32, 0x09, 0x2b, 0xaf, 0x94, 0x6d, 0xec, 0xd4, 0xf7, 0x96, 0x3a, 0x1b, 0x8f,
	0x6f, 0x27, 0x5d, 0x8e, 0xaa, 0xf5, 0x5f, 0x1b, 0xfa, 0x31, 0x7b, 0xf3, 0x27, 0x00, 0x00, 0xff,
	0xff, 0xec, 0xc2, 0x12, 0x6d, 0xe8, 0x04, 0x00, 0x00,
}
