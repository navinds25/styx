// Code generated by protoc-gen-go. DO NOT EDIT.
// source: filetransfer.proto

package filetransfer

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

func init() { proto.RegisterFile("filetransfer.proto", fileDescriptor_85d5b4bd112d6203) }

var fileDescriptor_85d5b4bd112d6203 = []byte{
	// 198 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4a, 0xcb, 0xcc, 0x49,
	0x2d, 0x29, 0x4a, 0xcc, 0x2b, 0x4e, 0x4b, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2,
	0x41, 0x16, 0x93, 0x92, 0x44, 0xe6, 0xc5, 0xa7, 0x96, 0xa5, 0xe6, 0x95, 0x14, 0x43, 0x14, 0x1a,
	0xad, 0x64, 0xe4, 0xe2, 0x74, 0x0b, 0x09, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0x15, 0x72, 0xe1,
	0x62, 0xf7, 0xc9, 0x2c, 0x2e, 0x71, 0xc9, 0x2c, 0x12, 0x92, 0xd5, 0x43, 0x31, 0x36, 0x28, 0x35,
	0x37, 0xbf, 0x24, 0xd5, 0x25, 0xb3, 0x28, 0x35, 0xb9, 0x24, 0xbf, 0xa8, 0x52, 0x4a, 0x02, 0x9b,
	0xb4, 0x5b, 0x66, 0x4e, 0xaa, 0x01, 0xa3, 0x50, 0x18, 0x97, 0x00, 0xd4, 0x14, 0xe7, 0xfc, 0xbc,
	0x94, 0xcc, 0x92, 0xcc, 0xfc, 0x3c, 0x21, 0x35, 0xbc, 0xc6, 0xc1, 0xd5, 0xe1, 0x33, 0xd7, 0xe8,
	0x2c, 0x23, 0x17, 0x3f, 0x54, 0x00, 0xee, 0xe2, 0x50, 0x2e, 0x9e, 0x10, 0xa8, 0x62, 0x90, 0x2a,
	0x21, 0x55, 0xbc, 0xf6, 0xc0, 0x94, 0x4a, 0xc9, 0xa0, 0x2a, 0x83, 0x89, 0x07, 0x97, 0x24, 0x96,
	0x94, 0x16, 0x1b, 0x30, 0x0a, 0x25, 0x73, 0x09, 0xc2, 0xc4, 0x10, 0x7e, 0xd0, 0x23, 0xce, 0x0f,
	0xc4, 0x5a, 0x92, 0xc4, 0x06, 0x8e, 0x02, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x2e, 0x79,
	0x9c, 0xc3, 0xc1, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FTServiceClient is the client API for FTService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FTServiceClient interface {
	ListDir(ctx context.Context, in *RemoteDirectory, opts ...grpc.CallOption) (FTService_ListDirClient, error)
	ListDirCondition(ctx context.Context, in *RemoteDirectoryCondition, opts ...grpc.CallOption) (FTService_ListDirConditionClient, error)
}

type fTServiceClient struct {
	cc *grpc.ClientConn
}

func NewFTServiceClient(cc *grpc.ClientConn) FTServiceClient {
	return &fTServiceClient{cc}
}

func (c *fTServiceClient) ListDir(ctx context.Context, in *RemoteDirectory, opts ...grpc.CallOption) (FTService_ListDirClient, error) {
	stream, err := c.cc.NewStream(ctx, &_FTService_serviceDesc.Streams[0], "/filetransfer.FTService/ListDir", opts...)
	if err != nil {
		return nil, err
	}
	x := &fTServiceListDirClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type FTService_ListDirClient interface {
	Recv() (*RemoteFile, error)
	grpc.ClientStream
}

type fTServiceListDirClient struct {
	grpc.ClientStream
}

func (x *fTServiceListDirClient) Recv() (*RemoteFile, error) {
	m := new(RemoteFile)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *fTServiceClient) ListDirCondition(ctx context.Context, in *RemoteDirectoryCondition, opts ...grpc.CallOption) (FTService_ListDirConditionClient, error) {
	stream, err := c.cc.NewStream(ctx, &_FTService_serviceDesc.Streams[1], "/filetransfer.FTService/ListDirCondition", opts...)
	if err != nil {
		return nil, err
	}
	x := &fTServiceListDirConditionClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type FTService_ListDirConditionClient interface {
	Recv() (*RemoteFile, error)
	grpc.ClientStream
}

type fTServiceListDirConditionClient struct {
	grpc.ClientStream
}

func (x *fTServiceListDirConditionClient) Recv() (*RemoteFile, error) {
	m := new(RemoteFile)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// FTServiceServer is the server API for FTService service.
type FTServiceServer interface {
	ListDir(*RemoteDirectory, FTService_ListDirServer) error
	ListDirCondition(*RemoteDirectoryCondition, FTService_ListDirConditionServer) error
}

func RegisterFTServiceServer(s *grpc.Server, srv FTServiceServer) {
	s.RegisterService(&_FTService_serviceDesc, srv)
}

func _FTService_ListDir_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RemoteDirectory)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FTServiceServer).ListDir(m, &fTServiceListDirServer{stream})
}

type FTService_ListDirServer interface {
	Send(*RemoteFile) error
	grpc.ServerStream
}

type fTServiceListDirServer struct {
	grpc.ServerStream
}

func (x *fTServiceListDirServer) Send(m *RemoteFile) error {
	return x.ServerStream.SendMsg(m)
}

func _FTService_ListDirCondition_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RemoteDirectoryCondition)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FTServiceServer).ListDirCondition(m, &fTServiceListDirConditionServer{stream})
}

