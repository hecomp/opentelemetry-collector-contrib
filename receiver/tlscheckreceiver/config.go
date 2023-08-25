package tlscheckreceiver

import (
	"fmt"
	"github.com/juju/errors"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/tlscheckreceiver/internal/configtls"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/tlscheckreceiver/internal/metadata"
	"go.opentelemetry.io/collector/receiver/scraperhelper"
	"net/url"
)

var (
	errInvalidEndpoint = errors.New(`"endpoint" must be in the form of <scheme>://<hostname>:<port>`)
	errInvalidCertPath = errors.New(`"local cert path invalid"`)
)

const defaultEndpoint = "http://localhost:433"

type Config struct {
	scraperhelper.ScraperControllerSettings `mapstructure:",squash"`
	configtls.TLSCertsClientSettings        `mapstructure:",squash"`

	metadata.MetricsBuilderConfig `mapstructure:",squash"`
}

// Validate validates the configuration by checking for missing or invalid fields
func (cfg *Config) Validate() error {
	_, err := url.Parse(cfg.Endpoint)
	if err != nil {
		return fmt.Errorf("invalid endpoint: '%s': %w", cfg.Endpoint, err)
	}

	if cfg.TLSSetting.TLSSetting.CertFile == "" {
		return errInvalidCertPath
	}

	return nil
}
