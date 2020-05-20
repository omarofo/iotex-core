// Code generated by MockGen. DO NOT EDIT.
// Source: ./api/apitestclient.go

// Package mock_apiserviceclient is a generated GoMock package.
package mock_apiserviceclient

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	iotexapi "github.com/iotexproject/iotex-proto/golang/iotexapi"
	grpc "google.golang.org/grpc"
	reflect "reflect"
)

// MockServiceClient is a mock of ServiceClient interface
type MockServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockServiceClientMockRecorder
}

// MockServiceClientMockRecorder is the mock recorder for MockServiceClient
type MockServiceClientMockRecorder struct {
	mock *MockServiceClient
}

// NewMockServiceClient creates a new mock instance
func NewMockServiceClient(ctrl *gomock.Controller) *MockServiceClient {
	mock := &MockServiceClient{ctrl: ctrl}
	mock.recorder = &MockServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockServiceClient) EXPECT() *MockServiceClientMockRecorder {
	return m.recorder
}

// GetAccount mocks base method
func (m *MockServiceClient) GetAccount(ctx context.Context, in *iotexapi.GetAccountRequest, opts ...grpc.CallOption) (*iotexapi.GetAccountResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetAccount", varargs...)
	ret0, _ := ret[0].(*iotexapi.GetAccountResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccount indicates an expected call of GetAccount
func (mr *MockServiceClientMockRecorder) GetAccount(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccount", reflect.TypeOf((*MockServiceClient)(nil).GetAccount), varargs...)
}

// GetActions mocks base method
func (m *MockServiceClient) GetActions(ctx context.Context, in *iotexapi.GetActionsRequest, opts ...grpc.CallOption) (*iotexapi.GetActionsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetActions", varargs...)
	ret0, _ := ret[0].(*iotexapi.GetActionsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetActions indicates an expected call of GetActions
func (mr *MockServiceClientMockRecorder) GetActions(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetActions", reflect.TypeOf((*MockServiceClient)(nil).GetActions), varargs...)
}

// GetBlockMetas mocks base method
func (m *MockServiceClient) GetBlockMetas(ctx context.Context, in *iotexapi.GetBlockMetasRequest, opts ...grpc.CallOption) (*iotexapi.GetBlockMetasResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBlockMetas", varargs...)
	ret0, _ := ret[0].(*iotexapi.GetBlockMetasResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlockMetas indicates an expected call of GetBlockMetas
func (mr *MockServiceClientMockRecorder) GetBlockMetas(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlockMetas", reflect.TypeOf((*MockServiceClient)(nil).GetBlockMetas), varargs...)
}

// GetChainMeta mocks base method
func (m *MockServiceClient) GetChainMeta(ctx context.Context, in *iotexapi.GetChainMetaRequest, opts ...grpc.CallOption) (*iotexapi.GetChainMetaResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetChainMeta", varargs...)
	ret0, _ := ret[0].(*iotexapi.GetChainMetaResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetChainMeta indicates an expected call of GetChainMeta
func (mr *MockServiceClientMockRecorder) GetChainMeta(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetChainMeta", reflect.TypeOf((*MockServiceClient)(nil).GetChainMeta), varargs...)
}

// GetServerMeta mocks base method
func (m *MockServiceClient) GetServerMeta(ctx context.Context, in *iotexapi.GetServerMetaRequest, opts ...grpc.CallOption) (*iotexapi.GetServerMetaResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetServerMeta", varargs...)
	ret0, _ := ret[0].(*iotexapi.GetServerMetaResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetServerMeta indicates an expected call of GetServerMeta
func (mr *MockServiceClientMockRecorder) GetServerMeta(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetServerMeta", reflect.TypeOf((*MockServiceClient)(nil).GetServerMeta), varargs...)
}

// SendAction mocks base method
func (m *MockServiceClient) SendAction(ctx context.Context, in *iotexapi.SendActionRequest, opts ...grpc.CallOption) (*iotexapi.SendActionResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SendAction", varargs...)
	ret0, _ := ret[0].(*iotexapi.SendActionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendAction indicates an expected call of SendAction
func (mr *MockServiceClientMockRecorder) SendAction(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendAction", reflect.TypeOf((*MockServiceClient)(nil).SendAction), varargs...)
}

// GetReceiptByAction mocks base method
func (m *MockServiceClient) GetReceiptByAction(ctx context.Context, in *iotexapi.GetReceiptByActionRequest, opts ...grpc.CallOption) (*iotexapi.GetReceiptByActionResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetReceiptByAction", varargs...)
	ret0, _ := ret[0].(*iotexapi.GetReceiptByActionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetReceiptByAction indicates an expected call of GetReceiptByAction
func (mr *MockServiceClientMockRecorder) GetReceiptByAction(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReceiptByAction", reflect.TypeOf((*MockServiceClient)(nil).GetReceiptByAction), varargs...)
}

// ReadContract mocks base method
func (m *MockServiceClient) ReadContract(ctx context.Context, in *iotexapi.ReadContractRequest, opts ...grpc.CallOption) (*iotexapi.ReadContractResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ReadContract", varargs...)
	ret0, _ := ret[0].(*iotexapi.ReadContractResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadContract indicates an expected call of ReadContract
func (mr *MockServiceClientMockRecorder) ReadContract(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadContract", reflect.TypeOf((*MockServiceClient)(nil).ReadContract), varargs...)
}

// SuggestGasPrice mocks base method
func (m *MockServiceClient) SuggestGasPrice(ctx context.Context, in *iotexapi.SuggestGasPriceRequest, opts ...grpc.CallOption) (*iotexapi.SuggestGasPriceResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SuggestGasPrice", varargs...)
	ret0, _ := ret[0].(*iotexapi.SuggestGasPriceResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SuggestGasPrice indicates an expected call of SuggestGasPrice
func (mr *MockServiceClientMockRecorder) SuggestGasPrice(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SuggestGasPrice", reflect.TypeOf((*MockServiceClient)(nil).SuggestGasPrice), varargs...)
}

// EstimateGasForAction mocks base method
func (m *MockServiceClient) EstimateGasForAction(ctx context.Context, in *iotexapi.EstimateGasForActionRequest, opts ...grpc.CallOption) (*iotexapi.EstimateGasForActionResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "EstimateGasForAction", varargs...)
	ret0, _ := ret[0].(*iotexapi.EstimateGasForActionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EstimateGasForAction indicates an expected call of EstimateGasForAction
func (mr *MockServiceClientMockRecorder) EstimateGasForAction(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EstimateGasForAction", reflect.TypeOf((*MockServiceClient)(nil).EstimateGasForAction), varargs...)
}

// EstimateActionGasConsumption mocks base method
func (m *MockServiceClient) EstimateActionGasConsumption(ctx context.Context, in *iotexapi.EstimateActionGasConsumptionRequest, opts ...grpc.CallOption) (*iotexapi.EstimateActionGasConsumptionResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "EstimateActionGasConsumption", varargs...)
	ret0, _ := ret[0].(*iotexapi.EstimateActionGasConsumptionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EstimateActionGasConsumption indicates an expected call of EstimateActionGasConsumption
func (mr *MockServiceClientMockRecorder) EstimateActionGasConsumption(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EstimateActionGasConsumption", reflect.TypeOf((*MockServiceClient)(nil).EstimateActionGasConsumption), varargs...)
}

// ReadState mocks base method
func (m *MockServiceClient) ReadState(ctx context.Context, in *iotexapi.ReadStateRequest, opts ...grpc.CallOption) (*iotexapi.ReadStateResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ReadState", varargs...)
	ret0, _ := ret[0].(*iotexapi.ReadStateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadState indicates an expected call of ReadState
func (mr *MockServiceClientMockRecorder) ReadState(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadState", reflect.TypeOf((*MockServiceClient)(nil).ReadState), varargs...)
}

// GetEpochMeta mocks base method
func (m *MockServiceClient) GetEpochMeta(ctx context.Context, in *iotexapi.GetEpochMetaRequest, opts ...grpc.CallOption) (*iotexapi.GetEpochMetaResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetEpochMeta", varargs...)
	ret0, _ := ret[0].(*iotexapi.GetEpochMetaResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEpochMeta indicates an expected call of GetEpochMeta
func (mr *MockServiceClientMockRecorder) GetEpochMeta(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEpochMeta", reflect.TypeOf((*MockServiceClient)(nil).GetEpochMeta), varargs...)
}

// GetRawBlocks mocks base method
func (m *MockServiceClient) GetRawBlocks(ctx context.Context, in *iotexapi.GetRawBlocksRequest, opts ...grpc.CallOption) (*iotexapi.GetRawBlocksResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetRawBlocks", varargs...)
	ret0, _ := ret[0].(*iotexapi.GetRawBlocksResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRawBlocks indicates an expected call of GetRawBlocks
func (mr *MockServiceClientMockRecorder) GetRawBlocks(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRawBlocks", reflect.TypeOf((*MockServiceClient)(nil).GetRawBlocks), varargs...)
}

// GetLogs mocks base method
func (m *MockServiceClient) GetLogs(ctx context.Context, in *iotexapi.GetLogsRequest, opts ...grpc.CallOption) (*iotexapi.GetLogsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetLogs", varargs...)
	ret0, _ := ret[0].(*iotexapi.GetLogsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLogs indicates an expected call of GetLogs
func (mr *MockServiceClientMockRecorder) GetLogs(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLogs", reflect.TypeOf((*MockServiceClient)(nil).GetLogs), varargs...)
}

// GetEvmTransfersByActionHash mocks base method
func (m *MockServiceClient) GetEvmTransfersByActionHash(ctx context.Context, in *iotexapi.GetEvmTransfersByActionHashRequest, opts ...grpc.CallOption) (*iotexapi.GetEvmTransfersByActionHashResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetEvmTransfersByActionHash", varargs...)
	ret0, _ := ret[0].(*iotexapi.GetEvmTransfersByActionHashResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEvmTransfersByActionHash indicates an expected call of GetEvmTransfersByActionHash
func (mr *MockServiceClientMockRecorder) GetEvmTransfersByActionHash(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEvmTransfersByActionHash", reflect.TypeOf((*MockServiceClient)(nil).GetEvmTransfersByActionHash), varargs...)
}

// GetEvmTransfersByBlockHeight mocks base method
func (m *MockServiceClient) GetEvmTransfersByBlockHeight(ctx context.Context, in *iotexapi.GetEvmTransfersByBlockHeightRequest, opts ...grpc.CallOption) (*iotexapi.GetEvmTransfersByBlockHeightResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetEvmTransfersByBlockHeight", varargs...)
	ret0, _ := ret[0].(*iotexapi.GetEvmTransfersByBlockHeightResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEvmTransfersByBlockHeight indicates an expected call of GetEvmTransfersByBlockHeight
func (mr *MockServiceClientMockRecorder) GetEvmTransfersByBlockHeight(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEvmTransfersByBlockHeight", reflect.TypeOf((*MockServiceClient)(nil).GetEvmTransfersByBlockHeight), varargs...)
}

// StreamBlocks mocks base method
func (m *MockServiceClient) StreamBlocks(ctx context.Context, in *iotexapi.StreamBlocksRequest, opts ...grpc.CallOption) (iotexapi.APIService_StreamBlocksClient, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "StreamBlocks", varargs...)
	ret0, _ := ret[0].(iotexapi.APIService_StreamBlocksClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StreamBlocks indicates an expected call of StreamBlocks
func (mr *MockServiceClientMockRecorder) StreamBlocks(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StreamBlocks", reflect.TypeOf((*MockServiceClient)(nil).StreamBlocks), varargs...)
}

// StreamLogs mocks base method
func (m *MockServiceClient) StreamLogs(ctx context.Context, in *iotexapi.StreamLogsRequest, opts ...grpc.CallOption) (iotexapi.APIService_StreamLogsClient, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "StreamLogs", varargs...)
	ret0, _ := ret[0].(iotexapi.APIService_StreamLogsClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StreamLogs indicates an expected call of StreamLogs
func (mr *MockServiceClientMockRecorder) StreamLogs(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StreamLogs", reflect.TypeOf((*MockServiceClient)(nil).StreamLogs), varargs...)
}

// GetElectionBuckets mocks base method
func (m *MockServiceClient) GetElectionBuckets(ctx context.Context, in *iotexapi.GetElectionBucketsRequest, opts ...grpc.CallOption) (*iotexapi.GetElectionBucketsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetElectionBuckets", varargs...)
	ret0, _ := ret[0].(*iotexapi.GetElectionBucketsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetElectionBuckets indicates an expected call of GetElectionBuckets
func (mr *MockServiceClientMockRecorder) GetElectionBuckets(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetElectionBuckets", reflect.TypeOf((*MockServiceClient)(nil).GetElectionBuckets), varargs...)
}
