package app

import (
	"context"
	"github.com/dmitryburov/word-of-wisdom/config"
	"github.com/dmitryburov/word-of-wisdom/internal/repository"
	"github.com/dmitryburov/word-of-wisdom/internal/usecase"
	"github.com/dmitryburov/word-of-wisdom/pkg/logger/zap"
	"github.com/dmitryburov/word-of-wisdom/pkg/pow/hashcash"
	"github.com/dmitryburov/word-of-wisdom/utils"
)

const (
	ErrPowInit   = "failed to initialize pow"
	ErrRunServer = "failed server run"
)

// RunServer started server application
func RunServer(ctx context.Context, cfg *config.ServerConfig) {
	loggerProvider := zap.NewZapLogger(cfg.Logger)
	loggerProvider.InitLogger(utils.ApplicationServerName)

	powProvider, err := hashcash.NewPOW(cfg.Pow.Complexity)
	if err != nil {
		utils.FatalApplication(ErrPowInit, err)
	}

	repo := repository.NewRepositories()
	server := usecase.NewServer(cfg, loggerProvider, powProvider, repo)

	if err = server.Run(ctx); err != nil {
		utils.FatalApplication(ErrRunServer, err)
	}
}
