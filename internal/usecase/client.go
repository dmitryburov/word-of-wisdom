package usecase

import (
	"context"
	"fmt"
	"net"

	"github.com/dmitryburov/word-of-wisdom/config"
	"github.com/dmitryburov/word-of-wisdom/pkg/logger"
	"github.com/dmitryburov/word-of-wisdom/pkg/pow"
	"github.com/dmitryburov/word-of-wisdom/utils"
)

// client represents a client
type client struct {
	conf   *config.ClientConfig
	logger logger.Logger
	solver pow.Solver
}

// NewClient creates a new client
func NewClient(
	conf *config.ClientConfig,
	logger logger.Logger,
	solver pow.Solver) *client {
	return &client{
		conf:   conf,
		logger: logger,
		solver: solver,
	}
}

// Start started fetch
func (c *client) Start(ctx context.Context, count int) error {
	for i := 0; i < count; i++ {
		if ctx.Err() != nil {
			break
		}

		q, err := c.GetQuote(ctx)
		if err != nil {
			c.logger.Error("failed to get quote", err.Error())
		} else {
			c.logger.Info(string(q))
		}
	}

	return nil
}

// GetQuote returns a quote from the server
func (c *client) GetQuote(ctx context.Context) ([]byte, error) {
	var dialer net.Dialer
	conn, err := dialer.DialContext(ctx, "tcp", c.conf.ServerAddr)
	if err != nil {
		return nil, fmt.Errorf("failed to dial: %w", err)
	}
	defer func() {
		if err := conn.Close(); err != nil {
			c.logger.Error("failed to close connection", err.Error())
		}
	}()

	// challenge request
	if err := utils.WriteMessage(conn, []byte("challenge")); err != nil {
		return nil, fmt.Errorf("send challenge request err: %w", err)
	}

	// receive challenge
	challenge, err := utils.ReadMessage(conn)
	if err != nil {
		return nil, fmt.Errorf("receive challenge err: %w", err)
	}

	// send solution
	solution := c.solver.Solve(challenge)
	if err := utils.WriteMessage(conn, solution); err != nil {
		return nil, fmt.Errorf("send solution err: %w", err)
	}

	// receive quote
	quote, err := utils.ReadMessage(conn)
	if err != nil {
		return nil, fmt.Errorf("receive quote err: %w", err)
	}

	return quote, nil
}
