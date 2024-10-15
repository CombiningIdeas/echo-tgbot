package app

import (
	"context"
	"os/signal"
	"syscall"
)

func Startup() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT,
		syscall.SIGTERM, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGABRT)

	defer stop()
	runServer(ctx)
}

func runServer(ctx context.Context) {

}

func printStartMessage(cfg *config.Config) {

}
