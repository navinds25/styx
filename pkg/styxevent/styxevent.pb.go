// Code generated by protoc-gen-go. DO NOT EDIT.
// source: styxevent.proto

package styxevent

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type SftpTransferConfig_TransferType int32

const (
	SftpTransferConfig_Pull SftpTransferConfig_TransferType = 0
	SftpTransferConfig_Push SftpTransferConfig_TransferType = 1
)

var SftpTransferConfig_TransferType_name = map[int32]string{
	0: "Pull",
	1: "Push",
}

var SftpTransferConfig_TransferType_value = map[string]int32{
	"Pull": 0,
	"Push": 1,
}

func (x SftpTransferConfig_TransferType) String() string {
	return proto.EnumName(SftpTransferConfig_TransferType_name, int32(x))
}

func (SftpTransferConfig_TransferType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1257f2c7de005776, []int{6, 0}
}

type Noparams struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Noparams) Reset()         { *m = Noparams{} }
func (m *Noparams) String() string { return proto.CompactTextString(m) }
func (*Noparams) ProtoMessage()    {}
func (*Noparams) Descriptor() ([]byte, []int) {
	return fileDescriptor_1257f2c7de005776, []int{0}
}

func (m *Noparams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Noparams.Unmarshal(m, b)
}
func (m *Noparams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Noparams.Marshal(b, m, deterministic)
}
func (m *Noparams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Noparams.Merge(m, src)
}
func (m *Noparams) XXX_Size() int {
	return xxx_messageInfo_Noparams.Size(m)
}
func (m *Noparams) XXX_DiscardUnknown() {
	xxx_messageInfo_Noparams.DiscardUnknown(m)
}

var xxx_messageInfo_Noparams proto.InternalMessageInfo

type Ack struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Error                string   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ack) Reset()         { *m = Ack{} }
func (m *Ack) String() string { return proto.CompactTextString(m) }
func (*Ack) ProtoMessage()    {}
func (*Ack) Descriptor() ([]byte, []int) {
	return fileDescriptor_1257f2c7de005776, []int{1}
}

func (m *Ack) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ack.Unmarshal(m, b)
}
func (m *Ack) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ack.Marshal(b, m, deterministic)
}
func (m *Ack) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ack.Merge(m, src)
}
func (m *Ack) XXX_Size() int {
	return xxx_messageInfo_Ack.Size(m)
}
func (m *Ack) XXX_DiscardUnknown() {
	xxx_messageInfo_Ack.DiscardUnknown(m)
}

var xxx_messageInfo_Ack proto.InternalMessageInfo

func (m *Ack) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Ack) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type SearchFileInfo struct {
	Jid                  string   `protobuf:"bytes,1,opt,name=jid,proto3" json:"jid,omitempty"`
	Type                 string   `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Filename             string   `protobuf:"bytes,3,opt,name=filename,proto3" json:"filename,omitempty"`
	Directory            string   `protobuf:"bytes,4,opt,name=directory,proto3" json:"directory,omitempty"`
	Server               string   `protobuf:"bytes,5,opt,name=server,proto3" json:"server,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchFileInfo) Reset()         { *m = SearchFileInfo{} }
func (m *SearchFileInfo) String() string { return proto.CompactTextString(m) }
func (*SearchFileInfo) ProtoMessage()    {}
func (*SearchFileInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_1257f2c7de005776, []int{2}
}

func (m *SearchFileInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchFileInfo.Unmarshal(m, b)
}
func (m *SearchFileInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchFileInfo.Marshal(b, m, deterministic)
}
func (m *SearchFileInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchFileInfo.Merge(m, src)
}
func (m *SearchFileInfo) XXX_Size() int {
	return xxx_messageInfo_SearchFileInfo.Size(m)
}
func (m *SearchFileInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchFileInfo.DiscardUnknown(m)
}

var xxx_messageInfo_SearchFileInfo proto.InternalMessageInfo

func (m *SearchFileInfo) GetJid() string {
	if m != nil {
		return m.Jid
	}
	return ""
}

func (m *SearchFileInfo) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *SearchFileInfo) GetFilename() string {
	if m != nil {
		return m.Filename
	}
	return ""
}

