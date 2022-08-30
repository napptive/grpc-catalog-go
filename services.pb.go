// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0-devel
// 	protoc        v3.13.0
// source: catalog/services.proto

package grpc_catalog_go

import (
	context "context"
	grpc_catalog_common_go "github.com/napptive/grpc-catalog-common-go"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

var File_catalog_services_proto protoreflect.FileDescriptor

var file_catalog_services_proto_rawDesc = []byte{
	0x0a, 0x16, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f,
	0x67, 0x1a, 0x16, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x69, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x63, 0x61, 0x74, 0x61, 0x6c,
	0x6f, 0x67, 0x2d, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69,
	0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xe2, 0x04, 0x0a, 0x07, 0x43, 0x61, 0x74, 0x61, 0x6c,
	0x6f, 0x67, 0x12, 0x5f, 0x0a, 0x03, 0x41, 0x64, 0x64, 0x12, 0x1e, 0x2e, 0x63, 0x61, 0x74, 0x61,
	0x6c, 0x6f, 0x67, 0x2e, 0x41, 0x64, 0x64, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x63, 0x61, 0x74, 0x61,
	0x6c, 0x6f, 0x67, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4f, 0x70, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x22, 0x0f, 0x2f,
	0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2f, 0x76, 0x30, 0x2f, 0x61, 0x64, 0x64, 0x3a, 0x01,
	0x2a, 0x28, 0x01, 0x12, 0x65, 0x0a, 0x08, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x12,
	0x23, 0x2e, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f,
	0x61, 0x64, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x46,
	0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x22,
	0x14, 0x2f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2f, 0x76, 0x30, 0x2f, 0x64, 0x6f, 0x77,
	0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x3a, 0x01, 0x2a, 0x30, 0x01, 0x12, 0x66, 0x0a, 0x06, 0x52, 0x65,
	0x6d, 0x6f, 0x76, 0x65, 0x12, 0x21, 0x2e, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x52,
	0x65, 0x6d, 0x6f, 0x76, 0x65, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f,
	0x67, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4f, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x22, 0x12, 0x2f, 0x63, 0x61,
	0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2f, 0x76, 0x30, 0x2f, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x3a,
	0x01, 0x2a, 0x12, 0x5f, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x20, 0x2e, 0x63, 0x61, 0x74,
	0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x63,
	0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x22, 0x10,
	0x2f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2f, 0x76, 0x30, 0x2f, 0x6c, 0x69, 0x73, 0x74,
	0x3a, 0x01, 0x2a, 0x12, 0x66, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1f, 0x2e, 0x63, 0x61,
	0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x63,
	0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x41, 0x70, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1b,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x22, 0x10, 0x2f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67,
	0x2f, 0x76, 0x30, 0x2f, 0x69, 0x6e, 0x66, 0x6f, 0x3a, 0x01, 0x2a, 0x12, 0x5e, 0x0a, 0x07, 0x53,
	0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x1c, 0x2e, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67,
	0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x53,
	0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1b,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x12, 0x13, 0x2f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67,
	0x2f, 0x76, 0x30, 0x2f, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x32, 0xfe, 0x01, 0x0a, 0x17,
	0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73,
	0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x47, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x12, 0x1f, 0x2e, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x5f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x4f, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x54, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x2e, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e,
	0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x63, 0x61, 0x74, 0x61, 0x6c,
	0x6f, 0x67, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4f, 0x70, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x20,
	0x2e, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x70, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x18, 0x2e, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x41, 0x70, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x32, 0xfe, 0x01, 0x0a,
	0x0c, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x6a, 0x0a,
	0x06, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x12, 0x21, 0x2e, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f,
	0x67, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x63, 0x61, 0x74,
	0x61, 0x6c, 0x6f, 0x67, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4f, 0x70, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x22, 0x16,
	0x2f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2f, 0x76, 0x30, 0x2f, 0x61, 0x70, 0x70, 0x2f,
	0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x3a, 0x01, 0x2a, 0x12, 0x81, 0x01, 0x0a, 0x10, 0x47, 0x65,
	0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20,
	0x2e, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x21, 0x2e, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x28, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x22, 0x22, 0x1d, 0x2f, 0x63, 0x61,
	0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2f, 0x76, 0x30, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x01, 0x2a, 0x42, 0x35, 0x5a,
	0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x61, 0x70, 0x70,
	0x74, 0x69, 0x76, 0x65, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f,
	0x67, 0x2d, 0x67, 0x6f, 0x3b, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f,
	0x67, 0x5f, 0x67, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_catalog_services_proto_goTypes = []interface{}{
	(*AddApplicationRequest)(nil),               // 0: catalog.AddApplicationRequest
	(*DownloadApplicationRequest)(nil),          // 1: catalog.DownloadApplicationRequest
	(*RemoveApplicationRequest)(nil),            // 2: catalog.RemoveApplicationRequest
	(*ListApplicationsRequest)(nil),             // 3: catalog.ListApplicationsRequest
	(*InfoApplicationRequest)(nil),              // 4: catalog.InfoApplicationRequest
	(*grpc_catalog_common_go.EmptyRequest)(nil), // 5: catalog_common.EmptyRequest
	(*DeleteNamespaceRequest)(nil),              // 6: catalog.DeleteNamespaceRequest
	(*DeployApplicationRequest)(nil),            // 7: catalog.DeployApplicationRequest
	(*GetConfigurationRequest)(nil),             // 8: catalog.GetConfigurationRequest
	(*grpc_catalog_common_go.OpResponse)(nil),   // 9: catalog_common.OpResponse
	(*FileInfo)(nil),                            // 10: catalog.FileInfo
	(*ApplicationList)(nil),                     // 11: catalog.ApplicationList
	(*InfoApplicationResponse)(nil),             // 12: catalog.InfoApplicationResponse
	(*SummaryResponse)(nil),                     // 13: catalog.SummaryResponse
	(*GetConfigurationResponse)(nil),            // 14: catalog.GetConfigurationResponse
}
var file_catalog_services_proto_depIdxs = []int32{
	0,  // 0: catalog.Catalog.Add:input_type -> catalog.AddApplicationRequest
	1,  // 1: catalog.Catalog.Download:input_type -> catalog.DownloadApplicationRequest
	2,  // 2: catalog.Catalog.Remove:input_type -> catalog.RemoveApplicationRequest
	3,  // 3: catalog.Catalog.List:input_type -> catalog.ListApplicationsRequest
	4,  // 4: catalog.Catalog.Info:input_type -> catalog.InfoApplicationRequest
	5,  // 5: catalog.Catalog.Summary:input_type -> catalog_common.EmptyRequest
	6,  // 6: catalog.NamespaceAdministration.Delete:input_type -> catalog.DeleteNamespaceRequest
	2,  // 7: catalog.NamespaceAdministration.DeleteApplication:input_type -> catalog.RemoveApplicationRequest
	3,  // 8: catalog.NamespaceAdministration.List:input_type -> catalog.ListApplicationsRequest
	7,  // 9: catalog.Applications.Deploy:input_type -> catalog.DeployApplicationRequest
	8,  // 10: catalog.Applications.GetConfiguration:input_type -> catalog.GetConfigurationRequest
	9,  // 11: catalog.Catalog.Add:output_type -> catalog_common.OpResponse
	10, // 12: catalog.Catalog.Download:output_type -> catalog.FileInfo
	9,  // 13: catalog.Catalog.Remove:output_type -> catalog_common.OpResponse
	11, // 14: catalog.Catalog.List:output_type -> catalog.ApplicationList
	12, // 15: catalog.Catalog.Info:output_type -> catalog.InfoApplicationResponse
	13, // 16: catalog.Catalog.Summary:output_type -> catalog.SummaryResponse
	9,  // 17: catalog.NamespaceAdministration.Delete:output_type -> catalog_common.OpResponse
	9,  // 18: catalog.NamespaceAdministration.DeleteApplication:output_type -> catalog_common.OpResponse
	11, // 19: catalog.NamespaceAdministration.List:output_type -> catalog.ApplicationList
	9,  // 20: catalog.Applications.Deploy:output_type -> catalog_common.OpResponse
	14, // 21: catalog.Applications.GetConfiguration:output_type -> catalog.GetConfigurationResponse
	11, // [11:22] is the sub-list for method output_type
	0,  // [0:11] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_catalog_services_proto_init() }
func file_catalog_services_proto_init() {
	if File_catalog_services_proto != nil {
		return
	}
	file_catalog_entities_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_catalog_services_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   3,
		},
		GoTypes:           file_catalog_services_proto_goTypes,
		DependencyIndexes: file_catalog_services_proto_depIdxs,
	}.Build()
	File_catalog_services_proto = out.File
	file_catalog_services_proto_rawDesc = nil
	file_catalog_services_proto_goTypes = nil
	file_catalog_services_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CatalogClient is the client API for Catalog service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CatalogClient interface {
	// Add a new application to the catalog.
	Add(ctx context.Context, opts ...grpc.CallOption) (Catalog_AddClient, error)
	// Download an application from the catalog.
	Download(ctx context.Context, in *DownloadApplicationRequest, opts ...grpc.CallOption) (Catalog_DownloadClient, error)
	//Remove an application from the catalog.
	Remove(ctx context.Context, in *RemoveApplicationRequest, opts ...grpc.CallOption) (*grpc_catalog_common_go.OpResponse, error)
	// List the applications in the catalog.
	List(ctx context.Context, in *ListApplicationsRequest, opts ...grpc.CallOption) (*ApplicationList, error)
	// Info returns the detail of a given application.
	Info(ctx context.Context, in *InfoApplicationRequest, opts ...grpc.CallOption) (*InfoApplicationResponse, error)
	// Summary returns the catalog summary
	Summary(ctx context.Context, in *grpc_catalog_common_go.EmptyRequest, opts ...grpc.CallOption) (*SummaryResponse, error)
}

type catalogClient struct {
	cc grpc.ClientConnInterface
}

func NewCatalogClient(cc grpc.ClientConnInterface) CatalogClient {
	return &catalogClient{cc}
}

func (c *catalogClient) Add(ctx context.Context, opts ...grpc.CallOption) (Catalog_AddClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Catalog_serviceDesc.Streams[0], "/catalog.Catalog/Add", opts...)
	if err != nil {
		return nil, err
	}
	x := &catalogAddClient{stream}
	return x, nil
}

type Catalog_AddClient interface {
	Send(*AddApplicationRequest) error
	CloseAndRecv() (*grpc_catalog_common_go.OpResponse, error)
	grpc.ClientStream
}

type catalogAddClient struct {
	grpc.ClientStream
}

func (x *catalogAddClient) Send(m *AddApplicationRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *catalogAddClient) CloseAndRecv() (*grpc_catalog_common_go.OpResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(grpc_catalog_common_go.OpResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *catalogClient) Download(ctx context.Context, in *DownloadApplicationRequest, opts ...grpc.CallOption) (Catalog_DownloadClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Catalog_serviceDesc.Streams[1], "/catalog.Catalog/Download", opts...)
	if err != nil {
		return nil, err
	}
	x := &catalogDownloadClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Catalog_DownloadClient interface {
	Recv() (*FileInfo, error)
	grpc.ClientStream
}

type catalogDownloadClient struct {
	grpc.ClientStream
}

func (x *catalogDownloadClient) Recv() (*FileInfo, error) {
	m := new(FileInfo)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *catalogClient) Remove(ctx context.Context, in *RemoveApplicationRequest, opts ...grpc.CallOption) (*grpc_catalog_common_go.OpResponse, error) {
	out := new(grpc_catalog_common_go.OpResponse)
	err := c.cc.Invoke(ctx, "/catalog.Catalog/Remove", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogClient) List(ctx context.Context, in *ListApplicationsRequest, opts ...grpc.CallOption) (*ApplicationList, error) {
	out := new(ApplicationList)
	err := c.cc.Invoke(ctx, "/catalog.Catalog/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogClient) Info(ctx context.Context, in *InfoApplicationRequest, opts ...grpc.CallOption) (*InfoApplicationResponse, error) {
	out := new(InfoApplicationResponse)
	err := c.cc.Invoke(ctx, "/catalog.Catalog/Info", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogClient) Summary(ctx context.Context, in *grpc_catalog_common_go.EmptyRequest, opts ...grpc.CallOption) (*SummaryResponse, error) {
	out := new(SummaryResponse)
	err := c.cc.Invoke(ctx, "/catalog.Catalog/Summary", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CatalogServer is the server API for Catalog service.
type CatalogServer interface {
	// Add a new application to the catalog.
	Add(Catalog_AddServer) error
	// Download an application from the catalog.
	Download(*DownloadApplicationRequest, Catalog_DownloadServer) error
	//Remove an application from the catalog.
	Remove(context.Context, *RemoveApplicationRequest) (*grpc_catalog_common_go.OpResponse, error)
	// List the applications in the catalog.
	List(context.Context, *ListApplicationsRequest) (*ApplicationList, error)
	// Info returns the detail of a given application.
	Info(context.Context, *InfoApplicationRequest) (*InfoApplicationResponse, error)
	// Summary returns the catalog summary
	Summary(context.Context, *grpc_catalog_common_go.EmptyRequest) (*SummaryResponse, error)
}

// UnimplementedCatalogServer can be embedded to have forward compatible implementations.
type UnimplementedCatalogServer struct {
}

func (*UnimplementedCatalogServer) Add(Catalog_AddServer) error {
	return status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (*UnimplementedCatalogServer) Download(*DownloadApplicationRequest, Catalog_DownloadServer) error {
	return status.Errorf(codes.Unimplemented, "method Download not implemented")
}
func (*UnimplementedCatalogServer) Remove(context.Context, *RemoveApplicationRequest) (*grpc_catalog_common_go.OpResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Remove not implemented")
}
func (*UnimplementedCatalogServer) List(context.Context, *ListApplicationsRequest) (*ApplicationList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (*UnimplementedCatalogServer) Info(context.Context, *InfoApplicationRequest) (*InfoApplicationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Info not implemented")
}
func (*UnimplementedCatalogServer) Summary(context.Context, *grpc_catalog_common_go.EmptyRequest) (*SummaryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Summary not implemented")
}

func RegisterCatalogServer(s *grpc.Server, srv CatalogServer) {
	s.RegisterService(&_Catalog_serviceDesc, srv)
}

func _Catalog_Add_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CatalogServer).Add(&catalogAddServer{stream})
}

type Catalog_AddServer interface {
	SendAndClose(*grpc_catalog_common_go.OpResponse) error
	Recv() (*AddApplicationRequest, error)
	grpc.ServerStream
}

type catalogAddServer struct {
	grpc.ServerStream
}

func (x *catalogAddServer) SendAndClose(m *grpc_catalog_common_go.OpResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *catalogAddServer) Recv() (*AddApplicationRequest, error) {
	m := new(AddApplicationRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Catalog_Download_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(DownloadApplicationRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CatalogServer).Download(m, &catalogDownloadServer{stream})
}

type Catalog_DownloadServer interface {
	Send(*FileInfo) error
	grpc.ServerStream
}

type catalogDownloadServer struct {
	grpc.ServerStream
}

func (x *catalogDownloadServer) Send(m *FileInfo) error {
	return x.ServerStream.SendMsg(m)
}

func _Catalog_Remove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveApplicationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServer).Remove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/catalog.Catalog/Remove",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServer).Remove(ctx, req.(*RemoveApplicationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Catalog_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListApplicationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/catalog.Catalog/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServer).List(ctx, req.(*ListApplicationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Catalog_Info_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InfoApplicationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServer).Info(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/catalog.Catalog/Info",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServer).Info(ctx, req.(*InfoApplicationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Catalog_Summary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(grpc_catalog_common_go.EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServer).Summary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/catalog.Catalog/Summary",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServer).Summary(ctx, req.(*grpc_catalog_common_go.EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Catalog_serviceDesc = grpc.ServiceDesc{
	ServiceName: "catalog.Catalog",
	HandlerType: (*CatalogServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Remove",
			Handler:    _Catalog_Remove_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Catalog_List_Handler,
		},
		{
			MethodName: "Info",
			Handler:    _Catalog_Info_Handler,
		},
		{
			MethodName: "Summary",
			Handler:    _Catalog_Summary_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Add",
			Handler:       _Catalog_Add_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Download",
			Handler:       _Catalog_Download_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "catalog/services.proto",
}

// NamespaceAdministrationClient is the client API for NamespaceAdministration service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NamespaceAdministrationClient interface {
	// Delete a namespace so that the applications contained on it are not longer available.
	Delete(ctx context.Context, in *DeleteNamespaceRequest, opts ...grpc.CallOption) (*grpc_catalog_common_go.OpResponse, error)
	// DeleteApplication deletes an application
	DeleteApplication(ctx context.Context, in *RemoveApplicationRequest, opts ...grpc.CallOption) (*grpc_catalog_common_go.OpResponse, error)
	// List all the applications in the catalog
	List(ctx context.Context, in *ListApplicationsRequest, opts ...grpc.CallOption) (*ApplicationList, error)
}

type namespaceAdministrationClient struct {
	cc grpc.ClientConnInterface
}

func NewNamespaceAdministrationClient(cc grpc.ClientConnInterface) NamespaceAdministrationClient {
	return &namespaceAdministrationClient{cc}
}

func (c *namespaceAdministrationClient) Delete(ctx context.Context, in *DeleteNamespaceRequest, opts ...grpc.CallOption) (*grpc_catalog_common_go.OpResponse, error) {
	out := new(grpc_catalog_common_go.OpResponse)
	err := c.cc.Invoke(ctx, "/catalog.NamespaceAdministration/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *namespaceAdministrationClient) DeleteApplication(ctx context.Context, in *RemoveApplicationRequest, opts ...grpc.CallOption) (*grpc_catalog_common_go.OpResponse, error) {
	out := new(grpc_catalog_common_go.OpResponse)
	err := c.cc.Invoke(ctx, "/catalog.NamespaceAdministration/DeleteApplication", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *namespaceAdministrationClient) List(ctx context.Context, in *ListApplicationsRequest, opts ...grpc.CallOption) (*ApplicationList, error) {
	out := new(ApplicationList)
	err := c.cc.Invoke(ctx, "/catalog.NamespaceAdministration/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NamespaceAdministrationServer is the server API for NamespaceAdministration service.
type NamespaceAdministrationServer interface {
	// Delete a namespace so that the applications contained on it are not longer available.
	Delete(context.Context, *DeleteNamespaceRequest) (*grpc_catalog_common_go.OpResponse, error)
	// DeleteApplication deletes an application
	DeleteApplication(context.Context, *RemoveApplicationRequest) (*grpc_catalog_common_go.OpResponse, error)
	// List all the applications in the catalog
	List(context.Context, *ListApplicationsRequest) (*ApplicationList, error)
}

// UnimplementedNamespaceAdministrationServer can be embedded to have forward compatible implementations.
type UnimplementedNamespaceAdministrationServer struct {
}

func (*UnimplementedNamespaceAdministrationServer) Delete(context.Context, *DeleteNamespaceRequest) (*grpc_catalog_common_go.OpResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (*UnimplementedNamespaceAdministrationServer) DeleteApplication(context.Context, *RemoveApplicationRequest) (*grpc_catalog_common_go.OpResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteApplication not implemented")
}
func (*UnimplementedNamespaceAdministrationServer) List(context.Context, *ListApplicationsRequest) (*ApplicationList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}

func RegisterNamespaceAdministrationServer(s *grpc.Server, srv NamespaceAdministrationServer) {
	s.RegisterService(&_NamespaceAdministration_serviceDesc, srv)
}

func _NamespaceAdministration_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteNamespaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NamespaceAdministrationServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/catalog.NamespaceAdministration/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NamespaceAdministrationServer).Delete(ctx, req.(*DeleteNamespaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NamespaceAdministration_DeleteApplication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveApplicationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NamespaceAdministrationServer).DeleteApplication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/catalog.NamespaceAdministration/DeleteApplication",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NamespaceAdministrationServer).DeleteApplication(ctx, req.(*RemoveApplicationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NamespaceAdministration_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListApplicationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NamespaceAdministrationServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/catalog.NamespaceAdministration/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NamespaceAdministrationServer).List(ctx, req.(*ListApplicationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _NamespaceAdministration_serviceDesc = grpc.ServiceDesc{
	ServiceName: "catalog.NamespaceAdministration",
	HandlerType: (*NamespaceAdministrationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Delete",
			Handler:    _NamespaceAdministration_Delete_Handler,
		},
		{
			MethodName: "DeleteApplication",
			Handler:    _NamespaceAdministration_DeleteApplication_Handler,
		},
		{
			MethodName: "List",
			Handler:    _NamespaceAdministration_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "catalog/services.proto",
}

// ApplicationsClient is the client API for Applications service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ApplicationsClient interface {
	// Deploy an application on a target Playground platform. This endpoint
	// will gather the application information and send it to the target
	// playground platform.
	Deploy(ctx context.Context, in *DeployApplicationRequest, opts ...grpc.CallOption) (*grpc_catalog_common_go.OpResponse, error)
	GetConfiguration(ctx context.Context, in *GetConfigurationRequest, opts ...grpc.CallOption) (*GetConfigurationResponse, error)
}

type applicationsClient struct {
	cc grpc.ClientConnInterface
}

func NewApplicationsClient(cc grpc.ClientConnInterface) ApplicationsClient {
	return &applicationsClient{cc}
}

func (c *applicationsClient) Deploy(ctx context.Context, in *DeployApplicationRequest, opts ...grpc.CallOption) (*grpc_catalog_common_go.OpResponse, error) {
	out := new(grpc_catalog_common_go.OpResponse)
	err := c.cc.Invoke(ctx, "/catalog.Applications/Deploy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationsClient) GetConfiguration(ctx context.Context, in *GetConfigurationRequest, opts ...grpc.CallOption) (*GetConfigurationResponse, error) {
	out := new(GetConfigurationResponse)
	err := c.cc.Invoke(ctx, "/catalog.Applications/GetConfiguration", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ApplicationsServer is the server API for Applications service.
type ApplicationsServer interface {
	// Deploy an application on a target Playground platform. This endpoint
	// will gather the application information and send it to the target
	// playground platform.
	Deploy(context.Context, *DeployApplicationRequest) (*grpc_catalog_common_go.OpResponse, error)
	GetConfiguration(context.Context, *GetConfigurationRequest) (*GetConfigurationResponse, error)
}

// UnimplementedApplicationsServer can be embedded to have forward compatible implementations.
type UnimplementedApplicationsServer struct {
}

func (*UnimplementedApplicationsServer) Deploy(context.Context, *DeployApplicationRequest) (*grpc_catalog_common_go.OpResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Deploy not implemented")
}
func (*UnimplementedApplicationsServer) GetConfiguration(context.Context, *GetConfigurationRequest) (*GetConfigurationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConfiguration not implemented")
}

func RegisterApplicationsServer(s *grpc.Server, srv ApplicationsServer) {
	s.RegisterService(&_Applications_serviceDesc, srv)
}

func _Applications_Deploy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeployApplicationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationsServer).Deploy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/catalog.Applications/Deploy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationsServer).Deploy(ctx, req.(*DeployApplicationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Applications_GetConfiguration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetConfigurationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationsServer).GetConfiguration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/catalog.Applications/GetConfiguration",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationsServer).GetConfiguration(ctx, req.(*GetConfigurationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Applications_serviceDesc = grpc.ServiceDesc{
	ServiceName: "catalog.Applications",
	HandlerType: (*ApplicationsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Deploy",
			Handler:    _Applications_Deploy_Handler,
		},
		{
			MethodName: "GetConfiguration",
			Handler:    _Applications_GetConfiguration_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "catalog/services.proto",
}
