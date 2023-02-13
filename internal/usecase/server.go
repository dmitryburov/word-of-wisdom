package usecase

import (
	"context"
	"errors"
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/dmitryburov/word-of-wisdom/config"
	"github.com/dmitryburov/word-of-wisdom/internal/repository"
	"github.com/dmitryburov/word-of-wisdom/pkg/logger"
	"github.com/dmitryburov/word-of-wisdom/pkg/pow"
	"github.com/dmitryburov/word-of-wisdom/utils"
)

// server represents a server
type server struct {
	cfg      *config.ServerConfig
	logger   logger.Logger
	verifier pow.Verifier
	repo     repository.Repositories
	listener net.Listener
	wg       sync.WaitGroup
	cancel   context.CancelFunc
}

// NewServer creates a new server
func NewServer(
	cfg *config.ServerConfig,
	logger logger.Logger,
	verifier pow.Verifier,
	repo repository.Repositories,
) *server {
	return &server{
		cfg:      cfg,
		logger:   logger,
		verifier: verifier,
		repo:     repo,
	}
}

// Run starts the server
func (s *server) Run(ctx context.Context) (err error) {
	ctx, s.cancel = context.WithCancel(ctx)
	defer s.cancel()

	lc := net.ListenConfig{
		KeepAlive: s.cfg.KeepAlive,
	}
	s.listener, err = lc.Listen(ctx, "tcp", s.cfg.Addr)
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}

	s.logger.Info(fmt.Sprintf("server started on port %s", s.listener.Addr().String()))

	s.wg.Add(1)
	go s.serve(ctx)
	s.wg.Wait()

	s.logger.Info("server stopped")

	return nil
}

// Stop stops the server
func (s *server) Stop() {
	s.cancel()
}

func (s *server) serve(ctx context.Context) {
	defer s.wg.Done()

	go func() {
		<-ctx.Done()
		err := s.listener.Close()
		if err != nil && !errors.Is(err, net.ErrClosed) {
			s.logger.Error("failed to close listener: ", err.Error())
		}
	}()

	for {
		conn, err := s.listener.Accept()
		if errors.Is(err, net.ErrClosed) {
			s.logger.Debug("listener closed")
			return
		} else if err != nil {
			s.logger.Error("failed to accept connection: ", err.Error())
			continue
		}

		s.wg.Add(1)
		go func(conn net.Conn) {
			defer s.wg.Done()

			if err = s.handle(conn); err != nil {
				s.logger.Error("handle error: ", err.Error())
			}
		}(conn)
	}
}

func (s *server) handle(conn net.Conn) error {
	defer conn.Close()
	_ = conn.SetDeadline(time.Now().Add(s.cfg.Deadline))

	// receive challenge request
	if _, err := utils.ReadMessage(conn); err != nil {
		return fmt.Errorf("read message err: %w", err)
	}

	// send challenge
	challenge := s.verifier.Challenge()
	if err := utils.WriteMessage(conn, challenge); err != nil {
		return fmt.Errorf("send challenge err: %w", err)
	}

	// receive solution
	solution, err := utils.ReadMessage(conn)
	if err != nil {
		return fmt.Errorf("receive proof err: %w", err)
	}

	// verify solution
	if err = s.verifier.Verify(challenge, solution); err != nil {
		return fmt.Errorf("invalid verify: %w", err)
	}

	// send result
	quote, err := s.repo.Quotes.GetQuote()
	if err != nil {
		return fmt.Errorf("get quote err: %w", err)
	}

	if err = utils.WriteMessage(conn, []byte(quote)); err != nil {
		return fmt.Errorf("send quote err: %w", err)
	}

	return nil
}
