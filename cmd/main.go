package main

import (
	"context"
	"github.com/glamostoffer/ValinorGateway/internal/app"
	"github.com/glamostoffer/ValinorGateway/internal/config"
	"github.com/glamostoffer/ValinorGateway/utils/logger"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg := config.LoadConfig()
	lg := setupPrettySlog()

	lg.Info(
		"Config loaded",
		slog.Any("cfg", *cfg),
	)

	a := app.New(*cfg, lg)

	go func() {
		if err := a.Start(context.Background()); err != nil {
			panic(err.Error())
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<-stop

	stopCtx, stopCancel := context.WithTimeout(context.Background(), cfg.StopTimeout)
	defer stopCancel()

	if err := a.Stop(stopCtx); err != nil {
		panic(err.Error())
	}
}

func setupPrettySlog() *slog.Logger {
	opts := logger.HandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}

	handler := opts.NewPrettyHandler(os.Stdout)

	return slog.New(handler)
}
