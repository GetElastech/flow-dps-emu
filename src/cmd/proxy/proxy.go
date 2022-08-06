package proxy

import (
	"context"
	"github.com/onflow/flow-dps/api/dps"
	"github.com/onflow/flow-dps/codec/zbor"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func NewDPSService() *DPSService {
	return &DPSService{
		codec: zbor.NewCodec(),
	}
}

// This is a quick implementation of DPS gRPC server that
// forwards requests to a live node.
type DPSService struct {
	dps.APIServer
	codec *zbor.Codec
}

func (DPSService) GetFirst(context.Context, *dps.GetFirstRequest) (*dps.GetFirstResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFirst not implemented")
}
func (DPSService) GetLast(context.Context, *dps.GetLastRequest) (*dps.GetLastResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLast not implemented")
}
func (DPSService) GetHeightForBlock(context.Context, *dps.GetHeightForBlockRequest) (*dps.GetHeightForBlockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHeightForBlock not implemented")
}
func (DPSService) GetCommit(context.Context, *dps.GetCommitRequest) (*dps.GetCommitResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCommit not implemented")
}
func (DPSService) GetHeader(context.Context, *dps.GetHeaderRequest) (*dps.GetHeaderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHeader not implemented")
}
func (DPSService) GetEvents(context.Context, *dps.GetEventsRequest) (*dps.GetEventsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEvents not implemented")
}
func (DPSService) GetRegisterValues(context.Context, *dps.GetRegisterValuesRequest) (*dps.GetRegisterValuesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRegisterValues not implemented")
}
func (DPSService) GetCollection(context.Context, *dps.GetCollectionRequest) (*dps.GetCollectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCollection not implemented")
}
func (DPSService) ListCollectionsForHeight(context.Context, *dps.ListCollectionsForHeightRequest) (*dps.ListCollectionsForHeightResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCollectionsForHeight not implemented")
}
func (DPSService) GetGuarantee(context.Context, *dps.GetGuaranteeRequest) (*dps.GetGuaranteeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGuarantee not implemented")
}
func (DPSService) GetTransaction(context.Context, *dps.GetTransactionRequest) (*dps.GetTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransaction not implemented")
}
func (DPSService) GetHeightForTransaction(context.Context, *dps.GetHeightForTransactionRequest) (*dps.GetHeightForTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHeightForTransaction not implemented")
}
func (DPSService) ListTransactionsForHeight(context.Context, *dps.ListTransactionsForHeightRequest) (*dps.ListTransactionsForHeightResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTransactionsForHeight not implemented")
}
func (DPSService) GetResult(context.Context, *dps.GetResultRequest) (*dps.GetResultResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetResult not implemented")
}
func (DPSService) GetSeal(context.Context, *dps.GetSealRequest) (*dps.GetSealResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSeal not implemented")
}
func (DPSService) ListSealsForHeight(context.Context, *dps.ListSealsForHeightRequest) (*dps.ListSealsForHeightResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSealsForHeight not implemented")
}
