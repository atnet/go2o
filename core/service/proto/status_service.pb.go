// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v5.26.1
// source: status_service.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_status_service_proto protoreflect.FileDescriptor

var file_status_service_proto_rawDesc = []byte{
	0x0a, 0x14, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x32, 0x47, 0x0a, 0x0d, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x19, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x06, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x07, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x22, 0x00,
	0x12, 0x1b, 0x0a, 0x05, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x07, 0x2e, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x1a, 0x07, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x22, 0x00, 0x42, 0x1f, 0x0a,
	0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x67, 0x6f, 0x32, 0x6f,
	0x2e, 0x72, 0x70, 0x63, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_status_service_proto_goTypes = []interface{}{
	(*Empty)(nil),  // 0: Empty
	(*String)(nil), // 1: String
}
var file_status_service_proto_depIdxs = []int32{
	0, // 0: StatusService.Ping:input_type -> Empty
	1, // 1: StatusService.Hello:input_type -> String
	1, // 2: StatusService.Ping:output_type -> String
	1, // 3: StatusService.Hello:output_type -> String
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_status_service_proto_init() }
func file_status_service_proto_init() {
	if File_status_service_proto != nil {
		return
	}
	file_global_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_status_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_status_service_proto_goTypes,
		DependencyIndexes: file_status_service_proto_depIdxs,
	}.Build()
	File_status_service_proto = out.File
	file_status_service_proto_rawDesc = nil
	file_status_service_proto_goTypes = nil
	file_status_service_proto_depIdxs = nil
}
