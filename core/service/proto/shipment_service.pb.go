// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.27.0
// source: shipment_service.proto

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

var File_shipment_service_proto protoreflect.FileDescriptor

var file_shipment_service_proto_rawDesc = []byte{
	0x0a, 0x16, 0x73, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x64, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1a, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x73, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e,
	0x74, 0x5f, 0x64, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x92, 0x02, 0x0a, 0x0f,
	0x53, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x31, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x76, 0x65, 0x72, 0x61, 0x67,
	0x65, 0x41, 0x72, 0x65, 0x61, 0x5f, 0x12, 0x0f, 0x2e, 0x53, 0x43, 0x6f, 0x76, 0x65, 0x72, 0x61,
	0x67, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x07, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x22, 0x00, 0x12, 0x3b, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x68,
	0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x08, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49,
	0x64, 0x1a, 0x1a, 0x2e, 0x53, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x45, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x46, 0x6c,
	0x6f, 0x77, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x12, 0x19, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x73, 0x74,
	0x69, 0x63, 0x46, 0x6c, 0x6f, 0x77, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x10, 0x2e, 0x53, 0x53, 0x68, 0x69, 0x70, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x54,
	0x72, 0x61, 0x63, 0x6b, 0x22, 0x00, 0x12, 0x48, 0x0a, 0x16, 0x53, 0x68, 0x69, 0x70, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x54, 0x72, 0x61, 0x63, 0x6b,
	0x12, 0x1a, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63,
	0x54, 0x72, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x53,
	0x53, 0x68, 0x69, 0x70, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x22, 0x00,
	0x42, 0x1f, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x67,
	0x6f, 0x32, 0x6f, 0x2e, 0x72, 0x70, 0x63, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_shipment_service_proto_goTypes = []interface{}{
	(*SCoverageValue)(nil),            // 0: SCoverageValue
	(*OrderId)(nil),                   // 1: OrderId
	(*LogisticFlowTrackRequest)(nil),  // 2: LogisticFlowTrackRequest
	(*OrderLogisticTrackRequest)(nil), // 3: OrderLogisticTrackRequest
	(*Result)(nil),                    // 4: Result
	(*ShipmentOrderListResponse)(nil), // 5: ShipmentOrderListResponse
	(*SShipOrderTrack)(nil),           // 6: SShipOrderTrack
}
var file_shipment_service_proto_depIdxs = []int32{
	0, // 0: ShipmentService.CreateCoverageArea_:input_type -> SCoverageValue
	1, // 1: ShipmentService.GetOrderShipments:input_type -> OrderId
	2, // 2: ShipmentService.GetLogisticFlowTrack:input_type -> LogisticFlowTrackRequest
	3, // 3: ShipmentService.ShipOrderLogisticTrack:input_type -> OrderLogisticTrackRequest
	4, // 4: ShipmentService.CreateCoverageArea_:output_type -> Result
	5, // 5: ShipmentService.GetOrderShipments:output_type -> ShipmentOrderListResponse
	6, // 6: ShipmentService.GetLogisticFlowTrack:output_type -> SShipOrderTrack
	6, // 7: ShipmentService.ShipOrderLogisticTrack:output_type -> SShipOrderTrack
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_shipment_service_proto_init() }
func file_shipment_service_proto_init() {
	if File_shipment_service_proto != nil {
		return
	}
	file_global_proto_init()
	file_message_order_dto_proto_init()
	file_message_shipment_dto_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_shipment_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_shipment_service_proto_goTypes,
		DependencyIndexes: file_shipment_service_proto_depIdxs,
	}.Build()
	File_shipment_service_proto = out.File
	file_shipment_service_proto_rawDesc = nil
	file_shipment_service_proto_goTypes = nil
	file_shipment_service_proto_depIdxs = nil
}
