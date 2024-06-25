//*
// This file is auto generated by tto v0.4.5 !
// If you want to modify this code, please read the guide
// to modify code template.
//
// Get started: https://github.com/ixre/tto
//
// Copyright (C) 2021 <no value>, All rights reserved.
//
// name : template_service.proto
// author : jarrysix
// date : 2021/12/02 10:37:45
// description : 条码服务
// history :

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.27.0
// source: code_service.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	CodeService_SaveQrTemplate_FullMethodName   = "/CodeService/SaveQrTemplate"
	CodeService_GetQrTemplate_FullMethodName    = "/CodeService/GetQrTemplate"
	CodeService_DeleteQrTemplate_FullMethodName = "/CodeService/DeleteQrTemplate"
)

// CodeServiceClient is the client API for CodeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CodeServiceClient interface {
	// 保存CommQrTemplate
	SaveQrTemplate(ctx context.Context, in *SaveQrTemplateRequest, opts ...grpc.CallOption) (*SaveQrTemplateResponse, error)
	// 获取CommQrTemplate
	GetQrTemplate(ctx context.Context, in *CommQrTemplateId, opts ...grpc.CallOption) (*SQrTemplate, error)
	// 删除CommQrTemplate
	DeleteQrTemplate(ctx context.Context, in *CommQrTemplateId, opts ...grpc.CallOption) (*Result, error)
}

type codeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCodeServiceClient(cc grpc.ClientConnInterface) CodeServiceClient {
	return &codeServiceClient{cc}
}

func (c *codeServiceClient) SaveQrTemplate(ctx context.Context, in *SaveQrTemplateRequest, opts ...grpc.CallOption) (*SaveQrTemplateResponse, error) {
	out := new(SaveQrTemplateResponse)
	err := c.cc.Invoke(ctx, CodeService_SaveQrTemplate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *codeServiceClient) GetQrTemplate(ctx context.Context, in *CommQrTemplateId, opts ...grpc.CallOption) (*SQrTemplate, error) {
	out := new(SQrTemplate)
	err := c.cc.Invoke(ctx, CodeService_GetQrTemplate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *codeServiceClient) DeleteQrTemplate(ctx context.Context, in *CommQrTemplateId, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, CodeService_DeleteQrTemplate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CodeServiceServer is the server API for CodeService service.
// All implementations must embed UnimplementedCodeServiceServer
// for forward compatibility
type CodeServiceServer interface {
	// 保存CommQrTemplate
	SaveQrTemplate(context.Context, *SaveQrTemplateRequest) (*SaveQrTemplateResponse, error)
	// 获取CommQrTemplate
	GetQrTemplate(context.Context, *CommQrTemplateId) (*SQrTemplate, error)
	// 删除CommQrTemplate
	DeleteQrTemplate(context.Context, *CommQrTemplateId) (*Result, error)
	mustEmbedUnimplementedCodeServiceServer()
}

// UnimplementedCodeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCodeServiceServer struct {
}

func (UnimplementedCodeServiceServer) SaveQrTemplate(context.Context, *SaveQrTemplateRequest) (*SaveQrTemplateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveQrTemplate not implemented")
}
func (UnimplementedCodeServiceServer) GetQrTemplate(context.Context, *CommQrTemplateId) (*SQrTemplate, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetQrTemplate not implemented")
}
func (UnimplementedCodeServiceServer) DeleteQrTemplate(context.Context, *CommQrTemplateId) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteQrTemplate not implemented")
}
func (UnimplementedCodeServiceServer) mustEmbedUnimplementedCodeServiceServer() {}

// UnsafeCodeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CodeServiceServer will
// result in compilation errors.
type UnsafeCodeServiceServer interface {
	mustEmbedUnimplementedCodeServiceServer()
}

func RegisterCodeServiceServer(s grpc.ServiceRegistrar, srv CodeServiceServer) {
	s.RegisterService(&CodeService_ServiceDesc, srv)
}

func _CodeService_SaveQrTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveQrTemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CodeServiceServer).SaveQrTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CodeService_SaveQrTemplate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CodeServiceServer).SaveQrTemplate(ctx, req.(*SaveQrTemplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CodeService_GetQrTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommQrTemplateId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CodeServiceServer).GetQrTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CodeService_GetQrTemplate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CodeServiceServer).GetQrTemplate(ctx, req.(*CommQrTemplateId))
	}
	return interceptor(ctx, in, info, handler)
}

func _CodeService_DeleteQrTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommQrTemplateId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CodeServiceServer).DeleteQrTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CodeService_DeleteQrTemplate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CodeServiceServer).DeleteQrTemplate(ctx, req.(*CommQrTemplateId))
	}
	return interceptor(ctx, in, info, handler)
}

// CodeService_ServiceDesc is the grpc.ServiceDesc for CodeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CodeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "CodeService",
	HandlerType: (*CodeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SaveQrTemplate",
			Handler:    _CodeService_SaveQrTemplate_Handler,
		},
		{
			MethodName: "GetQrTemplate",
			Handler:    _CodeService_GetQrTemplate_Handler,
		},
		{
			MethodName: "DeleteQrTemplate",
			Handler:    _CodeService_DeleteQrTemplate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "code_service.proto",
}
