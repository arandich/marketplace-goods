package main

import (
	"context"
	"github.com/arandich/marketplace-goods/internal/config"
	logs "github.com/arandich/marketplace-sdk/logger"
	"log"
	"math/rand"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg := config.Config{}
	err := config.New(ctx, &cfg)
	if err != nil {
		log.Fatalln(err)
	}

	logger := logs.NewLogger(cfg.Logger.Level).With().
		CallerWithSkipFrameCount(2).
		Str("app", cfg.App.Name).
		Logger()

	logger.Info().Msg(cfg.App.Name + " started the app")

	// Random
	rand.NewSource(time.Now().UnixNano())

	// Run App
	runApp(logger.WithContext(ctx), cfg)
}