func (m *SearchFileInfo) GetDirectory() string {
	if m != nil {
		return m.Directory
	}
	return ""
}

func (m *SearchFileInfo) GetServer() string {
	if m != nil {
		return m.Server
	}
	return ""
}

type FoundFile struct {
	Jid                  string   `protobuf:"bytes,1,opt,name=jid,proto3" json:"jid,omitempty"`
	Type                 string   `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Match                string   `protobuf:"bytes,3,opt,name=match,proto3" json:"match,omitempty"`
	Regmatch             string   `protobuf:"bytes,4,opt,name=regmatch,proto3" json:"regmatch,omitempty"`
	Error                string   `protobuf:"bytes,5,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FoundFile) Reset()         { *m = FoundFile{} }
func (m *FoundFile) String() string { return proto.CompactTextString(m) }
func (*FoundFile) ProtoMessage()    {}
func (*FoundFile) Descriptor() ([]byte, []int) {
	return fileDescriptor_1257f2c7de005776, []int{3}
}

func (m *FoundFile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FoundFile.Unmarshal(m, b)
}
func (m *FoundFile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FoundFile.Marshal(b, m, deterministic)
}
func (m *FoundFile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FoundFile.Merge(m, src)
}
func (m *FoundFile) XXX_Size() int {
	return xxx_messageInfo_FoundFile.Size(m)
}
func (m *FoundFile) XXX_DiscardUnknown() {
	xxx_messageInfo_FoundFile.DiscardUnknown(m)
}

var xxx_messageInfo_FoundFile proto.InternalMessageInfo

func (m *FoundFile) GetJid() string {
	if m != nil {
		return m.Jid
	}
	return ""
}

func (m *FoundFile) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *FoundFile) GetMatch() string {
	if m != nil {
		return m.Match
	}
	return ""
}

func (m *FoundFile) GetRegmatch() string {
	if m != nil {
		return m.Regmatch
	}
	return ""
}

func (m *FoundFile) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type GetRemoteFile struct {
	Jid                  string   `protobuf:"bytes,1,opt,name=jid,proto3" json:"jid,omitempty"`
	Type                 string   `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Sourcefile           string   `protobuf:"bytes,3,opt,name=sourcefile,proto3" json:"sourcefile,omitempty"`
	Destinationfile      string   `protobuf:"bytes,4,opt,name=destinationfile,proto3" json:"destinationfile,omitempty"`
	Sourceserver         string   `protobuf:"bytes,5,opt,name=sourceserver,proto3" json:"sourceserver,omitempty"`
	Authmethod           string   `protobuf:"bytes,6,opt,name=authmethod,proto3" json:"authmethod,omitempty"`
	Username             string   `protobuf:"bytes,7,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,8,opt,name=password,proto3" json:"password,omitempty"`
	Key                  string   `protobuf:"bytes,9,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRemoteFile) Reset()         { *m = GetRemoteFile{} }
func (m *GetRemoteFile) String() string { return proto.CompactTextString(m) }
func (*GetRemoteFile) ProtoMessage()    {}
func (*GetRemoteFile) Descriptor() ([]byte, []int) {
	return fileDescriptor_1257f2c7de005776, []int{4}
}

func (m *GetRemoteFile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRemoteFile.Unmarshal(m, b)
}
func (m *GetRemoteFile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRemoteFile.Marshal(b, m, deterministic)
}
func (m *GetRemoteFile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRemoteFile.Merge(m, src)
}
func (m *GetRemoteFile) XXX_Size() int {
	return xxx_messageInfo_GetRemoteFile.Size(m)
}
func (m *GetRemoteFile) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRemoteFile.DiscardUnknown(m)
}

var xxx_messageInfo_GetRemoteFile proto.InternalMessageInfo

func (m *GetRemoteFile) GetJid() string {
	if m != nil {
		return m.Jid
	}
	return ""
}

func (m *GetRemoteFile) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *GetRemoteFile) GetSourcefile() string {
	if m != nil {
		return m.Sourcefile
	}
	return ""
}

