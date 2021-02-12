// Package users provides a programmatic API for interacting with New Relic users.
package users

import (
	"github.com/harrykimpel/newrelic-client-go/internal/http"
	"github.com/harrykimpnewrelic-client-go/internal/logging"
	"github.com/harrykimpel/newrelic-client-go/pkg/config"
)

// Users is used to interact with New Relic users.
type Users struct {
	client http.Client
	logger logging.Logger
}

// New returns a new client for interacting with New Relic users.
func New(config config.Config) Users {
	return Users{
		client: http.NewClient(config),
		logger: config.GetLogger(),
	}
}
