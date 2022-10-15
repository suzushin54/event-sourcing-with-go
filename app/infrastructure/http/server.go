package http

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type Server struct {
	server *http.Server
}

func NewServer(mux http.Handler) *Server {
	return &Server{
		server: &http.Server{Handler: mux},
	}
}

func (s *Server) Run(ctx context.Context, port int) error {
	ctx, stop := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	defer stop()

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return fmt.Errorf("failed to build listener: %v", err)
	}

	url := fmt.Sprintf("http://%s", l.Addr().String())
	log.Printf("start with: %v", url)

	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(
		func() error {
			if err := s.server.Serve(l); err != nil &&
				err != http.ErrServerClosed { // 正常終了
				log.Printf("failed to close: %+v", err)
				return err
			}
			return nil
		},
	)

	// channel からの終了通知を待機する
	<-ctx.Done()
	if err := s.server.Shutdown(context.Background()); err != nil {
		log.Printf("failed to shutdown: %+v", err)
	}
	// goroutineの終了を待つ
	return eg.Wait()
}
