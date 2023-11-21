// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.15.8
// source: rpc_services.proto

package pb

import (
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

type EmptyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyRequest) Reset() {
	*x = EmptyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_services_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyRequest) ProtoMessage() {}

func (x *EmptyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_services_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyRequest.ProtoReflect.Descriptor instead.
func (*EmptyRequest) Descriptor() ([]byte, []int) {
	return file_rpc_services_proto_rawDescGZIP(), []int{0}
}

var File_rpc_services_proto protoreflect.FileDescriptor

var file_rpc_services_proto_rawDesc = []byte{
	0x0a, 0x12, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x13, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x5f,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x72,
	0x70, 0x63, 0x5f, 0x67, 0x65, 0x74, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x0f, 0x72, 0x70, 0x63, 0x5f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x10, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x0e, 0x0a, 0x0c, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x32, 0xc6, 0x01, 0x0a, 0x11, 0x47, 0x72, 0x70, 0x63, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3f, 0x0a, 0x06, 0x53, 0x69,
	0x67, 0x6e, 0x55, 0x70, 0x12, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x75, 0x70,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x19,
	0x2e, 0x70, 0x62, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x05, 0x6c,
	0x6f, 0x67, 0x69, 0x6e, 0x12, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x18, 0x2e,
	0x70, 0x62, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x07, 0x47, 0x65, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x12, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x3d, 0x5a,
	0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x52, 0x61, 0x6b, 0x73,
	0x68, 0x69, 0x74, 0x4b, 0x75, 0x6d, 0x61, 0x72, 0x30, 0x34, 0x2f, 0x43, 0x68, 0x61, 0x74, 0x4d,
	0x61, 0x74, 0x65, 0x2f, 0x74, 0x72, 0x65, 0x65, 0x2f, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x2f,
	0x67, 0x6f, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_services_proto_rawDescOnce sync.Once
	file_rpc_services_proto_rawDescData = file_rpc_services_proto_rawDesc
)

func file_rpc_services_proto_rawDescGZIP() []byte {
	file_rpc_services_proto_rawDescOnce.Do(func() {
		file_rpc_services_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_services_proto_rawDescData)
	})
	return file_rpc_services_proto_rawDescData
}

var file_rpc_services_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_rpc_services_proto_goTypes = []interface{}{
	(*EmptyRequest)(nil),          // 0: pb.EmptyRequest
	(*SignupRequestMessage)(nil),  // 1: pb.SignupRequestMessage
	(*LoginRequestMessage)(nil),   // 2: pb.LoginRequestMessage
	(*SignupResponseMessage)(nil), // 3: pb.SignupResponseMessage
	(*LoginResponseMessage)(nil),  // 4: pb.LoginResponseMessage
	(*GetUserResponse)(nil),       // 5: pb.GetUserResponse
}
var file_rpc_services_proto_depIdxs = []int32{
	1, // 0: pb.GrpcServerService.SignUp:input_type -> pb.SignupRequestMessage
	2, // 1: pb.GrpcServerService.login:input_type -> pb.LoginRequestMessage
	0, // 2: pb.GrpcServerService.GetUser:input_type -> pb.EmptyRequest
	3, // 3: pb.GrpcServerService.SignUp:output_type -> pb.SignupResponseMessage
	4, // 4: pb.GrpcServerService.login:output_type -> pb.LoginResponseMessage
	5, // 5: pb.GrpcServerService.GetUser:output_type -> pb.GetUserResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_rpc_services_proto_init() }
func file_rpc_services_proto_init() {
	if File_rpc_services_proto != nil {
		return
	}
	file_empty_request_proto_init()
	file_rpc_get_user_proto_init()
	file_rpc_login_proto_init()
	file_rpc_signup_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_rpc_services_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyRequest); i {
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
			RawDescriptor: file_rpc_services_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rpc_services_proto_goTypes,
		DependencyIndexes: file_rpc_services_proto_depIdxs,
		MessageInfos:      file_rpc_services_proto_msgTypes,
	}.Build()
	File_rpc_services_proto = out.File
	file_rpc_services_proto_rawDesc = nil
	file_rpc_services_proto_goTypes = nil
	file_rpc_services_proto_depIdxs = nil
}
