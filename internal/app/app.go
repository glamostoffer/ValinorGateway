package app

import (
	"context"
	"errors"
	"fmt"
	"github.com/glamostoffer/ValinorGateway/internal/config"
	"github.com/glamostoffer/ValinorGateway/internal/delivery"
	"github.com/glamostoffer/ValinorGateway/internal/middleware"
	"github.com/glamostoffer/ValinorGateway/internal/server"
	"github.com/glamostoffer/ValinorGateway/internal/usecase"
	"github.com/glamostoffer/ValinorGateway/pkg/consts"
	authclient "github.com/glamostoffer/ValinorProtos/auth"
	chatclient "github.com/glamostoffer/ValinorProtos/chat"
	"log/slog"
)

type (
	App struct {
		cfg        config.Config
		components []component
		log        *slog.Logger
	}
	component struct {
		Service Lifecycle
		Name    string
	}
	Lifecycle interface {
		Start(ctx context.Context) error
		Stop(ctx context.Context) error
	}
)

func New(cfg config.Config, logger *slog.Logger) *App {
	return &App{
		cfg: cfg,
		log: logger,
	}
}

func (a *App) Start(ctx context.Context) error {
	log := a.log.With(slog.String("op", "app.Start"))

	auth := authclient.New(a.cfg.AuthCfg)
	chat := chatclient.New(a.cfg.ChatCfg)

	useCase := usecase.New(auth, chat)

	handler := delivery.New(a.cfg.RouteConfig, useCase)

	mw := middleware.New(auth)

	serv := server.New(
		a.cfg.HTTPServer,
		mw,
		useCase,
		handler,
		a.log,
		auth,
	)

	a.components = append(a.components,
		component{auth, "Auth Client"},
		component{serv, "HTTP Server"},
	)

	okChan := make(chan struct{})
	errChan := make(chan error)

	go func() {
		var err error
		for _, c := range a.components {
			log.Info(consts.FmtStarting, slog.Any("name", c.Name))

			err = c.Service.Start(context.Background())
			if err != nil {
				log.Error(consts.FmtErrOnStarting, c.Name, err.Error())
				errChan <- errors.New(
					fmt.Sprintf("%s %s: %s", consts.FmtCannotStart, c.Name, err.Error()),
				)

				return
			}
		}
		okChan <- struct{}{}
	}()

	select {
	case <-ctx.Done():
		return errors.New("start timeout")
	case err := <-errChan:
		return err
	case <-okChan:
		log.Info("application started!")
		return nil
	}
}

func (a *App) Stop(ctx context.Context) error {
	log := a.log.With(slog.String("op", "app.Stop"))
	okChan := make(chan struct{})
	errChan := make(chan error)

	go func() {
		var err error
		for i := len(a.components) - 1; i >= 0; i-- {
			log.Info(
				consts.FmtStopping,
				slog.Any("name", a.components[i].Name),
			)

			err = a.components[i].Service.Stop(ctx)
			if err != nil {
				log.Error(consts.FmtErrOnStopping, a.components[i].Name, err.Error())
				errChan <- errors.New(
					fmt.Sprintf(
						"%s %s: %s",
						consts.FmtCannotStop,
						a.components[i].Name,
						err.Error(),
					),
				)

				return
			}
		}
		okChan <- struct{}{}
	}()

	select {
	case <-ctx.Done():
		return errors.New("stop timeout")
	case err := <-errChan:
		return err
	case <-okChan:
		log.Info("application stopped!")
		return nil
	}
}
