package app

import (
	"Calendar/internal/config"
	"Calendar/internal/repository"
	"Calendar/internal/service"
	"Calendar/internal/transport"
	"Calendar/pkg/logger"
	"Calendar/pkg/postgres"
	"context"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

type App struct {
	SubscriptionServer *transport.CalendarServer
	cfg                *config.Config
	ctx                context.Context
	wg                 sync.WaitGroup
	cancel             context.CancelFunc
}

func New(ctx context.Context, cfg *config.Config) *App {
	db, err := postgres.New(cfg.Postgres)
	if err != nil {
		panic(err)
	}
	repo := repository.NewCalendarRepository(ctx, db)
	srv := service.NewCalendarService(ctx, repo)
	server := transport.NewCalendarServer(ctx, cfg, srv)
	return &App{
		SubscriptionServer: server,
		cfg:                cfg,
		ctx:                ctx,
	}
}

func (s *App) MustRun() {
	if err := s.Run(); err != nil {
		panic(err)
	}
}

func (s *App) Run() error {
	errCh := make(chan error, 1)
	s.wg.Add(1)
	go func() {
		logger.GetLoggerFromCtx(s.ctx).Info("Server started on address", zap.Any("address", s.cfg.Host+":"+s.cfg.Port))
		defer s.wg.Done()
		if err := s.SubscriptionServer.Run(); err != nil {
			errCh <- err
			s.cancel()
		}
	}()
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	select {
	case err := <-errCh:
		logger.GetLoggerFromCtx(s.ctx).Error("error running app", zap.Error(err))
		return err
	case <-s.ctx.Done():
		logger.GetLoggerFromCtx(s.ctx).Info("context done")
	}

	return nil
}