func (m *GetRemoteFile) GetDestinationfile() string {
	if m != nil {
		return m.Destinationfile
	}
	return ""
}

func (m *GetRemoteFile) GetSourceserver() string {
	if m != nil {
		return m.Sourceserver
	}
	return ""
}

func (m *GetRemoteFile) GetAuthmethod() string {
	if m != nil {
		return m.Authmethod
	}
	return ""
}

func (m *GetRemoteFile) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *GetRemoteFile) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *GetRemoteFile) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type RequestFileTransfer struct {
	Jid                  string   `protobuf:"bytes,1,opt,name=jid,proto3" json:"jid,omitempty"`
	Type                 string   `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Sourcefile           string   `protobuf:"bytes,3,opt,name=sourcefile,proto3" json:"sourcefile,omitempty"`
	Destinationfile      string   `protobuf:"bytes,4,opt,name=destinationfile,proto3" json:"destinationfile,omitempty"`
	Sourceserver         string   `protobuf:"bytes,5,opt,name=sourceserver,proto3" json:"sourceserver,omitempty"`
	Destinationserver    string   `protobuf:"bytes,6,opt,name=destinationserver,proto3" json:"destinationserver,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestFileTransfer) Reset()         { *m = RequestFileTransfer{} }
func (m *RequestFileTransfer) String() string { return proto.CompactTextString(m) }
func (*RequestFileTransfer) ProtoMessage()    {}
func (*RequestFileTransfer) Descriptor() ([]byte, []int) {
	return fileDescriptor_1257f2c7de005776, []int{5}
}

func (m *RequestFileTransfer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestFileTransfer.Unmarshal(m, b)
}
func (m *RequestFileTransfer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestFileTransfer.Marshal(b, m, deterministic)
}
func (m *RequestFileTransfer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestFileTransfer.Merge(m, src)
}
func (m *RequestFileTransfer) XXX_Size() int {
	return xxx_messageInfo_RequestFileTransfer.Size(m)
}
func (m *RequestFileTransfer) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestFileTransfer.DiscardUnknown(m)
}

var xxx_messageInfo_RequestFileTransfer proto.InternalMessageInfo

func (m *RequestFileTransfer) GetJid() string {
	if m != nil {
		return m.Jid
	}
	return ""
}

func (m *RequestFileTransfer) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *RequestFileTransfer) GetSourcefile() string {
	if m != nil {
		return m.Sourcefile
	}
	return ""
}

func (m *RequestFileTransfer) GetDestinationfile() string {
	if m != nil {
		return m.Destinationfile
	}
	return ""
}

func (m *RequestFileTransfer) GetSourceserver() string {
	if m != nil {
		return m.Sourceserver
	}
	return ""
}

func (m *RequestFileTransfer) GetDestinationserver() string {
	if m != nil {
		return m.Destinationserver
	}
	return ""
}

