package main

import (
	"Calendar/internal/app"
	"Calendar/internal/config"
	"Calendar/pkg/logger"
	"context"
)

func main() {
	ctx := context.Background()
	ctx, err := logger.New(ctx)
	if err != nil {
		panic(err)
	}
	cfg, err := config.NewConfig()
	if err != nil {
		panic(err)
	}
	a := app.New(ctx, cfg)
	a.MustRun()
}
