package proxy

import "github.com/onflow/flow-dps/api/dps"

func NewDPSService() *DPSService {
	return &DPSService{}
}

type DPSService struct {
	dps.APIServer
}
