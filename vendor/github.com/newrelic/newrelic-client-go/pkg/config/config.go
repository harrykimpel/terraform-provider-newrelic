package config

import (
	"net/http"
	"time"

	"github.com/newrelic/newrelic-client-go/internal/logging"
	"github.com/newrelic/newrelic-client-go/internal/version"
)

// Config contains all the configuration data for the API Client.
type Config struct {
	// PersonalAPIKey to authenticate API requests
	// see: https://docs.newrelic.com/docs/apis/get-started/intro-apis/types-new-relic-api-keys#personal-api-key
	PersonalAPIKey string

	// AdminAPIKey to authenticate API requests
	// Note this will be deprecated in the future!
	// see: https://docs.newrelic.com/docs/apis/get-started/intro-apis/types-new-relic-api-keys#admin
	AdminAPIKey string

	// Region of the New Relic platform to use
	// Valid values are: US, EU
	Region string

	// HTTP
	Timeout               *time.Duration
	HTTPTransport         http.RoundTripper
	UserAgent             string
	BaseURL               string
	SyntheticsBaseURL     string
	InfrastructureBaseURL string
	NerdGraphBaseURL      string
	ServiceName           string

	// LogLevel can be one of the following values:
	// "panic", "fatal", "error", "warn", "info", "debug", "trace"
	LogLevel string
	LogJSON  bool
	Logger   logging.Logger
}

// GetLogger returns a logger instance based on the config values.
func (c *Config) GetLogger() logging.Logger {
	if c.Logger != nil {
		return c.Logger
	}

	l := logging.NewStructuredLogger().
		SetDefaultFields(map[string]string{"newrelic-client-go": version.Version}).
		LogJSON(c.LogJSON).
		SetLogLevel(c.LogLevel)

	c.Logger = l
	return l
}