type SftpTransferConfig struct {
	Transferid           string                          `protobuf:"bytes,1,opt,name=transferid,proto3" json:"transferid,omitempty"`
	Description          string                          `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Type                 SftpTransferConfig_TransferType `protobuf:"varint,3,opt,name=type,proto3,enum=styxevent.SftpTransferConfig_TransferType" json:"type,omitempty"`
	Localfile            string                          `protobuf:"bytes,4,opt,name=localfile,proto3" json:"localfile,omitempty"`
	Localpath            string                          `protobuf:"bytes,5,opt,name=localpath,proto3" json:"localpath,omitempty"`
	Remotefile           string                          `protobuf:"bytes,6,opt,name=remotefile,proto3" json:"remotefile,omitempty"`
	Remotepath           string                          `protobuf:"bytes,7,opt,name=remotepath,proto3" json:"remotepath,omitempty"`
	Remotehost           string                          `protobuf:"bytes,8,opt,name=remotehost,proto3" json:"remotehost,omitempty"`
	Remoteport           int32                           `protobuf:"varint,9,opt,name=remoteport,proto3" json:"remoteport,omitempty"`
	Remoteuser           string                          `protobuf:"bytes,10,opt,name=remoteuser,proto3" json:"remoteuser,omitempty"`
	Remotepassword       string                          `protobuf:"bytes,11,opt,name=remotepassword,proto3" json:"remotepassword,omitempty"`
	Remotekey            string                          `protobuf:"bytes,12,opt,name=remotekey,proto3" json:"remotekey,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *SftpTransferConfig) Reset()         { *m = SftpTransferConfig{} }
func (m *SftpTransferConfig) String() string { return proto.CompactTextString(m) }
func (*SftpTransferConfig) ProtoMessage()    {}
func (*SftpTransferConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_1257f2c7de005776, []int{6}
}

func (m *SftpTransferConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SftpTransferConfig.Unmarshal(m, b)
}
func (m *SftpTransferConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SftpTransferConfig.Marshal(b, m, deterministic)
}
func (m *SftpTransferConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SftpTransferConfig.Merge(m, src)
}
func (m *SftpTransferConfig) XXX_Size() int {
	return xxx_messageInfo_SftpTransferConfig.Size(m)
}
func (m *SftpTransferConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_SftpTransferConfig.DiscardUnknown(m)
}

var xxx_messageInfo_SftpTransferConfig proto.InternalMessageInfo

func (m *SftpTransferConfig) GetTransferid() string {
	if m != nil {
		return m.Transferid
	}
	return ""
}

func (m *SftpTransferConfig) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *SftpTransferConfig) GetType() SftpTransferConfig_TransferType {
	if m != nil {
		return m.Type
	}
	return SftpTransferConfig_Pull
}

func (m *SftpTransferConfig) GetLocalfile() string {
	if m != nil {
		return m.Localfile
	}
	return ""
}

func (m *SftpTransferConfig) GetLocalpath() string {
	if m != nil {
		return m.Localpath
	}
	return ""
}

func (m *SftpTransferConfig) GetRemotefile() string {
	if m != nil {
		return m.Remotefile
	}
	return ""
}

func (m *SftpTransferConfig) GetRemotepath() string {
	if m != nil {
		return m.Remotepath
	}
	return ""
}

func (m *SftpTransferConfig) GetRemotehost() string {
	if m != nil {
		return m.Remotehost
	}
	return ""
}

func (m *SftpTransferConfig) GetRemoteport() int32 {
	if m != nil {
		return m.Remoteport
	}
	return 0
}

func (m *SftpTransferConfig) GetRemoteuser() string {
	if m != nil {
		return m.Remoteuser
	}
	return ""
}

func (m *SftpTransferConfig) GetRemotepassword() string {
	if m != nil {
		return m.Remotepassword
	}
	return ""
}

func (m *SftpTransferConfig) GetRemotekey() string {
	if m != nil {
		return m.Remotekey
	}
	return ""
}

func init() {
	proto.RegisterEnum("styxevent.SftpTransferConfig_TransferType", SftpTransferConfig_TransferType_name, SftpTransferConfig_TransferType_value)
	proto.RegisterType((*Noparams)(nil), "styxevent.Noparams")
	proto.RegisterType((*Ack)(nil), "styxevent.Ack")
	proto.RegisterType((*SearchFileInfo)(nil), "styxevent.SearchFileInfo")
	proto.RegisterType((*FoundFile)(nil), "styxevent.FoundFile")
	proto.RegisterType((*GetRemoteFile)(nil), "styxevent.GetRemoteFile")
	proto.RegisterType((*RequestFileTransfer)(nil), "styxevent.RequestFileTransfer")
	proto.RegisterType((*SftpTransferConfig)(nil), "styxevent.SftpTransferConfig")
}

func init() { proto.RegisterFile("styxevent.proto", fileDescriptor_1257f2c7de005776) }

