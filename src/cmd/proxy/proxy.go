package proxy

import (
	"context"
	"fmt"
	"github.com/onflow/flow-dps/api/dps"
	"github.com/onflow/flow-dps/codec/zbor"
	"github.com/onflow/flow-go/engine/access/rpc/backend"
	"github.com/onflow/flow-go/model/flow"
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
		"127.0.0.1:4500",
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
	last, err := d.GetLast(ctx, &dps.GetLastRequest{})
	if err != nil {
		return nil, err
	}
	for height := uint64(0); height < last.GetHeight(); height++ {
		res, err := d.client.GetBlockByHeight(ctx, &access.GetBlockByHeightRequest{
			Height: height,
		})
		if err == nil && res.Block.Height == height {
			return &dps.GetFirstResponse{
				Height: height,
			}, nil
		}
	}
	return nil, status.Errorf(codes.NotFound, "starting height not found")
}

func (d *DPSService) GetLast(ctx context.Context, req *dps.GetLastRequest) (*dps.GetLastResponse, error) {
	res, err := d.client.GetLatestBlockHeader(ctx, &access.GetLatestBlockHeaderRequest{
		IsSealed: true,
	})
	if err != nil {
		return nil, err
	}
	return &dps.GetLastResponse{
		Height: res.Block.Height,
	}, nil
}

func (d *DPSService) GetHeightForBlock(ctx context.Context, r *dps.GetHeightForBlockRequest) (*dps.GetHeightForBlockResponse, error) {
	res, err := d.client.GetBlockByID(ctx, &access.GetBlockByIDRequest{
		Id: r.BlockID,
	})
	if err != nil {
		return nil, err
	}
	return &dps.GetHeightForBlockResponse{
		Height: res.Block.Height,
	}, nil
}

func (d *DPSService) GetCommit(ctx context.Context, r *dps.GetCommitRequest) (*dps.GetCommitResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCommit not implemented")
}

func (d *DPSService) GetHeader(ctx context.Context, req *dps.GetHeaderRequest) (*dps.GetHeaderResponse, error) {
	res, err := d.client.GetBlockHeaderByHeight(ctx, &access.GetBlockHeaderByHeightRequest{
		Height: req.Height,
	})
	if err != nil {
		return nil, err
	}

	hdr := flow.Header{
		ChainID:            flow.ChainID(res.Block.ChainId),
		ParentID:           flow.MustHexStringToIdentifier(string(res.Block.ParentId)),
		Height:             res.Block.Height,
		PayloadHash:        flow.MustHexStringToIdentifier(string(res.Block.PayloadHash)),
		Timestamp:          (*res.Block.Timestamp).AsTime(),
		View:               res.Block.View,
		ParentVoterIndices: res.Block.ParentVoterIndices,
		ParentVoterSigData: res.Block.ParentVoterSigData,
		ProposerID:         flow.MustHexStringToIdentifier(string(res.Block.ProposerId)),
		ProposerSigData:    res.Block.ProposerSigData,
	}

	b, err := d.codec.Marshal(hdr)
	return &dps.GetHeaderResponse{
		Data:   b,
		Height: res.Block.Height,
	}, nil
}
func (d *DPSService) GetEvents(ctx context.Context, r *dps.GetEventsRequest) (*dps.GetEventsResponse, error) {
	list := make([]flow.Event, 0)
	for _, t := range r.GetTypes() {
		res, err := d.client.GetEventsForHeightRange(ctx, &access.GetEventsForHeightRangeRequest{
			StartHeight: r.Height,
			EndHeight:   r.Height,
			Type:        t,
		})
		if err != nil {
			return nil, err
		}
		for _, r := range res.GetResults() {
			for _, e := range r.Events {
				a := flow.Event{
					Type:             flow.EventType(e.Type),
					TransactionIndex: e.TransactionIndex,
					TransactionID:    flow.HashToID(e.TransactionId),
					EventIndex:       e.EventIndex,
					Payload:          e.Payload,
				}
				list = append(list, a)
			}
		}
	}

	data, err := d.codec.Marshal(list)
	if err != nil {
		return nil, fmt.Errorf("could not encode events: %w", err)
	}

	res := dps.GetEventsResponse{
		Height: r.Height,
		Types:  r.Types,
		Data:   data,
	}

	return &res, nil
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
