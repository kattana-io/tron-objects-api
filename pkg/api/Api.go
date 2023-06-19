package api

import (
	"github.com/kattana-io/tron-objects-api/pkg/url"
	"go.uber.org/zap"
)

type API struct {
	endpoint string
	log      *zap.Logger
	provider url.APIURLProvider
}

func NewAPI(nodeURL string, logger *zap.Logger, provider url.APIURLProvider) *API {
	return &API{
		endpoint: nodeURL,
		log:      logger,
		provider: provider,
	}
}

// "000000000000000000000000a614f803b6fd780986a42c78ec9c7f77e6ded13c" -> "a614f803b6fd780986a42c78ec9c7f77e6ded13c"
func TrimZeroes(address string) string {
	idx := 0
	for ; idx < len(address); idx++ {
		if address[idx] != '0' {
			break
		}
	}
	return address[idx:]
}