var fileDescriptor_1257f2c7de005776 = []byte{
	// 653 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x55, 0xdd, 0x4e, 0xd4, 0x40,
	0x14, 0xb6, 0xec, 0x2e, 0x6c, 0x0f, 0xb8, 0xac, 0x03, 0x31, 0x75, 0x83, 0x84, 0xf4, 0xc2, 0x10,
	0x62, 0x08, 0x41, 0xbd, 0x31, 0xd1, 0x04, 0x49, 0x20, 0x26, 0xc6, 0x90, 0x5d, 0x5e, 0x60, 0x6c,
	0xcf, 0xd2, 0xca, 0xb6, 0x53, 0x66, 0x66, 0x91, 0xbe, 0x81, 0x37, 0x3e, 0x8d, 0x0f, 0xe3, 0x8d,
	0x0f, 0x63, 0xe6, 0xa7, 0xed, 0x74, 0xd7, 0x10, 0xb9, 0xf3, 0x6e, 0xce, 0xf7, 0xcd, 0x37, 0xe7,
	0x7c, 0xa7, 0x67, 0xa6, 0xb0, 0x29, 0x64, 0x79, 0x87, 0xb7, 0x98, 0xcb, 0xc3, 0x82, 0x33, 0xc9,
	0x88, 0x5f, 0x03, 0x21, 0x40, 0xff, 0x33, 0x2b, 0x28, 0xa7, 0x99, 0x08, 0xdf, 0x40, 0xe7, 0x24,
	0xba, 0x26, 0x01, 0xac, 0x65, 0x28, 0x04, 0xbd, 0xc2, 0xc0, 0xdb, 0xf3, 0xf6, 0xfd, 0x71, 0x15,
	0x92, 0x6d, 0xe8, 0x21, 0xe7, 0x8c, 0x07, 0x2b, 0x1a, 0x37, 0x41, 0xf8, 0xdd, 0x83, 0xc1, 0x04,
	0x29, 0x8f, 0x92, 0xb3, 0x74, 0x86, 0x1f, 0xf3, 0x29, 0x23, 0x43, 0xe8, 0x7c, 0x4d, 0x63, 0x2b,
	0x57, 0x4b, 0x42, 0xa0, 0x2b, 0xcb, 0x02, 0xad, 0x52, 0xaf, 0xc9, 0x08, 0xfa, 0xd3, 0x74, 0x86,
	0x39, 0xcd, 0x30, 0xe8, 0x68, 0xbc, 0x8e, 0xc9, 0x0e, 0xf8, 0x71, 0xca, 0x31, 0x92, 0x8c, 0x97,
	0x41, 0x57, 0x93, 0x0d, 0x40, 0x9e, 0xc2, 0xaa, 0x40, 0x7e, 0x8b, 0x3c, 0xe8, 0x69, 0xca, 0x46,
	0x61, 0x09, 0xfe, 0x19, 0x9b, 0xe7, 0xb1, 0x2a, 0xe4, 0x1f, 0x8b, 0xd8, 0x86, 0x5e, 0x46, 0x65,
	0x94, 0xd8, 0x0a, 0x4c, 0xa0, 0x4a, 0xe3, 0x78, 0x65, 0x08, 0x93, 0xbd, 0x8e, 0x9b, 0x2e, 0xf4,
	0xdc, 0x2e, 0xfc, 0x58, 0x81, 0xc7, 0xe7, 0x28, 0xc7, 0x98, 0x31, 0x89, 0x0f, 0xc8, 0xbf, 0x0b,
	0x20, 0xd8, 0x9c, 0x47, 0xa8, 0xac, 0xdb, 0x22, 0x1c, 0x84, 0xec, 0xc3, 0x66, 0x8c, 0x42, 0xa6,
	0x39, 0x95, 0x29, 0xcb, 0xf5, 0x26, 0x53, 0xd0, 0x22, 0x4c, 0x42, 0xd8, 0x30, 0xba, 0x56, 0x6b,
	0x5a, 0x98, 0xca, 0x46, 0xe7, 0x32, 0xc9, 0x50, 0x26, 0x2c, 0x0e, 0x56, 0x4d, 0xb6, 0x06, 0x51,
	0xbe, 0xe7, 0x02, 0xb9, 0xfe, 0x24, 0x6b, 0xc6, 0x77, 0x15, 0x2b, 0xae, 0xa0, 0x42, 0x7c, 0x63,
	0x3c, 0x0e, 0xfa, 0x86, 0xab, 0x62, 0xe5, 0xf5, 0x1a, 0xcb, 0xc0, 0x37, 0x5e, 0xaf, 0xb1, 0x0c,
	0x7f, 0x79, 0xb0, 0x35, 0xc6, 0x9b, 0x39, 0x0a, 0xa9, 0xba, 0x71, 0xc9, 0x69, 0x2e, 0xa6, 0xc8,
	0xff, 0xcb, 0xae, 0xbc, 0x84, 0x27, 0x8e, 0xcc, 0x6e, 0x34, 0xcd, 0x59, 0x26, 0xc2, 0xdf, 0x1d,
	0x20, 0x93, 0xa9, 0x2c, 0x2a, 0x4b, 0xa7, 0x2a, 0xd3, 0x95, 0x2a, 0x59, 0x5a, 0xa4, 0xf6, 0xe7,
	0x20, 0x64, 0x0f, 0xd6, 0x63, 0x14, 0x11, 0x4f, 0x0b, 0x75, 0x96, 0x75, 0xeb, 0x42, 0xe4, 0xbd,
	0x6d, 0x84, 0xb2, 0x3b, 0x38, 0x3e, 0x38, 0x6c, 0xae, 0xed, 0x72, 0xba, 0xc3, 0x2a, 0xbc, 0x2c,
	0x0b, 0xb4, 0x4d, 0xdb, 0x01, 0x7f, 0xc6, 0x22, 0x3a, 0x73, 0xda, 0xd1, 0x00, 0x35, 0x5b, 0x50,
	0x99, 0xd8, 0x2e, 0x34, 0x80, 0xaa, 0x9e, 0xeb, 0xd1, 0xd5, 0x62, 0x3b, 0x18, 0x0d, 0xd2, 0xf0,
	0x5a, 0xbe, 0xe6, 0xf2, 0x6d, 0x7d, 0xc2, 0x84, 0xb4, 0xe3, 0xe1, 0x20, 0x8e, 0x9e, 0x71, 0xa9,
	0xe7, 0xa4, 0x37, 0x76, 0x90, 0x86, 0x57, 0xe3, 0x16, 0x80, 0xab, 0x57, 0x08, 0x79, 0x01, 0x83,
	0x2a, 0x9b, 0x1d, 0xc1, 0x75, 0xbd, 0x67, 0x01, 0x55, 0x2e, 0x0d, 0xa2, 0xc6, 0x71, 0xc3, 0xb8,
	0xac, 0x81, 0x30, 0x84, 0x0d, 0xb7, 0x6f, 0xa4, 0x0f, 0xdd, 0x8b, 0xf9, 0x6c, 0x36, 0x7c, 0x64,
	0x56, 0x22, 0x19, 0x7a, 0xc7, 0x3f, 0x3b, 0xd0, 0x9d, 0xc8, 0xf2, 0x8e, 0xbc, 0x03, 0x50, 0x93,
	0x6b, 0x9e, 0x36, 0xf2, 0xcc, 0xfd, 0x1c, 0xad, 0xd7, 0x6e, 0xb4, 0xed, 0x50, 0xcd, 0xf3, 0x73,
	0x0e, 0xc3, 0x46, 0x3e, 0x91, 0x1c, 0x69, 0xf6, 0xe0, 0x43, 0xf6, 0xbd, 0x23, 0x8f, 0x9c, 0xc0,
	0xe6, 0x18, 0x6f, 0x5a, 0x97, 0x68, 0xd7, 0xd9, 0xfc, 0x97, 0x4b, 0x36, 0x1a, 0x38, 0xbc, 0x7a,
	0xd2, 0x5f, 0x43, 0x5f, 0xf9, 0xd4, 0x75, 0x05, 0x0e, 0xd7, 0x7a, 0xb0, 0x96, 0x54, 0x07, 0xd0,
	0x3f, 0x1f, 0x5f, 0x9c, 0x5e, 0xa2, 0x90, 0x64, 0x81, 0x5b, 0xda, 0xfb, 0x16, 0xfc, 0x93, 0x38,
	0xb6, 0x57, 0xe1, 0xf9, 0xbd, 0xa3, 0xbb, 0xa4, 0xfd, 0x00, 0xf0, 0x29, 0x15, 0xd2, 0x8a, 0xb7,
	0x1c, 0xb6, 0xfa, 0x35, 0x8d, 0xee, 0x3f, 0xf1, 0xc8, 0xfb, 0xb2, 0xaa, 0xff, 0x6c, 0xaf, 0xfe,
	0x04, 0x00, 0x00, 0xff, 0xff, 0x3c, 0x9e, 0x74, 0xb0, 0xec, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// StyxClient is the client API for Styx service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StyxClient interface {
	FileSearch(ctx context.Context, in *SearchFileInfo, opts ...grpc.CallOption) (*FoundFile, error)
	FileSearchStream(ctx context.Context, opts ...grpc.CallOption) (Styx_FileSearchStreamClient, error)
	ReqFileTransfer(ctx context.Context, in *RequestFileTransfer, opts ...grpc.CallOption) (*Ack, error)
	PullFile(ctx context.Context, in *GetRemoteFile, opts ...grpc.CallOption) (*Ack, error)
	// sftpmgmt functions
	GRPCTest(ctx context.Context, in *Ack, opts ...grpc.CallOption) (*Ack, error)
	AddConfig(ctx context.Context, in *SftpTransferConfig, opts ...grpc.CallOption) (*Ack, error)
	ListConfig(ctx context.Context, in *Noparams, opts ...grpc.CallOption) (Styx_ListConfigClient, error)
}

