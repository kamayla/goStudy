package main

import (
	"context"
	"errors"
	"golang.org/x/sync/errgroup"
	"log"
	"net"
	"net/http"
)

type Server struct {
	srv *http.Server
	l   net.Listener
}

func NewServer(l net.Listener, mux http.Handler) *Server {
	return &Server{
		srv: &http.Server{Handler: mux},
		l:   l,
	}
}

func (s *Server) Run(ctx context.Context) error {
	// TODO: NotifyContextをやるとサーバーが落ちちゃうのであとで調査
	//ctx, stop := signal.NotifyContext(ctx)
	//defer stop()

	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		if err := s.srv.Serve(s.l); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Printf("failed to close: %+v", err)
			return err
		}
		return nil
	})

	<-ctx.Done()
	if err := s.srv.Shutdown(context.Background()); err != nil {
		log.Printf("failed to shutdown: %+v", err)
		return nil
	}
	// グレースフルシャットダウンの終了を待つ
	return eg.Wait()
}
