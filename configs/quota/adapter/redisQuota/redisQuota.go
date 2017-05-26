package redisQuota

import (
	"fmt"

	"../../quota"
)

type Adapter struct{}
type Processor struct{}

// this will likely need refinement in approach
func RegisterAdapter() quota.Adapter {
	return newAdapter()
}

// this is where adapter config would go
func newAdapter() *Adapter {
	return &Adapter{}
}

func (a *Adapter) NewProcessor(config quota.Config, templates []*quota.Template) (quota.Processor, error) {
	return &Processor{}, nil
}

func (Adapter) Close() error { return nil }

func (p *Processor) Process(instances []*quota.Instance) (QuotaResponse, error) {
	return QuotaResponse{}, nil
}

func (Processor) Close() error { return nil }
