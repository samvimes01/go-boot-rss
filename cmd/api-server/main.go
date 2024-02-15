package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/samvimes01/go-rss/internal/api"
	"github.com/samvimes01/go-rss/internal/api/routes"
	"github.com/samvimes01/go-rss/internal/env"
	rss_parser "github.com/samvimes01/go-rss/internal/rss-parser"
)

func main() {
	ctx := context.Background()
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	config := env.NewEnv()
	cfg := routes.NewAPIConfig(config, &ctx)

	go rss_parser.CrawlFeeds(config, cfg)

	if err := api.Run(config, cfg); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
