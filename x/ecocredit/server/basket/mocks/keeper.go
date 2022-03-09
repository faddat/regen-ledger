// Code generated by MockGen. DO NOT EDIT.
// Source: x/ecocredit/server/basket/keeper.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	types "github.com/cosmos/cosmos-sdk/types"
	gomock "github.com/golang/mock/gomock"
	types0 "github.com/regen-network/regen-ledger/types"
	ecocredit "github.com/regen-network/regen-ledger/x/ecocredit"
)

// MockEcocreditKeeper is a mock of EcocreditKeeper interface.
type MockEcocreditKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockEcocreditKeeperMockRecorder
}

// MockEcocreditKeeperMockRecorder is the mock recorder for MockEcocreditKeeper.
type MockEcocreditKeeperMockRecorder struct {
	mock *MockEcocreditKeeper
}

// NewMockEcocreditKeeper creates a new mock instance.
func NewMockEcocreditKeeper(ctrl *gomock.Controller) *MockEcocreditKeeper {
	mock := &MockEcocreditKeeper{ctrl: ctrl}
	mock.recorder = &MockEcocreditKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEcocreditKeeper) EXPECT() *MockEcocreditKeeperMockRecorder {
	return m.recorder
}

// AllowedAskDenoms mocks base method.
func (m *MockEcocreditKeeper) AllowedAskDenoms(arg0 context.Context, arg1 *ecocredit.QueryAllowedAskDenomsRequest) (*ecocredit.QueryAllowedAskDenomsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllowedAskDenoms", arg0, arg1)
	ret0, _ := ret[0].(*ecocredit.QueryAllowedAskDenomsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AllowedAskDenoms indicates an expected call of AllowedAskDenoms.
func (mr *MockEcocreditKeeperMockRecorder) AllowedAskDenoms(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllowedAskDenoms", reflect.TypeOf((*MockEcocreditKeeper)(nil).AllowedAskDenoms), arg0, arg1)
}

// Balance mocks base method.
func (m *MockEcocreditKeeper) Balance(arg0 context.Context, arg1 *ecocredit.QueryBalanceRequest) (*ecocredit.QueryBalanceResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Balance", arg0, arg1)
	ret0, _ := ret[0].(*ecocredit.QueryBalanceResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Balance indicates an expected call of Balance.
func (mr *MockEcocreditKeeperMockRecorder) Balance(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Balance", reflect.TypeOf((*MockEcocreditKeeper)(nil).Balance), arg0, arg1)
}

// Basket mocks base method.
func (m *MockEcocreditKeeper) Basket(arg0 context.Context, arg1 *ecocredit.QueryBasketRequest) (*ecocredit.QueryBasketResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Basket", arg0, arg1)
	ret0, _ := ret[0].(*ecocredit.QueryBasketResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Basket indicates an expected call of Basket.
func (mr *MockEcocreditKeeperMockRecorder) Basket(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Basket", reflect.TypeOf((*MockEcocreditKeeper)(nil).Basket), arg0, arg1)
}

// BasketCredits mocks base method.
func (m *MockEcocreditKeeper) BasketCredits(arg0 context.Context, arg1 *ecocredit.QueryBasketCreditsRequest) (*ecocredit.QueryBasketCreditsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BasketCredits", arg0, arg1)
	ret0, _ := ret[0].(*ecocredit.QueryBasketCreditsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BasketCredits indicates an expected call of BasketCredits.
func (mr *MockEcocreditKeeperMockRecorder) BasketCredits(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BasketCredits", reflect.TypeOf((*MockEcocreditKeeper)(nil).BasketCredits), arg0, arg1)
}

// Baskets mocks base method.
func (m *MockEcocreditKeeper) Baskets(arg0 context.Context, arg1 *ecocredit.QueryBasketsRequest) (*ecocredit.QueryBasketsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Baskets", arg0, arg1)
	ret0, _ := ret[0].(*ecocredit.QueryBasketsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Baskets indicates an expected call of Baskets.
func (mr *MockEcocreditKeeperMockRecorder) Baskets(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Baskets", reflect.TypeOf((*MockEcocreditKeeper)(nil).Baskets), arg0, arg1)
}

// BatchInfo mocks base method.
func (m *MockEcocreditKeeper) BatchInfo(arg0 context.Context, arg1 *ecocredit.QueryBatchInfoRequest) (*ecocredit.QueryBatchInfoResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BatchInfo", arg0, arg1)
	ret0, _ := ret[0].(*ecocredit.QueryBatchInfoResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BatchInfo indicates an expected call of BatchInfo.
func (mr *MockEcocreditKeeperMockRecorder) BatchInfo(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchInfo", reflect.TypeOf((*MockEcocreditKeeper)(nil).BatchInfo), arg0, arg1)
}

// Batches mocks base method.
func (m *MockEcocreditKeeper) Batches(arg0 context.Context, arg1 *ecocredit.QueryBatchesRequest) (*ecocredit.QueryBatchesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Batches", arg0, arg1)
	ret0, _ := ret[0].(*ecocredit.QueryBatchesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Batches indicates an expected call of Batches.
func (mr *MockEcocreditKeeperMockRecorder) Batches(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Batches", reflect.TypeOf((*MockEcocreditKeeper)(nil).Batches), arg0, arg1)
}

// BuyOrder mocks base method.
func (m *MockEcocreditKeeper) BuyOrder(arg0 context.Context, arg1 *ecocredit.QueryBuyOrderRequest) (*ecocredit.QueryBuyOrderResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BuyOrder", arg0, arg1)
	ret0, _ := ret[0].(*ecocredit.QueryBuyOrderResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BuyOrder indicates an expected call of BuyOrder.
func (mr *MockEcocreditKeeperMockRecorder) BuyOrder(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BuyOrder", reflect.TypeOf((*MockEcocreditKeeper)(nil).BuyOrder), arg0, arg1)
}

// BuyOrders mocks base method.
func (m *MockEcocreditKeeper) BuyOrders(arg0 context.Context, arg1 *ecocredit.QueryBuyOrdersRequest) (*ecocredit.QueryBuyOrdersResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BuyOrders", arg0, arg1)
	ret0, _ := ret[0].(*ecocredit.QueryBuyOrdersResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BuyOrders indicates an expected call of BuyOrders.
func (mr *MockEcocreditKeeperMockRecorder) BuyOrders(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BuyOrders", reflect.TypeOf((*MockEcocreditKeeper)(nil).BuyOrders), arg0, arg1)
}

// BuyOrdersByAddress mocks base method.
func (m *MockEcocreditKeeper) BuyOrdersByAddress(arg0 context.Context, arg1 *ecocredit.QueryBuyOrdersByAddressRequest) (*ecocredit.QueryBuyOrdersByAddressResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BuyOrdersByAddress", arg0, arg1)
	ret0, _ := ret[0].(*ecocredit.QueryBuyOrdersByAddressResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BuyOrdersByAddress indicates an expected call of BuyOrdersByAddress.
func (mr *MockEcocreditKeeperMockRecorder) BuyOrdersByAddress(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BuyOrdersByAddress", reflect.TypeOf((*MockEcocreditKeeper)(nil).BuyOrdersByAddress), arg0, arg1)
}

// ClassInfo mocks base method.
func (m *MockEcocreditKeeper) ClassInfo(arg0 context.Context, arg1 *ecocredit.QueryClassInfoRequest) (*ecocredit.QueryClassInfoResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClassInfo", arg0, arg1)
	ret0, _ := ret[0].(*ecocredit.QueryClassInfoResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ClassInfo indicates an expected call of ClassInfo.
func (mr *MockEcocreditKeeperMockRecorder) ClassInfo(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClassInfo", reflect.TypeOf((*MockEcocreditKeeper)(nil).ClassInfo), arg0, arg1)
}

// Classes mocks base method.
func (m *MockEcocreditKeeper) Classes(arg0 context.Context, arg1 *ecocredit.QueryClassesRequest) (*ecocredit.QueryClassesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Classes", arg0, arg1)
	ret0, _ := ret[0].(*ecocredit.QueryClassesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Classes indicates an expected call of Classes.
func (mr *MockEcocreditKeeperMockRecorder) Classes(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Classes", reflect.TypeOf((*MockEcocreditKeeper)(nil).Classes), arg0, arg1)
}

// CreditTypes mocks base method.
func (m *MockEcocreditKeeper) CreditTypes(arg0 context.Context, arg1 *ecocredit.QueryCreditTypesRequest) (*ecocredit.QueryCreditTypesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreditTypes", arg0, arg1)
	ret0, _ := ret[0].(*ecocredit.QueryCreditTypesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreditTypes indicates an expected call of CreditTypes.
func (mr *MockEcocreditKeeperMockRecorder) CreditTypes(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreditTypes", reflect.TypeOf((*MockEcocreditKeeper)(nil).CreditTypes), arg0, arg1)
}

// GetCreateBasketFee mocks base method.
func (m *MockEcocreditKeeper) GetCreateBasketFee(ctx context.Context) types.Coins {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCreateBasketFee", ctx)
	ret0, _ := ret[0].(types.Coins)
	return ret0
}

// GetCreateBasketFee indicates an expected call of GetCreateBasketFee.
func (mr *MockEcocreditKeeperMockRecorder) GetCreateBasketFee(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCreateBasketFee", reflect.TypeOf((*MockEcocreditKeeper)(nil).GetCreateBasketFee), ctx)
}

// HasClassInfo mocks base method.
func (m *MockEcocreditKeeper) HasClassInfo(ctx types0.Context, classID string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasClassInfo", ctx, classID)
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasClassInfo indicates an expected call of HasClassInfo.
func (mr *MockEcocreditKeeperMockRecorder) HasClassInfo(ctx, classID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasClassInfo", reflect.TypeOf((*MockEcocreditKeeper)(nil).HasClassInfo), ctx, classID)
}

// Params mocks base method.
func (m *MockEcocreditKeeper) Params(arg0 context.Context, arg1 *ecocredit.QueryParamsRequest) (*ecocredit.QueryParamsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Params", arg0, arg1)
	ret0, _ := ret[0].(*ecocredit.QueryParamsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Params indicates an expected call of Params.
func (mr *MockEcocreditKeeperMockRecorder) Params(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Params", reflect.TypeOf((*MockEcocreditKeeper)(nil).Params), arg0, arg1)
}

// ProjectInfo mocks base method.
func (m *MockEcocreditKeeper) ProjectInfo(arg0 context.Context, arg1 *ecocredit.QueryProjectInfoRequest) (*ecocredit.QueryProjectInfoResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProjectInfo", arg0, arg1)
	ret0, _ := ret[0].(*ecocredit.QueryProjectInfoResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProjectInfo indicates an expected call of ProjectInfo.
func (mr *MockEcocreditKeeperMockRecorder) ProjectInfo(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProjectInfo", reflect.TypeOf((*MockEcocreditKeeper)(nil).ProjectInfo), arg0, arg1)
}

// Projects mocks base method.
func (m *MockEcocreditKeeper) Projects(arg0 context.Context, arg1 *ecocredit.QueryProjectsRequest) (*ecocredit.QueryProjectsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Projects", arg0, arg1)
	ret0, _ := ret[0].(*ecocredit.QueryProjectsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Projects indicates an expected call of Projects.
func (mr *MockEcocreditKeeperMockRecorder) Projects(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Projects", reflect.TypeOf((*MockEcocreditKeeper)(nil).Projects), arg0, arg1)
}

// SellOrder mocks base method.
func (m *MockEcocreditKeeper) SellOrder(arg0 context.Context, arg1 *ecocredit.QuerySellOrderRequest) (*ecocredit.QuerySellOrderResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SellOrder", arg0, arg1)
	ret0, _ := ret[0].(*ecocredit.QuerySellOrderResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SellOrder indicates an expected call of SellOrder.
func (mr *MockEcocreditKeeperMockRecorder) SellOrder(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SellOrder", reflect.TypeOf((*MockEcocreditKeeper)(nil).SellOrder), arg0, arg1)
}

// SellOrders mocks base method.
func (m *MockEcocreditKeeper) SellOrders(arg0 context.Context, arg1 *ecocredit.QuerySellOrdersRequest) (*ecocredit.QuerySellOrdersResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SellOrders", arg0, arg1)
	ret0, _ := ret[0].(*ecocredit.QuerySellOrdersResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SellOrders indicates an expected call of SellOrders.
func (mr *MockEcocreditKeeperMockRecorder) SellOrders(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SellOrders", reflect.TypeOf((*MockEcocreditKeeper)(nil).SellOrders), arg0, arg1)
}

// SellOrdersByAddress mocks base method.
func (m *MockEcocreditKeeper) SellOrdersByAddress(arg0 context.Context, arg1 *ecocredit.QuerySellOrdersByAddressRequest) (*ecocredit.QuerySellOrdersByAddressResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SellOrdersByAddress", arg0, arg1)
	ret0, _ := ret[0].(*ecocredit.QuerySellOrdersByAddressResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SellOrdersByAddress indicates an expected call of SellOrdersByAddress.
func (mr *MockEcocreditKeeperMockRecorder) SellOrdersByAddress(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SellOrdersByAddress", reflect.TypeOf((*MockEcocreditKeeper)(nil).SellOrdersByAddress), arg0, arg1)
}

// SellOrdersByBatchDenom mocks base method.
func (m *MockEcocreditKeeper) SellOrdersByBatchDenom(arg0 context.Context, arg1 *ecocredit.QuerySellOrdersByBatchDenomRequest) (*ecocredit.QuerySellOrdersByBatchDenomResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SellOrdersByBatchDenom", arg0, arg1)
	ret0, _ := ret[0].(*ecocredit.QuerySellOrdersByBatchDenomResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SellOrdersByBatchDenom indicates an expected call of SellOrdersByBatchDenom.
func (mr *MockEcocreditKeeperMockRecorder) SellOrdersByBatchDenom(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SellOrdersByBatchDenom", reflect.TypeOf((*MockEcocreditKeeper)(nil).SellOrdersByBatchDenom), arg0, arg1)
}

// Supply mocks base method.
func (m *MockEcocreditKeeper) Supply(arg0 context.Context, arg1 *ecocredit.QuerySupplyRequest) (*ecocredit.QuerySupplyResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Supply", arg0, arg1)
	ret0, _ := ret[0].(*ecocredit.QuerySupplyResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Supply indicates an expected call of Supply.
func (mr *MockEcocreditKeeperMockRecorder) Supply(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Supply", reflect.TypeOf((*MockEcocreditKeeper)(nil).Supply), arg0, arg1)
}