type styxClient struct {
	cc *grpc.ClientConn
}

func NewStyxClient(cc *grpc.ClientConn) StyxClient {
	return &styxClient{cc}
}

func (c *styxClient) FileSearch(ctx context.Context, in *SearchFileInfo, opts ...grpc.CallOption) (*FoundFile, error) {
	out := new(FoundFile)
	err := c.cc.Invoke(ctx, "/styxevent.Styx/FileSearch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *styxClient) FileSearchStream(ctx context.Context, opts ...grpc.CallOption) (Styx_FileSearchStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Styx_serviceDesc.Streams[0], "/styxevent.Styx/FileSearchStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &styxFileSearchStreamClient{stream}
	return x, nil
}

type Styx_FileSearchStreamClient interface {
	Send(*SearchFileInfo) error
	Recv() (*FoundFile, error)
	grpc.ClientStream
}

type styxFileSearchStreamClient struct {
	grpc.ClientStream
}

func (x *styxFileSearchStreamClient) Send(m *SearchFileInfo) error {
	return x.ClientStream.SendMsg(m)
}

func (x *styxFileSearchStreamClient) Recv() (*FoundFile, error) {
	m := new(FoundFile)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *styxClient) ReqFileTransfer(ctx context.Context, in *RequestFileTransfer, opts ...grpc.CallOption) (*Ack, error) {
	out := new(Ack)
	err := c.cc.Invoke(ctx, "/styxevent.Styx/ReqFileTransfer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *styxClient) PullFile(ctx context.Context, in *GetRemoteFile, opts ...grpc.CallOption) (*Ack, error) {
	out := new(Ack)
	err := c.cc.Invoke(ctx, "/styxevent.Styx/PullFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *styxClient) GRPCTest(ctx context.Context, in *Ack, opts ...grpc.CallOption) (*Ack, error) {
	out := new(Ack)
	err := c.cc.Invoke(ctx, "/styxevent.Styx/GRPCTest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *styxClient) AddConfig(ctx context.Context, in *SftpTransferConfig, opts ...grpc.CallOption) (*Ack, error) {
	out := new(Ack)
	err := c.cc.Invoke(ctx, "/styxevent.Styx/AddConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *styxClient) ListConfig(ctx context.Context, in *Noparams, opts ...grpc.CallOption) (Styx_ListConfigClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Styx_serviceDesc.Streams[1], "/styxevent.Styx/ListConfig", opts...)
	if err != nil {
		return nil, err
	}
	x := &styxListConfigClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Styx_ListConfigClient interface {
	Recv() (*SftpTransferConfig, error)
	grpc.ClientStream
}

type styxListConfigClient struct {
	grpc.ClientStream
}

func (x *styxListConfigClient) Recv() (*SftpTransferConfig, error) {
	m := new(SftpTransferConfig)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StyxServer is the server API for Styx service.
type StyxServer interface {
	FileSearch(context.Context, *SearchFileInfo) (*FoundFile, error)
	FileSearchStream(Styx_FileSearchStreamServer) error
	ReqFileTransfer(context.Context, *RequestFileTransfer) (*Ack, error)
	PullFile(context.Context, *GetRemoteFile) (*Ack, error)
	// sftpmgmt functions
	GRPCTest(context.Context, *Ack) (*Ack, error)
	AddConfig(context.Context, *SftpTransferConfig) (*Ack, error)
	ListConfig(*Noparams, Styx_ListConfigServer) error
}

func RegisterStyxServer(s *grpc.Server, srv StyxServer) {
	s.RegisterService(&_Styx_serviceDesc, srv)
}

func _Styx_FileSearch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchFileInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StyxServer).FileSearch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/styxevent.Styx/FileSearch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StyxServer).FileSearch(ctx, req.(*SearchFileInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Styx_FileSearchStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StyxServer).FileSearchStream(&styxFileSearchStreamServer{stream})
}

type Styx_FileSearchStreamServer interface {
	Send(*FoundFile) error
	Recv() (*SearchFileInfo, error)
	grpc.ServerStream
}

type styxFileSearchStreamServer struct {
	grpc.ServerStream
}

func (x *styxFileSearchStreamServer) Send(m *FoundFile) error {
	return x.ServerStream.SendMsg(m)
}

func (x *styxFileSearchStreamServer) Recv() (*SearchFileInfo, error) {
	m := new(SearchFileInfo)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Styx_ReqFileTransfer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestFileTransfer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StyxServer).ReqFileTransfer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/styxevent.Styx/ReqFileTransfer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StyxServer).ReqFileTransfer(ctx, req.(*RequestFileTransfer))
	}
	return interceptor(ctx, in, info, handler)
}

func _Styx_PullFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRemoteFile)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StyxServer).PullFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/styxevent.Styx/PullFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StyxServer).PullFile(ctx, req.(*GetRemoteFile))
	}
	return interceptor(ctx, in, info, handler)
}

