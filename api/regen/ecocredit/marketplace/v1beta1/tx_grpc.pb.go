// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: regen/ecocredit/marketplace/v1beta1/tx.proto

package marketplacev1beta1

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
	// Sell creates new sell orders.
	Sell(ctx context.Context, in *MsgSell, opts ...grpc.CallOption) (*MsgSellResponse, error)
	// UpdateSellOrders updates existing sell orders.
	UpdateSellOrders(ctx context.Context, in *MsgUpdateSellOrders, opts ...grpc.CallOption) (*MsgUpdateSellOrdersResponse, error)
	// Buy creates credit buy orders.
	Buy(ctx context.Context, in *MsgBuy, opts ...grpc.CallOption) (*MsgBuyResponse, error)
	// AllowAskDenom is a governance operation which authorizes a new ask denom to be used in sell orders
	AllowAskDenom(ctx context.Context, in *MsgAllowAskDenom, opts ...grpc.CallOption) (*MsgAllowAskDenomResponse, error)
}

type msgClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgClient(cc grpc.ClientConnInterface) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) Sell(ctx context.Context, in *MsgSell, opts ...grpc.CallOption) (*MsgSellResponse, error) {
	out := new(MsgSellResponse)
	err := c.cc.Invoke(ctx, "/regen.ecocredit.marketplace.v1beta1.Msg/Sell", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateSellOrders(ctx context.Context, in *MsgUpdateSellOrders, opts ...grpc.CallOption) (*MsgUpdateSellOrdersResponse, error) {
	out := new(MsgUpdateSellOrdersResponse)
	err := c.cc.Invoke(ctx, "/regen.ecocredit.marketplace.v1beta1.Msg/UpdateSellOrders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) Buy(ctx context.Context, in *MsgBuy, opts ...grpc.CallOption) (*MsgBuyResponse, error) {
	out := new(MsgBuyResponse)
	err := c.cc.Invoke(ctx, "/regen.ecocredit.marketplace.v1beta1.Msg/Buy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) AllowAskDenom(ctx context.Context, in *MsgAllowAskDenom, opts ...grpc.CallOption) (*MsgAllowAskDenomResponse, error) {
	out := new(MsgAllowAskDenomResponse)
	err := c.cc.Invoke(ctx, "/regen.ecocredit.marketplace.v1beta1.Msg/AllowAskDenom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
// All implementations must embed UnimplementedMsgServer
// for forward compatibility
type MsgServer interface {
	// Sell creates new sell orders.
	Sell(context.Context, *MsgSell) (*MsgSellResponse, error)
	// UpdateSellOrders updates existing sell orders.
	UpdateSellOrders(context.Context, *MsgUpdateSellOrders) (*MsgUpdateSellOrdersResponse, error)
	// Buy creates credit buy orders.
	Buy(context.Context, *MsgBuy) (*MsgBuyResponse, error)
	// AllowAskDenom is a governance operation which authorizes a new ask denom to be used in sell orders
	AllowAskDenom(context.Context, *MsgAllowAskDenom) (*MsgAllowAskDenomResponse, error)
	mustEmbedUnimplementedMsgServer()
}

// UnimplementedMsgServer must be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (UnimplementedMsgServer) Sell(context.Context, *MsgSell) (*MsgSellResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sell not implemented")
}
func (UnimplementedMsgServer) UpdateSellOrders(context.Context, *MsgUpdateSellOrders) (*MsgUpdateSellOrdersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSellOrders not implemented")
}
func (UnimplementedMsgServer) Buy(context.Context, *MsgBuy) (*MsgBuyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Buy not implemented")
}
func (UnimplementedMsgServer) AllowAskDenom(context.Context, *MsgAllowAskDenom) (*MsgAllowAskDenomResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AllowAskDenom not implemented")
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

func _Msg_Sell_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSell)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Sell(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/regen.ecocredit.marketplace.v1beta1.Msg/Sell",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Sell(ctx, req.(*MsgSell))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateSellOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateSellOrders)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateSellOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/regen.ecocredit.marketplace.v1beta1.Msg/UpdateSellOrders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateSellOrders(ctx, req.(*MsgUpdateSellOrders))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_Buy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgBuy)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Buy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/regen.ecocredit.marketplace.v1beta1.Msg/Buy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Buy(ctx, req.(*MsgBuy))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_AllowAskDenom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgAllowAskDenom)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).AllowAskDenom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/regen.ecocredit.marketplace.v1beta1.Msg/AllowAskDenom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).AllowAskDenom(ctx, req.(*MsgAllowAskDenom))
	}
	return interceptor(ctx, in, info, handler)
}

// Msg_ServiceDesc is the grpc.ServiceDesc for Msg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "regen.ecocredit.marketplace.v1beta1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Sell",
			Handler:    _Msg_Sell_Handler,
		},
		{
			MethodName: "UpdateSellOrders",
			Handler:    _Msg_UpdateSellOrders_Handler,
		},
		{
			MethodName: "Buy",
			Handler:    _Msg_Buy_Handler,
		},
		{
			MethodName: "AllowAskDenom",
			Handler:    _Msg_AllowAskDenom_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "regen/ecocredit/marketplace/v1beta1/tx.proto",
}
