package client

import (
	"fmt"
	"os"
	"time"

	"github.com/rs/zerolog"
)

type (
	Spec struct {
		Token     string     `json:"token,omitempty"`
		Domain    string     `json:"domain,omitempty"`
		RateLimit *RateLimit `json:"rate_limit,omitempty"`
		Debug     bool       `json:"debug,omitempty"`
	}
	RateLimit struct {
		MaxBackoff time.Duration `json:"max_backoff,omitempty"`
		MaxRetries int32         `json:"max_retries,omitempty"`
	}
)

const (
	OktaAPIToken = "OKTA_API_TOKEN"
)

func (s *Spec) setDefaults(logger *zerolog.Logger) {
	const (
		minRetries = int32(2)
		minBackOff = 30 * time.Second
	)

	if s.RateLimit == nil {
		s.RateLimit = new(RateLimit)
	}

	if s.RateLimit.MaxRetries < minRetries {
		s.RateLimit.MaxRetries = minRetries
	}

	if s.RateLimit.MaxBackoff < minBackOff {
		s.RateLimit.MaxBackoff = minBackOff
	}

	if len(s.Token) == 0 {
		logger.Warn().Msgf("usage of %q environment variable value is deprecated and will be dropped in a future release", OktaAPIToken)
		s.Token = os.Getenv(OktaAPIToken)
	}
}

func (s *Spec) validate() error {
	if len(s.Token) == 0 {
		return fmt.Errorf("missing API token (should be set in the configuration or as %q environment variable)", OktaAPIToken)
	}

	const exampleDomain = "https://<CHANGE_THIS_TO_YOUR_OKTA_DOMAIN>.okta.com"
	switch s.Domain {
	case "", exampleDomain:
		return fmt.Errorf("missing \"domain\" in plugin configuration")
	}

	return nil
}
