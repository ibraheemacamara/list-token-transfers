package config

import (
	"errors"
	"os"
	"strings"
)

// Config is a struct representing the application's configuration.
type Config struct {
	hostname           string
	port               string
	logLevel           string
	env                string
	erc20TransferTopic string
	rpcProvider        string
	allowedMethods     []string
}

var cfg Config

// LoadConfig loads configuration settings from environment variables.
func LoadConfig() error {
	erc20TransferTopic := os.Getenv("ERC20_TRANSFER_TOPIC")
	rpcProvider := os.Getenv("RPC_PROVIDER")
	allowedMethodString := os.Getenv("ALLOWED_METHOD")
	allowedMethod := strings.Split(allowedMethodString, ",")
	if erc20TransferTopic == "" || rpcProvider == "" || len(allowedMethod) < 1 {
		return errors.New("config must be set")
	}

	logLevel := os.Getenv("LOG_LEVEL")
	if logLevel == "" {
		logLevel = "INFO"
	}

	hostname := os.Getenv("HOSTNAME")
	if hostname == "" {
		hostname = "localhost"
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	env := os.Getenv("ENV")
	if port == "" {
		env = "dev"
	}

	cfg = Config{
		erc20TransferTopic: erc20TransferTopic,
		rpcProvider:        rpcProvider,
		hostname:           hostname,
		port:               port,
		logLevel:           logLevel,
		env:                env,
		allowedMethods:     allowedMethod,
	}

	return nil
}

// GetConfig returns the loaded Config instance.
func GetConfig() Config {
	return cfg
}

// Hostname returns the hostname for the configuration.
func (c Config) Hostname() string {
	return c.hostname
}

// Port returns the application's port for the configuration.
func (c Config) Port() string {
	return c.port
}

// LogLevel returns the logging level for the configuration.
func (c Config) LogLevel() string {
	return c.logLevel
}

// ENV returns the env for the configuration.
func (c Config) ENV() string {
	return c.env
}

// Erc20TransferTopic returns the Erc20TransferTopic for the configuration.
func (c Config) Erc20TransferTopic() string {
	return c.erc20TransferTopic
}

// rpcProvider returns the rpcProvider for the configuration.
func (c Config) RpcProvider() string {
	return c.rpcProvider
}

// AllowedMethods returns the AllowedMethod for the configuration.
func (c Config) AllowedMethods() []string {
	return c.allowedMethods
}
