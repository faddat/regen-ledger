// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: regen/ecocredit/basket/v1beta1/tx.proto

package basketv1beta1

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
	// CreateBasket creates a bank denom which wraps credits.
	CreateBasket(ctx context.Context, in *MsgCreateBasket, opts ...grpc.CallOption) (*MsgCreateBasketResponse, error)
	// AddToBasket adds credits to a basket in return for basket tokens.
	AddToBasket(ctx context.Context, in *MsgAddToBasket, opts ...grpc.CallOption) (*MsgAddToBasketResponse, error)
	// TakeFromBasket takes credits from a basket without regard for which
	// credits they are. The credits will be auto-retired if disable_auto_retire
	// is false. Credits will be chosen randomly using the previous block hash
	// as a consensus source of randomness.
	// More concretely, the implementation is as follows:
	// - take the previous block hash and convert it into an uint64,
	// - given the total number of different credits within the basket `n`, the
	//   first credits that will get picked correspond to: hash modulo n (in
	//   terms of order),
	// - then if we need to take more credits, we get some from the next one and
	//   so on.
	TakeFromBasket(ctx context.Context, in *MsgTakeFromBasket, opts ...grpc.CallOption) (*MsgTakeFromBasketResponse, error)
	// PickFromBasket picks specific credits from a basket. If allow_picking is
	// set to false, then only an address which deposited credits in the basket
	// can pick those credits. All other addresses will be blocked from picking
	// those credits. The credits will be auto-retired if disable_auto_retire is
	// false unless the credits were previously put into the basket by the
	// address picking them from the basket, in which case they will remain
	// tradable. This functionality allows the owner of a credit to have more
	// control over the credits they are putting in baskets then ordinary users
	// to deal with the scenario where basket tokens end up being worth
	// significantly less than the credits on their own.
	PickFromBasket(ctx context.Context, in *MsgPickFromBasket, opts ...grpc.CallOption) (*MsgPickFromBasketResponse, error)
}

type msgClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgClient(cc grpc.ClientConnInterface) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) CreateBasket(ctx context.Context, in *MsgCreateBasket, opts ...grpc.CallOption) (*MsgCreateBasketResponse, error) {
	out := new(MsgCreateBasketResponse)
	err := c.cc.Invoke(ctx, "/regen.ecocredit.basket.v1beta1.Msg/CreateBasket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) AddToBasket(ctx context.Context, in *MsgAddToBasket, opts ...grpc.CallOption) (*MsgAddToBasketResponse, error) {
	out := new(MsgAddToBasketResponse)
	err := c.cc.Invoke(ctx, "/regen.ecocredit.basket.v1beta1.Msg/AddToBasket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) TakeFromBasket(ctx context.Context, in *MsgTakeFromBasket, opts ...grpc.CallOption) (*MsgTakeFromBasketResponse, error) {
	out := new(MsgTakeFromBasketResponse)
	err := c.cc.Invoke(ctx, "/regen.ecocredit.basket.v1beta1.Msg/TakeFromBasket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) PickFromBasket(ctx context.Context, in *MsgPickFromBasket, opts ...grpc.CallOption) (*MsgPickFromBasketResponse, error) {
	out := new(MsgPickFromBasketResponse)
	err := c.cc.Invoke(ctx, "/regen.ecocredit.basket.v1beta1.Msg/PickFromBasket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
// All implementations must embed UnimplementedMsgServer
// for forward compatibility
type MsgServer interface {
	// CreateBasket creates a bank denom which wraps credits.
	CreateBasket(context.Context, *MsgCreateBasket) (*MsgCreateBasketResponse, error)
	// AddToBasket adds credits to a basket in return for basket tokens.
	AddToBasket(context.Context, *MsgAddToBasket) (*MsgAddToBasketResponse, error)
	// TakeFromBasket takes credits from a basket without regard for which
	// credits they are. The credits will be auto-retired if disable_auto_retire
	// is false. Credits will be chosen randomly using the previous block hash
	// as a consensus source of randomness.
	// More concretely, the implementation is as follows:
	// - take the previous block hash and convert it into an uint64,
	// - given the total number of different credits within the basket `n`, the
	//   first credits that will get picked correspond to: hash modulo n (in
	//   terms of order),
	// - then if we need to take more credits, we get some from the next one and
	//   so on.
	TakeFromBasket(context.Context, *MsgTakeFromBasket) (*MsgTakeFromBasketResponse, error)
	// PickFromBasket picks specific credits from a basket. If allow_picking is
	// set to false, then only an address which deposited credits in the basket
	// can pick those credits. All other addresses will be blocked from picking
	// those credits. The credits will be auto-retired if disable_auto_retire is
	// false unless the credits were previously put into the basket by the
	// address picking them from the basket, in which case they will remain
	// tradable. This functionality allows the owner of a credit to have more
	// control over the credits they are putting in baskets then ordinary users
	// to deal with the scenario where basket tokens end up being worth
	// significantly less than the credits on their own.
	PickFromBasket(context.Context, *MsgPickFromBasket) (*MsgPickFromBasketResponse, error)
	mustEmbedUnimplementedMsgServer()
}

// UnimplementedMsgServer must be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (UnimplementedMsgServer) CreateBasket(context.Context, *MsgCreateBasket) (*MsgCreateBasketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBasket not implemented")
}
func (UnimplementedMsgServer) AddToBasket(context.Context, *MsgAddToBasket) (*MsgAddToBasketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddToBasket not implemented")
}
func (UnimplementedMsgServer) TakeFromBasket(context.Context, *MsgTakeFromBasket) (*MsgTakeFromBasketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TakeFromBasket not implemented")
}
func (UnimplementedMsgServer) PickFromBasket(context.Context, *MsgPickFromBasket) (*MsgPickFromBasketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PickFromBasket not implemented")
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

func _Msg_CreateBasket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateBasket)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateBasket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/regen.ecocredit.basket.v1beta1.Msg/CreateBasket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateBasket(ctx, req.(*MsgCreateBasket))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_AddToBasket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgAddToBasket)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).AddToBasket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/regen.ecocredit.basket.v1beta1.Msg/AddToBasket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).AddToBasket(ctx, req.(*MsgAddToBasket))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_TakeFromBasket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgTakeFromBasket)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).TakeFromBasket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/regen.ecocredit.basket.v1beta1.Msg/TakeFromBasket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).TakeFromBasket(ctx, req.(*MsgTakeFromBasket))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_PickFromBasket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgPickFromBasket)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).PickFromBasket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/regen.ecocredit.basket.v1beta1.Msg/PickFromBasket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).PickFromBasket(ctx, req.(*MsgPickFromBasket))
	}
	return interceptor(ctx, in, info, handler)
}

// Msg_ServiceDesc is the grpc.ServiceDesc for Msg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "regen.ecocredit.basket.v1beta1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBasket",
			Handler:    _Msg_CreateBasket_Handler,
		},
		{
			MethodName: "AddToBasket",
			Handler:    _Msg_AddToBasket_Handler,
		},
		{
			MethodName: "TakeFromBasket",
			Handler:    _Msg_TakeFromBasket_Handler,
		},
		{
			MethodName: "PickFromBasket",
			Handler:    _Msg_PickFromBasket_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "regen/ecocredit/basket/v1beta1/tx.proto",
}
