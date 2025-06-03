package rest

import (
	"github.com/kattana-io/tron-objects-api/pkg/client/rest"
	"go.uber.org/zap"
)

type API struct {
	endpoint string
	log      *zap.Logger
	provider rest.APIURLProvider
}

func NewAPI(nodeURL string, logger *zap.Logger, provider rest.APIURLProvider) *API {
	return &API{
		endpoint: nodeURL,
		log:      logger,
		provider: provider,
	}
}
