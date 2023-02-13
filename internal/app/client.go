package app

import (
	"context"
	"github.com/dmitryburov/word-of-wisdom/config"
	"github.com/dmitryburov/word-of-wisdom/internal/usecase"
	"github.com/dmitryburov/word-of-wisdom/pkg/logger/zap"
	"github.com/dmitryburov/word-of-wisdom/pkg/pow/hashcash"
	"github.com/dmitryburov/word-of-wisdom/utils"
)

const (
	ErrClientFetch = "client error fetch"
)

// RunClient started client application
func RunClient(ctx context.Context, cfg *config.ClientConfig) {
	loggerProvider := zap.NewZapLogger(cfg.Logger)
	loggerProvider.InitLogger(utils.ApplicationClientName)

	powProvider, err := hashcash.NewPOW(cfg.Pow.Complexity)
	if err != nil {
		utils.FatalApplication(ErrPowInit, err)
	}

	client := usecase.NewClient(cfg, loggerProvider, powProvider)
	if err = client.Start(ctx, cfg.RequestCount); err != nil {
		utils.FatalApplication(ErrClientFetch, err)
	}
}
