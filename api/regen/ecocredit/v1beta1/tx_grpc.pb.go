// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: regen/ecocredit/v1beta1/tx.proto

package ecocreditv1beta1

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

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MsgClient interface {
	// CreateClass creates a new credit class with an approved list of issuers and
	// optional metadata.
	CreateClass(ctx context.Context, in *MsgCreateClass, opts ...grpc.CallOption) (*MsgCreateClassResponse, error)
	// CreateProject creates a new project within a credit class.
	CreateProject(ctx context.Context, in *MsgCreateProject, opts ...grpc.CallOption) (*MsgCreateProjectResponse, error)
	// CreateBatch creates a new batch of credits for an existing project.
	// This will create a new batch denom with a fixed supply. Issued credits can
	// be distributed to recipients in either tradable or retired form.
	CreateBatch(ctx context.Context, in *MsgCreateBatch, opts ...grpc.CallOption) (*MsgCreateBatchResponse, error)
	// Send sends tradable credits from one account to another account. Sent
	// credits can either be tradable or retired on receipt.
	Send(ctx context.Context, in *MsgSend, opts ...grpc.CallOption) (*MsgSendResponse, error)
	// Retire retires a specified number of credits in the holder's account.
	Retire(ctx context.Context, in *MsgRetire, opts ...grpc.CallOption) (*MsgRetireResponse, error)
	// Cancel removes a number of credits from the holder's account and also
	// deducts them from the tradable supply, effectively cancelling their
	// issuance on Regen Ledger
	Cancel(ctx context.Context, in *MsgCancel, opts ...grpc.CallOption) (*MsgCancelResponse, error)
	// UpdateClassAdmin updates the credit class admin
	UpdateClassAdmin(ctx context.Context, in *MsgUpdateClassAdmin, opts ...grpc.CallOption) (*MsgUpdateClassAdminResponse, error)
	// UpdateClassIssuers updates the credit class issuer list
	UpdateClassIssuers(ctx context.Context, in *MsgUpdateClassIssuers, opts ...grpc.CallOption) (*MsgUpdateClassIssuersResponse, error)
	// UpdateClassMetadata updates the credit class metadata
	UpdateClassMetadata(ctx context.Context, in *MsgUpdateClassMetadata, opts ...grpc.CallOption) (*MsgUpdateClassMetadataResponse, error)
}

type msgClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgClient(cc grpc.ClientConnInterface) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) CreateClass(ctx context.Context, in *MsgCreateClass, opts ...grpc.CallOption) (*MsgCreateClassResponse, error) {
	out := new(MsgCreateClassResponse)
	err := c.cc.Invoke(ctx, "/regen.ecocredit.v1beta1.Msg/CreateClass", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) CreateProject(ctx context.Context, in *MsgCreateProject, opts ...grpc.CallOption) (*MsgCreateProjectResponse, error) {
	out := new(MsgCreateProjectResponse)
	err := c.cc.Invoke(ctx, "/regen.ecocredit.v1beta1.Msg/CreateProject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) CreateBatch(ctx context.Context, in *MsgCreateBatch, opts ...grpc.CallOption) (*MsgCreateBatchResponse, error) {
	out := new(MsgCreateBatchResponse)
	err := c.cc.Invoke(ctx, "/regen.ecocredit.v1beta1.Msg/CreateBatch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) Send(ctx context.Context, in *MsgSend, opts ...grpc.CallOption) (*MsgSendResponse, error) {
	out := new(MsgSendResponse)
	err := c.cc.Invoke(ctx, "/regen.ecocredit.v1beta1.Msg/Send", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) Retire(ctx context.Context, in *MsgRetire, opts ...grpc.CallOption) (*MsgRetireResponse, error) {
	out := new(MsgRetireResponse)
	err := c.cc.Invoke(ctx, "/regen.ecocredit.v1beta1.Msg/Retire", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) Cancel(ctx context.Context, in *MsgCancel, opts ...grpc.CallOption) (*MsgCancelResponse, error) {
	out := new(MsgCancelResponse)
	err := c.cc.Invoke(ctx, "/regen.ecocredit.v1beta1.Msg/Cancel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateClassAdmin(ctx context.Context, in *MsgUpdateClassAdmin, opts ...grpc.CallOption) (*MsgUpdateClassAdminResponse, error) {
	out := new(MsgUpdateClassAdminResponse)
	err := c.cc.Invoke(ctx, "/regen.ecocredit.v1beta1.Msg/UpdateClassAdmin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateClassIssuers(ctx context.Context, in *MsgUpdateClassIssuers, opts ...grpc.CallOption) (*MsgUpdateClassIssuersResponse, error) {
	out := new(MsgUpdateClassIssuersResponse)
	err := c.cc.Invoke(ctx, "/regen.ecocredit.v1beta1.Msg/UpdateClassIssuers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateClassMetadata(ctx context.Context, in *MsgUpdateClassMetadata, opts ...grpc.CallOption) (*MsgUpdateClassMetadataResponse, error) {
	out := new(MsgUpdateClassMetadataResponse)
	err := c.cc.Invoke(ctx, "/regen.ecocredit.v1beta1.Msg/UpdateClassMetadata", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
// All implementations must embed UnimplementedMsgServer
// for forward compatibility
type MsgServer interface {
	// CreateClass creates a new credit class with an approved list of issuers and
	// optional metadata.
	CreateClass(context.Context, *MsgCreateClass) (*MsgCreateClassResponse, error)
	// CreateProject creates a new project within a credit class.
	CreateProject(context.Context, *MsgCreateProject) (*MsgCreateProjectResponse, error)
	// CreateBatch creates a new batch of credits for an existing project.
	// This will create a new batch denom with a fixed supply. Issued credits can
	// be distributed to recipients in either tradable or retired form.
	CreateBatch(context.Context, *MsgCreateBatch) (*MsgCreateBatchResponse, error)
	// Send sends tradable credits from one account to another account. Sent
	// credits can either be tradable or retired on receipt.
	Send(context.Context, *MsgSend) (*MsgSendResponse, error)
	// Retire retires a specified number of credits in the holder's account.
	Retire(context.Context, *MsgRetire) (*MsgRetireResponse, error)
	// Cancel removes a number of credits from the holder's account and also
	// deducts them from the tradable supply, effectively cancelling their
	// issuance on Regen Ledger
	Cancel(context.Context, *MsgCancel) (*MsgCancelResponse, error)
	// UpdateClassAdmin updates the credit class admin
	UpdateClassAdmin(context.Context, *MsgUpdateClassAdmin) (*MsgUpdateClassAdminResponse, error)
	// UpdateClassIssuers updates the credit class issuer list
	UpdateClassIssuers(context.Context, *MsgUpdateClassIssuers) (*MsgUpdateClassIssuersResponse, error)
	// UpdateClassMetadata updates the credit class metadata
	UpdateClassMetadata(context.Context, *MsgUpdateClassMetadata) (*MsgUpdateClassMetadataResponse, error)
	mustEmbedUnimplementedMsgServer()
}

// UnimplementedMsgServer must be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (UnimplementedMsgServer) CreateClass(context.Context, *MsgCreateClass) (*MsgCreateClassResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateClass not implemented")
}
func (UnimplementedMsgServer) CreateProject(context.Context, *MsgCreateProject) (*MsgCreateProjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProject not implemented")
}
func (UnimplementedMsgServer) CreateBatch(context.Context, *MsgCreateBatch) (*MsgCreateBatchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBatch not implemented")
}
func (UnimplementedMsgServer) Send(context.Context, *MsgSend) (*MsgSendResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Send not implemented")
}
func (UnimplementedMsgServer) Retire(context.Context, *MsgRetire) (*MsgRetireResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Retire not implemented")
}
func (UnimplementedMsgServer) Cancel(context.Context, *MsgCancel) (*MsgCancelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Cancel not implemented")
}
func (UnimplementedMsgServer) UpdateClassAdmin(context.Context, *MsgUpdateClassAdmin) (*MsgUpdateClassAdminResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateClassAdmin not implemented")
}
func (UnimplementedMsgServer) UpdateClassIssuers(context.Context, *MsgUpdateClassIssuers) (*MsgUpdateClassIssuersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateClassIssuers not implemented")
}
func (UnimplementedMsgServer) UpdateClassMetadata(context.Context, *MsgUpdateClassMetadata) (*MsgUpdateClassMetadataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateClassMetadata not implemented")
}
func (UnimplementedMsgServer) mustEmbedUnimplementedMsgServer() {}

// UnsafeMsgServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MsgServer will
// result in compilation errors.
type UnsafeMsgServer interface {
	mustEmbedUnimplementedMsgServer()
}

func RegisterMsgServer(s grpc.ServiceRegistrar, srv MsgServer) {
	s.RegisterService(&Msg_ServiceDesc, srv)
}

func _Msg_CreateClass_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateClass)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateClass(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/regen.ecocredit.v1beta1.Msg/CreateClass",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateClass(ctx, req.(*MsgCreateClass))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_CreateProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateProject)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/regen.ecocredit.v1beta1.Msg/CreateProject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateProject(ctx, req.(*MsgCreateProject))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_CreateBatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateBatch)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateBatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/regen.ecocredit.v1beta1.Msg/CreateBatch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateBatch(ctx, req.(*MsgCreateBatch))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSend)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/regen.ecocredit.v1beta1.Msg/Send",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Send(ctx, req.(*MsgSend))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_Retire_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRetire)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Retire(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/regen.ecocredit.v1beta1.Msg/Retire",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Retire(ctx, req.(*MsgRetire))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_Cancel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCancel)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Cancel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/regen.ecocredit.v1beta1.Msg/Cancel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Cancel(ctx, req.(*MsgCancel))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateClassAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateClassAdmin)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateClassAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/regen.ecocredit.v1beta1.Msg/UpdateClassAdmin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateClassAdmin(ctx, req.(*MsgUpdateClassAdmin))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateClassIssuers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateClassIssuers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateClassIssuers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/regen.ecocredit.v1beta1.Msg/UpdateClassIssuers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateClassIssuers(ctx, req.(*MsgUpdateClassIssuers))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateClassMetadata_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateClassMetadata)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateClassMetadata(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/regen.ecocredit.v1beta1.Msg/UpdateClassMetadata",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateClassMetadata(ctx, req.(*MsgUpdateClassMetadata))
	}
	return interceptor(ctx, in, info, handler)
}

// Msg_ServiceDesc is the grpc.ServiceDesc for Msg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "regen.ecocredit.v1beta1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateClass",
			Handler:    _Msg_CreateClass_Handler,
		},
		{
			MethodName: "CreateProject",
			Handler:    _Msg_CreateProject_Handler,
		},
		{
			MethodName: "CreateBatch",
			Handler:    _Msg_CreateBatch_Handler,
		},
		{
			MethodName: "Send",
			Handler:    _Msg_Send_Handler,
		},
		{
			MethodName: "Retire",
			Handler:    _Msg_Retire_Handler,
		},
		{
			MethodName: "Cancel",
			Handler:    _Msg_Cancel_Handler,
		},
		{
			MethodName: "UpdateClassAdmin",
			Handler:    _Msg_UpdateClassAdmin_Handler,
		},
		{
			MethodName: "UpdateClassIssuers",
			Handler:    _Msg_UpdateClassIssuers_Handler,
		},
		{
			MethodName: "UpdateClassMetadata",
			Handler:    _Msg_UpdateClassMetadata_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "regen/ecocredit/v1beta1/tx.proto",
}
