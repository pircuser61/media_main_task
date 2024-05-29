package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/pircuser61/media_main_task/config"
	"github.com/pircuser61/media_main_task/internal/transport/rest"
)

func main() {

	opts := slog.HandlerOptions{}
	opts.Level = config.GetLogLevel()

	logger := slog.New(slog.NewTextHandler(os.Stdout, &opts))
	logger.Info("start")
	logger.Info("LogLevel", slog.Any("level", opts.Level.Level()))
	logger.Info(config.GetAddr())

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	// если хочется в отдельном потоке запустить ...
	/*
		doneCh := make(chan struct{})
		go func() {
			defer func() {
				doneCh <- struct{}{}
			}()
			rest.RunSv(ctx, logger)
		}()
		<-doneCh
	*/

	rest.RunSv(ctx, logger)
	logger.Debug("done")
}
