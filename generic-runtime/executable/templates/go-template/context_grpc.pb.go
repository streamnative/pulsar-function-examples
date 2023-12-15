//*
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.9
// source: context.proto

package main

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	ContextService_Publish_FullMethodName       = "/context.ContextService/publish"
	ContextService_CurrentRecord_FullMethodName = "/context.ContextService/currentRecord"
	ContextService_RecordMetrics_FullMethodName = "/context.ContextService/recordMetrics"
	ContextService_Seek_FullMethodName          = "/context.ContextService/seek"
	ContextService_Pause_FullMethodName         = "/context.ContextService/pause"
	ContextService_Resume_FullMethodName        = "/context.ContextService/resume"
	ContextService_GetState_FullMethodName      = "/context.ContextService/getState"
	ContextService_PutState_FullMethodName      = "/context.ContextService/putState"
	ContextService_DeleteState_FullMethodName   = "/context.ContextService/deleteState"
	ContextService_GetCounter_FullMethodName    = "/context.ContextService/getCounter"
	ContextService_IncrCounter_FullMethodName   = "/context.ContextService/incrCounter"
)

// ContextServiceClient is the client API for ContextService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ContextServiceClient interface {
	Publish(ctx context.Context, in *PulsarMessage, opts ...grpc.CallOption) (*emptypb.Empty, error)
	CurrentRecord(ctx context.Context, in *MessageId, opts ...grpc.CallOption) (*Record, error)
	RecordMetrics(ctx context.Context, in *MetricData, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Seek(ctx context.Context, in *Partition, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Pause(ctx context.Context, in *Partition, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Resume(ctx context.Context, in *Partition, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetState(ctx context.Context, in *StateKey, opts ...grpc.CallOption) (*StateResult, error)
	PutState(ctx context.Context, in *StateKeyValue, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteState(ctx context.Context, in *StateKey, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetCounter(ctx context.Context, in *StateKey, opts ...grpc.CallOption) (*Counter, error)
	IncrCounter(ctx context.Context, in *IncrStateKey, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type contextServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewContextServiceClient(cc grpc.ClientConnInterface) ContextServiceClient {
	return &contextServiceClient{cc}
}

func (c *contextServiceClient) Publish(ctx context.Context, in *PulsarMessage, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, ContextService_Publish_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contextServiceClient) CurrentRecord(ctx context.Context, in *MessageId, opts ...grpc.CallOption) (*Record, error) {
	out := new(Record)
	err := c.cc.Invoke(ctx, ContextService_CurrentRecord_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contextServiceClient) RecordMetrics(ctx context.Context, in *MetricData, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, ContextService_RecordMetrics_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contextServiceClient) Seek(ctx context.Context, in *Partition, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, ContextService_Seek_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contextServiceClient) Pause(ctx context.Context, in *Partition, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, ContextService_Pause_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contextServiceClient) Resume(ctx context.Context, in *Partition, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, ContextService_Resume_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contextServiceClient) GetState(ctx context.Context, in *StateKey, opts ...grpc.CallOption) (*StateResult, error) {
	out := new(StateResult)
	err := c.cc.Invoke(ctx, ContextService_GetState_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contextServiceClient) PutState(ctx context.Context, in *StateKeyValue, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, ContextService_PutState_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contextServiceClient) DeleteState(ctx context.Context, in *StateKey, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, ContextService_DeleteState_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contextServiceClient) GetCounter(ctx context.Context, in *StateKey, opts ...grpc.CallOption) (*Counter, error) {
	out := new(Counter)
	err := c.cc.Invoke(ctx, ContextService_GetCounter_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contextServiceClient) IncrCounter(ctx context.Context, in *IncrStateKey, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, ContextService_IncrCounter_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ContextServiceServer is the server API for ContextService service.
// All implementations should embed UnimplementedContextServiceServer
// for forward compatibility
type ContextServiceServer interface {
	Publish(context.Context, *PulsarMessage) (*emptypb.Empty, error)
	CurrentRecord(context.Context, *MessageId) (*Record, error)
	RecordMetrics(context.Context, *MetricData) (*emptypb.Empty, error)
	Seek(context.Context, *Partition) (*emptypb.Empty, error)
	Pause(context.Context, *Partition) (*emptypb.Empty, error)
	Resume(context.Context, *Partition) (*emptypb.Empty, error)
	GetState(context.Context, *StateKey) (*StateResult, error)
	PutState(context.Context, *StateKeyValue) (*emptypb.Empty, error)
	DeleteState(context.Context, *StateKey) (*emptypb.Empty, error)
	GetCounter(context.Context, *StateKey) (*Counter, error)
	IncrCounter(context.Context, *IncrStateKey) (*emptypb.Empty, error)
}

// UnimplementedContextServiceServer should be embedded to have forward compatible implementations.
type UnimplementedContextServiceServer struct {
}

func (UnimplementedContextServiceServer) Publish(context.Context, *PulsarMessage) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Publish not implemented")
}
func (UnimplementedContextServiceServer) CurrentRecord(context.Context, *MessageId) (*Record, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CurrentRecord not implemented")
}
func (UnimplementedContextServiceServer) RecordMetrics(context.Context, *MetricData) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RecordMetrics not implemented")
}
func (UnimplementedContextServiceServer) Seek(context.Context, *Partition) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Seek not implemented")
}
func (UnimplementedContextServiceServer) Pause(context.Context, *Partition) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Pause not implemented")
}
func (UnimplementedContextServiceServer) Resume(context.Context, *Partition) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Resume not implemented")
}
func (UnimplementedContextServiceServer) GetState(context.Context, *StateKey) (*StateResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetState not implemented")
}
func (UnimplementedContextServiceServer) PutState(context.Context, *StateKeyValue) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutState not implemented")
}
func (UnimplementedContextServiceServer) DeleteState(context.Context, *StateKey) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteState not implemented")
}
func (UnimplementedContextServiceServer) GetCounter(context.Context, *StateKey) (*Counter, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCounter not implemented")
}
func (UnimplementedContextServiceServer) IncrCounter(context.Context, *IncrStateKey) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IncrCounter not implemented")
}

// UnsafeContextServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ContextServiceServer will
// result in compilation errors.
type UnsafeContextServiceServer interface {
	mustEmbedUnimplementedContextServiceServer()
}

func RegisterContextServiceServer(s grpc.ServiceRegistrar, srv ContextServiceServer) {
	s.RegisterService(&ContextService_ServiceDesc, srv)
}

func _ContextService_Publish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PulsarMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContextServiceServer).Publish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ContextService_Publish_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContextServiceServer).Publish(ctx, req.(*PulsarMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContextService_CurrentRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContextServiceServer).CurrentRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ContextService_CurrentRecord_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContextServiceServer).CurrentRecord(ctx, req.(*MessageId))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContextService_RecordMetrics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MetricData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContextServiceServer).RecordMetrics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ContextService_RecordMetrics_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContextServiceServer).RecordMetrics(ctx, req.(*MetricData))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContextService_Seek_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Partition)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContextServiceServer).Seek(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ContextService_Seek_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContextServiceServer).Seek(ctx, req.(*Partition))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContextService_Pause_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Partition)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContextServiceServer).Pause(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ContextService_Pause_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContextServiceServer).Pause(ctx, req.(*Partition))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContextService_Resume_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Partition)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContextServiceServer).Resume(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ContextService_Resume_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContextServiceServer).Resume(ctx, req.(*Partition))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContextService_GetState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StateKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContextServiceServer).GetState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ContextService_GetState_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContextServiceServer).GetState(ctx, req.(*StateKey))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContextService_PutState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StateKeyValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContextServiceServer).PutState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ContextService_PutState_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContextServiceServer).PutState(ctx, req.(*StateKeyValue))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContextService_DeleteState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StateKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContextServiceServer).DeleteState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ContextService_DeleteState_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContextServiceServer).DeleteState(ctx, req.(*StateKey))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContextService_GetCounter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StateKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContextServiceServer).GetCounter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ContextService_GetCounter_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContextServiceServer).GetCounter(ctx, req.(*StateKey))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContextService_IncrCounter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IncrStateKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContextServiceServer).IncrCounter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ContextService_IncrCounter_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContextServiceServer).IncrCounter(ctx, req.(*IncrStateKey))
	}
	return interceptor(ctx, in, info, handler)
}

// ContextService_ServiceDesc is the grpc.ServiceDesc for ContextService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ContextService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "context.ContextService",
	HandlerType: (*ContextServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "publish",
			Handler:    _ContextService_Publish_Handler,
		},
		{
			MethodName: "currentRecord",
			Handler:    _ContextService_CurrentRecord_Handler,
		},
		{
			MethodName: "recordMetrics",
			Handler:    _ContextService_RecordMetrics_Handler,
		},
		{
			MethodName: "seek",
			Handler:    _ContextService_Seek_Handler,
		},
		{
			MethodName: "pause",
			Handler:    _ContextService_Pause_Handler,
		},
		{
			MethodName: "resume",
			Handler:    _ContextService_Resume_Handler,
		},
		{
			MethodName: "getState",
			Handler:    _ContextService_GetState_Handler,
		},
		{
			MethodName: "putState",
			Handler:    _ContextService_PutState_Handler,
		},
		{
			MethodName: "deleteState",
			Handler:    _ContextService_DeleteState_Handler,
		},
		{
			MethodName: "getCounter",
			Handler:    _ContextService_GetCounter_Handler,
		},
		{
			MethodName: "incrCounter",
			Handler:    _ContextService_IncrCounter_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "context.proto",
}
