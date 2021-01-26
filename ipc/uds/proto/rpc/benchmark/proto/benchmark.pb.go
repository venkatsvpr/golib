// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.7.0
// source: benchmark.proto

package benchmark

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type BenchmarkReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *BenchmarkReq) Reset() {
	*x = BenchmarkReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_benchmark_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BenchmarkReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BenchmarkReq) ProtoMessage() {}

func (x *BenchmarkReq) ProtoReflect() protoreflect.Message {
	mi := &file_benchmark_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BenchmarkReq.ProtoReflect.Descriptor instead.
func (*BenchmarkReq) Descriptor() ([]byte, []int) {
	return file_benchmark_proto_rawDescGZIP(), []int{0}
}

func (x *BenchmarkReq) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type BenchmarkResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *BenchmarkResp) Reset() {
	*x = BenchmarkResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_benchmark_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BenchmarkResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BenchmarkResp) ProtoMessage() {}

func (x *BenchmarkResp) ProtoReflect() protoreflect.Message {
	mi := &file_benchmark_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BenchmarkResp.ProtoReflect.Descriptor instead.
func (*BenchmarkResp) Descriptor() ([]byte, []int) {
	return file_benchmark_proto_rawDescGZIP(), []int{1}
}

func (x *BenchmarkResp) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_benchmark_proto protoreflect.FileDescriptor

var file_benchmark_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x62, 0x65, 0x6e, 0x63, 0x68, 0x6d, 0x61, 0x72, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x22, 0x0a, 0x0c, 0x42, 0x65, 0x6e, 0x63, 0x68, 0x6d, 0x61, 0x72, 0x6b, 0x52, 0x65,
	0x71, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x23, 0x0a, 0x0d, 0x42, 0x65, 0x6e, 0x63, 0x68, 0x6d, 0x61,
	0x72, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0x39, 0x0a, 0x09, 0x42, 0x65,
	0x6e, 0x63, 0x68, 0x6d, 0x61, 0x72, 0x6b, 0x12, 0x2c, 0x0a, 0x09, 0x42, 0x65, 0x6e, 0x63, 0x68,
	0x6d, 0x61, 0x72, 0x6b, 0x12, 0x0d, 0x2e, 0x42, 0x65, 0x6e, 0x63, 0x68, 0x6d, 0x61, 0x72, 0x6b,
	0x52, 0x65, 0x71, 0x1a, 0x0e, 0x2e, 0x42, 0x65, 0x6e, 0x63, 0x68, 0x6d, 0x61, 0x72, 0x6b, 0x52,
	0x65, 0x73, 0x70, 0x22, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_benchmark_proto_rawDescOnce sync.Once
	file_benchmark_proto_rawDescData = file_benchmark_proto_rawDesc
)

func file_benchmark_proto_rawDescGZIP() []byte {
	file_benchmark_proto_rawDescOnce.Do(func() {
		file_benchmark_proto_rawDescData = protoimpl.X.CompressGZIP(file_benchmark_proto_rawDescData)
	})
	return file_benchmark_proto_rawDescData
}

var file_benchmark_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_benchmark_proto_goTypes = []interface{}{
	(*BenchmarkReq)(nil),  // 0: BenchmarkReq
	(*BenchmarkResp)(nil), // 1: BenchmarkResp
}
var file_benchmark_proto_depIdxs = []int32{
	0, // 0: Benchmark.Benchmark:input_type -> BenchmarkReq
	1, // 1: Benchmark.Benchmark:output_type -> BenchmarkResp
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_benchmark_proto_init() }
func file_benchmark_proto_init() {
	if File_benchmark_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_benchmark_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BenchmarkReq); i {
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
		file_benchmark_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BenchmarkResp); i {
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
			RawDescriptor: file_benchmark_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_benchmark_proto_goTypes,
		DependencyIndexes: file_benchmark_proto_depIdxs,
		MessageInfos:      file_benchmark_proto_msgTypes,
	}.Build()
	File_benchmark_proto = out.File
	file_benchmark_proto_rawDesc = nil
	file_benchmark_proto_goTypes = nil
	file_benchmark_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// BenchmarkClient is the client API for Benchmark service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BenchmarkClient interface {
	Benchmark(ctx context.Context, in *BenchmarkReq, opts ...grpc.CallOption) (*BenchmarkResp, error)
}

type benchmarkClient struct {
	cc grpc.ClientConnInterface
}

func NewBenchmarkClient(cc grpc.ClientConnInterface) BenchmarkClient {
	return &benchmarkClient{cc}
}

func (c *benchmarkClient) Benchmark(ctx context.Context, in *BenchmarkReq, opts ...grpc.CallOption) (*BenchmarkResp, error) {
	out := new(BenchmarkResp)
	err := c.cc.Invoke(ctx, "/Benchmark/Benchmark", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BenchmarkServer is the server API for Benchmark service.
type BenchmarkServer interface {
	Benchmark(context.Context, *BenchmarkReq) (*BenchmarkResp, error)
}

// UnimplementedBenchmarkServer can be embedded to have forward compatible implementations.
type UnimplementedBenchmarkServer struct {
}

func (*UnimplementedBenchmarkServer) Benchmark(context.Context, *BenchmarkReq) (*BenchmarkResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Benchmark not implemented")
}

func RegisterBenchmarkServer(s *grpc.Server, srv BenchmarkServer) {
	s.RegisterService(&_Benchmark_serviceDesc, srv)
}

func _Benchmark_Benchmark_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BenchmarkReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BenchmarkServer).Benchmark(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Benchmark/Benchmark",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BenchmarkServer).Benchmark(ctx, req.(*BenchmarkReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Benchmark_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Benchmark",
	HandlerType: (*BenchmarkServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Benchmark",
			Handler:    _Benchmark_Benchmark_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "benchmark.proto",
}
