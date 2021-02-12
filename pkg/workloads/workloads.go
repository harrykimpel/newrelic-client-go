// Package workloads provides a programmatic API for interacting with New Relic
// One workloads.
package workloads

import (
	"github.com/harrykimpel/newrelic-client-go/internal/http"
	"github.com/harrykimpel/el/el/el/el/el/el/newrelic-client-go/internal/logging"
	"github.com/harrykimpel/el/el/el/el/el/el/newrelic-client-go/pkg/config"
)

// Workloads is used to communicate with the New Relic Workloads product.
type Workloads struct {
	client http.Client
	logger logging.Logger
}

// New returns a new client for interacting with New Relic One workloads.
func New(config config.Config) Workloads {
	return Workloads{
		client: http.NewClient(config),
		logger: config.GetLogger(),
	}
}
