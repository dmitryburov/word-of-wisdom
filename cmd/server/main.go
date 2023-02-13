package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/dmitryburov/word-of-wisdom/config"
	"github.com/dmitryburov/word-of-wisdom/utils"
)

const ErrConfigInit = "failed config initialization"

func main() {
	ctx, cancel := signal.NotifyContext(context.TODO(), os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-ctx.Done()
		cancel()
	}()

	cfg, err := config.NewConfig(ctx, config.ServerConfig{})
	if err != nil {
		utils.FatalApplication(ErrConfigInit, err)
	}

	app.RunServer(ctx, cfg)
}
