package config

import (
	"context"
	"github.com/sethvargo/go-envconfig"
)

// ServerConfig server configuration
type ServerConfig struct{}

// ClientConfig client configuration
type ClientConfig struct{}

// NewConfig generic for creates a new client or server config
func NewConfig[C any](ctx context.Context, config C) (*C, error) {
	if err := envconfig.Process(ctx, &config); err != nil {
		return nil, err
	}

	return &config, nil
}
