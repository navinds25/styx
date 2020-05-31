// Code generated by protoc-gen-go. DO NOT EDIT.
// source: execute.proto

package execute

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

type Executable struct {
	Command              string   `protobuf:"bytes,1,opt,name=command,proto3" json:"command,omitempty"`
	Arguments            []string `protobuf:"bytes,2,rep,name=arguments,proto3" json:"arguments,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Executable) Reset()         { *m = Executable{} }
func (m *Executable) String() string { return proto.CompactTextString(m) }
func (*Executable) ProtoMessage()    {}
func (*Executable) Descriptor() ([]byte, []int) {
	return fileDescriptor_58179d2e1720ec81, []int{0}
}

func (m *Executable) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Executable.Unmarshal(m, b)
}
func (m *Executable) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Executable.Marshal(b, m, deterministic)
}
func (m *Executable) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Executable.Merge(m, src)
}
func (m *Executable) XXX_Size() int {
	return xxx_messageInfo_Executable.Size(m)
}
func (m *Executable) XXX_DiscardUnknown() {
	xxx_messageInfo_Executable.DiscardUnknown(m)
}

var xxx_messageInfo_Executable proto.InternalMessageInfo

func (m *Executable) GetCommand() string {
	if m != nil {
		return m.Command
	}
	return ""
}

func (m *Executable) GetArguments() []string {
	if m != nil {
		return m.Arguments
	}
	return nil
}

type Output struct {
	Output               string   `protobuf:"bytes,1,opt,name=output,proto3" json:"output,omitempty"`
	Error                string   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Output) Reset()         { *m = Output{} }
func (m *Output) String() string { return proto.CompactTextString(m) }
func (*Output) ProtoMessage()    {}
func (*Output) Descriptor() ([]byte, []int) {
	return fileDescriptor_58179d2e1720ec81, []int{1}
}

func (m *Output) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Output.Unmarshal(m, b)
}
func (m *Output) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Output.Marshal(b, m, deterministic)
}
func (m *Output) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Output.Merge(m, src)
}
func (m *Output) XXX_Size() int {
	return xxx_messageInfo_Output.Size(m)
}
func (m *Output) XXX_DiscardUnknown() {
	xxx_messageInfo_Output.DiscardUnknown(m)
}

var xxx_messageInfo_Output proto.InternalMessageInfo

func (m *Output) GetOutput() string {
	if m != nil {
		return m.Output
	}
	return ""
}

func (m *Output) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func init() {
	proto.RegisterType((*Executable)(nil), "execute.Executable")
	proto.RegisterType((*Output)(nil), "execute.Output")
}

func init() { proto.RegisterFile("execute.proto", fileDescriptor_58179d2e1720ec81) }

var fileDescriptor_58179d2e1720ec81 = []byte{
	// 197 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0xad, 0x48, 0x4d,
	0x2e, 0x2d, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x87, 0x72, 0x95, 0x5c, 0xb8,
	0xb8, 0x5c, 0xc1, 0xcc, 0xc4, 0xa4, 0x9c, 0x54, 0x21, 0x09, 0x2e, 0xf6, 0xe4, 0xfc, 0xdc, 0xdc,
	0xc4, 0xbc, 0x14, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x18, 0x57, 0x48, 0x86, 0x8b, 0x33,
	0xb1, 0x28, 0xbd, 0x34, 0x37, 0x35, 0xaf, 0xa4, 0x58, 0x82, 0x49, 0x81, 0x59, 0x83, 0x33, 0x08,
	0x21, 0xa0, 0x64, 0xc6, 0xc5, 0xe6, 0x5f, 0x5a, 0x52, 0x50, 0x5a, 0x22, 0x24, 0xc6, 0xc5, 0x96,
	0x0f, 0x66, 0x41, 0x0d, 0x80, 0xf2, 0x84, 0x44, 0xb8, 0x58, 0x53, 0x8b, 0x8a, 0xf2, 0x8b, 0x24,
	0x98, 0xc0, 0xc2, 0x10, 0x8e, 0x51, 0x05, 0x17, 0x1f, 0xc4, 0xf6, 0xd4, 0xe0, 0xd4, 0xa2, 0xb2,
	0xcc, 0xe4, 0x54, 0x21, 0x6d, 0x2e, 0xe6, 0xa0, 0xd2, 0x3c, 0x21, 0x61, 0x3d, 0x98, 0x7b, 0x11,
	0xae, 0x93, 0xe2, 0x87, 0x0b, 0x42, 0x2d, 0xb3, 0xe4, 0xe2, 0x0f, 0x2a, 0xcd, 0x0b, 0x2e, 0x29,
	0x4a, 0x4d, 0xcc, 0x85, 0x0a, 0x11, 0xa5, 0xd1, 0x80, 0xd1, 0x89, 0x3b, 0x8a, 0x53, 0xcf, 0x1a,
	0x2a, 0x9a, 0xc4, 0x06, 0x0e, 0x14, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x24, 0xf0, 0x32,
	0x15, 0x25, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ExecuteServiceClient is the client API for ExecuteService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ExecuteServiceClient interface {
	Run(ctx context.Context, in *Executable, opts ...grpc.CallOption) (*Output, error)
	RunStreamOutput(ctx context.Context, in *Executable, opts ...grpc.CallOption) (ExecuteService_RunStreamOutputClient, error)
}

type executeServiceClient struct {
	cc *grpc.ClientConn
}

func NewExecuteServiceClient(cc *grpc.ClientConn) ExecuteServiceClient {
	return &executeServiceClient{cc}
}

func (c *executeServiceClient) Run(ctx context.Context, in *Executable, opts ...grpc.CallOption) (*Output, error) {
	out := new(Output)
	err := c.cc.Invoke(ctx, "/execute.ExecuteService/Run", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *executeServiceClient) RunStreamOutput(ctx context.Context, in *Executable, opts ...grpc.CallOption) (ExecuteService_RunStreamOutputClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ExecuteService_serviceDesc.Streams[0], "/execute.ExecuteService/RunStreamOutput", opts...)
	if err != nil {
		return nil, err
	}
	x := &executeServiceRunStreamOutputClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ExecuteService_RunStreamOutputClient interface {
	Recv() (*Output, error)
	grpc.ClientStream
}

type executeServiceRunStreamOutputClient struct {
	grpc.ClientStream
}

func (x *executeServiceRunStreamOutputClient) Recv() (*Output, error) {
	m := new(Output)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ExecuteServiceServer is the server API for ExecuteService service.
type ExecuteServiceServer interface {
	Run(context.Context, *Executable) (*Output, error)
	RunStreamOutput(*Executable, ExecuteService_RunStreamOutputServer) error
}

func RegisterExecuteServiceServer(s *grpc.Server, srv ExecuteServiceServer) {
	s.RegisterService(&_ExecuteService_serviceDesc, srv)
}

func _ExecuteService_Run_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Executable)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExecuteServiceServer).Run(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/execute.ExecuteService/Run",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExecuteServiceServer).Run(ctx, req.(*Executable))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExecuteService_RunStreamOutput_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Executable)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ExecuteServiceServer).RunStreamOutput(m, &executeServiceRunStreamOutputServer{stream})
}

type ExecuteService_RunStreamOutputServer interface {
	Send(*Output) error
	grpc.ServerStream
}

type executeServiceRunStreamOutputServer struct {
	grpc.ServerStream
}

func (x *executeServiceRunStreamOutputServer) Send(m *Output) error {
	return x.ServerStream.SendMsg(m)
}

var _ExecuteService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "execute.ExecuteService",
	HandlerType: (*ExecuteServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Run",
			Handler:    _ExecuteService_Run_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "RunStreamOutput",
			Handler:       _ExecuteService_RunStreamOutput_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "execute.proto",
}
