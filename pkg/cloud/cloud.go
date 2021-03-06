package cloud

import (
	"github.com/harrykimpel/newrelic-client-go/internal/http"
	"github.com/harrykimpel/newrelic-client-go/internal/logging"
	"github.com/harrykimpel/newrelic-client-go/pkg/config"
)

type Cloud struct {
	client http.Client
	config config.Config
	logger logging.Logger
	pager  http.Pager
}

func New(config config.Config) Cloud {

	client := http.NewClient(config)
	client.SetAuthStrategy(&http.PersonalAPIKeyCapableV2Authorizer{})

	pkg := Cloud{
		client: client,
		config: config,
		logger: config.GetLogger(),
		pager:  &http.LinkHeaderPager{},
	}

	return pkg
}
