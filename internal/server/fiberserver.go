package server

import (
	"context"
	"errors"
	"github.com/glamostoffer/ValinorGateway/internal/delivery"
	"github.com/glamostoffer/ValinorGateway/internal/middleware"
	"github.com/glamostoffer/ValinorGateway/internal/usecase"
	"github.com/glamostoffer/ValinorProtos/auth"
	"github.com/glamostoffer/ValinorProtos/chat"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log/slog"
	"time"
)

type FiberServer struct {
	cfg     Config
	fb      *fiber.App
	mw      middleware.Middleware
	uc      usecase.UseCase
	handler *delivery.Handler
	logger  *slog.Logger
	auth    *auth.Connector
	chat    *chat.Connector
}

func New(
	cfg Config,
	mw middleware.Middleware,
	uc usecase.UseCase,
	handler *delivery.Handler,
	logger *slog.Logger,
	auth *auth.Connector,
	chat *chat.Connector,
) *FiberServer {
	return &FiberServer{
		cfg: cfg,
		fb: fiber.New(
			fiber.Config{
				DisableStartupMessage: true,
			},
		),
		mw:      mw,
		uc:      uc,
		handler: handler,
		logger:  logger,
		auth:    auth,
		chat:    chat,
	}
}

func (f *FiberServer) Start(ctx context.Context) error {
	log := f.logger.With(slog.String("op", "server.Start"))

	f.fb.Use(cors.New(
		cors.Config{
			AllowOrigins:     f.cfg.AllowOrigins,
			AllowCredentials: true,
			AllowHeaders:     f.cfg.AllowHeaders,
			ExposeHeaders:    f.cfg.ExposeHeaders,
		},
	))

	route := f.fb.Group("/")
	delivery.MapRoutes(route, f.mw, f.handler)

	go func() {
		if err := f.fb.Listen(f.cfg.Host + ":" + f.cfg.Port); err != nil {
			log.Error(err.Error())
			panic(err.Error())
		}
	}()

	return nil
}

func (f *FiberServer) Stop(_ context.Context) error {
	log := f.logger.With(slog.String("op", "server.Stop"))

	okCh, errCh := make(chan struct{}), make(chan error)
	go func() {
		if err := f.fb.Shutdown(); err != nil {
			errCh <- err
		}
		okCh <- struct{}{}
	}()

	select {
	case <-okCh:
		return nil
	case err := <-errCh:
		log.Error(err.Error())
		return err
	case <-time.After(f.cfg.StopTimeout):
		err := errors.New("stop timeout")
		log.Error(err.Error())
		return err
	}
}
