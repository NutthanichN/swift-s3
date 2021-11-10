package server

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/rs/zerolog/log"
	"runtime"
)

type Config struct {
	LogPretty                           bool   `envconfig:"LOG_PRETTY" default:"true"`
	LogSampling                         bool   `envconfig:"LOG_SAMPLING" default:"false"`
	LogRequests                         bool   `envconfig:"LOG_REQUESTS" default:"false"`
	ServerAddress                       string `envconfig:"SERVER_ADDRESS" default:":8080"`
	SentryDSN                           string `envconfig:"SENTRY_DSN" default:""`
	TestStorageAccessKeyID              string `envconfig:"TEST_STORAGE_ACCESS_KEY_ID" default:""`
	TestStorageAccessKeySecret          string `envconfig:"TEST_STORAGE_ACCESS_KEY_SECRET" default:""`
	AuthTokenSecret                     string `envconfig:"AUTH_TOKEN_SECRET" default:""`
	StorageCredentialsBaseUrl           string `envconfig:"CRED_PROVIDER_BASE_URL" default:""`
	StorageCredentialsUrlPath           string `envconfig:"CRED_PROVIDER_ACCESS_KEY_PATH" default:""`
	StorageAccessProviderActivationPath string `envconfig:"CRED_PROVIDER_ACTIVATION_PATH" default:""`
	SwiftDomain                         string `envconfig:"SWIFT_DOMAIN"`
	SwiftAccessKey                      string `envconfig:"SWIFT_ACCESS_KEY"`
	SwiftUser                           string `envconfig:"SWIFT_USER"`
	SwiftAuthUrl                        string `envconfig:"SWIFT_AUTH_URL"`
}

var globalConfig = &Config{}

// LoadConfig will load config from env
func LoadConfig() (*Config, error) {
	// This error is ignored, since .env file may not exists,
	// and should not need to be created.
	_ = godotenv.Load(".env")

	var config = globalConfig
	err := envconfig.Process("STORAGE_", config)

	return config, err
}

// LogGeneralAppInfo will log basic information for the app
func LogGeneralAppInfo() {
	log.Info().
		Int("GOMAXPROCS", runtime.GOMAXPROCS(0)).
		Msg("Starting application")
}

func GetConfig() *Config {
	return globalConfig
}
