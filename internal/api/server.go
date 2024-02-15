package api

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"sync"

	"github.com/samvimes01/go-rss/internal/api/middleware"
	"github.com/samvimes01/go-rss/internal/api/routes"
	"github.com/samvimes01/go-rss/internal/env"
)

func Run(config *env.Env, apiCfg *routes.APIConfig) error {
	ctx := *(apiCfg.GetCtx())

	mux := http.NewServeMux()
	routes.Setup(config, mux, apiCfg)
	srv := middleware.Cors(mux)

	httpServer := &http.Server{
		Addr:    net.JoinHostPort(config.Host, config.Port),
		Handler: srv,
	}

	go func() {
		log.Printf("Server listening on %s\n", httpServer.Addr)
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Fprintf(os.Stderr, "error listening and serving: %s\n", err)
		}
	}()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		<-ctx.Done()
		if err := httpServer.Shutdown(ctx); err != nil {
			fmt.Fprintf(os.Stderr, "error shutting down http server: %s\n", err)
		} else {
      fmt.Println("Server shut down due to SIGINT")
    }
	}()
	wg.Wait()
	return nil
}