type FTService_ListDirConditionServer interface {
	Send(*RemoteFile) error
	grpc.ServerStream
}

type fTServiceListDirConditionServer struct {
	grpc.ServerStream
}

func (x *fTServiceListDirConditionServer) Send(m *RemoteFile) error {
	return x.ServerStream.SendMsg(m)
}

var _FTService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "filetransfer.FTService",
	HandlerType: (*FTServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListDir",
			Handler:       _FTService_ListDir_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ListDirCondition",
			Handler:       _FTService_ListDirCondition_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "filetransfer.proto",
}

// RemoteFTServiceClient is the client API for RemoteFTService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RemoteFTServiceClient interface {
	TransferFile(ctx context.Context, in *RemoteDirectoryTransfer, opts ...grpc.CallOption) (RemoteFTService_TransferFileClient, error)
	TransferCondition(ctx context.Context, in *RemoteDirectoryConditionTransfer, opts ...grpc.CallOption) (RemoteFTService_TransferConditionClient, error)
}

type remoteFTServiceClient struct {
	cc *grpc.ClientConn
}

func NewRemoteFTServiceClient(cc *grpc.ClientConn) RemoteFTServiceClient {
	return &remoteFTServiceClient{cc}
}

func (c *remoteFTServiceClient) TransferFile(ctx context.Context, in *RemoteDirectoryTransfer, opts ...grpc.CallOption) (RemoteFTService_TransferFileClient, error) {
	stream, err := c.cc.NewStream(ctx, &_RemoteFTService_serviceDesc.Streams[0], "/filetransfer.RemoteFTService/TransferFile", opts...)
	if err != nil {
		return nil, err
	}
	x := &remoteFTServiceTransferFileClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type RemoteFTService_TransferFileClient interface {
	Recv() (*TransferStatus, error)
	grpc.ClientStream
}

type remoteFTServiceTransferFileClient struct {
	grpc.ClientStream
}

func (x *remoteFTServiceTransferFileClient) Recv() (*TransferStatus, error) {
	m := new(TransferStatus)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *remoteFTServiceClient) TransferCondition(ctx context.Context, in *RemoteDirectoryConditionTransfer, opts ...grpc.CallOption) (RemoteFTService_TransferConditionClient, error) {
	stream, err := c.cc.NewStream(ctx, &_RemoteFTService_serviceDesc.Streams[1], "/filetransfer.RemoteFTService/TransferCondition", opts...)
	if err != nil {
		return nil, err
	}
	x := &remoteFTServiceTransferConditionClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type RemoteFTService_TransferConditionClient interface {
	Recv() (*TransferStatus, error)
	grpc.ClientStream
}

type remoteFTServiceTransferConditionClient struct {
	grpc.ClientStream
}

func (x *remoteFTServiceTransferConditionClient) Recv() (*TransferStatus, error) {
	m := new(TransferStatus)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RemoteFTServiceServer is the server API for RemoteFTService service.
type RemoteFTServiceServer interface {
	TransferFile(*RemoteDirectoryTransfer, RemoteFTService_TransferFileServer) error
	TransferCondition(*RemoteDirectoryConditionTransfer, RemoteFTService_TransferConditionServer) error
}

func RegisterRemoteFTServiceServer(s *grpc.Server, srv RemoteFTServiceServer) {
	s.RegisterService(&_RemoteFTService_serviceDesc, srv)
}

func _RemoteFTService_TransferFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RemoteDirectoryTransfer)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RemoteFTServiceServer).TransferFile(m, &remoteFTServiceTransferFileServer{stream})
}

type RemoteFTService_TransferFileServer interface {
	Send(*TransferStatus) error
	grpc.ServerStream
}

type remoteFTServiceTransferFileServer struct {
	grpc.ServerStream
}

func (x *remoteFTServiceTransferFileServer) Send(m *TransferStatus) error {
	return x.ServerStream.SendMsg(m)
}

func _RemoteFTService_TransferCondition_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RemoteDirectoryConditionTransfer)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RemoteFTServiceServer).TransferCondition(m, &remoteFTServiceTransferConditionServer{stream})
}

type RemoteFTService_TransferConditionServer interface {
	Send(*TransferStatus) error
	grpc.ServerStream
}

type remoteFTServiceTransferConditionServer struct {
	grpc.ServerStream
}

func (x *remoteFTServiceTransferConditionServer) Send(m *TransferStatus) error {
	return x.ServerStream.SendMsg(m)
}

var _RemoteFTService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "filetransfer.RemoteFTService",
	HandlerType: (*RemoteFTServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "TransferFile",
			Handler:       _RemoteFTService_TransferFile_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "TransferCondition",
			Handler:       _RemoteFTService_TransferCondition_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "filetransfer.proto",
}
