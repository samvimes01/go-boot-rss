package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/samvimes01/go-rss/internal/api/middleware"
	"github.com/samvimes01/go-rss/internal/api/routes"
	"github.com/samvimes01/go-rss/internal/env"
	rss_parser "github.com/samvimes01/go-rss/internal/rss-parser"
)

func main() {
	ctx := context.Background()
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	envir := env.NewEnv()
	cfg := routes.NewAPIConfig(envir, &ctx)

	mux := http.NewServeMux()

	routes.Setup(envir, mux, cfg)

	corsHandler := middleware.Cors(mux)

	server := http.Server{
		Handler: corsHandler,
		Addr:    envir.Host + ":" + envir.Port,
	}
	
	go rss_parser.CrawlFeeds(envir, cfg)

	log.Println("Server started on: ", server.Addr)
	
	if err := server.ListenAndServe(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
