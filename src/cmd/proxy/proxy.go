package proxy

import (
	"context"
	"github.com/onflow/flow-dps/api/dps"
	"github.com/onflow/flow-dps/codec/zbor"
	"github.com/onflow/flow-go/engine/access/rpc/backend"
	"github.com/onflow/flow-go/utils/grpcutils"
	"github.com/onflow/flow/protobuf/go/flow/access"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

func NewDPSService(logger zerolog.Logger) *DPSService {
	// No public key means an insecure channel
	clientRPCConnection, err := grpc.Dial(
		"access:9000",
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(grpcutils.DefaultMaxMsgSize)),
		grpc.WithInsecure(), //nolint:staticcheck
		backend.WithClientUnaryInterceptor(3*time.Second))
	if err != nil {
		return nil
	}

	return &DPSService{
		codec:  zbor.NewCodec(),
		client: access.NewAccessAPIClient(clientRPCConnection),
		logger: logger,
	}
}

// This is a quick implementation of DPS gRPC server that
// forwards requests to a live node.
type DPSService struct {
	dps.APIServer
	codec  *zbor.Codec
	client access.AccessAPIClient
	logger zerolog.Logger
}

func (d *DPSService) GetFirst(ctx context.Context, r *dps.GetFirstRequest) (*dps.GetFirstResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method not implemented")
}

func (d *DPSService) GetLast(ctx context.Context, req *dps.GetLastRequest) (*dps.GetLastResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method not implemented")
}

func (d *DPSService) GetHeightForBlock(ctx context.Context, r *dps.GetHeightForBlockRequest) (*dps.GetHeightForBlockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method not implemented")
}

func (d *DPSService) GetCommit(ctx context.Context, r *dps.GetCommitRequest) (*dps.GetCommitResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method not implemented")
}

func (d *DPSService) GetHeader(ctx context.Context, req *dps.GetHeaderRequest) (*dps.GetHeaderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method not implemented")
}
func (d *DPSService) GetEvents(ctx context.Context, r *dps.GetEventsRequest) (*dps.GetEventsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method not implemented")
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
func (d *DPSService) GetSeal(ctx context.Context, rr *dps.GetSealRequest) (*dps.GetSealResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSeal not implemented")
}
func (d *DPSService) ListSealsForHeight(ctx context.Context, r *dps.ListSealsForHeightRequest) (*dps.ListSealsForHeightResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSealsForHeight not implemented")
}