func _Styx_GRPCTest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ack)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StyxServer).GRPCTest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/styxevent.Styx/GRPCTest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StyxServer).GRPCTest(ctx, req.(*Ack))
	}
	return interceptor(ctx, in, info, handler)
}

func _Styx_AddConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SftpTransferConfig)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StyxServer).AddConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/styxevent.Styx/AddConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StyxServer).AddConfig(ctx, req.(*SftpTransferConfig))
	}
	return interceptor(ctx, in, info, handler)
}

func _Styx_ListConfig_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Noparams)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StyxServer).ListConfig(m, &styxListConfigServer{stream})
}

type Styx_ListConfigServer interface {
	Send(*SftpTransferConfig) error
	grpc.ServerStream
}

type styxListConfigServer struct {
	grpc.ServerStream
}

func (x *styxListConfigServer) Send(m *SftpTransferConfig) error {
	return x.ServerStream.SendMsg(m)
}

var _Styx_serviceDesc = grpc.ServiceDesc{
	ServiceName: "styxevent.Styx",
	HandlerType: (*StyxServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FileSearch",
			Handler:    _Styx_FileSearch_Handler,
		},
		{
			MethodName: "ReqFileTransfer",
			Handler:    _Styx_ReqFileTransfer_Handler,
		},
		{
			MethodName: "PullFile",
			Handler:    _Styx_PullFile_Handler,
		},
		{
			MethodName: "GRPCTest",
			Handler:    _Styx_GRPCTest_Handler,
		},
		{
			MethodName: "AddConfig",
			Handler:    _Styx_AddConfig_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "FileSearchStream",
			Handler:       _Styx_FileSearchStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "ListConfig",
			Handler:       _Styx_ListConfig_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "styxevent.proto",
}
