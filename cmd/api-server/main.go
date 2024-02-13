package main

import (
	"net/http"

	"github.com/samvimes01/go-rss/internal/api/middlewares"
	"github.com/samvimes01/go-rss/internal/api/routes"
	"github.com/samvimes01/go-rss/internal/env"
)

var envir *env.Env
var cfg *routes.APIConfig

func init() {
	envir = env.NewEnv()
	cfg = routes.NewAPIConfig(envir)
}

func main() {
	mux := http.NewServeMux()
	routes.Setup(envir, mux, cfg)
	corsHandler := middlewares.CorsMiddleware(mux)
	server := http.Server{
		Handler: corsHandler,
		Addr:    envir.Host + ":" + envir.Port,
	}

	server.ListenAndServe()
}
